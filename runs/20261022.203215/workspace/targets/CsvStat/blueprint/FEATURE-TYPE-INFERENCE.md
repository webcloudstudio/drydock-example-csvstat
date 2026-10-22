# FEATURE: Type Inference

| Field       | Value |
|-------------|-------|
| Version     | 20261022 V1 |
| Description | Infer the ordered type of each validated CSV column. |
| Depends On  | FEATURE-CSV-VALIDATION.md |
| Provides    | ordered column type inference |
| Consumes    | validated CSV documents and null classification |

## Purpose

Classify each column using the first applicable type in this order: integer, number, date, boolean, text, or empty.

## Rules

- Integer values match `-?[0-9]+`.
- Number values match `-?[0-9]+(\.[0-9]+)?`.
- Dates match `YYYY-MM-DD` and represent real calendar dates.
- Booleans equal `true` or `false`, ignoring case.
- Null values do not participate in inference.
- A column with no non-null values is `empty`.
- Preserve normalized source strings for later statistics and output.

## Programmatic Acceptance

=== AC type-inference-integer ===
Intent: A column containing signed whole numbers is classified as integer.
import csv
import subprocess
import tempfile
from pathlib import Path

source = "value\n-2\n0\n17\n"
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    result = subprocess.run(
        ["go", "run", "./cmd/csvstat", "--column", "value", str(path)],
        capture_output=True, text=True
    )
    assert result.returncode == 0
    rows = list(csv.DictReader(result.stdout.splitlines()))
    assert rows[0]["type"] == "integer"
=== END AC type-inference-integer ===

=== AC type-inference-date-validation ===
Intent: Impossible calendar dates prevent date classification.
import csv
import subprocess
import tempfile
from pathlib import Path

source = "value\n2025-02-30\n2025-03-01\n"
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    result = subprocess.run(
        ["go", "run", "./cmd/csvstat", "--column", "value", str(path)],
        capture_output=True, text=True
    )
    assert result.returncode == 0
    rows = list(csv.DictReader(result.stdout.splitlines()))
    assert rows[0]["type"] == "text"
=== END AC type-inference-date-validation ===

=== AC type-inference-null-and-empty ===
Intent: Null values are ignored and an all-null column is classified as empty.
import csv
import subprocess
import tempfile
from pathlib import Path

source = "value\nNA\nNULL\n\n"
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    result = subprocess.run(
        ["go", "run", "./cmd/csvstat", "--column", "value", str(path)],
        capture_output=True, text=True
    )
    assert result.returncode == 0
    rows = list(csv.DictReader(result.stdout.splitlines()))
    assert rows[0]["type"] == "empty"
=== END AC type-inference-null-and-empty ===

## User Acceptance

- None.

## Guardrails

- Type inference must operate only on validated, normalized records.
- Application code must not read supplied grader files or embed expected statistics.
