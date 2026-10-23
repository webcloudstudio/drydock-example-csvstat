# FEATURE: CLI Columns

| Field       | Value |
|-------------|-------|
| Version     | 20261023 V1 |
| Description | Resolves named columns and normalizes null values for profiling commands. |
| Depends On  | FEATURE-CLI-INPUT.md |
| Provides    | column lookup, null normalization |
| Consumes    | RFC 4180 CSV parsing, trimmed records, structural validation |

## Null and Column Rules

After trimming, a value is null when it is empty or equals `NA`, `N/A`, or `NULL`, ignoring case. Named columns are resolved against the validated, trimmed header. An absent named column is rejected with exit code 3 and no standard output.

## Programmatic Acceptance

=== AC cli-columns-existing ===
Intent: A named existing column is accepted for profiling.
from pathlib import Path
import subprocess

asset = Path("sources/employees.csv")
result = subprocess.run(
    ["./csvstat-ac", "--column", "salary", str(asset)],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
assert result.stdout != ""
=== END AC cli-columns-existing ===

=== AC cli-columns-missing-profile ===
Intent: A missing profile column is rejected without standard output.
from pathlib import Path
import subprocess

asset = Path("sources/orders.csv")
result = subprocess.run(
    ["./csvstat-ac", "--column", "nosuch", str(asset)],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 3
assert result.stdout == ""
=== END AC cli-columns-missing-profile ===

=== AC cli-columns-missing-frequency ===
Intent: A missing frequency column is rejected without standard output.
from pathlib import Path
import subprocess

asset = Path("sources/orders.csv")
result = subprocess.run(
    ["./csvstat-ac", "--freq", "nosuch", str(asset)],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 3
assert result.stdout == ""
=== END AC cli-columns-missing-frequency ===

## User Acceptance

- None.

## Guardrails

- Null recognition occurs only after trimming.
- Column lookup uses header names from the examined file.
- No statistics or value lists are hardcoded.
