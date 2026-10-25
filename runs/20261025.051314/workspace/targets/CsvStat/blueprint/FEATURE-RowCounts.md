# FEATURE: Row Counts

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Count data records, null values, and distinct non-null values for one column. |
| Depends On  | FEATURE-NullValues.md, FEATURE-InputRejection.md |
| Provides    | profile.Counts |
| Consumes    | csvdata.IsNull |

## Purpose

For each column the profile reports `rows`, `nulls`, and `distinct`. This story implements the
counting function in `internal/profile`. Rendering the counts into the profile CSV belongs to the
profile output story.

## Interface

```go
// Counts holds the three counts reported for one column.
type CountResult struct {
    Rows     int
    Nulls    int
    Distinct int
}

// Counts takes every trimmed value of one column, in file order, and returns its counts.
func Counts(values []string) CountResult
```

## Behavior

- `Rows` is the number of data records, that is `len(values)`. A header-only file gives zero values
  and therefore `Rows` 0.
- `Nulls` is the number of values for which `csvdata.IsNull` is true (empty, `NA`, `N/A`, `NULL`
  ignoring case).
- `Distinct` is the number of different non-null values, compared as exact strings after trimming.
  `Globex` and `globex` are different values. Null values are never counted as distinct.
- A column whose values are all null has `Distinct` 0.

## Required Tests

The project suite in `internal/profile` includes these tests by exact name, each written against
real output:

- `TestCountsRowsNullsDistinct` — mixed null and non-null values give the expected three counts.
- `TestCountsHeaderOnly` — no values give all three counts zero.
- `TestCountsDistinctExactStrings` — values differing only in case are distinct; repeated values
  count once.
- `TestCountsAllNull` — a column of only null spellings has `Distinct` 0 and `Nulls` equal to `Rows`.

## Programmatic Acceptance

=== AC rowcounts-rows-nulls-distinct ===
Intent: Counts reports rows, nulls, and distinct for a mixed column.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestCountsRowsNullsDistinct")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC rowcounts-rows-nulls-distinct ===

=== AC rowcounts-header-only ===
Intent: A column with no data records reports zero for every count.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestCountsHeaderOnly")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC rowcounts-header-only ===

=== AC rowcounts-distinct-exact ===
Intent: Distinct compares exact strings and counts repeats once.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestCountsDistinctExactStrings")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC rowcounts-distinct-exact ===

=== AC rowcounts-all-null ===
Intent: An all-null column has zero distinct values and every row null.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestCountsAllNull")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC rowcounts-all-null ===

## User Acceptance

- None.

## Guardrails

- Counts are computed from the values given; no count is embedded or read from `sources/`.
- Distinct never merges values by case folding or further trimming.
- Application code never reads `sources/run.sh`.
