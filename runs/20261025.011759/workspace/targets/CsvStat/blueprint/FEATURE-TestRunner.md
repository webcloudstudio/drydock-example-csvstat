# FEATURE: Test Runner

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Provide the POSIX bin/test.sh that runs go vet and go test and exits zero only when both pass. |
| Depends On  | FEATURE-FrequencyTop.md, FEATURE-ColumnSelect.md, FEATURE-ReadFailure.md, FEATURE-Build.md |
| Provides    | bin/test.sh |
| Consumes    | go vet, go test |

## Purpose

`bin/test.sh` is the single command that gates every behavioral story. The project's own Go test
suite is written beside the code and covers each type, each statistic, null handling, rounding,
each output form, and each exit code, using files the tests write themselves.

## Behavior

- POSIX sh only; no bashisms. It changes to the application root (the directory above `bin/`) so it
  works from any current directory.
- It runs `go vet ./...` and then `go test ./...`, and exits zero only when both succeed. A failure
  of `go vet` stops the script with a non-zero status.
- Tests call `cli.Run(args, &stdout, &stderr)` with `bytes.Buffer` writers and write their fixture
  files under `t.TempDir()`. They never read `sources/run.sh`.
- The suite covers: each of the five types, min/max per type, sum/mean/median, half-away-from-zero
  rounding including negatives, null classification, quoting of fields with commas, frequency
  ordering and `--top`, and exit codes 1, 2, and 3.

## Programmatic Acceptance

=== AC test-script-passes ===
Intent: bin/test.sh runs go vet and go test over the application and exits zero.
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(["sh", "bin/test.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
=== END AC test-script-passes ===

=== AC test-script-is-posix-sh ===
Intent: bin/test.sh parses under a POSIX shell without executing it.
Requires: executable=sh; scope=test

import subprocess

result = subprocess.run(["sh", "-n", "bin/test.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
=== END AC test-script-is-posix-sh ===

=== AC test-script-works-from-other-directory ===
Intent: bin/test.sh succeeds when invoked from a directory other than the application root.
Requires: executable=go; scope=test

import os
import subprocess
import tempfile

script = os.path.abspath("bin/test.sh")
with tempfile.TemporaryDirectory() as elsewhere:
    result = subprocess.run(["sh", script], cwd=elsewhere, capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
=== END AC test-script-works-from-other-directory ===

=== AC test-suite-exists ===
Intent: The module contains automated Go tests that go test discovers and runs successfully.
Requires: executable=go; scope=test

import subprocess

listed = subprocess.run(["go", "test", "-list", ".", "./..."], capture_output=True, text=True)
print(listed.stdout)
print(listed.stderr)
assert listed.returncode == 0, listed.returncode
names = [line for line in listed.stdout.splitlines() if line.startswith("Test")]
assert len(names) > 0, names
=== END AC test-suite-exists ===

## User Acceptance

- None.

## Guardrails

- The script uses only POSIX sh constructs and the Go toolchain; no vendored toolchain.
- Tests never read `sources/run.sh` or the supplied data files, and embed no answers from them.
