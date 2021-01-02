package client

import (
	"reflect"
	"testing"

	"github.com/pasdam/go-yahoo-finance-client/internal/pkg/requests/quote"
)

func Test_mapQuoteResponseToQuote(t *testing.T) {
	type args struct {
		value *quote.Response
	}
	tests := []struct {
		name string
		args args
		want *ExchangeRate
	}{
		{
			name: "Should correctly map a value",
			args: args{
				value: &quote.Response{
					Response: &quote.QuoteResponse{
						Results: []quote.Result{
							{
								RegularMarketPrice: 10.1,
								RegularMarketTime:  100,
							},
						},
					},
				},
			},
			want: &ExchangeRate{
				Buy:       10.1,
				Sell:      10.1,
				Timestamp: 100,
			},
		},
		{
			name: "Should correctly map a different value",
			args: args{
				value: &quote.Response{
					Response: &quote.QuoteResponse{
						Results: []quote.Result{
							{
								RegularMarketPrice: 20.1,
								RegularMarketTime:  200,
							},
						},
					},
				},
			},
			want: &ExchangeRate{
				Buy:       20.1,
				Sell:      20.1,
				Timestamp: 200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapQuoteResponseToQuote(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapQuoteResponseToQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}
