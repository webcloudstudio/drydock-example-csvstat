# FEATURE: Null Values

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Classifies empty, NA, N/A, and NULL values as null after trimming. |
| Depends On  | FEATURE-CsvReader.md |
| Provides    | internal/profile.IsNull |
| Consumes    | internal/csvio.Read |

## Purpose

Every statistic, type, and frequency rule applies only to non-null values. This feature supplies the single predicate that decides whether a value is null, so counts, typing, statistics, and frequency lists agree.

## Behavior

- Package `internal/profile` exports `IsNull(value string) bool`.
- The value is trimmed of leading and trailing spaces before the test, so the predicate gives the same answer for a padded value and its trimmed form.
- A value is null when the trimmed value is empty, or equals `NA`, `N/A`, or `NULL` ignoring case (`na`, `n/a`, `Null`, `null` are null).
- Every other value is non-null, including near misses such as `NAN`, `N/B`, `NULLS`, `none`, and `0`.
- The predicate is pure: no I/O, no package state.

## Required project tests

The builder writes these tests in `internal/profile` beside the code, with exactly these names:

- `TestIsNullEmpty`: the empty string and whitespace-only strings are null.
- `TestIsNullSpellings`: `NA`, `N/A`, `NULL` in upper, lower, and mixed case are null, including when padded with spaces.
- `TestIsNullNonNull`: near misses and ordinary values are not null.

## Programmatic Acceptance

=== AC nullvalues-empty ===
Intent: Empty and whitespace-only values are classified as null by the delivered predicate.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestIsNullEmpty")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC nullvalues-empty ===

=== AC nullvalues-spellings ===
Intent: NA, N/A, and NULL are null in any letter case, padded or not.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestIsNullSpellings")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC nullvalues-spellings ===

=== AC nullvalues-nonnull ===
Intent: Near misses and ordinary values are not null.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestIsNullNonNull")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC nullvalues-nonnull ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no third-party modules.
- The predicate never reads a data file and never embeds a value taken from a supplied data file.
- Null spellings are exactly the empty string, `NA`, `N/A`, and `NULL` ignoring case; no other spelling is added.
