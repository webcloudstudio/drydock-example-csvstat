# FEATURE: Profile Extremes

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Compute the min and max of every column according to its inferred type. |
| Depends On  | FEATURE-Decimal-Rounding.md, FEATURE-Profile-Counts.md |
| Provides    | profile min, max |
| Consumes    | infer types, internal/decimal |

## Purpose

Fill the `min` and `max` fields of each profile row. Null values are ignored throughout.

## Behavior

- `integer`: smallest and largest compared numerically, written as plain integers.
- `number`: smallest and largest compared exactly (`big.Rat`), written with the two-decimal formatter (`10.125` is written `10.13`).
- `date`: earliest and latest.
- `boolean`: `min` is `false` if any value is false (any letter case), else `true`; `max` is `true` if any value is true, else `false`. Output is lower case.
- `text`: smallest and largest by byte value of the trimmed string (`B` sorts before `a`; `10` sorts before `9`).
- `empty`: both fields are empty.
- A value that needs quoting (a comma, quote, or newline) is quoted by the shared CSV writer.

## Programmatic Acceptance

=== AC prf003-integer-number ===
Intent: Integer extremes compare numerically and number extremes compare exactly and are written with two decimals.
Suite: scoped

import csv
import io
import os
import tempfile
from fractions import Fraction
from drydock_go import cli_runner

def run_profile(columns, *opts):
    names = list(columns)
    n = len(columns[names[0]])
    buf = io.StringIO()
    w = csv.writer(buf, lineterminator="\n")
    w.writerow(names)
    for i in range(n):
        w.writerow([columns[c][i] for c in names])
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write(buf.getvalue())
        code, out, err = cli_runner(*opts, path)
    print(out)
    print(err)
    assert code == 0, code
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}

def fmt2(fr):
    scaled = fr * 100
    n, d = scaled.numerator, scaled.denominator
    r = (abs(n) * 2 + d) // (2 * d)
    sign = "-" if (n < 0 and r != 0) else ""
    return sign + str(r // 100) + "." + str(r % 100).zfill(2)

def real(values):
    return [Fraction(v) for v in values if v.lower() not in ("", "na", "n/a", "null")]

columns = {
    "ints": ["10", "9", "-3", "NA", "100"],
    "nums": ["1.5", "-2.25", "10.125", "NULL", "0.0"],
    "trap": ["9.5", "10.25", "", "N/A", "9.75"],
}
rows = run_profile(columns)
iv = real(columns["ints"])
assert rows["ints"]["type"] == "integer", rows["ints"]
assert rows["ints"]["min"] == str(int(min(iv))), rows["ints"]
assert rows["ints"]["max"] == str(int(max(iv))), rows["ints"]
for name in ("nums", "trap"):
    vals = real(columns[name])
    assert rows[name]["type"] == "number", (name, rows[name])
    assert rows[name]["min"] == fmt2(min(vals)), (name, rows[name])
    assert rows[name]["max"] == fmt2(max(vals)), (name, rows[name])
=== END AC prf003-integer-number ===

=== AC prf003-date ===
Intent: Date columns report the earliest and latest date, ignoring nulls.
Suite: scoped

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

def run_profile(columns, *opts):
    names = list(columns)
    n = len(columns[names[0]])
    buf = io.StringIO()
    w = csv.writer(buf, lineterminator="\n")
    w.writerow(names)
    for i in range(n):
        w.writerow([columns[c][i] for c in names])
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write(buf.getvalue())
        code, out, err = cli_runner(*opts, path)
    print(out)
    print(err)
    assert code == 0, code
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}

raw = ["2025-03-01", "2024-12-31", "2025-01-15", "", "NA"]
vals = [v for v in raw if v not in ("", "NA")]
rows = run_profile({"d": raw})
assert rows["d"]["type"] == "date", rows["d"]
assert rows["d"]["min"] == min(vals), rows["d"]
assert rows["d"]["max"] == max(vals), rows["d"]
=== END AC prf003-date ===

=== AC prf003-boolean ===
Intent: Boolean min is false when any value is false and max is true when any value is true, in lower case.
Suite: scoped

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

def run_profile(columns, *opts):
    names = list(columns)
    n = len(columns[names[0]])
    buf = io.StringIO()
    w = csv.writer(buf, lineterminator="\n")
    w.writerow(names)
    for i in range(n):
        w.writerow([columns[c][i] for c in names])
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write(buf.getvalue())
        code, out, err = cli_runner(*opts, path)
    print(out)
    print(err)
    assert code == 0, code
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}

columns = {
    "all_true": ["true", "TRUE", "True"],
    "mixed": ["true", "False", "NA"],
    "all_false": ["FALSE", "false", ""],
}
rows = run_profile(columns)
for name, raw in columns.items():
    vals = [v.lower() for v in raw if v.lower() not in ("", "na", "n/a", "null")]
    expected_min = "false" if "false" in vals else "true"
    expected_max = "true" if "true" in vals else "false"
    assert rows[name]["type"] == "boolean", (name, rows[name])
    assert rows[name]["min"] == expected_min, (name, rows[name])
    assert rows[name]["max"] == expected_max, (name, rows[name])
=== END AC prf003-boolean ===

=== AC prf003-text-bytes ===
Intent: Text extremes are compared by byte value of the trimmed strings and ignore nulls.
Suite: scoped

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

def run_profile(columns, *opts):
    names = list(columns)
    n = len(columns[names[0]])
    buf = io.StringIO()
    w = csv.writer(buf, lineterminator="\n")
    w.writerow(names)
    for i in range(n):
        w.writerow([columns[c][i] for c in names])
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write(buf.getvalue())
        code, out, err = cli_runner(*opts, path)
    print(out)
    print(err)
    assert code == 0, code
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}

columns = {
    "cased": ["b", "B", "a", "Z", "NA", " pad "],
    "digits": ["10", "9", "x", "", "NULL", "ab"],
    "commas": ["a,b", "a", "ab", "NA", "", "b"],
}
rows = run_profile(columns)
for name, raw in columns.items():
    vals = [v.strip() for v in raw if v.strip().lower() not in ("", "na", "n/a", "null")]
    expected_min = min(vals, key=lambda s: s.encode())
    expected_max = max(vals, key=lambda s: s.encode())
    assert rows[name]["type"] == "text", (name, rows[name])
    assert rows[name]["min"] == expected_min, (name, rows[name])
    assert rows[name]["max"] == expected_max, (name, rows[name])
=== END AC prf003-text-bytes ===

=== AC prf003-empty-and-header-only ===
Intent: Empty columns, including every column of a header-only file, have empty min and max.
Suite: scoped

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
assert code == 0, code
rows = list(csv.DictReader(io.StringIO(out)))
assert len(rows) == len(names)
for r in rows:
    assert r["type"] == "empty", r
    assert r["min"] == "", r
    assert r["max"] == "", r

buf = io.StringIO()
w = csv.writer(buf, lineterminator="\n")
w.writerow(["only_null"])
for v in ["", "NA", "NULL"]:
    w.writerow([v])
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in2.csv")
    with open(path, "w", newline="") as f:
        f.write(buf.getvalue())
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, code
row = list(csv.DictReader(io.StringIO(out)))[0]
assert row["type"] == "empty", row
assert row["min"] == "", row
assert row["max"] == "", row
=== END AC prf003-empty-and-header-only ===

## User Acceptance

- None.

## Guardrails

- Number extremes are compared exactly, never through `float64`, and are written only through `internal/decimal`.
- Text comparison is by byte value; do not use locale or case-folding collation.
- Do not hardcode any extreme; every value comes from the file examined.
