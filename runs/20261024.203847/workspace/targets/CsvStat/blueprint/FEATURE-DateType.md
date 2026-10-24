# FEATURE: Date Type

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Recognize date columns whose every non-null value is a real YYYY-MM-DD calendar date. |
| Depends On  | FEATURE-TypeInference.md, ARCHITECTURE.md |
| Provides    | typing.IsDate |
| Consumes    | typing.Infer |

## Purpose

Let a column of calendar dates be typed `date`, while near-misses such as `2025-02-30` leave the column as `text`.

## Behavior

- `typing.IsDate(s string) bool` is true only when `s` matches `YYYY-MM-DD` (four, two, two digits, hyphen-separated) and names a real calendar date, leap years included.
- Validation uses `time.Parse` with layout `2006-01-02` plus a check that the formatted result equals the input, or an equivalent exact check, so no value is normalized into validity.
- The date hook in `typing.Infer` is bound to `typing.IsDate`. A column is `date` only when every non-null value satisfies it and not every value is an integer or number.
- One non-date value makes the column fall to the next matching type (`boolean` or `text`).
- Timestamps such as `2025-03-01T08:00` are not dates.

## Programmatic Acceptance

=== AC date-valid-column ===
Intent: A column of real calendar dates, including a leap day, is typed date.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

values = ["2025-01-31", "2024-02-29", "2015-08-27", "2025-12-01"]
text = "when\n" + "\n".join(values) + "\n"
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.DictReader(io.StringIO(out)))
assert len(rows) == 1
assert rows[0]["type"] == "date"
=== END AC date-valid-column ===

=== AC date-invalid-calendar-day ===
Intent: A single impossible calendar date makes the whole column text.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

good = ["2025-03-18", "2025-03-05"]
bad = ["2025-02-30"]
text = "shipped\n" + "\n".join(good + bad) + "\n"
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.DictReader(io.StringIO(out)))
assert rows[0]["type"] == "text"
=== END AC date-invalid-calendar-day ===

=== AC date-non-leap-and-month-bounds ===
Intent: Feb 29 in a non-leap year, month 13, and day 00 each fail date typing.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

cases = {
    "nonleap": ["2025-02-29"],
    "month13": ["2025-13-01"],
    "day00": ["2025-01-00"],
}
names = list(cases)
lines = [",".join(names), ",".join(cases[n][0] for n in names)]
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
    assert rows[n]["type"] == "text", n
=== END AC date-non-leap-and-month-bounds ===

=== AC date-nulls-and-timestamps ===
Intent: Nulls do not disturb a date column, while timestamp values are not dates.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

cols = {
    "d": ["2025-03-01", "", "NA", "2025-03-02"],
    "ts": ["2025-03-01T08:00", "2025-03-01T09:00", "2025-03-02T10:00", "2025-03-03T11:00"],
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
assert rows["d"]["type"] == "date"
assert rows["ts"]["type"] == "text"
=== END AC date-nulls-and-timestamps ===

## User Acceptance

- None.

## Guardrails

- Calendar validity is decided by the standard library, not by a hardcoded list of dates.
- Standard library only.
