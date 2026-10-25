package infer

import (
	"regexp"
	"strings"
	"time"
)

// Type names written to the type field of a profile row.
const (
	Integer = "integer"
	Number  = "number"
	Date    = "date"
	Boolean = "boolean"
	Text    = "text"
	Empty   = "empty"
)

var (
	integerPattern = regexp.MustCompile(`^-?[0-9]+$`)
	numberPattern  = regexp.MustCompile(`^-?[0-9]+(\.[0-9]+)?$`)
	datePattern    = regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`)
)

// ladder is the first-match order of the typed rules. A column takes the first
// type every one of its values satisfies; text is the fallback.
var ladder = []struct {
	typ   string
	match func(string) bool
}{
	{Integer, integerPattern.MatchString},
	{Number, numberPattern.MatchString},
	{Date, isDate},
	{Boolean, isBoolean},
}

// Classify returns the type of a column from its non-null, trimmed values. A
// column with no values is Empty. The caller removes null values first.
func Classify(values []string) string {
	if len(values) == 0 {
		return Empty
	}
	for _, rule := range ladder {
		if allMatch(values, rule.match) {
			return rule.typ
		}
	}
	return Text
}

func allMatch(values []string, match func(string) bool) bool {
	for _, v := range values {
		if !match(v) {
			return false
		}
	}
	return true
}

// isDate checks the YYYY-MM-DD shape, then the calendar with time.Parse so that
// impossible dates and non-leap February 29 are rejected.
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
