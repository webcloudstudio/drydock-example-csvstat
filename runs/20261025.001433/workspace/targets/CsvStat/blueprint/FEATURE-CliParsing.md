# FEATURE: Command Line Parsing

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Accept the two csvstat command forms and the --help text, producing a typed request for dispatch. |
| Depends On  | ARCHITECTURE.md |
| Provides    | csvstat [--column <name>] <file>, csvstat --freq <name> [--top <n>] <file>, csvstat --help |
| Consumes    | internal/cli.Run |

## Purpose

`internal/cli` parses arguments with a local `flag.FlagSet` created with `flag.ContinueOnError`,
whose output is directed to standard error. Options precede exactly one file.

| Option | Meaning |
|---|---|
| `--column <name>` | Profile only the named column. |
| `--freq <name>` | List the most frequent values of the named column. |
| `--top <n>` | With `--freq`, list at most n values; default 5. |
| `--help` | Write usage text to standard output and exit 0. |

## Behavior

1. Parsing yields a typed request: mode (profile or frequency), column name, frequency name, top
   count, and file path.
2. `--help` writes usage text naming both command forms and every option to standard output and
   returns 0, writing nothing to standard error.
3. Profile and frequency modes are dispatched through small interfaces held by `Run`, so later
   stories plug in without changing parsing.
4. Names given to `--column` and `--freq` are kept exactly as given; they are matched later,
   exactly and case-sensitively, against trimmed header names.
5. Usage-error exit mapping is owned by the usage-errors story; this story only has to produce the
   request for a well-formed command line.

## Required Go Tests

The build adds these tests in `internal/cli` (names are exact): `TestParseHelp`,
`TestParseColumnForm`, `TestParseFreqForm`.

## Programmatic Acceptance

=== AC cliparse-help-exits-zero ===
Intent: csvstat --help exits 0 and writes usage naming the command and every option to standard output.

from drydock_go import cli_runner

code, out, err = cli_runner("--help")
print(out)
print(err)
assert code == 0
for token in ["csvstat", "--column", "--freq", "--top"]:
    assert token in out, token
=== END AC cliparse-help-exits-zero ===

=== AC cliparse-help-test ===
Intent: The package test for the help form passes.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseHelp")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cliparse-help-test ===

=== AC cliparse-column-form-test ===
Intent: Parsing the --column form yields a profile request with the named column and the file path.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseColumnForm")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cliparse-column-form-test ===

=== AC cliparse-freq-form-test ===
Intent: Parsing the --freq form yields a frequency request with the name, default top of 5, and the file path.
Requires: executable=go; scope=test

from drydock_go import run_test

result = run_test("./internal/cli", "TestParseFreqForm")
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC cliparse-freq-form-test ===

## User Acceptance

- None.

## Guardrails

- `Run` never calls `os.Exit`; only `cmd/csvstat/main.go` does.
- The flag set writes its own diagnostics to standard error, never standard output.
- Standard library only.
