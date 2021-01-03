package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pasdam/go-rest-util/pkg/restutil"
	"github.com/pasdam/go-yahoo-finance-client/internal/pkg/requests/chart"
	"github.com/pasdam/go-yahoo-finance-client/internal/pkg/requests/quote"
	"github.com/pasdam/mockit/matchers/argument"
	"github.com/pasdam/mockit/mockit"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient()

	assert.Equal(t, "https://query1.finance.yahoo.com", c.baseURL.String())
}

func TestNewClientWithURL(t *testing.T) {
	type args struct {
		u string
	}
	tests := []struct {
		args    args
		wantErr error
	}{
		{
			args: args{
				u: "localhost:123",
			},
		},
		{
			args: args{
				u: "\t invalid url",
			},
			wantErr: errors.New("parse \"\\t invalid url\": net/url: invalid control character in URL"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.u, func(t *testing.T) {
			got, err := NewClientWithURL(tt.args.u)

			if tt.wantErr != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tt.args.u, got.baseURL.String())
			}
		})
	}
}

func TestClient_Quotes_Real(t *testing.T) {
	t.SkipNow()
	c := NewClient()
	quotes, err := c.Quotes("USD", "EUR", 1588766400, 1588767300)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(quotes))

	expected := []*PriceQuote{
		{
			Timestamp: 1588766400,
			Open:      0.9247000217437744,
			High:      0.9251000285148621,
			Low:       0.9243000149726868,
			Close:     0.925000011920929,
			Volume:    0,
		},
		{
			Timestamp: 1588767300,
			Open:      0,
			High:      0,
			Low:       0,
			Close:     0,
			Volume:    0,
		},
	}
	assert.Equal(t, expected, quotes)
}

