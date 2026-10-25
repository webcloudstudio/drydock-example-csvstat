# FEATURE: Input Rejection

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Reject an empty, malformed, or ragged CSV file with exit 3, nothing on standard output, and one diagnostic line on standard error. |
| Depends On  | FEATURE-CsvReader.md, ARCHITECTURE.md |
| Provides    | exit code 3 for rejected input |
| Consumes    | internal/csvio.Read |

## Purpose

After reading the file, the application rejects it when any of these holds:

- the file is empty (no header record);
- a header name is empty after trimming;
- a header name is repeated;
- any data record has a different number of fields from the header;
- `encoding/csv` reports a syntax error (a `csv.ParseError`), such as a bare quote.

A rejected file exits 3, writes nothing to standard output, and writes one line to standard error naming the file and, for a record, its line number. Standard error text is never compared.

## Behavior

- Validation completes before anything is written to standard output.
- The ragged record is identified by the line number the reader recorded for it.
- A file with a header and no data records is valid and is not rejected.
- An open or read failure is not a rejection; it is exit 1 (FEATURE-ReadFailure.md).

## Programmatic Acceptance

=== AC reject-ragged-record ===
Intent: A record with more fields than the header exits 3 with nothing on standard output.

import os
from drydock_go import cli_runner

path = os.path.abspath(os.path.join("sources", "ragged.csv"))
assert os.path.isfile(path)
code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3
assert out == ""
=== END AC reject-ragged-record ===

=== AC reject-short-record ===
Intent: A record with fewer fields than the header exits 3 with nothing on standard output.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as work:
    path = os.path.join(work, "short.csv")
    with open(path, "w", encoding="utf-8", newline="") as handle:
        handle.write("a,b,c\n1,2,3\n4,5\n")
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 3
    assert out == ""
=== END AC reject-short-record ===

=== AC reject-empty-file ===
Intent: An empty file exits 3 with nothing on standard output.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as work:
    path = os.path.join(work, "empty.csv")
    with open(path, "wb"):
        pass
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 3
    assert out == ""
=== END AC reject-empty-file ===

=== AC reject-empty-header-name ===
Intent: A header with an empty name exits 3 with nothing on standard output.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as work:
    path = os.path.join(work, "emptyname.csv")
    with open(path, "w", encoding="utf-8", newline="") as handle:
        handle.write("a,,c\n1,2,3\n")
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 3
    assert out == ""
=== END AC reject-empty-header-name ===

=== AC reject-repeated-header-name ===
Intent: A header with a repeated name, including one that differs only by padding, exits 3 with nothing on standard output.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as work:
    path = os.path.join(work, "repeated.csv")
    with open(path, "w", encoding="utf-8", newline="") as handle:
        handle.write("a, a ,c\n1,2,3\n")
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 3
    assert out == ""
=== END AC reject-repeated-header-name ===

=== AC reject-valid-file-not-rejected ===
Intent: A well-formed file is not rejected, so the rejection path is not taken for every input.

import os
from drydock_go import cli_runner

path = os.path.abspath(os.path.join("sources", "employees.csv"))
assert os.path.isfile(path)
code, out, err = cli_runner(path)
print(err)
assert code != 3
=== END AC reject-valid-file-not-rejected ===

=== AC reject-header-only-not-rejected ===
Intent: A file with a header and no data records is not rejected.

import os
from drydock_go import cli_runner

path = os.path.abspath(os.path.join("sources", "header_only.csv"))
assert os.path.isfile(path)
code, out, err = cli_runner(path)
print(err)
assert code != 3
=== END AC reject-header-only-not-rejected ===

## User Acceptance

- None.

## Guardrails

- A rejected file writes nothing to standard output.
- Diagnostics go to standard error only and are never compared.
- The list of rejection conditions is exhaustive; invalid UTF-8 is not one of them.
- Do not modify any file under `sources/`.
