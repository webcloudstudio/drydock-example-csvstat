package profile

import "testing"

func expectType(t *testing.T, want string, values ...string) {
	t.Helper()
	if got := InferType(values); got != want {
		t.Fatalf("InferType(%q) = %q, want %q", values, got, want)
	}
}

func TestInferIntegerColumn(t *testing.T) {
	expectType(t, TypeInteger, "0", "-12", "345")
}

func TestInferNumberColumn(t *testing.T) {
	expectType(t, TypeNumber, "1", "2.5", "-0.75")
}

func TestIntegerPreferredOverNumber(t *testing.T) {
	expectType(t, TypeInteger, "10", "20", "-30")
	expectType(t, TypeInteger, "10")
}

func TestNullsIgnoredWhenInferringNumeric(t *testing.T) {
	expectType(t, TypeNumber, "1", "NA", "", "N/A", "null", "2.5")
	expectType(t, TypeInteger, "7", "null", "NULL", "")
}

func TestMalformedNumericsAreNotNumeric(t *testing.T) {
	for _, value := range []string{"+1", "1.", ".5", "1e3", "1,5", "0x10", "-"} {
		t.Run(value, func(t *testing.T) {
			expectType(t, TypeText, value)
		})
	}
	expectType(t, TypeText, "1", "+1")
}

func TestInferDateColumn(t *testing.T) {
	expectType(t, TypeDate, "2025-01-31", "2024-02-29", "1999-12-31")
}

func TestInvalidCalendarDateIsNotDate(t *testing.T) {
	expectType(t, TypeText, "2025-02-30")
	expectType(t, TypeText, "2025-13-01")
}

func TestLeapDayHandling(t *testing.T) {
	expectType(t, TypeDate, "2024-02-29")
	expectType(t, TypeText, "2025-02-29")
}

func TestDateShapeIsStrict(t *testing.T) {
	for _, value := range []string{"2025-2-3", "2025/02/03", "2025-02-03T10:00"} {
		t.Run(value, func(t *testing.T) {
			expectType(t, TypeText, value)
		})
	}
	// All digits: integer is tested before date, so this is integer, never date.
	expectType(t, TypeInteger, "20250203")
}

func TestNullsIgnoredWhenInferringDate(t *testing.T) {
	expectType(t, TypeDate, "2025-01-01", "NA", "", "n/a")
}

func TestInferBooleanColumnMixedCase(t *testing.T) {
	expectType(t, TypeBoolean, "true", "FALSE", "False", "TRUE")
	expectType(t, TypeBoolean, "NA", "True")
}

func TestInferTextFallback(t *testing.T) {
	expectType(t, TypeText, "apple", "2025-01-01")
	expectType(t, TypeText, "1", "x")
}

func TestInferEmptyWhenAllNull(t *testing.T) {
	expectType(t, TypeEmpty, "NA", "", "null")
}

func TestInferEmptyWhenNoRows(t *testing.T) {
	expectType(t, TypeEmpty)
	if got := InferType([]string{}); got != TypeEmpty {
		t.Fatalf("InferType(empty slice) = %q, want %q", got, TypeEmpty)
	}
}

func TestBooleanLookalikesAreText(t *testing.T) {
	for _, value := range []string{"yes", "no", "t", "f", "y", "n", "on", "off"} {
		t.Run(value, func(t *testing.T) {
			expectType(t, TypeText, value)
		})
	}
}

func TestTypeOrderPrecedence(t *testing.T) {
	expectType(t, TypeInteger, "0", "1")
	expectType(t, TypeNumber, "1", "2.5")
	expectType(t, TypeDate, "2025-01-01")
	expectType(t, TypeBoolean, "true")
	expectType(t, TypeText, "2025-01-01", "true")
}
