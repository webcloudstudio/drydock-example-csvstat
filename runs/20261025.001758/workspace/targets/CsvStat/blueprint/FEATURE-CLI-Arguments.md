# FEATURE: CLI Arguments

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Parse csvstat options that precede exactly one file argument and reject a missing or extra file with exit 2. |
| Depends On  | ARCHITECTURE.md, FEATURE-Executable-Build.md |
| Provides    | csvstat [--column <name>] <file> |
| Consumes    | internal/cli.Run |

## Purpose

Command forms:

```text
csvstat [--column <name>] <file>
csvstat --freq <name> [--top <n>] <file>
```

Options precede the file. Exactly one file is required.

## Behavior

- `internal/cli` uses a local `flag.FlagSet` with `flag.ContinueOnError`, its output directed to the injected stderr writer, and defines `--column`, `--freq`, and `--top`.
- Parsing stops at the first non-option argument, so an option placed after the file is an extra positional and is a usage error.
- No file argument, or more than one file argument, exits 2 with nothing on standard output and one diagnostic line on standard error.
- A single file with valid options is not a usage error; what happens next is owned by later stories.
- Validation of option values and combinations is owned by FEATURE-CLI-Option-Validation.md.

## Reads

- The argument list passed to `Run`.

## Writes

- One standard error line on a usage error; nothing on standard output.

## Programmatic Acceptance

=== AC args-no-file ===
Intent: Running with no arguments exits 2 and writes nothing to standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner(package="./cmd/csvstat")
assert code == 2, (code, out, err)
assert out == ""
=== END AC args-no-file ===

=== AC args-extra-file ===
Intent: Two file arguments exit 2 and write nothing to standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

first = "sources/orders.csv"
second = "sources/sensors.csv"
code, out, err = cli_runner(first, second, package="./cmd/csvstat")
assert code == 2, (code, out, err)
assert out == ""
=== END AC args-extra-file ===

=== AC args-option-after-file ===
Intent: An option placed after the file is an extra positional and exits 2 with empty standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

path = "sources/orders.csv"
code, out, err = cli_runner(path, "--column", "qty", package="./cmd/csvstat")
assert code == 2, (code, out, err)
assert out == ""
=== END AC args-option-after-file ===

=== AC args-option-then-file-accepted ===
Intent: Options followed by exactly one file are not a usage error.
Requires: executable=go; scope=test

from drydock_go import cli_runner

path = "sources/orders.csv"
code, out, err = cli_runner("--column", "qty", path, package="./cmd/csvstat")
print(out)
print(err)
assert code != 2, (code, out, err)
=== END AC args-option-then-file-accepted ===

=== AC args-single-file-accepted ===
Intent: A single file with no options is not a usage error.
Requires: executable=go; scope=test

from drydock_go import cli_runner

path = "sources/sensors.csv"
code, out, err = cli_runner(path, package="./cmd/csvstat")
print(out)
print(err)
assert code != 2, (code, out, err)
=== END AC args-single-file-accepted ===

## User Acceptance

- None.

## Guardrails

- `Run` never calls `os.Exit`; it returns the exit code.
- A usage error never writes to standard output.
- Exactly one diagnostic line is written to standard error per failing command.
