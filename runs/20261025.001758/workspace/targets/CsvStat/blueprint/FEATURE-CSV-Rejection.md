# FEATURE: CSV Rejection

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Reject an empty file, an empty or repeated header name, or a ragged record with exit 3, no standard output and one diagnostic line. |
| Depends On  | ARCHITECTURE.md, FEATURE-CSV-Reading.md |
| Provides    | rejected-input exit 3 |
| Consumes    | csvdata.Read |

## Purpose

A malformed file is rejected as a whole. Nothing is profiled and nothing reaches standard output,
so a caller never mistakes partial output for a result.

## Behavior

- After reading, `cli.Run` validates the table and exits `3` with nothing on standard output when:
  - the file has no records at all (zero bytes, or only blank content);
  - a header name is empty after trimming (including a header of only spaces);
  - a header name appears more than once (exact string comparison after trimming);
  - any data record has a field count different from the header;
  - `encoding/csv` reports a syntax error (for example a bare quote); this is rejected content, so
    it is exit `3`, not `1`.
- Exactly one diagnostic line is written to standard error. It names the file and, for a ragged
  record, the record's line number (`Record.Line`).
- A header followed by no data records is valid and is not rejected.
- Exit `1` (unreadable file) is owned by FEATURE-Read-Failure.md; exit `2` by the CLI features.

## Required unit tests

The build agent writes this test in `internal/cli` with exactly this name:

- `TestRejectionDiagnosticNamesFileAndLine` runs `Run` with `bytes.Buffer` writers on a ragged file
  and asserts stdout is empty, stderr is a single line, and that line contains the file path and the
  line number of the ragged record.

## Programmatic Acceptance

=== AC csv-rejection-ragged-supplied-file ===
Intent: The supplied ragged file is rejected with exit 3 and nothing on standard output.

import os
from drydock_go import cli_runner

path = os.path.abspath(os.path.join("sources", "ragged.csv"))
code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3
assert out == ""
=== END AC csv-rejection-ragged-supplied-file ===

=== AC csv-rejection-empty-file ===
Intent: A zero-byte file is rejected with exit 3 and nothing on standard output.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "empty.csv")
    open(path, "w").close()
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 3
    assert out == ""
=== END AC csv-rejection-empty-file ===

=== AC csv-rejection-empty-header-name ===
Intent: A header name that is empty or only spaces after trimming is rejected with exit 3.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    for i, header in enumerate(["a,,c", "a,   ,c"]):
        path = os.path.join(d, f"h{i}.csv")
        with open(path, "w", newline="") as f:
            f.write(header + "\n1,2,3\n")
        code, out, err = cli_runner(path)
        print(out)
        print(err)
        assert code == 3
        assert out == ""
=== END AC csv-rejection-empty-header-name ===

=== AC csv-rejection-repeated-header-name ===
Intent: A repeated header name, including one that repeats only after trimming, is rejected with exit 3.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    for i, header in enumerate(["a,b,a", "a, a ,b"]):
        path = os.path.join(d, f"r{i}.csv")
        with open(path, "w", newline="") as f:
            f.write(header + "\n1,2,3\n")
        code, out, err = cli_runner(path)
        print(out)
        print(err)
        assert code == 3
        assert out == ""
=== END AC csv-rejection-repeated-header-name ===

=== AC csv-rejection-short-and-long-records ===
Intent: A record with too few or too many fields is rejected with exit 3 and nothing on standard output.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    for i, row in enumerate(["1,2", "1,2,3,4"]):
        path = os.path.join(d, f"w{i}.csv")
        with open(path, "w", newline="") as f:
            f.write("a,b,c\n" + row + "\n")
        code, out, err = cli_runner(path)
        print(out)
        print(err)
        assert code == 3
        assert out == ""
=== END AC csv-rejection-short-and-long-records ===

=== AC csv-rejection-header-only-is-valid ===
Intent: A header with no data records is not rejected.

import os
from drydock_go import cli_runner

path = os.path.abspath(os.path.join("sources", "header_only.csv"))
code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0
=== END AC csv-rejection-header-only-is-valid ===

=== AC csv-rejection-diagnostic-names-file-and-line ===
Intent: The single diagnostic line names the file and the ragged record's line number.

from drydock_go import run_test

result = run_test("./internal/cli", "TestRejectionDiagnosticNamesFileAndLine")
print(result.stdout)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC csv-rejection-diagnostic-names-file-and-line ===

## User Acceptance

- None.

## Guardrails

- A rejected file never produces partial standard output.
- Only one diagnostic line is written to standard error per failing command.
- Rejected content is exit 3; it is never reported as exit 1 or 2.
