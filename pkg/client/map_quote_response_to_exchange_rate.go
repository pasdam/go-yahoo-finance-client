package client

import (
	"github.com/pasdam/go-yahoo-finance-client/internal/pkg/requests/quote"
)

func mapQuoteResponseToQuote(value *quote.Response) *ExchangeRate {
	return &ExchangeRate{
		Buy:       value.Response.Results[0].RegularMarketPrice,
		Sell:      value.Response.Results[0].RegularMarketPrice,
		Timestamp: value.Response.Results[0].RegularMarketTime,
	}
}
