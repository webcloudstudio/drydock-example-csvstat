# FEATURE: Frequency Limit

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Limits frequency output with --top, defaulting to 5 and accepting a positive integer only. |
| Depends On  | FEATURE-Frequency.md, FEATURE-UsageErrors.md, ARCHITECTURE.md |
| Provides    | csvstat --freq <name> --top <n> <file> |
| Consumes    | value,count output |

## Purpose

`--top <n>` bounds how many rows `--freq` writes.

## Behavior

- Default limit is 5 when `--top` is absent.
- `--top` accepts a positive integer only. Zero, negative numbers, non-numeric text, and an empty
  value are usage errors (exit 2, nothing on standard output).
- When fewer distinct non-null values exist than the limit, all of them are written.
- `--top` without `--freq` is a usage error (exit 2, nothing on standard output).
- A very large positive integer that fits no machine int is accepted as unbounded rather than
  rejected.

## Programmatic Acceptance

=== AC top-default-is-five ===
Intent: Without --top, a column with many distinct values lists exactly five rows.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import csv
import io
import os
from collections import Counter

from drydock_go import cli_runner

path = os.path.abspath("sources/orders.csv")
col = "order_date"
with open(path, newline="") as fh:
    rows = list(csv.reader(fh))
header = [h.strip() for h in rows[0]]
idx = header.index(col)
vals = [r[idx].strip() for r in rows[1:]]
vals = [v for v in vals if v != "" and v.upper() not in ("NA", "N/A", "NULL")]
ordered = sorted(Counter(vals).items(), key=lambda kv: (-kv[1], kv[0].encode()))
default_limit = 5
assert len(ordered) > default_limit
want = [[v, str(c)] for v, c in ordered[:default_limit]]

code, out, err = cli_runner("--freq", col, path)
print(out)
print(err)
assert code == 0, (code, out, err)
got = list(csv.reader(io.StringIO(out)))
assert got[1:] == want, (got[1:], want)
=== END AC top-default-is-five ===

=== AC top-limits-rows ===
Intent: An explicit --top n writes at most n rows, matching the ordered counts derived from the file.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import csv
import io
import os
from collections import Counter

from drydock_go import cli_runner

path = os.path.abspath("sources/employees.csv")
col = "manager_id"
top = 3
with open(path, newline="") as fh:
    rows = list(csv.reader(fh))
header = [h.strip() for h in rows[0]]
idx = header.index(col)
vals = [r[idx].strip() for r in rows[1:]]
vals = [v for v in vals if v != "" and v.upper() not in ("NA", "N/A", "NULL")]
ordered = sorted(Counter(vals).items(), key=lambda kv: (-kv[1], kv[0].encode()))
want = [[v, str(c)] for v, c in ordered[:top]]

code, out, err = cli_runner("--freq", col, "--top", str(top), path)
print(out)
print(err)
assert code == 0, (code, out, err)
got = list(csv.reader(io.StringIO(out)))
assert got[1:] == want, (got[1:], want)
assert len(got) - 1 <= top
=== END AC top-limits-rows ===

=== AC top-larger-than-distinct ===
Intent: A limit above the distinct count lists every distinct non-null value.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import csv
import io
import os
from collections import Counter

from drydock_go import cli_runner

path = os.path.abspath("sources/employees.csv")
col = "department"
top = 100
with open(path, newline="") as fh:
    rows = list(csv.reader(fh))
header = [h.strip() for h in rows[0]]
idx = header.index(col)
vals = [r[idx].strip() for r in rows[1:]]
vals = [v for v in vals if v != "" and v.upper() not in ("NA", "N/A", "NULL")]
counts = Counter(vals)
ordered = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode()))
want = [[v, str(c)] for v, c in ordered]

code, out, err = cli_runner("--freq", col, "--top", str(top), path)
print(out)
print(err)
assert code == 0, (code, out, err)
got = list(csv.reader(io.StringIO(out)))
assert got[1:] == want, (got[1:], want)
assert len(got) - 1 == len(counts)
=== END AC top-larger-than-distinct ===

=== AC top-rejects-non-positive ===
Intent: A --top that is not a positive integer exits 2 with nothing on standard output.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

path = os.path.abspath("sources/sensors.csv")
for bad in ("0", "-1", "abc", "2.5", ""):
    code, out, err = cli_runner("--freq", "sensor", "--top", bad, path)
    print(bad, code, err)
    assert code == 2, (bad, code, out, err)
    assert out == "", (bad, out)
=== END AC top-rejects-non-positive ===

=== AC top-requires-freq ===
Intent: --top without --freq is a usage error with exit 2 and empty standard output.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

path = os.path.abspath("sources/sensors.csv")
code, out, err = cli_runner("--top", "3", path)
print(err)
assert code == 2, (code, out, err)
assert out == "", out

code, out, err = cli_runner("--column", "sensor", "--top", "3", path)
print(err)
assert code == 2, (code, out, err)
assert out == "", out
=== END AC top-requires-freq ===

## User Acceptance

- None.

## Guardrails

- `--top` never accepts zero or a negative number.
- A rejected `--top` writes nothing to standard output.
- The default of 5 is applied only when `--top` is absent.
