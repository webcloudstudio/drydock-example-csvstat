package profile

import (
	"math/big"
	"sort"
	"strconv"
	"strings"

	"csvstat/internal/csvdata"
	"csvstat/internal/decimal"
	"csvstat/internal/infer"
)

// Header is the fixed first row of every profile.
var Header = []string{"column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"}

// Rows returns one profile row per column of t, in file order. Each row has one
// field per Header name.
func Rows(t csvdata.Table) [][]string {
	rows := make([][]string, len(t.Header))
	for i, name := range t.Header {
		values := make([]string, len(t.Records))
		for j, rec := range t.Records {
			values[j] = rec.Fields[i]
		}
		rows[i] = column(name, values)
	}
	return rows
}

// column computes the profile row for one column from all of its trimmed values,
// nulls included.
func column(name string, values []string) []string {
	var present []string
	nulls := 0
	seen := make(map[string]bool)
	for _, v := range values {
		if csvdata.IsNull(v) {
			nulls++
			continue
		}
		present = append(present, v)
		seen[v] = true
	}

	typ := infer.Classify(present)
	row := make([]string, len(Header))
	row[0] = name
	row[1] = typ
	row[2] = strconv.Itoa(len(values))
	row[3] = strconv.Itoa(nulls)
	row[4] = strconv.Itoa(len(seen))
	row[5], row[6] = extremes(typ, present)
	if typ == infer.Integer || typ == infer.Number {
		row[7], row[8], row[9] = numeric(typ, present)
	}
	return row
}

// extremes returns the min and max of non-null values by the column's type. An
// empty column has no extremes, so both fields are blank.
func extremes(typ string, present []string) (string, string) {
	switch typ {
	case infer.Integer, infer.Number:
		sorted := sortedRats(present)
		lo, hi := sorted[0], sorted[len(sorted)-1]
		if typ == infer.Integer {
			return lo.Num().String(), hi.Num().String()
		}
		return decimal.Format(lo), decimal.Format(hi)
	case infer.Date, infer.Text:
		lo, hi := present[0], present[0]
		for _, v := range present[1:] {
			if v < lo {
				lo = v
			}
			if v > hi {
				hi = v
			}
		}
		return lo, hi
	case infer.Boolean:
		lo, hi := "true", "false"
		for _, v := range present {
			lower := strings.ToLower(v)
			if lower == "false" {
				lo = "false"
			}
			if lower == "true" {
				hi = "true"
			}
		}
		return lo, hi
	default:
		return "", ""
	}
}

// numeric returns the sum, mean, and median of an integer or number column. The
// sum of an integer column is a plain integer; every other figure is written to
// two decimal places. Arithmetic is exact, and rounding happens only in decimal.Format.
func numeric(typ string, present []string) (string, string, string) {
	sorted := sortedRats(present)

	total := new(big.Rat)
	for _, r := range sorted {
		total.Add(total, r)
	}
	sum := decimal.Format(total)
	if typ == infer.Integer {
		sum = total.Num().String()
	}

	mean := new(big.Rat).Quo(total, big.NewRat(int64(len(sorted)), 1))

	middle := len(sorted) / 2
	median := new(big.Rat).Set(sorted[middle])
	if len(sorted)%2 == 0 {
		median.Add(sorted[middle-1], sorted[middle])
		median.Quo(median, big.NewRat(2, 1))
	}

	return sum, decimal.Format(mean), decimal.Format(median)
}

// sortedRats parses values exactly and sorts them numerically. The caller passes
// only values already classified as integer or number, so every value parses.
func sortedRats(values []string) []*big.Rat {
	rats := make([]*big.Rat, len(values))
	for i, v := range values {
		r, ok := new(big.Rat).SetString(v)
		if !ok {
			panic("profile: value " + strconv.Quote(v) + " classified numeric but does not parse")
		}
		rats[i] = r
	}
	sort.Slice(rats, func(i, j int) bool { return rats[i].Cmp(rats[j]) < 0 })
	return rats
}
