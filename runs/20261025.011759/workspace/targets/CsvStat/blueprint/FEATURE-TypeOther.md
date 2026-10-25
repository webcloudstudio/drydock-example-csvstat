# FEATURE: Type Other

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infers the boolean, text, and empty column types. |
| Depends On  | FEATURE-TypeNumeric.md |
| Provides    | boolean, text, empty types |
| Consumes    | internal/profile.InferType |

## Purpose

This feature completes the typing ladder: `boolean` for true/false columns, `text` as the catch-all, and `empty` for columns with no non-null value.

## Behavior

- `InferType` returns `"empty"` when no value in the column is non-null. This includes a column with no rows at all, a column of only empty strings, and a column of only `NA`, `N/A`, or `NULL` spellings.
- `InferType` returns `"boolean"` when the column is not integer, number, or date, and every non-null value equals `true` or `false` ignoring case (`TRUE`, `False`).
- `InferType` returns `"text"` for every other column, including a column mixing booleans with other words, or booleans with numbers.
- Null values are ignored for every rung; a column of booleans and nulls is `boolean`.
- `yes`, `no`, `1`, `0`, `t`, `f` are not booleans.
- `empty` takes precedence over every other type because it is decided first, by the absence of non-null values.

## Required project tests

The builder writes these tests in `internal/profile` with exactly these names:

- `TestInferTypeBoolean`: mixed-case true/false columns, with and without nulls.
- `TestInferTypeText`: ordinary words, mixed booleans and words, padded and punctuated strings.
- `TestInferTypeEmpty`: no values, only empty strings, only null spellings.

## Programmatic Acceptance

=== AC typeother-boolean ===
Intent: Columns whose non-null values are all true or false in any case are typed boolean.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferTypeBoolean")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC typeother-boolean ===

=== AC typeother-text ===
Intent: Columns that satisfy no earlier rung are typed text.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferTypeText")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC typeother-text ===

=== AC typeother-empty ===
Intent: Columns with no non-null value are typed empty.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferTypeEmpty")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC typeother-empty ===

## User Acceptance

- None.

## Guardrails

- Boolean spellings are exactly `true` and `false` ignoring case; no other spelling is added.
- Typing never reads a data file and never hardcodes a column name or a type for a column.
- Standard library only.
