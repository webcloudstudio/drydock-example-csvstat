package csvio

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// writeFile stores content in a fresh temporary file and returns its path.
func writeFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "in.csv")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

// loadRejected loads content and asserts it is rejected with a RejectError, returning that error.
func loadRejected(t *testing.T, content string) *RejectError {
	t.Helper()
	_, err := Load(writeFile(t, content))
	var reject *RejectError
	if !errors.As(err, &reject) {
		t.Fatalf("Load(%q) error = %v, want *RejectError", content, err)
	}
	return reject
}

func TestTrimHeaderNames(t *testing.T) {
	table, err := Load(writeFile(t, "  alpha , beta  \n1,2\n"))
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"alpha", "beta"}
	if !reflect.DeepEqual(table.Header, want) {
		t.Errorf("Header = %q, want %q", table.Header, want)
	}
}

func TestTrimValues(t *testing.T) {
	table, err := Load(writeFile(t, "a,b\n  x ,\t y \n"))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := table.Column(0), []string{"x"}; !reflect.DeepEqual(got, want) {
		t.Errorf("column a = %q, want %q", got, want)
	}
	if got, want := table.Column(1), []string{"\t y"}; !reflect.DeepEqual(got, want) {
		t.Errorf("column b = %q, want %q (tabs are not trimmed)", got, want)
	}
}

func TestRejectEmptyFile(t *testing.T) {
	for _, content := range []string{"", "\n"} {
		reject := loadRejected(t, content)
		if reject.Line != 0 {
			t.Errorf("Load(%q) line = %d, want 0 for a whole-file rejection", content, reject.Line)
		}
	}
}

func TestRejectBadHeader(t *testing.T) {
	cases := map[string]int{
		",b\n1,2\n":   1, // empty name
		"a,a\n1,2\n":  1, // repeated
		"a, a\n1,2\n": 1, // repeated only after trimming
		"a, \n1,2\n":  1, // empty only after trimming
	}
	for content, line := range cases {
		reject := loadRejected(t, content)
		if reject.Line != line {
			t.Errorf("Load(%q) line = %d, want %d", content, reject.Line, line)
		}
	}
}

func TestRejectRaggedRecord(t *testing.T) {
	cases := map[string]int{
		"a,b\n1\n":        2, // short
		"a,b\n1,2,3\n":    2, // long
		"a,b\n1,2\n3,4,5": 3, // later line
	}
	for content, line := range cases {
		reject := loadRejected(t, content)
		if reject.Line != line {
			t.Errorf("Load(%q) line = %d, want %d", content, reject.Line, line)
		}
	}
}

func TestRejectCSVSyntaxError(t *testing.T) {
	reject := loadRejected(t, "a,b\n1,x\"y\n")
	if reject.Line != 2 {
		t.Errorf("line = %d, want 2", reject.Line)
	}
}

func TestHeaderOnlyFileIsValid(t *testing.T) {
	table, err := Load(writeFile(t, "sku,price\n"))
	if err != nil {
		t.Fatal(err)
	}
	if len(table.Records) != 0 {
		t.Errorf("Records = %d, want 0", len(table.Records))
	}
	if want := []string{"sku", "price"}; !reflect.DeepEqual(table.Header, want) {
		t.Errorf("Header = %q, want %q", table.Header, want)
	}
}

func TestByteOrderMarkIsKept(t *testing.T) {
	const bom = "\ufeff"
	table, err := Load(writeFile(t, bom+"a,b\n1,2\n"))
	if err != nil {
		t.Fatal(err)
	}
	if table.Header[0] != bom+"a" {
		t.Errorf("first header = %q, want BOM kept", table.Header[0])
	}
}

func TestInvalidUTF8IsNotRejected(t *testing.T) {
	if _, err := Load(writeFile(t, "a,b\n\xff,2\n")); err != nil {
		t.Errorf("Load() error = %v, want nil", err)
	}
}

func TestQuotedCommaIsOneValue(t *testing.T) {
	table, err := Load(writeFile(t, "a,b\n\"x, y\",2\n"))
	if err != nil {
		t.Fatal(err)
	}
	if got := table.Column(0)[0]; got != "x, y" {
		t.Errorf("value = %q, want %q", got, "x, y")
	}
}

func TestUnreadablePathIsReadError(t *testing.T) {
	_, err := Load(filepath.Join(t.TempDir(), "missing.csv"))
	var read *ReadError
	if !errors.As(err, &read) {
		t.Fatalf("Load(missing) error = %v, want *ReadError", err)
	}
}

func TestDirectoryIsReadError(t *testing.T) {
	_, err := Load(t.TempDir())
	var read *ReadError
	if !errors.As(err, &read) {
		t.Fatalf("Load(directory) error = %v, want *ReadError", err)
	}
}

func TestIndexIsExactAndCaseSensitive(t *testing.T) {
	table, err := Load(writeFile(t, "Alpha,beta\n1,2\n"))
	if err != nil {
		t.Fatal(err)
	}
	if i, ok := table.Index("beta"); !ok || i != 1 {
		t.Errorf("Index(beta) = %d, %v; want 1, true", i, ok)
	}
	if _, ok := table.Index("alpha"); ok {
		t.Errorf("Index(alpha) found, want absent (case-sensitive)")
	}
}
