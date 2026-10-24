package reader

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
)

// Table is a parsed CSV file: the first record is the header and every later record is a row.
type Table struct {
	Header []string   // first record
	Rows   [][]string // later records, in file order
	Lines  []int      // 1-based input line on which each row starts; len(Lines) == len(Rows)
}

// Read parses r with encoding/csv. Field counts are not checked here; see Validate.
// An empty input returns a zero Table and a nil error.
//
// encoding/csv skips blank lines, but in a single-column file a blank line is a record holding
// one empty value. Blank lines are therefore made explicit (as "") before parsing, and are
// dropped again afterwards when the header has more than one column.
func Read(r io.Reader) (Table, error) {
	raw, err := io.ReadAll(r)
	if err != nil {
		return Table{}, err
	}
	data, blank := markBlankLines(raw)

	cr := csv.NewReader(bytes.NewReader(data))
	cr.FieldsPerRecord = -1

	var t Table
	for {
		rec, err := cr.Read()
		if errors.Is(err, io.EOF) {
			return t, nil
		}
		if err != nil {
			return Table{}, err
		}
		line, _ := cr.FieldPos(0)
		if t.Header == nil {
			if blank[line] {
				continue
			}
			t.Header = rec
			continue
		}
		if blank[line] && len(t.Header) != 1 {
			continue
		}
		t.Rows = append(t.Rows, rec)
		t.Lines = append(t.Lines, line)
	}
}

// markBlankLines inserts "" at the start of every blank line that lies outside a quoted field
// and returns the 1-based numbers of the lines it changed. Line numbers are unchanged.
func markBlankLines(raw []byte) ([]byte, map[int]bool) {
	blank := make(map[int]bool)
	out := make([]byte, 0, len(raw))
	line, inQuote, atStart := 1, false, true
	for i := 0; i < len(raw); i++ {
		c := raw[i]
		if atStart && !inQuote && (c == '\n' || (c == '\r' && i+1 < len(raw) && raw[i+1] == '\n')) {
			out = append(out, '"', '"')
			blank[line] = true
		}
		atStart = false
		switch c {
		case '"':
			inQuote = !inQuote
		case '\n':
			line++
			atStart = true
		}
		out = append(out, c)
	}
	return out, blank
}
