package csvdata

import (
	"encoding/csv"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func mustRead(t *testing.T, input string) *Table {
	t.Helper()
	table, err := Read(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Read(%q): unexpected error: %v", input, err)
	}
	return table
}

func TestReadHeaderAndRecords(t *testing.T) {
	table := mustRead(t, "name,qty\nwidget,2\ngadget,5\n")

	if want := []string{"name", "qty"}; !reflect.DeepEqual(table.Header, want) {
		t.Fatalf("Header = %q, want %q", table.Header, want)
	}
	got := make([][]string, len(table.Records))
	for i, rec := range table.Records {
		got[i] = rec.Fields
	}
	want := [][]string{{"widget", "2"}, {"gadget", "5"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("record fields = %q, want %q", got, want)
	}
}

func TestReadQuotedComma(t *testing.T) {
	table := mustRead(t, "name,note\n\"Smith, J\",ok\n")

	if len(table.Records) != 1 {
		t.Fatalf("got %d records, want 1", len(table.Records))
	}
	if want := []string{"Smith, J", "ok"}; !reflect.DeepEqual(table.Records[0].Fields, want) {
		t.Fatalf("Fields = %q, want %q", table.Records[0].Fields, want)
	}
}

func TestReadTrailingNewline(t *testing.T) {
	withNewline := mustRead(t, "a,b\n1,2\n")
	withoutNewline := mustRead(t, "a,b\n1,2")

	if !reflect.DeepEqual(withNewline, withoutNewline) {
		t.Fatalf("with newline = %+v, without = %+v; want equal", withNewline, withoutNewline)
	}
	if len(withNewline.Records) != 1 {
		t.Fatalf("got %d records, want 1", len(withNewline.Records))
	}
}

func TestReadRecordLineNumbers(t *testing.T) {
	// The first data record starts on line 2 and spans lines 2-3 through the
	// quoted newline; the next record starts on line 4.
	table := mustRead(t, "a,b\n1,\"multi\nline\"\n3,4\n")

	lines := make([]int, len(table.Records))
	for i, rec := range table.Records {
		lines[i] = rec.Line
	}
	if want := []int{2, 4}; !reflect.DeepEqual(lines, want) {
		t.Fatalf("record lines = %v, want %v", lines, want)
	}
}

func TestReadHeaderOnly(t *testing.T) {
	for _, input := range []string{"a,b\n", "a,b"} {
		table := mustRead(t, input)
		if want := []string{"a", "b"}; !reflect.DeepEqual(table.Header, want) {
			t.Errorf("Read(%q) Header = %q, want %q", input, table.Header, want)
		}
		if len(table.Records) != 0 {
			t.Errorf("Read(%q) got %d records, want 0", input, len(table.Records))
		}
	}
}

func TestReadNoRecords(t *testing.T) {
	for _, input := range []string{"", "\n\n\n"} {
		table := mustRead(t, input)
		if table.Header != nil {
			t.Errorf("Read(%q) Header = %q, want nil", input, table.Header)
		}
		if len(table.Records) != 0 {
			t.Errorf("Read(%q) got %d records, want 0", input, len(table.Records))
		}
	}
}

func TestReadBareQuoteReportsParseError(t *testing.T) {
	_, err := Read(strings.NewReader("a,b\n1\"x,2\n"))

	var parseErr *csv.ParseError
	if !errors.As(err, &parseErr) {
		t.Fatalf("Read error = %v, want a *csv.ParseError", err)
	}
	if parseErr.Line != 2 {
		t.Fatalf("ParseError.Line = %d, want 2", parseErr.Line)
	}
}
