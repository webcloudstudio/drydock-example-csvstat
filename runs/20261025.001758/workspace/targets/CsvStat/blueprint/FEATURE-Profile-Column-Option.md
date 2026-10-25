# FEATURE: Profile Column Option

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Profile a single named column with the --column option. |
| Depends On  | FEATURE-Profile-Extremes.md, FEATURE-CLI-Unknown-Column.md |
| Provides    | csvstat --column <name> <file> |
| Consumes    | column lookup, profile rows |

## Purpose

`csvstat --column <name> <file>` writes the profile header and exactly one row: the row of the
named column. That row is identical to the same column's row in the full profile. An unknown column
name is owned by FEATURE-CLI-Unknown-Column.md (exit 3), and combining `--column` with `--freq` is
owned by FEATURE-CLI-Option-Validation.md (exit 2).

## Behavior

- Output is the header `column,type,rows,nulls,distinct,min,max,sum,mean,median` followed by one
  row for the named column.
- The column is matched by exact trimmed header name.
- The row is computed by the same code path as the full profile, so `rows` still counts every data
  record and every statistic covers the whole column.
- A header-only file with `--column` yields the header and one row with `rows` 0.

## Programmatic Acceptance

=== AC column-row-equals-full-profile ===
Intent: For every column of employees.csv, --column writes the header plus exactly that column's row from the full profile.

import csv
import io
from drydock_go import cli_runner

path = "sources/employees.csv"
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
    assert code == 0, (name, code)
    got = list(csv.reader(io.StringIO(out)))
    assert got == [header, row], (name, got)
=== END AC column-row-equals-full-profile ===

=== AC column-only-one-row-sensors ===
Intent: On sensors.csv, --column for each column writes exactly two records and names the requested column.

import csv
import io
from drydock_go import cli_runner

path = "sources/sensors.csv"
with open(path, newline="", encoding="utf-8") as f:
    names = [c.strip() for c in next(csv.reader(f))]
for name in names:
    code, out, err = cli_runner("--column", name, path)
    print(out)
    print(err)
    assert code == 0, (name, code)
    got = list(csv.reader(io.StringIO(out)))
    assert len(got) == 2, (name, got)
    assert got[1][0] == name
=== END AC column-only-one-row-sensors ===

=== AC column-header-only-file ===
Intent: --column on a header-only file writes one row whose rows count equals the number of data records in the file.

import csv
import io
from drydock_go import cli_runner

path = "sources/header_only.csv"
with open(path, newline="", encoding="utf-8") as f:
    records = list(csv.reader(f))
data_rows = len(records) - 1
name = "price"
code, out, err = cli_runner("--column", name, path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert len(got) == 2
assert got[1][0] == name
assert got[1][2] == str(data_rows)
=== END AC column-header-only-file ===

=== AC column-matches-padded-header-orders ===
Intent: --column on orders.csv for a column with padded text values agrees with the full profile row for that column.

import csv
import io
from drydock_go import cli_runner

path = "sources/orders.csv"
code, out, err = cli_runner(path)
assert code == 0, err
full = list(csv.reader(io.StringIO(out)))
by_name = {r[0]: r for r in full[1:]}
name = "customer"
code, out, err = cli_runner("--column", name, path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert got == [full[0], by_name[name]]
=== END AC column-matches-padded-header-orders ===

## User Acceptance

- None.

## Guardrails

- `--column` never adds, removes, or reorders profile fields.
- Application code never reads `sources/run.sh` or any supplied data file and never hardcodes a statistic.
- Nothing is written to standard output when the command fails.
