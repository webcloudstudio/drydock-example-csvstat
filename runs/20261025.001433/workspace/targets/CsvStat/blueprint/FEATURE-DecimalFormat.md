# FEATURE: Decimal Format

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Write numeric statistics with exactly two decimals, rounded half away from zero once from the exact value. |
| Depends On  | FEATURE-SumMean.md, FEATURE-Median.md |
| Provides    | decimal formatter for min, max, sum, mean, and median |
| Consumes    | exact big.Rat values from sum, mean, median, min, max |

## Purpose

Numeric statistics are written in one canonical shape. Rounding is performed exactly once, at
output, from the exact rational value, so that results are reproducible and independent of
binary floating point.

## Trigger

Every profile row of an `integer` or `number` column writes its numeric cells through this
formatter.

## Behavior

- Package: `internal/decimal` exposes a function taking a `*big.Rat` and returning a string.
- For a `number` column: `min`, `max`, `sum`, `mean`, and `median` are two-decimal strings.
- For an `integer` column: `mean` and `median` are two-decimal strings; `min`, `max`, and `sum`
  are plain integers with no decimal point.
- Two-decimal form: scale the absolute value by 100, round half up on the absolute value
  (ties move away from zero) using `big.Int` arithmetic, then insert the decimal point. There is
  always at least one digit before the point and exactly two after it.
- A leading `-` appears only when the rounded result is non-zero and the value is negative. A
  negative value that rounds to zero is written without a sign.
- No thousands separator, no exponent, no leading `+`.
- Rounding is applied once to the exact value; an already rounded value is never rounded again.
- Non-numeric column types are unaffected and keep their own rules.

## Reads

- Exact `big.Rat` values produced by the sum, mean, median, and min/max computations.

## Writes

- The numeric cells of the profile row on standard output.

## Programmatic Acceptance

=== AC decimalformat-number-columns ===
Intent: Every number min, max, sum, mean, and median is written with exactly two decimals rounded half away from zero from the exact value.
Suite: scoped
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from decimal import Decimal, ROUND_HALF_UP, getcontext
from fractions import Fraction

from drydock_go import cli_runner


def fmt(fr):
    getcontext().prec = 200
    d = Decimal(fr.numerator) / Decimal(fr.denominator)
    return str(d.quantize(Decimal("0.01"), rounding=ROUND_HALF_UP))


