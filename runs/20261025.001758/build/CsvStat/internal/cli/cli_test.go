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

// writeFile creates a file in a test-owned directory and returns its path.
func writeFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "input.csv")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}

func TestRunUsageErrors(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{name: "no arguments", args: nil},
		{name: "two files", args: []string{"a.csv", "b.csv"}},
		{name: "unknown option", args: []string{"--bogus", "a.csv"}},
		{name: "option missing value", args: []string{"--column"}},
		{name: "freq missing value", args: []string{"--freq"}},
		{name: "top zero with freq", args: []string{"--freq", "name", "--top", "0", "a.csv"}},
		{name: "top negative with freq", args: []string{"--freq", "name", "--top=-1", "a.csv"}},
		{name: "top not an integer", args: []string{"--freq", "name", "--top", "abc", "a.csv"}},
		{name: "top without freq", args: []string{"--top", "3", "a.csv"}},
		{name: "column with freq", args: []string{"--freq", "name", "--column", "x", "a.csv"}},
		{name: "freq without file", args: []string{"--freq", "name"}},
		{name: "option after file", args: []string{"a.csv", "--column", "x"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := Run(tt.args, &stdout, &stderr)
			if code != ExitUsage {
				t.Fatalf("Run(%q) = %d, want %d", tt.args, code, ExitUsage)
			}
			if stdout.Len() != 0 {
				t.Errorf("Run(%q) wrote %q to stdout, want nothing", tt.args, stdout.String())
			}
			if n := bytes.Count(stderr.Bytes(), []byte("\n")); n != 1 {
				t.Errorf("Run(%q) wrote %d stderr lines, want 1", tt.args, n)
			}
		})
	}
}

func TestRunHelpExitsOK(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := Run([]string{"-h"}, &stdout, &stderr)
	if code != ExitOK {
		t.Fatalf("Run(-h) = %d, want %d", code, ExitOK)
	}
	if stdout.Len() == 0 {
		t.Error("Run(-h) wrote nothing to stdout, want usage")
	}
}

func TestRunValidUsageIsNotUsageError(t *testing.T) {
	valid := [][]string{
		{"a.csv"},
		{"--column", "name", "a.csv"},
		{"--freq", "name", "--top", "5", "a.csv"},
		{"--freq", "name", "a.csv"},
	}
	for _, args := range valid {
		var stdout, stderr bytes.Buffer
		code := Run(args, &stdout, &stderr)
		if code == ExitUsage {
			t.Errorf("Run(%q) = %d, want a non-usage code", args, code)
		}
	}
}

func TestRunValidFileExitsZero(t *testing.T) {
	path := writeFile(t, " id , name \n 1 , Ada \n 2 , Grace \n")
	var stdout, stderr bytes.Buffer
	code := Run([]string{path}, &stdout, &stderr)
	if code != ExitOK {
		t.Fatalf("Run(valid) = %d, want %d; stderr %q", code, ExitOK, stderr.String())
	}
	if stderr.Len() != 0 {
		t.Errorf("Run(valid) wrote %q to stderr, want nothing", stderr.String())
	}
}

func TestRunMissingFileExitsOneWithoutStdout(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no-such-file.csv")
	var stdout, stderr bytes.Buffer
	code := Run([]string{missing}, &stdout, &stderr)
	if code != ExitOperational {
		t.Fatalf("Run(missing) = %d, want %d", code, ExitOperational)
	}
	if stdout.Len() != 0 {
		t.Errorf("Run(missing) wrote %q to stdout, want nothing", stdout.String())
	}
}

func TestRunDirectoryExitsOneWithoutStdout(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := Run([]string{t.TempDir()}, &stdout, &stderr)
	if code != ExitOperational {
		t.Fatalf("Run(directory) = %d, want %d", code, ExitOperational)
	}
	if stdout.Len() != 0 {
		t.Errorf("Run(directory) wrote %q to stdout, want nothing", stdout.String())
	}
}

func TestRunRejectedContentExitsThreeWithoutStdout(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{name: "empty file", content: ""},
		{name: "blank content", content: "\n\n"},
		{name: "empty header name", content: "a,,c\n1,2,3\n"},
		{name: "repeated header name", content: "a,b,a\n1,2,3\n"},
		{name: "ragged record", content: "a,b,c\n1,2\n"},
		{name: "bare quote", content: "a,b\n1,x\"y\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := writeFile(t, tt.content)
			var stdout, stderr bytes.Buffer
			code := Run([]string{path}, &stdout, &stderr)
			if code != ExitRejected {
				t.Fatalf("Run(%s) = %d, want %d", tt.name, code, ExitRejected)
			}
			if stdout.Len() != 0 {
				t.Errorf("Run(%s) wrote %q to stdout, want nothing", tt.name, stdout.String())
			}
			if n := bytes.Count(stderr.Bytes(), []byte("\n")); n != 1 {
				t.Errorf("Run(%s) wrote %d stderr lines, want 1", tt.name, n)
			}
		})
	}
}

