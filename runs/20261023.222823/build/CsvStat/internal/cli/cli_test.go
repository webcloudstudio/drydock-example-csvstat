package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRunRejectsMissingInput(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := Run(nil, &stdout, &stderr)

	if code != ExitUsage {
		t.Fatalf("Run(nil) = %d, want %d", code, ExitUsage)
	}
	if stdout.Len() != 0 {
		t.Fatalf("stdout length = %d, want 0", stdout.Len())
	}
}

func TestRunPrintsHelp(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := Run([]string{"--help"}, &stdout, &stderr)

	if code != ExitOK {
		t.Fatalf("Run(-help) = %d, want %d", code, ExitOK)
	}
	if stdout.Len() == 0 {
		t.Fatal("help output is empty")
	}
}

func TestRunAcceptsColumnProfileInvocation(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "employees.csv")
	if err := os.WriteFile(path, []byte("age\n36\n"), 0600); err != nil {
		t.Fatal(err)
	}
	var stdout, stderr bytes.Buffer

	code := Run([]string{"--column", "age", path}, &stdout, &stderr)

	if code != ExitOK {
		t.Fatalf("profile invocation = %d, want %d", code, ExitOK)
	}
}

func TestRunAcceptsFrequencyInvocationAndTop(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "employees.csv")
	if err := os.WriteFile(path, []byte("department\nSales\nSupport\nSales\n"), 0600); err != nil {
		t.Fatal(err)
	}
	var stdout, stderr bytes.Buffer

	code := Run([]string{"--freq", "department", "--top", "3", path}, &stdout, &stderr)

	if code != ExitOK {
		t.Fatalf("frequency invocation = %d, want %d", code, ExitOK)
	}
}

func TestRunRejectsConflictingModes(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := Run([]string{"--column", "age", "--freq", "department", "employees.csv"}, &stdout, &stderr)

	if code != ExitUsage {
		t.Fatalf("conflicting modes = %d, want %d", code, ExitUsage)
	}
}

func TestRunRejectsNegativeTop(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := Run([]string{"--freq", "department", "--top", "-1", "employees.csv"}, &stdout, &stderr)

	if code != ExitUsage {
		t.Fatalf("negative top = %d, want %d", code, ExitUsage)
	}
}
