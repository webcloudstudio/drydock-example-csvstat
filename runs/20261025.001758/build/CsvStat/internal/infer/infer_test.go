package infer

import "testing"

func TestClassifyEachType(t *testing.T) {
	tests := []struct {
		name   string
		values []string
		want   string
	}{
		{name: "integers", values: []string{"1", "-20", "0", "007"}, want: Integer},
		{name: "decimals", values: []string{"1.5", "-0.25", "3"}, want: Number},
		{name: "compact date is integer", values: []string{"20250115", "20250116"}, want: Integer},
		{name: "decimals only", values: []string{"0.5", "10.0"}, want: Number},
		{name: "dates", values: []string{"2025-01-15", "2024-02-29"}, want: Date},
		{name: "booleans", values: []string{"true", "FALSE", "True"}, want: Boolean},
		{name: "words", values: []string{"abc", "def"}, want: Text},
		{name: "no values", values: nil, want: Empty},
		{name: "empty slice", values: []string{}, want: Empty},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Classify(tt.values); got != tt.want {
				t.Errorf("Classify(%q) = %q, want %q", tt.values, got, tt.want)
			}
		})
	}
}

func TestClassifyPrecedence(t *testing.T) {
	tests := []struct {
		name   string
		values []string
		want   string
	}{
		{name: "zero and one are integer", values: []string{"0", "1", "1", "0"}, want: Integer},
		{name: "integer-shaped with decimal is number", values: []string{"1", "2.50"}, want: Number},
		{name: "number-shaped dates are not number", values: []string{"2025-01-01", "2025-01-02"}, want: Date},
		{name: "date-shaped boolean is text", values: []string{"true", "2025-01-01"}, want: Text},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Classify(tt.values); got != tt.want {
				t.Errorf("Classify(%q) = %q, want %q", tt.values, got, tt.want)
			}
		})
	}
}

func TestClassifyRejectsMalformedForms(t *testing.T) {
	tests := []struct {
		name   string
		values []string
	}{
		{name: "trailing dot", values: []string{"1.", "2"}},
		{name: "leading dot", values: []string{".5", "2"}},
		{name: "plus sign", values: []string{"+1", "2"}},
		{name: "exponent", values: []string{"1e3", "2"}},
		{name: "comma decimal", values: []string{"1,5", "2"}},
		{name: "lone minus", values: []string{"-", "2"}},
		{name: "impossible date", values: []string{"2025-02-30", "2025-01-01"}},
		{name: "non-leap feb 29", values: []string{"2023-02-29", "2025-03-01"}},
		{name: "month 13", values: []string{"2025-13-01", "2025-03-01"}},
		{name: "unpadded date", values: []string{"2025-1-5", "2025-03-01"}},
		{name: "yes and no", values: []string{"yes", "no"}},
		{name: "t and f", values: []string{"t", "f"}},
		{name: "boolean with stray word", values: []string{"true", "maybe"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Classify(tt.values); got != Text {
				t.Errorf("Classify(%q) = %q, want %q", tt.values, got, Text)
			}
		})
	}
}

func TestClassifyLeapDayValidOnlyInLeapYears(t *testing.T) {
	if got := Classify([]string{"2024-02-29", "2000-02-29"}); got != Date {
		t.Errorf("leap days = %q, want %q", got, Date)
	}
	if got := Classify([]string{"1900-02-29"}); got != Text {
		t.Errorf("1900-02-29 = %q, want %q", got, Text)
	}
}
