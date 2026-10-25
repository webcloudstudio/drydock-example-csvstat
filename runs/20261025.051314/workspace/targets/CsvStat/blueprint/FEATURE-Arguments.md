# FEATURE: Arguments

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Parse csvstat options and exactly one file argument, with options preceding the file. |
| Depends On  | FEATURE-Executable.md |
| Provides    | cli.ParseArgs |
| Consumes    | cli.Run |

## Purpose

csvstat accepts two command forms:

```text
csvstat [--column <name>] <file>
csvstat --freq <name> [--top <n>] <file>
```

Options precede the file and exactly one file is required. This story delivers the parser and its
wiring into `Run`. Cross-option validation (`--top` without `--freq`, `--column` with `--freq`,
non-positive `--top`) belongs to the usage-errors story.

## Behavior

- `internal/cli` exports `ParseArgs(args []string) (Options, error)`.
- `Options` has fields `Column string`, `Freq string`, `Top int`, `TopSet bool`, `File string`, and
  flags are recorded as set or not set (a presence flag for `--column` and `--freq`, so an empty value
  is distinguishable from an absent option).
- Parsing uses a local `flag.FlagSet` created with `flag.ContinueOnError` and its output discarded.
  Options take a separate value (`--top 3`); a repeated option is a parse error.
- Parsing stops at the first non-option argument. That argument is the file. Any argument after it,
  including a later option, is an extra argument and a parse error. Options after the file are
  therefore rejected.
- A missing file, an extra argument, an unknown option, and an option with no value each return an
  error. `--top` is parsed as a decimal integer; a non-integer value is a parse error.
- `Run` calls `ParseArgs`. When it returns an error, `Run` writes one line to `stderr`, writes nothing
  to `stdout`, and returns 2.

## Tests For The Project Suite

The builder writes these `internal/cli` tests, named exactly as listed, using the named inputs:

- `TestParseArgsFileOnly`: `[file]` yields File set and no options.
- `TestParseArgsColumn`: `[--column name file]` yields Column and File.
- `TestParseArgsFreqTop`: `[--freq name --top 3 file]` yields Freq, Top 3, TopSet, and File.
- `TestParseArgsMissingFile`: `[--column name]` is an error.
- `TestParseArgsExtraFile`: two files is an error.
- `TestParseArgsUnknownOption`: `[--frobnicate file]` is an error.
- `TestParseArgsMissingValue`: `[file-less --top]` forms with no value are an error.
- `TestParseArgsOptionsAfterFile`: `[file --column name]` is an error.

## Programmatic Acceptance

=== AC args-file-only ===
Intent: A lone file argument parses.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

result = run_test("./internal/cli", "TestParseArgsFileOnly")
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC args-file-only ===

=== AC args-column ===
Intent: The --column option and its value parse before the file.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

result = run_test("./internal/cli", "TestParseArgsColumn")
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC args-column ===

=== AC args-freq-top ===
Intent: The --freq and --top options and their values parse before the file.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

result = run_test("./internal/cli", "TestParseArgsFreqTop")
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC args-freq-top ===

=== AC args-structural-errors ===
Intent: Missing file, extra file, unknown option, missing value, and options after the file are parse errors.
Requires: executable=go; scope=test

import sys
from drydock_go import run_test

names = [
    "TestParseArgsMissingFile",
    "TestParseArgsExtraFile",
    "TestParseArgsUnknownOption",
    "TestParseArgsMissingValue",
    "TestParseArgsOptionsAfterFile",
]
for name in names:
    result = run_test("./internal/cli", name)
    print(name, result.status)
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.passed, (name, result.status, result.stdout, result.stderr)
=== END AC args-structural-errors ===

=== AC args-valid-forms-not-usage-errors ===
Intent: The valid command forms are accepted by the executable and never exit with the usage-error code.
Requires: executable=go; scope=test

import csv
import os
import tempfile
from drydock_go import cli_runner

usage_exit = 2
with tempfile.TemporaryDirectory() as workdir:
    path = os.path.join(workdir, "data.csv")
    with open(path, "w", newline="", encoding="utf-8") as handle:
        writer = csv.writer(handle, lineterminator="\n")
        writer.writerow(["a", "b"])
        writer.writerow(["1", "2"])
    forms = [
        [path],
        ["--column", "a", path],
        ["--freq", "a", path],
        ["--freq", "a", "--top", "3", path],
    ]
    for form in forms:
        code, out, err = cli_runner(*form)
        print(form, code, out, err)
        assert code != usage_exit, form
=== END AC args-valid-forms-not-usage-errors ===

=== AC args-options-after-file-exit ===
Intent: An option placed after the file is rejected by the executable with the usage-error exit status and no standard output.

import csv
import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as workdir:
    path = os.path.join(workdir, "data.csv")
    with open(path, "w", newline="", encoding="utf-8") as handle:
        writer = csv.writer(handle, lineterminator="\n")
        writer.writerow(["a", "b"])
        writer.writerow(["1", "2"])
    code, out, err = cli_runner(path, "--column", "a")
    print(code, out, err)
    assert code == 2
    assert out == ""
=== END AC args-options-after-file-exit ===

## User Acceptance

- None.

## Guardrails

- Options precede the file; exactly one file is required.
- A parse failure writes nothing to standard output.
- Use a local `flag.FlagSet` with `flag.ContinueOnError`; never the global flag set.
