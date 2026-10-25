# FEATURE: Sum and Mean

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Compute the exact sum and mean of integer and number columns. |
| Depends On  | FEATURE-BooleanTextTypes.md |
| Provides    | profile.SumMean |
| Consumes    | profile.InferType |

## Purpose

Numeric columns report `sum` and `mean`. Exactness matters because the output is rounded once, at
write time, from the exact value. This story computes the exact values; formatting belongs to the
decimal formatting story.

## Interface

```go
// SumMean takes the column type and the trimmed non-null values of the column. ok is false when
// the type is not integer or number, or when there are no values.
func SumMean(typ string, values []string) (sum, mean *big.Rat, ok bool)
```

## Behavior

- Only `integer` and `number` columns have a sum and mean. For every other type, including `empty`,
  `ok` is false.
- Parse every value into a `big.Rat` and accumulate exactly with `math/big`. No `float64` is used
  at any step, so `0.1` plus `0.2` is exactly three tenths.
- `sum` is the exact total of the non-null values.
- `mean` is the exact total divided by the non-null count. Null values are neither summed nor
  counted.
- The returned values are exact rationals. They are not rounded here.

## Required Tests

The project suite in `internal/profile` includes these tests by exact name:

- `TestSumIntegerExact` — an integer total, including negatives, equals the big-integer total.
- `TestSumNumberNoFloatDrift` — decimals such as `0.1` and `0.2` sum exactly to their rational total.
- `TestMeanExactRational` — total over count is an exact rational, for example a repeating third.
- `TestSumMeanIgnoresNulls` — the divisor is the non-null count, not the row count.
- `TestSumMeanNonNumeric` — text, date, boolean, and empty columns give `ok` false.

## Programmatic Acceptance

=== AC summean-integer-sum ===
Intent: Integer sums are exact, including negative values.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestSumIntegerExact")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC summean-integer-sum ===

=== AC summean-number-no-drift ===
Intent: Decimal sums carry no floating-point drift.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestSumNumberNoFloatDrift")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC summean-number-no-drift ===

=== AC summean-mean-exact ===
Intent: The mean is the exact total divided by the non-null count.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMeanExactRational")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC summean-mean-exact ===

=== AC summean-ignores-nulls ===
Intent: The mean divisor is the non-null count.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestSumMeanIgnoresNulls")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC summean-ignores-nulls ===

=== AC summean-non-numeric ===
Intent: Non-numeric column types have no sum or mean.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestSumMeanNonNumeric")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC summean-non-numeric ===

## User Acceptance

- None.

## Guardrails

- No `float64` arithmetic in sum or mean.
- Values are rounded once, at output time, never here.
- No total is embedded in code or read from `sources/`.
