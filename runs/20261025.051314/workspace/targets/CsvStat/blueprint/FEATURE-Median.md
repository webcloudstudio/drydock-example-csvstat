# FEATURE: Median

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Compute the exact median of integer and number columns. |
| Depends On  | FEATURE-SumMean.md |
| Provides    | profile.Median |
| Consumes    | profile.SumMean |

## Purpose

Numeric columns report a `median`. It is computed exactly and rounded once at output time by the
decimal formatting story.

## Interface

```go
// Median takes the column type and the trimmed non-null values of the column. ok is false when
// the type is not integer or number, or when there are no values.
func Median(typ string, values []string) (median *big.Rat, ok bool)
```

## Behavior

- Parse values into `big.Rat` and sort them numerically (`9` before `10`, `-3` before `2.5`),
  never lexically.
- With an odd count the median is the middle value.
- With an even count the median is the exact mean of the two middle values.
- Null values are excluded before counting.
- Non-numeric types and columns with no non-null value give `ok` false.
- The input slice is not reordered for the caller.

## Required Tests

The project suite in `internal/profile` includes these tests by exact name:

- `TestMedianOddCount` — the middle of an unsorted odd-length list.
- `TestMedianEvenCountExact` — the exact mean of the two middle values, including a case whose
  result is not representable in binary floating point.
- `TestMedianNumericOrder` — numeric rather than lexical ordering.
- `TestMedianIgnoresNulls` — nulls are excluded from the count.
- `TestMedianNonNumeric` — non-numeric and empty columns give `ok` false.

## Programmatic Acceptance

=== AC median-odd ===
Intent: An odd count yields the middle value of the sorted values.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMedianOddCount")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC median-odd ===

=== AC median-even ===
Intent: An even count yields the exact mean of the two middle values.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMedianEvenCountExact")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC median-even ===

=== AC median-numeric-order ===
Intent: Values are sorted numerically, not lexically.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMedianNumericOrder")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC median-numeric-order ===

=== AC median-ignores-nulls ===
Intent: Null values do not count toward the median.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMedianIgnoresNulls")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC median-ignores-nulls ===

=== AC median-non-numeric ===
Intent: Non-numeric and empty columns have no median.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMedianNonNumeric")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC median-non-numeric ===

## User Acceptance

- None.

## Guardrails

- No `float64` arithmetic in the median.
- The median is never rounded before output.
- No median is embedded in code or read from `sources/`.
