package profile

import (
	"regexp"
	"strings"
	"time"
)

var (
	integerPattern = regexp.MustCompile(`^-?[0-9]+$`)
	numberPattern  = regexp.MustCompile(`^-?[0-9]+(\.[0-9]+)?$`)
	datePattern    = regexp.MustCompile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`)
)

// inferType returns the first of integer, number, date, boolean, and text that every present
// (non-null) value satisfies, or empty when no value is present.
func inferType(present []string) string {
	switch {
	case len(present) == 0:
		return typeEmpty
	case allMatch(present, integerPattern.MatchString):
		return typeInteger
	case allMatch(present, numberPattern.MatchString):
		return typeNumber
	case allMatch(present, isDate):
		return typeDate
	case allMatch(present, isBoolean):
		return typeBoolean
	}
	return typeText
}

func allMatch(values []string, ok func(string) bool) bool {
	for _, v := range values {
		if !ok(v) {
			return false
		}
	}
	return true
}

// isDate accepts exactly YYYY-MM-DD naming a real calendar date. The length and pattern checks
// reject non-padded forms that time.Parse alone would accept.
func isDate(v string) bool {
	if len(v) != 10 || !datePattern.MatchString(v) {
		return false
	}
	_, err := time.Parse("2006-01-02", v)
	return err == nil
}

// isBoolean accepts true or false in any letter case.
func isBoolean(v string) bool {
	lower := strings.ToLower(v)
	return lower == "true" || lower == "false"
}
