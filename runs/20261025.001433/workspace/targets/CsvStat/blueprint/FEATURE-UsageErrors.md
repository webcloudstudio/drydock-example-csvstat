# FEATURE: Usage Errors

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Maps every malformed command line to exit code 2 with nothing on standard output. |
| Depends On  | FEATURE-CliParsing.md, ARCHITECTURE.md |
| Provides    | exit code 2 |
| Consumes    | csvstat [--column <name>] <file>, csvstat --freq <name> [--top <n>] <file> |

## Purpose

Any command line that does not match one of the two command forms is a usage error.

## Behavior

Exit code 2, nothing on standard output, and a diagnostic on standard error, for each of:

- no file argument;
- more than one file argument;
- an unknown option;
- an option missing its value (`--column`, `--freq`, or `--top` as the last argument);
- `--top` not a positive integer;
- `--top` without `--freq`;
- `--column` together with `--freq`.

The flag set uses `flag.ContinueOnError` with its output directed to standard error, so the parser
never writes to standard output on failure. Usage errors are decided before the file is opened, so
a usage error for a missing file path never becomes exit 1. `--help` is not a usage error and is
owned by FEATURE-CliParsing.md.

## Programmatic Acceptance

=== AC usage-missing-or-extra-file ===
Intent: No file, or more than one file, exits 2 with empty standard output.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

orders = os.path.abspath("sources/orders.csv")
sensors = os.path.abspath("sources/sensors.csv")

code, out, err = cli_runner()
print(err)
assert code == 2, (code, out, err)
assert out == "", out

code, out, err = cli_runner(orders, sensors)
print(err)
assert code == 2, (code, out, err)
assert out == "", out

code, out, err = cli_runner("--column", "qty", orders, sensors)
print(err)
assert code == 2, (code, out, err)
assert out == "", out
=== END AC usage-missing-or-extra-file ===

=== AC usage-unknown-option ===
Intent: An unknown option exits 2 with empty standard output.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

orders = os.path.abspath("sources/orders.csv")
code, out, err = cli_runner("--frobnicate", orders)
print(err)
assert code == 2, (code, out, err)
assert out == "", out
=== END AC usage-unknown-option ===

=== AC usage-missing-option-value ===
Intent: Each value-taking option given as the last argument exits 2 with empty standard output.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

orders = os.path.abspath("sources/orders.csv")
for opt in ("--column", "--freq"):
    code, out, err = cli_runner(orders, opt)
    print(opt, code, err)
    assert code == 2, (opt, code, out, err)
    assert out == "", (opt, out)

code, out, err = cli_runner("--freq", "qty", orders, "--top")
print(err)
assert code == 2, (code, out, err)
assert out == "", out
=== END AC usage-missing-option-value ===

=== AC usage-bad-top ===
Intent: A --top of zero or non-numeric text exits 2 with empty standard output.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

sensors = os.path.abspath("sources/sensors.csv")
for bad in ("0", "-4", "many"):
    code, out, err = cli_runner("--freq", "sensor", "--top", bad, sensors)
    print(bad, code, err)
    assert code == 2, (bad, code, out, err)
    assert out == "", (bad, out)
=== END AC usage-bad-top ===

=== AC usage-top-without-freq ===
Intent: --top without --freq exits 2 with empty standard output.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

sensors = os.path.abspath("sources/sensors.csv")
code, out, err = cli_runner("--top", "3", sensors)
print(err)
assert code == 2, (code, out, err)
assert out == "", out
=== END AC usage-top-without-freq ===

=== AC usage-column-with-freq ===
Intent: --column combined with --freq exits 2 with empty standard output.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

orders = os.path.abspath("sources/orders.csv")
code, out, err = cli_runner("--column", "qty", "--freq", "qty", orders)
print(err)
assert code == 2, (code, out, err)
assert out == "", out
=== END AC usage-column-with-freq ===

=== AC usage-before-file-access ===
Intent: A usage error on a nonexistent path is exit 2, not exit 1, because arguments are validated before the file is opened.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

missing = os.path.abspath("no-such-file-for-usage-check.csv")
assert not os.path.exists(missing)
code, out, err = cli_runner("--top", "3", missing)
print(err)
assert code == 2, (code, out, err)
assert out == "", out
=== END AC usage-before-file-access ===

## User Acceptance

- None.

## Guardrails

- A usage error never writes to standard output.
- A usage error never exits 1 or 3.
- Diagnostics go to standard error and are never compared.
