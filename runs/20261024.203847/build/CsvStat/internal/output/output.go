package output

import (
	"encoding/csv"
	"io"
	"strconv"

	"csvstat/internal/freq"
	"csvstat/internal/profile"
)

var (
	profileHeader = []string{"column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"}
	freqHeader    = []string{"value", "count"}
)

// WriteProfile writes the profile header and one row per column as RFC 4180 CSV with LF endings.
func WriteProfile(w io.Writer, cols []profile.Column) error {
	cw := csv.NewWriter(w)
	if err := cw.Write(profileHeader); err != nil {
		return err
	}
	for _, c := range cols {
		rec := []string{
			c.Name, c.Type,
			strconv.Itoa(c.Rows), strconv.Itoa(c.Nulls), strconv.Itoa(c.Distinct),
			c.Min, c.Max, c.Sum, c.Mean, c.Median,
		}
		if err := cw.Write(rec); err != nil {
			return err
		}
	}
	cw.Flush()
	return cw.Error()
}

// WriteFreq writes the frequency header and one row per entry as RFC 4180 CSV with LF endings.
func WriteFreq(w io.Writer, entries []freq.Entry) error {
	cw := csv.NewWriter(w)
	if err := cw.Write(freqHeader); err != nil {
		return err
	}
	for _, e := range entries {
		if err := cw.Write([]string{e.Value, strconv.Itoa(e.Count)}); err != nil {
			return err
		}
	}
	cw.Flush()
	return cw.Error()
}
