# FEATURE: CLI Option Validation

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Reject unknown options, missing option values, and invalid option combinations with exit 2 before the file is opened. |
| Depends On  | ARCHITECTURE.md, FEATURE-CLI-Arguments.md |
| Provides    | usage-error exit 2 |
| Consumes    | internal/cli.Run |

## Purpose

Every malformed invocation is a usage error: exit 2, nothing on standard output, one diagnostic line on standard error.

## Usage errors

| Invocation | Reason |
|---|---|
| `--frobnicate <file>` | Unknown option |
| `--column` or `--freq` or `--top` with no value | Missing option value |
| `--freq <name> --top 0 <file>` | `--top` is not a positive integer |
| `--freq <name> --top -1 <file>` | `--top` is negative |
| `--freq <name> --top abc <file>` | `--top` is not an integer |
| `--top 3 <file>` | `--top` without `--freq` |
| `--column <name> --freq <name> <file>` | `--column` combined with `--freq` |

## Behavior

- Validation completes before the file is opened, so a usage error outranks a missing or unreadable file.
- `--top` defaults to 5 when `--freq` is given and `--top` is not.
- Frequency behavior itself is owned by the Frequency stories; this story owns only the rejection of invalid usage.

## Reads

- The argument list passed to `Run`.

## Writes

- One standard error line on a usage error; nothing on standard output.

## Programmatic Acceptance

=== AC opts-unknown-option ===
Intent: An unknown option exits 2 and writes nothing to standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

path = "sources/orders.csv"
code, out, err = cli_runner("--frobnicate", path, package="./cmd/csvstat")
assert code == 2, (code, out, err)
assert out == ""
=== END AC opts-unknown-option ===

=== AC opts-missing-value ===
Intent: Each value-taking option given without its value exits 2 with empty standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

path = "sources/orders.csv"
cases = [
    ["--column"],
    ["--freq"],
    ["--freq", "qty", "--top"],
    ["--column", "--freq"],
]
for args in cases:
    code, out, err = cli_runner(*args, package="./cmd/csvstat")
    assert code == 2, (args, code, out, err)
    assert out == "", args
code, out, err = cli_runner(path, "--column", package="./cmd/csvstat")
assert code == 2, (code, out, err)
assert out == ""
=== END AC opts-missing-value ===

=== AC opts-top-not-positive-integer ===
Intent: --top of zero, a negative number, or a non-integer exits 2 with empty standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

path = "sources/sensors.csv"
for bad in ["0", "-1", "abc", "1.5", ""]:
    code, out, err = cli_runner(
        "--freq", "sensor", "--top", bad, path, package="./cmd/csvstat"
    )
    assert code == 2, (bad, code, out, err)
    assert out == "", bad
=== END AC opts-top-not-positive-integer ===

=== AC opts-top-without-freq ===
Intent: --top without --freq exits 2 with empty standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

path = "sources/sensors.csv"
code, out, err = cli_runner("--top", "3", path, package="./cmd/csvstat")
assert code == 2, (code, out, err)
assert out == ""
=== END AC opts-top-without-freq ===

=== AC opts-column-with-freq ===
Intent: --column combined with --freq exits 2 with empty standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

path = "sources/orders.csv"
name = "qty"
code, out, err = cli_runner(
    "--column", name, "--freq", name, path, package="./cmd/csvstat"
)
assert code == 2, (code, out, err)
assert out == ""
=== END AC opts-column-with-freq ===

=== AC opts-usage-outranks-missing-file ===
Intent: A usage error is reported as exit 2 even when the named file does not exist.
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as work:
    missing = os.path.join(work, "no-such-file.csv")
    code, out, err = cli_runner(
        "--freq", "x", "--top", "0", missing, package="./cmd/csvstat"
    )
    assert code == 2, (code, out, err)
    assert out == ""
=== END AC opts-usage-outranks-missing-file ===

## User Acceptance

- None.

## Guardrails

- Never open the input file before all usage validation has passed.
- A usage error never writes to standard output.
- Do not assert on the wording of the diagnostic; only the exit status and empty standard output are contract.
