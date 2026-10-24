package reader

import (
	"encoding/csv"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func mustRead(t *testing.T, in string) Table {
	t.Helper()
	tbl, err := Read(strings.NewReader(in))
	if err != nil {
		t.Fatalf("Read(%q) error: %v", in, err)
	}
	return tbl
}

func TestReadHeaderThenRows(t *testing.T) {
	tbl := mustRead(t, "a,b\n1,2\n3,4\n")
	if !reflect.DeepEqual(tbl.Header, []string{"a", "b"}) {
		t.Errorf("Header = %q", tbl.Header)
	}
	want := [][]string{{"1", "2"}, {"3", "4"}}
	if !reflect.DeepEqual(tbl.Rows, want) {
		t.Errorf("Rows = %q, want %q", tbl.Rows, want)
	}
}

func TestReadQuotedComma(t *testing.T) {
	tbl := mustRead(t, "note\n\"remote, part-time\"\n")
	if len(tbl.Rows) != 1 || len(tbl.Rows[0]) != 1 || tbl.Rows[0][0] != "remote, part-time" {
		t.Errorf("Rows = %q, want one value %q", tbl.Rows, "remote, part-time")
	}
}

func TestReadInvalidUTF8NotRejected(t *testing.T) {
	_, err := Read(strings.NewReader("a\n\xff\xfe\n"))
	if err != nil {
		t.Errorf("Read returned error %v for invalid UTF-8, want nil", err)
	}
}

func TestReadRaggedRecordReturned(t *testing.T) {
	tbl, err := Read(strings.NewReader("a,b\n1,2,3\n"))
	if err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if len(tbl.Rows) != 1 || !reflect.DeepEqual(tbl.Rows[0], []string{"1", "2", "3"}) {
		t.Errorf("Rows = %q, want the extra-field record unchanged", tbl.Rows)
	}
}

func TestReadLineNumbers(t *testing.T) {
	tbl := mustRead(t, "a,b\n1,\"two\nlines\"\n3,4\n")
	want := []int{2, 4}
	if !reflect.DeepEqual(tbl.Lines, want) {
		t.Errorf("Lines = %v, want %v", tbl.Lines, want)
	}
}

func TestReadEmptyInput(t *testing.T) {
	tbl, err := Read(strings.NewReader(""))
	if err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if tbl.Header != nil || len(tbl.Rows) != 0 || len(tbl.Lines) != 0 {
		t.Errorf("Read(empty) = %+v, want zero Table", tbl)
	}
}

func TestReadHeaderOnlyHasNoRows(t *testing.T) {
	tbl := mustRead(t, "a,b,c\n")
	if len(tbl.Header) != 3 || len(tbl.Rows) != 0 {
		t.Errorf("Read(header only) = %+v", tbl)
	}
}

func TestReadSyntaxErrorReturnedUnchanged(t *testing.T) {
	_, err := Read(strings.NewReader("a,b\n1,x\"y\n"))
	var pe *csv.ParseError
	if !errors.As(err, &pe) {
		t.Errorf("Read error = %v, want a *csv.ParseError", err)
	}
}

func TestReadBlankLineInSingleColumnIsEmptyValue(t *testing.T) {
	tbl := mustRead(t, "fruit\napple\n\nNA\n")
	want := [][]string{{"apple"}, {""}, {"NA"}}
	if !reflect.DeepEqual(tbl.Rows, want) {
		t.Errorf("Rows = %q, want %q", tbl.Rows, want)
	}
	if !reflect.DeepEqual(tbl.Lines, []int{2, 3, 4}) {
		t.Errorf("Lines = %v", tbl.Lines)
	}
}

func TestReadBlankLineInMultiColumnIsSkipped(t *testing.T) {
	tbl := mustRead(t, "a,b\n1,2\n\n3,4\n\n")
	if len(tbl.Rows) != 2 {
		t.Errorf("Rows = %q, want the blank lines skipped", tbl.Rows)
	}
}

func TestReadBlankLineInsideQuotedFieldKept(t *testing.T) {
	tbl := mustRead(t, "a\n\"x\n\ny\"\n")
	if len(tbl.Rows) != 1 || tbl.Rows[0][0] != "x\n\ny" {
		t.Errorf("Rows = %q", tbl.Rows)
	}
}
