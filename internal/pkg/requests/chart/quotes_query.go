package chart

import (
	"net/url"
	"strconv"
)

func QuotesQuery(symbol string, fromTimestamp uint64, toTimestamp uint64, interval string) string {
	values := url.Values{}
	values.Add("period1", strconv.FormatUint(fromTimestamp, 10))
	values.Add("period2", strconv.FormatUint(toTimestamp, 10))
	values.Add("interval", interval)
	values.Add("symbol", symbol)
	return values.Encode()
}
