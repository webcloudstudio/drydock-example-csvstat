package main

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestModuleDeclaresNoRequirementsOrChecksums(t *testing.T) {
	mod, err := os.ReadFile(filepath.Join("..", "..", "go.mod"))
	if err != nil {
		t.Fatalf("read go.mod: %v", err)
	}
	for _, directive := range []string{`(?m)^\s*require\b`, `(?m)^\s*replace\b`} {
		if regexp.MustCompile(directive).Match(mod) {
			t.Errorf("go.mod contains %s", directive)
		}
	}
	if _, err := os.Stat(filepath.Join("..", "..", "go.sum")); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("go.sum exists or cannot be checked: %v", err)
	}
}

func TestBuiltBinaryExitsWithUsageWhenNoArguments(t *testing.T) {
	bin := filepath.Join(t.TempDir(), "csvstat")
	if out, err := exec.Command("go", "build", "-o", bin, ".").CombinedOutput(); err != nil {
		t.Fatalf("go build failed: %v\n%s", err, out)
	}
	cmd := exec.Command(bin)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) || exitErr.ExitCode() != 2 {
		t.Fatalf("exit error = %v, want exit status 2", err)
	}
	if stdout.Len() != 0 || !strings.Contains(stderr.String(), "csvstat [--column <name>] <file>") {
		t.Errorf("stdout %q, stderr %q", stdout.String(), stderr.String())
	}
}
