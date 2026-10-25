# FEATURE: Read Failure

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Report a file that cannot be opened or read, including a directory path, with exit 1 and nothing on standard output. |
| Depends On  | FEATURE-CsvReader.md, ARCHITECTURE.md |
| Provides    | exit code 1 for operational failure |
| Consumes    | internal/csvio.Read |

## Purpose

When the file path cannot be opened or read, the application exits 1, writes nothing to standard output, and writes one diagnostic line to standard error naming the path. This covers a missing file, a permission failure, and a path that names a directory. The application must not panic; a Go panic exits 2, which would be misreported as a usage error.

## Behavior

- An open failure from `os.Open` maps to exit 1.
- A read failure after opening, such as reading a directory, maps to exit 1.
- A valid path with a malformed argument list is a usage error (exit 2), decided before the file is touched.
- A CSV syntax or structure problem is a rejection (exit 3), not exit 1.

## Programmatic Acceptance

=== AC readfail-missing-file ===
Intent: A path that does not exist exits 1 with nothing on standard output.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as work:
    path = os.path.join(work, "no-such-file.csv")
    assert not os.path.exists(path)
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 1
    assert out == ""
=== END AC readfail-missing-file ===

=== AC readfail-directory ===
Intent: A path naming a directory exits 1 with nothing on standard output and does not panic.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as work:
    assert os.path.isdir(work)
    code, out, err = cli_runner(work)
    print(out)
    print(err)
    assert code == 1
    assert out == ""
=== END AC readfail-directory ===

=== AC readfail-distinct-from-rejection ===
Intent: An unreadable path exits with a different code from a rejected file, so exit 1 and exit 3 are not conflated.

import os
import tempfile
from drydock_go import cli_runner

ragged = os.path.abspath(os.path.join("sources", "ragged.csv"))
assert os.path.isfile(ragged)
rejected_code, rejected_out, rejected_err = cli_runner(ragged)

with tempfile.TemporaryDirectory() as work:
    missing = os.path.join(work, "absent.csv")
    missing_code, missing_out, missing_err = cli_runner(missing)

print(rejected_err)
print(missing_err)
assert missing_code != rejected_code
assert missing_code == 1
assert rejected_code == 3
=== END AC readfail-distinct-from-rejection ===

## User Acceptance

- None.

## Guardrails

- A read failure writes nothing to standard output.
- The application never panics on an unreadable path.
- Diagnostics go to standard error only and are never compared.
- Do not modify any file under `sources/`.
