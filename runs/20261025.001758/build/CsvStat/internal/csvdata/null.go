package csvdata

import "strings"

// IsNull reports whether a trimmed value is null: the empty string, or NA, N/A or
// NULL in any letter case. Every statistic uses this one rule.
func IsNull(value string) bool {
	return value == "" ||
		strings.EqualFold(value, "NA") ||
		strings.EqualFold(value, "N/A") ||
		strings.EqualFold(value, "NULL")
}
