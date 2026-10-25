package cli

import (
	"fmt"
	"math/big"
	"strconv"

	"csvstat/internal/csvdata"
	"csvstat/internal/profile"
)

// profileHeader is the first row of the default profile output.
var profileHeader = []string{"column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"}

// profileRows returns the header and one row per column, in file order.
func profileRows(t *csvdata.Table) [][]string {
	rows := [][]string{profileHeader}
	for c, name := range t.Header {
		rows = append(rows, profileRow(name, columnValues(t, c)))
	}
	return rows
}

// columnRows returns the header and the profile row of the named column. The name must match a
// trimmed header name exactly; an absent name is an error.
func columnRows(t *csvdata.Table, name string) ([][]string, error) {
	c := columnIndex(t.Header, name)
	if c < 0 {
		return nil, fmt.Errorf("no column named %q", name)
	}
	return [][]string{profileHeader, profileRow(name, columnValues(t, c))}, nil
}

// profileRow computes one column's profile. Statistics that the column type does not carry are
// left blank.
func profileRow(name string, values []string) []string {
	typ := profile.InferType(values)
	counts := profile.Counts(values)
	extremes := profile.MinMax(typ, values)

	row := []string{
		name,
		typ,
		strconv.Itoa(counts.Rows),
		strconv.Itoa(counts.Nulls),
		strconv.Itoa(counts.Distinct),
		extremeText(typ, extremes.Min, extremes.MinRat),
		extremeText(typ, extremes.Max, extremes.MaxRat),
		"", "", "",
	}
	if sum, mean, ok := profile.SumMean(typ, values); ok {
		row[7] = sumText(typ, sum)
		row[8] = profile.FormatDecimal(mean)
	}
	if median, ok := profile.Median(typ, values); ok {
		row[9] = profile.FormatDecimal(median)
	}
	return row
}

// extremeText renders a minimum or maximum. Number extremes are written with two decimals;
// every other type is written as its text form.
func extremeText(typ string, text string, exact *big.Rat) string {
	if typ == profile.TypeNumber && exact != nil {
		return profile.FormatDecimal(exact)
	}
	return text
}

// sumText renders an exact sum: integer columns as plain integers, number columns with two decimals.
func sumText(typ string, sum *big.Rat) string {
	if typ == profile.TypeInteger {
		return profile.FormatInteger(sum)
	}
	return profile.FormatDecimal(sum)
}

// columnValues returns the trimmed value of column c for every data record, in file order.
func columnValues(t *csvdata.Table, c int) []string {
	values := make([]string, len(t.Records))
	for r, record := range t.Records {
		values[r] = record.Fields[c]
	}
	return values
}
