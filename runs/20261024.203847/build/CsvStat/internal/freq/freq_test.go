package freq

import (
	"reflect"
	"testing"
)

func TestTopOrdersByCountThenValue(t *testing.T) {
	values := []string{"b", "a", "b", "c", "a", "d"}
	got := Top(values, 10)
	want := []Entry{{"a", 2}, {"b", 2}, {"c", 1}, {"d", 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Top = %v, want %v", got, want)
	}
}

func TestTopLimit(t *testing.T) {
	values := []string{"a", "a", "a", "b", "b", "c"}
	got := Top(values, 2)
	want := []Entry{{"a", 3}, {"b", 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Top = %v, want %v", got, want)
	}
}

func TestTopLimitLargerThanDistinct(t *testing.T) {
	got := Top([]string{"x", "y"}, 50)
	if len(got) != 2 {
		t.Errorf("Top returned %d entries, want 2", len(got))
	}
}

func TestTopSkipsNulls(t *testing.T) {
	got := Top([]string{"", "NA", "NULL", "n/a", "  ", "v"}, 5)
	want := []Entry{{"v", 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Top = %v, want %v", got, want)
	}
}

func TestTopTrimsBeforeCounting(t *testing.T) {
	got := Top([]string{"Globex", "  Globex  ", "Globex "}, 5)
	want := []Entry{{"Globex", 3}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Top = %v, want %v", got, want)
	}
}

func TestTopCaseSensitive(t *testing.T) {
	got := Top([]string{"Acme", "acme"}, 5)
	if len(got) != 2 {
		t.Errorf("Top = %v, want two distinct entries", got)
	}
}

func TestTopAllNullIsEmpty(t *testing.T) {
	if got := Top([]string{"", "NULL"}, 5); len(got) != 0 {
		t.Errorf("Top = %v, want none", got)
	}
}

func TestTopNegativeLimitIsEmpty(t *testing.T) {
	if got := Top([]string{"a"}, -1); len(got) != 0 {
		t.Errorf("Top = %v, want none", got)
	}
}
