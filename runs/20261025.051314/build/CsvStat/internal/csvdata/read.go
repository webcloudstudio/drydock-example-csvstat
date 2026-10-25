package csvdata

import (
	"encoding/csv"
	"fmt"
	"io"
)

// Record is one data record and the 1-based source line on which it starts.
type Record struct {
	Line   int
	Fields []string
}

// Table is a parsed CSV input: the header record and the data records after it.
type Table struct {
	Header  []string
	Records []Record
}

// Read parses r as RFC 4180 CSV. The first record is the header and every later
// record is a data record. Values are returned as parsed, untrimmed. Input with
// no record returns a nil Header and no error; Validate rejects it.
func Read(r io.Reader) (*Table, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1
	cr.LazyQuotes = false

	var t Table
	header, err := cr.Read()
	if err == io.EOF {
		return &t, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read csv: %w", err)
	}
	t.Header = header

	for {
		fields, err := cr.Read()
		if err == io.EOF {
			return &t, nil
		}
		if err != nil {
			return nil, fmt.Errorf("read csv: %w", err)
		}
		line, _ := cr.FieldPos(0)
		t.Records = append(t.Records, Record{Line: line, Fields: fields})
	}
}
