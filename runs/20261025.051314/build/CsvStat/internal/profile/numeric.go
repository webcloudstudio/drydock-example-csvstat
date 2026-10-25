package profile

import (
	"math/big"

	"csvstat/internal/csvdata"
)

// isNumeric reports whether a column type carries sum, mean, and median.
func isNumeric(typ string) bool {
	return typ == TypeInteger || typ == TypeNumber
}

// nonNull returns the values that are not null spellings, in their original order.
func nonNull(values []string) []string {
	var present []string
	for _, value := range values {
		if !csvdata.IsNull(value) {
			present = append(present, value)
		}
	}
	return present
}

// parseRats parses every value as an exact decimal into a new big.Rat. ok is
// false when any value is not a plain decimal number, so no float parsing is
// ever involved and the caller's slice is never reordered.
func parseRats(values []string) (rats []*big.Rat, ok bool) {
	rats = make([]*big.Rat, len(values))
	for i, value := range values {
		if !numberPattern.MatchString(value) {
			return nil, false
		}
		rat, valid := new(big.Rat).SetString(value)
		if !valid {
			return nil, false
		}
		rats[i] = rat
	}
	return rats, true
}
