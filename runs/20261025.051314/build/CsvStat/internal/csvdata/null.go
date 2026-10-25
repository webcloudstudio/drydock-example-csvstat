package csvdata

import "strings"

// IsNull reports whether value, already trimmed of leading and trailing spaces,
// is a null spelling: empty, or NA, N/A or NULL ignoring case. No other token
// is null.
func IsNull(value string) bool {
	return value == "" ||
		strings.EqualFold(value, "NA") ||
		strings.EqualFold(value, "N/A") ||
		strings.EqualFold(value, "NULL")
}
