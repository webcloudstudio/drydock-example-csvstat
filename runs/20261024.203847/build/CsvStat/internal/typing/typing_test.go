package typing

import "testing"

func TestInferInteger(t *testing.T) {
	if got := Infer([]string{"1", "-20", "300", "0"}); got != "integer" {
		t.Errorf("Infer = %q, want integer", got)
	}
}

func TestInferNumber(t *testing.T) {
	if got := Infer([]string{"1", "2.5", "-3", "4.25"}); got != "number" {
		t.Errorf("Infer = %q, want number", got)
	}
}

func TestInferBooleanCaseInsensitive(t *testing.T) {
	if got := Infer([]string{"true", "FALSE", "True"}); got != "boolean" {
		t.Errorf("Infer = %q, want boolean", got)
	}
}

func TestInferBooleanMixedWithTextIsText(t *testing.T) {
	if got := Infer([]string{"true", "false", "yes"}); got != "text" {
		t.Errorf("Infer = %q, want text", got)
	}
}

func TestInferTextFallback(t *testing.T) {
	if got := Infer([]string{"1", "2", "x"}); got != "text" {
		t.Errorf("Infer = %q, want text", got)
	}
}

func TestInferNullsIgnored(t *testing.T) {
	if got := Infer([]string{"5", "", "NA", "n/a", "Null", "7"}); got != "integer" {
		t.Errorf("Infer = %q, want integer", got)
	}
}

func TestInferAllNullIsEmpty(t *testing.T) {
	if got := Infer([]string{"", "NA", "N/A", "NULL", "null", ""}); got != "empty" {
		t.Errorf("Infer = %q, want empty", got)
	}
}

func TestInferNoValuesIsEmpty(t *testing.T) {
	if got := Infer(nil); got != "empty" {
		t.Errorf("Infer(nil) = %q, want empty", got)
	}
}

func TestInferDateColumn(t *testing.T) {
	if got := Infer([]string{"2025-01-31", "2024-02-29", "2025-12-01"}); got != "date" {
		t.Errorf("Infer = %q, want date", got)
	}
}

func TestInferDateNeedsEveryValue(t *testing.T) {
	if got := Infer([]string{"2025-03-18", "2025-02-30"}); got != "text" {
		t.Errorf("Infer = %q, want text", got)
	}
}

func TestInferTimestampIsText(t *testing.T) {
	if got := Infer([]string{"2025-03-01T08:00", "2025-03-01T09:00"}); got != "text" {
		t.Errorf("Infer = %q, want text", got)
	}
}

func TestIsDateAcceptsLeapDay(t *testing.T) {
	if !IsDate("2024-02-29") {
		t.Errorf("IsDate(2024-02-29) = false, want true")
	}
}

func TestIsDateRejectsNonLeapDay(t *testing.T) {
	if IsDate("2025-02-29") {
		t.Errorf("IsDate(2025-02-29) = true, want false")
	}
}

func TestIsDateRejectsMonthAndDayBounds(t *testing.T) {
	for _, v := range []string{"2025-13-01", "2025-00-10", "2025-01-00", "2025-04-31"} {
		if IsDate(v) {
			t.Errorf("IsDate(%q) = true, want false", v)
		}
	}
}

func TestIsDateRejectsWrongShape(t *testing.T) {
	for _, v := range []string{"2025-3-1", "25-03-01", "2025/03/01", " 2025-03-01", "2025-03-01T00:00"} {
		if IsDate(v) {
			t.Errorf("IsDate(%q) = true, want false", v)
		}
	}
}

func TestInferIntegerBeforeNumberAndDate(t *testing.T) {
	if got := Infer([]string{"20250301"}); got != "integer" {
		t.Errorf("Infer = %q, want integer", got)
	}
}
