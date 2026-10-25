# FEATURE: Frequency List

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | List the most frequent non-null values of one column with --freq. |
| Depends On  | FEATURE-CLI-Unknown-Column.md, FEATURE-Null-Values.md |
| Provides    | csvstat --freq <name> <file> |
| Consumes    | column lookup, csvdata.IsNull |

## Purpose

`csvstat --freq <name> <file>` lists the distinct non-null values of the named column with their
counts instead of profiling.

## Behavior

- Output header is `value,count`.
- One row per distinct non-null trimmed value, ordered by count descending, then value ascending by
  byte value.
- Null values (empty, `NA`, `N/A`, `NULL` ignoring case) are never listed.
- Values are trimmed before counting and compared as exact strings, so case differences are
  different values.
- At most 5 rows are written by default; the `--top` option is owned by FEATURE-Frequency-Top.md.
- A column name absent from the header exits 3 through the shared lookup.

## Programmatic Acceptance

=== AC freq-department-order ===
Intent: --freq department on employees.csv lists every distinct value ordered by count descending then value by byte order.

import csv
import io
from drydock_go import cli_runner

def load(path):
    with open(path, newline="", encoding="utf-8") as f:
        rows = list(csv.reader(f))
    return [[c.strip() for c in r] for r in rows]

def is_null(v):
    return v == "" or v.lower() in ("na", "n/a", "null")

def expected(path, col):
    rows = load(path)
    i = rows[0].index(col)
    counts = {}
    for r in rows[1:]:
        if not is_null(r[i]):
            counts[r[i]] = counts.get(r[i], 0) + 1
    return sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))

path = "sources/employees.csv"
col = "department"
code, out, err = cli_runner("--freq", col, path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
want = expected(path, col)
assert len(want) <= 5
assert got[0] == ["value", "count"]
assert got[1:] == [[v, str(c)] for v, c in want]
=== END AC freq-department-order ===

=== AC freq-default-limit-five ===
Intent: Without --top, --freq writes at most five value rows, the first five of the full ordering.

import csv
import io
from drydock_go import cli_runner

def load(path):
    with open(path, newline="", encoding="utf-8") as f:
        rows = list(csv.reader(f))
    return [[c.strip() for c in r] for r in rows]

def is_null(v):
    return v == "" or v.lower() in ("na", "n/a", "null")

def expected(path, col):
    rows = load(path)
    i = rows[0].index(col)
    counts = {}
    for r in rows[1:]:
        if not is_null(r[i]):
            counts[r[i]] = counts.get(r[i], 0) + 1
    return sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))

path = "sources/employees.csv"
col = "name"
want = expected(path, col)
assert len(want) > 5
code, out, err = cli_runner("--freq", col, path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert len(got) == 1 + 5
assert got[1:] == [[v, str(c)] for v, c in want[:5]]
=== END AC freq-default-limit-five ===

=== AC freq-excludes-nulls ===
Intent: Null values, including NA and empty cells, are never listed; the remaining values and counts match the file.

import csv
import io
from drydock_go import cli_runner

def load(path):
    with open(path, newline="", encoding="utf-8") as f:
        rows = list(csv.reader(f))
    return [[c.strip() for c in r] for r in rows]

def is_null(v):
    return v == "" or v.lower() in ("na", "n/a", "null")

def expected(path, col):
    rows = load(path)
    i = rows[0].index(col)
    counts = {}
    for r in rows[1:]:
        if not is_null(r[i]):
            counts[r[i]] = counts.get(r[i], 0) + 1
    return sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))

path = "sources/employees.csv"
col = "manager_id"
want = expected(path, col)
assert len(want) <= 5
code, out, err = cli_runner("--freq", col, path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert all(not is_null(r[0]) for r in got[1:])
assert got[1:] == [[v, str(c)] for v, c in want]
=== END AC freq-excludes-nulls ===

=== AC freq-ties-by-byte-value ===
Intent: Values with equal counts are ordered ascending by byte value, and padded values are trimmed before counting.

import csv
import io
from drydock_go import cli_runner

def load(path):
    with open(path, newline="", encoding="utf-8") as f:
        rows = list(csv.reader(f))
    return [[c.strip() for c in r] for r in rows]

def is_null(v):
    return v == "" or v.lower() in ("na", "n/a", "null")

def expected(path, col):
    rows = load(path)
    i = rows[0].index(col)
    counts = {}
    for r in rows[1:]:
        if not is_null(r[i]):
            counts[r[i]] = counts.get(r[i], 0) + 1
    return sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))

for path, col in (("sources/sensors.csv", "sensor"), ("sources/orders.csv", "customer")):
    want = expected(path, col)
    assert len(want) <= 5
    code, out, err = cli_runner("--freq", col, path)
    print(out)
    print(err)
    assert code == 0
    got = list(csv.reader(io.StringIO(out)))
    assert got[1:] == [[v, str(c)] for v, c in want], (path, col)
=== END AC freq-ties-by-byte-value ===

=== AC freq-case-sensitive-values ===
Intent: Values that differ only by letter case are distinct, and the counts come from the file supplied by the criterion.

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

values = ["Red", "red", "Red", "RED", "red", "Red"]
d = tempfile.mkdtemp()
path = os.path.join(d, "colors.csv")
with open(path, "w", newline="", encoding="utf-8") as f:
    w = csv.writer(f, lineterminator="\n")
    w.writerow(["color"])
    for v in values:
        w.writerow([v])
counts = {}
for v in values:
    counts[v] = counts.get(v, 0) + 1
want = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))
code, out, err = cli_runner("--freq", "color", path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert got[1:] == [[v, str(c)] for v, c in want]
=== END AC freq-case-sensitive-values ===

## User Acceptance

- None.

## Guardrails

- Null values are never listed in frequency output.
- Frequency output is never combined with profile output.
- Application code never reads `sources/run.sh` or any supplied data file and never hardcodes a value list.
