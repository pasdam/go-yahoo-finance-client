package quote

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Quote_ShouldParseQuoteJSON(t *testing.T) {
	jsonBody := "{\"quoteResponse\":{\"result\":[{\"language\":\"en-US\",\"region\":\"US\",\"quoteType\":\"CURRENCY\",\"quoteSourceName\":\"Delayed Quote\",\"triggerable\":true,\"currency\":\"USD\",\"regularMarketChange\":0.0,\"regularMarketChangePercent\":0.0,\"regularMarketPrice\":1.223092,\"regularMarketDayHigh\":1.2316787,\"regularMarketDayLow\":1.2214487,\"regularMarketPreviousClose\":1.223092,\"bid\":1.223092,\"ask\":1.2223445,\"fiftyTwoWeekLow\":1.2214487,\"fiftyTwoWeekHigh\":1.2316787,\"regularMarketTime\":1609453789,\"regularMarketDayRange\":\"1.2214487 - 1.2316787\",\"regularMarketVolume\":0,\"bidSize\":0,\"askSize\":0,\"fullExchangeName\":\"CCY\",\"fiftyTwoWeekLowChange\":0.005699992,\"fiftyTwoWeekLowChangePercent\":0.0046665836,\"fiftyTwoWeekRange\":\"1.2214487 - 1.2316787\",\"fiftyTwoWeekHighChange\":-0.0011000037,\"fiftyTwoWeekHighChangePercent\":-8.9309306E-4,\"sourceInterval\":15,\"exchangeDataDelayedBy\":0,\"firstTradeDateMilliseconds\":1070236800000,\"priceHint\":4,\"tradeable\":false,\"marketState\":\"CLOSED\",\"exchange\":\"CCY\",\"exchangeTimezoneName\":\"Europe/London\",\"exchangeTimezoneShortName\":\"GMT\",\"gmtOffSetMilliseconds\":0,\"market\":\"ccy_market\",\"esgPopulated\":false,\"symbol\":\"EURUSD=X\"}],\"error\":null}}"
	resp := &Response{}

	json.Unmarshal([]byte(jsonBody), resp)

	assert.Equal(t, &Response{
		Response: &QuoteResponse{
			Results: []Result{
				{
					EsgPopulated:                  false,
					Exchange:                      "CCY",
					ExchangeDataDelayedBy:         0,
					ExchangeTimezoneName:          "Europe/London",
					ExchangeTimezoneShortName:     "GMT",
					FiftyTwoWeekHigh:              1.2316787,
					FiftyTwoWeekHighChange:        -0.0011000037,
					FiftyTwoWeekHighChangePercent: -8.9309306e-4,
					FiftyTwoWeekLow:               1.2214487,
					FiftyTwoWeekLowChange:         0.005699992,
					FiftyTwoWeekLowChangePercent:  0.0046665836,
					FiftyTwoWeekRange:             "1.2214487 - 1.2316787",
					FirstTradeDateMilliseconds:    1070236800000,
					FullExchangeName:              "CCY",
					GmtOffSetMilliseconds:         0,
					Language:                      "en-US",
					Market:                        "ccy_market",
					MarketState:                   "CLOSED",
					QuoteSourceName:               "Delayed Quote",
					QuoteType:                     "CURRENCY",
					Region:                        "US",
					RegularMarketChange:           0.0,
					RegularMarketChangePercent:    -0.0,
					RegularMarketDayHigh:          1.2316787,
					RegularMarketDayLow:           1.2214487,
					RegularMarketDayRange:         "1.2214487 - 1.2316787",
					RegularMarketPreviousClose:    1.223092,
					RegularMarketPrice:            1.223092,
					RegularMarketTime:             1609453789,
					RegularMarketVolume:           0,
					SourceInterval:                15,
					Symbol:                        "EURUSD=X",
					Tradeable:                     false,
					Triggerable:                   true,
				},
			},
		},
	}, resp)
}
