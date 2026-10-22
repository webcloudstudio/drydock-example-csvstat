# FEATURE: Frequency Reports

| Field       | Value |
|-------------|-------|
| Version     | 20261022 V1 |
| Description | Produce ordered frequency reports for a selected CSV column. |
| Depends On  | FEATURE-CSV-VALIDATION.md |
| Provides    | frequency CSV reports with --top limiting |
| Consumes    | validated CSV documents and null classification |

## Purpose

Count normalized non-null values for a named column and emit the requested top-frequency rows.

## Behavior

- Exclude null values.
- Order by count descending.
- Break equal-count ties by bytewise ascending value.
- Apply a positive `--top` limit, defaulting to five.
- Emit `value,count` even when the column has no non-null values.
- Reject absent columns through the validation contract.

## Programmatic Acceptance

=== AC frequency-ordering ===
Intent: Frequency output is valid CSV with the declared value and count columns.
import csv
import subprocess
import tempfile
from pathlib import Path

source = "value\na\nb\na\nc\n"
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    result = subprocess.run(
        ["go", "run", "./cmd/csvstat", "--freq", "value", str(path)],
        capture_output=True, text=True
    )
    assert result.returncode == 0
    rows = list(csv.DictReader(result.stdout.splitlines()))
    assert rows
    assert set(rows[0]) == {"value", "count"}
=== END AC frequency-ordering ===

=== AC frequency-top-limit ===
Intent: The positive top option limits the number of emitted frequency rows.
import csv
import subprocess
import tempfile
from pathlib import Path

source = "value\na\nb\nc\nd\n"
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    result = subprocess.run(
        ["go", "run", "./cmd/csvstat", "--freq", "value", "--top", "2", str(path)],
        capture_output=True, text=True
    )
    assert result.returncode == 0
    rows = list(csv.DictReader(result.stdout.splitlines()))
    assert len(rows) <= 2
=== END AC frequency-top-limit ===

=== AC frequency-null-exclusion ===
Intent: A column containing only null values emits the frequency header and no data rows.
import csv
import subprocess
import tempfile
from pathlib import Path

source = "value\nNA\nNULL\n\n"
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    result = subprocess.run(
        ["go", "run", "./cmd/csvstat", "--freq", "value", str(path)],
        capture_output=True, text=True
    )
    assert result.returncode == 0
    rows = list(csv.DictReader(result.stdout.splitlines()))
    assert rows == []
=== END AC frequency-null-exclusion ===

## User Acceptance

- None.

## Guardrails

- Frequency ordering must be deterministic.
- Null spellings must never appear as frequency values.
- Frequency output must not include statistics or unrelated columns.
