package quote

type Result struct {
	Language                      string  `json:"language"`
	Region                        string  `json:"region"`
	QuoteType                     string  `json:"quoteType"`
	QuoteSourceName               string  `json:"quoteSourceName"`
	Triggerable                   bool    `json:"triggerable"`
	FiftyTwoWeekLow               float64 `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh              float64 `json:"fiftyTwoWeekHigh"`
	SourceInterval                uint32  `json:"sourceInterval"`
	ExchangeDataDelayedBy         uint32  `json:"exchangeDataDelayedBy"`
	Tradeable                     bool    `json:"tradeable"`
	FirstTradeDateMilliseconds    uint64  `json:"firstTradeDateMilliseconds"`
	CirculatingSupply             uint64  `json:"circulatingSupply"`
	EsgPopulated                  bool    `json:"esgPopulated"`
	RegularMarketChange           float64 `json:"regularMarketChange"`
	RegularMarketChangePercent    float64 `json:"regularMarketChangePercent"`
	RegularMarketTime             uint64  `json:"regularMarketTime"`
	RegularMarketPrice            float64 `json:"regularMarketPrice"`
	RegularMarketDayHigh          float64 `json:"regularMarketDayHigh"`
	RegularMarketDayRange         string  `json:"regularMarketDayRange"`
	RegularMarketDayLow           float64 `json:"regularMarketDayLow"`
	RegularMarketVolume           float64 `json:"regularMarketVolume"`
	RegularMarketPreviousClose    float64 `json:"regularMarketPreviousClose"`
	FullExchangeName              string  `json:"fullExchangeName"`
	FiftyTwoWeekLowChange         float64 `json:"fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent  float64 `json:"fiftyTwoWeekLowChangePercent"`
	FiftyTwoWeekRange             string  `json:"fiftyTwoWeekRange"`
	FiftyTwoWeekHighChange        float64 `json:"fiftyTwoWeekHighChange"`
	FiftyTwoWeekHighChangePercent float64 `json:"fiftyTwoWeekHighChangePercent"`
	Exchange                      string  `json:"exchange"`
	ExchangeTimezoneName          string  `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName     string  `json:"exchangeTimezoneShortName"`
	GmtOffSetMilliseconds         uint64  `json:"gmtOffSetMilliseconds"`
	Market                        string  `json:"market"`
	MarketState                   string  `json:"marketState"`
	Symbol                        string  `json:"symbol"`
}
