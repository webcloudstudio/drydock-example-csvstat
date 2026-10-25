package profile

import (
	"math/big"
	"sort"
)

// Median takes the column type and the trimmed non-null values of the column. ok is false when
// the type is not integer or number, or when there are no values.
func Median(typ string, values []string) (median *big.Rat, ok bool) {
	if !isNumeric(typ) {
		return nil, false
	}
	rats, ok := parseRats(nonNull(values))
	if !ok || len(rats) == 0 {
		return nil, false
	}
	sort.Slice(rats, func(i, j int) bool { return rats[i].Cmp(rats[j]) < 0 })

	n := len(rats)
	if n%2 == 1 {
		return rats[n/2], true
	}
	middle := new(big.Rat).Add(rats[n/2-1], rats[n/2])
	return middle.Quo(middle, big.NewRat(2, 1)), true
}
