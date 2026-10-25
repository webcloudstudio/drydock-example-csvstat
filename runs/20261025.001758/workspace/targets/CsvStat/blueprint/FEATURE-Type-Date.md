# FEATURE: Type Date

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infer date columns whose non-null values are real YYYY-MM-DD calendar dates. |
| Depends On  | FEATURE-Type-Number.md |
| Provides    | type date |
| Consumes    | csvdata.IsNull |

## Purpose

Extend the type ladder so columns of calendar dates are typed `date`.

## Behavior

- A column that is neither `integer` nor `number` is `date` when every non-null trimmed value
  matches `YYYY-MM-DD` (four digits, two digits, two digits, hyphen separated) and is a real
  calendar date.
- Validate the shape first, then the calendar with `time.Parse` on the `2006-01-02` layout.
  Leap days are valid only in leap years.
- A single impossible date (for example a February 30) means the column is not `date`; it falls
  through to later types.
- Null values are ignored; a column with no non-null value is never `date`.

## Programmatic Acceptance

=== AC typ003-date-detected ===
Intent: Columns of real calendar dates, including a leap day, padding, and nulls, are typed date.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

columns = {
    "plain": ["2025-01-15", "2024-12-31", "2019-06-01", "2020-02-10"],
    "leap_day": ["2024-02-29", "2025-03-01", "2025-03-02", "2025-03-03"],
    "padded": [" 2025-01-15 ", "2025-01-16", "2025-01-17", "2025-01-18"],
    "with_nulls": ["2025-01-15", "", "NA", "NULL"],
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
    assert rows[name]["type"] == "date", name
=== END AC typ003-date-detected ===

=== AC typ003-invalid-dates-rejected ===
Intent: A column with an impossible date, a non-leap leap day, a bad month, or a non-padded form is not typed date.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

columns = {
    "feb30": ["2025-01-15", "2025-02-30", "2025-03-01", "2025-03-02"],
    "feb29_non_leap": ["2023-02-29", "2025-03-01", "2025-03-02", "2025-03-03"],
    "month13": ["2025-13-01", "2025-03-01", "2025-03-02", "2025-03-03"],
    "unpadded": ["2025-1-5", "2025-03-01", "2025-03-02", "2025-03-03"],
    "compact": ["20250115", "20250116", "20250117", "20250118"],
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
    assert rows[name]["type"] != "date", name
=== END AC typ003-invalid-dates-rejected ===

## User Acceptance

- None.

## Guardrails

- Inference uses only the values of the file examined.
- Integer and number take precedence over date whenever their patterns are satisfied.
