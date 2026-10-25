// Package csvio reads, trims, and structurally validates the input CSV file.
package csvio

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// ErrUnreadable marks a path that cannot be opened or read, including a directory.
var ErrUnreadable = errors.New("cannot read file")

// ErrRejected marks input that was read but is not a valid CSV table.
var ErrRejected = errors.New("rejected input")

// Table is a trimmed CSV file: the header and its data records in file order.
type Table struct {
	Header []string
	Rows   [][]string
	// Lines holds the 1-based starting line of each data record.
	Lines []int
}

// Read parses the CSV file at path. Open and read failures wrap ErrUnreadable; an empty,
// malformed, or ragged file wraps ErrRejected.
func Read(path string) (*Table, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnreadable, err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %v", ErrUnreadable, path, err)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("%w: %s is a directory", ErrUnreadable, path)
	}
	return parse(path, f)
}

func parse(name string, r io.Reader) (*Table, error) {
	cr := csv.NewReader(r)
	// Field-count mismatches are checked here so the diagnostic can name the line.
	cr.FieldsPerRecord = -1

	header, err := cr.Read()
	if errors.Is(err, io.EOF) {
		return nil, rejectedf(name, "no header record, file is empty")
	}
	if err != nil {
		return nil, classify(name, err)
	}

	table := &Table{Header: trimAll(header)}
	seen := make(map[string]bool, len(table.Header))
	for i, h := range table.Header {
		if h == "" {
			return nil, rejectedf(name, "header column %d is empty", i+1)
		}
		if seen[h] {
			return nil, rejectedf(name, "header name %q is repeated", h)
		}
		seen[h] = true
	}

	for {
		record, err := cr.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, classify(name, err)
		}
		line, _ := cr.FieldPos(0)
		if len(record) != len(table.Header) {
			return nil, rejectedf(name, "line %d has %d fields, header has %d", line, len(record), len(table.Header))
		}
		table.Rows = append(table.Rows, trimAll(record))
		table.Lines = append(table.Lines, line)
	}
	return table, nil
}

// classify maps an encoding/csv error to a rejection when it is a syntax error, and to an
// unreadable file when it came from the underlying reader.
func classify(name string, err error) error {
	var parseErr *csv.ParseError
	if errors.As(err, &parseErr) {
		return fmt.Errorf("%w: %s: %v", ErrRejected, name, err)
	}
	return fmt.Errorf("%w: %s: %v", ErrUnreadable, name, err)
}

func rejectedf(name, format string, args ...any) error {
	return fmt.Errorf("%w: %s: %s", ErrRejected, name, fmt.Sprintf(format, args...))
}

func trimAll(values []string) []string {
	out := make([]string, len(values))
	for i, v := range values {
		out[i] = strings.TrimSpace(v)
	}
	return out
}
