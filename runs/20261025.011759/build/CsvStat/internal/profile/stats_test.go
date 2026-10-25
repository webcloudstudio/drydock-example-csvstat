package profile

import (
	"math/big"
	"testing"
)

func TestCountsRows(t *testing.T) {
	rows, _, _ := countValues([]string{"1", "NA", "", "3"})
	if rows != 4 {
		t.Errorf("rows = %d, want 4", rows)
	}
}

func TestCountsNulls(t *testing.T) {
	_, nulls, _ := countValues([]string{"", "NA", "N/A", "NULL", "null", "x", "NAN", "0"})
	if nulls != 5 {
		t.Errorf("nulls = %d, want 5", nulls)
	}
}

func TestCountsDistinct(t *testing.T) {
	_, _, distinct := countValues([]string{"true", "TRUE", "true", "NA", "", "false"})
	if distinct != 3 {
		t.Errorf("distinct = %d, want 3 (true, TRUE, false; nulls excluded)", distinct)
	}
}

func TestCountsHeaderOnly(t *testing.T) {
	rows, nulls, distinct := countValues(nil)
	if rows != 0 || nulls != 0 || distinct != 0 {
		t.Errorf("counts = (%d, %d, %d), want (0, 0, 0)", rows, nulls, distinct)
	}
}

func TestSumInteger(t *testing.T) {
	sum, mean, _ := numericStats(TypeInteger, []string{"10", "-3", "5"})
	if sum != "12" {
		t.Errorf("sum = %q, want %q (plain integer)", sum, "12")
	}
	if mean != "4.00" {
		t.Errorf("mean = %q, want %q", mean, "4.00")
	}
}

func TestSumNumber(t *testing.T) {
	sum, _, _ := numericStats(TypeNumber, []string{"0.10", "0.20", "0.30"})
	if sum != "0.60" {
		t.Errorf("sum = %q, want %q", sum, "0.60")
	}
	// The exact sum 2.009 rounds once, to 2.01.
	sum, _, _ = numericStats(TypeNumber, []string{"1.005", "1.004"})
	if sum != "2.01" {
		t.Errorf("sum = %q, want %q", sum, "2.01")
	}
}

func TestMeanRounding(t *testing.T) {
	// 1 + 2 + 2 = 5 and 5/3 = 1.666..., which rounds to 1.67.
	_, mean, _ := numericStats(TypeInteger, []string{"1", "2", "2"})
	if mean != "1.67" {
		t.Errorf("mean = %q, want %q", mean, "1.67")
	}
	// The exact mean 0.125 is half a hundredth and rounds away from zero to 0.13.
	_, mean, _ = numericStats(TypeNumber, []string{"0.12", "0.13"})
	if mean != "0.13" {
		t.Errorf("mean = %q, want %q", mean, "0.13")
	}
}

func TestMeanNegative(t *testing.T) {
	// The exact mean -1.005 rounds away from zero to -1.01.
	_, mean, _ := numericStats(TypeNumber, []string{"-1.01", "-1.00"})
	if mean != "-1.01" {
		t.Errorf("mean = %q, want %q", mean, "-1.01")
	}
	_, mean, _ = numericStats(TypeInteger, []string{"-1", "0"})
	if mean != "-0.50" {
		t.Errorf("mean = %q, want %q", mean, "-0.50")
	}
	_, mean, _ = numericStats(TypeNumber, []string{"-0.001", "0.001"})
	if mean != "0.00" {
		t.Errorf("mean = %q, want %q (no negative zero)", mean, "0.00")
	}
}

func TestSumMeanNonNumeric(t *testing.T) {
	cases := []struct {
		typ    string
		values []string
	}{
		{TypeDate, []string{"2025-01-01"}},
		{TypeBoolean, []string{"true"}},
		{TypeText, []string{"word"}},
		{TypeEmpty, []string{"NA"}},
	}
	for _, c := range cases {
		col := ProfileColumn("c", c.values)
		if col.Type != c.typ {
			t.Fatalf("column type = %q, want %q", col.Type, c.typ)
		}
		if col.Sum != "" || col.Mean != "" || col.Median != "" {
			t.Errorf("%s statistics = (%q, %q, %q), want all blank", c.typ, col.Sum, col.Mean, col.Median)
		}
	}
}

