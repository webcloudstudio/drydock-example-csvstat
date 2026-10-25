package cli

import (
	"fmt"
	"strconv"

	"csvstat/internal/csvdata"
	"csvstat/internal/freq"
)

// defaultTop is the number of frequency rows listed when --top is not given.
const defaultTop = 5

// freqRows returns the header and the limited frequency listing of the --freq column.
func freqRows(t *csvdata.Table, opts Options) ([][]string, error) {
	c := columnIndex(t.Header, opts.Freq)
	if c < 0 {
		return nil, fmt.Errorf("no column named %q", opts.Freq)
	}

	limit := defaultTop
	if opts.TopSet {
		limit = opts.Top
	}
	entries := freq.Top(freq.Count(columnValues(t, c)), limit)

	rows := [][]string{{"value", "count"}}
	for _, entry := range entries {
		rows = append(rows, []string{entry.Value, strconv.Itoa(entry.Count)})
	}
	return rows, nil
}

// columnIndex returns the position of the header name, or -1 when no column has that name.
func columnIndex(header []string, name string) int {
	for i, column := range header {
		if column == name {
			return i
		}
	}
	return -1
}
