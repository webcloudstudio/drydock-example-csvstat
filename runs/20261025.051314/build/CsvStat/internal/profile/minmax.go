package profile

import (
	"math/big"
	"strings"
)

// Extremes holds the minimum and maximum of one column.
type Extremes struct {
	Min, Max       string   // text form, see MinMax
	MinRat, MaxRat *big.Rat // exact values for integer and number columns, else nil
}

// MinMax takes the column type and the trimmed values of the column. Null values are ignored.
// Integer extremes are plain integers; number extremes keep their source text and carry exact
// values in MinRat and MaxRat. Date text orders chronologically because it is fixed-width
// YYYY-MM-DD, and text orders by byte value.
func MinMax(typ string, values []string) Extremes {
	present := nonNull(values)
	if len(present) == 0 {
		return Extremes{}
	}
	switch typ {
	case TypeInteger, TypeNumber:
		return numericExtremes(typ, present)
	case TypeDate, TypeText:
		lo, hi := present[0], present[0]
		for _, value := range present[1:] {
			if value < lo {
				lo = value
			}
			if value > hi {
				hi = value
			}
		}
		return Extremes{Min: lo, Max: hi}
	case TypeBoolean:
		return booleanExtremes(present)
	default:
		return Extremes{}
	}
}

// numericExtremes compares the values exactly as big.Rat. Integer extremes are rewritten as plain
// integers; number extremes keep the source text of the extreme value.
func numericExtremes(typ string, values []string) Extremes {
	rats, ok := parseRats(values)
	if !ok {
		return Extremes{}
	}
	lo, hi := 0, 0
	for i := range rats {
		if rats[i].Cmp(rats[lo]) < 0 {
			lo = i
		}
		if rats[i].Cmp(rats[hi]) > 0 {
			hi = i
		}
	}
	result := Extremes{Min: values[lo], Max: values[hi], MinRat: rats[lo], MaxRat: rats[hi]}
	if typ == TypeInteger {
		result.Min = FormatInteger(rats[lo])
		result.Max = FormatInteger(rats[hi])
	}
	return result
}

// booleanExtremes reports false as the minimum if any value is false, else true; and true as the
// maximum if any value is true, else false. The output is lowercase whatever the input case.
func booleanExtremes(values []string) Extremes {
	lo, hi := "true", "false"
	for _, value := range values {
		if strings.EqualFold(value, "false") {
			lo = "false"
		}
		if strings.EqualFold(value, "true") {
			hi = "true"
		}
	}
	return Extremes{Min: lo, Max: hi}
}
