package cli

import (
	"bytes"
	"os"
	"path/filepath"
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

func TestRunValidFileExitsZeroWithoutStdout(t *testing.T) {
	path := writeFile(t, " id , name \n 1 , Ada \n 2 , Grace \n")
	var stdout, stderr bytes.Buffer
	code := Run([]string{path}, &stdout, &stderr)
	if code != ExitOK {
		t.Fatalf("Run(valid) = %d, want %d; stderr %q", code, ExitOK, stderr.String())
	}
	if stdout.Len() != 0 {
		t.Errorf("Run(valid) wrote %q to stdout, want nothing until profiling is built", stdout.String())
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
		{name: "column", args: []string{"--column", "name", "a.csv"}, want: options{file: "a.csv", column: "name", top: defaultTop}},
		{name: "freq defaults top", args: []string{"--freq", "name", "a.csv"}, want: options{file: "a.csv", freq: "name", top: defaultTop}},
		{name: "freq with top", args: []string{"--freq", "name", "--top", "7", "a.csv"}, want: options{file: "a.csv", freq: "name", top: 7}},
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
