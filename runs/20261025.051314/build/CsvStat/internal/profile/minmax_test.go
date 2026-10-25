package profile

import "testing"

func expectExtremes(t *testing.T, typ string, values []string, wantMin, wantMax string) Extremes {
	t.Helper()
	got := MinMax(typ, values)
	if got.Min != wantMin || got.Max != wantMax {
		t.Fatalf("MinMax(%q, %q) = (%q, %q), want (%q, %q)", typ, values, got.Min, got.Max, wantMin, wantMax)
	}
	return got
}

func TestMinMaxInteger(t *testing.T) {
	got := expectExtremes(t, TypeInteger, []string{"10", "9", "-3", "100"}, "-3", "100")
	if got.MinRat == nil || got.MaxRat == nil {
		t.Fatal("integer extremes have nil exact values")
	}
}

func TestMinMaxNumber(t *testing.T) {
	got := expectExtremes(t, TypeNumber, []string{"2.5", "10.25", "-0.5", "-0.25", "0.125"}, "-0.5", "10.25")
	if got.MinRat.RatString() != "-1/2" || got.MaxRat.RatString() != "41/4" {
		t.Fatalf("exact extremes = %s, %s; want -1/2, 41/4", got.MinRat.RatString(), got.MaxRat.RatString())
	}
}

func TestMinMaxDate(t *testing.T) {
	expectExtremes(t, TypeDate, []string{"2025-01-31", "2019-12-01", "2024-02-29"}, "2019-12-01", "2025-01-31")
}

func TestMinMaxBoolean(t *testing.T) {
	cases := []struct {
		name     string
		values   []string
		min, max string
	}{
		{"mixed case", []string{"TRUE", "false", "True"}, "false", "true"},
		{"all true", []string{"TRUE", "True"}, "true", "true"},
		{"all false", []string{"false", "FALSE"}, "false", "false"},
		{"mixed", []string{"False", "true"}, "false", "true"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expectExtremes(t, TypeBoolean, c.values, c.min, c.max)
		})
	}
}

func TestMinMaxText(t *testing.T) {
	// Byte order: upper case before lower case, space before letters, "apple" before "apple2".
	expectExtremes(t, TypeText, []string{"apple", "Banana", "a b", "Zed", "apple2"}, "Banana", "apple2")
}

func TestMinMaxEmpty(t *testing.T) {
	got := expectExtremes(t, TypeEmpty, []string{"NA", ""}, "", "")
	if got.MinRat != nil || got.MaxRat != nil {
		t.Fatalf("empty extremes have exact values: %+v", got)
	}
	expectExtremes(t, TypeEmpty, nil, "", "")
}

func TestMinMaxSingleValue(t *testing.T) {
	expectExtremes(t, TypeInteger, []string{"NA", "7"}, "7", "7")
}

func TestMinMaxIntegerLargeExact(t *testing.T) {
	// Beyond float64 precision the two values differ only in the last digit.
	expectExtremes(t, TypeInteger, []string{"9007199254740993", "9007199254740992"},
		"9007199254740992", "9007199254740993")
}
