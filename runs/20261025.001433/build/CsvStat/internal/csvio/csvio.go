package csvio

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadError reports a file that cannot be opened or read. The CLI maps it to exit code 1.
type ReadError struct {
	Path string
	Err  error
}

func (e *ReadError) Error() string { return fmt.Sprintf("cannot read %s: %v", e.Path, e.Err) }

// Unwrap returns the underlying operating-system or reader error.
func (e *ReadError) Unwrap() error { return e.Err }

// RejectError reports input that is not accepted as a CSV table. The CLI maps it to exit code 3.
// Line is the 1-based line where the problem starts, or 0 when the problem concerns the whole file.
type RejectError struct {
	Path   string
	Line   int
	Reason string
}

func (e *RejectError) Error() string {
	if e.Line == 0 {
		return fmt.Sprintf("%s: %s", e.Path, e.Reason)
	}
	return fmt.Sprintf("%s line %d: %s", e.Path, e.Line, e.Reason)
}

// Table is a validated CSV file. Header names and values are trimmed of leading and trailing
// spaces, and Records keep file order.
type Table struct {
	Path    string
	Header  []string
	Records [][]string
}

// rawRecord is one record as encoding/csv returned it, with the line where it starts.
type rawRecord struct {
	line   int
	fields []string
}

// Load reads and validates the CSV file at path.
func Load(path string) (*Table, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, &ReadError{Path: path, Err: err}
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return nil, &ReadError{Path: path, Err: err}
	}
	if info.IsDir() {
		return nil, &ReadError{Path: path, Err: errors.New("is a directory")}
	}
	return read(path, f)
}

// Index returns the position of the header name, which must match exactly and case-sensitively.
func (t *Table) Index(name string) (int, bool) {
	for i, h := range t.Header {
		if h == name {
			return i, true
		}
	}
	return 0, false
}

// Column returns the trimmed values of column i in record order.
func (t *Table) Column(i int) []string {
	values := make([]string, len(t.Records))
	for r, rec := range t.Records {
		values[r] = rec[i]
	}
	return values
}

// read parses every record from r and validates the table structure. path names the source in errors.
func read(path string, r io.Reader) (*Table, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1

	var raw []rawRecord
	for {
		fields, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			var pe *csv.ParseError
			if errors.As(err, &pe) {
				return nil, &RejectError{Path: path, Line: pe.Line, Reason: pe.Err.Error()}
			}
			return nil, &ReadError{Path: path, Err: err}
		}
		line, _ := cr.FieldPos(0)
		raw = append(raw, rawRecord{line: line, fields: fields})
	}

	if len(raw) == 0 {
		return nil, &RejectError{Path: path, Reason: "file has no records"}
	}

	header := trimAll(raw[0].fields)
	seen := make(map[string]bool, len(header))
	for _, name := range header {
		if name == "" {
			return nil, &RejectError{Path: path, Line: raw[0].line, Reason: "empty header name"}
		}
		if seen[name] {
			return nil, &RejectError{Path: path, Line: raw[0].line, Reason: fmt.Sprintf("repeated header name %q", name)}
		}
		seen[name] = true
	}

	t := &Table{Path: path, Header: header}
	for _, rec := range raw[1:] {
		if len(rec.fields) != len(header) {
			return nil, &RejectError{
				Path:   path,
				Line:   rec.line,
				Reason: fmt.Sprintf("record has %d fields, header has %d", len(rec.fields), len(header)),
			}
		}
		t.Records = append(t.Records, trimAll(rec.fields))
	}
	return t, nil
}

// trimAll returns a copy of fields with only the space character trimmed from each one.
func trimAll(fields []string) []string {
	out := make([]string, len(fields))
	for i, f := range fields {
		out[i] = strings.Trim(f, " ")
	}
	return out
}
