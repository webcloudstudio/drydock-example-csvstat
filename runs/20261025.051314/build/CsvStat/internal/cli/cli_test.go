package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

// runWith writes content to a temporary file and runs the command on it.
func runWith(t *testing.T, content string) (code int, stdout *bytes.Buffer, stderr *bytes.Buffer) {
	t.Helper()
	path := filepath.Join(t.TempDir(), "input.csv")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write input: %v", err)
	}
	return runArgs(t, path)
}

func runArgs(t *testing.T, args ...string) (code int, stdout *bytes.Buffer, stderr *bytes.Buffer) {
	t.Helper()
	stdout, stderr = &bytes.Buffer{}, &bytes.Buffer{}
	code = Run(args, stdout, stderr)
	return code, stdout, stderr
}

func TestRunValidFileExitsZero(t *testing.T) {
	code, _, _ := runWith(t, "a,b\n1,2\n")
	if code != exitOK {
		t.Fatalf("exit code = %d, want %d", code, exitOK)
	}
}

func TestRunHeaderOnlyFileExitsZero(t *testing.T) {
	code, _, _ := runWith(t, "a,b\n")
	if code != exitOK {
		t.Fatalf("exit code = %d, want %d", code, exitOK)
	}
}

func TestRunRejectedInputExitsThreeWithEmptyStdout(t *testing.T) {
	cases := map[string]string{
		"empty file":        "",
		"blank lines only":  "\n\n\n",
		"empty header name": "id,,name\n1,x,y\n",
		"repeated header":   "id,name,id\n1,x,2\n",
		"ragged record":     "a,b\n1,2\n3,4,5\n6,7\n",
		"bare quote":        "a,b\n1\"x,2\n",
	}
	for name, content := range cases {
		t.Run(name, func(t *testing.T) {
			code, stdout, _ := runWith(t, content)
			if code != exitRejected {
				t.Fatalf("exit code = %d, want %d", code, exitRejected)
			}
			if stdout.Len() != 0 {
				t.Fatalf("stdout has %d bytes, want none", stdout.Len())
			}
		})
	}
}

func TestRunUsageErrorExitsTwoWithEmptyStdout(t *testing.T) {
	path := filepath.Join(t.TempDir(), "input.csv")
	if err := os.WriteFile(path, []byte("a,b\n1,2\n"), 0o600); err != nil {
		t.Fatalf("write input: %v", err)
	}
	cases := map[string][]string{
		"no arguments":      {},
		"option after file": {path, "--column", "a"},
		"unknown option":    {"--frobnicate", path},
	}
	for name, args := range cases {
		t.Run(name, func(t *testing.T) {
			code, stdout, _ := runArgs(t, args...)
			if code != exitUsage {
				t.Fatalf("exit code = %d, want %d", code, exitUsage)
			}
			if stdout.Len() != 0 {
				t.Fatalf("stdout has %d bytes, want none", stdout.Len())
			}
		})
	}
}

func TestRunMissingFileExitsOne(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "absent.csv")
	code, stdout, _ := runArgs(t, missing)
	if code != exitOperational {
		t.Fatalf("exit code = %d, want %d", code, exitOperational)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout has %d bytes, want none", stdout.Len())
	}
}

func TestRunValidFormsAreNotUsageErrors(t *testing.T) {
	path := filepath.Join(t.TempDir(), "input.csv")
	if err := os.WriteFile(path, []byte("a,b\n1,2\n"), 0o600); err != nil {
		t.Fatalf("write input: %v", err)
	}
	forms := [][]string{
		{path},
		{"--column", "a", path},
		{"--freq", "a", path},
		{"--freq", "a", "--top", "3", path},
	}
	for _, form := range forms {
		if code, _, _ := runArgs(t, form...); code == exitUsage {
			t.Errorf("Run(%q) exit code = %d, want non-usage", form, code)
		}
	}
}
