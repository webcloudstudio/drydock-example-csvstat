# FEATURE: Column Select

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Profile only the column named by --column and reject an absent column name with exit 3. |
| Depends On  | FEATURE-CliOptions.md, FEATURE-NullTrim.md, ARCHITECTURE.md |
| Provides    | --column behavior |
| Consumes    | --column, reader.Trim |

## Purpose

`--column <name>` narrows the profile to a single column. The name is matched against the trimmed header names, exactly and case-sensitively. A name absent from the header is a rejected input.

## Behavior

- The `--column` value is compared to each trimmed header name by exact, case-sensitive string equality.
- A match produces the profile header plus one row for that column only, with statistics computed exactly as in the full profile.
- No match exits 3, writes nothing to standard output, and writes one diagnostic to standard error naming the file and the missing column.
- The check runs after the file has been read and validated, so a malformed file still exits 3 for its own reason.

## Programmatic Acceptance

=== AC column-select-one-row ===
Intent: --column outputs the profile header and a single row for the named column.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

chosen = "qty"
work = tempfile.mkdtemp()
path = os.path.join(work, "data.csv")
with open(path, "w", encoding="utf-8") as handle:
    handle.write("id,qty,note\n1,10,a\n2,20,b\n")

code, out, err = cli_runner("--column", chosen, path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert rows[1][0] == chosen, rows
=== END AC column-select-one-row ===

=== AC column-select-matches-full-profile-row ===
Intent: The row produced by --column equals that column's row in the full profile.
Requires: executable=go; scope=test

import csv
import io

from drydock_go import cli_runner

chosen = "salary"
path = "sources/employees.csv"

code_all, out_all, err_all = cli_runner(path)
print(err_all)
assert code_all == 0, (code_all, err_all)
full = {row[0]: row for row in csv.reader(io.StringIO(out_all))}

code_one, out_one, err_one = cli_runner("--column", chosen, path)
print(out_one)
print(err_one)
assert code_one == 0, (code_one, err_one)
rows = list(csv.reader(io.StringIO(out_one)))
assert len(rows) == 2, rows
assert rows[1] == full[chosen], (rows[1], full[chosen])
=== END AC column-select-matches-full-profile-row ===

=== AC column-select-absent-name ===
Intent: A column name not in the header exits 3 with nothing on standard output.
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

work = tempfile.mkdtemp()
path = os.path.join(work, "data.csv")
with open(path, "w", encoding="utf-8") as handle:
    handle.write("id,qty\n1,10\n")

code, out, err = cli_runner("--column", "nosuch", path)
print(out)
print(err)
assert code == 3, (code, out, err)
assert out == "", out
=== END AC column-select-absent-name ===

=== AC column-select-case-sensitive ===
Intent: Column names match case-sensitively, so a differently cased name is absent and exits 3.
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

actual = "qty"
work = tempfile.mkdtemp()
path = os.path.join(work, "data.csv")
with open(path, "w", encoding="utf-8") as handle:
    handle.write("id," + actual + "\n1,10\n")

code, out, err = cli_runner("--column", actual.upper(), path)
print(out)
print(err)
assert code == 3, (code, out, err)
assert out == "", out
=== END AC column-select-case-sensitive ===

=== AC column-select-trimmed-header ===
Intent: A header name padded with spaces in the file is matched by its trimmed name.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

chosen = "qty"
work = tempfile.mkdtemp()
path = os.path.join(work, "data.csv")
with open(path, "w", encoding="utf-8") as handle:
    handle.write("id,  " + chosen + "  \n1,10\n")

code, out, err = cli_runner("--column", chosen, path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert rows[1][0] == chosen, rows
=== END AC column-select-trimmed-header ===

## User Acceptance

- None.

## Guardrails

- Matching is exact and case-sensitive on trimmed header names; there is no case folding or prefix matching.
- An absent column never produces partial standard output.
