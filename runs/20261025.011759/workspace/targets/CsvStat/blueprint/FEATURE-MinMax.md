# FEATURE: MinMax

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Report the minimum and maximum for every column type. |
| Depends On  | FEATURE-TypeDate.md, FEATURE-TypeOther.md |
| Provides    | min, max |
| Consumes    | internal/profile.InferType |

## Purpose

Give every column a `min` and `max` appropriate to its inferred type.

## Behavior

Statistics consider non-null values only.

| Type | min | max |
|---|---|---|
| `integer` | smallest value, written as a plain integer | largest value, written as a plain integer |
| `number` | smallest value, exactly two decimals | largest value, exactly two decimals |
| `date` | earliest date | latest date |
| `boolean` | `false` if any value is false, else `true` | `true` if any value is true, else `false` |
| `text` | smallest by byte value | largest by byte value |
| `empty` | empty string | empty string |

- Integer and number comparison is numeric and exact, using `math/big`. Values such as `-3` and `12` compare as numbers, not strings.
- Boolean values compare ignoring case, and the written result is lowercase.
- Number output has a leading `-` when negative, no thousands separator, and is rounded once half away from zero from the exact value.

## Interface

- Package `internal/profile` computes `min` and `max` strings for a column given its type and values.
- The project test suite in `internal/profile` carries tests named `TestMinMaxInteger`, `TestMinMaxNumber`, `TestMinMaxDate`, `TestMinMaxBoolean`, `TestMinMaxText`, and `TestMinMaxEmpty`.

## Programmatic Acceptance

=== AC minmax-integer ===
Intent: Integer min and max compare numerically and are written as plain integers.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxInteger")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-integer ===

=== AC minmax-number ===
Intent: Number min and max compare numerically and are written with two decimals.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxNumber")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-number ===

=== AC minmax-date ===
Intent: Date min and max give the earliest and latest date.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxDate")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-date ===

=== AC minmax-boolean ===
Intent: Boolean min is false if any value is false and max is true if any value is true, ignoring case.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxBoolean")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-boolean ===

=== AC minmax-text ===
Intent: Text min and max compare by byte value.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxText")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-text ===

=== AC minmax-empty ===
Intent: An empty column has blank min and max.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxEmpty")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-empty ===

## User Acceptance

- None.

## Guardrails

- Numeric comparison never uses string order or floating point.
- No min or max value is hardcoded; every value comes from the file examined.
