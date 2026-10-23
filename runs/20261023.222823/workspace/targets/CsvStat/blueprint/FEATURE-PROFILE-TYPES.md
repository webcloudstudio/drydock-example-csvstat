# FEATURE: Profile Types

| Field       | Value |
|-------------|-------|
| Version     | 20261023 V1 |
| Description | Infers integer, number, date, boolean, text, and empty column types. |
| Depends On  | FEATURE-CLI-COLUMNS.md |
| Provides    | integer, number, date, boolean, text, and empty inference |
| Consumes    | column lookup, null normalization |

## Type Inference

For non-null values, classify using this precedence:

1. `integer`: `-?[0-9]+`
2. `number`: `-?[0-9]+(\.[0-9]+)?`
3. `date`: a real `YYYY-MM-DD` calendar date
4. `boolean`: `true` or `false`, ignoring case
5. `text`: any remaining value

A column with no non-null values is `empty`.

## Programmatic Acceptance

=== AC profile-types-fixtures ===
Intent: The supplied profile fixtures are accepted and produce profile rows.
from pathlib import Path
import subprocess

for name in ("employees.csv", "sensors.csv", "orders.csv", "header_only.csv"):
    result = subprocess.run(
        ["./csvstat-ac", str(Path("sources") / name)],
        capture_output=True,
        text=True,
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
    assert result.stdout != ""
=== END AC profile-types-fixtures ===

=== AC profile-types-all-null ===
Intent: An all-null column remains a valid profile input.
from pathlib import Path
import subprocess

result = subprocess.run(
    ["./csvstat-ac", "--column", "battery", "sources/sensors.csv"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
assert result.stdout != ""
=== END AC profile-types-all-null ===

=== AC profile-types-invalid-date-text ===
Intent: A column containing an invalid calendar date remains valid input rather than causing command failure.
import subprocess

result = subprocess.run(
    ["./csvstat-ac", "--column", "shipped", "sources/orders.csv"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
assert result.stdout != ""
=== END AC profile-types-invalid-date-text ===

## User Acceptance

- None.

## Guardrails

- Type inference ignores null values.
- Date validation rejects impossible calendar dates.
- Inference follows the specified precedence.
- No fixture-specific type or value is hardcoded.
