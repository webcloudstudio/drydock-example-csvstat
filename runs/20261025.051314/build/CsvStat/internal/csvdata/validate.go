package csvdata

import (
	"errors"
	"fmt"
)

// ErrRejected matches every input rejection with errors.Is.
var ErrRejected = errors.New("input rejected")

// RejectedError describes why an input was rejected. Line is the 1-based source
// line of the offending record, or 0 when the rejection concerns the header or
// the whole input.
type RejectedError struct {
	Line   int
	Reason string
}

func (e *RejectedError) Error() string {
	if e.Line == 0 {
		return e.Reason
	}
	return fmt.Sprintf("line %d: %s", e.Line, e.Reason)
}

// Unwrap makes errors.Is(err, ErrRejected) true for every RejectedError.
func (e *RejectedError) Unwrap() error {
	return ErrRejected
}

// Validate rejects input that must produce no output: no header record (empty
// or blank-only input), an empty or repeated header name, and any record whose
// field count differs from the header. Call it after Trim, so names compare as
// trimmed. A header-only table is valid.
func Validate(t *Table) error {
	if t.Header == nil {
		return &RejectedError{Reason: "no header record: input is empty or blank"}
	}

	seen := make(map[string]bool, len(t.Header))
	for i, name := range t.Header {
		if name == "" {
			return &RejectedError{Reason: fmt.Sprintf("header column %d is empty", i+1)}
		}
		if seen[name] {
			return &RejectedError{Reason: fmt.Sprintf("header name %q is repeated", name)}
		}
		seen[name] = true
	}

	for _, rec := range t.Records {
		if len(rec.Fields) != len(t.Header) {
			return &RejectedError{
				Line:   rec.Line,
				Reason: fmt.Sprintf("record has %d fields, header has %d", len(rec.Fields), len(t.Header)),
			}
		}
	}
	return nil
}
