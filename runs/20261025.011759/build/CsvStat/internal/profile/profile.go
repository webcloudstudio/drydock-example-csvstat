// Package profile classifies nulls, infers column types, and computes counts, statistics, and frequencies.
package profile

import (
	"regexp"
	"strings"
	"time"
)

// Column types, in the order InferType tests them.
const (
	TypeInteger = "integer"
	TypeNumber  = "number"
	TypeDate    = "date"
	TypeBoolean = "boolean"
	TypeText    = "text"
	TypeEmpty   = "empty"
)

// nullSpellings are the non-empty values classified as null, compared ignoring case.
var nullSpellings = []string{"NA", "N/A", "NULL"}

var (
	integerPattern = regexp.MustCompile(`^-?[0-9]+$`)
	numberPattern  = regexp.MustCompile(`^-?[0-9]+(\.[0-9]+)?$`)
	datePattern    = regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`)
)

// Column is the profile of one column.
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

// IsNull reports whether a value is null: empty after trimming, or NA, N/A, or NULL ignoring case.
func IsNull(value string) bool {
	v := strings.TrimSpace(value)
	if v == "" {
		return true
	}
	for _, spelling := range nullSpellings {
		if strings.EqualFold(v, spelling) {
			return true
		}
	}
	return false
}

// InferType returns the first of integer, number, date, boolean, or text that every non-null
// value satisfies, or empty when no value is non-null.
func InferType(values []string) string {
	present := nonNull(values)
	if len(present) == 0 {
		return TypeEmpty
	}
	if allMatch(present, integerPattern.MatchString) {
		return TypeInteger
	}
	if allMatch(present, numberPattern.MatchString) {
		return TypeNumber
	}
	if allMatch(present, isDate) {
		return TypeDate
	}
	if allMatch(present, isBoolean) {
		return TypeBoolean
	}
	return TypeText
}

// ProfileColumn computes the full profile of one column from its trimmed values.
func ProfileColumn(name string, values []string) Column {
	typ := InferType(values)
	rows, nulls, distinct := countValues(values)
	present := nonNull(values)

	col := Column{
		Name:     name,
		Type:     typ,
		Rows:     rows,
		Nulls:    nulls,
		Distinct: distinct,
	}
	col.Min, col.Max = minMax(typ, present)
	if typ == TypeInteger || typ == TypeNumber {
		col.Sum, col.Mean, col.Median = numericStats(typ, present)
	}
	return col
}

// countValues returns the row count, null count, and distinct non-null count.
func countValues(values []string) (rows, nulls, distinct int) {
	seen := make(map[string]bool)
	for _, v := range values {
		if IsNull(v) {
			nulls++
			continue
		}
		seen[strings.TrimSpace(v)] = true
	}
	return len(values), nulls, len(seen)
}

// nonNull returns the trimmed values that are not null.
func nonNull(values []string) []string {
	out := make([]string, 0, len(values))
	for _, v := range values {
		if !IsNull(v) {
			out = append(out, strings.TrimSpace(v))
		}
	}
	return out
}

func allMatch(values []string, ok func(string) bool) bool {
	for _, v := range values {
		if !ok(v) {
			return false
		}
	}
	return true
}

// isDate requires the YYYY-MM-DD shape and a real calendar date, so leap years and month
// lengths are checked by the standard library.
func isDate(v string) bool {
	if !datePattern.MatchString(v) {
		return false
	}
	_, err := time.Parse("2006-01-02", v)
	return err == nil
}

func isBoolean(v string) bool {
	return strings.EqualFold(v, "true") || strings.EqualFold(v, "false")
}
