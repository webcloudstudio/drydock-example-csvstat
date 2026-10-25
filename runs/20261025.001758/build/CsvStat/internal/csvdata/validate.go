package csvdata

import "fmt"

// Validate rejects a table that cannot be profiled: no records at all, an empty
// or repeated header name, or a record whose field count differs from the header.
// A header with no data records is valid.
func Validate(t Table) error {
	if len(t.Header) == 0 {
		return &RejectError{Reason: "file has no records"}
	}
	seen := make(map[string]bool, len(t.Header))
	for i, name := range t.Header {
		if name == "" {
			return &RejectError{Reason: fmt.Sprintf("header name %d is empty", i+1)}
		}
		if seen[name] {
			return &RejectError{Reason: fmt.Sprintf("header name %q is repeated", name)}
		}
		seen[name] = true
	}
	for _, rec := range t.Records {
		if len(rec.Fields) != len(t.Header) {
			return &RejectError{
				Line:   rec.Line,
				Reason: fmt.Sprintf("record has %d fields, header has %d", len(rec.Fields), len(t.Header)),
			}
		}
	}
	return nil
}
