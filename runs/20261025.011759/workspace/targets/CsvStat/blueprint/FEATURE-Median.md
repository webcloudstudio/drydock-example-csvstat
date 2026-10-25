# FEATURE: Median

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Compute the exact median of numeric columns. |
| Depends On  | FEATURE-TypeNumeric.md |
| Provides    | median |
| Consumes    | internal/profile.InferType |

## Purpose

Give `integer` and `number` columns an exact `median`.

## Behavior

- Sort the non-null values exactly (`big.Rat`), then take the middle value; when the count is even, take the exact mean of the two middle values.
- The result is written with exactly two decimal places, rounded once half away from zero, with a leading `-` when negative and no thousands separator. This applies to `integer` columns too.
- For every non-numeric type `median` is an empty string.

## Interface

- Package `internal/profile` computes the median string for a typed column.
- The project test suite in `internal/profile` carries tests named `TestMedianOdd`, `TestMedianEven`, `TestMedianUnsorted`, `TestMedianNumber`, and `TestMedianNonNumeric`.

## Programmatic Acceptance

=== AC median-odd ===
Intent: With an odd count the median is the middle sorted value.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMedianOdd")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC median-odd ===

=== AC median-even ===
Intent: With an even count the median is the exact mean of the two middle values.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMedianEven")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC median-even ===

=== AC median-unsorted ===
Intent: Input order does not affect the median and the sort is numeric, not lexical.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMedianUnsorted")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC median-unsorted ===

=== AC median-number ===
Intent: A number median is written with two decimals rounded half away from zero.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMedianNumber")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC median-number ===

=== AC median-non-numeric ===
Intent: Median is blank for non-numeric columns.

from drydock_go import run_test

result = run_test("./internal/profile", "TestMedianNonNumeric")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC median-non-numeric ===

## User Acceptance

- None.

## Guardrails

- Sorting and averaging are exact; no floating point is used.
- Null values never take part in the median.
- No median value is hardcoded.
