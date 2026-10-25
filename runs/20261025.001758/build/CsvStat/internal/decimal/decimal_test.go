package decimal

import (
	"math/big"
	"testing"
)

func rat(t *testing.T, s string) *big.Rat {
	t.Helper()
	r, ok := new(big.Rat).SetString(s)
	if !ok {
		t.Fatalf("SetString(%q) failed", s)
	}
	return r
}

func TestFormatRoundsHalfAwayFromZero(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"0.005", "0.01"},
		{"-0.005", "-0.01"},
		{"0.004", "0.00"},
		{"-0.004", "0.00"},
		{"0.994", "0.99"},
		{"0.995", "1.00"},
		{"-0.995", "-1.00"},
		{"2.675", "2.68"},
		{"-2.675", "-2.68"},
		{"1.005", "1.01"},
		{"10.125", "10.13"},
		{"-10.125", "-10.13"},
	}
	for _, tt := range tests {
		if got := Format(rat(t, tt.in)); got != tt.want {
			t.Errorf("Format(%s) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestFormatNeverWritesNegativeZero(t *testing.T) {
	for _, in := range []string{"0", "-0", "0.0", "-0.0", "-0.001", "-0.004"} {
		if got := Format(rat(t, in)); got != "0.00" {
			t.Errorf("Format(%s) = %q, want 0.00", in, got)
		}
	}
}

func TestFormatPadsToTwoDecimals(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"0", "0.00"},
		{"7", "7.00"},
		{"-3", "-3.00"},
		{"1.5", "1.50"},
		{"-0.1", "-0.10"},
	}
	for _, tt := range tests {
		if got := Format(rat(t, tt.in)); got != tt.want {
			t.Errorf("Format(%s) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestFormatExactBeyondFloat(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"12345678901234567890.12", "12345678901234567890.12"},
		{"123456789012345678.995", "123456789012345679.00"},
		{"9007199254740993.001", "9007199254740993.00"},
		{"-9007199254740993.005", "-9007199254740993.01"},
	}
	for _, tt := range tests {
		if got := Format(rat(t, tt.in)); got != tt.want {
			t.Errorf("Format(%s) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestFormatSumOfTenthsIsExact(t *testing.T) {
	total := new(big.Rat)
	for i := 0; i < 10; i++ {
		total.Add(total, rat(t, "0.1"))
	}
	if got := Format(total); got != "1.00" {
		t.Errorf("Format(sum of ten 0.1) = %q, want 1.00", got)
	}
}
