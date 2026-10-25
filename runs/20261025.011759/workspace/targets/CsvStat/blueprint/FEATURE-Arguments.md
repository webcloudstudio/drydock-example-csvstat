# FEATURE: Arguments

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Parses the csvstat command line and rejects malformed usage with exit 2. |
| Depends On  | FEATURE-Build.md |
| Provides    | --column, --freq, --top, exit code 2 |
| Consumes    | internal/cli.Run |

## Purpose

The command forms are `csvstat [--column <name>] <file>` and `csvstat --freq <name> [--top <n>] <file>`. This feature parses them and turns every usage mistake into exit code 2 with nothing on standard output.

## Behavior

- `internal/cli.Run(args []string, stdout, stderr io.Writer) int` parses `args` with a local `flag.FlagSet` created with `flag.ContinueOnError`, its own output set to `io.Discard`.
- Options precede the file. Exactly one positional argument, the file, is required. Anything after the first positional argument counts as an extra argument.
- Recognized options: `--column <name>`, `--freq <name>`, `--top <n>`.
- Exit 2 with empty standard output for each of:
  - no file, or more than one file;
  - an unknown option;
  - an option given without its value;
  - `--column` together with `--freq`;
  - `--top` without `--freq`.
- On exit 2, write exactly one line to standard error naming both command forms. There is no `--help` option.
- A well-formed command line is accepted and handed to the later behavior; it never exits 2.
- A `--top` value that is not a positive integer is owned by the `--top` feature (FEATURE-FrequencyTop.md), not this one.

## Required project tests

The builder writes tests calling `cli.Run` with `bytes.Buffer` writers for every case above, asserting the exit code, empty standard output, and a single-line diagnostic.

## Programmatic Acceptance

=== AC arguments-file-count ===
Intent: A missing file or an extra file exits 2 with nothing on standard output.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    first = os.path.join(d, "first.csv")
    second = os.path.join(d, "second.csv")
    for path in (first, second):
        with open(path, "w", newline="") as f:
            f.write("a,b\n1,2\n")

    code, out, err = cli_runner()
    print(out)
    print(err)
    assert code == 2, (code, out, err)
    assert out == ""

    code, out, err = cli_runner(first, second)
    print(out)
    print(err)
    assert code == 2, (code, out, err)
    assert out == ""

    code, out, err = cli_runner("--column", "a")
    print(out)
    print(err)
    assert code == 2, (code, out, err)
    assert out == ""
=== END AC arguments-file-count ===

=== AC arguments-bad-options ===
Intent: An unknown option and an option without its value exit 2 with nothing on standard output.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "data.csv")
    with open(path, "w", newline="") as f:
        f.write("a,b\n1,2\n")

    cases = [
        ["--frobnicate", path],
        ["--column"],
        ["--freq"],
        ["--freq", "a", "--top"],
    ]
    for args in cases:
        code, out, err = cli_runner(*args)
        print(args, code)
        print(out)
        print(err)
        assert code == 2, (args, code, out, err)
        assert out == "", (args, out)
=== END AC arguments-bad-options ===

=== AC arguments-option-combinations ===
Intent: --column with --freq, and --top without --freq, exit 2 with nothing on standard output.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "data.csv")
    with open(path, "w", newline="") as f:
        f.write("a,b\n1,2\n")

    cases = [
        ["--column", "a", "--freq", "a", path],
        ["--top", "3", path],
        ["--column", "a", "--top", "3", path],
    ]
    for args in cases:
        code, out, err = cli_runner(*args)
        print(args, code)
        print(out)
        print(err)
        assert code == 2, (args, code, out, err)
        assert out == "", (args, out)
=== END AC arguments-option-combinations ===

=== AC arguments-options-precede-file ===
Intent: An option placed after the file is an extra argument and exits 2.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "data.csv")
    with open(path, "w", newline="") as f:
        f.write("a,b\n1,2\n")

    code, out, err = cli_runner(path, "--column", "a")
    print(out)
    print(err)
    assert code == 2, (code, out, err)
    assert out == ""
=== END AC arguments-options-precede-file ===

=== AC arguments-usage-line ===
Intent: A usage error writes exactly one non-empty diagnostic line to standard error.
Suite: scoped
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner()
print(out)
print(err)
assert code == 2, (code, out, err)
assert out == ""
lines = [line for line in err.splitlines() if line.strip()]
assert len(lines) == 1, err
=== END AC arguments-usage-line ===

=== AC arguments-wellformed-accepted ===
Intent: Well-formed command lines are never rejected as usage errors.
Suite: scoped
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "data.csv")
    with open(path, "w", newline="") as f:
        f.write("a,b\n1,2\n")

    cases = [
        [path],
        ["--column", "a", path],
        ["--freq", "a", path],
        ["--freq", "a", "--top", "2", path],
    ]
    for args in cases:
        code, out, err = cli_runner(*args)
        print(args, code)
        print(err)
        assert code != 2, (args, code, out, err)
=== END AC arguments-wellformed-accepted ===

## User Acceptance

- None.

## Guardrails

- No `--help` option is added.
- Only `main` calls `os.Exit`; `Run` returns the exit code.
- Flag parsing never writes to the process's standard output or standard error directly; all output goes through the writers passed to `Run`.
- Standard library only.
