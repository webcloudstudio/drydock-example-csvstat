# FEATURE: Type Date

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infers the date column type with real calendar validation. |
| Depends On  | FEATURE-TypeNumeric.md |
| Provides    | date type |
| Consumes    | internal/profile.InferType |

## Purpose

A column is a `date` column when every non-null value is a real calendar date written `YYYY-MM-DD`. A value such as `2025-02-30` has the right shape but is not a date, so it keeps its column out of this type.

## Behavior

- `InferType` returns `"date"` when the column is not integer or number and every non-null value matches `^[0-9]{4}-[0-9]{2}-[0-9]{2}$` and parses with `time.Parse("2006-01-02", value)` without error.
- The shape check is strict: the value must be exactly ten characters in that form; a single-digit month or day, a time suffix such as `2025-03-01T08:00`, or a slash separator is not a date.
- Calendar validity includes month length and leap years: `2024-02-29` is a date, `2025-02-29` and `2025-04-31` are not.
- A single invalid date among valid ones sends the column to the later rungs (boolean, text), so it becomes `text`.
- Null values are ignored.
- A column that is entirely integers or numbers keeps that earlier type.

## Required project tests

The builder writes these tests in `internal/profile` with exactly these names:

- `TestInferTypeDate`: valid dates, leap day, and dates mixed with nulls are typed date.
- `TestInferTypeInvalidDate`: impossible calendar dates, wrong shapes, and time suffixes make the column text.

## Programmatic Acceptance

=== AC typedate-valid ===
Intent: Columns of real calendar dates, including a leap day and nulls, are typed date.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferTypeDate")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC typedate-valid ===

=== AC typedate-invalid ===
Intent: One impossible date or malformed shape keeps a column from being typed date.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferTypeInvalidDate")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC typedate-invalid ===

## User Acceptance

- None.

## Guardrails

- Calendar validity is decided by the standard library, not by a hand-written month table.
- Typing never reads a data file and never hardcodes a value.
- Standard library only.