func TestRejectionDiagnosticNamesFileAndLine(t *testing.T) {
	path := writeFile(t, "sku,price,in_stock\nA1,9.99,true\nA2,4.50,false\nA3,1.25,true,extra\nA4,3.00,false\n")
	var stdout, stderr bytes.Buffer
	code := Run([]string{path}, &stdout, &stderr)
	if code != ExitRejected {
		t.Fatalf("Run(ragged) = %d, want %d", code, ExitRejected)
	}
	if stdout.Len() != 0 {
		t.Errorf("Run(ragged) wrote %q to stdout, want nothing", stdout.String())
	}
	lines := strings.Split(strings.TrimSuffix(stderr.String(), "\n"), "\n")
	if len(lines) != 1 {
		t.Fatalf("Run(ragged) wrote %d stderr lines, want 1: %q", len(lines), stderr.String())
	}
	if !strings.Contains(lines[0], path) {
		t.Errorf("diagnostic %q does not name file %q", lines[0], path)
	}
	if !strings.Contains(lines[0], "line 4") {
		t.Errorf("diagnostic %q does not name ragged record line 4", lines[0])
	}
}

func TestParseCapturesOptions(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want options
	}{
		{name: "file only", args: []string{"a.csv"}, want: options{file: "a.csv", top: defaultTop}},
		{name: "column", args: []string{"--column", "name", "a.csv"}, want: options{file: "a.csv", column: "name", top: defaultTop, columnSet: true}},
		{name: "freq defaults top", args: []string{"--freq", "name", "a.csv"}, want: options{file: "a.csv", freq: "name", top: defaultTop, freqSet: true}},
		{name: "freq with top", args: []string{"--freq", "name", "--top", "7", "a.csv"}, want: options{file: "a.csv", freq: "name", top: 7, freqSet: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args, &bytes.Buffer{})
			if err != nil {
				t.Fatalf("parse(%q) error = %v", tt.args, err)
			}
			if got != tt.want {
				t.Errorf("parse(%q) = %+v, want %+v", tt.args, got, tt.want)
			}
		})
	}
}

// runCSV runs the command on a file and parses stdout back as CSV records.
func runCSV(t *testing.T, args ...string) (int, [][]string) {
	t.Helper()
	var stdout, stderr bytes.Buffer
	code := Run(args, &stdout, &stderr)
	if code != ExitOK {
		return code, nil
	}
	records, err := csv.NewReader(&stdout).ReadAll()
	if err != nil {
		t.Fatalf("stdout is not CSV: %v; stdout %q", err, stdout.String())
	}
	return code, records
}

var profileHeader = []string{"column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"}

func TestRunWritesHeaderAndOneRowPerColumnInFileOrder(t *testing.T) {
	path := writeFile(t, "name,id, score \nAda,1,2.5\nGrace,2,\n")
	code, records := runCSV(t, path)
	if code != ExitOK {
		t.Fatalf("Run = %d, want %d", code, ExitOK)
	}
	if got := records[0]; !reflect.DeepEqual(got, profileHeader) {
		t.Fatalf("header = %q, want %q", got, profileHeader)
	}
	var names []string
	for _, r := range records[1:] {
		names = append(names, r[0])
	}
	if want := []string{"name", "id", "score"}; !reflect.DeepEqual(names, want) {
		t.Errorf("column names = %q, want %q", names, want)
	}
	for _, r := range records {
		if len(r) != len(profileHeader) {
			t.Errorf("record %q has %d fields, want %d", r, len(r), len(profileHeader))
		}
	}
}

func TestRunProfilesEachTypeEndToEnd(t *testing.T) {
	path := writeFile(t, "n,i,d,b,t,e\n2.50,1,2025-01-15,true,abc,\n-1.25,2,2024-02-29,FALSE,xyz,NA\n10,3,2025-03-01,True,,NULL\n")
	code, records := runCSV(t, path)
	if code != ExitOK {
		t.Fatalf("Run = %d, want %d", code, ExitOK)
	}
	want := map[string][]string{
		"n": {"n", "number", "3", "0", "3", "-1.25", "10.00", "11.25", "3.75", "2.50"},
		"i": {"i", "integer", "3", "0", "3", "1", "3", "6", "2.00", "2.00"},
		"d": {"d", "date", "3", "0", "3", "2024-02-29", "2025-03-01", "", "", ""},
		"b": {"b", "boolean", "3", "0", "3", "false", "true", "", "", ""},
		"t": {"t", "text", "3", "1", "2", "abc", "xyz", "", "", ""},
		"e": {"e", "empty", "3", "3", "0", "", "", "", "", ""},
	}
	for _, r := range records[1:] {
		if !reflect.DeepEqual(r, want[r[0]]) {
			t.Errorf("row %q = %q, want %q", r[0], r, want[r[0]])
		}
	}
}

