package client

import (
	"net/url"
	"path"

	"github.com/pasdam/go-rest-util/pkg/restutil"
	"github.com/pasdam/go-yahoo-finance-client/internal/pkg/requests/chart"
	"github.com/pasdam/go-yahoo-finance-client/internal/pkg/requests/quote"
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

// Quotes retrieves the price quotes for the specified symbol within the
// required timeframe
func (c *Client) Quotes(symbol string, fromTimestamp uint64, toTimestamp uint64, interval Interval) ([]*PriceQuote, error) {

	c.baseURL.Path = path.Join("v8/finance/chart", symbol)
	c.baseURL.RawQuery = chart.QuotesQuery(symbol, fromTimestamp, toTimestamp, interval.String())

	response := &chart.QuotesResponseContent{}
	err := restutil.GetJSON(c.baseURL.String(), nil, response)
	if err != nil {
		return nil, err
	}

	return mapYahooResponseToQuote(response), nil
}

// CurrentRate retrieves the current exchange rate for the specified symbol
// pair
func (c *Client) CurrentRate(symbol string) (*ExchangeRate, error) {

	c.baseURL.Path = "v7/finance/quote"
	c.baseURL.RawQuery = quote.Query(symbol)

	response := &quote.Response{}
	err := restutil.GetJSON(c.baseURL.String(), nil, response)
	if err != nil {
		return nil, err
	}

	return mapQuoteResponseToQuote(response), nil
}
