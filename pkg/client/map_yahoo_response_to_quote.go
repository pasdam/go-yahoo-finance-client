package client

import "github.com/pasdam/go-yahoo-finance-client/internal/pkg/requests/chart"

func mapYahooResponseToQuote(content *chart.QuotesResponseContent) (quotes []*PriceQuote) {
	timestamps := content.Chart.Results[0].Timestamps
	volumes := content.Chart.Results[0].Indicators.Quotes[0].Volume
	open := content.Chart.Results[0].Indicators.Quotes[0].Open
	high := content.Chart.Results[0].Indicators.Quotes[0].High
	low := content.Chart.Results[0].Indicators.Quotes[0].Low
	close := content.Chart.Results[0].Indicators.Quotes[0].Close

	result := make([]*PriceQuote, len(timestamps))
	for i := 0; i < len(result); i++ {
		result[i] = &PriceQuote{
			Timestamp: timestamps[i],
			Volume:    volumes[i],
			Open:      open[i],
			High:      high[i],
			Low:       low[i],
			Close:     close[i],
		}
	}

	return result
}
