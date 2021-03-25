package currency

import (
	"fmt"
)

// PairSymbol returns the symbol given a specific currency pair
func PairSymbol(from string, to string) string {
	if from == "USD" {
		return fmt.Sprintf("%s=X", to)
	}
	return fmt.Sprintf("%s%s=X", from, to)
}
