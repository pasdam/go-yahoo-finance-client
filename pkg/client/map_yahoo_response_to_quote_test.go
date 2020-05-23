package client

import (
	"reflect"
	"testing"
)

func Test_mapYahooResponseToQuote(t *testing.T) {
	type args struct {
		content *quotesResponseContent
	}
	tests := []struct {
		name       string
		args       args
		wantQuotes []*PriceQuote
	}{
		{
			name: "",
			args: args{
				content: &quotesResponseContent{
					Chart: quotesChart{
						Results: []result{
							{
								Timestamps: []uint64{0, 1, 2},
								Indicators: indicators{
									Quotes: []*indicatorQuote{
										{
											Open:   []float64{1, 11, 111},
											High:   []float64{2, 22, 222},
											Low:    []float64{3, 33, 333},
											Close:  []float64{4, 44, 444},
											Volume: []float64{5, 55, 555},
										},
									},
								},
							},
						},
					},
				},
			},
			wantQuotes: []*PriceQuote{
				{
					Timestamp: 0,
					Open:      1,
					High:      2,
					Low:       3,
					Close:     4,
					Volume:    5,
				},
				{
					Timestamp: 1,
					Open:      11,
					High:      22,
					Low:       33,
					Close:     44,
					Volume:    55,
				},
				{
					Timestamp: 2,
					Open:      111,
					High:      222,
					Low:       333,
					Close:     444,
					Volume:    555,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotQuotes := mapYahooResponseToQuote(tt.args.content); !reflect.DeepEqual(gotQuotes, tt.wantQuotes) {
				t.Errorf("mapYahooResponseToQuote() = %v, want %v", gotQuotes, tt.wantQuotes)
			}
		})
	}
}
