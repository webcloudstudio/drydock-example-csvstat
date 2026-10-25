# FEATURE: Median

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Report the exact median of the non-null values of integer and number columns. |
| Depends On  | FEATURE-TypesNumeric.md, FEATURE-MinMax.md |
| Provides    | profile field median |
| Consumes    | column type decision, null rule, trimmed values |

## Purpose

The `median` field of a profile row gives the middle of a numeric column's distribution. It is
computed from exact values so ordering and the even-count average are never subject to floating
point error.

## Trigger

Every profile run computes `median` for every column whose inferred type is `integer` or `number`.

## Behavior

- Only non-null values participate (null: empty, `NA`, `N/A`, `NULL` ignoring case, after trimming).
- Values are parsed into `big.Rat` and sorted ascending by exact value (`Rat.Cmp`).
- With an odd count the median is the middle element.
- With an even count the median is the mean of the two middle elements, as an exact `big.Rat`.
- Both integer and number columns write the median through the decimal formatter, with exactly two
  decimals (see `FEATURE-DecimalFormat.md`).
- A column of type `text`, `date`, `boolean`, or `empty` leaves `median` blank.
- Input order does not matter; the column is sorted on a copy so no other statistic is disturbed.

## Reads

- The trimmed, null-classified values of the column and its inferred type.

## Writes

- The `median` cell of the profile row on standard output. No files, no persistence.

## Programmatic Acceptance

=== AC median-odd-and-even-integers ===
Intent: The median is the middle value of an odd count and the mean of the two middle values of an even count, regardless of input order.
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
        return Fraction(s[n // 2])
    return Fraction(s[n // 2 - 1] + s[n // 2], 2)


def run(text, *opts):
    with tempfile.TemporaryDirectory() as d:
        p = os.path.join(d, "data.csv")
        with open(p, "w", newline="") as fh:
            fh.write(text)
        code, out, err = cli_runner(*opts, p)
        assert code == 0
        return code, out, err


for vals in ([9, 1, 5, 7, 3], [8, 2, 6, 4], [42], [3, -10, 7, 7, 0, -4]):
    text = "v\n" + "\n".join(str(v) for v in vals) + "\n"
    code, out, err = run(text)
    print(out)
    print(err)
    assert code == 0
    row = list(csv.reader(io.StringIO(out)))[1]
    assert row[1] == "integer"
    assert row[9] == fmt(median(vals))
=== END AC median-odd-and-even-integers ===

=== AC median-number-exact-rounding ===
Intent: A number median that falls on a tie is rounded half away from zero from the exact value, for positive and negative values.
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


cases = [["0.01", "0.02"], ["-0.01", "-0.02"], ["1.5", "2.25", "9.125"], ["2.675", "2.675"], ["100", "0.5"]]
for vals in cases:
    text = "v\n" + "\n".join(vals) + "\n"
    code, out, err = run(text)
    print(out)
    print(err)
    assert code == 0
    row = list(csv.reader(io.StringIO(out)))[1]
    assert row[1] == "number"
    assert row[9] == fmt(median([Fraction(v) for v in vals]))
=== END AC median-number-exact-rounding ===

=== AC median-nulls-skipped ===
Intent: Null values are excluded when counting and ordering values for the median.
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


cells = ["50", "NA", "", "10", "N/A", "30", "NULL", "20"]
text = "k,v\n" + "\n".join("a," + c for c in cells) + "\n"
code, out, err = run(text)
print(out)
print(err)
assert code == 0
row = list(csv.reader(io.StringIO(out)))[2]
nonnull = sorted([50, 10, 30, 20])
expected = Fraction(nonnull[1] + nonnull[2], 2)
assert row[0] == "v"
assert row[9] == fmt(expected)
=== END AC median-nulls-skipped ===

=== AC median-blank-for-other-types ===
Intent: Text, date, boolean, and empty columns leave median blank.
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


text = "t,d,b,e,n\nabc,2025-01-31,true,NA,1\nxyz,2024-12-01,False,,2\n"
code, out, err = run(text)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
by_name = {r[0]: r for r in rows[1:]}
for name in ("t", "d", "b", "e"):
    assert by_name[name][9] == ""
assert by_name["n"][9] != ""
=== END AC median-blank-for-other-types ===

=== AC median-fixture-sensors-orders ===
Intent: The supplied sensors and orders files report medians computed independently from their own values.
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


checks = [
    ("sources/sensors.csv", "temp_c"),
    ("sources/sensors.csv", "humidity"),
    ("sources/sensors.csv", "reading_id"),
    ("sources/orders.csv", "qty"),
    ("sources/orders.csv", "unit_price"),
    ("sources/employees.csv", "salary"),
]
for path, name in checks:
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0
    rows = {r[0]: r for r in csv.reader(io.StringIO(out))}
    assert rows[name][9] == fmt(median(values(path, name)))
=== END AC median-fixture-sensors-orders ===

## User Acceptance

- None.

## Guardrails

- No `float32` or `float64` on the data path of the median.
- The median of a column never mutates the order of values used by other statistics.
- Median is blank for every type other than `integer` and `number`.
- Application code never reads `sources/run.sh` or any supplied data file and never embeds an answer.
