# FEATURE: Decimal Format

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Format exact rational results as two-decimal strings rounded half away from zero, and integers as plain integers. |
| Depends On  | FEATURE-MinMax.md, FEATURE-SumMean.md, FEATURE-Median.md |
| Provides    | profile.FormatDecimal, profile.FormatInteger |
| Consumes    | profile.MinMax, profile.SumMean, profile.Median |

## Purpose

Exact values are rounded once, when written. This story provides the formatting functions the
profile output story applies to each statistic.

## Interface

```go
// FormatDecimal writes r with exactly two decimal places.
func FormatDecimal(r *big.Rat) string

// FormatInteger writes an integer-valued r as a plain integer.
func FormatInteger(r *big.Rat) string
```

## Behavior

- `FormatDecimal` rounds the exact value to two places, half away from zero: an exact tie such as
  one thousandth short of a half cent rounds by magnitude upward for both signs (`0.005` becomes
  `0.01`, `-0.005` becomes `-0.01`). Rounding is done on the exact rational with integer
  arithmetic, never through `float64` or `fmt` rounding of a float.
- Exactly two decimal places are always written, so a whole value has `.00`.
- A leading `-` is written only when the rounded result is non-zero. A value that rounds to zero
  is written `0.00` with no minus sign.
- No thousands separator is written.
- Very large magnitudes keep every digit.
- `FormatInteger` writes the integer with no decimal point and no separator, with `-` for negatives
  and no negative zero.
- Which statistics use which function: `number` min, max, sum, mean, and median, and `integer`
  mean and median use `FormatDecimal`; `integer` min, max, and sum use `FormatInteger`. `date`,
  `boolean`, `text` extremes are written as given.

## Required Tests

The project suite in `internal/profile` includes these tests by exact name:

- `TestFormatDecimalTwoPlaces` — whole, one-place, and many-place values all gain exactly two places.
- `TestFormatDecimalHalfAwayFromZero` — positive and negative exact ties round away from zero.
- `TestFormatDecimalNoNegativeZero` — a small negative value that rounds to zero is `0.00`.
- `TestFormatDecimalNoSeparator` — a value above one thousand has no thousands separator.
- `TestFormatDecimalLargeExact` — a magnitude beyond `float64` precision keeps all digits.
- `TestFormatInteger` — plain integers, negatives, and zero.

## Programmatic Acceptance

=== AC decfmt-two-places ===
Intent: FormatDecimal always writes exactly two decimal places.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestFormatDecimalTwoPlaces")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC decfmt-two-places ===

=== AC decfmt-half-away ===
Intent: Exact ties round half away from zero for both signs.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestFormatDecimalHalfAwayFromZero")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC decfmt-half-away ===

=== AC decfmt-no-negative-zero ===
Intent: A negative value that rounds to zero is written without a minus sign.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestFormatDecimalNoNegativeZero")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC decfmt-no-negative-zero ===

=== AC decfmt-no-separator ===
Intent: Large values are written without a thousands separator.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestFormatDecimalNoSeparator")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC decfmt-no-separator ===

=== AC decfmt-large-exact ===
Intent: Values beyond float64 precision keep every digit.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestFormatDecimalLargeExact")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC decfmt-large-exact ===

=== AC decfmt-integer ===
Intent: FormatInteger writes plain integers with no decimal point.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestFormatInteger")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC decfmt-integer ===

## User Acceptance

- None.

## Guardrails

- Rounding happens once, on the exact rational, never through `float64`.
- No negative zero in any output.
- No thousands separator.
- No formatted value is embedded in code or read from `sources/`.
