# FEATURE: CSV Ingestion

| Field       | Value |
|-------------|-------|
| Version     | 20261022 V1 |
| Description | Reads RFC 4180 CSV input and exposes trimmed records in source order. |
| Depends On  | ARCHITECTURE.md |
| Provides    | normalized CSV document model |
| Consumes    | Go module, ./cmd/csvstat entry point |

## Ingestion Rules

- Read the first CSV record as the header.
- Parse RFC 4180 quoting with `encoding/csv`.
- Trim leading and trailing spaces from every header and value after parsing.
- Preserve record order.
- Preserve commas, quotes, and newlines contained in quoted fields.
- Pass the normalized document to validation and profiling.

## Programmatic Acceptance

=== AC csv-ingestion-valid ===
Intent: A valid RFC 4180 document with padded fields and a quoted comma is accepted for profiling.

import pathlib
import subprocess
import tempfile

with tempfile.TemporaryDirectory() as directory:
    path = pathlib.Path(directory) / "quoted.csv"
    path.write_text(
        ' name ,note,value\n Alice ,"hello, world", 7 \n',
        encoding="utf-8",
    )

    result = subprocess.run(
        ["go", "run", "./cmd/csvstat", str(path)],
        capture_output=True,
        text=True,
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
    assert result.stdout != ""
=== END AC csv-ingestion-valid ===

=== AC csv-ingestion-multiline ===
Intent: A quoted field containing a newline is parsed as one CSV field rather than a malformed record.

import pathlib
import subprocess
import tempfile

with tempfile.TemporaryDirectory() as directory:
    path = pathlib.Path(directory) / "multiline.csv"
    path.write_text(
        'name,note\nAlice,"first line\nsecond line"\n',
        encoding="utf-8",
    )

    result = subprocess.run(
        ["go", "run", "./cmd/csvstat", "--column", "note", str(path)],
        capture_output=True,
        text=True,
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
    assert result.stdout != ""
=== END AC csv-ingestion-multiline ===

## User Acceptance

- None.

## Guardrails

- Use `encoding/csv`; do not split records manually on commas or newlines.
- Normalize only after CSV parsing.
- Preserve the input record order.
