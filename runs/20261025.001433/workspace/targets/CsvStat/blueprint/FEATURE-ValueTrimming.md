# FEATURE: Value Trimming

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Trim leading and trailing spaces from every header name and every value before any other rule applies. |
| Depends On  | FEATURE-CsvReading.md, ARCHITECTURE.md |
| Provides    | trimmed header names and values |
| Consumes    | profile output header and rows |

## Purpose

Leading and trailing spaces are removed from every header name and every value before null, type,
distinct, and output rules apply. Trimming is done once, in `internal/csvio`, so every later rule
sees trimmed text.

## Behavior

1. Only the space character is trimmed (`strings.Trim(s, " ")`); tabs and other whitespace remain.
2. Header names are trimmed before the empty-name and repeated-name checks.
3. Values are trimmed before the null test, so a value of spaces only is empty and therefore null.
4. Column names written in the profile output are the trimmed names.

## Required Go Tests

The build adds these tests in `internal/csvio` (names are exact): `TestTrimHeaderNames`,
`TestTrimValues`.

## Programmatic Acceptance

=== AC trim-header-names ===
Intent: Padded header names are written trimmed in the profile output.

import csv
import io
import tempfile
from pathlib import Path
from drydock_go import cli_runner

names = ["alpha", "beta"]
with tempfile.TemporaryDirectory() as tmp:
    target = Path(tmp) / "padded.csv"
    target.write_text("  " + names[0] + " , " + names[1] + "  \n1,2\n")
    code, out, err = cli_runner(str(target))
    print(out)
    print(err)
    assert code == 0
    rows = list(csv.reader(io.StringIO(out)))
    assert [row[0] for row in rows[1:]] == names
=== END AC trim-header-names ===

=== AC trim-only-spaces ===
Intent: Only spaces are trimmed; a trailing tab stays part of the header name.

import csv
import io
import tempfile
from pathlib import Path
from drydock_go import cli_runner

name = "alpha\t"
other = "beta"
with tempfile.TemporaryDirectory() as tmp:
    target = Path(tmp) / "tab.csv"
    target.write_text(" " + name + " ," + other + "\n1,2\n")
    code, out, err = cli_runner(str(target))
    assert code == 0
    rows = list(csv.reader(io.StringIO(out)))
    assert rows[1][0] == name
    assert rows[2][0] == other
=== END AC trim-only-spaces ===

=== AC trim-header-test ===
Intent: The package test for header trimming passes.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvio", "TestTrimHeaderNames")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC trim-header-test ===

=== AC trim-values-test ===
Intent: The package test for value trimming passes.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvio", "TestTrimValues")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC trim-values-test ===

## User Acceptance

- None.

## Guardrails

- Tabs, non-breaking spaces, and other whitespace are not trimmed.
- Trimming happens before every other rule; no later package re-trims or sees untrimmed text.
- Standard library only.
