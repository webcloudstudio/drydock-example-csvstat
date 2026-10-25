# FEATURE: CSV Reading

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Read a UTF-8 RFC 4180 CSV file with encoding/csv and trim every header name and value of leading and trailing spaces. |
| Depends On  | ARCHITECTURE.md, FEATURE-CLI-Arguments.md |
| Provides    | csvdata.Read |
| Consumes    | internal/cli.Run |

## Purpose

`csvstat` profiles a CSV file. Every later rule (null detection, type inference, statistics,
frequencies) works on the trimmed table this feature produces, so trimming happens here, before
any other rule applies.

## Behavior

- `internal/csvdata` exposes `Read(path string) (Table, error)`.
- `Table` carries `Header []string` and `Records []Record`. `Record` carries `Line int` (the 1-based
  line on which the record starts, as reported by `csv.Reader.FieldPos`) and `Fields []string`.
- The file is read with `encoding/csv`. `FieldsPerRecord` is set to `-1` so that a record whose
  width differs from the header is returned to the caller intact; the rejection rules are owned by
  FEATURE-CSV-Rejection.md.
- The first record is the header. Every header name and every value has leading and trailing
  spaces trimmed (`strings.Trim(s, " ")`) before any other rule.
- Quoted fields decode per RFC 4180: a quoted field may contain commas, doubled quotes and newlines.
- `Read` returns an error that wraps the underlying `*os.PathError` when the file cannot be opened
  or read, so callers can tell an operational failure from a content problem.
- After argument validation, `cli.Run` reads the file with `csvdata.Read`. A successfully read file
  exits `0`; what is written to standard output is defined by later stories.

## Required unit tests

The build agent writes these tests in `internal/csvdata` with exactly these names:

- `TestReadTrimsHeaderAndValues`
- `TestReadQuotedFieldWithComma`
- `TestReadReturnsRaggedRecordsWithLineNumbers`

## Programmatic Acceptance

=== AC csv-reading-valid-file-exits-zero ===
Intent: A readable CSV whose header and values carry padding spaces is read successfully by the built executable.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "padded.csv")
    with open(path, "w", newline="") as f:
        f.write(" id , name \n 1 , Ada \n 2 , Grace \n")
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0
=== END AC csv-reading-valid-file-exits-zero ===

=== AC csv-reading-quoted-comma-exits-zero ===
Intent: A quoted field containing a comma and a doubled quote does not make the file ragged.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "quoted.csv")
    with open(path, "w", newline="") as f:
        f.write('id,notes\n1,"relocating to Austin, TX"\n2,"say ""hi"""\n')
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0
=== END AC csv-reading-quoted-comma-exits-zero ===

=== AC csv-reading-trims-header-and-values ===
Intent: The reader trims the header and every value before any other rule.

from drydock_go import run_test

result = run_test("./internal/csvdata", "TestReadTrimsHeaderAndValues")
print(result.stdout)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC csv-reading-trims-header-and-values ===

=== AC csv-reading-quoted-field ===
Intent: The reader decodes RFC 4180 quoted fields containing commas.

from drydock_go import run_test

result = run_test("./internal/csvdata", "TestReadQuotedFieldWithComma")
print(result.stdout)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC csv-reading-quoted-field ===

=== AC csv-reading-ragged-records-returned ===
Intent: The reader returns records of differing width with their line numbers so rejection can name the line.

from drydock_go import run_test

result = run_test("./internal/csvdata", "TestReadReturnsRaggedRecordsWithLineNumbers")
print(result.stdout)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC csv-reading-ragged-records-returned ===

## User Acceptance

- None.

## Guardrails

- Standard library only; read files with `encoding/csv`.
- Application code never reads `sources/run.sh` or any supplied data file and never embeds an answer.
- Trimming removes spaces only from the ends of a header name or value, never from its interior.