func TestClient_Quotes(t *testing.T) {
	type mocks struct {
		getBody *chart.QuotesResponseContent
		getErr  error
		query   string
	}
	type args struct {
		baseCurrency  string
		quoteCurrency string
		fromTimestamp uint64
		toTimestamp   uint64
	}
	tests := []struct {
		name    string
		mocks   mocks
		args    args
		want    []*PriceQuote
		wantErr error
	}{
		{
			name: "Should return parsed response when the request is successfull",
			mocks: mocks{
				getBody: &chart.QuotesResponseContent{
					Chart: chart.QuotesChart{
						Results: []chart.Result{
							{
								Meta:       chart.Meta{},
								Timestamps: []uint64{1000, 2000},
								Indicators: chart.Indicators{
									Quotes: []*chart.IndicatorQuote{
										{
											Volume: []float64{10, 100},
											Open:   []float64{20, 200},
											High:   []float64{30, 300},
											Low:    []float64{40, 400},
											Close:  []float64{50, 500},
										},
									},
								},
							},
						},
					},
				},
				query: "/EUR=X?interval=15m&period1=100&period2=200&symbol=EUR%3DX",
			},
			args: args{
				baseCurrency:  "USD",
				quoteCurrency: "EUR",
				fromTimestamp: 100,
				toTimestamp:   200,
			},
			want: []*PriceQuote{
				{
					Volume:    10,
					Open:      20,
					High:      30,
					Low:       40,
					Close:     50,
					Timestamp: 1000,
				},
				{
					Volume:    100,
					Open:      200,
					High:      300,
					Low:       400,
					Close:     500,
					Timestamp: 2000,
				},
			},
			wantErr: nil,
		},
		{
			name: "Should propagate error if restutil.GetJSON raises it",
			mocks: mocks{
				getErr: errors.New("some-get-json-error"),
				query:  "/AUD=X?interval=15m&period1=300&period2=400&symbol=AUD%3DX",
			},
			args: args{
				baseCurrency:  "USD",
				quoteCurrency: "AUD",
				fromTimestamp: 300,
				toTimestamp:   400,
			},
			want:    nil,
			wantErr: errors.New("some-get-json-error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost/v8/finance/chart%s", tt.mocks.query)
			if tt.mocks.getBody != nil {
				mock := mockit.MockFunc(t, restutil.Get)
				jsonBody, _ := json.Marshal(tt.mocks.getBody)
				mock.With(url, nil).Return(jsonBody, nil)
				mock.With(argument.Any, argument.Any).Return(nil, errors.New("Unexpected arguments"))
			}
			if tt.mocks.getErr != nil {
				mock := mockit.MockFunc(t, restutil.GetJSON)
				mock.With(url, nil, argument.Any).Return(tt.mocks.getErr)
				mock.With(argument.Any, argument.Any, argument.Any).Return(errors.New("Unexpected arguments"))
			}
			c, err := NewClientWithURL("http://localhost")
			assert.Nil(t, err)

			got, err := c.Quotes(tt.args.baseCurrency, tt.args.quoteCurrency, tt.args.fromTimestamp, tt.args.toTimestamp)

			if tt.wantErr == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
			if tt.want == nil {
				assert.Nil(t, got)
			} else {
				assert.NotNil(t, got)
				for i := 0; i < len(tt.want); i++ {
					assert.Equal(t, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestClient_CurrentRate(t *testing.T) {
	type mocks struct {
		getBody *quote.Response
		getErr  error
		query   string
	}
	type args struct {
		baseCurrency  string
		quoteCurrency string
	}
	tests := []struct {
		name    string
		mocks   mocks
		args    args
		want    *ExchangeRate
		wantErr error
	}{
		{
			name: "Should return parsed response when the request is successfull",
			mocks: mocks{
				getBody: &quote.Response{
					Response: &quote.QuoteResponse{
						Results: []quote.Result{
							{
								RegularMarketPrice: 10.1,
								RegularMarketTime:  100,
							},
						},
					},
				},
				query: "?&symbols=EUR=X&fields=extendedMarketChange,extendedMarketChangePercent,extendedMarketPrice,extendedMarketTime,regularMarketChange,regularMarketChangePercent,regularMarketPrice,regularMarketTime,circulatingSupply,ask,askSize,bid,bidSize,dayHigh,dayLow,regularMarketDayHigh,regularMarketDayLow,regularMarketVolume,volume",
			},
			args: args{
				baseCurrency:  "USD",
				quoteCurrency: "EUR",
			},
			want: &ExchangeRate{
				Buy:       10.1,
				Sell:      10.1,
				Timestamp: 100,
			},
			wantErr: nil,
		},
		{
			name: "Should propagate error if restutil.GetJSON raises it",
			mocks: mocks{
				getErr: errors.New("some-get-json-error"),
				query:  "?&symbols=AUD=X&fields=extendedMarketChange,extendedMarketChangePercent,extendedMarketPrice,extendedMarketTime,regularMarketChange,regularMarketChangePercent,regularMarketPrice,regularMarketTime,circulatingSupply,ask,askSize,bid,bidSize,dayHigh,dayLow,regularMarketDayHigh,regularMarketDayLow,regularMarketVolume,volume",
			},
			args: args{
				baseCurrency:  "USD",
				quoteCurrency: "AUD",
			},
			want:    nil,
			wantErr: errors.New("some-get-json-error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost/v7/finance/quote%s", tt.mocks.query)
			if tt.mocks.getBody != nil {
				mock := mockit.MockFunc(t, restutil.Get)
				jsonBody, _ := json.Marshal(tt.mocks.getBody)
				mock.With(url, nil).Return(jsonBody, nil)
				mock.With(argument.Any, argument.Any).Return(nil, errors.New("Unexpected arguments"))
			}
			if tt.mocks.getErr != nil {
				mock := mockit.MockFunc(t, restutil.GetJSON)
				mock.With(url, nil, argument.Any).Return(tt.mocks.getErr)
				mock.With(argument.Any, argument.Any, argument.Any).Return(errors.New("Unexpected arguments"))
			}
			c, err := NewClientWithURL("http://localhost")
			assert.Nil(t, err)

			got, err := c.CurrentRate(tt.args.baseCurrency, tt.args.quoteCurrency)

			if tt.wantErr == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestClient_IT_CurrentRate_ShouldParseValueFromServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"quoteResponse\":{\"result\":[{\"language\":\"en-US\",\"region\":\"US\",\"quoteType\":\"CURRENCY\",\"quoteSourceName\":\"Delayed Quote\",\"triggerable\":true,\"currency\":\"USD\",\"regularMarketChange\":0.0,\"regularMarketChangePercent\":-0.0,\"regularMarketPrice\":1.223092,\"regularMarketDayHigh\":1.2316787,\"regularMarketDayLow\":1.2214487,\"regularMarketPreviousClose\":1.223092,\"bid\":1.223092,\"ask\":1.2223445,\"fiftyTwoWeekLow\":1.2214487,\"fiftyTwoWeekHigh\":1.2316787,\"regularMarketTime\":1609453789,\"regularMarketDayRange\":\"1.2214487 - 1.2316787\",\"regularMarketVolume\":0,\"bidSize\":0,\"askSize\":0,\"fullExchangeName\":\"CCY\",\"fiftyTwoWeekLowChange\":0.005699992,\"fiftyTwoWeekLowChangePercent\":0.0046665836,\"fiftyTwoWeekRange\":\"1.2214487 - 1.2316787\",\"fiftyTwoWeekHighChange\":-0.0011000037,\"fiftyTwoWeekHighChangePercent\":-8.9309306E-4,\"sourceInterval\":15,\"exchangeDataDelayedBy\":0,\"firstTradeDateMilliseconds\":1070236800000,\"priceHint\":4,\"tradeable\":false,\"marketState\":\"CLOSED\",\"exchange\":\"CCY\",\"exchangeTimezoneName\":\"Europe/London\",\"exchangeTimezoneShortName\":\"GMT\",\"gmtOffSetMilliseconds\":0,\"market\":\"ccy_market\",\"esgPopulated\":false,\"symbol\":\"EURUSD=X\"}],\"error\":null}}")
	}))
	defer ts.Close()

	client, err := NewClientWithURL(ts.URL)
	assert.Nil(t, err)

	rate, err := client.CurrentRate("EUR", "USD")
	assert.Nil(t, err)

	assert.Equal(t, &ExchangeRate{
		Buy:       1.223092,
		Sell:      1.223092,
		Timestamp: 1609453789,
	}, rate)
}