func TestRunHeaderOnlyFileHasZeroRowsAndEmptyType(t *testing.T) {
	path := writeFile(t, "sku,price,in_stock\n")
	_, records := runCSV(t, path)
	wantRows := [][]string{
		{"sku", "empty", "0", "0", "0", "", "", "", "", ""},
		{"price", "empty", "0", "0", "0", "", "", "", "", ""},
		{"in_stock", "empty", "0", "0", "0", "", "", "", "", ""},
	}
	if !reflect.DeepEqual(records[1:], wantRows) {
		t.Errorf("rows = %q, want %q", records[1:], wantRows)
	}
}

func TestRunColumnWritesHeaderAndOnlyThatRow(t *testing.T) {
	path := writeFile(t, "id,qty\n1,10\n2,20\n3,30\n")
	code, records := runCSV(t, "--column", "qty", path)
	if code != ExitOK {
		t.Fatalf("Run = %d, want %d", code, ExitOK)
	}
	want := [][]string{profileHeader, {"qty", "integer", "3", "0", "3", "10", "30", "60", "20.00", "20.00"}}
	if !reflect.DeepEqual(records, want) {
		t.Errorf("records = %q, want %q", records, want)
	}
}

func TestRunColumnMatchesFullProfileRow(t *testing.T) {
	path := writeFile(t, "id,name,when\n1,Ada,2025-01-01\n2,,2025-01-02\n3,Grace,2025-01-03\n")
	_, full := runCSV(t, path)
	for _, row := range full[1:] {
		_, got := runCSV(t, "--column", row[0], path)
		if !reflect.DeepEqual(got, [][]string{profileHeader, row}) {
			t.Errorf("--column %s = %q, want header and %q", row[0], got, row)
		}
	}
}

func TestRunUnknownColumnExitsThreeWithoutStdout(t *testing.T) {
	path := writeFile(t, "id,qty\n1,10\n")
	tests := []struct {
		name string
		args []string
	}{
		{name: "column", args: []string{"--column", "nosuch", path}},
		{name: "freq", args: []string{"--freq", "nosuch", path}},
		{name: "wrong case", args: []string{"--column", "QTY", path}},
		{name: "empty name", args: []string{"--column", "", path}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			code := Run(tt.args, &stdout, &stderr)
			if code != ExitRejected {
				t.Fatalf("Run(%q) = %d, want %d", tt.args, code, ExitRejected)
			}
			if stdout.Len() != 0 {
				t.Errorf("Run(%q) wrote %q to stdout, want nothing", tt.args, stdout.String())
			}
			if n := bytes.Count(stderr.Bytes(), []byte("\n")); n != 1 {
				t.Errorf("Run(%q) wrote %d stderr lines, want 1", tt.args, n)
			}
		})
	}
}

func TestRunRaggedFileReportsRejectionBeforeColumnLookup(t *testing.T) {
	path := writeFile(t, "a,b\n1,2\n3,4,5\n")
	var stdout, stderr bytes.Buffer
	code := Run([]string{"--column", "a", path}, &stdout, &stderr)
	if code != ExitRejected {
		t.Fatalf("Run = %d, want %d", code, ExitRejected)
	}
	if stdout.Len() != 0 {
		t.Errorf("Run wrote %q to stdout, want nothing", stdout.String())
	}
	if !strings.Contains(stderr.String(), "line 3") {
		t.Errorf("diagnostic %q does not name ragged line 3", stderr.String())
	}
}

func TestRunOutputUsesUnixLineEndings(t *testing.T) {
	path := writeFile(t, "a,b\r\n1,x\r\n")
	var stdout, stderr bytes.Buffer
	if code := Run([]string{path}, &stdout, &stderr); code != ExitOK {
		t.Fatalf("Run = %d, want %d", code, ExitOK)
	}
	if strings.Contains(stdout.String(), "\r") {
		t.Errorf("stdout contains carriage return: %q", stdout.String())
	}
	if !strings.HasSuffix(stdout.String(), "\n") {
		t.Errorf("stdout %q does not end with newline", stdout.String())
	}
}
