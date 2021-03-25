package currency

import "testing"

func Test_CurrencyPairSymbol(t *testing.T) {
	type args struct {
		from string
		to   string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				from: "USD",
				to:   "EUR",
			},
			want: "EUR=X",
		},
		{
			args: args{
				from: "USD",
				to:   "JPY",
			},
			want: "JPY=X",
		},
		{
			args: args{
				from: "EUR",
				to:   "USD",
			},
			want: "EURUSD=X",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := PairSymbol(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("PairSymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
