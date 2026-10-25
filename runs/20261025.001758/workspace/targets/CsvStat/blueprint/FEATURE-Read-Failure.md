# FEATURE: Read Failure

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Exit 1 with nothing on standard output when the file cannot be opened or read. |
| Depends On  | ARCHITECTURE.md, FEATURE-CSV-Rejection.md |
| Provides    | operational-failure exit 1 |
| Consumes    | csvdata.Read |

## Purpose

An operational failure (the file is unavailable) is distinguished from rejected content so callers
can tell a missing input from a malformed one.

## Behavior

- When `csvdata.Read` fails to open or read the path (nonexistent path, a directory, no
  permission), `cli.Run` exits `1`.
- Nothing is written to standard output; one diagnostic line naming the file is written to standard
  error.
- Exit `1` is reserved for this case. Content that was read but rejected is exit `3`
  (FEATURE-CSV-Rejection.md); usage errors are exit `2` (FEATURE-CLI-Arguments.md and
  FEATURE-CLI-Option-Validation.md).

## Programmatic Acceptance

=== AC read-failure-missing-file ===
Intent: A path that does not exist exits 1 with nothing on standard output.

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    missing = os.path.join(d, "no-such-file.csv")
    code, out, err = cli_runner(missing)
    print(out)
    print(err)
    assert code == 1
    assert out == ""
=== END AC read-failure-missing-file ===

=== AC read-failure-directory ===
Intent: A path that is a directory cannot be read and exits 1 with nothing on standard output.

import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    code, out, err = cli_runner(d)
    print(out)
    print(err)
    assert code == 1
    assert out == ""
=== END AC read-failure-directory ===

=== AC read-failure-distinct-from-rejection ===
Intent: A readable but ragged file exits 3 while an unreadable one exits 1, so the two codes are not conflated.

import os
import tempfile
from drydock_go import cli_runner

ragged = os.path.abspath(os.path.join("sources", "ragged.csv"))
with tempfile.TemporaryDirectory() as d:
    missing = os.path.join(d, "absent.csv")
    rejected_code, _, _ = cli_runner(ragged)
    failed_code, out, err = cli_runner(missing)
    print(out)
    print(err)
    assert rejected_code == 3
    assert failed_code == 1
    assert failed_code != rejected_code
=== END AC read-failure-distinct-from-rejection ===

## User Acceptance

- None.

## Guardrails

- Exit `1` never accompanies output on standard output.
- Usage errors are not handled here and keep exit `2`.
