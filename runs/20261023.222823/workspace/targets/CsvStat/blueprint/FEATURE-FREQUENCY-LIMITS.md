# FEATURE: Frequency Limits

| Field       | Value |
|-------------|-------|
| Version     | 20261023 V1 |
| Description | Validate frequency limits and render bounded frequency output. |
| Depends On  | FEATURE-FREQUENCY-VALUES.md |
| Provides    | --top validation, default frequency limit, frequency CSV output |
| Consumes    | frequency aggregation and deterministic ordering |

## Purpose

The frequency command accepts `--freq <name> [--top <n>] <file>`. The default limit is five. `--top` must be a positive integer and is valid only with `--freq`. The limit is applied after frequency ordering.

## Behavior

- Emit the CSV header `value,count`.
- Emit no more than the requested number of frequency rows.
- Reject zero, negative, or non-integer `--top` values with usage exit status 2.
- Reject `--top` without `--freq` with usage exit status 2.
- Preserve RFC-compliant CSV rendering.

## Programmatic Acceptance

=== AC frequency-limit-success ===
Intent: A valid positive limit executes the frequency command successfully and bounds the parsed result.
import csv
import io
import subprocess

column = "department"
limit = "2"
result = subprocess.run(
    ["./csvstat", "--freq", column, "--top", limit, "sources/employees.csv"],
    capture_output=True,
    text=True,
)
assert result.returncode == 0
rows = list(csv.reader(io.StringIO(result.stdout)))
assert rows[0] == ["value", "count"]
assert len(rows) <= int(limit) + 1
=== END AC frequency-limit-success ===

=== AC frequency-limit-default ===
Intent: Omitting --top uses the default limit and returns a successful frequency result.
import csv
import io
import subprocess

result = subprocess.run(
    ["./csvstat", "--freq", "department", "sources/employees.csv"],
    capture_output=True,
    text=True,
)
assert result.returncode == 0
rows = list(csv.reader(io.StringIO(result.stdout)))
assert rows[0] == ["value", "count"]
assert len(rows) <= 6
=== END AC frequency-limit-default ===

=== AC frequency-limit-invalid ===
Intent: Non-positive and malformed limits are rejected as usage errors.
import subprocess

cases = [
    ["--freq", "department", "--top", "0", "sources/employees.csv"],
    ["--freq", "department", "--top", "-1", "sources/employees.csv"],
    ["--freq", "department", "--top", "x", "sources/employees.csv"],
]
for args in cases:
    result = subprocess.run(["./csvstat", *args], capture_output=True, text=True)
    assert result.returncode == 2
    assert result.stdout == ""
=== END AC frequency-limit-invalid ===

=== AC top-requires-frequency ===
Intent: --top without --freq is rejected as a usage error.
import subprocess

result = subprocess.run(
    ["./csvstat", "--top", "3", "sources/employees.csv"],
    capture_output=True,
    text=True,
)
assert result.returncode == 2
assert result.stdout == ""
=== END AC top-requires-frequency ===

## User Acceptance

- None.

## Guardrails

- Do not include null values in frequency output.
- Do not change the ordering supplied by frequency aggregation when applying the limit.
- Do not read or embed expected answers from supplied grading assets.
