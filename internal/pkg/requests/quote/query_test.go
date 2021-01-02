package quote

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	type args struct {
		symbol string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should return correct query for specified symbol",
			args: args{
				symbol: "BTC-USD",
			},
			want: "&symbols=BTC-USD&fields=extendedMarketChange,extendedMarketChangePercent,extendedMarketPrice,extendedMarketTime,regularMarketChange,regularMarketChangePercent,regularMarketPrice,regularMarketTime,circulatingSupply,ask,askSize,bid,bidSize,dayHigh,dayLow,regularMarketDayHigh,regularMarketDayLow,regularMarketVolume,volume",
		},
		{
			name: "Should return correct query for a different symbol",
			args: args{
				symbol: "BTC-EUR",
			},
			want: "&symbols=BTC-EUR&fields=extendedMarketChange,extendedMarketChangePercent,extendedMarketPrice,extendedMarketTime,regularMarketChange,regularMarketChangePercent,regularMarketPrice,regularMarketTime,circulatingSupply,ask,askSize,bid,bidSize,dayHigh,dayLow,regularMarketDayHigh,regularMarketDayLow,regularMarketVolume,volume",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Query(tt.args.symbol)

			assert.Equal(t, tt.want, got)
		})
	}
}
