package client

import (
	"net/url"
	"path"

	"github.com/pasdam/go-rest-util/pkg/restutil"
	"github.com/pasdam/go-yahoo-finance-client/internal/pkg/requests/chart"
)

// Client is the client to use to interact with Yahoo Finance API
type Client struct {
	baseURL *url.URL
}

// NewClient creates a new client with the default URL
func NewClient() *Client {
	u, _ := NewClientWithURL("https://query1.finance.yahoo.com")
	return u
}

// NewClientWithURL creates a new client with the specified URL
func NewClientWithURL(u string) (*Client, error) {
	uu, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	return &Client{
		baseURL: uu,
	}, nil
}

// Quotes retrieves the price quotes for the specified currency pair within the
// required timeframe
func (c *Client) Quotes(baseCurrency string, quoteCurrency string, fromTimestamp uint64, toTimestamp uint64) ([]*PriceQuote, error) {

	symbol := currencyPairSymbol(baseCurrency, quoteCurrency)

	c.baseURL.Path = path.Join("v8/finance/chart", symbol)
	c.baseURL.RawQuery = chart.QuotesQuery(symbol, fromTimestamp, toTimestamp)

	response := &chart.QuotesResponseContent{}
	err := restutil.GetJSON(c.baseURL.String(), response)
	if err != nil {
		return nil, err
	}

	return mapYahooResponseToQuote(response), nil
}
