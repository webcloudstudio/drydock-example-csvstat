# FEATURE: Sum And Mean

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Report the exact sum and mean of the non-null values of integer and number columns. |
| Depends On  | FEATURE-TypesNumeric.md, FEATURE-MinMax.md |
| Provides    | profile field sum, profile field mean |
| Consumes    | column type decision, null rule, trimmed values |

## Purpose

The `sum` and `mean` fields of each profile row summarize numeric columns. They are computed with
exact arithmetic so that no digit string overflows and no binary floating point error enters the
result. Rounding happens once, when the value is written (see `FEATURE-DecimalFormat.md`).

## Trigger

Every profile run (`csvstat <file>` and `csvstat --column <name> <file>`) computes `sum` and `mean`
for every column whose inferred type is `integer` or `number`.

## Behavior

- Package: `internal/profile` accumulates, `internal/decimal` holds exact values.
- Only non-null values participate. Null values (empty, `NA`, `N/A`, `NULL` ignoring case, after
  trimming) are skipped.
- Each value is parsed exactly: an integer into `big.Int`, a number into `big.Rat` through
  `Rat.SetString`. `float32` and `float64` are never used on the data path.
- `sum` is the exact total of the non-null values.
- `mean` is the exact total divided by the count of non-null values, held as a `big.Rat`.
- An `integer` column writes `sum` as a plain integer (an arbitrary-precision digit string). Its
  `mean` is written by the decimal formatter.
- A `number` column writes `sum` and `mean` through the decimal formatter.
- A column of type `text`, `date`, `boolean`, or `empty` leaves `sum` and `mean` blank.
- Columns are independent: a numeric column's sum never depends on another column.

## Reads

- The trimmed, null-classified values of the column and its inferred type.

## Writes

- The `sum` and `mean` cells of the profile row on standard output. No files, no persistence.

## Programmatic Acceptance

=== AC summean-integer-exact ===
Intent: An integer column reports the exact integer sum and the two-decimal mean of its values.
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


vals = [4, -7, 12, 100, 3]
text = "v\n" + "\n".join(str(v) for v in vals) + "\n"
code, out, err = run(text)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
row = rows[1]
total = sum(vals)
assert row[1] == "integer"
assert row[7] == str(total)
assert row[8] == fmt(Fraction(total, len(vals)))
=== END AC summean-integer-exact ===

=== AC summean-integer-bigint ===
Intent: Integer sums beyond 64 bits are computed with arbitrary precision and never overflow.
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


vals = [10**30 + 7, 10**30 + 9, -(10**29), 99999999999999999999]
text = "v\n" + "\n".join(str(v) for v in vals) + "\n"
code, out, err = run(text)
print(out)
print(err)
assert code == 0
row = list(csv.reader(io.StringIO(out)))[1]
total = sum(vals)
assert row[7] == str(total)
assert row[8] == fmt(Fraction(total, len(vals)))
=== END AC summean-integer-bigint ===

=== AC summean-number-exact-decimal ===
Intent: Number sums and means come from exact decimal values and are rounded half away from zero once.
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


cases = [["2.675"], ["0.1", "0.2", "0.7"], ["-2.675", "1.005", "3"], ["1.0049", "1.0049"]]
for vals in cases:
    text = "v\n" + "\n".join(vals) + "\n"
    code, out, err = run(text)
    print(out)
    print(err)
    assert code == 0
    row = list(csv.reader(io.StringIO(out)))[1]
    frs = [Fraction(v) for v in vals]
    total = sum(frs)
    assert row[1] == "number"
    assert row[7] == fmt(total)
    assert row[8] == fmt(total / len(frs))
=== END AC summean-number-exact-decimal ===

=== AC summean-nulls-skipped ===
Intent: Null values are excluded from both the sum and the divisor of the mean.
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


cells = ["10", "", "NA", " n/a ", "20", "Null", "35"]
text = "k,v\n" + "\n".join("a," + c for c in cells) + "\n"
code, out, err = run(text)
print(out)
print(err)
assert code == 0
row = list(csv.reader(io.StringIO(out)))[2]
nonnull = [10, 20, 35]
assert row[0] == "v"
assert row[1] == "integer"
assert row[7] == str(sum(nonnull))
assert row[8] == fmt(Fraction(sum(nonnull), len(nonnull)))
=== END AC summean-nulls-skipped ===

=== AC summean-blank-for-other-types ===
Intent: Text, date, boolean, and empty columns leave sum and mean blank.
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
    assert by_name[name][7] == ""
    assert by_name[name][8] == ""
assert by_name["n"][7] != ""
assert by_name["n"][8] != ""
=== END AC summean-blank-for-other-types ===

=== AC summean-fixture-employees ===
Intent: The supplied employees file reports the exact sum and mean computed independently from its own values.
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


path = "sources/employees.csv"
with open(path, newline="") as fh:
    recs = list(csv.reader(fh))
heads = [h.strip() for h in recs[0]]
null = ("", "na", "n/a", "null")


def values(name):
    i = heads.index(name)
    cells = [r[i].strip() for r in recs[1:]]
    return [c for c in cells if c.lower() not in null]


code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0
rows = {r[0]: r for r in csv.reader(io.StringIO(out))}

ints = [int(v) for v in values("id")]
assert rows["id"][7] == str(sum(ints))
assert rows["id"][8] == fmt(Fraction(sum(ints), len(ints)))

mgr = [int(v) for v in values("manager_id")]
assert rows["manager_id"][7] == str(sum(mgr))
assert rows["manager_id"][8] == fmt(Fraction(sum(mgr), len(mgr)))

sal = [Fraction(v) for v in values("salary")]
assert rows["salary"][7] == fmt(sum(sal))
assert rows["salary"][8] == fmt(sum(sal) / len(sal))
=== END AC summean-fixture-employees ===

## User Acceptance

- None.

## Guardrails

- No `float32` or `float64` on the data path of sum or mean.
- Rounding is not applied before output; the exact `big.Rat` value is carried until the formatter.
- Sum and mean are blank for every type other than `integer` and `number`.
- Application code never reads `sources/run.sh` or any supplied data file and never embeds an answer.
