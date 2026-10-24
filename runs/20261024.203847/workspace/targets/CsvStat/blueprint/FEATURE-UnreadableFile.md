# FEATURE: Unreadable File

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Exit 1 with nothing on standard output when the input file cannot be opened or read. |
| Depends On  | FEATURE-NullTrim.md, ARCHITECTURE.md |
| Provides    | exit 1 on unreadable file |
| Consumes    | reader.Read |

## Purpose

An operational failure to open or read the file is distinct from a malformed file. It exits 1, writes nothing to standard output, and writes one diagnostic line to standard error naming the file.

## Behavior

- An error from opening the file exits 1.
- An error while reading the file's bytes exits 1.
- Standard output is empty in both cases.
- Standard error carries one diagnostic line naming the file path.
- Read errors are never reported as exit 3; exit 3 is reserved for content that was read but rejected.

## Programmatic Acceptance

=== AC unreadable-missing-file ===
Intent: A path that does not exist exits 1 with nothing on standard output.
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

work = tempfile.mkdtemp()
missing = os.path.join(work, "no-such-file.csv")
assert not os.path.exists(missing)

code, out, err = cli_runner(missing)
print(out)
print(err)
assert code == 1, (code, out, err)
assert out == "", out
=== END AC unreadable-missing-file ===

=== AC unreadable-directory-path ===
Intent: A path that names a directory cannot be read as a file and exits 1 with nothing on standard output.
Requires: executable=go; scope=test

import tempfile

from drydock_go import cli_runner

work = tempfile.mkdtemp()

code, out, err = cli_runner(work)
print(out)
print(err)
assert code == 1, (code, out, err)
assert out == "", out
=== END AC unreadable-directory-path ===

=== AC unreadable-not-confused-with-rejected ===
Intent: A readable but malformed file exits 3, not 1, so the two failure classes stay distinct.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("sources/ragged.csv")
print(out)
print(err)
assert code != 1, (code, out, err)
assert code == 3, (code, out, err)
=== END AC unreadable-not-confused-with-rejected ===

## User Acceptance

- The standard error diagnostic for an unreadable file names the file path.

## Guardrails

- An unreadable file never writes anything to standard output.
- The file is read once; no partial profile is written before a read error is detected.
