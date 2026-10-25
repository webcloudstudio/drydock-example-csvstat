# FEATURE: CSV Reading

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Read UTF-8 RFC 4180 CSV with the first record as the header, recording each record's source line. |
| Depends On  | ARCHITECTURE.md |
| Provides    | csvdata.Read |
| Consumes    | None |

## Purpose

`internal/csvdata` turns an input stream into a header and data records. Trimming, null recognition, and
rejection of malformed input are separate stories that build on this reader.

## Behavior

- `Read(r io.Reader) (*Table, error)` parses with `encoding/csv`, with `FieldsPerRecord = -1` so field
  counts are judged by the rejection story, and `LazyQuotes` off.
- `Table` has `Header []string` and `Records []Record`. `Record` has `Line int`, the 1-based source line
  on which the record starts, and `Fields []string`.
- The first record is the header; every later record is a data record. Values are returned as parsed
  with no trimming.
- A quoted field containing a comma is one value. A quoted field may contain doubled quotes and newlines.
- A line ending after the last record creates no extra record. Input with or without a final newline
  yields the same records.
- Input with no record at all (empty or only blank lines) returns a `Table` with a nil `Header` and no
  `Records` and no error; the input-rejection story turns that into a rejection.
- A syntax error reported by `encoding/csv` is returned wrapped so callers can recover the line number
  with `errors.As` and `*csv.ParseError`.
- `Read` does not open files; `cli` opens the file and passes the reader.

## Tests For The Project Suite

The builder writes these `internal/csvdata` tests, named exactly as listed:

- `TestReadHeaderAndRecords`: header is the first record; the rest are data records.
- `TestReadQuotedComma`: a quoted field with a comma is a single value.
- `TestReadTrailingNewline`: a final newline adds no record, matching the same input without one.
- `TestReadRecordLineNumbers`: each record carries the line on which it starts, including after a
  quoted field that spans lines.
- `TestReadHeaderOnly`: a header with no data yields zero records.
- `TestReadNoRecords`: empty input yields a nil header and no records.

## Programmatic Acceptance

=== AC read-header-and-records ===
Intent: The first record is the header and later records are data records.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

result = run_test("./internal/csvdata", "TestReadHeaderAndRecords")
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC read-header-and-records ===

=== AC read-quoted-comma ===
Intent: A quoted field containing a comma parses as one value.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

result = run_test("./internal/csvdata", "TestReadQuotedComma")
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC read-quoted-comma ===

=== AC read-trailing-newline ===
Intent: A line ending after the last record does not create an extra data record.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

result = run_test("./internal/csvdata", "TestReadTrailingNewline")
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC read-trailing-newline ===

=== AC read-line-numbers ===
Intent: Each record carries the source line on which it starts.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

result = run_test("./internal/csvdata", "TestReadRecordLineNumbers")
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC read-line-numbers ===

=== AC read-header-only-and-empty ===
Intent: A header-only input has zero records and an input with no record has no header.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

for name in ("TestReadHeaderOnly", "TestReadNoRecords"):
    result = run_test("./internal/csvdata", name)
    print(name, result.status)
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC read-header-only-and-empty ===

## User Acceptance

- None.

## Guardrails

- Read with Go's `encoding/csv`; do not write a custom CSV parser.
- The reader does not trim, classify, or reject values; those belong to later stories.
- Never read a supplied data file from application code; the file named on the command line is the only input.
