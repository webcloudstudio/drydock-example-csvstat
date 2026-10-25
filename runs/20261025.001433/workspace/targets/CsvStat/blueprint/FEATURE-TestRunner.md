# FEATURE: Test Runner

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Provides a POSIX sh bin/test.sh that runs go vet and go test and exits zero only when both pass. |
| Depends On  | FEATURE-CsvReading.md, ARCHITECTURE.md |
| Provides    | bin/test.sh |
| Consumes    | go vet ./..., go test ./... |

## Purpose

`bin/test.sh` is the single entry point for the project's automated checks.

## Behavior

- POSIX sh only: no bash-isms, and it passes `sh -n bin/test.sh`.
- It changes to the application root (the parent of `bin/`) so it works from any directory.
- It runs `go vet ./...` and then `go test ./...`, and exits zero only when both exit zero. A
  failure of either exits non-zero; a vet failure does not let the test run mask it.
- It uses whatever `go` is on `PATH` and never a vendored toolchain.
- The Go test suite it runs covers each type, each statistic, null handling, rounding, each output
  form, and each exit code, using files the tests write themselves (not the files under `sources/`).

## Programmatic Acceptance

=== AC testrunner-syntax ===
Intent: bin/test.sh exists and passes a POSIX syntax check.
Suite: scoped
Requires: executable=sh; scope=test

import os
import subprocess

assert os.path.isfile("bin/test.sh")
result = subprocess.run(["sh", "-n", "bin/test.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC testrunner-syntax ===

=== AC testrunner-runs-vet-and-test ===
Intent: With a passing toolchain stub, bin/test.sh exits zero and invokes both go vet and go test.
Suite: scoped
Requires: executable=sh; scope=test

import os
import subprocess
import tempfile

with tempfile.TemporaryDirectory() as tmp:
    log = os.path.join(tmp, "calls.log")
    stub = os.path.join(tmp, "go")
    with open(stub, "w") as fh:
        fh.write('#!/bin/sh\necho "$1" >> "$GO_LOG"\n[ "$1" = "$FAIL_ON" ] && exit 1\nexit 0\n')
    os.chmod(stub, 0o755)
    env = {**os.environ, "PATH": tmp + os.pathsep + os.environ["PATH"], "GO_LOG": log, "FAIL_ON": "none"}
    result = subprocess.run(["sh", "bin/test.sh"], capture_output=True, text=True, env=env)
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
    with open(log) as fh:
        calls = fh.read().split()
    assert "vet" in calls, calls
    assert "test" in calls, calls
=== END AC testrunner-runs-vet-and-test ===

=== AC testrunner-vet-failure ===
Intent: When go vet fails, bin/test.sh exits non-zero.
Suite: scoped
Requires: executable=sh; scope=test

import os
import subprocess
import tempfile

with tempfile.TemporaryDirectory() as tmp:
    log = os.path.join(tmp, "calls.log")
    stub = os.path.join(tmp, "go")
    with open(stub, "w") as fh:
        fh.write('#!/bin/sh\necho "$1" >> "$GO_LOG"\n[ "$1" = "$FAIL_ON" ] && exit 1\nexit 0\n')
    os.chmod(stub, 0o755)
    env = {**os.environ, "PATH": tmp + os.pathsep + os.environ["PATH"], "GO_LOG": log, "FAIL_ON": "vet"}
    result = subprocess.run(["sh", "bin/test.sh"], capture_output=True, text=True, env=env)
    print(result.stdout)
    print(result.stderr)
    assert result.returncode != 0
=== END AC testrunner-vet-failure ===

=== AC testrunner-test-failure ===
Intent: When go test fails after a passing go vet, bin/test.sh exits non-zero.
Suite: scoped
Requires: executable=sh; scope=test

import os
import subprocess
import tempfile

with tempfile.TemporaryDirectory() as tmp:
    log = os.path.join(tmp, "calls.log")
    stub = os.path.join(tmp, "go")
    with open(stub, "w") as fh:
        fh.write('#!/bin/sh\necho "$1" >> "$GO_LOG"\n[ "$1" = "$FAIL_ON" ] && exit 1\nexit 0\n')
    os.chmod(stub, 0o755)
    env = {**os.environ, "PATH": tmp + os.pathsep + os.environ["PATH"], "GO_LOG": log, "FAIL_ON": "test"}
    result = subprocess.run(["sh", "bin/test.sh"], capture_output=True, text=True, env=env)
    print(result.stdout)
    print(result.stderr)
    assert result.returncode != 0
=== END AC testrunner-test-failure ===

=== AC testrunner-works-from-other-directory ===
Intent: bin/test.sh locates the application root itself, so a stubbed run from another directory still invokes go vet and go test.
Suite: scoped
Requires: executable=sh; scope=test

import os
import subprocess
import tempfile

script = os.path.abspath("bin/test.sh")
with tempfile.TemporaryDirectory() as tmp:
    log = os.path.join(tmp, "calls.log")
    stub = os.path.join(tmp, "go")
    with open(stub, "w") as fh:
        fh.write('#!/bin/sh\necho "$1" >> "$GO_LOG"\n[ "$1" = "$FAIL_ON" ] && exit 1\nexit 0\n')
    os.chmod(stub, 0o755)
    env = {**os.environ, "PATH": tmp + os.pathsep + os.environ["PATH"], "GO_LOG": log, "FAIL_ON": "none"}
    result = subprocess.run(["sh", script], capture_output=True, text=True, env=env, cwd=tmp)
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
    with open(log) as fh:
        calls = fh.read().split()
    assert "vet" in calls and "test" in calls, calls
=== END AC testrunner-works-from-other-directory ===

## User Acceptance

- None.

## Guardrails

- `bin/test.sh` is POSIX sh and never relies on bash.
- It does not run `sources/run.sh`; the supplied grader is run only by the terminal verification story.
- Tests written for the project never modify, rewrite, or substitute any file under `sources/`.
