package output

import (
	"bytes"
	"encoding/csv"
	"reflect"
	"strings"
	"testing"

	"csvstat/internal/freq"
	"csvstat/internal/profile"
)

func readBack(t *testing.T, s string) [][]string {
	t.Helper()
	recs, err := csv.NewReader(strings.NewReader(s)).ReadAll()
	if err != nil {
		t.Fatalf("output is not valid CSV: %v\n%q", err, s)
	}
	return recs
}

func TestWriteProfileHeaderAndRow(t *testing.T) {
	var buf bytes.Buffer
	cols := []profile.Column{{Name: "qty", Type: "integer", Rows: 2, Nulls: 1, Distinct: 1, Min: "3", Max: "3", Sum: "3", Mean: "3.00", Median: "3.00"}}
	if err := WriteProfile(&buf, cols); err != nil {
		t.Fatalf("WriteProfile error: %v", err)
	}
	want := [][]string{
		{"column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"},
		{"qty", "integer", "2", "1", "1", "3", "3", "3", "3.00", "3.00"},
	}
	if got := readBack(t, buf.String()); !reflect.DeepEqual(got, want) {
		t.Errorf("rows = %q, want %q", got, want)
	}
}

func TestWriteProfileLineEndingsAreLF(t *testing.T) {
	var buf bytes.Buffer
	_ = WriteProfile(&buf, []profile.Column{{Name: "a", Type: "empty"}})
	if strings.Contains(buf.String(), "\r") || !strings.HasSuffix(buf.String(), "\n") {
		t.Errorf("output %q, want LF endings only", buf.String())
	}
}

func TestWriteProfileQuotesComma(t *testing.T) {
	var buf bytes.Buffer
	_ = WriteProfile(&buf, []profile.Column{{Name: "notes", Type: "text", Min: "remote, part-time", Max: "remote, part-time"}})
	rows := readBack(t, buf.String())
	if rows[1][5] != "remote, part-time" {
		t.Errorf("min = %q, want the comma value kept in one field", rows[1][5])
	}
}

func TestWriteProfileEmptyStatisticsAreEmptyFields(t *testing.T) {
	var buf bytes.Buffer
	_ = WriteProfile(&buf, []profile.Column{{Name: "b", Type: "text"}})
	rows := readBack(t, buf.String())
	for _, i := range []int{5, 6, 7, 8, 9} {
		if rows[1][i] != "" {
			t.Errorf("field %d = %q, want empty", i, rows[1][i])
		}
	}
}

func TestWriteFreqHeaderAndRows(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteFreq(&buf, []freq.Entry{{Value: "Sales", Count: 8}}); err != nil {
		t.Fatalf("WriteFreq error: %v", err)
	}
	want := [][]string{{"value", "count"}, {"Sales", "8"}}
	if got := readBack(t, buf.String()); !reflect.DeepEqual(got, want) {
		t.Errorf("rows = %q, want %q", got, want)
	}
}

func TestWriteFreqHeaderOnlyWhenEmpty(t *testing.T) {
	var buf bytes.Buffer
	_ = WriteFreq(&buf, nil)
	if buf.String() != "value,count\n" {
		t.Errorf("output = %q, want header only", buf.String())
	}
}

func TestWriteFreqQuotesComma(t *testing.T) {
	var buf bytes.Buffer
	_ = WriteFreq(&buf, []freq.Entry{{Value: "relocating to Austin, TX", Count: 2}})
	rows := readBack(t, buf.String())
	if !reflect.DeepEqual(rows[1], []string{"relocating to Austin, TX", "2"}) {
		t.Errorf("row = %q", rows[1])
	}
}
