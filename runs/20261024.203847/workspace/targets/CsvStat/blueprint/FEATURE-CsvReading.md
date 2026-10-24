# FEATURE: CSV Reading

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Read a UTF-8 RFC 4180 CSV stream with encoding/csv into a header, data records, and record line numbers. |
| Depends On  | ARCHITECTURE.md |
| Provides    | reader.Read |

## Purpose

`internal/reader` parses the input file into a table that later stories trim, validate, type, and profile.

## Interface

```text
package reader

type Table struct {
    Header []string   // first record
    Rows   [][]string // later records, in file order
    Lines  []int      // 1-based input line on which each row starts; len(Lines) == len(Rows)
}

func Read(r io.Reader) (Table, error)
```

## Behavior

- Parsing uses `encoding/csv` with `FieldsPerRecord = -1`, so a record whose field count differs from the header is returned as-is. Ragged detection belongs to the reject-malformed story.
- The first record is the header; every later record is a data row. Line numbers come from the csv reader's field positions.
- A quoted field containing a comma, a quote, or a newline stays inside one value.
- An empty input returns a zero Table and a nil error; the caller rejects it.
- Bytes that are not valid UTF-8 are not a rejection condition; they are read as `encoding/csv` reads them.
- A syntax error reported by `encoding/csv` is returned unchanged; the caller chooses the exit status.
- Values are returned untrimmed here; trimming is the null-trim story.

## Test Names

The build agent writes these tests in `internal/reader` with exactly these names:

- `TestReadHeaderThenRows`: first record is the header and later records are rows, in order.
- `TestReadQuotedComma`: a quoted field containing a comma is one value.
- `TestReadInvalidUTF8NotRejected`: a field containing a byte that is not valid UTF-8 is read with a nil error.
- `TestReadRaggedRecordReturned`: a record with an extra field is returned unchanged with a nil error.
- `TestReadLineNumbers`: Lines holds the starting line of each row, counting a multi-line quoted field correctly.
- `TestReadEmptyInput`: empty input returns an empty Table and a nil error.

## Programmatic Acceptance

=== AC csvreading-header-and-rows ===
Intent: The first record is the header and later records are data rows in order.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/reader", "TestReadHeaderThenRows")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC csvreading-header-and-rows ===

=== AC csvreading-quoted-comma ===
Intent: A quoted comma stays inside a single value.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/reader", "TestReadQuotedComma")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC csvreading-quoted-comma ===

=== AC csvreading-invalid-utf8-accepted ===
Intent: Invalid UTF-8 bytes do not cause rejection.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/reader", "TestReadInvalidUTF8NotRejected")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC csvreading-invalid-utf8-accepted ===

=== AC csvreading-ragged-returned ===
Intent: A record with a different field count is returned for later rejection rather than failing the read.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/reader", "TestReadRaggedRecordReturned")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC csvreading-ragged-returned ===

=== AC csvreading-line-numbers ===
Intent: Each row carries the input line on which it starts.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/reader", "TestReadLineNumbers")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC csvreading-line-numbers ===

=== AC csvreading-empty-input ===
Intent: Empty input yields an empty table and no error.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/reader", "TestReadEmptyInput")
print(result.stdout)
print(result.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC csvreading-empty-input ===

## User Acceptance

- None.

## Guardrails

- Use only `encoding/csv` for parsing; no hand-written CSV parser.
- Do not read any file under `sources/` from application code.
- Do not reject a file for containing bytes that are not valid UTF-8.
