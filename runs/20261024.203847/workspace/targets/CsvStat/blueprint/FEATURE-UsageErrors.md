# FEATURE: Usage Errors

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Reject invalid command lines with exit 2, nothing on standard output, and both synopses on standard error. |
| Depends On  | FEATURE-CliOptions.md, ARCHITECTURE.md |
| Provides    | exit 2 on usage errors |
| Consumes    | --column, --freq, --top |

## Purpose

Any malformed command line exits 2 before the file is opened. Standard output stays empty, and standard error receives both synopses so the user can see the valid forms.

## Usage Errors

- No file argument, or more than one file argument.
- An unknown option, including `--help` and `-h` (no help flag is defined).
- An option that requires a value but has none.
- `--top` whose value is not a positive integer (zero, negative, or non-numeric).
- `--top` given without `--freq`.
- `--column` given together with `--freq`.

## Behavior

- Exit code is 2 for every case above.
- Standard output is empty.
- Standard error receives both synopses: `csvstat [--column <name>] <file>` and `csvstat --freq <name> [--top <n>] <file>`.
- Usage errors are decided before the file is opened, so a nonexistent file never turns a usage error into exit 1.

## Programmatic Acceptance

=== AC usage-missing-or-extra-file ===
Intent: No file argument and two file arguments each exit 2 with nothing on standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

cases = {
    "none": [],
    "two": ["sources/orders.csv", "sources/sensors.csv"],
}
for label, args in cases.items():
    code, out, err = cli_runner(*args)
    print(label, code, out, err)
    assert code == 2, (label, code, out, err)
    assert out == "", (label, out)
=== END AC usage-missing-or-extra-file ===

=== AC usage-unknown-option ===
Intent: An unknown option, including --help and -h, exits 2 with nothing on standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

for flag in ("--frobnicate", "--help", "-h"):
    code, out, err = cli_runner(flag, "sources/orders.csv")
    print(flag, code, out, err)
    assert code == 2, (flag, code, out, err)
    assert out == "", (flag, out)
=== END AC usage-unknown-option ===

=== AC usage-missing-option-value ===
Intent: Options that need a value but have none exit 2 with nothing on standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

cases = {
    "column": ["--column"],
    "freq": ["--freq"],
    "top": ["--freq", "sensor", "sources/sensors.csv", "--top"],
}
for label, args in cases.items():
    code, out, err = cli_runner(*args)
    print(label, code, out, err)
    assert code == 2, (label, code, out, err)
    assert out == "", (label, out)
=== END AC usage-missing-option-value ===

=== AC usage-invalid-top ===
Intent: A --top that is zero, negative, or not a number exits 2 with nothing on standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

for value in ("0", "-1", "abc", "1.5"):
    code, out, err = cli_runner("--freq", "sensor", "--top", value, "sources/sensors.csv")
    print(value, code, out, err)
    assert code == 2, (value, code, out, err)
    assert out == "", (value, out)
=== END AC usage-invalid-top ===

=== AC usage-top-without-freq ===
Intent: --top without --freq exits 2 with nothing on standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("--top", "3", "sources/sensors.csv")
print(out)
print(err)
assert code == 2, (code, out, err)
assert out == "", out
=== END AC usage-top-without-freq ===

=== AC usage-column-with-freq ===
Intent: --column combined with --freq exits 2 with nothing on standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("--column", "qty", "--freq", "qty", "sources/orders.csv")
print(out)
print(err)
assert code == 2, (code, out, err)
assert out == "", out
=== END AC usage-column-with-freq ===

=== AC usage-decided-before-file-open ===
Intent: A usage error on a nonexistent file still exits 2, not 1, because usage is checked before the file is opened.
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

work = tempfile.mkdtemp()
missing = os.path.join(work, "no-such-file.csv")
assert not os.path.exists(missing)

code, out, err = cli_runner("--top", "3", missing)
print(out)
print(err)
assert code == 2, (code, out, err)
assert out == "", out
=== END AC usage-decided-before-file-open ===

## User Acceptance

- Standard error for a usage error shows both synopses, `csvstat [--column <name>] <file>` and `csvstat --freq <name> [--top <n>] <file>`.

## Guardrails

- No help flag exists; `--help` and `-h` are unknown options.
- A usage error never writes anything to standard output.
- Usage errors are never reported as exit 1 or exit 3.
