package client

import (
	"net/url"
	"strconv"
)

func quotesQuery(symbol string, fromTimestamp uint64, toTimestamp uint64) string {
	values := url.Values{}
	values.Add("period1", strconv.FormatUint(fromTimestamp, 10))
	values.Add("period2", strconv.FormatUint(toTimestamp, 10))
	values.Add("interval", "15m") // TODO make configurable
	values.Add("symbol", symbol)
	return values.Encode()
}
