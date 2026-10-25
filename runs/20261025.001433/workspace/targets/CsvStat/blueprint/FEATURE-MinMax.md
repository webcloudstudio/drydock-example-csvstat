# FEATURE: Min Max

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Report the min and max profile fields for every column type. |
| Depends On  | FEATURE-TypesOther.md |
| Provides    | profile fields min, max |
| Consumes    | column type decision from FEATURE-TypesOther.md |

## Purpose

Each profile row reports the smallest and largest non-null value of its column, interpreted according to the column's type.

## Behavior

- `integer`: the smallest and largest value as numbers, written as plain integers (`-3` before `12`). Comparison is numeric, not lexical, and is exact (`math/big`), so digit strings of any length are safe.
- `number`: the smallest and largest value by exact decimal comparison. Output formatting to exactly two decimals is owned by FEATURE-DecimalFormat.md; this story supplies the exact values.
- `date`: the earliest and latest value (`YYYY-MM-DD` strings order chronologically).
- `boolean`: `min` is `false` if any value is false, else `true`; `max` is `true` if any value is true, else `false`. Output is lower-case regardless of input case.
- `text`: the smallest and largest value by byte value (not locale, not case-folded), after trimming.
- `empty`: both fields are blank.
- Null values are ignored. Blank is the empty string.
- Min and max are computed from the file examined; nothing is hardcoded.

## Programmatic Acceptance

=== AC minmax-integer ===
Intent: Integer min and max are compared numerically and written as plain integers.

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


values = [9, 100, -3, 42, 7]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
lines.append("99,NA")
rows = profile(lines)
assert rows["v"]["type"] == "integer"
assert rows["v"]["min"] == str(min(values))
assert rows["v"]["max"] == str(max(values))
=== END AC minmax-integer ===

=== AC minmax-integer-huge ===
Intent: Integer min and max stay exact for digit strings beyond 64-bit range.

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


values = [10**30 + 1, 10**30, -(10**25), 5]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
assert rows["v"]["min"] == str(min(values))
assert rows["v"]["max"] == str(max(values))
=== END AC minmax-integer-huge ===

=== AC minmax-number ===
Intent: Number min and max are chosen by exact decimal comparison and equal the supplied values.

import csv
import io
import os
import tempfile
from decimal import Decimal

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


values = ["15.27", "231.18", "-12.21", "0.50", "29.81", "100.00"]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
assert rows["v"]["type"] == "number"
assert rows["v"]["min"] == min(values, key=Decimal)
assert rows["v"]["max"] == max(values, key=Decimal)
=== END AC minmax-number ===

=== AC minmax-date ===
Intent: Date min and max are the earliest and latest dates.

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


values = ["2022-05-14", "2018-11-08", "2025-06-02", "2015-08-27", "2024-02-29"]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
assert rows["v"]["type"] == "date"
assert rows["v"]["min"] == min(values)
assert rows["v"]["max"] == max(values)
=== END AC minmax-date ===

=== AC minmax-boolean ===
Intent: Boolean min is false when any value is false and max is true when any value is true; a single-valued column reports that value for both.

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


mixed = ["True", "FALSE", "true"]
only_true = ["True", "TRUE", "true"]
only_false = ["False", "FALSE", "false"]
lines = ["m,t,f"] + [f"{a},{b},{c}" for a, b, c in zip(mixed, only_true, only_false)]
rows = profile(lines)
assert rows["m"]["min"] == "false" and rows["m"]["max"] == "true"
assert rows["t"]["min"] == "true" and rows["t"]["max"] == "true"
assert rows["f"]["min"] == "false" and rows["f"]["max"] == "false"
=== END AC minmax-boolean ===

=== AC minmax-text-bytes ===
Intent: Text min and max compare by byte value, so upper-case sorts before lower-case, and padding is trimmed.

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


values = ["banana", "Zebra", "apple", "  Acme Corp ", "mentor"]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
assert rows["v"]["type"] == "text"
trimmed = [v.strip() for v in values]
assert rows["v"]["min"] == min(trimmed, key=lambda s: s.encode())
assert rows["v"]["max"] == max(trimmed, key=lambda s: s.encode())
=== END AC minmax-text-bytes ===

=== AC minmax-empty-and-nulls ===
Intent: An empty column reports blank min and max, and nulls are ignored when choosing extremes.

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


values = [5, 2, 9]
lines = ["id,v,gap"] + [f"{i},{v},NULL" for i, v in enumerate(values)]
lines.append("77,NA,")
rows = profile(lines)
assert rows["gap"]["type"] == "empty"
assert rows["gap"]["min"] == ""
assert rows["gap"]["max"] == ""
assert rows["v"]["min"] == str(min(values))
assert rows["v"]["max"] == str(max(values))
=== END AC minmax-empty-and-nulls ===

## User Acceptance

- None.

## Guardrails

- Min and max are computed from the file examined; no extreme is hardcoded.
- Integer and number comparison is exact; no binary floating point is used.
- Text comparison is by byte value, never locale or case folding.
- Standard library only.
