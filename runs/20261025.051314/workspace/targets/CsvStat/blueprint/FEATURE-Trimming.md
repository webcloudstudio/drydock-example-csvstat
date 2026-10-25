# FEATURE: Trimming

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Trim leading and trailing spaces from every header name and every value before any other rule applies. |
| Depends On  | FEATURE-CsvReading.md |
| Provides    | csvdata.Trim |
| Consumes    | csvdata.Read |

## Purpose

Every header name and every value is trimmed of leading and trailing spaces before null recognition,
typing, counting, or comparison. Padded text such as `  Globex  ` and `Acme Corp ` therefore count as
`Globex` and `Acme Corp`.

## Behavior

- `Trim(t *Table)` replaces, in place, each header name and each record field with the same string
  stripped of leading and trailing space characters (U+0020).
- Only U+0020 is trimmed. Tabs and other whitespace are preserved. Interior spaces are preserved.
- Trimming applies after CSV parsing, so spaces inside a quoted field are also trimmed.
- Record line numbers and field counts are unchanged.
- The cli pipeline calls `Read` then `Trim` before any other rule; the profile-output and
  frequency-mode stories wire that call.

## Tests For The Project Suite

The builder writes these `internal/csvdata` tests, named exactly as listed:

- `TestTrimHeaderNames`: padded header names are trimmed.
- `TestTrimValues`: padded values are trimmed on both sides.
- `TestTrimQuotedPadding`: spaces inside a quoted field are trimmed, and the comma inside the quotes is kept.
- `TestTrimOnlySpaces`: tabs are preserved and interior spaces are kept.
- `TestTrimKeepsLines`: line numbers and field counts are unchanged.

## Programmatic Acceptance

=== AC trim-header-names ===
Intent: Leading and trailing spaces are removed from header names.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

result = run_test("./internal/csvdata", "TestTrimHeaderNames")
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC trim-header-names ===

=== AC trim-values ===
Intent: Leading and trailing spaces are removed from every value.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

result = run_test("./internal/csvdata", "TestTrimValues")
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC trim-values ===

=== AC trim-quoted-padding ===
Intent: Padding inside a quoted field is trimmed while the quoted comma is kept.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

result = run_test("./internal/csvdata", "TestTrimQuotedPadding")
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC trim-quoted-padding ===

=== AC trim-scope ===
Intent: Only spaces are trimmed, interior spaces remain, and line numbers and field counts are unchanged.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

for name in ("TestTrimOnlySpaces", "TestTrimKeepsLines"):
    result = run_test("./internal/csvdata", name)
    print(name, result.status)
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC trim-scope ===

## User Acceptance

- None.

## Guardrails

- Trimming happens before null recognition, typing, counting, and comparison.
- Trim leading and trailing spaces only; do not alter interior content.
- Distinctness is compared as exact strings after trimming.
