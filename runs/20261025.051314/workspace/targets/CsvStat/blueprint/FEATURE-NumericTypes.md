# FEATURE: Numeric Types

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infer integer and number column types from the non-null values of a column. |
| Depends On  | FEATURE-NullValues.md |
| Provides    | profile.InferType (integer, number) |
| Consumes    | csvdata.IsNull |

## Purpose

Column typing decides which statistics apply. This feature starts `profile.InferType` in
`internal/profile` with the two numeric types. Later features extend the same function with
`date`, `boolean`, `text`, and `empty`.

## Behavior

- `func InferType(values []string) string` in package `profile` receives the trimmed values of one
  column and returns the type name.
- Null values (per `csvdata.IsNull`) are ignored.
- A column is `integer` when every non-null value matches `-?[0-9]+`.
- Otherwise it is `number` when every non-null value matches `-?[0-9]+(\.[0-9]+)?`.
- Integer is tested before number, so a column of whole numbers is `integer`; a single decimal
  value makes the column `number`.
- Matching is anchored over the whole value and uses no float parsing. Values such as `+1`, `1.`,
  `.5`, `1e3`, `1,5`, `0x10`, and `-` are not numeric.
- When a column has at least one non-null value that is not numeric, this feature returns the
  fallback `text` until later features add the remaining types.
- A column whose non-null values are all null returns `empty`.

## Tests the build must write

In `internal/profile`:

- `TestInferIntegerColumn`
- `TestInferNumberColumn`
- `TestIntegerPreferredOverNumber`
- `TestNullsIgnoredWhenInferringNumeric`
- `TestMalformedNumericsAreNotNumeric`

## Programmatic Acceptance

=== AC numeric-integer ===
Intent: A column of integer-shaped values, including negatives, is typed integer.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferIntegerColumn")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC numeric-integer ===

=== AC numeric-number ===
Intent: A column with a decimal value among integers is typed number.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferNumberColumn")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC numeric-number ===

=== AC numeric-integer-preferred ===
Intent: Integer is chosen over number when both patterns match every value.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestIntegerPreferredOverNumber")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC numeric-integer-preferred ===

=== AC numeric-nulls-ignored ===
Intent: Null spellings among numeric values do not change the inferred type.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestNullsIgnoredWhenInferringNumeric")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC numeric-nulls-ignored ===

=== AC numeric-malformed ===
Intent: Values such as +1, 1., .5, 1e3 and 1,5 are not numeric.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMalformedNumericsAreNotNumeric")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC numeric-malformed ===

## User Acceptance

- None.

## Guardrails

- Classification never uses float parsing; values stay exact strings for later exact arithmetic.
- No type is hardcoded for any column or file; every type comes from the values examined.
- Only the standard library is used.
