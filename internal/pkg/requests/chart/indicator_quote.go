package chart

type IndicatorQuote struct {
	Volume []float64 `json:"volume"`
	Open   []float64 `json:"open"`
	High   []float64 `json:"high"`
	Low    []float64 `json:"low"`
	Close  []float64 `json:"close"`
}
