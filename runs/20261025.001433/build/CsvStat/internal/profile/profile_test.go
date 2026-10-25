package profile

import (
	"reflect"
	"testing"
)

func TestIsNullSpellings(t *testing.T) {
	for _, v := range []string{"", "NA", "na", "N/A", "n/a", "NULL", "null", "Null"} {
		if !IsNull(v) {
			t.Errorf("IsNull(%q) = false, want true", v)
		}
	}
	for _, v := range []string{"NAN", "N A", "NULLS", "none", "0", "x"} {
		if IsNull(v) {
			t.Errorf("IsNull(%q) = true, want false", v)
		}
	}
}

func TestSummarizeCountsAndNulls(t *testing.T) {
	row := Summarize("v", []string{"apple", "", "Apple", "apple", "NA", "pear"})
	if row.Rows != 6 || row.Nulls != 2 || row.Distinct != 3 {
		t.Errorf("rows/nulls/distinct = %d/%d/%d, want 6/2/3", row.Rows, row.Nulls, row.Distinct)
	}
}

func TestSummarizeTypes(t *testing.T) {
	cases := []struct {
		name   string
		values []string
		want   string
	}{
		{"integer", []string{"1", "-2", "007", "NA"}, "integer"},
		{"number", []string{"1", "-2.50", "3.14159"}, "number"},
		{"date", []string{"2025-02-01", "2024-02-29", ""}, "date"},
		{"boolean", []string{"true", "False", "TRUE", "NA"}, "boolean"},
		{"text", []string{"1", "2025-02-01"}, "text"},
		{"empty", []string{"", "NA", "NULL"}, "empty"},
		{"no records", nil, "empty"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Summarize("c", tc.values).Type; got != tc.want {
				t.Errorf("Type = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestInferTypeRejectsLookalikes(t *testing.T) {
	for _, v := range []string{"+5", ".5", "5.", "1e3", "1.2.3", "0x10", "--1", "1-", "12abc"} {
		if got := Summarize("c", []string{"1", "2.5", v}).Type; got == "integer" || got == "number" {
			t.Errorf("value %q typed %q, want not numeric", v, got)
		}
	}
}

func TestInferTypeDateCalendar(t *testing.T) {
	for _, v := range []string{"2025-02-30", "2025-02-29", "2025-13-01", "2025-00-10", "2025-2-3", "20250203"} {
		if got := Summarize("c", []string{"2025-02-01", v}).Type; got == "date" {
			t.Errorf("value %q typed date, want not date", v)
		}
	}
}

func TestSummarizeIntegerStatistics(t *testing.T) {
	row := Summarize("v", []string{"4", "-7", "12", "100", "3"})
	want := []string{"v", "integer", "5", "0", "5", "-7", "100", "112", "22.40", "4.00"}
	if got := row.Fields(); !reflect.DeepEqual(got, want) {
		t.Errorf("Fields() = %q, want %q", got, want)
	}
}

func TestSummarizeEvenCountMedianAndNumberFormat(t *testing.T) {
	row := Summarize("v", []string{"0.01", "0.02"})
	if row.Type != "number" || row.Median != "0.02" || row.Mean != "0.02" {
		t.Errorf("number median/mean = %q/%q, want 0.02/0.02 (0.015 rounds half up)", row.Median, row.Mean)
	}
	row = Summarize("v", []string{"1", "2", "3", "4"})
	if row.Median != "2.50" {
		t.Errorf("Median = %q, want 2.50", row.Median)
	}
}

func TestSummarizeIntegerHugeIsExact(t *testing.T) {
	row := Summarize("v", []string{"1000000000000000000000000000001", "1000000000000000000000000000000"})
	if row.Min != "1000000000000000000000000000000" || row.Max != "1000000000000000000000000000001" {
		t.Errorf("min/max = %q/%q", row.Min, row.Max)
	}
	if row.Sum != "2000000000000000000000000000001" {
		t.Errorf("Sum = %q", row.Sum)
	}
}

func TestSummarizeNumberMinMaxAreNumeric(t *testing.T) {
	row := Summarize("v", []string{"15.27", "231.18", "-12.21", "0.50"})
	if row.Min != "-12.21" || row.Max != "231.18" {
		t.Errorf("min/max = %q/%q, want -12.21/231.18", row.Min, row.Max)
	}
}

func TestSummarizeDateMinMax(t *testing.T) {
	row := Summarize("d", []string{"2022-05-14", "2018-11-08", "2025-06-02"})
	if row.Min != "2018-11-08" || row.Max != "2025-06-02" {
		t.Errorf("min/max = %q/%q", row.Min, row.Max)
	}
	if row.Sum != "" || row.Mean != "" || row.Median != "" {
		t.Errorf("date numeric fields = %q/%q/%q, want blank", row.Sum, row.Mean, row.Median)
	}
}

func TestSummarizeBooleanMinMax(t *testing.T) {
	cases := []struct {
		values   []string
		min, max string
	}{
		{[]string{"True", "FALSE", "true"}, "false", "true"},
		{[]string{"True", "TRUE"}, "true", "true"},
		{[]string{"False", "false"}, "false", "false"},
	}
	for _, tc := range cases {
		row := Summarize("b", tc.values)
		if row.Min != tc.min || row.Max != tc.max {
			t.Errorf("%q: min/max = %q/%q, want %q/%q", tc.values, row.Min, row.Max, tc.min, tc.max)
		}
	}
}

func TestSummarizeTextMinMaxByByteValue(t *testing.T) {
	row := Summarize("t", []string{"banana", "Zebra", "apple", "mentor"})
	if row.Min != "Zebra" || row.Max != "mentor" {
		t.Errorf("min/max = %q/%q, want Zebra/mentor", row.Min, row.Max)
	}
}

func TestSummarizeEmptyColumnIsBlank(t *testing.T) {
	row := Summarize("e", []string{"", "NULL"})
	if row.Min != "" || row.Max != "" || row.Sum != "" || row.Mean != "" || row.Median != "" {
		t.Errorf("empty column fields not blank: %q", row.Fields())
	}
}

func TestFrequenciesOrderAndLimit(t *testing.T) {
	values := []string{"b", "a", "b", "c", "a", "d", "NA", ""}
	got := Frequencies(values, 3)
	want := []Count{{"a", 2}, {"b", 2}, {"c", 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Frequencies() = %v, want %v", got, want)
	}
}

func TestFrequenciesAllNullIsEmpty(t *testing.T) {
	if got := Frequencies([]string{"", "NA"}, 5); len(got) != 0 {
		t.Errorf("Frequencies() = %v, want none", got)
	}
}
