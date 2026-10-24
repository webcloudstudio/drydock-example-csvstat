package profile

import "testing"

func TestCountsRowsNullsDistinct(t *testing.T) {
	values := []string{"apple", "pear", "apple", "", "NA", "n/a", "Null", "fig", "pear", "apple"}
	rows, nulls, distinct := Counts(values)
	if rows != 10 || nulls != 4 || distinct != 3 {
		t.Errorf("Counts = %d,%d,%d want 10,4,3", rows, nulls, distinct)
	}
}

func TestCountsDistinctIsCaseSensitive(t *testing.T) {
	_, _, distinct := Counts([]string{"Acme", "acme"})
	if distinct != 2 {
		t.Errorf("distinct = %d, want 2", distinct)
	}
}

func TestCountsAllNull(t *testing.T) {
	rows, nulls, distinct := Counts([]string{"", "NA", "NULL"})
	if rows != 3 || nulls != 3 || distinct != 0 {
		t.Errorf("Counts = %d,%d,%d want 3,3,0", rows, nulls, distinct)
	}
}

func TestCountsHeaderOnlyIsZero(t *testing.T) {
	rows, nulls, distinct := Counts(nil)
	if rows != 0 || nulls != 0 || distinct != 0 {
		t.Errorf("Counts(nil) = %d,%d,%d want zeros", rows, nulls, distinct)
	}
}

func TestMinMaxIntegerBeyond64Bits(t *testing.T) {
	lo, hi := MinMax("integer", []string{"42", "-7", "12345678901234567890123", "-99"})
	if lo != "-99" || hi != "12345678901234567890123" {
		t.Errorf("MinMax = %q,%q", lo, hi)
	}
}

func TestMinMaxNumberRoundsHalfAwayFromZero(t *testing.T) {
	lo, hi := MinMax("number", []string{"15.27", "2.005", "100", "7.5"})
	if lo != "2.01" || hi != "100.00" {
		t.Errorf("MinMax = %q,%q want 2.01,100.00", lo, hi)
	}
}

func TestMinMaxNumberNegativeOrdersNumerically(t *testing.T) {
	lo, hi := MinMax("number", []string{"-12.21", "-0.44", "29.81", "6.81", "-9.02"})
	if lo != "-12.21" || hi != "29.81" {
		t.Errorf("MinMax = %q,%q", lo, hi)
	}
}

func TestMinMaxBoolean(t *testing.T) {
	lo, hi := MinMax("boolean", []string{"true", "False", "true"})
	if lo != "false" || hi != "true" {
		t.Errorf("MinMax = %q,%q want false,true", lo, hi)
	}
	lo, hi = MinMax("boolean", []string{"TRUE", "true"})
	if lo != "true" || hi != "true" {
		t.Errorf("MinMax all-true = %q,%q", lo, hi)
	}
	lo, hi = MinMax("boolean", []string{"false", "FALSE"})
	if lo != "false" || hi != "false" {
		t.Errorf("MinMax all-false = %q,%q", lo, hi)
	}
}

func TestMinMaxTextByteOrder(t *testing.T) {
	lo, hi := MinMax("text", []string{"mentor", "Zed", "remote", "apple"})
	if lo != "Zed" || hi != "remote" {
		t.Errorf("MinMax = %q,%q", lo, hi)
	}
}

func TestMinMaxDateAsWritten(t *testing.T) {
	lo, hi := MinMax("date", []string{"2025-03-18", "2015-08-27", "2025-06-02"})
	if lo != "2015-08-27" || hi != "2025-06-02" {
		t.Errorf("MinMax = %q,%q", lo, hi)
	}
}

func TestMinMaxExcludesNulls(t *testing.T) {
	lo, hi := MinMax("integer", []string{"", "9", "NA", "3"})
	if lo != "3" || hi != "9" {
		t.Errorf("MinMax = %q,%q", lo, hi)
	}
}

func TestMinMaxEmptyIsBlank(t *testing.T) {
	lo, hi := MinMax("empty", []string{"", "NA"})
	if lo != "" || hi != "" {
		t.Errorf("MinMax = %q,%q want blanks", lo, hi)
	}
}

func TestNumericStatsIntegerEvenMedian(t *testing.T) {
	sum, mean, median := NumericStats("integer", []string{"1", "2", "3", "10"})
	if sum != "16" || mean != "4.00" || median != "2.50" {
		t.Errorf("NumericStats = %q,%q,%q", sum, mean, median)
	}
}

func TestNumericStatsIntegerOddMedian(t *testing.T) {
	_, _, median := NumericStats("integer", []string{"9", "1", "5", "100", "3"})
	if median != "5.00" {
		t.Errorf("median = %q, want 5.00", median)
	}
}

func TestNumericStatsExactLargeSum(t *testing.T) {
	sum, _, _ := NumericStats("integer", []string{"9223372036854775807", "9223372036854775807", "9223372036854775807"})
	if sum != "27670116110564327421" {
		t.Errorf("sum = %q", sum)
	}
}

func TestNumericStatsNumberRounding(t *testing.T) {
	sum, mean, median := NumericStats("number", []string{"0.125", "-12.21", "29.33", "", "N/A", "2.41", "0.005", "20.98"})
	if sum != "40.64" || mean != "6.77" || median != "1.27" {
		t.Errorf("NumericStats = %q,%q,%q", sum, mean, median)
	}
}

func TestNumericStatsHalfAwayFromZeroNegative(t *testing.T) {
	sum, _, _ := NumericStats("number", []string{"-1.005", "-2.000"})
	if sum != "-3.01" {
		t.Errorf("sum = %q, want -3.01", sum)
	}
}

func TestNumericStatsNegativeRoundingToZeroHasNoSign(t *testing.T) {
	sum, _, _ := NumericStats("number", []string{"-0.001"})
	if sum != "0.00" {
		t.Errorf("sum = %q, want 0.00", sum)
	}
}

func TestNumericStatsBlankForOtherTypes(t *testing.T) {
	for _, typ := range []string{"text", "boolean", "date", "empty"} {
		sum, mean, median := NumericStats(typ, []string{"1", "2"})
		if sum != "" || mean != "" || median != "" {
			t.Errorf("NumericStats(%s) = %q,%q,%q want blanks", typ, sum, mean, median)
		}
	}
}

func TestDescribeFullColumn(t *testing.T) {
	got := Describe("qty", []string{"1", "", "10", "NA"})
	want := Column{
		Name: "qty", Type: "integer", Rows: 4, Nulls: 2, Distinct: 2,
		Min: "1", Max: "10", Sum: "11", Mean: "5.50", Median: "5.50",
	}
	if got != want {
		t.Errorf("Describe = %+v, want %+v", got, want)
	}
}
