# FEATURE: SumMean

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Compute the exact sum and mean of numeric columns, rounded once half away from zero. |
| Depends On  | FEATURE-TypeNumeric.md |
| Provides    | sum, mean |
| Consumes    | internal/profile.InferType |

## Purpose

Give `integer` and `number` columns an exact `sum` and `mean`, free of floating point error.

## Behavior

- Arithmetic uses `math/big` (`big.Rat`) over non-null values. Rounding happens once, when writing.
- `integer` sum is the integer total written as a plain integer.
- `number` sum is written with exactly two decimal places.
- `mean` for `integer` and `number` is the total divided by the non-null count, written with exactly two decimals.
- Rounding is half away from zero from the exact value. A negative result has a leading `-`; there is no thousands separator.
- A numeric column with no non-null values is typed `empty`, so it never reaches sum or mean.
- For every other type `sum` and `mean` are empty strings.

## Interface

- Package `internal/profile` exposes a formatter that renders a `big.Rat` with two decimals, rounded half away from zero, and computes sum and mean strings for a typed column.
- The project test suite in `internal/profile` carries tests named `TestSumInteger`, `TestSumNumber`, `TestMeanRounding`, `TestMeanNegative`, and `TestSumMeanNonNumeric`.

## Programmatic Acceptance

=== AC sum-integer ===
Intent: Integer sum is the exact total written as a plain integer.

from drydock_go import run_test

result = run_test("./internal/profile", "TestSumInteger")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC sum-integer ===

=== AC sum-number ===
Intent: Number sum is exact and written with two decimals.

from drydock_go import run_test

result = run_test("./internal/profile", "TestSumNumber")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC sum-number ===

=== AC mean-rounding ===
Intent: Mean is total over non-null count, rounded once half away from zero to two decimals.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMeanRounding")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC mean-rounding ===

=== AC mean-negative ===
Intent: A negative mean keeps its sign and rounds half away from zero.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMeanNegative")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC mean-negative ===

=== AC summean-non-numeric ===
Intent: Sum and mean are blank for date, boolean, text, and empty columns.

from drydock_go import run_test

result = run_test("./internal/profile", "TestSumMeanNonNumeric")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC summean-non-numeric ===

## User Acceptance

- None.

## Guardrails

- No floating point arithmetic is used for sums, means, or rounding.
- Rounding happens exactly once, at write time.
- No statistic is hardcoded; every value is computed from the file examined.
