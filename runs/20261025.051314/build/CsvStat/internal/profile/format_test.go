package profile

import (
	"math/big"
	"testing"
)

// rat parses a decimal literal exactly for use as a formatting input.
func rat(t *testing.T, literal string) *big.Rat {
	t.Helper()
	r, ok := new(big.Rat).SetString(literal)
	if !ok {
		t.Fatalf("bad test literal %q", literal)
	}
	return r
}

func expectDecimal(t *testing.T, literal, want string) {
	t.Helper()
	if got := FormatDecimal(rat(t, literal)); got != want {
		t.Errorf("FormatDecimal(%s) = %q, want %q", literal, got, want)
	}
}

func TestFormatDecimalTwoPlaces(t *testing.T) {
	expectDecimal(t, "3", "3.00")
	expectDecimal(t, "1.5", "1.50")
	expectDecimal(t, "0", "0.00")
	expectDecimal(t, "1.23456", "1.23")
}

func TestFormatDecimalHalfAwayFromZero(t *testing.T) {
	expectDecimal(t, "0.005", "0.01")
	expectDecimal(t, "-0.005", "-0.01")
	expectDecimal(t, "0.125", "0.13")
	expectDecimal(t, "-0.125", "-0.13")
	expectDecimal(t, "2.345", "2.35")
	expectDecimal(t, "1.004", "1.00")
	expectDecimal(t, "1.005", "1.01")
}

func TestFormatDecimalNoNegativeZero(t *testing.T) {
	expectDecimal(t, "-0.001", "0.00")
	expectDecimal(t, "-0.004", "0.00")
	expectDecimal(t, "-0", "0.00")
}

func TestFormatDecimalNoSeparator(t *testing.T) {
	expectDecimal(t, "1234.5", "1234.50")
	expectDecimal(t, "1234567.891", "1234567.89")
}

func TestFormatDecimalLargeExact(t *testing.T) {
	// The tie at the third decimal rounds up, keeping every digit of a value that float64 cannot hold.
	expectDecimal(t, "12345678901234567890.125", "12345678901234567890.13")
	expectDecimal(t, "-900719925474099312.345", "-900719925474099312.35")
}

func TestFormatInteger(t *testing.T) {
	cases := map[string]string{
		"0":       "0",
		"-0":      "0",
		"-42":     "-42",
		"1234567": "1234567",
	}
	for literal, want := range cases {
		if got := FormatInteger(rat(t, literal)); got != want {
			t.Errorf("FormatInteger(%s) = %q, want %q", literal, got, want)
		}
	}
}
