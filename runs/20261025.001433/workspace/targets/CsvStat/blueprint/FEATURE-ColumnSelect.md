# FEATURE: Column Select

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Profile only the column named by --column, writing the profile header and one row. |
| Depends On  | FEATURE-CliParsing.md, FEATURE-CsvReading.md, FEATURE-DecimalFormat.md |
| Provides    | csvstat --column <name> <file> |
| Consumes    | csvstat <file> profile rows, trimmed header names |

## Purpose

Lets a user profile a single column of a wide file without reading every row of the full profile.

## Trigger

`csvstat --column <name> <file>`. `--column` cannot be combined with `--freq` (a usage error owned
by `FEATURE-UsageErrors.md`).

## Behavior

- The output is the profile header `column,type,rows,nulls,distinct,min,max,sum,mean,median`
  followed by exactly one row, the row of the named column, identical in every field to the row the
  full profile writes for that column.
- The `<name>` argument matches a trimmed header name exactly and case-sensitively. Header names
  are trimmed of leading and trailing spaces before matching, so a file header ` Name ` is
  selected by `Name`.
- The statistics of the other columns are not written; the named column's values are still
  validated like any other (a ragged record still rejects the file).
- A name that is not in the header is rejected with exit 3 (owned by `FEATURE-ColumnAbsent.md`).
- Output is CSV with `\n` line endings produced by `encoding/csv`'s writer, so a field containing a
  comma, quote, or newline is quoted.
- Success exits 0.

## Reads

- The CSV file given on the command line.

## Writes

- Standard output only.

## Programmatic Acceptance

=== AC colselect-matches-full-profile-row ===
Intent: For every column of the supplied files, --column writes the profile header and exactly the row the full profile writes for that column.
Suite: scoped
Requires: executable=go; scope=test

import csv
import io

from drydock_go import cli_runner

for path in ("sources/employees.csv", "sources/sensors.csv", "sources/orders.csv", "sources/header_only.csv"):
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0
    full = list(csv.reader(io.StringIO(out)))
    header = full[0]
    assert len(full) > 1
    for row in full[1:]:
        name = row[0]
        code, out, err = cli_runner("--column", name, path)
        print(out)
        print(err)
        assert code == 0
        picked = list(csv.reader(io.StringIO(out)))
        assert len(picked) == 2
        assert picked[0] == header
        assert picked[1] == row
=== END AC colselect-matches-full-profile-row ===

=== AC colselect-only-named-column ===
Intent: The output of --column contains one data row whose first cell is the name supplied, even for the last column of the file.
Suite: scoped
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def run(text, *opts):
    with tempfile.TemporaryDirectory() as d:
        p = os.path.join(d, "data.csv")
        with open(p, "w", newline="") as fh:
            fh.write(text)
        code, out, err = cli_runner(*opts, p)
        assert code == 0
        return code, out, err


names = ["alpha", "beta", "gamma", "delta"]
text = ",".join(names) + "\n1,x,true,2025-01-02\n2,y,false,2025-02-03\n"
for name in names:
    code, out, err = run(text, "--column", name)
    print(out)
    print(err)
    assert code == 0
    rows = list(csv.reader(io.StringIO(out)))
    assert len(rows) == 2
    assert rows[1][0] == name
=== END AC colselect-only-named-column ===

=== AC colselect-trimmed-header-and-values ===
Intent: A padded header is selected by its trimmed name and the padded values are trimmed before they are profiled.
Suite: scoped
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def run(text, *opts):
    with tempfile.TemporaryDirectory() as d:
        p = os.path.join(d, "data.csv")
        with open(p, "w", newline="") as fh:
            fh.write(text)
        code, out, err = cli_runner(*opts, p)
        assert code == 0
        return code, out, err


name = "Name"
text = "id,  " + name + "  \n1,  Acme  \n2,Acme\n3,Globex\n"
code, out, err = run(text, "--column", name)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2
assert rows[1][0] == name
assert rows[1][2] == "3"
assert rows[1][4] == "2"
=== END AC colselect-trimmed-header-and-values ===

=== AC colselect-output-format ===
Intent: Selected-column output uses newline line endings only and quotes a field containing a comma as one value.
Suite: scoped
Requires: executable=go; scope=test

import csv
import io

from drydock_go import cli_runner

path = "sources/employees.csv"
name = "notes"
code, out, err = cli_runner("--column", name, path)
print(out)
print(err)
assert code == 0
assert out.endswith("\n")
assert "\r" not in out
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2
assert len(rows[0]) == len(rows[1])
with open(path, newline="") as fh:
    recs = list(csv.reader(fh))
i = recs[0].index(name)
cells = [r[i].strip() for r in recs[1:]]
comma_cells = [c for c in cells if "," in c]
assert comma_cells
assert max(c for c in cells if c.lower() not in ("", "na", "n/a", "null")) == rows[1][6]
=== END AC colselect-output-format ===

=== AC colselect-header-only-file ===
Intent: Selecting a column of a file with a header and no rows succeeds and reports zero rows.
Suite: scoped
Requires: executable=go; scope=test

import csv
import io

from drydock_go import cli_runner

path = "sources/header_only.csv"
with open(path, newline="") as fh:
    name = next(csv.reader(fh))[1].strip()
code, out, err = cli_runner("--column", name, path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2
assert rows[1][0] == name
assert rows[1][1] == "empty"
assert rows[1][2] == str(0)
=== END AC colselect-header-only-file ===

## User Acceptance

- None.

## Guardrails

- `--column` never writes more than one data row.
- Column names are not matched case-insensitively and are not fuzzy matched.
- Application code never reads `sources/run.sh` or any supplied data file and never embeds an answer.