func TestMedianOdd(t *testing.T) {
	_, _, median := numericStats(TypeInteger, []string{"9", "1", "5"})
	if median != "5.00" {
		t.Errorf("median = %q, want %q", median, "5.00")
	}
}

func TestMedianEven(t *testing.T) {
	_, _, median := numericStats(TypeInteger, []string{"1", "2", "3", "10"})
	if median != "2.50" {
		t.Errorf("median = %q, want %q", median, "2.50")
	}
}

func TestMedianUnsorted(t *testing.T) {
	// Lexical order would place "10" before "2"; numeric order places 9 in the middle.
	_, _, median := numericStats(TypeInteger, []string{"10", "2", "9"})
	if median != "9.00" {
		t.Errorf("median = %q, want %q (numeric sort)", median, "9.00")
	}
}

func TestMedianNumber(t *testing.T) {
	// The exact median of 1.00 and 2.005 is 1.5025, which rounds to 1.50.
	_, _, median := numericStats(TypeNumber, []string{"1.00", "2.005"})
	if median != "1.50" {
		t.Errorf("median = %q, want %q", median, "1.50")
	}
	// The exact median of 0.01 and 0.02 is 0.015, which rounds away from zero to 0.02.
	_, _, median = numericStats(TypeNumber, []string{"0.01", "0.02"})
	if median != "0.02" {
		t.Errorf("median = %q, want %q", median, "0.02")
	}
}

func TestMedianNonNumeric(t *testing.T) {
	col := ProfileColumn("word", []string{"pear", "apple", "fig"})
	if col.Median != "" {
		t.Errorf("median = %q, want blank for text", col.Median)
	}
}

func TestMinMaxInteger(t *testing.T) {
	lo, hi := minMax(TypeInteger, []string{"-3", "12", "7"})
	if lo != "-3" || hi != "12" {
		t.Errorf("min,max = %q,%q, want -3,12 (numeric order)", lo, hi)
	}
}

func TestMinMaxNumber(t *testing.T) {
	lo, hi := minMax(TypeNumber, []string{"2.5", "-10.125", "9"})
	if lo != "-10.13" || hi != "9.00" {
		t.Errorf("min,max = %q,%q, want -10.13,9.00", lo, hi)
	}
}

func TestMinMaxDate(t *testing.T) {
	lo, hi := minMax(TypeDate, []string{"2025-03-01", "2024-12-31", "2025-02-14"})
	if lo != "2024-12-31" || hi != "2025-03-01" {
		t.Errorf("min,max = %q,%q, want 2024-12-31,2025-03-01", lo, hi)
	}
}

func TestMinMaxBoolean(t *testing.T) {
	cases := []struct {
		values []string
		lo, hi string
	}{
		{[]string{"true", "TRUE"}, "true", "true"},
		{[]string{"false", "True"}, "false", "true"},
		{[]string{"FALSE"}, "false", "false"},
	}
	for _, c := range cases {
		lo, hi := minMax(TypeBoolean, c.values)
		if lo != c.lo || hi != c.hi {
			t.Errorf("minMax(%q) = %q,%q, want %q,%q", c.values, lo, hi, c.lo, c.hi)
		}
	}
}

func TestMinMaxText(t *testing.T) {
	lo, hi := minMax(TypeText, []string{"pear", "Apple", "fig", "apple"})
	if lo != "Apple" || hi != "pear" {
		t.Errorf("min,max = %q,%q, want Apple,pear (byte order)", lo, hi)
	}
}

func TestMinMaxEmpty(t *testing.T) {
	lo, hi := minMax(TypeEmpty, nil)
	if lo != "" || hi != "" {
		t.Errorf("min,max = %q,%q, want blanks", lo, hi)
	}
}

func TestFormatRatHalfAwayFromZero(t *testing.T) {
	cases := []struct {
		name string
		in   *big.Rat
		want string
	}{
		{"half rounds up", big.NewRat(1, 200), "0.01"},
		{"exact half of a hundredth", big.NewRat(125, 1000), "0.13"},
		{"negative half rounds away", big.NewRat(-125, 1000), "-0.13"},
		{"whole number", big.NewRat(2, 1), "2.00"},
		{"tiny negative has no sign", big.NewRat(-1, 1000), "0.00"},
	}
	for _, c := range cases {
		if got := FormatRat(c.in); got != c.want {
			t.Errorf("%s: FormatRat(%s) = %q, want %q", c.name, c.in.RatString(), got, c.want)
		}
	}
}
