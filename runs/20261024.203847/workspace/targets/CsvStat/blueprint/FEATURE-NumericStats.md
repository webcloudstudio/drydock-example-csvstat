# FEATURE: Numeric Stats

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Compute exact sum, mean, and median for integer and number columns. |
| Depends On  | FEATURE-TypeInference.md, ARCHITECTURE.md |
| Provides    | profile.NumericStats |
| Consumes    | typing.Infer |

## Purpose

Supply the `sum`, `mean`, and `median` fields of each profile row for numeric columns.

## Behavior

- Applies to `integer` and `number` columns; every other type leaves all three fields blank.
- Parse non-null values into exact rationals with `math/big` (`big.Rat`). Never use floats. Integer totals cannot overflow.
- `sum`: integer columns write the plain integer total; number columns write the total with exactly two decimal places.
- `mean`: total divided by the non-null count, written with two decimal places for both integer and number columns.
- `median`: the middle of the sorted non-null values, or the mean of the two middle values when the count is even; two decimal places for both types.
- Rounding happens once, from the exact value, half away from zero. A leading `-` marks a negative result; no thousands separator.
- `profile.NumericStats(typ string, values []string) (sum, mean, median string)`.

## Programmatic Acceptance

=== AC numeric-integer-stats ===
Intent: Integer sum is a plain integer while mean and median carry two places, with an even count averaging the middle pair.

import csv
import io
import os
import tempfile
from fractions import Fraction

from drydock_go import cli_runner


def two(fr):
    neg = fr < 0
    q = int((abs(fr) * 100 + Fraction(1, 2)) // 1)
    sign = "-" if neg and q else ""
    return sign + str(q // 100) + "." + str(q % 100).zfill(2)


ints = [1, 2, 3, 10]
text = "n\n" + "\n".join(str(i) for i in ints) + "\n"
total = sum(ints)
mean = Fraction(total, len(ints))
s = sorted(ints)
mid = len(s) // 2
median = Fraction(s[mid - 1] + s[mid], 2)
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
r = list(csv.DictReader(io.StringIO(out)))[0]
assert r["type"] == "integer"
assert r["sum"] == str(total)
assert r["mean"] == two(mean)
assert r["median"] == two(median)
=== END AC numeric-integer-stats ===

=== AC numeric-integer-odd-median ===
Intent: An odd count takes the middle sorted value, independent of input order.

import csv
import io
import os
import tempfile
from fractions import Fraction

from drydock_go import cli_runner


def two(fr):
    neg = fr < 0
    q = int((abs(fr) * 100 + Fraction(1, 2)) // 1)
    sign = "-" if neg and q else ""
    return sign + str(q // 100) + "." + str(q % 100).zfill(2)


ints = [9, 1, 5, 100, 3]
text = "n\n" + "\n".join(str(i) for i in ints) + "\n"
s = sorted(ints)
median = Fraction(s[len(s) // 2])
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
r = list(csv.DictReader(io.StringIO(out)))[0]
assert r["median"] == two(median)
assert r["sum"] == str(sum(ints))
=== END AC numeric-integer-odd-median ===

=== AC numeric-exact-large-sum ===
Intent: Integer totals beyond 64 bits are exact rather than overflowed.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

ints = [9223372036854775807, 9223372036854775807, 9223372036854775807]
text = "n\n" + "\n".join(str(i) for i in ints) + "\n"
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
r = list(csv.DictReader(io.StringIO(out)))[0]
assert r["sum"] == str(sum(ints))
=== END AC numeric-exact-large-sum ===

=== AC numeric-number-stats-rounding ===
Intent: Number sum, mean, and median are two places rounded half away from zero from the exact value, with nulls excluded.

import csv
import io
import os
import tempfile
from fractions import Fraction

from drydock_go import cli_runner


def two(fr):
    neg = fr < 0
    q = int((abs(fr) * 100 + Fraction(1, 2)) // 1)
    sign = "-" if neg and q else ""
    return sign + str(q // 100) + "." + str(q % 100).zfill(2)


raw = ["0.125", "-12.21", "29.33", "", "N/A", "2.41", "0.005", "20.98"]
nonnull = [Fraction(v) for v in raw if v not in ("", "N/A")]
total = sum(nonnull)
mean = total / len(nonnull)
s = sorted(nonnull)
mid = len(s) // 2
median = s[mid] if len(s) % 2 else (s[mid - 1] + s[mid]) / 2
text = "t\n" + "\n".join(raw) + "\n"
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
r = list(csv.DictReader(io.StringIO(out)))[0]
assert r["type"] == "number"
assert r["sum"] == two(total)
assert r["mean"] == two(mean)
assert r["median"] == two(median)
=== END AC numeric-number-stats-rounding ===

=== AC numeric-half-away-negative ===
Intent: A negative exact tie rounds away from zero, so the magnitude rounds up.

import csv
import io
import os
import tempfile
from fractions import Fraction

from drydock_go import cli_runner


def two(fr):
    neg = fr < 0
    q = int((abs(fr) * 100 + Fraction(1, 2)) // 1)
    sign = "-" if neg and q else ""
    return sign + str(q // 100) + "." + str(q % 100).zfill(2)


raw = ["-1.005", "-2.000"]
fr = [Fraction(v) for v in raw]
total = sum(fr)
text = "t\n" + "\n".join(raw) + "\n"
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
r = list(csv.DictReader(io.StringIO(out)))[0]
assert r["sum"] == two(total)
assert r["mean"] == two(total / len(fr))
=== END AC numeric-half-away-negative ===

=== AC numeric-blank-for-non-numeric ===
Intent: Text, boolean, and empty columns leave sum, mean, and median blank.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

cols = {
    "word": ["a", "b", "c"],
    "flag": ["true", "false", "true"],
    "none": ["", "NA", "NULL"],
}
names = list(cols)
nrows = len(cols[names[0]])
lines = [",".join(names)]
for i in range(nrows):
    lines.append(",".join(cols[n][i] for n in names))
text = "\n".join(lines) + "\n"
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = {r["column"]: r for r in csv.DictReader(io.StringIO(out))}
for n in names:
    assert rows[n]["sum"] == "", n
    assert rows[n]["mean"] == "", n
    assert rows[n]["median"] == "", n
=== END AC numeric-blank-for-non-numeric ===

## User Acceptance

- None.

## Guardrails

- Sums, means, and medians are exact (`math/big`) and rounded once when written; no float arithmetic.
- No statistic is hardcoded; every value comes from the file examined.
- Application code never reads `sources/run.sh` or any supplied data file.
