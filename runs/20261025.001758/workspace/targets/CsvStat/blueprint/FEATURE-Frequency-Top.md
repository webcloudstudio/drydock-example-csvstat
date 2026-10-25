# FEATURE: Frequency Top Limit

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Limit the rows of a frequency listing with the --top option. |
| Depends On  | FEATURE-Frequency-List.md |
| Provides    | csvstat --freq <name> --top <n> <file> |
| Consumes    | frequency list |

## Purpose

`--top <n>` limits a `--freq` listing to at most `n` value rows. `n` is a positive integer and
defaults to 5. Rejection of invalid `--top` values (zero, negative, non-numeric, or `--top` without
`--freq`) is owned by FEATURE-CLI-Option-Validation.md and exits 2.

## Behavior

- At most `n` value rows are written after the `value,count` header.
- Rows are the first `n` of the full ordering: count descending, then value ascending by byte value.
- A `n` larger than the number of distinct non-null values writes every value.

## Programmatic Acceptance

=== AC top-limits-rows ===
Intent: --top n writes exactly the first n rows of the full frequency ordering for several values of n.

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
for n in (1, 2, 3, 8):
    code, out, err = cli_runner("--freq", col, "--top", str(n), path)
    print(out)
    print(err)
    assert code == 0, n
    got = list(csv.reader(io.StringIO(out)))
    assert len(got) == 1 + n
    assert got[1:] == [[v, str(c)] for v, c in want[:n]], n
=== END AC top-limits-rows ===

=== AC top-larger-than-distinct ===
Intent: A --top larger than the distinct count writes every distinct non-null value and no more.

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
want = expected(path, col)
n = len(want) + 50
code, out, err = cli_runner("--freq", col, "--top", str(n), path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert got[1:] == [[v, str(c)] for v, c in want]
=== END AC top-larger-than-distinct ===

=== AC top-above-default-lists-more ===
Intent: --top above the default of 5 lists more than five rows when the column has that many distinct values.

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

path = "sources/sensors.csv"
col = "temp_c"
want = expected(path, col)
n = 7
assert len(want) >= n
code, out, err = cli_runner("--freq", col, "--top", str(n), path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert got[1:] == [[v, str(c)] for v, c in want[:n]]
=== END AC top-above-default-lists-more ===

## User Acceptance

- None.

## Guardrails

- `--top` never changes the ordering, only the number of rows written.
- Application code never hardcodes a value list or count.
