# FEATURE: CLI Errors

| Field       | Value |
|-------------|-------|
| Version     | 20261023 V1 |
| Description | Map operational, usage, and rejected-input failures to the command contract. |
| Depends On  | FEATURE-CLI-COLUMNS.md, FEATURE-PROFILE-OUTPUT.md, FEATURE-FREQUENCY-LIMITS.md |
| Provides    | exit-code and diagnostic contract |
| Consumes    | CSV validation, column lookup, profile output, frequency output |

## Purpose

The command reports failures through exit status and diagnostics. It must never write profile or frequency CSV to standard output after a failure.

## Error Contract

- Exit 1 for a file that cannot be opened or read.
- Exit 2 for missing or extra files, unknown options, missing option values, invalid `--top`, `--top` without `--freq`, or incompatible options.
- Exit 3 for empty input, invalid or repeated headers, ragged records, or absent named columns.
- Write one diagnostic line to standard error for failures.
- Diagnostics identify the input file and identify the rejected record line when a record is structurally invalid.

## Programmatic Acceptance

=== AC operational-error ===
Intent: An unreadable input file returns the operational failure status and no CSV output.
import pathlib
import subprocess

missing = "sources/does-not-exist.csv"
result = subprocess.run(["./csvstat", missing], capture_output=True, text=True)
assert result.returncode == 1
assert result.stdout == ""
assert pathlib.Path(missing).name in result.stderr
=== END AC operational-error ===

=== AC usage-errors ===
Intent: Malformed command arguments return usage status and no CSV output.
import subprocess

cases = [
    ["./csvstat"],
    ["./csvstat", "sources/orders.csv", "sources/employees.csv"],
    ["./csvstat", "--unknown", "sources/orders.csv"],
    ["./csvstat", "--column"],
]
for command in cases:
    result = subprocess.run(command, capture_output=True, text=True)
    assert result.returncode == 2
    assert result.stdout == ""
=== END AC usage-errors ===

=== AC rejected-input ===
Intent: Structurally invalid input and absent columns return rejected-input status and no CSV output.
import subprocess

cases = [
    ["./csvstat", "sources/ragged.csv"],
    ["./csvstat", "--column", "missing", "sources/orders.csv"],
    ["./csvstat", "--freq", "missing", "sources/orders.csv"],
]
for command in cases:
    result = subprocess.run(command, capture_output=True, text=True)
    assert result.returncode == 3
    assert result.stdout == ""
    assert result.stderr.strip()
=== END AC rejected-input ===

=== AC successful-output-isolated ===
Intent: A valid command remains successful and emits CSV on standard output.
import subprocess

result = subprocess.run(
    ["./csvstat", "--column", "qty", "sources/orders.csv"],
    capture_output=True,
    text=True,
)
assert result.returncode == 0
assert result.stdout != ""
=== END AC successful-output-isolated ===

## User Acceptance

- None.

## Guardrails

- Never emit partial CSV output before input validation and command execution succeed.
- Keep diagnostics on standard error.
- Preserve the distinct exit statuses for operational, usage, and rejected-input failures.
