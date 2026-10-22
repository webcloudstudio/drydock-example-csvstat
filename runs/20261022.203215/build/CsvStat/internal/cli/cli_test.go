package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRunRequiresExactlyOneInputPath(t *testing.T) {
	path := writeCSV(t, "a\n1\n")
	tests := []struct {
		name string
		args []string
		want int
	}{
		{name: "missing", args: nil, want: 2},
		{name: "one path", args: []string{path}, want: 0},
		{name: "two paths", args: []string{"one.csv", "two.csv"}, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			if got := Run(tt.args, &stdout, &stderr); got != tt.want {
				t.Fatalf("Run(%v) = %d, want %d", tt.args, got, tt.want)
			}
		})
	}
}

func TestRunHelpIsStderrOnly(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if got := Run([]string{"--help"}, &stdout, &stderr); got != 0 {
		t.Fatalf("Run(--help) = %d, want 0", got)
	}
	if stdout.Len() != 0 || stderr.Len() == 0 {
		t.Fatalf("help streams = stdout %d bytes, stderr %d bytes", stdout.Len(), stderr.Len())
	}
}

func TestRunRejectsMalformedInvocationsWithoutOutput(t *testing.T) {
	path := writeCSV(t, "a\n1\n")
	for _, args := range [][]string{
		{}, {path, path}, {"--top", "0", path}, {"--top", "2", path},
		{"--column", "a", "--freq", "a", path}, {"--unknown", path},
		{"--help", "--unknown"},
		{"--freq", "a", "--top", "1", "--top", "2", path},
	} {
		var stdout, stderr bytes.Buffer
		if got := Run(args, &stdout, &stderr); got != 2 {
			t.Errorf("Run(%v) = %d, want 2", args, got)
		}
		if stdout.Len() != 0 {
			t.Errorf("Run(%v) wrote %d bytes to stdout", args, stdout.Len())
		}
	}
}

func TestRunDispatchesProfileAndFrequency(t *testing.T) {
	path := writeCSV(t, "a,b\n1,x\n2,y\n")
	for _, args := range [][]string{{path}, {"--freq", "b", "--top", "1", path}} {
		var stdout, stderr bytes.Buffer
		if got := Run(args, &stdout, &stderr); got != 0 {
			t.Fatalf("Run(%v) = %d, want 0; stderr=%q", args, got, stderr.String())
		}
		if stdout.Len() == 0 {
			t.Errorf("Run(%v) wrote no report", args)
		}
	}
}

func TestRunClassifiesInputFailures(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if got := Run([]string{"does-not-exist.csv"}, &stdout, &stderr); got != 1 {
		t.Fatalf("missing file exit = %d, want 1", got)
	}

	path := writeCSV(t, "a,b\n1\n")
	stdout.Reset()
	stderr.Reset()
	if got := Run([]string{path}, &stdout, &stderr); got != 3 {
		t.Fatalf("ragged CSV exit = %d, want 3", got)
	}

	stdout.Reset()
	stderr.Reset()
	if got := Run([]string{"--column", "missing", path}, &stdout, &stderr); got != 3 {
		t.Fatalf("missing column exit = %d, want 3", got)
	}
}

func writeCSV(t *testing.T, contents string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "data.csv")
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}
