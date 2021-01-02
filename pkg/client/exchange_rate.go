package client

// ExchangeRate contains the rate info about a specific financial instrument
type ExchangeRate struct {

	// Buy price
	Buy float64

	// Sell price
	Sell float64

	// Timestamp of the rate's valuation
	Timestamp uint64
}
