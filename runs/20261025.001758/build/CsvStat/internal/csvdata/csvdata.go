package csvdata

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
)

// Table is a trimmed CSV file: the header names and the data records beneath them.
type Table struct {
	Header  []string
	Records []Record
}

// Record is one data row. Line is the 1-based line on which the record starts.
type Record struct {
	Line   int
	Fields []string
}

// RejectError reports content that was read but cannot be profiled. Line is the
// 1-based line of the offending record, or 0 when the problem concerns the file as a whole.
type RejectError struct {
	Line   int
	Reason string
}

func (e *RejectError) Error() string {
	if e.Line == 0 {
		return e.Reason
	}
	return fmt.Sprintf("line %d: %s", e.Line, e.Reason)
}

// Read parses the CSV file at path. Every header name and value has leading and
// trailing spaces trimmed. A file that cannot be opened or read returns an error
// wrapping *fs.PathError; malformed CSV returns a *RejectError. Structural rules
// are applied by Validate.
func Read(path string) (Table, error) {
	f, err := os.Open(path)
	if err != nil {
		return Table{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1

	var table Table
	for {
		fields, err := r.Read()
		if err == io.EOF {
			return table, nil
		}
		if err != nil {
			return Table{}, readError(path, err)
		}
		trimAll(fields)
		if table.Header == nil {
			table.Header = fields
			continue
		}
		line, _ := r.FieldPos(0)
		table.Records = append(table.Records, Record{Line: line, Fields: fields})
	}
}

// readError classifies a failure from csv.Reader: a syntax error is rejected
// content, anything else is an operational read failure.
func readError(path string, err error) error {
	var parseErr *csv.ParseError
	if errors.As(err, &parseErr) {
		return &RejectError{Line: parseErr.Line, Reason: parseErr.Err.Error()}
	}
	var pathErr *fs.PathError
	if errors.As(err, &pathErr) {
		return err
	}
	return &fs.PathError{Op: "read", Path: path, Err: err}
}

// trimAll trims leading and trailing spaces from each value in place.
func trimAll(values []string) {
	for i, v := range values {
		values[i] = strings.Trim(v, " ")
	}
}
