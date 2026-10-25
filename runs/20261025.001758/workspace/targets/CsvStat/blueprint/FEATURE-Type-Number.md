# FEATURE: Type Number

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infer number columns for decimal values that are not all integers. |
| Depends On  | FEATURE-Type-Integer.md |
| Provides    | type number |
| Consumes    | csvdata.IsNull |

## Purpose

Extend the type ladder so columns with decimal values are typed `number`.

## Behavior

- A column that is not `integer` is `number` when every non-null trimmed value matches
  `-?[0-9]+(\.[0-9]+)?`.
- A column of only integer-shaped values stays `integer` (first match wins).
- Forms such as `1.`, `.5`, `+1`, `1e3`, and `1,5` are not numbers.
- Null values are ignored; a column with no non-null value is never `number`.

## Programmatic Acceptance

=== AC typ002-number-detected ===
Intent: Columns mixing integers and decimals, including negatives and padding and nulls, are typed number.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

columns = {
    "mixed": ["1", "2.50", "-3", "4"],
    "all_decimal": ["0.5", "-0.25", "3.14159", "10.0"],
    "padded": [" 1.5 ", "2", "3", "4"],
    "with_nulls": ["1.5", "", "NA", "N/A"],
}
names = list(columns)
count = len(columns[names[0]])
body = io.StringIO()
writer = csv.writer(body, lineterminator="\n")
writer.writerow(names)
for i in range(count):
    writer.writerow([columns[k][i] for k in names])

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "t.csv")
    with open(path, "w", newline="") as f:
        f.write(body.getvalue())
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0
rows = {r["column"]: r for r in csv.DictReader(io.StringIO(out))}
for name in names:
    assert rows[name]["type"] == "number", name
=== END AC typ002-number-detected ===

=== AC typ002-integer-wins-and-malformed-rejected ===
Intent: An all-integer column stays integer, and malformed decimal forms are typed neither number nor integer.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

columns = {
    "ints": ["1", "2", "3", "4"],
    "trailing_dot": ["1.", "2.5", "3", "4"],
    "leading_dot": [".5", "2.5", "3", "4"],
    "plus": ["+1.5", "2.5", "3", "4"],
    "exponent": ["1e3", "2.5", "3", "4"],
    "words": ["1.5", "abc", "3", "4"],
}
names = list(columns)
count = len(columns[names[0]])
body = io.StringIO()
writer = csv.writer(body, lineterminator="\n")
writer.writerow(names)
for i in range(count):
    writer.writerow([columns[k][i] for k in names])

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "t.csv")
    with open(path, "w", newline="") as f:
        f.write(body.getvalue())
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0
rows = {r["column"]: r for r in csv.DictReader(io.StringIO(out))}
assert rows["ints"]["type"] == "integer"
for name in names[1:]:
    assert rows[name]["type"] not in ("number", "integer"), name
=== END AC typ002-integer-wins-and-malformed-rejected ===

## User Acceptance

- None.

## Guardrails

- Inference uses only the values of the file examined.
- Integer takes precedence over number whenever both patterns are satisfied.
