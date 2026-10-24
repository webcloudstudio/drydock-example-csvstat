package reader

import "fmt"

// RejectError reports input that was read but rejected. It is a content failure and exits 3.
type RejectError struct {
	Line   int // 1-based input line of the offending record; 0 when the whole file is rejected
	Reason string
}

func (e *RejectError) Error() string {
	if e.Line == 0 {
		return e.Reason
	}
	return fmt.Sprintf("line %d: %s", e.Line, e.Reason)
}

// Validate checks a trimmed table. It rejects an empty file, an empty or repeated header name,
// and any record whose field count differs from the header. A header-only table is valid.
func Validate(t Table) error {
	if len(t.Header) == 0 {
		return &RejectError{Reason: "file is empty"}
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
	for i, row := range t.Rows {
		if len(row) != len(t.Header) {
			return &RejectError{
				Line:   t.Lines[i],
				Reason: fmt.Sprintf("record has %d fields, header has %d", len(row), len(t.Header)),
			}
		}
	}
	return nil
}
