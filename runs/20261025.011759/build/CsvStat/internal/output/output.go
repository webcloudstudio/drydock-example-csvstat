// Package output writes profile and frequency results as CSV.
package output

import (
	"encoding/csv"
	"io"
	"strconv"

	"csvstat/internal/profile"
)

var profileHeader = []string{"column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"}

var frequencyHeader = []string{"value", "count"}

// WriteProfile writes the profile header followed by one row per column, in the order given.
func WriteProfile(w io.Writer, cols []profile.Column) error {
	cw := csv.NewWriter(w)
	if err := cw.Write(profileHeader); err != nil {
		return err
	}
	for _, c := range cols {
		record := []string{
			c.Name,
			c.Type,
			strconv.Itoa(c.Rows),
			strconv.Itoa(c.Nulls),
			strconv.Itoa(c.Distinct),
			c.Min,
			c.Max,
			c.Sum,
			c.Mean,
			c.Median,
		}
		if err := cw.Write(record); err != nil {
			return err
		}
	}
	cw.Flush()
	return cw.Error()
}

// WriteFrequencies writes the frequency header followed by one row per value, in the order given.
func WriteFrequencies(w io.Writer, freqs []profile.Frequency) error {
	cw := csv.NewWriter(w)
	if err := cw.Write(frequencyHeader); err != nil {
		return err
	}
	for _, f := range freqs {
		if err := cw.Write([]string{f.Value, strconv.Itoa(f.Count)}); err != nil {
			return err
		}
	}
	cw.Flush()
	return cw.Error()
}
