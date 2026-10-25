# FEATURE: Test Gate

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Provides the POSIX `bin/test.sh` that runs `go vet ./...` and `go test ./...` and exits zero only when both pass. |
| Depends On  | FEATURE-ColumnProfile.md, FEATURE-FreqMode.md, FEATURE-UnreadableFile.md |
| Provides    | bin/test.sh |
| Consumes    | csvstat command, go vet, go test |

## Purpose

`bin/test.sh` is the project's single test entry point. It is a POSIX `sh` script that changes to the application root, runs `go vet ./...`, then `go test ./...`, and exits non-zero if either fails. The project suite, written beside the code, covers each type, each statistic, null handling, rounding, each output form, and each exit code, using files the tests write to temporary directories.

## Behavior

- The script uses only POSIX `sh` constructs and resolves the application root from its own location, so it works from any current directory.
- It runs `go vet ./...` first and stops with a non-zero status if vet fails.
- It then runs `go test ./...` and exits with that command's status.
- It never modifies files under `sources/`.
- The suite covers:
  - every column type (integer, number, date, boolean, text, empty);
  - rows, nulls, distinct, min, max, sum, mean, and median;
  - null spellings and trimming;
  - half-away-from-zero rounding and no negative zero;
  - the profile form, the `--column` form, and the `--freq` form;
  - exit codes 0, 1, 2, and 3.

## Programmatic Acceptance

=== AC testgate-script-exits-zero ===
Intent: Running bin/test.sh from the application root exits zero.
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(["sh", "bin/test.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
=== END AC testgate-script-exits-zero ===

=== AC testgate-runs-from-other-directory ===
Intent: bin/test.sh locates the application root itself and exits zero when started from another directory.
Requires: executable=go; scope=test

import os
import subprocess
import tempfile

script = os.path.abspath("bin/test.sh")
with tempfile.TemporaryDirectory() as tmp:
    result = subprocess.run(["sh", script], capture_output=True, text=True, cwd=tmp)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
=== END AC testgate-runs-from-other-directory ===

=== AC testgate-go-test-and-vet-pass ===
Intent: go vet and go test pass directly at the application root, the two commands the script wraps.
Requires: executable=go; scope=test

import subprocess

vet = subprocess.run(["go", "vet", "./..."], capture_output=True, text=True)
print(vet.stdout)
print(vet.stderr)
assert vet.returncode == 0, vet.returncode

tests = subprocess.run(["go", "test", "./..."], capture_output=True, text=True)
print(tests.stdout)
print(tests.stderr)
assert tests.returncode == 0, tests.returncode
=== END AC testgate-go-test-and-vet-pass ===

=== AC testgate-suite-has-tests ===
Intent: The project ships Go test files in more than one internal package, so the gate is not vacuous.
Requires: executable=go; scope=test

import pathlib

root = pathlib.Path("internal")
test_files = sorted(root.rglob("*_test.go"))
packages = {p.parent for p in test_files}
print([str(p) for p in test_files])
assert len(test_files) > 0
assert len(packages) > 1, packages
=== END AC testgate-suite-has-tests ===

## User Acceptance

- None.

## Guardrails

- `bin/test.sh` never edits, regenerates, or substitutes any file under `sources/`.
- The test suite never reads `sources/run.sh` and never embeds a grader answer.
- Tests write their own input files to temporary directories.
