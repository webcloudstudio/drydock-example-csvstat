# FEATURE: Profile Rows

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Emit the profile header and one row per column in file order, filling column and rows. |
| Depends On  | ARCHITECTURE.md, FEATURE-CSV-Reading.md, FEATURE-Null-Values.md, FEATURE-CSV-Rejection.md, FEATURE-CLI-Option-Validation.md |
| Provides    | profile output header |
| Consumes    | csvdata.Read, csvdata.IsNull |

## Purpose

This is the skeleton of the profile output. Later stories fill the remaining fields, so the output
shape is fixed here and never changes.

## Behavior

- Without `--freq`, standard output is CSV written by one `encoding/csv` writer with `\n` line
  endings.
- The first row is the header `column,type,rows,nulls,distinct,min,max,sum,mean,median` (ten fields).
- One row follows for each header column, in file order. Every row has ten fields.
- `column` is the trimmed header name. `rows` is the number of data records (not counting the
  header), written as a plain integer; a header-only file gives `0` for every column.
- The fields `type`, `nulls`, `distinct`, `min`, `max`, `sum`, `mean` and `median` are owned by
  later stories (Type-*, Profile-Counts, Profile-Extremes, Profile-Numeric-Stats). Until a later
  story is built, the field is written empty. The final application fills all of them.
- Exit `0` on success.

## Programmatic Acceptance

=== AC profile-rows-header-and-columns-in-file-order ===
Intent: The profile starts with the ten-field header and lists one row per input column in file order.

import csv
import io
import os
from drydock_go import cli_runner

data = os.path.abspath(os.path.join("sources", "employees.csv"))
with open(data, newline="") as f:
    source = list(csv.reader(f))
names = source[0]
code, out, err = cli_runner(data)
print(out)
print(err)
assert code == 0
parsed = list(csv.reader(io.StringIO(out)))
assert parsed[0] == ["column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"]
assert [row[0] for row in parsed[1:]] == names
assert all(len(row) == len(parsed[0]) for row in parsed)
=== END AC profile-rows-header-and-columns-in-file-order ===

=== AC profile-rows-count-matches-records ===
Intent: The rows field equals the number of data records for every column.

import csv
import io
import os
from drydock_go import cli_runner

for name in ("employees.csv", "sensors.csv", "orders.csv"):
    data = os.path.abspath(os.path.join("sources", name))
    with open(data, newline="") as f:
        source = list(csv.reader(f))
    record_count = len(source) - 1
    code, out, err = cli_runner(data)
    print(out)
    print(err)
    assert code == 0
    parsed = list(csv.reader(io.StringIO(out)))
    index = parsed[0].index("rows")
    assert len(parsed) - 1 == len(source[0])
    for row in parsed[1:]:
        assert row[index] == str(record_count)
=== END AC profile-rows-count-matches-records ===

=== AC profile-rows-header-only-file ===
Intent: A header-only file yields one row per column with rows equal to zero.

import csv
import io
import os
from drydock_go import cli_runner

data = os.path.abspath(os.path.join("sources", "header_only.csv"))
with open(data, newline="") as f:
    source = list(csv.reader(f))
code, out, err = cli_runner(data)
print(out)
print(err)
assert code == 0
parsed = list(csv.reader(io.StringIO(out)))
index = parsed[0].index("rows")
assert [row[0] for row in parsed[1:]] == source[0]
for row in parsed[1:]:
    assert row[index] == str(len(source) - 1)
=== END AC profile-rows-header-only-file ===

=== AC profile-rows-unix-line-endings ===
Intent: Profile output uses newline line endings only.

import os
from drydock_go import cli_runner

data = os.path.abspath(os.path.join("sources", "sensors.csv"))
code, out, err = cli_runner(data)
print(err)
assert code == 0
assert "\r" not in out
assert out.endswith("\n")
=== END AC profile-rows-unix-line-endings ===

=== AC profile-rows-trimmed-column-names ===
Intent: Padded header names are written trimmed, in file order.

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

names = ["alpha", "beta", "gamma"]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "padded.csv")
    with open(path, "w", newline="") as f:
        f.write(" " + " , ".join(names) + " \n1,2,3\n4,5,6\n")
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0
    parsed = list(csv.reader(io.StringIO(out)))
    assert [row[0] for row in parsed[1:]] == names
=== END AC profile-rows-trimmed-column-names ===

## User Acceptance

- None.

## Guardrails

- The header row is fixed; no field is added, removed or reordered by any story.
- Standard output is written only after the whole file has been read and validated.
- Application code never embeds a statistic, a type, or a value list.
