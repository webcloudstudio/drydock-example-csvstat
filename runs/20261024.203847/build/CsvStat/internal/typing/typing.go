package typing

import (
	"regexp"
	"strings"
	"time"

	"csvstat/internal/reader"
)

var (
	integerPattern = regexp.MustCompile(`^-?[0-9]+$`)
	numberPattern  = regexp.MustCompile(`^-?[0-9]+(\.[0-9]+)?$`)
	datePattern    = regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`)
)

// Infer returns the type token of a column from its trimmed values. Null values are skipped.
// The first of integer, number, date, boolean, text that every non-null value satisfies wins;
// a column with no non-null value is empty.
func Infer(values []string) string {
	found := false
	isInt, isNum, isDate, isBool := true, true, true, true
	for _, v := range values {
		if reader.IsNull(v) {
			continue
		}
		found = true
		isInt = isInt && integerPattern.MatchString(v)
		isNum = isNum && numberPattern.MatchString(v)
		isDate = isDate && IsDate(v)
		isBool = isBool && (strings.EqualFold(v, "true") || strings.EqualFold(v, "false"))
	}
	switch {
	case !found:
		return "empty"
	case isInt:
		return "integer"
	case isNum:
		return "number"
	case isDate:
		return "date"
	case isBool:
		return "boolean"
	default:
		return "text"
	}
}

// IsDate reports whether s is a real calendar date written YYYY-MM-DD. Leap years are honored
// and no value is normalized into validity.
func IsDate(s string) bool {
	if !datePattern.MatchString(s) {
		return false
	}
	t, err := time.Parse("2006-01-02", s)
	return err == nil && t.Format("2006-01-02") == s
}
