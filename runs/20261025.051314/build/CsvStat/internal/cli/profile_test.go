package cli

import (
	"bytes"
	"encoding/csv"
	"reflect"
	"testing"
)

// parseOutput parses stdout as RFC 4180 CSV, failing the test if it is malformed.
func parseOutput(t *testing.T, stdout *bytes.Buffer) [][]string {
	t.Helper()
	rows, err := csv.NewReader(stdout).ReadAll()
	if err != nil {
		t.Fatalf("stdout is not CSV: %v\n%s", err, stdout.String())
	}
	return rows
}

// profileOf runs the default invocation on content and returns the parsed rows.
func profileOf(t *testing.T, content string) [][]string {
	t.Helper()
	code, stdout, _ := runWith(t, content)
	if code != exitOK {
		t.Fatalf("exit code = %d, want %d", code, exitOK)
	}
	return parseOutput(t, stdout)
}

func TestProfileHeaderAndColumnOrder(t *testing.T) {
	rows := profileOf(t, "id,label,price\n1,a,1.50\n2,b,2.25\n4,c,3.00\n")

	wantHeader := []string{"column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"}
	if !reflect.DeepEqual(rows[0], wantHeader) {
		t.Fatalf("header = %q, want %q", rows[0], wantHeader)
	}
	var names, types []string
	for _, row := range rows[1:] {
		names = append(names, row[0])
		types = append(types, row[1])
	}
	if want := []string{"id", "label", "price"}; !reflect.DeepEqual(names, want) {
		t.Fatalf("columns = %q, want %q", names, want)
	}
	if want := []string{"integer", "text", "number"}; !reflect.DeepEqual(types, want) {
		t.Fatalf("types = %q, want %q", types, want)
	}
}

func TestProfileIntegerRow(t *testing.T) {
	rows := profileOf(t, "id,label,price\n1,a,1.50\n2,b,2.25\n4,c,3.00\n")

	want := []string{"id", "integer", "3", "0", "3", "1", "4", "7", "2.33", "2.00"}
	if !reflect.DeepEqual(rows[1], want) {
		t.Fatalf("integer row = %q, want %q", rows[1], want)
	}
}

func TestProfileNumberRow(t *testing.T) {
	rows := profileOf(t, "id,label,price\n1,a,1.50\n2,b,2.25\n4,c,3.00\n")

	want := []string{"price", "number", "3", "0", "3", "1.50", "3.00", "6.75", "2.25", "2.25"}
	if !reflect.DeepEqual(rows[3], want) {
		t.Fatalf("number row = %q, want %q", rows[3], want)
	}
}

func TestProfileTextRowHasBlankStatistics(t *testing.T) {
	rows := profileOf(t, "id,label,price\n1,a,1.50\n2,b,2.25\n4,c,3.00\n")

	want := []string{"label", "text", "3", "0", "3", "a", "c", "", "", ""}
	if !reflect.DeepEqual(rows[2], want) {
		t.Fatalf("text row = %q, want %q", rows[2], want)
	}
}

func TestProfileHeaderOnlyRowsAreEmpty(t *testing.T) {
	rows := profileOf(t, "sku,price\n")

	if len(rows) != 3 {
		t.Fatalf("got %d rows, want header plus 2 columns", len(rows))
	}
	want := []string{"sku", "empty", "0", "0", "0", "", "", "", "", ""}
	if !reflect.DeepEqual(rows[1], want) {
		t.Fatalf("header-only row = %q, want %q", rows[1], want)
	}
}

func TestProfileQuotedValueRoundTrips(t *testing.T) {
	// The value is padded and contains a comma and quotes; the trimmed value must come back intact.
	rows := profileOf(t, "note\n\"  x, \"\"y\"\"  \"\n")

	if got := rows[1][5]; got != `x, "y"` {
		t.Fatalf("min = %q, want %q", got, `x, "y"`)
	}
	if got := rows[1][6]; got != `x, "y"` {
		t.Fatalf("max = %q, want %q", got, `x, "y"`)
	}
}

func TestProfileNumberRoundsHalfAwayOnce(t *testing.T) {
	rows := profileOf(t, "pos,neg,tiny\n0.125,-0.125,0.001\n0.125,-0.125,-0.002\n")

	if got := rows[1][8]; got != "0.13" {
		t.Errorf("pos mean = %q, want 0.13", got)
	}
	if got := rows[2][8]; got != "-0.13" {
		t.Errorf("neg mean = %q, want -0.13", got)
	}
	if got := rows[3][7]; got != "0.00" {
		t.Errorf("tiny sum = %q, want 0.00", got)
	}
	if got := rows[3][8]; got != "0.00" {
		t.Errorf("tiny mean = %q, want 0.00", got)
	}
}

func TestProfileNullsCountedAndExcluded(t *testing.T) {
	rows := profileOf(t, "n\n1\nNA\n\"\"\n3\n")

	want := []string{"n", "integer", "4", "2", "2", "1", "3", "4", "2.00", "2.00"}
	if !reflect.DeepEqual(rows[1], want) {
		t.Fatalf("row = %q, want %q", rows[1], want)
	}
}

func TestProfileOutputEndsWithNewlineAndHasNoCarriageReturn(t *testing.T) {
	code, stdout, _ := runWith(t, "a,b\n1,x\n2,y\n")
	if code != exitOK {
		t.Fatalf("exit code = %d, want %d", code, exitOK)
	}
	if bytes.Contains(stdout.Bytes(), []byte("\r")) {
		t.Fatal("output contains a carriage return")
	}
	if stdout.Len() == 0 || stdout.Bytes()[stdout.Len()-1] != '\n' {
		t.Fatal("output does not end with a newline")
	}
	if rows := parseOutput(t, stdout); len(rows) != 3 {
		t.Fatalf("got %d rows, want header plus 2 columns", len(rows))
	}
}
