# FEATURE: Unreadable File

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Exits 1 with an empty standard output when the named file cannot be opened or read. |
| Depends On  | FEATURE-FailureDiagnostics.md |
| Provides    | exit code 1 |
| Consumes    | cli.Run, csvdata.Read |

## Purpose

Exit code 1 is the operational-failure code. It is used when the file named on the command line does not exist, is not permitted to be read, or fails while being read. It is distinct from exit 3, which means the file was read but its content was rejected.

## Behavior

- An open or read error from the operating system maps to exit 1, not exit 3.
- Standard output stays empty.
- One diagnostic line naming the file is written to standard error.
- This applies to every invocation form: profile, `--column`, and `--freq`.
- Usage errors (exit 2) take precedence over file access, because argument parsing completes before the file is opened.

## Programmatic Acceptance

=== AC unreadable-missing-file ===
Intent: A path that does not exist exits 1 and writes nothing to standard output.
Requires: executable=go; scope=test

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    missing = os.path.join(tmp, "no-such-file.csv")
    assert not os.path.exists(missing)
    code, out, err = cli_runner(missing)
print(out)
print(err)
assert code == 1, code
assert out == "", out
lines = [ln for ln in err.splitlines() if ln.strip()]
assert len(lines) == 1, lines
=== END AC unreadable-missing-file ===

=== AC unreadable-missing-file-freq ===
Intent: A missing file under --freq also exits 1 with nothing on standard output.
Requires: executable=go; scope=test

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    missing = os.path.join(tmp, "absent.csv")
    code, out, err = cli_runner("--freq", "anything", missing)
print(out)
print(err)
assert code == 1, code
assert out == "", out
=== END AC unreadable-missing-file-freq ===

=== AC unreadable-missing-file-column ===
Intent: A missing file under --column also exits 1 with nothing on standard output.
Requires: executable=go; scope=test

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    missing = os.path.join(tmp, "absent.csv")
    code, out, err = cli_runner("--column", "anything", missing)
print(out)
print(err)
assert code == 1, code
assert out == "", out
=== END AC unreadable-missing-file-column ===

=== AC unreadable-directory-as-file ===
Intent: A directory named as the file cannot be read and exits 1 with nothing on standard output.
Requires: executable=go; scope=test

import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    code, out, err = cli_runner(tmp)
print(out)
print(err)
assert code == 1, code
assert out == "", out
=== END AC unreadable-directory-as-file ===

=== AC unreadable-readable-file-still-succeeds ===
Intent: A readable fixture is not treated as unreadable and exits 0.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("sources/header_only.csv")
print(out)
print(err)
assert code == 0, code
assert out != ""
=== END AC unreadable-readable-file-still-succeeds ===

## User Acceptance

- None.

## Guardrails

- Exit 1 is reserved for files that cannot be opened or read; rejected content is exit 3.
- A failing command writes nothing to standard output.
