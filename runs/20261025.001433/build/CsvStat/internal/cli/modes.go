package cli

import (
	"strconv"

	"csvstat/internal/csvio"
	"csvstat/internal/profile"
)

// Profiler renders the profile rows, header first, for the chosen columns of a loaded table.
type Profiler interface {
	Profile(t *csvio.Table, columns []int) [][]string
}

// FrequencyLister renders the value,count rows, header first, for one column of a loaded table.
type FrequencyLister interface {
	Frequencies(t *csvio.Table, column, top int) [][]string
}

// profileMode is the default Profiler.
type profileMode struct{}

func (profileMode) Profile(t *csvio.Table, columns []int) [][]string {
	rows := [][]string{profile.Header}
	for _, c := range columns {
		rows = append(rows, profile.Summarize(t.Header[c], t.Column(c)).Fields())
	}
	return rows
}

// frequencyMode is the default FrequencyLister.
type frequencyMode struct{}

func (frequencyMode) Frequencies(t *csvio.Table, column, top int) [][]string {
	rows := [][]string{{"value", "count"}}
	for _, c := range profile.Frequencies(t.Column(column), top) {
		rows = append(rows, []string{c.Value, strconv.Itoa(c.Count)})
	}
	return rows
}
