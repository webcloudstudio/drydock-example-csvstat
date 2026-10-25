package freq

import (
	"sort"

	"csvstat/internal/csvdata"
)

// Entry is one distinct non-null value and the number of records that carry it.
type Entry struct {
	Value string
	Count int
}

// Count takes the trimmed values of one column and returns its distinct non-null values in listing
// order: count descending, then value ascending by byte value. Null values are never listed.
func Count(values []string) []Entry {
	counts := make(map[string]int)
	for _, value := range values {
		if csvdata.IsNull(value) {
			continue
		}
		counts[value]++
	}

	entries := make([]Entry, 0, len(counts))
	for value, count := range counts {
		entries = append(entries, Entry{Value: value, Count: count})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Count != entries[j].Count {
			return entries[i].Count > entries[j].Count
		}
		return entries[i].Value < entries[j].Value
	})
	return entries
}
