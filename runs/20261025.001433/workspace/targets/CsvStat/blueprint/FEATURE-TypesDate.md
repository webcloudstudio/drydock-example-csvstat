# FEATURE: Types Date

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infer the date column type and reject values that are not real calendar dates. |
| Depends On  | FEATURE-ProfileCounts.md |
| Provides    | profile field type value date |
| Consumes    | null rule from FEATURE-ProfileCounts.md |

## Purpose

A column is a `date` column when every non-null value is a real calendar date written `YYYY-MM-DD`. This specification owns the date test; the final type decision is composed in FEATURE-TypesOther.md.

## Behavior

- A value qualifies when, after trimming, it is exactly ten characters, matches `YYYY-MM-DD` (four digits, `-`, two digits, `-`, two digits), and names a real calendar date.
- A value such as `2025-02-30`, `2025-13-01`, or `2025-00-10` does not qualify. `2024-02-29` qualifies (leap year); `2025-02-29` does not.
- One non-qualifying non-null value makes the column not `date`; the column then falls through to the next type in order.
- Null values never disqualify the type.
- A column of integers or numbers is typed by FEATURE-TypesNumeric.md before the date test applies; a date value never matches the integer or number patterns.
- Implementation uses `time.Parse` with layout `2006-01-02` plus a length check so that non-padded forms such as `2025-2-3` are rejected.

## Programmatic Acceptance

=== AC types-date-valid ===
Intent: A column of real YYYY-MM-DD dates, including a leap day and nulls, is typed date.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def profile(lines):
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write("".join(line + "\n" for line in lines))
        code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}


values = ["2025-02-01", "2024-02-29", " 2023-12-31 ", "NA", "", "1999-01-01"]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
assert rows["v"]["type"] == "date"
=== END AC types-date-valid ===

=== AC types-date-invalid-calendar ===
Intent: Values that look like dates but are not real calendar dates keep the column from being date.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def profile(lines):
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write("".join(line + "\n" for line in lines))
        code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}


bad_values = ["2025-02-30", "2025-02-29", "2025-13-01", "2025-00-10", "2025-04-31", "2025-01-32"]
for bad in bad_values:
    values = ["2025-02-01", "2025-03-18", bad]
    lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
    rows = profile(lines)
    assert rows["v"]["type"] != "date", (bad, rows["v"]["type"])
=== END AC types-date-invalid-calendar ===

=== AC types-date-wrong-shape ===
Intent: Dates not written exactly as zero-padded YYYY-MM-DD do not qualify.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def profile(lines):
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write("".join(line + "\n" for line in lines))
        code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}


bad_values = ["2025-2-3", "2025/02/03", "03-02-2025", "2025-02-03T08:00", "20250203"]
for bad in bad_values:
    values = ["2025-02-01", bad]
    lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
    rows = profile(lines)
    assert rows["v"]["type"] != "date", (bad, rows["v"]["type"])
=== END AC types-date-wrong-shape ===

## User Acceptance

- None.

## Guardrails

- Date decisions are computed from the values of the file examined; no column name or file is special-cased.
- Only the `YYYY-MM-DD` form is a date; no other layout is accepted.
- Standard library only.
