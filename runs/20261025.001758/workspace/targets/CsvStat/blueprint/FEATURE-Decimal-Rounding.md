# FEATURE: Decimal Rounding

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Define the exact two-decimal formatting contract that rounds once from the exact value, half away from zero. |
| Depends On  | FEATURE-Profile-Numeric-Stats.md |
| Provides    | decimal formatting contract |
| Consumes    | internal/decimal |

## Purpose

Make `internal/decimal` exact and total. Every decimal the tool writes (number sum, every mean, every median, and later number min and max) goes through one formatter.

## Behavior

- Input is an exact `big.Rat`. The result is rounded once, half away from zero, to exactly two decimal places.
- A tie rounds away from zero: `0.005` becomes `0.01`, `-0.005` becomes `-0.01`, `2.675` becomes `2.68`, `1.005` becomes `1.01`.
- A leading `-` appears only when the rounded value is negative; a value that rounds to zero is written `0.00`, never `-0.00`.
- No thousands separator, no exponent, always two digits after the point, at any magnitude.
- Integer min, max, and sum remain plain integers and do not use this formatter.
- Number min and max use this formatter; they are verified by the extremes story, which is built after this one.

## Suite Pattern

The project's own `go test` suite exercises `internal/decimal` directly with table tests: ties in both signs, values just below and above a tie, zero, negative values that round to zero, and magnitudes beyond `float64`.

## Programmatic Acceptance

=== AC prf005-ties-and-sign ===
Intent: Ties round away from zero in both signs, negative values rounding to zero lose the sign, and float-hostile values round from the exact value.
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

singles = ["0.005", "-0.005", "0.004", "-0.004", "1.005", "2.675", "-2.675", "0.994", "0.0", "-0.0"]
columns = {"c" + str(i): [v] for i, v in enumerate(singles)}
rows = run_profile(columns)
for i, v in enumerate(singles):
    name = "c" + str(i)
    expected = fmt2(Fraction(v))
    assert rows[name]["type"] == "number", (name, rows[name])
    assert rows[name]["sum"] == expected, (name, v, rows[name])
    assert rows[name]["mean"] == expected, (name, v, rows[name])
    assert rows[name]["median"] == expected, (name, v, rows[name])
=== END AC prf005-ties-and-sign ===

=== AC prf005-integer-mean-ties ===
Intent: Integer column means that land on a tie round half away from zero while the integer sum stays plain.
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

pos = ["1", "0", "0", "0", "0", "0", "0", "0"]
neg = ["-1", "0", "0", "0", "0", "0", "0", "0"]
rows = run_profile({"pos": pos, "neg": neg})
for name, raw in (("pos", pos), ("neg", neg)):
    vals = [Fraction(v) for v in raw]
    assert rows[name]["type"] == "integer", (name, rows[name])
    assert rows[name]["sum"] == str(int(sum(vals))), (name, rows[name])
    assert rows[name]["mean"] == fmt2(sum(vals) / len(vals)), (name, rows[name])
    assert rows[name]["median"] == fmt2(Fraction(0)), (name, rows[name])
=== END AC prf005-integer-mean-ties ===

=== AC prf005-median-ties ===
Intent: A median that is the mean of two middle values and lands on a tie rounds half away from zero in both signs.
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

columns = {"up": ["0.00", "0.01"], "down": ["-0.01", "0.00"]}
rows = run_profile(columns)
for name, raw in columns.items():
    vals = sorted(Fraction(v) for v in raw)
    expected_median = fmt2((vals[0] + vals[1]) / 2)
    assert rows[name]["median"] == expected_median, (name, rows[name])
    assert rows[name]["mean"] == expected_median, (name, rows[name])
=== END AC prf005-median-ties ===

=== AC prf005-large-no-separator ===
Intent: Large values are written with two decimals, no thousands separator, and no exponent.
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

ints = ["1000000", "2345678"]
nums = ["1234567.891", "123456789012345678.995"]
rows = run_profile({"i": ints, "n": nums})
iv = [Fraction(v) for v in ints]
nv = [Fraction(v) for v in nums]
assert rows["i"]["sum"] == str(int(sum(iv))), rows["i"]
assert rows["i"]["mean"] == fmt2(sum(iv) / len(iv)), rows["i"]
assert rows["n"]["sum"] == fmt2(sum(nv)), rows["n"]
assert rows["n"]["mean"] == fmt2(sum(nv) / len(nv)), rows["n"]
assert rows["n"]["median"] == fmt2((nv[0] + nv[1]) / 2), rows["n"]
=== END AC prf005-large-no-separator ===

## User Acceptance

- None.

## Guardrails

- Round exactly once, from the exact rational value, half away from zero; never round through `float64`.
- Never write `-0.00`, a thousands separator, or an exponent.
- Integer min, max, and sum are never passed through the decimal formatter.
