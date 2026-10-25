package profile

import "testing"

func TestCountsRowsNullsDistinct(t *testing.T) {
	got := Counts([]string{"a", "NA", "a", "", "b", "n/a"})
	want := CountResult{Rows: 6, Nulls: 3, Distinct: 2}
	if got != want {
		t.Fatalf("Counts = %+v, want %+v", got, want)
	}
}

func TestCountsHeaderOnly(t *testing.T) {
	if got := Counts(nil); got != (CountResult{}) {
		t.Fatalf("Counts(nil) = %+v, want zero", got)
	}
	if got := Counts([]string{}); got != (CountResult{}) {
		t.Fatalf("Counts(empty) = %+v, want zero", got)
	}
}

func TestCountsDistinctExactStrings(t *testing.T) {
	got := Counts([]string{"Globex", "globex", "Globex", "GLOBEX"})
	want := CountResult{Rows: 4, Nulls: 0, Distinct: 3}
	if got != want {
		t.Fatalf("Counts = %+v, want %+v", got, want)
	}
}

func TestCountsAllNull(t *testing.T) {
	got := Counts([]string{"", "NA", "NULL"})
	want := CountResult{Rows: 3, Nulls: 3, Distinct: 0}
	if got != want {
		t.Fatalf("Counts = %+v, want %+v", got, want)
	}
}

func TestCountsNullNeverDistinct(t *testing.T) {
	if got := Counts([]string{"", "", "x"}); got.Distinct != 1 || got.Nulls != 2 {
		t.Fatalf("Counts = %+v, want Distinct 1 and Nulls 2", got)
	}
}
