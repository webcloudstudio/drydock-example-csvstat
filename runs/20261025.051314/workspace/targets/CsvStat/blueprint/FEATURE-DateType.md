# FEATURE: Date Type

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infer date columns whose non-null values are real YYYY-MM-DD calendar dates. |
| Depends On  | FEATURE-NumericTypes.md |
| Provides    | profile.InferType (date) |
| Consumes    | csvdata.IsNull, profile.InferType |

## Purpose

Extend `profile.InferType` so date columns are recognised and sort chronologically in later
statistics.

## Behavior

- After integer and number are ruled out, a column is `date` when every non-null value has the
  shape `YYYY-MM-DD` (four digits, hyphen, two digits, hyphen, two digits) and is a real calendar
  date.
- Real-date validation uses `time.Parse` with layout `2006-01-02` in addition to the digit-shape
  check, so `2025-02-30`, `2025-13-01`, and `2025-02-29` are not dates while `2024-02-29` is.
- Shapes such as `2025-2-3`, `2025/02/03`, `20250203`, and `2025-02-03T10:00` are not dates.
- Null values are ignored.
- One non-date value makes the column fall through to later types (`boolean`, then `text`).
  Integer and number take precedence, but no YYYY-MM-DD value can match them.

## Tests the build must write

In `internal/profile`:

- `TestInferDateColumn`
- `TestInvalidCalendarDateIsNotDate`
- `TestLeapDayHandling`
- `TestDateShapeIsStrict`
- `TestNullsIgnoredWhenInferringDate`

## Programmatic Acceptance

=== AC date-column ===
Intent: A column of valid YYYY-MM-DD values is typed date.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferDateColumn")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC date-column ===

=== AC date-invalid-calendar ===
Intent: A column containing an impossible date such as 2025-02-30 is not typed date.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInvalidCalendarDateIsNotDate")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC date-invalid-calendar ===

=== AC date-leap-day ===
Intent: A leap-year February 29 is a date and a non-leap-year February 29 is not.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestLeapDayHandling")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC date-leap-day ===

=== AC date-shape-strict ===
Intent: Values with the wrong shape (unpadded, slashes, timestamps) are not dates.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestDateShapeIsStrict")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC date-shape-strict ===

=== AC date-nulls-ignored ===
Intent: Null spellings among dates do not change the inferred date type.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestNullsIgnoredWhenInferringDate")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC date-nulls-ignored ===

## User Acceptance

- None.

## Guardrails

- Date validity is computed from the value, never from a list of known dates.
- Only the standard library (`time`, `regexp`) is used.
- The numeric rules from FEATURE-NumericTypes.md are not altered.
