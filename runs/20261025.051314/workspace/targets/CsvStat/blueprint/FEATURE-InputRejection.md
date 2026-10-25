# FEATURE: Input Rejection

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Reject empty files, blank-only files, empty or repeated header names, and ragged records with exit 3. |
| Depends On  | FEATURE-Trimming.md, FEATURE-CsvReading.md, FEATURE-Arguments.md |
| Provides    | csvdata.ErrRejected |
| Consumes    | csvdata.Read, cli.ParseArgs |

## Purpose

Malformed input must never produce partial output. Rejection is detected in `internal/csvdata`
and mapped to exit code 3 by `internal/cli`.

## Behavior

- `csvdata` exposes a rejection error type (sentinel or typed error matched with `errors.Is` /
  `errors.As`) named `ErrRejected`. For a ragged record it carries the 1-based line number.
- The file is rejected when:
  - it is empty (zero bytes);
  - it contains only blank lines;
  - any header name is empty after trimming;
  - any header name repeats (compared exactly after trimming);
  - any record has a different number of fields from the header.
- A header-only file is valid and yields zero data records.
- A trailing line ending after the last record creates no extra record.
- `cli.Run` reads the file named by the parsed arguments. On a rejection error it writes nothing
  to standard output, writes one diagnostic line to standard error naming the file (and the line
  number for a ragged record), and returns exit code 3.
- A CSV syntax error from `encoding/csv` (for example a bare quote) is also a rejection, exit 3.

## Tests the build must write

In `internal/csvdata`:

- `TestRejectEmptyFile`
- `TestRejectBlankLinesOnly`
- `TestRejectEmptyHeaderName`
- `TestRejectRepeatedHeaderName`
- `TestRejectRaggedRecord`
- `TestHeaderOnlyFileAccepted`
- `TestTrailingNewlineAddsNoRecord`

## Programmatic Acceptance

=== AC reject-unit-empty-and-blank ===
Intent: Empty files and files of only blank lines are rejected by the reader.
Requires: executable=go; scope=test

from drydock_go import run_test

for name in ("TestRejectEmptyFile", "TestRejectBlankLinesOnly"):
    result = run_test("./internal/csvdata", name)
    print(result.stdout)
    print(result.stderr)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC reject-unit-empty-and-blank ===

=== AC reject-unit-headers ===
Intent: Empty and repeated header names are rejected by the reader.
Requires: executable=go; scope=test

from drydock_go import run_test

for name in ("TestRejectEmptyHeaderName", "TestRejectRepeatedHeaderName"):
    result = run_test("./internal/csvdata", name)
    print(result.stdout)
    print(result.stderr)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC reject-unit-headers ===

=== AC reject-unit-ragged-and-valid ===
Intent: Ragged records are rejected while header-only and newline-terminated files are accepted.
Requires: executable=go; scope=test

from drydock_go import run_test

for name in (
    "TestRejectRaggedRecord",
    "TestHeaderOnlyFileAccepted",
    "TestTrailingNewlineAddsNoRecord",
):
    result = run_test("./internal/csvdata", name)
    print(result.stdout)
    print(result.stderr)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC reject-unit-ragged-and-valid ===

=== AC reject-cli-ragged ===
Intent: The executable exits 3 with empty standard output for a ragged record.
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "ragged.csv")
    with open(path, "w", newline="") as handle:
        handle.write("a,b\n1,2\n3,4,5\n6,7\n")
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3, (code, out, err)
assert out == "", out
=== END AC reject-cli-ragged ===

=== AC reject-cli-empty-file ===
Intent: The executable exits 3 with empty standard output for a zero-byte file.
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "empty.csv")
    with open(path, "w", newline="") as handle:
        handle.write("")
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3, (code, out, err)
assert out == "", out
=== END AC reject-cli-empty-file ===

=== AC reject-cli-blank-lines-only ===
Intent: The executable exits 3 with empty standard output for a file holding only blank lines.
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "blank.csv")
    with open(path, "w", newline="") as handle:
        handle.write("\n\n\n")
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3, (code, out, err)
assert out == "", out
=== END AC reject-cli-blank-lines-only ===

=== AC reject-cli-repeated-header ===
Intent: The executable exits 3 with empty standard output when a header name repeats.
Requires: executable=go; scope=test

import csv
import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "dup.csv")
    with open(path, "w", newline="") as handle:
        writer = csv.writer(handle, lineterminator="\n")
        writer.writerow(["id", "name", "id"])
        writer.writerow(["1", "x", "2"])
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3, (code, out, err)
assert out == "", out
=== END AC reject-cli-repeated-header ===

=== AC reject-cli-empty-header-name ===
Intent: The executable exits 3 with empty standard output when a header name is empty.
Requires: executable=go; scope=test

import csv
import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "emptyname.csv")
    with open(path, "w", newline="") as handle:
        writer = csv.writer(handle, lineterminator="\n")
        writer.writerow(["id", "", "name"])
        writer.writerow(["1", "x", "y"])
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3, (code, out, err)
assert out == "", out
=== END AC reject-cli-empty-header-name ===

## User Acceptance

- None.

## Guardrails

- A rejected command never writes anything to standard output.
- The diagnostic goes to standard error only and is never part of any contract.
- Rejection logic lives in `internal/csvdata`; `internal/cli` only maps it to exit 3.
- Only the standard library is used; supplied files under `sources/` are never modified.
