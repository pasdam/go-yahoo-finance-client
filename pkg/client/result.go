package client

type result struct {
	Meta       meta       `json:"meta"`
	Timestamps []uint64   `json:"timestamp"`
	Indicators indicators `json:"indicators"`
}
