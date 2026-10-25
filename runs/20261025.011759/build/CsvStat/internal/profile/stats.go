package profile

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

// numericStats returns the sum, mean, and median of validated integer or number values.
// The sum of an integer column is a plain integer; mean and median always have two decimals.
func numericStats(typ string, present []string) (sum, mean, median string) {
	nums := parseRats(present)
	total := new(big.Rat)
	for _, n := range nums {
		total.Add(total, n)
	}
	mean = FormatRat(new(big.Rat).Quo(total, big.NewRat(int64(len(nums)), 1)))
	median = FormatRat(medianOf(nums))
	sum = formatValue(typ, total)
	return sum, mean, median
}

// medianOf returns the exact median; for an even count it is the mean of the two middle values.
func medianOf(nums []*big.Rat) *big.Rat {
	sorted := make([]*big.Rat, len(nums))
	copy(sorted, nums)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Cmp(sorted[j]) < 0 })

	n := len(sorted)
	if n%2 == 1 {
		return sorted[n/2]
	}
	middle := new(big.Rat).Add(sorted[n/2-1], sorted[n/2])
	return middle.Quo(middle, big.NewRat(2, 1))
}

// minMax returns the smallest and largest non-null value for a column type, or blanks when the
// type has no ordering or no value.
func minMax(typ string, present []string) (lo, hi string) {
	if len(present) == 0 {
		return "", ""
	}
	switch typ {
	case TypeInteger, TypeNumber:
		nums := parseRats(present)
		least, most := nums[0], nums[0]
		for _, n := range nums[1:] {
			if n.Cmp(least) < 0 {
				least = n
			}
			if n.Cmp(most) > 0 {
				most = n
			}
		}
		return formatValue(typ, least), formatValue(typ, most)
	case TypeDate:
		// Validated YYYY-MM-DD text orders chronologically as bytes.
		least, most := present[0], present[0]
		for _, v := range present[1:] {
			if v < least {
				least = v
			}
			if v > most {
				most = v
			}
		}
		return least, most
	case TypeBoolean:
		hasFalse, hasTrue := false, false
		for _, v := range present {
			if strings.EqualFold(v, "false") {
				hasFalse = true
			} else {
				hasTrue = true
			}
		}
		// false sorts below true, so min is false when any value is false, and max is true when any is true.
		lo, hi = "true", "false"
		if hasFalse {
			lo = "false"
		}
		if hasTrue {
			hi = "true"
		}
		return lo, hi
	case TypeText:
		least, most := present[0], present[0]
		for _, v := range present[1:] {
			if v < least {
				least = v
			}
			if v > most {
				most = v
			}
		}
		return least, most
	}
	return "", ""
}

// parseRats converts validated integer or number text to exact rationals. InferType has already
// matched every value against the numeric pattern, so each parse succeeds.
func parseRats(values []string) []*big.Rat {
	out := make([]*big.Rat, len(values))
	for i, v := range values {
		r, ok := new(big.Rat).SetString(v)
		if !ok {
			panic(fmt.Sprintf("profile: numeric value %q passed InferType but does not parse", v))
		}
		out[i] = r
	}
	return out
}

// formatValue renders an exact value for a column type: integers as plain integers, numbers with
// two decimals.
func formatValue(typ string, r *big.Rat) string {
	if typ == TypeInteger {
		return r.Num().String()
	}
	return FormatRat(r)
}

// FormatRat renders r with exactly two decimal places, rounded once half away from zero.
func FormatRat(r *big.Rat) string {
	abs := new(big.Rat).Abs(r)
	// floor(abs*100 + 1/2) computed exactly as floor((200*num + den) / (2*den)).
	scaled := new(big.Int).Mul(abs.Num(), big.NewInt(200))
	scaled.Add(scaled, abs.Denom())
	hundredths := new(big.Int).Quo(scaled, new(big.Int).Mul(abs.Denom(), big.NewInt(2)))

	whole := new(big.Int)
	frac := new(big.Int)
	whole.QuoRem(hundredths, big.NewInt(100), frac)

	sign := ""
	if r.Sign() < 0 && hundredths.Sign() != 0 {
		sign = "-"
	}
	return fmt.Sprintf("%s%s.%02d", sign, whole.String(), frac.Int64())
}
