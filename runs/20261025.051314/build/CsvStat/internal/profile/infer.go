package profile

import (
	"regexp"
	"strings"
	"time"

	"csvstat/internal/csvdata"
)

// Column type names, in the order InferType tests them.
const (
	TypeEmpty   = "empty"
	TypeInteger = "integer"
	TypeNumber  = "number"
	TypeDate    = "date"
	TypeBoolean = "boolean"
	TypeText    = "text"
)

var (
	integerPattern = regexp.MustCompile(`^-?[0-9]+$`)
	numberPattern  = regexp.MustCompile(`^-?[0-9]+(\.[0-9]+)?$`)
	datePattern    = regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`)
)

// InferType returns the type of one column from its trimmed values. Null values
// are ignored. The first type in the order empty, integer, number, date,
// boolean, text that the non-null values all satisfy is returned. Classification
// never uses float parsing and never rewrites a value.
func InferType(values []string) string {
	var present []string
	for _, value := range values {
		if !csvdata.IsNull(value) {
			present = append(present, value)
		}
	}
	switch {
	case len(present) == 0:
		return TypeEmpty
	case allMatch(present, integerPattern.MatchString):
		return TypeInteger
	case allMatch(present, numberPattern.MatchString):
		return TypeNumber
	case allMatch(present, isDate):
		return TypeDate
	case allMatch(present, isBoolean):
		return TypeBoolean
	default:
		return TypeText
	}
}

func allMatch(values []string, match func(string) bool) bool {
	for _, value := range values {
		if !match(value) {
			return false
		}
	}
	return true
}

// isDate requires the YYYY-MM-DD shape and a real calendar date.
func isDate(value string) bool {
	if !datePattern.MatchString(value) {
		return false
	}
	_, err := time.Parse("2006-01-02", value)
	return err == nil
}

func isBoolean(value string) bool {
	return strings.EqualFold(value, "true") || strings.EqualFold(value, "false")
}
