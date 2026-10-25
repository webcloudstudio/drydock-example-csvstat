# FEATURE: Types Numeric

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infer the integer and number column types from the non-null values of a column. |
| Depends On  | FEATURE-ProfileCounts.md |
| Provides    | profile field type values integer and number |
| Consumes    | null rule from FEATURE-ProfileCounts.md |

## Purpose

A column's type is the first type in the order integer, number, date, boolean, text that every non-null value satisfies. This specification owns the first two members of that order.

## Behavior

- `integer`: every non-null value matches `-?[0-9]+` (the whole value, after trimming).
- `number`: every non-null value matches `-?[0-9]+(\.[0-9]+)?` (the whole value, after trimming) and the column is not `integer`.
- Matching is anchored at both ends. `+5`, `.5`, `5.`, `1e3`, `1,000`, and `0x10` match neither pattern.
- A single non-matching non-null value disqualifies the type for the column.
- Null values (see FEATURE-ProfileCounts.md) never disqualify a type.
- Leading zeros are permitted (`007` is an integer). Arbitrarily long digit strings are integers; no length limit applies.
- Matchers live in `internal/profile` and are hand-coded or use `regexp` anchored with `^` and `$`.

## Programmatic Acceptance

=== AC types-numeric-integer ===
Intent: A column whose non-null values all match the integer pattern is typed integer, including negatives, leading zeros, and nulls.

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


values = ["1", "-2", "007", " 40 ", "NA", "", "123456789012345678901234567890"]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
assert rows["v"]["type"] == "integer"
assert rows["id"]["type"] == "integer"
=== END AC types-numeric-integer ===

=== AC types-numeric-number ===
Intent: A column with at least one decimal value and otherwise integer or decimal values is typed number.

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


values = ["1", "-2.50", "3.14159", "0.0", "N/A", "10"]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
assert rows["v"]["type"] == "number"
=== END AC types-numeric-number ===

=== AC types-numeric-rejects-lookalikes ===
Intent: Each value that fails the anchored numeric patterns keeps its column from being integer or number.

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


bad_values = ["+5", ".5", "5.", "1e3", "1.2.3", "0x10", "--1", "1-", "12abc"]
for bad in bad_values:
    values = ["1", "2.5", bad]
    lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
    rows = profile(lines)
    assert rows["v"]["type"] not in ("integer", "number"), (bad, rows["v"]["type"])
=== END AC types-numeric-rejects-lookalikes ===

=== AC types-numeric-integer-preferred ===
Intent: An all-integer column is integer, not number, and adding one decimal changes only that column to number.

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


ints = ["1", "2", "3"]
mixed = ["1", "2", "3.5"]
lines = ["a,b"] + [f"{x},{y}" for x, y in zip(ints, mixed)]
rows = profile(lines)
assert rows["a"]["type"] == "integer"
assert rows["b"]["type"] == "number"
=== END AC types-numeric-integer-preferred ===

## User Acceptance

- None.

## Guardrails

- Type decisions are computed from the values of the file examined; no type is hardcoded for a column name or file.
- Null values never disqualify a type.
- No third-party modules; standard library only.
