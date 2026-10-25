# FEATURE: Column Absent

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Reject a --column or --freq name that is not in the header with exit code 3 and nothing on standard output. |
| Depends On  | FEATURE-ColumnSelect.md, FEATURE-RejectMalformed.md |
| Provides    | exit 3 for an absent column name |
| Consumes    | csvstat --column <name> <file>, csvstat --freq <name> <file>, rejection exit mapping |

## Purpose

A misspelled or absent column name is a rejected request, not an empty result. The caller must be
able to tell it apart from a usage error (exit 2) and from an unreadable file (exit 1).

## Trigger

`csvstat --column <name> <file>` or `csvstat --freq <name> [--top <n>] <file>` where `<name>`
matches no trimmed header name of the file.

## Behavior

- Matching is exact and case-sensitive against trimmed header names (`salary` does not select
  `Salary`).
- The check happens after the file is read and validated and before any output is written, so
  standard output stays empty.
- The process exits 3 and writes one diagnostic line to standard error naming the file.
- A name that is present is unaffected and continues to succeed.
- A file that cannot be read still exits 1 first, and a usage error still exits 2 first; this
  story only owns the absent-name rejection.
- A file that is itself rejected (empty, ragged, bad header) exits 3 for that reason, with the same
  empty standard output.

## Reads

- The CSV file header after trimming.

## Writes

- One diagnostic line on standard error. Nothing on standard output.

## Programmatic Acceptance

=== AC colabsent-column-exit-3 ===
Intent: --column with a name absent from the header exits 3 and writes nothing to standard output.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import cli_runner

for path in ("sources/orders.csv", "sources/employees.csv", "sources/sensors.csv", "sources/header_only.csv"):
    code, out, err = cli_runner("--column", "no_such_column", path)
    print(out)
    print(err)
    assert code == 3
    assert out == ""
=== END AC colabsent-column-exit-3 ===

=== AC colabsent-freq-exit-3 ===
Intent: --freq with a name absent from the header exits 3 and writes nothing to standard output, with and without --top.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import cli_runner

for path in ("sources/orders.csv", "sources/employees.csv", "sources/sensors.csv"):
    for extra in ([], ["--top", "2"]):
        code, out, err = cli_runner("--freq", "no_such_column", *extra, path)
        print(out)
        print(err)
        assert code == 3
        assert out == ""
=== END AC colabsent-freq-exit-3 ===

=== AC colabsent-case-sensitive ===
Intent: A name that differs from a real header name only by letter case is absent and is rejected.
Suite: scoped
Requires: executable=go; scope=test

import csv

from drydock_go import cli_runner

path = "sources/employees.csv"
with open(path, newline="") as fh:
    heads = [h.strip() for h in next(csv.reader(fh))]
for name in heads:
    wrong = name.upper()
    assert wrong != name
    for mode in ("--column", "--freq"):
        code, out, err = cli_runner(mode, wrong, path)
        print(out)
        print(err)
        assert code == 3
        assert out == ""
=== END AC colabsent-case-sensitive ===

=== AC colabsent-present-name-still-succeeds ===
Intent: Every real header name is accepted by --column and --freq, so rejection is specific to absent names.
Suite: scoped
Requires: executable=go; scope=test

import csv

from drydock_go import cli_runner

path = "sources/orders.csv"
with open(path, newline="") as fh:
    heads = [h.strip() for h in next(csv.reader(fh))]
for name in heads:
    for mode in ("--column", "--freq"):
        code, out, err = cli_runner(mode, name, path)
        print(out)
        print(err)
        assert code == 0
        assert out != ""
=== END AC colabsent-present-name-still-succeeds ===

=== AC colabsent-header-only-and-ragged ===
Intent: An absent name is rejected for a header-only file, and a ragged file is rejected with exit 3 whatever name is given.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("--column", "no_such_column", "sources/header_only.csv")
print(out)
print(err)
assert code == 3
assert out == ""

for args in (["--column", "no_such_column"], ["--freq", "no_such_column"], ["--column", "sku"]):
    code, out, err = cli_runner(*args, "sources/ragged.csv")
    print(out)
    print(err)
    assert code == 3
    assert out == ""
=== END AC colabsent-header-only-and-ragged ===

## User Acceptance

- None.

## Guardrails

- Nothing is written to standard output when the name is absent, not even a header.
- An absent name is exit 3, never exit 1 or exit 2.
- Application code never reads `sources/run.sh` or any supplied data file and never embeds an answer.
