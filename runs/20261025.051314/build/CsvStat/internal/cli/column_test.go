package cli

import (
	"bytes"
	"encoding/csv"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// columnFile writes content to a temporary file and returns its path.
func columnFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "input.csv")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write input: %v", err)
	}
	return path
}

// columnRowsOf runs --column on a file with content and returns the parsed rows.
func columnRowsOf(t *testing.T, content, name string) (int, [][]string) {
	t.Helper()
	code, stdout, _ := runArgs(t, "--column", name, columnFile(t, content))
	if code != exitOK {
		return code, nil
	}
	rows, err := csv.NewReader(bytes.NewReader(stdout.Bytes())).ReadAll()
	if err != nil {
		t.Fatalf("stdout is not CSV: %v", err)
	}
	return code, rows
}

const columnFixture = "id,price,label\n1,1.50,a\n2,2.25,b\n4,3.00,c\n"

func TestColumnWritesHeaderAndOneRow(t *testing.T) {
	_, rows := columnRowsOf(t, columnFixture, "price")
	if len(rows) != 2 {
		t.Fatalf("got %d rows, want header plus one", len(rows))
	}
	if want := profileHeader; !reflect.DeepEqual(rows[0], want) {
		t.Fatalf("header = %q, want %q", rows[0], want)
	}
	want := []string{"price", "number", "3", "0", "3", "1.50", "3.00", "6.75", "2.25", "2.25"}
	if !reflect.DeepEqual(rows[1], want) {
		t.Fatalf("row = %q, want %q", rows[1], want)
	}
}

func TestColumnRowMatchesFullProfile(t *testing.T) {
	full := profileOf(t, columnFixture)
	for _, row := range full[1:] {
		_, rows := columnRowsOf(t, columnFixture, row[0])
		if !reflect.DeepEqual(rows[1], row) {
			t.Errorf("column %q row = %q, want %q", row[0], rows[1], row)
		}
	}
}

func TestColumnUnknownNameExitsThreeWithEmptyStdout(t *testing.T) {
	code, stdout, _ := runArgs(t, "--column", "nosuch", columnFile(t, columnFixture))
	if code != exitRejected {
		t.Fatalf("exit code = %d, want %d", code, exitRejected)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout has %d bytes, want none", stdout.Len())
	}
}

func TestColumnNameIsCaseSensitive(t *testing.T) {
	code, _ := columnRowsOf(t, columnFixture, "PRICE")
	if code != exitRejected {
		t.Fatalf("exit code = %d, want %d", code, exitRejected)
	}
}

func TestColumnTrimmedHeaderNameMatches(t *testing.T) {
	_, rows := columnRowsOf(t, "  qty  ,other\n1,x\n", "qty")
	if len(rows) != 2 || rows[1][0] != "qty" || rows[1][1] != "integer" {
		t.Fatalf("rows = %q, want trimmed qty integer row", rows)
	}
}

func TestColumnRaggedInputExitsThreeWithEmptyStdout(t *testing.T) {
	code, stdout, _ := runArgs(t, "--column", "a", columnFile(t, "a,b\n1,2\n3,4,5\n"))
	if code != exitRejected {
		t.Fatalf("exit code = %d, want %d", code, exitRejected)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout has %d bytes, want none", stdout.Len())
	}
}
