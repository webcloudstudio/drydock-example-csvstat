package profile

import "sort"

// Count is one distinct non-null value and how many times it occurs.
type Count struct {
	Value string
	Count int
}

// Frequencies returns at most top distinct non-null values, most frequent first. Values with equal
// counts are ordered ascending by byte value. Null values are never listed.
func Frequencies(values []string, top int) []Count {
	counts := make(map[string]int)
	for _, v := range values {
		if !IsNull(v) {
			counts[v]++
		}
	}

	out := make([]Count, 0, len(counts))
	for v, n := range counts {
		out = append(out, Count{Value: v, Count: n})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Count != out[j].Count {
			return out[i].Count > out[j].Count
		}
		return out[i].Value < out[j].Value
	})
	if len(out) > top {
		out = out[:top]
	}
	return out
}
