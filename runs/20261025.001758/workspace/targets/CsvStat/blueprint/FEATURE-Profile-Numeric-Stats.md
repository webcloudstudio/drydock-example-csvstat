# FEATURE: Profile Numeric Stats

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Compute exact sum, mean, and median for integer and number columns using arbitrary-precision rational arithmetic. |
| Depends On  | FEATURE-Type-Text-Empty.md, FEATURE-Profile-Counts.md |
| Provides    | profile sum, mean, median |
| Consumes    | infer types, internal/decimal |

## Purpose

Fill the `sum`, `mean`, and `median` fields for columns typed `integer` or `number`. All arithmetic is exact (`math/big`, `big.Rat`); rounding happens once, when the value is written.

## Behavior

- Only non-null values participate. The count used for the mean is the non-null count.
- `sum`: integer columns write the integer total; number columns write the total with exactly two decimal places.
- `mean`: sum divided by the non-null count, written with two decimal places for both integer and number columns.
- `median`: the middle value of the numerically sorted non-null values, or the mean of the two middle values when the count is even; written with two decimal places for both types.
- Sorting is numeric, never lexical (`9` sorts before `10`).
- For every other type (`text`, `date`, `boolean`, `empty`) the three fields are empty.
- `internal/decimal` provides the formatter that rounds a `big.Rat` once, half away from zero, to two decimals. Its exact contract is owned by the decimal rounding story.

## Programmatic Acceptance

=== AC prf004-integer-stats ===
Intent: Integer columns report the exact integer sum and two-decimal mean and median over non-null values.
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

def median(vals):
    s = sorted(vals)
    m = len(s)
    return s[m // 2] if m % 2 else (s[m // 2 - 1] + s[m // 2]) / 2

raw = ["3", "1", "2", "10", "NA", ""]
vals = [Fraction(v) for v in raw if v not in ("NA", "")]
rows = run_profile({"n": raw})
expected_sum = str(int(sum(vals)))
expected_mean = fmt2(sum(vals) / len(vals))
expected_median = fmt2(median(vals))
assert rows["n"]["type"] == "integer", rows["n"]
assert rows["n"]["sum"] == expected_sum, rows["n"]
assert rows["n"]["mean"] == expected_mean, rows["n"]
assert rows["n"]["median"] == expected_median, rows["n"]
=== END AC prf004-integer-stats ===

=== AC prf004-number-stats ===
Intent: Number columns report exact two-decimal sum, mean, and median, including sums that binary floating point gets wrong.
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

def median(vals):
    s = sorted(vals)
    m = len(s)
    return s[m // 2] if m % 2 else (s[m // 2 - 1] + s[m // 2]) / 2

columns = {
    "floaty": ["0.1", "0.2", "0.3", "0.7", "0.6", "NULL"],
    "neg": ["-1.5", "-2.25", "-0.75", "", "NA", "N/A"],
    "mixed": ["1", "2.5", "-3.25", "10.10", "4", "null"],
}
rows = run_profile(columns)
for name, raw in columns.items():
    vals = [Fraction(v) for v in raw if v.lower() not in ("", "na", "n/a", "null")]
    assert rows[name]["type"] == "number", (name, rows[name])
    expected_sum = fmt2(sum(vals))
    expected_mean = fmt2(sum(vals) / len(vals))
    expected_median = fmt2(median(vals))
    assert rows[name]["sum"] == expected_sum, (name, rows[name])
    assert rows[name]["mean"] == expected_mean, (name, rows[name])
    assert rows[name]["median"] == expected_median, (name, rows[name])
=== END AC prf004-number-stats ===

=== AC prf004-median-parity-and-order ===
Intent: Median uses numeric order and averages the two middle values for even counts.
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

def median(vals):
    s = sorted(vals)
    m = len(s)
    return s[m // 2] if m % 2 else (s[m // 2 - 1] + s[m // 2]) / 2

cases = {
    "odd": ["5", "1", "9", "NA", "", "NULL"],
    "even": ["4", "1", "3", "2"],
    "lexical_trap": ["10", "9", "100", "2"],
}
# every column in one file needs equal length; pad shorter ones with nulls
width = max(len(v) for v in cases.values())
columns = {k: v + [""] * (width - len(v)) for k, v in cases.items()}
rows = run_profile(columns)
for name, raw in cases.items():
    vals = [Fraction(v) for v in raw if v.lower() not in ("", "na", "n/a", "null")]
    assert rows[name]["median"] == fmt2(median(vals)), (name, rows[name])
=== END AC prf004-median-parity-and-order ===

=== AC prf004-big-exact ===
Intent: Sums beyond float precision stay exact for integer and number columns.
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

big_ints = ["9007199254740993", "1"]
big_nums = ["12345678901234567890.12", "0.01"]
rows = run_profile({"bi": big_ints, "bn": big_nums})
iv = [Fraction(v) for v in big_ints]
nv = [Fraction(v) for v in big_nums]
assert rows["bi"]["sum"] == str(int(sum(iv))), rows["bi"]
assert rows["bi"]["mean"] == fmt2(sum(iv) / len(iv)), rows["bi"]
assert rows["bn"]["sum"] == fmt2(sum(nv)), rows["bn"]
assert rows["bn"]["mean"] == fmt2(sum(nv) / len(nv)), rows["bn"]
=== END AC prf004-big-exact ===

=== AC prf004-non-numeric-empty ===
Intent: Text, date, boolean, and empty columns leave sum, mean, and median empty.
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
    "t": ["abc", "1", "x"],
    "d": ["2025-01-01", "2025-02-03", "2024-12-31"],
    "b": ["true", "false", "TRUE"],
    "e": ["", "NA", "NULL"],
}
expected_types = {"t": "text", "d": "date", "b": "boolean", "e": "empty"}
rows = run_profile(columns)
for name in columns:
    assert rows[name]["type"] == expected_types[name], (name, rows[name])
    for field in ("sum", "mean", "median"):
        assert rows[name][field] == "", (name, field, rows[name])
=== END AC prf004-non-numeric-empty ===

## User Acceptance

- None.

## Guardrails

- Sums, means, and medians are computed exactly with `math/big`; never with `float64`.
- Round once, at output; never round an intermediate value.
- Do not hardcode any statistic; every value comes from the file examined.
