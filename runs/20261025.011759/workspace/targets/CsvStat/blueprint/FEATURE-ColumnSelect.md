# FEATURE: Column Select

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Profile only the column named by --column and reject an absent name with exit 3. |
| Depends On  | FEATURE-ProfileOutput.md, FEATURE-Arguments.md, FEATURE-InputRejection.md |
| Provides    | --column behavior |
| Consumes    | profile table on standard output |

## Purpose

`csvstat --column <name> <file>` writes the profile header followed by the single row for the named
column. The row is identical to the row the full profile writes for that column. A name that is not
in the header is a rejected input.

## Behavior

- Output is the header `column,type,rows,nulls,distinct,min,max,sum,mean,median` and exactly one
  row, for the named column, written with the same `encoding/csv` writer as the full profile.
- The name is compared to the trimmed header names as an exact string.
- A name absent from the header exits 3, writes nothing to standard output, and writes one
  diagnostic line naming the file to standard error.
- The whole file is still read and validated first, so a ragged file exits 3 regardless of the
  column named.

## Programmatic Acceptance

=== AC column-select-matches-full-profile ===
Intent: A named column yields the profile header plus exactly the row the full profile gives that column.
Requires: executable=go; scope=test

import csv
import io
from drydock_go import cli_runner

data = "sources/employees.csv"
code, out, err = cli_runner(data)
assert code == 0, (code, out, err)
full = list(csv.reader(io.StringIO(out)))
header = full[0]
by_name = {row[0]: row for row in full[1:]}

for name in ("salary", "id", "remote", "hire_date"):
    code, out, err = cli_runner("--column", name, data)
    assert code == 0, (name, code, out, err)
    rows = list(csv.reader(io.StringIO(out)))
    assert len(rows) == 2, (name, rows)
    assert rows[0] == header, (name, rows[0], header)
    assert rows[1] == by_name[name], (name, rows[1], by_name[name])
=== END AC column-select-matches-full-profile ===

=== AC column-select-quoted-field ===
Intent: A column whose value contains a comma is quoted by the writer and parses back to the same row as the full profile.
Requires: executable=go; scope=test

import csv
import io
from drydock_go import cli_runner

data = "sources/employees.csv"
code, out, err = cli_runner(data)
assert code == 0, (code, out, err)
full = list(csv.reader(io.StringIO(out)))
by_name = {row[0]: row for row in full[1:]}

name = "notes"
code, out, err = cli_runner("--column", name, data)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert rows[1] == by_name[name], (rows[1], by_name[name])
assert any("," in field for field in rows[1]), rows[1]
=== END AC column-select-quoted-field ===

=== AC column-select-absent-name ===
Intent: A column name missing from the header exits 3 and writes nothing to standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("--column", "nosuch", "sources/orders.csv")
assert code == 3, (code, out, err)
assert out == "", out
=== END AC column-select-absent-name ===

=== AC column-select-trimmed-header-name ===
Intent: The column name is matched against trimmed header names, so a padded header still resolves.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

name = "price"
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "padded.csv")
    with open(path, "w", newline="") as handle:
        handle.write(f"id,  {name}  \n1,2.50\n2,3.50\n")
    code, out, err = cli_runner("--column", name, path)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(io.StringIO(out)))
    assert len(rows) == 2, rows
    assert rows[1][0] == name, rows
=== END AC column-select-trimmed-header-name ===

## User Acceptance

- None.

## Guardrails

- The row for a named column is computed from the file, never hardcoded.
- Nothing is written to standard output when the command fails.
- Only the exact-string, trimmed header name selects a column; no prefix or case-insensitive match.
