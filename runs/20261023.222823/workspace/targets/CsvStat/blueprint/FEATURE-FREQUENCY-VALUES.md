# FEATURE: Frequency Values

| Field       | Value |
|-------------|-------|
| Version     | 20261023 V1 |
| Description | Aggregate and deterministically order non-null column values by frequency. |
| Depends On  | FEATURE-CLI-COLUMNS.md |
| Provides    | frequency aggregation and deterministic ordering |
| Consumes    | column lookup, null normalization |

## Purpose

The frequency capability counts each distinct non-null normalized value in the requested column. It excludes empty values and the case-insensitive null spellings `NA`, `N/A`, and `NULL`.

Rows are ordered by descending count and then by bytewise ascending value.

## Programmatic Acceptance

=== AC frequency-excludes-nulls ===
Intent: Frequency output omits every recognized null spelling.

import csv
import tempfile
from pathlib import Path
from drydock_go import cli_runner

source = "value\nalpha\nNA\nn/a\nNULL\n\nbeta\n"
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    code, output, error = cli_runner("--freq", "value", str(path), package="./cmd/csvstat")
    assert code == 0
    rows = list(csv.reader(output.splitlines()))
    values = [row[0] for row in rows[1:]]
    assert values == ["alpha", "beta"]
    assert all(value not in {"", "NA", "n/a", "NULL"} for value in values)
=== END AC frequency-excludes-nulls ===

=== AC frequency-orders-counts ===
Intent: Frequency rows are ordered by descending count before value ordering.

import csv
import tempfile
from pathlib import Path
from drydock_go import cli_runner

source = "value\nb\na\nb\nc\na\nb\n"
expected_counts = {"a": 2, "b": 3, "c": 1}
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    code, output, error = cli_runner("--freq", "value", str(path), package="./cmd/csvstat")
    assert code == 0
    rows = list(csv.reader(output.splitlines()))[1:]
    actual = [(row[0], int(row[1])) for row in rows]
    assert actual == sorted(expected_counts.items(), key=lambda item: (-item[1], item[0]))
=== END AC frequency-orders-counts ===

=== AC frequency-orders-ties ===
Intent: Equal-frequency values are ordered by ascending byte value.

import csv
import tempfile
from pathlib import Path
from drydock_go import cli_runner

source = "value\nz\na\nm\nz\na\nm\n"
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    code, output, error = cli_runner("--freq", "value", str(path), package="./cmd/csvstat")
    assert code == 0
    rows = list(csv.reader(output.splitlines()))[1:]
    assert [row[0] for row in rows] == sorted(["a", "m", "z"])
=== END AC frequency-orders-ties ===

## User Acceptance

- None.

## Guardrails

- Null values must never appear in frequency rows.
- Frequency ordering must be deterministic and independent of input map iteration order.
