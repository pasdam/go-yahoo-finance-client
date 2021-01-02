package quote

type QuoteResponse struct {
	Results []Result `json:"result"`
	Error   string   `json:"error"`
}
