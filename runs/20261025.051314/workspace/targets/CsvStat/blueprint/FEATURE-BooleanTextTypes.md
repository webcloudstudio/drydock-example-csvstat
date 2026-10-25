# FEATURE: Boolean and Text Types

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infer boolean, text, and empty column types, completing the type order. |
| Depends On  | FEATURE-DateType.md |
| Provides    | profile.InferType (boolean, text, empty) |
| Consumes    | csvdata.IsNull, profile.InferType |

## Purpose

Complete `profile.InferType` so every column receives exactly one type: the first of `integer`,
`number`, `date`, `boolean`, `text` that every non-null value satisfies, or `empty` when there is
no non-null value.

## Behavior

- A column is `boolean` when every non-null value equals `true` or `false` ignoring case
  (`True`, `TRUE`, `false`, `False`). Values such as `yes`, `1`, `0`, and `t` are not boolean.
- `text` is the fallback for any column that satisfies none of the earlier types.
- A column with no non-null value (every value null, or zero data rows) is `empty`.
- The order of evaluation is: empty check, integer, number, date, boolean, text.
- A column of `0` and `1` is `integer`, not boolean.
- Original value case is preserved for distinct counting; only classification ignores case. The
  min and max statistics later render booleans as lowercase `true` or `false` regardless of input
  case; that rendering belongs to the min/max feature.
- Null values are ignored during classification.

## Tests the build must write

In `internal/profile`:

- `TestInferBooleanColumnMixedCase`
- `TestInferTextFallback`
- `TestInferEmptyWhenAllNull`
- `TestInferEmptyWhenNoRows`
- `TestBooleanLookalikesAreText`
- `TestTypeOrderPrecedence`

## Programmatic Acceptance

=== AC type-boolean-mixed-case ===
Intent: Values true and false in any letter case form a boolean column.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferBooleanColumnMixedCase")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC type-boolean-mixed-case ===

=== AC type-text-fallback ===
Intent: A column matching no earlier type is text.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferTextFallback")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC type-text-fallback ===

=== AC type-empty-all-null ===
Intent: A column whose values are all null spellings is typed empty.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferEmptyWhenAllNull")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC type-empty-all-null ===

=== AC type-empty-no-rows ===
Intent: A column with zero data rows is typed empty.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestInferEmptyWhenNoRows")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC type-empty-no-rows ===

=== AC type-boolean-lookalikes ===
Intent: Values such as yes, 1 or t do not make a boolean column.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestBooleanLookalikesAreText")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC type-boolean-lookalikes ===

=== AC type-order-precedence ===
Intent: Earlier types win over later ones when several would match.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/profile", "TestTypeOrderPrecedence")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC type-order-precedence ===

## User Acceptance

- None.

## Guardrails

- No column name, file name, or value list is hardcoded to decide a type.
- Booleans are classified case-insensitively but the stored value is never rewritten here.
- Only the standard library is used.
