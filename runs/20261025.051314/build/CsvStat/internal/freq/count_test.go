package freq

import (
	"reflect"
	"testing"
)

func TestCountOrdersByCountThenValue(t *testing.T) {
	got := Count([]string{"b", "a", "b", "c", "a", "b", "z", "c", "B", "B"})
	want := []Entry{{"b", 3}, {"B", 2}, {"a", 2}, {"c", 2}, {"z", 1}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Count = %+v, want %+v", got, want)
	}
}

func TestCountSkipsNulls(t *testing.T) {
	got := Count([]string{"x", "NA", "null", "", "N/A", "x"})
	want := []Entry{{"x", 2}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Count = %+v, want %+v", got, want)
	}
}

func TestCountEmptyInput(t *testing.T) {
	if got := Count(nil); len(got) != 0 {
		t.Fatalf("Count(nil) = %+v, want no entries", got)
	}
	if got := Count([]string{"", "NA"}); len(got) != 0 {
		t.Fatalf("Count(all null) = %+v, want no entries", got)
	}
}

func TestCountKeepsCaseDistinct(t *testing.T) {
	got := Count([]string{"Globex", "globex"})
	if len(got) != 2 {
		t.Fatalf("Count = %+v, want 2 distinct entries", got)
	}
}
