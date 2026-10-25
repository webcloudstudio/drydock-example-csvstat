package profile

import "math/big"

// SumMean takes the column type and the trimmed non-null values of the column. ok is false when
// the type is not integer or number, or when there are no values.
func SumMean(typ string, values []string) (sum, mean *big.Rat, ok bool) {
	if !isNumeric(typ) {
		return nil, nil, false
	}
	rats, ok := parseRats(nonNull(values))
	if !ok || len(rats) == 0 {
		return nil, nil, false
	}
	sum = new(big.Rat)
	for _, rat := range rats {
		sum.Add(sum, rat)
	}
	count := new(big.Rat).SetInt64(int64(len(rats)))
	mean = new(big.Rat).Quo(sum, count)
	return sum, mean, true
}
