package profile

import (
	"reflect"
	"testing"
)

func TestIsNullEmpty(t *testing.T) {
	for _, v := range []string{"", " ", "\t", "   "} {
		if !IsNull(v) {
			t.Errorf("IsNull(%q) = false, want true", v)
		}
	}
}

func TestIsNullSpellings(t *testing.T) {
	for _, v := range []string{"NA", "na", "Na", "N/A", "n/a", "N/a", "NULL", "null", "Null", " NA ", " n/a\t", "  null  "} {
		if !IsNull(v) {
			t.Errorf("IsNull(%q) = false, want true", v)
		}
	}
}

func TestIsNullNonNull(t *testing.T) {
	for _, v := range []string{"NAN", "N/B", "NULLS", "none", "0", "nil", "N A", "x", "false"} {
		if IsNull(v) {
			t.Errorf("IsNull(%q) = true, want false", v)
		}
	}
}

func TestInferTypeInteger(t *testing.T) {
	cases := map[string][]string{
		"plain":         {"1", "22", "333"},
		"negatives":     {"-1", "0", "-75"},
		"leading zeros": {"007", "010", "-003"},
		"with nulls":    {"4", "", "NA", "-2", "null"},
	}
	for name, values := range cases {
		if got := InferType(values); got != TypeInteger {
			t.Errorf("%s: InferType(%q) = %q, want %q", name, values, got, TypeInteger)
		}
	}
}

func TestInferTypeNumber(t *testing.T) {
	cases := map[string][]string{
		"decimals":          {"1.50", "2.25"},
		"negative decimals": {"-2.25", "-0.5"},
		"mixed with ints":   {"3", "1.5", "-2"},
		"mixed with nulls":  {"1.5", "N/A", "", "2.0"},
	}
	for name, values := range cases {
		if got := InferType(values); got != TypeNumber {
			t.Errorf("%s: InferType(%q) = %q, want %q", name, values, got, TypeNumber)
		}
	}
}

func TestInferTypeNotNumeric(t *testing.T) {
	for _, v := range []string{"+5", ".5", "5.", "1e3", "1,000", "-", "1.2.3", "--1"} {
		got := InferType([]string{v})
		if got == TypeInteger || got == TypeNumber {
			t.Errorf("InferType(%q) = %q, want neither integer nor number", v, got)
		}
	}
}

func TestInferTypeDate(t *testing.T) {
	cases := map[string][]string{
		"valid dates":     {"2025-01-05", "2025-12-31"},
		"leap day":        {"2024-02-29", "2000-02-29"},
		"dates with nils": {"2025-03-01", "NA", "", "2025-02-14"},
	}
	for name, values := range cases {
		if got := InferType(values); got != TypeDate {
			t.Errorf("%s: InferType(%q) = %q, want %q", name, values, got, TypeDate)
		}
	}
}

func TestInferTypeInvalidDate(t *testing.T) {
	cases := map[string][]string{
		"impossible day":      {"2025-02-30", "2025-01-01"},
		"non-leap feb 29":     {"2025-02-29"},
		"april 31":            {"2025-04-31"},
		"single digit month":  {"2025-1-05"},
		"time suffix":         {"2025-03-01T08:00"},
		"slash separator":     {"2025/03/01"},
		"one bad among valid": {"2025-01-01", "2025-13-01"},
	}
	for name, values := range cases {
		if got := InferType(values); got != TypeText {
			t.Errorf("%s: InferType(%q) = %q, want %q", name, values, got, TypeText)
		}
	}
}

func TestInferTypeBoolean(t *testing.T) {
	cases := map[string][]string{
		"mixed case":   {"true", "FALSE", "False"},
		"with nulls":   {"TRUE", "", "NA", "false"},
		"single value": {"True"},
	}
	for name, values := range cases {
		if got := InferType(values); got != TypeBoolean {
			t.Errorf("%s: InferType(%q) = %q, want %q", name, values, got, TypeBoolean)
		}
	}
}

func TestInferTypeText(t *testing.T) {
	cases := map[string][]string{
		"words":            {"apple", "pear"},
		"booleans & words": {"true", "maybe"},
		"booleans & ints":  {"true", "1"},
		"punctuated":       {"a-b", "c.d!"},
		"yes and no":       {"yes", "no"},
		"ones and zeros":   {"1", "0", "t", "f"},
	}
	for name, values := range cases {
		if got := InferType(values); got != TypeText {
			t.Errorf("%s: InferType(%q) = %q, want %q", name, values, got, TypeText)
		}
	}
}

func TestInferTypeEmpty(t *testing.T) {
	cases := map[string][]string{
		"no values":         nil,
		"only empty":        {"", "  ", ""},
		"only null spelled": {"NA", "n/a", "NULL"},
	}
	for name, values := range cases {
		if got := InferType(values); got != TypeEmpty {
			t.Errorf("%s: InferType(%q) = %q, want %q", name, values, got, TypeEmpty)
		}
	}
}

func TestProfileColumnNumericFields(t *testing.T) {
	col := ProfileColumn("amt", []string{"1.50", "-2.25", "NA", "10.00"})
	want := Column{
		Name: "amt", Type: TypeNumber, Rows: 4, Nulls: 1, Distinct: 3,
		Min: "-2.25", Max: "10.00", Sum: "9.25", Mean: "3.08", Median: "1.50",
	}
	if !reflect.DeepEqual(col, want) {
		t.Errorf("ProfileColumn = %+v, want %+v", col, want)
	}
}

func TestProfileColumnNonNumericBlanksStatistics(t *testing.T) {
	col := ProfileColumn("when", []string{"2025-01-05", "2024-12-31"})
	if col.Sum != "" || col.Mean != "" || col.Median != "" {
		t.Errorf("date column statistics = (%q, %q, %q), want all blank", col.Sum, col.Mean, col.Median)
	}
}
