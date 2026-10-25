# FEATURE: Profile Counts

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Report the rows, nulls, and distinct counts for every column of a CSV file. |
| Depends On  | FEATURE-ValueTrimming.md |
| Provides    | profile fields rows, nulls, distinct |
| Consumes    | trimmed header names and values from FEATURE-ValueTrimming.md |

## Purpose

For each column, the profile row carries three counts computed from the trimmed values of that column: `rows`, `nulls`, and `distinct`. Later stories (types, statistics, frequency) reuse the null rule defined here.

## Behavior

- `rows` is the number of data records in the file (the header is not counted). Every column of a file reports the same `rows`.
- A value is **null** when, after trimming, it is empty or equals `NA`, `N/A`, or `NULL` ignoring case. Every other value is non-null.
- `nulls` is the number of null values in the column.
- `distinct` is the number of different non-null values in the column, compared as exact strings after trimming. Comparison is case-sensitive, so `Apple` and `apple` are two values. Null values are never counted.
- A file with a header and no data records reports `rows` 0, `nulls` 0, and `distinct` 0 for every column.
- The null rule is exposed as a single reusable function in `internal/profile` so that types, statistics, and frequency use one definition.
- All counts are computed from the file given. No count is hardcoded.

## Programmatic Acceptance

=== AC counts-rows-nulls-distinct ===
Intent: rows, nulls, and distinct are computed from the file, with padded values trimmed before comparison.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def profile(lines):
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write("".join(line + "\n" for line in lines))
        code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}


values = ["apple", " apple ", "pear", "Apple", "pear", "fig", "fig", "fig"]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
trimmed = [v.strip() for v in values]
assert rows["v"]["rows"] == str(len(values))
assert rows["v"]["nulls"] == str(0)
assert rows["v"]["distinct"] == str(len(set(trimmed)))
assert rows["id"]["rows"] == str(len(values))
assert rows["id"]["distinct"] == str(len(values))
=== END AC counts-rows-nulls-distinct ===

=== AC counts-null-spellings ===
Intent: Empty, NA, N/A, and NULL in any letter case are null and are excluded from distinct.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def profile(lines):
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write("".join(line + "\n" for line in lines))
        code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}


null_spellings = ["", "NA", "na", "N/A", "n/a", "NULL", "null", "Null", "   ", " NA "]
real_values = ["x", "y", "x"]
values = null_spellings + real_values
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
assert rows["v"]["rows"] == str(len(values))
assert rows["v"]["nulls"] == str(len(null_spellings))
assert rows["v"]["distinct"] == str(len(set(real_values)))
=== END AC counts-null-spellings ===

=== AC counts-non-null-lookalikes ===
Intent: Values that merely resemble a null spelling are not null and count as distinct values.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def profile(lines):
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write("".join(line + "\n" for line in lines))
        code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}


values = ["NAN", "N A", "NULLS", "none", "0", "NA"]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
non_null = [v for v in values if v != "NA"]
assert rows["v"]["nulls"] == str(len(values) - len(non_null))
assert rows["v"]["distinct"] == str(len(set(non_null)))
=== END AC counts-non-null-lookalikes ===

=== AC counts-all-null-column ===
Intent: A column whose values are all null reports distinct 0 and nulls equal to rows.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def profile(lines):
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write("".join(line + "\n" for line in lines))
        code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}


values = ["", "NA", "NULL", " "]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
assert rows["v"]["rows"] == str(len(values))
assert rows["v"]["nulls"] == rows["v"]["rows"]
assert rows["v"]["distinct"] == str(0)
=== END AC counts-all-null-column ===

=== AC counts-header-only ===
Intent: A file with a header and no data records reports zero counts for every column and exits 0.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

names = ["sku", "price", "in_stock"]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as f:
        f.write(",".join(names) + "\n")
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
parsed = list(csv.DictReader(io.StringIO(out)))
assert [r["column"] for r in parsed] == names
for r in parsed:
    assert r["rows"] == str(0)
    assert r["nulls"] == str(0)
    assert r["distinct"] == str(0)
=== END AC counts-header-only ===

## User Acceptance

- None.

## Guardrails

- Counts are computed from the file examined; no count, type, or value list is hardcoded.
- Application code never reads `sources/run.sh` or any supplied data file.
- Null values are never counted by `distinct`.
