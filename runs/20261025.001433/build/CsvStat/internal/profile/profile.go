package profile

import (
	"math/big"
	"sort"
	"strconv"
	"strings"

	"csvstat/internal/decimal"
)

// Header is the column order of every profile row.
var Header = []string{"column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"}

// Column types, in the order the type decision tests them.
const (
	typeEmpty   = "empty"
	typeInteger = "integer"
	typeNumber  = "number"
	typeDate    = "date"
	typeBoolean = "boolean"
	typeText    = "text"
)

// Row is the profile of one column.
type Row struct {
	Column   string
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

// Fields returns the row in Header order.
func (r Row) Fields() []string {
	return []string{
		r.Column,
		r.Type,
		strconv.Itoa(r.Rows),
		strconv.Itoa(r.Nulls),
		strconv.Itoa(r.Distinct),
		r.Min,
		r.Max,
		r.Sum,
		r.Mean,
		r.Median,
	}
}

// IsNull reports whether a trimmed value is a null spelling: empty, or NA, N/A, or NULL ignoring case.
func IsNull(v string) bool {
	switch strings.ToUpper(v) {
	case "", "NA", "N/A", "NULL":
		return true
	}
	return false
}

// Summarize computes the profile row for one column from its trimmed values in file order.
func Summarize(name string, values []string) Row {
	row := Row{Column: name, Rows: len(values)}
	present := make([]string, 0, len(values))
	distinct := make(map[string]struct{})
	for _, v := range values {
		if IsNull(v) {
			row.Nulls++
			continue
		}
		present = append(present, v)
		distinct[v] = struct{}{}
	}
	row.Distinct = len(distinct)
	row.Type = inferType(present)

	switch row.Type {
	case typeInteger, typeNumber:
		summarizeNumeric(&row, present)
	case typeDate, typeText:
		summarizeText(&row, present)
	case typeBoolean:
		summarizeBoolean(&row, present)
	}
	return row
}

// summarizeNumeric fills min, max, sum, mean, and median for an integer or number column.
// Every value is parsed exactly; inferType admitted only decimal literals, so parsing cannot fail.
func summarizeNumeric(row *Row, present []string) {
	values := make([]*big.Rat, len(present))
	for i, v := range present {
		values[i], _ = new(big.Rat).SetString(v)
	}

	sorted := append([]*big.Rat(nil), values...)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Cmp(sorted[j]) < 0 })
	n := len(sorted)

	sum := new(big.Rat)
	for _, v := range values {
		sum.Add(sum, v)
	}
	mean := new(big.Rat).Quo(sum, new(big.Rat).SetInt64(int64(n)))

	var median *big.Rat
	if n%2 == 1 {
		median = sorted[n/2]
	} else {
		median = new(big.Rat).Add(sorted[n/2-1], sorted[n/2])
		median.Quo(median, big.NewRat(2, 1))
	}

	format := decimal.Format
	if row.Type == typeInteger {
		format = decimal.Whole
		row.Sum = decimal.Whole(sum)
	} else {
		row.Sum = decimal.Format(sum)
	}
	row.Min = format(sorted[0])
	row.Max = format(sorted[n-1])
	row.Mean = decimal.Format(mean)
	row.Median = decimal.Format(median)
}

// summarizeText fills min and max by byte value for a date or text column.
func summarizeText(row *Row, present []string) {
	row.Min, row.Max = present[0], present[0]
	for _, v := range present[1:] {
		if v < row.Min {
			row.Min = v
		}
		if v > row.Max {
			row.Max = v
		}
	}
}

// summarizeBoolean fills min and max for a boolean column, in lower case.
func summarizeBoolean(row *Row, present []string) {
	anyFalse, anyTrue := false, false
	for _, v := range present {
		if strings.ToLower(v) == "true" {
			anyTrue = true
		} else {
			anyFalse = true
		}
	}
	row.Min = "true"
	if anyFalse {
		row.Min = "false"
	}
	row.Max = "false"
	if anyTrue {
		row.Max = "true"
	}
}
