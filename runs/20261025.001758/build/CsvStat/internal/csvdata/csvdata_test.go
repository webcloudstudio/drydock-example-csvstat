package csvdata

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// writeFile creates a file in a test-owned directory and returns its path.
func writeFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "input.csv")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}

func mustRead(t *testing.T, path string) Table {
	t.Helper()
	table, err := Read(path)
	if err != nil {
		t.Fatalf("Read(%s) error = %v", path, err)
	}
	return table
}

func TestReadTrimsHeaderAndValues(t *testing.T) {
	path := writeFile(t, " id , first name \n 1 ,  Ada Lovelace  \n   ,Grace\n")
	table := mustRead(t, path)

	wantHeader := []string{"id", "first name"}
	if !reflect.DeepEqual(table.Header, wantHeader) {
		t.Errorf("Header = %q, want %q", table.Header, wantHeader)
	}
	wantRecords := []Record{
		{Line: 2, Fields: []string{"1", "Ada Lovelace"}},
		{Line: 3, Fields: []string{"", "Grace"}},
	}
	if !reflect.DeepEqual(table.Records, wantRecords) {
		t.Errorf("Records = %+v, want %+v", table.Records, wantRecords)
	}
}

func TestReadQuotedFieldWithComma(t *testing.T) {
	path := writeFile(t, "id,notes\n1,\"relocating to Austin, TX\"\n2,\"say \"\"hi\"\"\"\n")
	table := mustRead(t, path)

	wantRecords := []Record{
		{Line: 2, Fields: []string{"1", "relocating to Austin, TX"}},
		{Line: 3, Fields: []string{"2", `say "hi"`}},
	}
	if !reflect.DeepEqual(table.Records, wantRecords) {
		t.Errorf("Records = %+v, want %+v", table.Records, wantRecords)
	}
}

func TestReadReturnsRaggedRecordsWithLineNumbers(t *testing.T) {
	path := writeFile(t, "a,b,c\n1,2,3\n4,5\n6,7,8,9\n")
	table := mustRead(t, path)

	if len(table.Records) != 3 {
		t.Fatalf("len(Records) = %d, want 3", len(table.Records))
	}
	wantWidths := []int{3, 2, 4}
	wantLines := []int{2, 3, 4}
	for i, rec := range table.Records {
		if got := len(rec.Fields); got != wantWidths[i] {
			t.Errorf("Records[%d] width = %d, want %d", i, got, wantWidths[i])
		}
		if rec.Line != wantLines[i] {
			t.Errorf("Records[%d].Line = %d, want %d", i, rec.Line, wantLines[i])
		}
	}
}

func TestReadEmptyFileHasNoHeaderOrRecords(t *testing.T) {
	table := mustRead(t, writeFile(t, ""))
	if len(table.Header) != 0 || len(table.Records) != 0 {
		t.Errorf("Read(empty) = %+v, want no header and no records", table)
	}
}

func TestReadSyntaxErrorIsRejectedWithLine(t *testing.T) {
	path := writeFile(t, "a,b\n1,2\n3,x\"y\n")
	_, err := Read(path)

	var reject *RejectError
	if !errors.As(err, &reject) {
		t.Fatalf("Read(bare quote) error = %v, want *RejectError", err)
	}
	if reject.Line != 3 {
		t.Errorf("RejectError.Line = %d, want 3", reject.Line)
	}
}

func TestReadMissingFileWrapsPathError(t *testing.T) {
	_, err := Read(filepath.Join(t.TempDir(), "absent.csv"))

	var pathErr *fs.PathError
	if !errors.As(err, &pathErr) {
		t.Fatalf("Read(missing) error = %v, want *fs.PathError", err)
	}
	var reject *RejectError
	if errors.As(err, &reject) {
		t.Errorf("Read(missing) error = %v, must not be a content rejection", err)
	}
}

func TestReadDirectoryWrapsPathError(t *testing.T) {
	_, err := Read(t.TempDir())

	var pathErr *fs.PathError
	if !errors.As(err, &pathErr) {
		t.Fatalf("Read(directory) error = %v, want *fs.PathError", err)
	}
	var reject *RejectError
	if errors.As(err, &reject) {
		t.Errorf("Read(directory) error = %v, must not be a content rejection", err)
	}
}
