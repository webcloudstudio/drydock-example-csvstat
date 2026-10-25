package csvio

import (
	"encoding/csv"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func writeFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "data.csv")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestReadTrimsHeaderAndValues(t *testing.T) {
	path := writeFile(t, "  id , name  \n  1 ,  Globex  \n")
	table, err := Read(path)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if want := []string{"id", "name"}; !reflect.DeepEqual(table.Header, want) {
		t.Errorf("Header = %q, want %q", table.Header, want)
	}
	if want := [][]string{{"1", "Globex"}}; !reflect.DeepEqual(table.Rows, want) {
		t.Errorf("Rows = %q, want %q", table.Rows, want)
	}
}

func TestReadKeepsQuotedCommaInOneField(t *testing.T) {
	path := writeFile(t, "a,b\n\"x, y\",2\n")
	table, err := Read(path)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if len(table.Rows) != 1 || len(table.Rows[0]) != 2 {
		t.Fatalf("Rows = %q, want one record of two fields", table.Rows)
	}
	if got := table.Rows[0][0]; got != "x, y" {
		t.Errorf("first field = %q, want %q", got, "x, y")
	}
}

func TestReadReturnsRecordLines(t *testing.T) {
	path := writeFile(t, "a,b\n1,2\n\n3,4\n")
	table, err := Read(path)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if want := []int{2, 4}; !reflect.DeepEqual(table.Lines, want) {
		t.Errorf("Lines = %v, want %v (blank line skipped, record lines kept)", table.Lines, want)
	}
	if len(table.Lines) != len(table.Rows) {
		t.Errorf("len(Lines) = %d, len(Rows) = %d, want equal", len(table.Lines), len(table.Rows))
	}
}

func TestReadAcceptsInvalidUTF8(t *testing.T) {
	path := writeFile(t, "a,b\n\xff,2\n")
	table, err := Read(path)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if got := table.Rows[0][0]; got != "\xff" {
		t.Errorf("first field = %q, want the invalid byte preserved", got)
	}
}

func TestReadMissingFileIsOpenError(t *testing.T) {
	path := filepath.Join(t.TempDir(), "no-such-file.csv")
	_, err := Read(path)
	if !errors.Is(err, ErrUnreadable) {
		t.Fatalf("Read error = %v, want ErrUnreadable", err)
	}
	var parseErr *csv.ParseError
	if errors.As(err, &parseErr) {
		t.Errorf("Read error %v is a CSV parse error, want an open error", err)
	}
}

func TestReadDirectoryIsUnreadable(t *testing.T) {
	_, err := Read(t.TempDir())
	if !errors.Is(err, ErrUnreadable) {
		t.Fatalf("Read(dir) error = %v, want ErrUnreadable", err)
	}
}

func TestReadHeaderOnlyIsValid(t *testing.T) {
	path := writeFile(t, "a,b\n")
	table, err := Read(path)
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	if len(table.Rows) != 0 || len(table.Header) != 2 {
		t.Errorf("got header %q and %d rows, want two columns and no rows", table.Header, len(table.Rows))
	}
}

func TestReadRejections(t *testing.T) {
	cases := map[string]string{
		"empty file":        "",
		"blank lines only":  "\n\n",
		"empty header name": "a,,c\n1,2,3\n",
		"repeated header":   "a, a ,c\n1,2,3\n",
		"short record":      "a,b,c\n1,2,3\n4,5\n",
		"long record":       "a,b\n1,2\n3,4,5\n",
		"bare quote":        "a,b\n1,\"2\n",
	}
	for name, content := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := Read(writeFile(t, content))
			if !errors.Is(err, ErrRejected) {
				t.Errorf("Read error = %v, want ErrRejected", err)
			}
		})
	}
}
