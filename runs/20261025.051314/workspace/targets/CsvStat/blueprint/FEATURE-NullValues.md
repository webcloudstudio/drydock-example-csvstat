# FEATURE: Null Values

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Recognise null spellings (empty, NA, N/A, NULL ignoring case) on trimmed values. |
| Depends On  | FEATURE-Trimming.md, FEATURE-CsvReading.md |
| Provides    | csvdata.IsNull |
| Consumes    | csvdata.Read |

## Purpose

Statistics, type inference and frequency counting all need one shared definition of "null". This
feature provides it as a function in `internal/csvdata` so no other package re-implements it.

## Behavior

- `func IsNull(value string) bool` in package `csvdata` receives a value that has already been
  trimmed of leading and trailing spaces.
- A value is null when it is the empty string, or equals `NA`, `N/A`, or `NULL` ignoring case
  (`na`, `n/a`, `Null` and `null` are null).
- Every other value is non-null, including `0`, `false`, `NAN`, `N/B`, `NULLS` and `NA ` is not
  an input (trimming precedes this rule), so only exact spellings match.
- A quoted empty field (`""`) is read by `encoding/csv` as an empty string, so it is null.
- A field made only of spaces is empty after trimming and is therefore null.
- `IsNull` is pure: it never reads files and never allocates state.

## Tests the build must write

The project's own suite in `internal/csvdata` contains these test functions (names are fixed so
acceptance can run them):

- `TestIsNullEmpty`
- `TestIsNullSpellings`
- `TestIsNullNonNullValues`
- `TestQuotedEmptyFieldIsNull`
- `TestSpacesOnlyFieldIsNull`

## Programmatic Acceptance

=== AC null-empty ===
Intent: The empty string is reported null by csvdata.IsNull.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvdata", "TestIsNullEmpty")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC null-empty ===

=== AC null-spellings ===
Intent: NA, N/A and NULL are null in any letter case.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvdata", "TestIsNullSpellings")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC null-spellings ===

=== AC null-non-null-values ===
Intent: Ordinary values such as 0, false and near-miss spellings are not null.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvdata", "TestIsNullNonNullValues")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC null-non-null-values ===

=== AC null-quoted-empty ===
Intent: A quoted empty field read from a CSV file is null.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvdata", "TestQuotedEmptyFieldIsNull")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC null-quoted-empty ===

=== AC null-spaces-only ===
Intent: A field of only spaces is null once trimmed.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvdata", "TestSpacesOnlyFieldIsNull")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC null-spaces-only ===

## User Acceptance

- None.

## Guardrails

- `IsNull` never reads `sources/run.sh` or any supplied data file and embeds no answer.
- Only the standard library is used.
- Null spellings are exactly those named in the source; no other token is treated as null.
