# FEATURE: Unreadable File

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Maps a file that cannot be opened or read to exit code 1 with nothing on standard output. |
| Depends On  | FEATURE-CliParsing.md, ARCHITECTURE.md |
| Provides    | exit code 1 |
| Consumes    | csvstat [--column <name>] <file>, csvstat --freq <name> [--top <n>] <file> |

## Purpose

An operational failure to obtain the input is reported as exit 1, distinct from a usage error (2)
and from rejected content (3).

## Behavior

- A path that does not exist, cannot be opened, is a directory, or fails while being read exits 1.
- Nothing is written to standard output.
- One diagnostic line on standard error names the file.
- The rule applies to the profile form and the `--freq` form alike.
- Arguments are validated first, so a well-formed command line naming an unreadable file is exit 1
  and a malformed one is exit 2.

## Programmatic Acceptance

=== AC unreadable-missing-path ===
Intent: A path that does not exist exits 1 with empty standard output.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

missing = os.path.abspath("no-such-file-for-unreadable-check.csv")
assert not os.path.exists(missing)
code, out, err = cli_runner(missing)
print(err)
assert code == 1, (code, out, err)
assert out == "", out
=== END AC unreadable-missing-path ===

=== AC unreadable-directory ===
Intent: A path that is a directory exits 1 with empty standard output.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

directory = os.path.abspath("sources")
assert os.path.isdir(directory)
code, out, err = cli_runner(directory)
print(err)
assert code == 1, (code, out, err)
assert out == "", out
=== END AC unreadable-directory ===

=== AC unreadable-with-options ===
Intent: The --column and --freq forms also exit 1 for an unreadable file.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

missing = os.path.abspath("no-such-file-for-unreadable-check.csv")
assert not os.path.exists(missing)

code, out, err = cli_runner("--column", "qty", missing)
print(err)
assert code == 1, (code, out, err)
assert out == "", out

code, out, err = cli_runner("--freq", "qty", "--top", "2", missing)
print(err)
assert code == 1, (code, out, err)
assert out == "", out
=== END AC unreadable-with-options ===

=== AC unreadable-diagnostic-on-stderr ===
Intent: The failure is reported on standard error and names the file.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

missing = os.path.abspath("no-such-file-for-unreadable-check.csv")
code, out, err = cli_runner(missing)
print(err)
assert code == 1, (code, out, err)
assert out == "", out
assert os.path.basename(missing) in err
=== END AC unreadable-diagnostic-on-stderr ===

## User Acceptance

- None.

## Guardrails

- An unreadable file never produces standard output.
- An unreadable file is never exit 2 or 3.
- Application code never reads `sources/run.sh` or any supplied data file.
