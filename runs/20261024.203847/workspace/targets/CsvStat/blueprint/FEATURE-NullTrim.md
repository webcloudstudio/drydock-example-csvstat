# FEATURE: Null and Trim

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Trim header names and values of surrounding spaces and classify null values. |
| Depends On  | FEATURE-CsvReading.md |
| Provides    | reader.Trim, reader.IsNull |

## Purpose

Every later rule operates on trimmed values and on the null classification defined here.

## Interface

```text
package reader

func TrimTable(t Table) Table   // trims every header name and every value
func Trim(s string) string      // removes leading and trailing spaces
func IsNull(s string) bool      // s is already trimmed
```

## Behavior

- `Trim` removes leading and trailing space characters (U+0020) only, and is applied to every header name and every value before any other rule.
- `IsNull` is true when the trimmed value is empty, or equals `NA`, `N/A`, or `NULL` ignoring case. Every other value is non-null.
- A value consisting only of spaces trims to empty and is therefore null.
- Distinctness elsewhere is compared on the trimmed string, so a padded value and its trimmed form are the same value.

## Test Names

The build agent writes these tests in `internal/reader` with exactly these names:

- `TestTrimSpaces`: leading and trailing spaces are removed and interior spaces are kept.
- `TestIsNullSpellings`: empty, NA, N/A, NULL in upper, lower, and mixed case are null; a near miss such as "NAN" or "N A" is not.
- `TestTrimBeforeNull`: a value of only spaces is null after trimming.
- `TestTrimTableHeaderAndRows`: TrimTable trims header names and every row value and preserves Lines.

## Programmatic Acceptance

=== AC nulltrim-trim-spaces ===
Intent: Trim removes surrounding spaces only.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/reader", "TestTrimSpaces")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC nulltrim-trim-spaces ===

=== AC nulltrim-null-spellings ===
Intent: Empty, NA, N/A, and NULL in any case are null and nothing else is.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/reader", "TestIsNullSpellings")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC nulltrim-null-spellings ===

=== AC nulltrim-trim-before-null ===
Intent: A value made only of spaces is null once trimmed.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/reader", "TestTrimBeforeNull")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC nulltrim-trim-before-null ===

=== AC nulltrim-trim-table ===
Intent: TrimTable trims header names and all row values and keeps line numbers.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/reader", "TestTrimTableHeaderAndRows")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC nulltrim-trim-table ===

## User Acceptance

- None.

## Guardrails

- Null spellings are exactly empty, `NA`, `N/A`, and `NULL` ignoring case; do not add others.
- Do not read any file under `sources/` from application code.
