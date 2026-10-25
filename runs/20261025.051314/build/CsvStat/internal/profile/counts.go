package profile

import "csvstat/internal/csvdata"

// CountResult holds the three counts reported for one column.
type CountResult struct {
	Rows     int
	Nulls    int
	Distinct int
}

// Counts takes every trimmed value of one column, in file order, and returns its counts.
func Counts(values []string) CountResult {
	result := CountResult{Rows: len(values)}
	distinct := make(map[string]struct{})
	for _, value := range values {
		if csvdata.IsNull(value) {
			result.Nulls++
			continue
		}
		distinct[value] = struct{}{}
	}
	result.Distinct = len(distinct)
	return result
}
