package profile

import (
	"reflect"
	"testing"
)

// wantMedian asserts the median of values as a rational string.
func wantMedian(t *testing.T, typ string, values []string, want string) {
	t.Helper()
	median, ok := Median(typ, values)
	if !ok {
		t.Fatalf("Median(%q, %q) ok = false, want true", typ, values)
	}
	if got := median.RatString(); got != want {
		t.Fatalf("Median(%q) = %s, want %s", values, got, want)
	}
}

func TestMedianOddCount(t *testing.T) {
	wantMedian(t, TypeInteger, []string{"5", "1", "3"}, "3")
}

func TestMedianEvenCountExact(t *testing.T) {
	// The mean of 0.1 and 0.2 is 3/20, which has no exact binary float form.
	wantMedian(t, TypeNumber, []string{"0.1", "0.2"}, "3/20")
	wantMedian(t, TypeInteger, []string{"1", "2", "3", "4"}, "5/2")
}

func TestMedianNumericOrder(t *testing.T) {
	// Lexically sorted this is -3, 10, 2, 9 and would give 6; numerically the middle pair is 2 and 9.
	wantMedian(t, TypeInteger, []string{"10", "9", "-3", "2"}, "11/2")
}

func TestMedianIgnoresNulls(t *testing.T) {
	wantMedian(t, TypeInteger, []string{"1", "NA", "3", "", "5"}, "3")
}

func TestMedianNonNumeric(t *testing.T) {
	for _, typ := range []string{TypeText, TypeDate, TypeBoolean, TypeEmpty} {
		if _, ok := Median(typ, []string{"1", "2", "3"}); ok {
			t.Errorf("Median(%q) ok = true, want false", typ)
		}
	}
	if _, ok := Median(TypeInteger, []string{"NA"}); ok {
		t.Error("Median with only nulls ok = true, want false")
	}
}

func TestMedianDoesNotReorderInput(t *testing.T) {
	values := []string{"3", "1", "2"}
	Median(TypeInteger, values)
	if want := []string{"3", "1", "2"}; !reflect.DeepEqual(values, want) {
		t.Fatalf("input reordered to %q, want %q", values, want)
	}
}
