package client

import (
	"fmt"
)

func currencyPairSymbol(from string, to string) string {
	if from == "USD" {
		return fmt.Sprintf("%s=X", to)
	}
	return fmt.Sprintf("%s%s=X", from, to)
}
