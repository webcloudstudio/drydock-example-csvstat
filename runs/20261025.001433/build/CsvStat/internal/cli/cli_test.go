package cli

import (
	"bytes"
	"encoding/csv"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// writeCSV stores content in a fresh temporary file and returns its path.
func writeCSV(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "in.csv")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

// parseRecords reads stdout back as CSV so tests compare values, not text.
func parseRecords(t *testing.T, out string) [][]string {
	t.Helper()
	records, err := csv.NewReader(strings.NewReader(out)).ReadAll()
	if err != nil {
		t.Fatalf("stdout is not CSV: %v\n%s", err, out)
	}
	return records
}

func TestRunExitCodes(t *testing.T) {
	good := writeCSV(t, "a,b\n1,2\n")
	ragged := writeCSV(t, "a,b\n1\n")
	tests := []struct {
		name string
		args []string
		want int
	}{
		{name: "no arguments", args: nil, want: exitUsage},
		{name: "two files", args: []string{good, good}, want: exitUsage},
		{name: "unknown flag", args: []string{"--no-such-flag", good}, want: exitUsage},
		{name: "help flag", args: []string{"--help"}, want: exitOK},
		{name: "profile one file", args: []string{good}, want: exitOK},
		{name: "unreadable file", args: []string{filepath.Join(t.TempDir(), "missing.csv")}, want: exitRead},
		{name: "directory", args: []string{t.TempDir()}, want: exitRead},
		{name: "ragged file", args: []string{ragged}, want: exitReject},
		{name: "absent column", args: []string{"--column", "zz", good}, want: exitReject},
		{name: "absent freq name", args: []string{"--freq", "zz", good}, want: exitReject},
		{name: "usage before file access", args: []string{"--top", "3", "missing.csv"}, want: exitUsage},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			if got := Run(tt.args, &stdout, &stderr); got != tt.want {
				t.Errorf("Run(%q) = %d, want %d (stderr %q)", tt.args, got, tt.want, stderr.String())
			}
		})
	}
}

func TestRunWritesNothingToStdoutOnFailure(t *testing.T) {
	ragged := writeCSV(t, "a,b\n1\n")
	cases := [][]string{
		nil,
		{"--frobnicate", ragged},
		{filepath.Join(t.TempDir(), "missing.csv")},
		{ragged},
		{"--column", "zz", writeCSV(t, "a\n1\n")},
	}
	for _, args := range cases {
		var stdout, stderr bytes.Buffer
		if code := Run(args, &stdout, &stderr); code == exitOK {
			t.Fatalf("Run(%q) = 0, want a failure", args)
		}
		if stdout.Len() != 0 {
			t.Errorf("Run(%q) stdout = %q, want empty", args, stdout.String())
		}
	}
}

func TestRunHelpWritesUsageToStdoutOnly(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if code := Run([]string{"--help"}, &stdout, &stderr); code != exitOK {
		t.Fatalf("Run(--help) = %d, want 0", code)
	}
	if stderr.Len() != 0 {
		t.Errorf("stderr = %q, want empty", stderr.String())
	}
	for _, token := range []string{"csvstat", "--column", "--freq", "--top", "--help"} {
		if !strings.Contains(stdout.String(), token) {
			t.Errorf("usage does not name %q", token)
		}
	}
}

func TestRunProfilesEveryColumnInFileOrder(t *testing.T) {
	path := writeCSV(t, " id , name \n1, Ada \n2,NA\n")
	var stdout, stderr bytes.Buffer
	if code := Run([]string{path}, &stdout, &stderr); code != exitOK {
		t.Fatalf("Run() = %d, stderr %q", code, stderr.String())
	}
	records := parseRecords(t, stdout.String())
	want := [][]string{
		{"column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"},
		{"id", "integer", "2", "0", "2", "1", "2", "3", "1.50", "1.50"},
		{"name", "text", "2", "1", "1", "Ada", "Ada", "", "", ""},
	}
	if !reflect.DeepEqual(records, want) {
		t.Errorf("profile = %q, want %q", records, want)
	}
}

func TestRunProfileSingleColumn(t *testing.T) {
	path := writeCSV(t, "a,b\n1,x\n2,y\n")
	var stdout, stderr bytes.Buffer
	if code := Run([]string{"--column", "b", path}, &stdout, &stderr); code != exitOK {
		t.Fatalf("Run() = %d, stderr %q", code, stderr.String())
	}
	records := parseRecords(t, stdout.String())
	if len(records) != 2 || records[1][0] != "b" {
		t.Errorf("records = %q, want header plus the b row", records)
	}
}

func TestRunFrequencyListsAndLimits(t *testing.T) {
	path := writeCSV(t, "s\nnorth\nsouth\nnorth\neast\nNA\nsouth\nnorth\n")
	var stdout, stderr bytes.Buffer
	if code := Run([]string{"--freq", "s", "--top", "2", path}, &stdout, &stderr); code != exitOK {
		t.Fatalf("Run() = %d, stderr %q", code, stderr.String())
	}
	records := parseRecords(t, stdout.String())
	want := [][]string{{"value", "count"}, {"north", "3"}, {"south", "2"}}
	if !reflect.DeepEqual(records, want) {
		t.Errorf("frequencies = %q, want %q", records, want)
	}
}

func TestRunFrequencyDefaultLimitIsFive(t *testing.T) {
	path := writeCSV(t, "s\na\nb\nc\nd\ne\nf\ng\n")
	var stdout, stderr bytes.Buffer
	if code := Run([]string{"--freq", "s", path}, &stdout, &stderr); code != exitOK {
		t.Fatalf("Run() = %d, stderr %q", code, stderr.String())
	}
	if got := len(parseRecords(t, stdout.String())) - 1; got != 5 {
		t.Errorf("rows = %d, want 5", got)
	}
}

func TestRunFrequencyAllNullWritesHeaderOnly(t *testing.T) {
	path := writeCSV(t, "s,b\n1,\n2,NA\n")
	var stdout, stderr bytes.Buffer
	if code := Run([]string{"--freq", "b", path}, &stdout, &stderr); code != exitOK {
		t.Fatalf("Run() = %d, stderr %q", code, stderr.String())
	}
	if got, want := parseRecords(t, stdout.String()), [][]string{{"value", "count"}}; !reflect.DeepEqual(got, want) {
		t.Errorf("records = %q, want %q", got, want)
	}
}
