package profile

import "sort"

// Frequency is one distinct non-null value of a column and how many times it occurs.
type Frequency struct {
	Value string
	Count int
}

// Frequencies lists the distinct non-null values ordered by count descending, then by value
// ascending by byte value.
func Frequencies(values []string) []Frequency {
	counts := make(map[string]int)
	for _, v := range nonNull(values) {
		counts[v]++
	}

	out := make([]Frequency, 0, len(counts))
	for value, count := range counts {
		out = append(out, Frequency{Value: value, Count: count})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Count != out[j].Count {
			return out[i].Count > out[j].Count
		}
		return out[i].Value < out[j].Value
	})
	return out
}
