# FEATURE: Min and Max

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Compute the minimum and maximum of a column according to its inferred type. |
| Depends On  | FEATURE-BooleanTextTypes.md |
| Provides    | profile.MinMax |
| Consumes    | profile.InferType |

## Purpose

The profile reports `min` and `max` for every column. The rule depends on the column type. This
story implements the extremes computation in `internal/profile`; two-decimal formatting of
`number` results belongs to the decimal formatting story.

## Interface

```go
// Extremes holds the minimum and maximum of one column.
type Extremes struct {
    Min, Max       string   // text form, see below
    MinRat, MaxRat *big.Rat // exact values for integer and number columns, else nil
}

// MinMax takes the column type and the trimmed non-null values of the column.
func MinMax(typ string, values []string) Extremes
```

`typ` is one of `integer`, `number`, `date`, `boolean`, `text`, `empty`.

## Behavior

- `integer`: compare exactly with `math/big`, never as `float64`. `Min` and `Max` are the smallest
  and largest values written as plain integers. `MinRat` and `MaxRat` are set.
- `number`: compare exactly as `big.Rat` (so `2.5` is less than `10.25`, and `-0.5` less than
  `-0.25`). `MinRat` and `MaxRat` are set; `Min` and `Max` hold the source text of the extreme
  values, and the decimal formatting story renders them with two places.
- `date`: `Min` is the earliest and `Max` the latest. `YYYY-MM-DD` text orders chronologically.
- `boolean`: `Min` is `false` if any value is false (ignoring case), else `true`. `Max` is `true`
  if any value is true (ignoring case), else `false`. Output is lowercase whatever the input case.
- `text`: smallest and largest by byte value, not locale collation. Upper case sorts before lower
  case; a space sorts before letters.
- `empty`: `Min` and `Max` are blank and the `Rat` fields are nil.
- A column with exactly one non-null value reports that value as both minimum and maximum.

## Required Tests

The project suite in `internal/profile` includes these tests by exact name:

- `TestMinMaxInteger` — signed integers, including a negative, are ordered numerically, not
  lexically (`9` is less than `10`).
- `TestMinMaxNumber` — decimals are ordered exactly, including negatives and differing digit counts.
- `TestMinMaxDate` — earliest and latest dates.
- `TestMinMaxBoolean` — mixed case input yields lowercase output; all-true gives `true` for both;
  all-false gives `false` for both; mixed gives `false` and `true`.
- `TestMinMaxText` — byte-value ordering, including mixed case.
- `TestMinMaxEmpty` — an empty column yields blank extremes.
- `TestMinMaxSingleValue` — one value is both minimum and maximum.

## Programmatic Acceptance

=== AC minmax-integer ===
Intent: Integer extremes are ordered numerically with exact arithmetic.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxInteger")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-integer ===

=== AC minmax-number ===
Intent: Number extremes are ordered exactly as decimals.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxNumber")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-number ===

=== AC minmax-date ===
Intent: Date extremes are the earliest and latest dates.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxDate")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-date ===

=== AC minmax-boolean ===
Intent: Boolean extremes follow the false/true rule and are lowercase regardless of input case.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxBoolean")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-boolean ===

=== AC minmax-text ===
Intent: Text extremes are ordered by byte value.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxText")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-text ===

=== AC minmax-empty ===
Intent: An empty column has blank minimum and maximum.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxEmpty")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-empty ===

=== AC minmax-single-value ===
Intent: A single non-null value is both the minimum and the maximum.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestMinMaxSingleValue")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC minmax-single-value ===

## User Acceptance

- None.

## Guardrails

- Numeric comparison never uses `float64`.
- Text comparison is by byte value only.
- No extreme is embedded in code or read from `sources/`; every value comes from the input given.
EOF_NOT_USED
