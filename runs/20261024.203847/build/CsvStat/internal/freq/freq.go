package freq

import (
	"sort"

	"csvstat/internal/reader"
)

// Entry is one distinct non-null value and how many times it occurs.
type Entry struct {
	Value string
	Count int
}

// Top returns at most n of the most frequent non-null values. Values are trimmed before
// counting. Entries are ordered by count descending, then value ascending by byte value.
func Top(values []string, n int) []Entry {
	counts := make(map[string]int)
	for _, raw := range values {
		v := reader.Trim(raw)
		if reader.IsNull(v) {
			continue
		}
		counts[v]++
	}
	entries := make([]Entry, 0, len(counts))
	for v, c := range counts {
		entries = append(entries, Entry{Value: v, Count: c})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Count != entries[j].Count {
			return entries[i].Count > entries[j].Count
		}
		return entries[i].Value < entries[j].Value
	})
	if n < 0 {
		n = 0
	}
	if n < len(entries) {
		entries = entries[:n]
	}
	return entries
}
