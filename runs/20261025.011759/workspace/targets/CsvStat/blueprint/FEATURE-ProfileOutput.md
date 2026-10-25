# FEATURE: ProfileOutput

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Write the profile table with one row per column in file order. |
| Depends On  | FEATURE-Counts.md, FEATURE-MinMax.md, FEATURE-SumMean.md, FEATURE-Median.md, FEATURE-TypeDate.md, FEATURE-TypeOther.md |
| Provides    | profile table on standard output |
| Consumes    | internal/profile |

## Purpose

Assemble the per-column type and statistics into the CSV table the command writes when profiling a file.

## Behavior

- With a file argument and no `--freq`, write the header `column,type,rows,nulls,distinct,min,max,sum,mean,median`.
- Then write one row per column in file order, using the trimmed header name.
- Output uses `encoding/csv`'s writer with `\n` line endings, so a field containing a comma, a double quote, or a newline is quoted per RFC 4180.
- The command exits 0 on success.
- On any failure nothing is written to standard output; the table is built fully before the first byte is written.

## Interface

- `internal/cli.Run` calls the reader, null classifier, type inference, and the statistics, then writes the table to its `stdout` writer.
- The project test suite carries `TestProfileHeader`, `TestProfileRowOrder`, and `TestProfileQuoting` in `internal/cli`.

## Programmatic Acceptance

=== AC profile-columns-and-stats ===
Intent: Profiling a generated file yields one row per column in file order with counts, types, and statistics matching values computed from that file's own inputs.

import csv
import io
import math
import os
import tempfile
from fractions import Fraction

from drydock_go import cli_runner


def fmt(fr):
    n = math.floor(abs(fr) * 100 + Fraction(1, 2))
    s = f"{n // 100}.{n % 100:02d}"
    return "-" + s if fr < 0 and n != 0 else s


header = ["id", "amt", "when", "flag", "word", "gone"]
records = [
    ["3", "1.50", "2025-01-05", "true", "pear", "NA"],
    ["1", "-2.25", "2024-12-31", "FALSE", "apple", ""],
    ["2", "10.00", "2025-03-01", "true", "pear", "null"],
    ["10", "N/A", "2025-02-14", "true", "fig", "NULL"],
    ["7", "4.75", "NA", "false", "apple", "N/A"],
]
cols = list(zip(*records))

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as fh:
        w = csv.writer(fh)
        w.writerow(header)
        w.writerows(records)
    code, out, err = cli_runner(path)

print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[0][0] == "column"
assert [r[0] for r in rows[1:]] == header

nulls = {"", "na", "n/a", "null"}


def nonnull(col):
    return [v for v in col if v.strip().lower() not in nulls]


by = {r[0]: r for r in rows[1:]}

ids = [Fraction(int(v)) for v in nonnull(cols[0])]
r = by["id"]
assert r[1] == "integer"
assert r[2] == str(len(records))
assert r[3] == str(len(cols[0]) - len(ids))
assert r[4] == str(len(set(nonnull(cols[0]))))
assert r[5] == str(int(min(ids)))
assert r[6] == str(int(max(ids)))
assert r[7] == str(int(sum(ids)))
assert r[8] == fmt(sum(ids) / len(ids))
s = sorted(ids)
mid = len(s) // 2
med = s[mid] if len(s) % 2 else (s[mid - 1] + s[mid]) / 2
assert r[9] == fmt(med)

amts = [Fraction(v) for v in nonnull(cols[1])]
r = by["amt"]
assert r[1] == "number"
assert r[3] == str(len(cols[1]) - len(amts))
assert r[5] == fmt(min(amts))
assert r[6] == fmt(max(amts))
assert r[7] == fmt(sum(amts))
assert r[8] == fmt(sum(amts) / len(amts))
s = sorted(amts)
mid = len(s) // 2
med = s[mid] if len(s) % 2 else (s[mid - 1] + s[mid]) / 2
assert r[9] == fmt(med)

dates = nonnull(cols[2])
r = by["when"]
assert r[1] == "date"
assert r[3] == str(len(cols[2]) - len(dates))
assert r[5] == min(dates)
assert r[6] == max(dates)
assert r[7:10] == ["", "", ""]

r = by["flag"]
assert r[1] == "boolean"
assert r[4] == str(len(set(cols[3])))
assert r[5] == "false"
assert r[6] == "true"
assert r[7:10] == ["", "", ""]

words = nonnull(cols[4])
r = by["word"]
assert r[1] == "text"
assert r[5] == min(words, key=lambda v: v.encode())
assert r[6] == max(words, key=lambda v: v.encode())
assert r[7:10] == ["", "", ""]

r = by["gone"]
assert r[1] == "empty"
assert r[2] == str(len(records))
assert r[3] == str(len(records))
assert r[4] == "0"
assert r[5:10] == ["", "", "", "", ""]
=== END AC profile-columns-and-stats ===

=== AC profile-quoting-and-trim ===
Intent: A value containing a comma is quoted on output, and padding spaces in header names and values are trimmed.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

name = "label"
value = "red, green"
raw = f"  {name}  ,n\n  {value}  ,1\n{value},2\n"

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as fh:
        fh.write(raw.replace(value, '"' + value + '"'))
    code, out, err = cli_runner(path)

print(out)
print(err)
assert code == 0, (code, out, err)
assert out.endswith("\n")
assert "\r" not in out
rows = list(csv.reader(io.StringIO(out)))
by = {r[0]: r for r in rows[1:]}
assert name in by
assert by[name][1] == "text"
assert by[name][4] == "1"
assert by[name][5] == value
assert by[name][6] == value
=== END AC profile-quoting-and-trim ===

=== AC profile-header-only ===
Intent: A header-only file reports every column as empty with zero counts and exits 0.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

header = ["sku", "price", "in_stock"]

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as fh:
        csv.writer(fh).writerow(header)
    code, out, err = cli_runner(path)

print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert [r[0] for r in rows[1:]] == header
for r in rows[1:]:
    assert r[1] == "empty"
    assert r[2:5] == ["0", "0", "0"]
    assert r[5:10] == ["", "", "", "", ""]
=== END AC profile-header-only ===

=== AC profile-failure-silent-stdout ===
Intent: A rejected file writes nothing to standard output and exits 3.

import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "ragged.csv")
    with open(path, "w", newline="") as fh:
        fh.write("a,b\n1,2\n3,4,5\n")
    code, out, err = cli_runner(path)

print(out)
print(err)
assert code == 3, (code, out, err)
assert out == ""
=== END AC profile-failure-silent-stdout ===

## User Acceptance

- None.

## Guardrails

- Every value in the table is computed from the file examined; no statistic, type, or value list is hardcoded.
- The application never reads `sources/run.sh` or any supplied data file.
- Nothing is written to standard output when the command fails.