def median(vals):
    s = sorted(vals)
    n = len(s)
    if n % 2:
        return s[n // 2]
    return (s[n // 2 - 1] + s[n // 2]) / 2


def run(text, *opts):
    with tempfile.TemporaryDirectory() as d:
        p = os.path.join(d, "data.csv")
        with open(p, "w", newline="") as fh:
            fh.write(text)
        code, out, err = cli_runner(*opts, p)
        assert code == 0
        return code, out, err


cases = [["1.255", "3", "4.5"], ["2.675"], ["-2.675", "-1.004", "10.125"], ["1234567.891", "0.5"]]
for vals in cases:
    text = "v\n" + "\n".join(vals) + "\n"
    code, out, err = run(text)
    print(out)
    print(err)
    assert code == 0
    row = list(csv.reader(io.StringIO(out)))[1]
    frs = [Fraction(v) for v in vals]
    assert row[1] == "number"
    assert row[5] == fmt(min(frs))
    assert row[6] == fmt(max(frs))
    assert row[7] == fmt(sum(frs))
    assert row[8] == fmt(sum(frs) / len(frs))
    assert row[9] == fmt(median(frs))
=== END AC decimalformat-number-columns ===

=== AC decimalformat-integer-columns ===
Intent: Integer min, max, and sum are plain integers while integer mean and median carry exactly two decimals.
Suite: scoped
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from decimal import Decimal, ROUND_HALF_UP, getcontext
from fractions import Fraction

from drydock_go import cli_runner


def fmt(fr):
    getcontext().prec = 200
    d = Decimal(fr.numerator) / Decimal(fr.denominator)
    return str(d.quantize(Decimal("0.01"), rounding=ROUND_HALF_UP))


def run(text, *opts):
    with tempfile.TemporaryDirectory() as d:
        p = os.path.join(d, "data.csv")
        with open(p, "w", newline="") as fh:
            fh.write(text)
        code, out, err = cli_runner(*opts, p)
        assert code == 0
        return code, out, err


vals = [-5, 12, 7, 1000, 2]
text = "v\n" + "\n".join(str(v) for v in vals) + "\n"
code, out, err = run(text)
print(out)
print(err)
assert code == 0
row = list(csv.reader(io.StringIO(out)))[1]
s = sorted(vals)
assert row[1] == "integer"
assert row[5] == str(min(vals))
assert row[6] == str(max(vals))
assert row[7] == str(sum(vals))
assert row[8] == fmt(Fraction(sum(vals), len(vals)))
assert row[9] == fmt(Fraction(s[len(s) // 2]))
for cell in (row[5], row[6], row[7]):
    assert "." not in cell
for cell in (row[8], row[9]):
    assert len(cell.split(".")[1]) == 2
=== END AC decimalformat-integer-columns ===

=== AC decimalformat-single-rounding ===
Intent: Rounding happens once from the exact value, so a value that double rounding would push up stays down.
Suite: scoped
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from decimal import Decimal, ROUND_HALF_UP, getcontext
from fractions import Fraction

from drydock_go import cli_runner


def fmt(fr):
    getcontext().prec = 200
    d = Decimal(fr.numerator) / Decimal(fr.denominator)
    return str(d.quantize(Decimal("0.01"), rounding=ROUND_HALF_UP))


def run(text, *opts):
    with tempfile.TemporaryDirectory() as d:
        p = os.path.join(d, "data.csv")
        with open(p, "w", newline="") as fh:
            fh.write(text)
        code, out, err = cli_runner(*opts, p)
        assert code == 0
        return code, out, err


vals = ["1.0049", "1.0049"]
text = "v\n" + "\n".join(vals) + "\n"
code, out, err = run(text)
print(out)
print(err)
assert code == 0
row = list(csv.reader(io.StringIO(out)))[1]
exact = Fraction(vals[0])
assert row[5] == fmt(exact)
assert row[8] == fmt(exact)
assert row[9] == fmt(exact)
assert Decimal(row[8]) < Decimal(vals[0]) + Decimal("0.005")
=== END AC decimalformat-single-rounding ===

=== AC decimalformat-no-negative-zero ===
Intent: A negative value that rounds to zero is written as an unsigned zero with two decimals.
Suite: scoped
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def run(text, *opts):
    with tempfile.TemporaryDirectory() as d:
        p = os.path.join(d, "data.csv")
        with open(p, "w", newline="") as fh:
            fh.write(text)
        code, out, err = cli_runner(*opts, p)
        assert code == 0
        return code, out, err


vals = ["-0.001", "-0.002"]
text = "v\n" + "\n".join(vals) + "\n"
code, out, err = run(text)
print(out)
print(err)
assert code == 0
row = list(csv.reader(io.StringIO(out)))[1]
assert row[1] == "number"
for cell in row[5:10]:
    assert not cell.startswith("-")
    assert len(cell.split(".")[1]) == 2
    assert int(cell.split(".")[0]) == 0
    assert int(cell.split(".")[1]) == 0
=== END AC decimalformat-no-negative-zero ===

=== AC decimalformat-fixture-files ===
Intent: Number and integer columns of the supplied files are written in the canonical shape from values computed independently of the application.
Suite: scoped
Requires: executable=go; scope=test

import csv
import io
from decimal import Decimal, ROUND_HALF_UP, getcontext
from fractions import Fraction

from drydock_go import cli_runner


def fmt(fr):
    getcontext().prec = 200
    d = Decimal(fr.numerator) / Decimal(fr.denominator)
    return str(d.quantize(Decimal("0.01"), rounding=ROUND_HALF_UP))


def median(vals):
    s = sorted(vals)
    n = len(s)
    if n % 2:
        return s[n // 2]
    return (s[n // 2 - 1] + s[n // 2]) / 2


null = ("", "na", "n/a", "null")


def values(path, name):
    with open(path, newline="") as fh:
        recs = list(csv.reader(fh))
    i = [h.strip() for h in recs[0]].index(name)
    cells = [r[i].strip() for r in recs[1:]]
    return [Fraction(c) for c in cells if c.lower() not in null]


for path, name in (
    ("sources/employees.csv", "salary"),
    ("sources/sensors.csv", "temp_c"),
    ("sources/orders.csv", "unit_price"),
):
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0
    rows = {r[0]: r for r in csv.reader(io.StringIO(out))}
    frs = values(path, name)
    row = rows[name]
    assert row[5] == fmt(min(frs))
    assert row[6] == fmt(max(frs))
    assert row[7] == fmt(sum(frs))
    assert row[8] == fmt(sum(frs) / len(frs))
    assert row[9] == fmt(median(frs))

code, out, err = cli_runner("sources/sensors.csv")
assert code == 0
rows = {r[0]: r for r in csv.reader(io.StringIO(out))}
hum = values("sources/sensors.csv", "humidity")
row = rows["humidity"]
assert row[5] == str(int(min(hum)))
assert row[6] == str(int(max(hum)))
assert row[7] == str(int(sum(hum)))
assert row[8] == fmt(sum(hum) / len(hum))
assert row[9] == fmt(median(hum))
=== END AC decimalformat-fixture-files ===

## User Acceptance

- None.

## Guardrails

- Rounding is performed exactly once, from the exact `big.Rat`, never from a rounded intermediate.
- No `float32` or `float64`, and no `fmt` verb that formats a float, produces a statistic.
- Output carries no thousands separator, exponent, or `+` sign.
- Application code never reads `sources/run.sh` or any supplied data file and never embeds an answer.
