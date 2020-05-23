package client

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient()

	assert.Equal(t, "https://query1.finance.yahoo.com", c.baseURL.String())
}

func TestNewClientWithURL(t *testing.T) {
	type args struct {
		u string
	}
	tests := []struct {
		args    args
		wantErr error
	}{
		{
			args: args{
				u: "localhost:123",
			},
		},
		{
			args: args{
				u: "\t invalid url",
			},
			wantErr: errors.New("parse \"\\t invalid url\": net/url: invalid control character in URL"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.u, func(t *testing.T) {
			got, err := NewClientWithURL(tt.args.u)

			if tt.wantErr != nil {
				assert.NotNil(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
				assert.Nil(t, got)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tt.args.u, got.baseURL.String())
			}
		})
	}
}

func TestClient_Quotes_Real(t *testing.T) {
	t.SkipNow()
	c := NewClient()
	quotes, err := c.Quotes("USD", "EUR", 1588766400, 1588767300)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(quotes))

	expected := []*PriceQuote{
		{
			Timestamp: 1588766400,
			Open:      0.9247000217437744,
			High:      0.9251000285148621,
			Low:       0.9243000149726868,
			Close:     0.925000011920929,
			Volume:    0,
		},
		{
			Timestamp: 1588767300,
			Open:      0,
			High:      0,
			Low:       0,
			Close:     0,
			Volume:    0,
		},
	}
	assert.Equal(t, expected, quotes)
}
