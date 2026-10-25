# FEATURE: CSV Reading and Profile Output

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Read a well-formed CSV file and write the profile header and one row per column in file order. |
| Depends On  | FEATURE-CliParsing.md, ARCHITECTURE.md |
| Provides    | profile output header and rows |
| Consumes    | csvstat [--column <name>] <file> |

## Purpose

`internal/csvio` opens the file and reads it with `encoding/csv` as RFC 4180 defines it; the first
record is the header. The profile mode writes CSV to standard output with `\n` line endings.

## Behavior

1. The reader sets `FieldsPerRecord = -1` so field-count validation is done by the application,
   which is owned by the rejection story.
2. Standard output starts with the header `column,type,rows,nulls,distinct,min,max,sum,mean,median`,
   then one row per column in file order. Each row has exactly ten fields; the first is the column
   name. The remaining fields are produced by the later counts, types, and statistics stories; this
   story establishes the row structure.
3. A quoted field containing a comma is read as one value. A field containing a comma, a double
   quote, or a newline is quoted on output by the `encoding/csv` writer.
4. A file with a header and no data records succeeds with one row per header column.
5. Bytes are compared as bytes: a byte-order mark is not stripped and the file is not rejected for
   invalid UTF-8.
6. Output is written to a buffer and copied to standard output only after reading succeeds.

## Programmatic Acceptance

=== AC csvread-header-and-columns ===
Intent: Profiling a well-formed file exits 0 and writes the declared header then one ten-field row per input column in file order.

import csv
import io
from drydock_go import cli_runner

path = "sources/employees.csv"
with open(path, newline="") as handle:
    input_header = next(csv.reader(handle))
declared = "column,type,rows,nulls,distinct,min,max,sum,mean,median".split(",")
code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == declared
assert [row[0] for row in rows[1:]] == input_header
for row in rows:
    assert len(row) == len(declared)
=== END AC csvread-header-and-columns ===

=== AC csvread-line-endings ===
Intent: Profile output uses newline line endings only.

from drydock_go import cli_runner

code, out, err = cli_runner("sources/orders.csv")
assert code == 0
assert "\r" not in out
assert out.endswith("\n")
=== END AC csvread-line-endings ===

=== AC csvread-quoted-comma ===
Intent: A quoted field containing a comma is one value, so the file keeps its declared column count.

import csv
import io
import tempfile
from pathlib import Path
from drydock_go import cli_runner

names = ["a", "b"]
with tempfile.TemporaryDirectory() as tmp:
    target = Path(tmp) / "quoted.csv"
    target.write_text(names[0] + "," + names[1] + "\n" + '"x, y",2\n')
    code, out, err = cli_runner(str(target))
    print(out)
    print(err)
    assert code == 0
    rows = list(csv.reader(io.StringIO(out)))
    assert [row[0] for row in rows[1:]] == names
=== END AC csvread-quoted-comma ===

=== AC csvread-header-only ===
Intent: A file with a header and no data records succeeds with one row per header column.

import csv
import io
from drydock_go import cli_runner

path = "sources/header_only.csv"
with open(path, newline="") as handle:
    input_header = next(csv.reader(handle))
code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert [row[0] for row in rows[1:]] == input_header
=== END AC csvread-header-only ===

=== AC csvread-bom-kept ===
Intent: A byte-order mark is not stripped from the first header name.

import csv
import io
import tempfile
from pathlib import Path
from drydock_go import cli_runner

bom = "\ufeff"
first = "a"
second = "b"
payload = (bom + first + "," + second + "\n1,2\n").encode("utf-8")
with tempfile.TemporaryDirectory() as tmp:
    target = Path(tmp) / "bom.csv"
    target.write_bytes(payload)
    code, out, err = cli_runner(str(target))
    assert code == 0
    rows = list(csv.reader(io.StringIO(out)))
    assert rows[1][0] == bom + first
    assert rows[2][0] == second
=== END AC csvread-bom-kept ===

## User Acceptance

- None.

## Guardrails

- Application code never reads `sources/run.sh` or any supplied data file by name.
- A byte-order mark is not stripped and invalid UTF-8 is not a rejection reason.
- Standard output receives nothing until the whole file has been read and validated.
- Standard library only.
