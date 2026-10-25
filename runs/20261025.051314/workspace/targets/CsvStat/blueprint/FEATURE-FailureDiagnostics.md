# FEATURE: Failure Diagnostics

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Ensures every failing csvstat command writes nothing to standard output and exactly one diagnostic line to standard error. |
| Depends On  | FEATURE-InputRejection.md, FEATURE-UsageErrors.md, FEATURE-ProfileOutput.md |
| Provides    | stderr diagnostics, exit code 3 |
| Consumes    | csvdata.ErrRejected, cli.Run |

## Purpose

A command that fails reports the failure on standard error as one line and leaves standard output empty. Rejected input exits 3. The diagnostic names the file and, for a rejected record, its line number. Standard error is never compared by the grader, so its wording is free; its shape (one line, nothing on standard output) is the contract.

## Behavior

- Rejected input (empty file, file containing only blank lines, empty header name, repeated header name, ragged record) exits 3.
- Standard output is written only after the whole result is built, so no partial output precedes a failure.
- Exactly one non-blank line is written to standard error per failure.
- A ragged record's diagnostic includes the line number of the offending record.
- Exit code 1 (unreadable file) and exit code 2 (usage) share the same stdout-empty, one-line-stderr rule and are owned by their own stories.

## Programmatic Acceptance

=== AC diag-ragged-fixture ===
Intent: The supplied ragged fixture exits 3, writes nothing to standard output, and writes one diagnostic line.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("sources/ragged.csv")
print(out)
print(err)
assert code == 3, code
assert out == "", out
lines = [ln for ln in err.splitlines() if ln.strip()]
assert len(lines) == 1, lines
=== END AC diag-ragged-fixture ===

=== AC diag-ragged-line-number ===
Intent: The diagnostic for a ragged record carries the number of the line holding that record.
Requires: executable=go; scope=test

import csv
import os
import re
import tempfile
from drydock_go import cli_runner

rows = [["a", "b"], ["1", "2"], ["3", "4"], ["5", "6", "7"], ["8", "9"]]
bad_line = 4
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "ragged_case.csv")
    with open(path, "w", newline="", encoding="utf-8") as fh:
        csv.writer(fh, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3, code
assert out == "", out
assert str(bad_line) in re.findall(r"\d+", err), err
=== END AC diag-ragged-line-number ===

=== AC diag-empty-file ===
Intent: An empty file is rejected with exit 3 and nothing on standard output.
Requires: executable=go; scope=test

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "empty.csv")
    with open(path, "w", encoding="utf-8") as fh:
        fh.write("")
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3, code
assert out == "", out
lines = [ln for ln in err.splitlines() if ln.strip()]
assert len(lines) == 1, lines
=== END AC diag-empty-file ===

=== AC diag-blank-lines-only ===
Intent: A file containing only blank lines is rejected as empty with exit 3.
Requires: executable=go; scope=test

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "blank.csv")
    with open(path, "w", encoding="utf-8", newline="") as fh:
        fh.write("\n\n\n")
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3, code
assert out == "", out
=== END AC diag-blank-lines-only ===

=== AC diag-empty-header-name ===
Intent: An empty header name is rejected with exit 3 and nothing on standard output.
Requires: executable=go; scope=test

import csv
import os
import tempfile
from drydock_go import cli_runner

rows = [["a", "", "c"], ["1", "2", "3"]]
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "emptyname.csv")
    with open(path, "w", newline="", encoding="utf-8") as fh:
        csv.writer(fh, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3, code
assert out == "", out
=== END AC diag-empty-header-name ===

=== AC diag-repeated-header-name ===
Intent: A repeated header name, including one repeated only after trimming, is rejected with exit 3.
Requires: executable=go; scope=test

import csv
import os
import tempfile
from drydock_go import cli_runner

name = "dup"
rows = [[name, "x", "  " + name + "  "], ["1", "2", "3"]]
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "repeated.csv")
    with open(path, "w", newline="", encoding="utf-8") as fh:
        csv.writer(fh, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3, code
assert out == "", out
=== END AC diag-repeated-header-name ===

=== AC diag-failure-with-column-option ===
Intent: A rejected file under --column still exits 3 with nothing on standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("--column", "sku", "sources/ragged.csv")
print(out)
print(err)
assert code == 3, code
assert out == "", out
=== END AC diag-failure-with-column-option ===

## User Acceptance

- None.

## Guardrails

- A failing command never writes a partial profile or frequency listing to standard output.
- Diagnostics go to standard error only and are never compared by the grader.
- Rejected files are never modified.
