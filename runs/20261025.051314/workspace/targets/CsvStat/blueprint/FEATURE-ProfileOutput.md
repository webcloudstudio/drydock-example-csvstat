# FEATURE: Profile Output

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Write the per-column profile of a CSV file as RFC 4180 CSV when neither --column nor --freq is given. |
| Depends On  | FEATURE-DecimalFormat.md, FEATURE-RowCounts.md, FEATURE-Arguments.md, FEATURE-CsvReading.md |
| Provides    | csvstat <file> |
| Consumes    | csvdata.Read, profile.InferType, profile.Counts, profile.FormatDecimal, cli.ParseArgs |

## Purpose

`csvstat <file>` profiles every column of the file. The output is CSV on standard output with
the header `column,type,rows,nulls,distinct,min,max,sum,mean,median`, followed by one row per
column in file order. Line endings are `\n`.

## Behavior

- The command reads the file, trims, rejects malformed input, infers each column's type, and
  computes the counts and statistics defined by the stories it consumes.
- The output is built in memory and written to standard output only after every column has been
  computed, so a failing command writes nothing to standard output.
- Output uses Go's `encoding/csv` writer with `UseCRLF` off, so a field containing a comma, a
  double quote, or a newline is quoted as RFC 4180 specifies.
- A header-only file produces one row per header column with type `empty`, `rows` 0, and every
  other field zero or blank.
- Integer `min`, `max`, and `sum` are plain integers. Number `min`, `max`, and `sum`, and every
  `mean` and `median`, carry exactly two decimals produced by `profile.FormatDecimal`.
- Exit status is 0 on success; a rejected file exits 3 with empty standard output.

## Programmatic Acceptance

=== AC profile-output-columns ===
Intent: The profile has the declared header and one row per column in file order, with types, counts, integer sum, and decimal sum and mean derived from the input.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from decimal import Decimal, ROUND_HALF_UP
from drydock_go import cli_runner

rows = [["id", "label", "price"], ["1", "a", "1.50"], ["2", "b", "2.25"], ["4", "c", "3.00"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, code
parsed = list(csv.reader(io.StringIO(out)))
header = ["column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"]
assert parsed[0] == header
body = parsed[1:]
assert [r[0] for r in body] == rows[0]
data = rows[1:]
assert all(r[2] == str(len(data)) for r in body)
assert all(r[3] == str(0) for r in body)
assert [r[1] for r in body] == ["integer", "text", "number"]
ids = [int(r[0]) for r in data]
assert body[0][7] == str(sum(ids))
assert body[0][5] == str(min(ids)) and body[0][6] == str(max(ids))
q = Decimal("0.01")
mean = (Decimal(sum(ids)) / len(ids)).quantize(q, rounding=ROUND_HALF_UP)
assert body[0][8] == str(mean)
prices = [Decimal(r[2]) for r in data]
assert body[2][7] == str(sum(prices).quantize(q, rounding=ROUND_HALF_UP))
assert body[1][7] == "" and body[1][8] == "" and body[1][9] == ""
=== END AC profile-output-columns ===

=== AC profile-output-quoting ===
Intent: A field containing a comma and a double quote is written quoted so the parsed value equals the trimmed input value.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

value = '  x, "y"  '
rows = [["note"], [value]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, code
parsed = list(csv.reader(io.StringIO(out)))
assert len(parsed) == 2
row = parsed[1]
assert row[0] == "note" and row[1] == "text"
assert row[5] == value.strip()
assert row[6] == value.strip()
assert row[4] == "1"
=== END AC profile-output-quoting ===

=== AC profile-output-line-endings ===
Intent: Output uses newline line endings with one line per column plus the header and a trailing newline.
Requires: executable=go; scope=test

import csv
import os
import tempfile
from drydock_go import cli_runner

rows = [["a", "b", "c"], ["1", "x", "true"], ["2", "y", "false"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, code
assert "\r" not in out
assert out.endswith("\n")
assert len(out.splitlines()) == 1 + len(rows[0])
=== END AC profile-output-line-endings ===

=== AC profile-output-header-only ===
Intent: A header-only file yields one empty-typed row per column with zero rows.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

rows = [["sku", "price", "in_stock"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, code
body = list(csv.reader(io.StringIO(out)))[1:]
assert [r[0] for r in body] == rows[0]
assert all(r[1] == "empty" for r in body)
assert all(r[2] == str(0) for r in body)
assert all(r[3] == str(0) and r[4] == str(0) for r in body)
assert all(r[5:] == [""] * 5 for r in body)
=== END AC profile-output-header-only ===

=== AC profile-output-rounding ===
Intent: Number results round half away from zero once, and a total or mean that rounds to zero carries no minus sign.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from decimal import Decimal, ROUND_HALF_UP
from drydock_go import cli_runner

rows = [["pos", "neg", "tiny"], ["0.125", "-0.125", "0.001"], ["0.125", "-0.125", "-0.002"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, code
body = list(csv.reader(io.StringIO(out)))[1:]
q = Decimal("0.01")
pos_mean = Decimal("0.125").quantize(q, rounding=ROUND_HALF_UP)
neg_mean = Decimal("-0.125").quantize(q, rounding=ROUND_HALF_UP)
assert body[0][8] == str(pos_mean)
assert body[1][8] == str(neg_mean)
zero = f"{0:.2f}"
assert body[2][7] == zero
assert body[2][8] == zero
=== END AC profile-output-rounding ===

=== AC profile-output-rejected-input ===
Intent: A ragged file exits 3 and writes nothing to standard output.
Requires: executable=go; scope=test

import csv
import os
import tempfile
from drydock_go import cli_runner

rows = [["a", "b"], ["1", "2"], ["3", "4", "5"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "ragged.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner(path)
print(err)
assert code == 3, code
assert out == ""
=== END AC profile-output-rejected-input ===

## User Acceptance

- None.

## Guardrails

- Standard output is written only on success; a failing command leaves it empty.
- Application code never reads `sources/run.sh` or embeds a statistic, a type, or a value list.
- Nothing under `sources/` is modified.
