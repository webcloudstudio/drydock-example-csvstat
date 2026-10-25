# FEATURE: Reject Malformed Input

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Reject empty files, bad header names, ragged records, and CSV syntax errors with exit code 3 and nothing on standard output. |
| Depends On  | FEATURE-CsvReading.md, ARCHITECTURE.md |
| Provides    | exit code 3 for rejected input |
| Consumes    | internal/csvio |

## Purpose

Malformed input is rejected before any output is produced. Rejection exits 3, writes nothing to
standard output, and writes one line to standard error naming the file and, for a record problem,
its line number.

## Behavior

1. An empty file, or a file with no records at all, is rejected.
2. A header name that is empty after trimming, or repeated after trimming, is rejected.
3. A record whose field count differs from the header is rejected, naming its line number.
4. A syntax error reported by `encoding/csv`, such as a bare quote in an unquoted field, is
   rejected with exit 3 and names the file and line number.
5. All output is buffered; nothing reaches standard output unless validation of the whole file
   succeeded.
6. A header-only file is valid and is not rejected.

## Required Go Tests

The build adds these tests in `internal/csvio` (names are exact): `TestRejectEmptyFile`,
`TestRejectBadHeader`, `TestRejectRaggedRecord`.

## Programmatic Acceptance

=== AC reject-ragged-supplied ===
Intent: A record with a different field count than the header exits 3 with nothing on standard output and the file named on standard error.

from drydock_go import cli_runner

path = "sources/ragged.csv"
code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3
assert out == ""
assert path in err
=== END AC reject-ragged-supplied ===

=== AC reject-empty-file ===
Intent: A zero-byte file and a file holding only a blank line both exit 3 with nothing on standard output.

import tempfile
from pathlib import Path
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    for index, payload in enumerate([b"", b"\n"]):
        target = Path(tmp) / ("empty" + str(index) + ".csv")
        target.write_bytes(payload)
        code, out, err = cli_runner(str(target))
        print(out)
        print(err)
        assert code == 3
        assert out == ""
=== END AC reject-empty-file ===

=== AC reject-bad-header-names ===
Intent: An empty header name, a repeated header name, and a name repeated only after trimming each exit 3 with nothing on standard output.

import tempfile
from pathlib import Path
from drydock_go import cli_runner

cases = [",b\n1,2\n", "a,a\n1,2\n", "a, a\n1,2\n", "a, \n1,2\n"]
with tempfile.TemporaryDirectory() as tmp:
    for index, text in enumerate(cases):
        target = Path(tmp) / ("header" + str(index) + ".csv")
        target.write_text(text)
        code, out, err = cli_runner(str(target))
        print(out)
        print(err)
        assert code == 3, text
        assert out == ""
=== END AC reject-bad-header-names ===

=== AC reject-short-record ===
Intent: A record with fewer fields than the header exits 3 with nothing on standard output.

import tempfile
from pathlib import Path
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    target = Path(tmp) / "short.csv"
    target.write_text("a,b\n1\n")
    code, out, err = cli_runner(str(target))
    print(out)
    print(err)
    assert code == 3
    assert out == ""
    assert str(target) in err
=== END AC reject-short-record ===

=== AC reject-csv-syntax-error ===
Intent: A bare quote inside an unquoted field is a CSV syntax error that exits 3 with nothing on standard output.

import tempfile
from pathlib import Path
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    target = Path(tmp) / "bare_quote.csv"
    target.write_text("a,b\n1,x" + '"' + "y\n")
    code, out, err = cli_runner(str(target))
    print(out)
    print(err)
    assert code == 3
    assert out == ""
    assert str(target) in err
=== END AC reject-csv-syntax-error ===

=== AC reject-valid-files-pass ===
Intent: Well-formed files, including a header-only file, are not rejected.

from drydock_go import cli_runner

for path in ["sources/employees.csv", "sources/header_only.csv"]:
    code, out, err = cli_runner(path)
    print(err)
    assert code == 0, path
=== END AC reject-valid-files-pass ===

=== AC reject-empty-file-test ===
Intent: The package test for empty-file rejection passes.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvio", "TestRejectEmptyFile")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC reject-empty-file-test ===

=== AC reject-bad-header-test ===
Intent: The package test for header-name rejection passes.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvio", "TestRejectBadHeader")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC reject-bad-header-test ===

=== AC reject-ragged-test ===
Intent: The package test for ragged-record rejection passes.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/csvio", "TestRejectRaggedRecord")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC reject-ragged-test ===

## User Acceptance

- None.

## Guardrails

- A rejected input never produces partial standard output.
- Only the diagnostic on standard error names the file and line; its wording is not a contract.
- A header-only file is valid input and must not be rejected.
- Standard library only.
