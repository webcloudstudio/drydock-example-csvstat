# FEATURE: Frequency Top

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Limit frequency rows with --top and reject a non-positive or non-integer value with exit 2. |
| Depends On  | FEATURE-Frequency.md |
| Provides    | --top behavior |
| Consumes    | --freq behavior |

## Purpose

`--top <n>` bounds the number of rows `--freq` writes.

## Behavior

- `--top n` writes at most `n` frequency rows; the default is 5.
- `n` must be a positive integer. Zero, a negative number, a non-numeric value, or a decimal exits
  2 with nothing on standard output.
- `--top` without `--freq` exits 2 with nothing on standard output.
- When `n` exceeds the number of distinct non-null values, all of them are listed.
- The limit applies after ordering, so the rows kept are the most frequent.

## Programmatic Acceptance

=== AC top-limits-rows ===
Intent: --top n lists the first n rows of the full ordering, and a larger n lists every distinct value.
Requires: executable=go; scope=test

import csv
import io
from drydock_go import cli_runner

data = "sources/employees.csv"
column = "department"
with open(data, newline="") as handle:
    records = list(csv.reader(handle))
index = [name.strip() for name in records[0]].index(column)
distinct = {row[index].strip() for row in records[1:]}

code, out, err = cli_runner("--freq", column, "--top", str(len(distinct) + 10), data)
assert code == 0, (code, out, err)
full = list(csv.reader(io.StringIO(out)))[1:]
assert len(full) == len(distinct), (full, distinct)

limit = 2
code, out, err = cli_runner("--freq", column, "--top", str(limit), data)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["value", "count"], rows[0]
assert rows[1:] == full[:limit], (rows[1:], full[:limit])
=== END AC top-limits-rows ===

=== AC top-with-column-order ===
Intent: --top may precede --freq and still limits the listing.
Requires: executable=go; scope=test

import csv
import io
from drydock_go import cli_runner

data = "sources/sensors.csv"
limit = 3
code, out, err = cli_runner("--top", str(limit), "--freq", "sensor", data)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))[1:]
assert len(rows) == limit, rows
=== END AC top-with-column-order ===

=== AC top-rejects-non-positive ===
Intent: A --top value that is not a positive integer exits 2 and writes nothing to standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

data = "sources/sensors.csv"
for bad in ("0", "-1", "abc", "2.5", ""):
    code, out, err = cli_runner("--freq", "sensor", "--top", bad, data)
    assert code == 2, (bad, code, out, err)
    assert out == "", (bad, out)
=== END AC top-rejects-non-positive ===

=== AC top-without-freq ===
Intent: --top without --freq exits 2 and writes nothing to standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("--top", "3", "sources/sensors.csv")
assert code == 2, (code, out, err)
assert out == "", out
=== END AC top-without-freq ===

=== AC top-missing-value ===
Intent: --top given as the last option without a value exits 2.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("--freq", "sensor", "--top")
assert code == 2, (code, out, err)
assert out == "", out
=== END AC top-missing-value ===

## User Acceptance

- None.

## Guardrails

- Usage errors never write to standard output.
- The default of 5 is applied only when `--top` is absent.
