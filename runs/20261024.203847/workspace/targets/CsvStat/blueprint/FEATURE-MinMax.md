# FEATURE: Min and Max

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Compute the minimum and maximum of each column according to its inferred type. |
| Depends On  | FEATURE-TypeInference.md, FEATURE-DateType.md, ARCHITECTURE.md |
| Provides    | profile.MinMax |
| Consumes    | typing.Infer, typing.IsDate |

## Purpose

Supply the `min` and `max` fields of each profile row.

## Behavior

- Extremes are taken over non-null trimmed values, using exact arithmetic (`math/big`) for numeric types, never floats.
- `integer`: smallest and largest written as plain integers. Values beyond 64 bits are handled exactly.
- `number`: smallest and largest written with exactly two decimal places, rounded half away from zero from the exact value.
- `date`: earliest and latest, as written.
- `boolean`: `min` is `false` if any value is false, else `true`; `max` is `true` if any value is true, else `false`. Output is lowercase.
- `text`: smallest and largest by byte value of the trimmed strings.
- `empty`: both blank.
- `profile.MinMax(typ string, values []string) (min, max string)`.

## Programmatic Acceptance

=== AC minmax-integer ===
Intent: Integer min and max, including negatives and a value beyond 64 bits, are exact plain integers.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

ints = [42, -7, 0, 12345678901234567890123, -99, 5]
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
assert r["type"] == "integer"
assert r["min"] == str(min(ints))
assert r["max"] == str(max(ints))
=== END AC minmax-integer ===

=== AC minmax-number-two-places ===
Intent: Number min and max are written with two decimal places rounded half away from zero.

import csv
import io
import os
import tempfile
from decimal import Decimal, ROUND_HALF_UP

from drydock_go import cli_runner

values = ["15.27", "231.18", "2.005", "100", "7.5"]
dec = [Decimal(v) for v in values]
cent = Decimal("0.01")
want_min = str(min(dec).quantize(cent, rounding=ROUND_HALF_UP))
want_max = str(max(dec).quantize(cent, rounding=ROUND_HALF_UP))
text = "p\n" + "\n".join(values) + "\n"
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
assert r["min"] == want_min
assert r["max"] == want_max
=== END AC minmax-number-two-places ===

=== AC minmax-number-negative ===
Intent: Negative decimals order numerically, not lexically, and keep their sign.

import csv
import io
import os
import tempfile
from decimal import Decimal, ROUND_HALF_UP

from drydock_go import cli_runner

values = ["-12.21", "-0.44", "29.81", "6.81", "-9.02"]
dec = [Decimal(v) for v in values]
cent = Decimal("0.01")
want_min = str(min(dec).quantize(cent, rounding=ROUND_HALF_UP))
want_max = str(max(dec).quantize(cent, rounding=ROUND_HALF_UP))
text = "t\n" + "\n".join(values) + "\n"
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
r = list(csv.DictReader(io.StringIO(out)))[0]
assert r["min"] == want_min
assert r["max"] == want_max
=== END AC minmax-number-negative ===

=== AC minmax-date-boolean-text ===
Intent: Date, boolean, and text columns report earliest/latest, any-false/any-true, and byte-order extremes.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

cols = {
    "when": ["2025-03-18", "2015-08-27", "2025-06-02", "2020-01-18"],
    "allt": ["true", "TRUE", "True", "true"],
    "mix": ["true", "False", "true", "false"],
    "word": ["mentor", "Zed", "remote", "apple"],
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
assert rows["when"]["type"] == "date"
assert rows["when"]["min"] == min(cols["when"])
assert rows["when"]["max"] == max(cols["when"])
assert rows["allt"]["type"] == "boolean"
assert rows["allt"]["min"] == "true"
assert rows["allt"]["max"] == "true"
assert rows["mix"]["type"] == "boolean"
assert rows["mix"]["min"] == "false"
assert rows["mix"]["max"] == "true"
key = lambda s: s.encode("utf-8")
assert rows["word"]["type"] == "text"
assert rows["word"]["min"] == min(cols["word"], key=key)
assert rows["word"]["max"] == max(cols["word"], key=key)
=== END AC minmax-date-boolean-text ===

=== AC minmax-empty-and-nulls ===
Intent: An empty column has blank min and max, and nulls are excluded from extremes.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

cols = {
    "none": ["", "NA", "N/A", "NULL"],
    "n": ["", "9", "NA", "3"],
}
names = list(cols)
nrows = len(cols[names[0]])
lines = [",".join(names)]
for i in range(nrows):
    lines.append(",".join(cols[n][i] for n in names))
text = "\n".join(lines) + "\n"
nonnull = [int(v) for v in cols["n"] if v not in ("", "NA")]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = {r["column"]: r for r in csv.DictReader(io.StringIO(out))}
assert rows["none"]["type"] == "empty"
assert rows["none"]["min"] == ""
assert rows["none"]["max"] == ""
assert rows["n"]["min"] == str(min(nonnull))
assert rows["n"]["max"] == str(max(nonnull))
=== END AC minmax-empty-and-nulls ===

## User Acceptance

- None.

## Guardrails

- No floating-point arithmetic for numeric extremes.
- No statistic or value is hardcoded; every value comes from the file examined.
