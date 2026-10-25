# FEATURE: Type Numeric

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infers the integer and number column types from a column's non-null values. |
| Depends On  | FEATURE-NullValues.md |
| Provides    | internal/profile.InferType |
| Consumes    | internal/profile.IsNull |

## Purpose

A column's type is the first of `integer`, `number`, `date`, `boolean`, `text` that every non-null value satisfies, or `empty` when there is no non-null value. This feature creates the typing function and implements its first two rungs; later features add the remaining rungs to the same function.

## Behavior

- Package `internal/profile` exports `InferType(values []string) string`. Values are already trimmed; null values are skipped using `IsNull`.
- Returns `"integer"` when every non-null value matches `^-?[0-9]+$`.
- Returns `"number"` when every non-null value matches `^-?[0-9]+(\.[0-9]+)?$` and at least one value is not an integer.
- Integer text with leading zeros or a leading `-` and digits is still an integer; `+5`, `.5`, `5.`, `1e3`, `1,000`, and a bare `-` are not numeric.
- Typing considers only non-null values: a column of integers and nulls is `integer`.
- The ladder order is integer, number, date, boolean, text; a column that satisfies an earlier rung never takes a later one.
- Non-numeric columns fall through to later rungs, implemented by FEATURE-TypeDate.md and FEATURE-TypeOther.md. Until they exist, a non-numeric column returns `"text"`.

## Required project tests

The builder writes these tests in `internal/profile` with exactly these names:

- `TestInferTypeInteger`: integer columns, negative values, leading zeros, and integers mixed with nulls.
- `TestInferTypeNumber`: decimals, negative decimals, decimals mixed with integers, and decimals mixed with nulls.
- `TestInferTypeNotNumeric`: forms such as `+5`, `.5`, `5.`, `1e3`, `1,000`, and `-` are not integer or number.

## Programmatic Acceptance

=== AC typenumeric-integer ===
Intent: Integer columns, including negatives and nulls, are typed integer by the delivered function.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferTypeInteger")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC typenumeric-integer ===

=== AC typenumeric-number ===
Intent: Columns with at least one decimal value and otherwise numeric values are typed number.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferTypeNumber")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC typenumeric-number ===

=== AC typenumeric-not-numeric ===
Intent: Malformed numeric spellings never type a column as integer or number.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferTypeNotNumeric")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC typenumeric-not-numeric ===

## User Acceptance

- None.

## Guardrails

- Type names are exactly `integer`, `number`, `date`, `boolean`, `text`, and `empty`.
- Typing never reads a data file and never hardcodes a column name or a type for a column.
- Standard library only.
