package reader

import (
	"reflect"
	"strings"
	"testing"
)

func TestTrimSpaces(t *testing.T) {
	if got := Trim("  a b  "); got != "a b" {
		t.Errorf("Trim = %q, want %q", got, "a b")
	}
}

func TestTrimLeavesTabs(t *testing.T) {
	if got := Trim("\ta\t"); got != "\ta\t" {
		t.Errorf("Trim = %q, want tabs kept", got)
	}
}

func TestIsNullSpellings(t *testing.T) {
	nulls := []string{"", "NA", "na", "Na", "N/A", "n/a", "NULL", "null", "Null"}
	for _, v := range nulls {
		if !IsNull(v) {
			t.Errorf("IsNull(%q) = false, want true", v)
		}
	}
	for _, v := range []string{"NAN", "N A", "nil", "0", "NA ", " NA"} {
		if IsNull(v) {
			t.Errorf("IsNull(%q) = true, want false", v)
		}
	}
}

func TestTrimBeforeNull(t *testing.T) {
	if !IsNull(Trim("   ")) {
		t.Errorf("a value of only spaces is not null after trimming")
	}
}

func TestTrimTableHeaderAndRows(t *testing.T) {
	in, err := Read(strings.NewReader(" a , b \n x , 1 \n"))
	if err != nil {
		t.Fatalf("Read error: %v", err)
	}
	out := TrimTable(in)
	if !reflect.DeepEqual(out.Header, []string{"a", "b"}) {
		t.Errorf("Header = %q", out.Header)
	}
	if !reflect.DeepEqual(out.Rows, [][]string{{"x", "1"}}) {
		t.Errorf("Rows = %q", out.Rows)
	}
	if !reflect.DeepEqual(out.Lines, in.Lines) {
		t.Errorf("Lines = %v, want %v", out.Lines, in.Lines)
	}
}
