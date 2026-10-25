# FEATURE: CSV Reader

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Read a CSV file with encoding/csv, taking the first record as the header and trimming every header name and value. |
| Depends On  | ARCHITECTURE.md |
| Provides    | internal/csvio.Read |
| Consumes    | None |

## Purpose

`internal/csvio` reads the file as RFC 4180 CSV with `encoding/csv`. The first record is the header. Every header name and every value is trimmed of leading and trailing spaces before any other rule applies. Structural rejection (empty file, bad header, ragged record) belongs to the InputRejection feature; this feature supplies the parsed, trimmed records and their line numbers.

## Interface

`csvio.Read(path string) (*Table, error)` where `Table` carries:

- `Header []string`: trimmed header names.
- `Rows [][]string`: trimmed data records in file order.
- `Lines []int`: the 1-based starting line of each data record (from `csv.Reader.FieldPos`), kept for diagnostics.

The reader sets `FieldsPerRecord = -1` so that a field-count mismatch is detected by the application and not by `encoding/csv`, which lets the caller name the offending line. `LazyQuotes` stays false. A quoted field containing a comma, a quote, or a newline is returned as one field.

## Behavior

- Trim spaces from header names and values, for example a value `  Globex  ` becomes `Globex`.
- A quoted field that contains a comma stays one field.
- Blank lines are skipped by `encoding/csv`.
- Invalid UTF-8 bytes are read and not rejected.
- Failure to open or read the file is returned as an error distinguishable from a CSV syntax error, so the caller can map it to exit 1 or exit 3.

## Required Project Tests

The build writes these tests in `internal/csvio` using files the tests create. Acceptance runs them by exact name.

- `TestReadTrimsHeaderAndValues`: padded header names and padded values come back trimmed.
- `TestReadKeepsQuotedCommaInOneField`: a quoted field containing a comma is one field.
- `TestReadReturnsRecordLines`: `Lines` has one entry per data record and reflects the file line.
- `TestReadAcceptsInvalidUTF8`: a file with a byte that is not valid UTF-8 is read without error.
- `TestReadMissingFileIsOpenError`: a missing path returns an error that is not a CSV parse error.

## Programmatic Acceptance

=== AC reader-trims ===
Intent: Header names and values are trimmed of leading and trailing spaces.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvio", "TestReadTrimsHeaderAndValues")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC reader-trims ===

=== AC reader-quoted-comma ===
Intent: A quoted field containing a comma is read as a single field.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvio", "TestReadKeepsQuotedCommaInOneField")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC reader-quoted-comma ===

=== AC reader-record-lines ===
Intent: Each data record keeps its line number for diagnostics.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvio", "TestReadReturnsRecordLines")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC reader-record-lines ===

=== AC reader-invalid-utf8 ===
Intent: A file containing bytes that are not valid UTF-8 is read and not rejected.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvio", "TestReadAcceptsInvalidUTF8")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC reader-invalid-utf8 ===

=== AC reader-open-error ===
Intent: A missing file produces an open error that is distinct from a CSV parse error.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvio", "TestReadMissingFileIsOpenError")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC reader-open-error ===

## User Acceptance

- None.

## Guardrails

- Use `encoding/csv` from the standard library only.
- Application code never reads `sources/run.sh` or any supplied data file.
- Invalid UTF-8 is not a rejection condition.
- Do not modify any file under `sources/`.
