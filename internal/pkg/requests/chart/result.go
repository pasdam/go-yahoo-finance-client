package chart

type Result struct {
	Meta       Meta       `json:"meta"`
	Timestamps []uint64   `json:"timestamp"`
	Indicators Indicators `json:"indicators"`
}
