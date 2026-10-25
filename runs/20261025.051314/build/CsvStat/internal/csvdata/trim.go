package csvdata

import "strings"

// Trim removes leading and trailing spaces (U+0020) from every header name and
// every record field, in place. Other whitespace and interior spaces are kept.
// Line numbers and field counts are unchanged.
func Trim(t *Table) {
	for i, name := range t.Header {
		t.Header[i] = strings.Trim(name, " ")
	}
	for r := range t.Records {
		fields := t.Records[r].Fields
		for i, value := range fields {
			fields[i] = strings.Trim(value, " ")
		}
	}
}
