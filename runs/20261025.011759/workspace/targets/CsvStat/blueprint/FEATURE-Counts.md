# FEATURE: Counts

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Count rows, nulls, and distinct non-null values for each column. |
| Depends On  | FEATURE-NullValues.md, FEATURE-TypeNumeric.md |
| Provides    | rows, nulls, distinct |
| Consumes    | internal/profile.IsNull |

## Purpose

Give every column three counts used by the profile table. The counts are computed from the file examined and never hardcoded.

## Behavior

- `rows` is the number of data records, that is, every record after the header. It is the same for every column.
- `nulls` is the number of values in the column for which `internal/profile.IsNull` is true.
- `distinct` is the number of different non-null values, compared as exact trimmed strings. Case is significant: `true` and `TRUE` are two values. Null values are never counted.
- A file with a header and no records yields `rows` 0, `nulls` 0, and `distinct` 0 for every column.

## Interface

- Package `internal/profile` exposes a function that takes a column's trimmed values and returns the three counts.
- The project test suite in `internal/profile` carries tests named `TestCountsRows`, `TestCountsNulls`, `TestCountsDistinct`, and `TestCountsHeaderOnly`, written by the implementer against real output.

## Programmatic Acceptance

=== AC counts-rows ===
Intent: The row count equals the number of data records, including a header-only input.

from drydock_go import run_test

result = run_test("./internal/profile", "TestCountsRows")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC counts-rows ===

=== AC counts-nulls ===
Intent: Empty, NA, N/A, and NULL values are counted as nulls and nothing else is.

from drydock_go import run_test

result = run_test("./internal/profile", "TestCountsNulls")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC counts-nulls ===

=== AC counts-distinct ===
Intent: Distinct counts different non-null trimmed strings with case significant and nulls excluded.

from drydock_go import run_test

result = run_test("./internal/profile", "TestCountsDistinct")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC counts-distinct ===

=== AC counts-header-only ===
Intent: A header-only input yields zero for all three counts.

from drydock_go import run_test

result = run_test("./internal/profile", "TestCountsHeaderOnly")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC counts-header-only ===

## User Acceptance

- None.

## Guardrails

- Counts are computed from the file examined; no count is hardcoded.
- Distinct comparison uses exact trimmed strings, never case folding or numeric equality.
- Null values are never counted as distinct values.
