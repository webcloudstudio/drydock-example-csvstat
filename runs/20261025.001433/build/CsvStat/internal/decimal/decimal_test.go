package decimal

import (
	"math/big"
	"testing"
)

func rat(t *testing.T, s string) *big.Rat {
	t.Helper()
	r, ok := new(big.Rat).SetString(s)
	if !ok {
		t.Fatalf("bad rational literal %q", s)
	}
	return r
}

func TestFormatTwoDecimals(t *testing.T) {
	cases := map[string]string{
		"3":          "3.00",
		"0":          "0.00",
		"1234567.89": "1234567.89",
		"0.5":        "0.50",
		"-3":         "-3.00",
		"100.1":      "100.10",
	}
	for in, want := range cases {
		if got := Format(rat(t, in)); got != want {
			t.Errorf("Format(%s) = %q, want %q", in, got, want)
		}
	}
}

func TestFormatRoundsHalfAwayFromZeroOnce(t *testing.T) {
	cases := map[string]string{
		"2.675":   "2.68",  // exact tie rounds up
		"-2.675":  "-2.68", // exact tie rounds away from zero
		"1.255":   "1.26",
		"0.005":   "0.01",
		"-0.005":  "-0.01",
		"1.0049":  "1.00",
		"2.674":   "2.67",
		"1/3":     "0.33",
		"-2/3":    "-0.67",
		"9.995":   "10.00",
		"-9.995":  "-10.00",
		"0.00499": "0.00",
	}
	for in, want := range cases {
		if got := Format(rat(t, in)); got != want {
			t.Errorf("Format(%s) = %q, want %q", in, got, want)
		}
	}
}

func TestFormatNegativeZeroHasNoSign(t *testing.T) {
	for _, in := range []string{"-0.001", "-0.004", "-0"} {
		if got := Format(rat(t, in)); got != "0.00" {
			t.Errorf("Format(%s) = %q, want 0.00", in, got)
		}
	}
}

func TestFormatHugeValueIsExact(t *testing.T) {
	r := new(big.Rat).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(30), nil))
	r.Add(r, big.NewRat(1, 4))
	if got, want := Format(r), "1000000000000000000000000000000.25"; got != want {
		t.Errorf("Format() = %q, want %q", got, want)
	}
}

func TestWholeWritesPlainInteger(t *testing.T) {
	if got := Whole(rat(t, "-42")); got != "-42" {
		t.Errorf("Whole(-42) = %q", got)
	}
	if got := Whole(rat(t, "007")); got != "7" {
		t.Errorf("Whole(007) = %q, want 7", got)
	}
}
