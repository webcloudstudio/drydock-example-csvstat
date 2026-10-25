# FEATURE: Release Verification

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Terminal assembly gate that runs the supplied grader and the project test entry point over the finished csvstat application. |
| Depends On  | FEATURE-Test-Entry.md, FEATURE-Read-Failure.md, FEATURE-Frequency-Top.md, FEATURE-Frequency-Empty.md, FEATURE-CLI-CSV-Output.md, FEATURE-Profile-Column-Option.md |
| Provides    | release verification |
| Consumes    | csvstat executable, bin/test.sh |

## Purpose

This is the terminal story. Every implementation story has closed before it runs. It adds no
application behavior. It confirms that the assembled `./cmd/csvstat` builds and that the supplied
grader `sources/run.sh` passes all 22 checks, which are the sole acceptance verdict. It also
confirms that `bin/test.sh` exits zero. The grader is read-only and is executed, never edited.
Application code never reads `sources/run.sh` or the supplied data files.

## Assembly

- `sh sources/run.sh` is run from the application root. It builds `./cmd/csvstat` with `go build`
  and compares exit code and standard output for 22 commands (A1-A4, B1-B3, C1-C5, E1-E4, U1-U6).
- `sh bin/test.sh` runs `go vet ./...` and `go test ./...`.
- The built executable is exercised directly through `drydock_go.cli_runner` for one usage error
  and one profile of a staged data file.

## Programmatic Acceptance

=== AC verification-grader ===
Intent: The supplied grader builds csvstat and passes every one of its 22 checks, so it exits zero.
Suite: full
Requires: executable=go; scope=test
Requires: executable=sh; scope=test

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/run.sh"],
    capture_output=True, text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0, result.returncode
=== END AC verification-grader ===

=== AC verification-test-entry ===
Intent: The project test entry point runs go vet and go test and exits zero.
Requires: executable=go; scope=test
Requires: executable=sh; scope=test

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "bin/test.sh"],
    capture_output=True, text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0, result.returncode
=== END AC verification-test-entry ===

=== AC verification-usage-error ===
Intent: The built executable exits 2 with nothing on standard output when no file is given.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner()
print(out)
print(err)
assert code == 2, code
assert out == "", out
=== END AC verification-usage-error ===

=== AC verification-profile-header-only ===
Intent: The built executable profiles a header-only file with one parsed row per header column, each reporting zero rows.
Requires: executable=go; scope=test

import csv
import io
import os

from drydock_go import cli_runner

data = os.path.abspath("sources/header_only.csv")
with open(data, newline="", encoding="utf-8") as handle:
    header = next(csv.reader(handle))

code, out, err = cli_runner(data)
print(out)
print(err)
assert code == 0, code

records = list(csv.DictReader(io.StringIO(out)))
assert [r["column"] for r in records] == header
assert all(r["rows"] == str(0) for r in records)
assert all(r["nulls"] == str(0) for r in records)
=== END AC verification-profile-header-only ===

=== AC verification-rejected-input ===
Intent: The built executable exits 3 with nothing on standard output for the supplied ragged file.
Requires: executable=go; scope=test

import os

from drydock_go import cli_runner

data = os.path.abspath("sources/ragged.csv")
code, out, err = cli_runner(data)
print(out)
print(err)
assert code == 3, code
assert out == "", out
=== END AC verification-rejected-input ===

## User Acceptance

- None.

## Guardrails

- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- Application code never reads `sources/run.sh` or any supplied data file, and never hardcodes a statistic, a type, or a value list.
- Only this terminal story runs the whole grader; no other story invokes `sh sources/run.sh`.
- The story adds no application behavior; a failure here is repaired in the story that owns the behavior.
