# FEATURE: CLI Input

| Field       | Value |
|-------------|-------|
| Version     | 20261023 V1 |
| Description | Reads, normalizes, and structurally validates CSV input. |
| Depends On  | FEATURE-CLI-FOUNDATION.md |
| Provides    | RFC 4180 CSV parsing, trimmed records, structural validation |
| Consumes    | ./cmd/csvstat |

## Input Rules

Read UTF-8 CSV with `encoding/csv`. Trim every header and value before null, type, comparison, and aggregation processing. Preserve quoted commas and reject empty files, empty or repeated headers, and records whose field count differs from the header.

## Diagnostics

Rejected records retain their source line information so a diagnostic can identify the file and rejected record line.

## Programmatic Acceptance

=== AC cli-input-quoted ===
Intent: A supplied CSV containing quoted commas is accepted.
from pathlib import Path
import subprocess

asset = Path("sources/employees.csv")
result = subprocess.run(
    ["./csvstat-ac", str(asset)],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
assert result.stdout != ""
=== END AC cli-input-quoted ===

=== AC cli-input-ragged ===
Intent: A ragged CSV is rejected without standard output.
from pathlib import Path
import subprocess

asset = Path("sources/ragged.csv")
result = subprocess.run(
    ["./csvstat-ac", str(asset)],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 3
assert result.stdout == ""
assert result.stderr != ""
=== END AC cli-input-ragged ===

=== AC cli-input-header-only ===
Intent: A header-only CSV is accepted as a valid empty data set.
from pathlib import Path
import subprocess

asset = Path("sources/header_only.csv")
result = subprocess.run(
    ["./csvstat-ac", str(asset)],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
assert result.stdout != ""
=== END AC cli-input-header-only ===

## User Acceptance

- None.

## Guardrails

- Supplied source assets remain byte-for-byte unchanged.
- Application code does not read `sources/run.sh`.
- CSV parsing uses the standard library.
