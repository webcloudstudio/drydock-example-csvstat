package profile

import (
	"fmt"
	"math/big"
	"sort"
	"strings"

	"csvstat/internal/reader"
	"csvstat/internal/typing"
)

// Column is one row of the profile. Statistic fields are empty when they do not apply.
type Column struct {
	Name     string
	Type     string
	Rows     int
	Nulls    int
	Distinct int
	Min      string
	Max      string
	Sum      string
	Mean     string
	Median   string
}

// Describe profiles one column from its trimmed values.
func Describe(name string, values []string) Column {
	rows, nulls, distinct := Counts(values)
	typ := typing.Infer(values)
	lo, hi := MinMax(typ, values)
	sum, mean, median := NumericStats(typ, values)
	return Column{
		Name: name, Type: typ,
		Rows: rows, Nulls: nulls, Distinct: distinct,
		Min: lo, Max: hi, Sum: sum, Mean: mean, Median: median,
	}
}

// Counts returns the number of data values, the number of null values, and the number of
// distinct non-null values. The values must already be trimmed; comparison is exact.
func Counts(values []string) (rows, nulls, distinct int) {
	seen := make(map[string]struct{})
	for _, v := range values {
		if reader.IsNull(v) {
			nulls++
			continue
		}
		seen[v] = struct{}{}
	}
	return len(values), nulls, len(seen)
}

// MinMax returns the smallest and largest non-null value for a column of type typ.
// Integers and numbers compare exactly; numbers are written with two decimal places.
// Booleans report false/true in lowercase. Text and dates compare by byte value.
func MinMax(typ string, values []string) (lo, hi string) {
	nonNull := nonNullValues(values)
	if len(nonNull) == 0 {
		return "", ""
	}
	switch typ {
	case "integer", "number":
		low, high := ratOf(nonNull[0]), ratOf(nonNull[0])
		for _, v := range nonNull[1:] {
			r := ratOf(v)
			if r.Cmp(low) < 0 {
				low = r
			}
			if r.Cmp(high) > 0 {
				high = r
			}
		}
		if typ == "integer" {
			return low.Num().String(), high.Num().String()
		}
		return fixed2(low), fixed2(high)
	case "date", "text":
		lo, hi = nonNull[0], nonNull[0]
		for _, v := range nonNull[1:] {
			if v < lo {
				lo = v
			}
			if v > hi {
				hi = v
			}
		}
		return lo, hi
	case "boolean":
		lo, hi = "true", "false"
		for _, v := range nonNull {
			if strings.EqualFold(v, "false") {
				lo = "false"
			}
			if strings.EqualFold(v, "true") {
				hi = "true"
			}
		}
		return lo, hi
	default:
		return "", ""
	}
}

// NumericStats returns the sum, mean, and median of an integer or number column. Every other
// type returns three empty strings. Arithmetic is exact; rounding happens once when writing.
func NumericStats(typ string, values []string) (sum, mean, median string) {
	if typ != "integer" && typ != "number" {
		return "", "", ""
	}
	nonNull := nonNullValues(values)
	if len(nonNull) == 0 {
		return "", "", ""
	}
	nums := make([]*big.Rat, len(nonNull))
	total := new(big.Rat)
	for i, v := range nonNull {
		nums[i] = ratOf(v)
		total.Add(total, nums[i])
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i].Cmp(nums[j]) < 0 })

	n := len(nums)
	var mid *big.Rat
	if n%2 == 1 {
		mid = nums[n/2]
	} else {
		mid = new(big.Rat).Add(nums[n/2-1], nums[n/2])
		mid.Quo(mid, big.NewRat(2, 1))
	}
	avg := new(big.Rat).Quo(total, big.NewRat(int64(n), 1))

	if typ == "integer" {
		sum = total.Num().String()
	} else {
		sum = fixed2(total)
	}
	return sum, fixed2(avg), fixed2(mid)
}

func nonNullValues(values []string) []string {
	out := make([]string, 0, len(values))
	for _, v := range values {
		if !reader.IsNull(v) {
			out = append(out, v)
		}
	}
	return out
}

// ratOf parses a value already accepted by the integer or number pattern. Those patterns are
// a subset of big.Rat syntax, so the parse cannot fail for a value that reached here.
func ratOf(s string) *big.Rat {
	r, _ := new(big.Rat).SetString(s)
	return r
}

// fixed2 writes x with exactly two decimal places, rounding once, half away from zero.
// A negative result that rounds to zero is written without a sign.
func fixed2(x *big.Rat) string {
	a := new(big.Rat).Abs(x)
	// floor(a*100 + 1/2) = floor((200*num + den) / (2*den)) for a non-negative a.
	num := new(big.Int).Mul(a.Num(), big.NewInt(200))
	num.Add(num, a.Denom())
	q := new(big.Int).Quo(num, new(big.Int).Mul(a.Denom(), big.NewInt(2)))

	whole, frac := new(big.Int).QuoRem(q, big.NewInt(100), new(big.Int))
	s := fmt.Sprintf("%s.%02d", whole.String(), frac.Int64())
	if x.Sign() < 0 && q.Sign() != 0 {
		s = "-" + s
	}
	return s
}
