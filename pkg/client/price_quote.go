package client

// PriceQuote contains all the quote info about a specific financial instrument's price for a specific timeframe
type PriceQuote struct {

	// Indicates the closing price
	Close float64

	// Indicates the highest price
	High float64

	// Indicates the lowest price
	Low float64

	// Indicates the open price
	Open float64

	// Indicates the open timestamp
	Timestamp uint64

	// Indicates the traded volume
	Volume float64
}
