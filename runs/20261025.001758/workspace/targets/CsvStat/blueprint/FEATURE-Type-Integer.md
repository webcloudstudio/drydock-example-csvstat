# FEATURE: Type Integer

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infer integer columns and report the type in the profile row. |
| Depends On  | FEATURE-Profile-Rows.md |
| Provides    | type integer |
| Consumes    | csvdata.IsNull |

## Purpose

Start the first-match type ladder in `internal/infer` and fill the `type` field of each profile row
for integer columns.

## Behavior

- A column is `integer` when it has at least one non-null value and every non-null value, after
  trimming, matches `-?[0-9]+`.
- Null values (empty, `NA`, `N/A`, `NULL` ignoring case) are ignored for inference.
- Leading zeros are allowed. A decimal point, a leading `+`, an exponent, or any other character
  makes the column not integer.
- Columns this story does not classify leave `type` empty until later type stories are built.
  Column with no non-null value is never `integer`.

## Interfaces

- `internal/infer` exposes a classification function over a column's non-null trimmed values; the
  profile writer calls it per column.

## Programmatic Acceptance

=== AC typ001-integer-detected ===
Intent: Columns whose non-null values all match the integer pattern, including negatives, leading zeros, padding, and nulls, are typed integer.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

columns = {
    "plain": ["1", "-20", "0", "007"],
    "padded": [" 5 ", "6", "-7", "8"],
    "with_nulls": ["1", "", "NA", "N/A"],
    "with_null_word": ["10", "NULL", "null", "-3"],
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
    assert rows[name]["type"] == "integer", name
=== END AC typ001-integer-detected ===

=== AC typ001-non-integer-rejected ===
Intent: Columns with a decimal, a plus sign, an exponent, or a non-numeric value, and columns with no non-null value, are not typed integer.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

columns = {
    "decimal": ["1", "2.5", "3", "4"],
    "plus": ["+1", "2", "3", "4"],
    "exponent": ["1e3", "2", "3", "4"],
    "word": ["1", "two", "3", "4"],
    "all_null": ["", "NA", "N/A", "NULL"],
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
    assert rows[name]["type"] != "integer", name
=== END AC typ001-non-integer-rejected ===

## User Acceptance

- None.

## Guardrails

- Inference uses only the values of the file examined; no column name or value is hardcoded.
- Null values never influence the type.
