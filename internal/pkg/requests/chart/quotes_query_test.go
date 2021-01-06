package chart

import "testing"

func Test_quotesQuery(t *testing.T) {
	type args struct {
		symbol        string
		fromTimestamp uint64
		toTimestamp   uint64
		interval      string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				symbol:        "EUR=X",
				fromTimestamp: 123,
				toTimestamp:   234,
				interval:      "15m",
			},
			want: "interval=15m&period1=123&period2=234&symbol=EUR%3DX",
		},
		{
			args: args{
				symbol:        "JPY=X",
				fromTimestamp: 345,
				toTimestamp:   567,
				interval:      "30m",
			},
			want: "interval=30m&period1=345&period2=567&symbol=JPY%3DX",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := QuotesQuery(tt.args.symbol, tt.args.fromTimestamp, tt.args.toTimestamp, tt.args.interval); got != tt.want {
				t.Errorf("quotesQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
