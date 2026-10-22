# FEATURE: CSV Validation

| Field       | Value |
|-------------|-------|
| Version     | 20261022 V1 |
| Description | Validates CSV structure, classifies nulls, and maps rejected and operational failures to exit codes. |
| Depends On  | FEATURE-CSV-INGESTION.md |
| Provides    | validated CSV documents and null classification |
| Consumes    | normalized CSV document model |

## Validation Rules

Reject the input with exit code `3` when:

- The file is empty.
- A header name is empty.
- A header name is repeated.
- A record has a different field count from the header.
- A requested column is absent.

Treat a value as null when it is empty or equals `NA`, `N/A`, or `NULL`, ignoring case, after trimming.
All other values are non-null.

Return exit code `1` when the file cannot be opened or read. Rejected and operational failures write no
standard output.

## Programmatic Acceptance

=== AC csv-validation-ragged ===
Intent: A record with the wrong field count is rejected with exit code 3 and no standard output.

import pathlib
import subprocess
import tempfile

with tempfile.TemporaryDirectory() as directory:
    path = pathlib.Path(directory) / "ragged.csv"
    path.write_text("a,b\n1,2\n3,4,5\n", encoding="utf-8")

    result = subprocess.run(
        ["go", "run", "./cmd/csvstat", str(path)],
        capture_output=True,
        text=True,
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 3
    assert result.stdout == ""
=== END AC csv-validation-ragged ===

=== AC csv-validation-column ===
Intent: A requested column absent from the header is rejected with exit code 3 and no standard output.

import pathlib
import subprocess
import tempfile

with tempfile.TemporaryDirectory() as directory:
    path = pathlib.Path(directory) / "data.csv"
    path.write_text("a,b\n1,2\n", encoding="utf-8")

    result = subprocess.run(
        ["go", "run", "./cmd/csvstat", "--column", "missing", str(path)],
        capture_output=True,
        text=True,
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 3
    assert result.stdout == ""
=== END AC csv-validation-column ===

=== AC csv-validation-operational ===
Intent: An unreadable or missing input path returns the operational-failure exit code and no standard output.

import pathlib
import subprocess
import tempfile

with tempfile.TemporaryDirectory() as directory:
    missing = pathlib.Path(directory) / "does-not-exist.csv"

    result = subprocess.run(
        ["go", "run", "./cmd/csvstat", str(missing)],
        capture_output=True,
        text=True,
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 1
    assert result.stdout == ""
=== END AC csv-validation-operational ===

=== AC csv-validation-nulls ===
Intent: Empty and recognized null spellings are accepted as valid input rather than rejected as malformed CSV.

import pathlib
import subprocess
import tempfile

with tempfile.TemporaryDirectory() as directory:
    path = pathlib.Path(directory) / "nulls.csv"
    path.write_text("a,b,c,d\n,NA,n/A,NULL\n1,2,3,4\n", encoding="utf-8")

    result = subprocess.run(
        ["go", "run", "./cmd/csvstat", str(path)],
        capture_output=True,
        text=True,
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
    assert result.stdout != ""
=== END AC csv-validation-nulls ===

## User Acceptance

- None.

## Guardrails

- Do not treat a ragged record as a usable partial record.
- Do not classify arbitrary text as null.
- Do not write partial output before validation completes.
- Diagnostics may identify the file and rejected record line, but diagnostics are never written to
  standard output.
