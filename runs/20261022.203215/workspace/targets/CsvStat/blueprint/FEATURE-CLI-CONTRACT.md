# FEATURE: CLI Contract

| Field       | Value |
|-------------|-------|
| Version     | 20261022 V1 |
| Description | Defines csvstat command parsing, option validation, dispatch, and exit-code behavior. |
| Depends On  | ARCHITECTURE.md |
| Provides    | csvstat argument parser and command dispatch |
| Consumes    | Go module, ./cmd/csvstat entry point |

## Command Forms

```text
csvstat [--column <name>] <file>
csvstat --freq <name> [--top <n>] <file>
```

Options precede the single required file argument. `--column` and `--freq` are mutually
exclusive. `--top` is valid only with `--freq` and must be a positive integer.

## Exit Contract

- `0`: successful profile or frequency report.
- `1`: the file cannot be opened or read.
- `2`: malformed arguments, unknown options, missing values, extra files, invalid `--top`, or
  incompatible options.
- `3`: rejected CSV input or a requested column absent from the header.

Failures write no bytes to standard output. Diagnostics are written to standard error.

## Programmatic Acceptance

=== AC cli-contract-usage ===
Intent: Malformed invocations return the usage-error exit status and produce no standard output.

import pathlib
import subprocess
import tempfile

with tempfile.TemporaryDirectory() as directory:
    path = pathlib.Path(directory) / "data.csv"
    path.write_text("a\n1\n", encoding="utf-8")

    cases = [
        ["go", "run", "./cmd/csvstat"],
        ["go", "run", "./cmd/csvstat", str(path), str(path)],
        ["go", "run", "./cmd/csvstat", "--top", "0", str(path)],
        ["go", "run", "./cmd/csvstat", "--top", "2", str(path)],
        ["go", "run", "./cmd/csvstat", "--column", "a", "--freq", "a", str(path)],
        ["go", "run", "./cmd/csvstat", "--unknown", str(path)],
    ]

    for command in cases:
        result = subprocess.run(command, capture_output=True, text=True)
        print(result.stdout)
        print(result.stderr)
        assert result.returncode == 2
        assert result.stdout == ""
=== END AC cli-contract-usage ===

=== AC cli-contract-dispatch ===
Intent: Valid profile and frequency command forms dispatch successfully.

import pathlib
import subprocess
import tempfile

with tempfile.TemporaryDirectory() as directory:
    path = pathlib.Path(directory) / "data.csv"
    path.write_text("a,b\n1,x\n2,y\n", encoding="utf-8")

    profile = subprocess.run(
        ["go", "run", "./cmd/csvstat", str(path)],
        capture_output=True,
        text=True,
    )
    print(profile.stdout)
    print(profile.stderr)
    assert profile.returncode == 0
    assert profile.stdout != ""

    frequency = subprocess.run(
        ["go", "run", "./cmd/csvstat", "--freq", "b", "--top", "1", str(path)],
        capture_output=True,
        text=True,
    )
    print(frequency.stdout)
    print(frequency.stderr)
    assert frequency.returncode == 0
    assert frequency.stdout != ""
=== END AC cli-contract-dispatch ===

## User Acceptance

- None.

## Guardrails

- Options must precede the file argument.
- Exactly one file is accepted.
- Never write partial or diagnostic content to standard output on failure.
- Do not silently reinterpret unknown options.
