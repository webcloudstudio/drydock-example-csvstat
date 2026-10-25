package profile

import "testing"

// wantSumMean asserts the exact sum and mean rendered as rationals.
func wantSumMean(t *testing.T, typ string, values []string, wantSum, wantMean string) {
	t.Helper()
	sum, mean, ok := SumMean(typ, values)
	if !ok {
		t.Fatalf("SumMean(%q, %q) ok = false, want true", typ, values)
	}
	if got := sum.RatString(); got != wantSum {
		t.Errorf("sum = %s, want %s", got, wantSum)
	}
	if got := mean.RatString(); got != wantMean {
		t.Errorf("mean = %s, want %s", got, wantMean)
	}
}

func TestSumIntegerExact(t *testing.T) {
	// The total exceeds float64 precision, and includes negatives.
	wantSumMean(t, TypeInteger,
		[]string{"-5", "10", "25", "1000000000000000000000"},
		"1000000000000000000030", "500000000000000000015/2")
}

func TestSumNumberNoFloatDrift(t *testing.T) {
	// In float64, 0.1 + 0.2 is not 0.3; exactly it is three tenths.
	wantSumMean(t, TypeNumber, []string{"0.1", "0.2"}, "3/10", "3/20")
}

func TestMeanExactRational(t *testing.T) {
	wantSumMean(t, TypeInteger, []string{"1", "1", "2"}, "4", "4/3")
}

func TestSumMeanIgnoresNulls(t *testing.T) {
	// Two non-null values: the divisor is 2, not the row count of 5.
	wantSumMean(t, TypeInteger, []string{"2", "NA", "", "4", "null"}, "6", "3")
}

func TestSumMeanNonNumeric(t *testing.T) {
	for _, typ := range []string{TypeText, TypeDate, TypeBoolean, TypeEmpty} {
		if _, _, ok := SumMean(typ, []string{"1", "2"}); ok {
			t.Errorf("SumMean(%q) ok = true, want false", typ)
		}
	}
}

func TestSumMeanNoNonNullValues(t *testing.T) {
	if _, _, ok := SumMean(TypeInteger, []string{"NA", ""}); ok {
		t.Fatal("SumMean with only nulls ok = true, want false")
	}
	if _, _, ok := SumMean(TypeInteger, nil); ok {
		t.Fatal("SumMean with no values ok = true, want false")
	}
}

func TestSumMeanRejectsNonDecimalText(t *testing.T) {
	// A fraction form parses as big.Rat but is not a decimal number.
	if _, _, ok := SumMean(TypeNumber, []string{"1/2"}); ok {
		t.Fatal("SumMean(\"1/2\") ok = true, want false")
	}
}
