package chart

type Meta struct {
	Currency             string `json:"currency"`
	Symbol               string `json:"symbol"`
	DataGranularity      string `json:"dataGranularity"`
	RangeVal             string `json:"range"`
	Timezone             string `json:"timezone"`
	ExchangeTimezoneName string `json:"exchangeTimezoneName"`
}
