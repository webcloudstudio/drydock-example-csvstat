# FEATURE: Test Entry Point

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | A POSIX-compatible bin/test.sh runs go vet and go test and exits zero only when both pass. |
| Depends On  | ARCHITECTURE.md, FEATURE-Executable-Build.md |
| Provides    | bin/test.sh |
| Consumes    | — |

## Purpose

`bin/test.sh` is the single project test command. It runs from the application root regardless of the caller's working directory.

## Behavior

- Written in POSIX sh with no bash-only syntax.
- Changes to the application root, runs `go vet ./...`, then `go test ./...`.
- Exits non-zero if either command fails; exits zero only when both pass.
- The Go test suite is non-empty from this story on. Later stories add tests beside their code and cover each type, statistic, null rule, rounding case, output form, and exit code using files the tests write.

## Reads

- The Go module in the application root.

## Writes

- Nothing persistent.

## Programmatic Acceptance

=== AC test-entry-syntax ===
Intent: bin/test.sh exists and is syntactically valid POSIX sh.

import pathlib
import subprocess

script = pathlib.Path("bin/test.sh")
assert script.is_file()
result = subprocess.run(["sh", "-n", "bin/test.sh"], capture_output=True, text=True)
print(result.stderr)
assert result.returncode == 0
=== END AC test-entry-syntax ===

=== AC test-entry-passes ===
Intent: sh bin/test.sh exits 0 when go vet and go test both pass.
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(["sh", "bin/test.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC test-entry-passes ===

=== AC test-entry-runs-from-elsewhere ===
Intent: bin/test.sh exits 0 when invoked from a different working directory.
Requires: executable=go; scope=test

import os
import subprocess
import tempfile

script = os.path.abspath("bin/test.sh")
with tempfile.TemporaryDirectory() as elsewhere:
    result = subprocess.run(["sh", script], cwd=elsewhere, capture_output=True, text=True)
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
=== END AC test-entry-runs-from-elsewhere ===

=== AC test-entry-suite-nonempty ===
Intent: The module contains at least one Go test file.

import pathlib

tests = list(pathlib.Path(".").rglob("*_test.go"))
assert tests
=== END AC test-entry-suite-nonempty ===

## User Acceptance

- None.

## Guardrails

- POSIX sh only; no bash-specific features.
- The script never modifies files under `sources/`.
- The script must not mask a failing `go vet` or `go test` exit status.
