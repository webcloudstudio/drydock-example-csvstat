package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// freqRun runs --freq on the column of a file with content, followed by extra arguments.
func freqRun(t *testing.T, content, column string, extra ...string) (int, [][]string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "input.csv")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write input: %v", err)
	}
	args := append([]string{"--freq", column}, extra...)
	args = append(args, path)
	code, stdout, _ := runArgs(t, args...)
	if code != exitOK {
		return code, nil
	}
	return code, parseOutput(t, stdout)
}

// sevenValues returns a column whose values v0..v6 occur 7, 6, ..., 1 times.
func sevenValues() string {
	var b strings.Builder
	b.WriteString("v\n")
	for i := 0; i < 7; i++ {
		for j := 0; j < 7-i; j++ {
			fmt.Fprintf(&b, "v%d\n", i)
		}
	}
	return b.String()
}

func TestFreqListsCountsInOrder(t *testing.T) {
	_, rows := freqRun(t, "v\nb\na\nb\nc\na\nb\n", "v")

	want := [][]string{{"value", "count"}, {"b", "3"}, {"a", "2"}, {"c", "1"}}
	if !reflect.DeepEqual(rows, want) {
		t.Fatalf("rows = %q, want %q", rows, want)
	}
}

func TestFreqDefaultTopIsFive(t *testing.T) {
	_, rows := freqRun(t, sevenValues(), "v")

	if len(rows) != 1+5 {
		t.Fatalf("got %d rows, want header plus 5", len(rows))
	}
	if want := []string{"v0", "7"}; !reflect.DeepEqual(rows[1], want) {
		t.Fatalf("first row = %q, want %q", rows[1], want)
	}
	if want := []string{"v4", "3"}; !reflect.DeepEqual(rows[5], want) {
		t.Fatalf("fifth row = %q, want %q", rows[5], want)
	}
}

func TestFreqExplicitTopLimitsAfterOrdering(t *testing.T) {
	_, rows := freqRun(t, sevenValues(), "v", "--top", "2")

	want := [][]string{{"value", "count"}, {"v0", "7"}, {"v1", "6"}}
	if !reflect.DeepEqual(rows, want) {
		t.Fatalf("rows = %q, want %q", rows, want)
	}
}

func TestFreqTopLargerThanDistinctListsAll(t *testing.T) {
	_, rows := freqRun(t, "v\na\nb\n", "v", "--top", "50")

	if len(rows) != 3 {
		t.Fatalf("got %d rows, want header plus 2", len(rows))
	}
}

func TestFreqHeaderOnlyListsHeaderOnly(t *testing.T) {
	_, rows := freqRun(t, "sku,price\n", "sku")

	if want := [][]string{{"value", "count"}}; !reflect.DeepEqual(rows, want) {
		t.Fatalf("rows = %q, want %q", rows, want)
	}
}

func TestFreqAllNullColumnListsHeaderOnly(t *testing.T) {
	_, rows := freqRun(t, "id,gap\n1,\n2,NA\n3, \n", "gap", "--top", "3")

	if want := [][]string{{"value", "count"}}; !reflect.DeepEqual(rows, want) {
		t.Fatalf("rows = %q, want %q", rows, want)
	}
}

func TestFreqQuotedValueRoundTrips(t *testing.T) {
	_, rows := freqRun(t, "notes\n\"relocating, TX\"\n\"relocating, TX\"\nother\n", "notes")

	if want := []string{"relocating, TX", "2"}; !reflect.DeepEqual(rows[1], want) {
		t.Fatalf("row 1 = %q, want %q", rows[1], want)
	}
	if want := []string{"other", "1"}; !reflect.DeepEqual(rows[2], want) {
		t.Fatalf("row 2 = %q, want %q", rows[2], want)
	}
}

func TestFreqUnknownColumnExitsThreeWithEmptyStdout(t *testing.T) {
	path := filepath.Join(t.TempDir(), "input.csv")
	if err := os.WriteFile(path, []byte("a,b\n1,2\n"), 0o600); err != nil {
		t.Fatalf("write input: %v", err)
	}
	code, stdout, _ := runArgs(t, "--freq", "missing", path)
	if code != exitRejected {
		t.Fatalf("exit code = %d, want %d", code, exitRejected)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout has %d bytes, want none", stdout.Len())
	}
}

func TestFreqNameIsCaseSensitive(t *testing.T) {
	path := filepath.Join(t.TempDir(), "input.csv")
	if err := os.WriteFile(path, []byte("qty,b\n1,2\n"), 0o600); err != nil {
		t.Fatalf("write input: %v", err)
	}
	code, stdout, _ := runArgs(t, "--freq", "QTY", path)
	if code != exitRejected {
		t.Fatalf("exit code = %d, want %d", code, exitRejected)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout has %d bytes, want none", stdout.Len())
	}
}

func TestFreqRaggedInputExitsThreeWithEmptyStdout(t *testing.T) {
	path := filepath.Join(t.TempDir(), "ragged.csv")
	if err := os.WriteFile(path, []byte("a,b\n1,2\n3\n"), 0o600); err != nil {
		t.Fatalf("write input: %v", err)
	}
	code, stdout, _ := runArgs(t, "--freq", "a", path)
	if code != exitRejected {
		t.Fatalf("exit code = %d, want %d", code, exitRejected)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout has %d bytes, want none", stdout.Len())
	}
}
