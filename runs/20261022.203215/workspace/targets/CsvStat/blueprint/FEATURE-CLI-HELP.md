# FEATURE: CLI Help

| Field       | Value |
|-------------|-------|
| Version     | 20261022 V1 |
| Description | Provides concise usage assistance and keeps diagnostics on standard error. |
| Depends On  | FEATURE-CLI-CONTRACT.md |
| Provides    | csvstat usage and diagnostic output |
| Consumes    | csvstat argument parser and command dispatch |

## Help Behavior

The conventional `--help` option succeeds and prints concise usage information to standard error.
The help describes both command forms, `--column`, `--freq`, `--top`, the required file argument,
and the principal exit-code categories.

Usage failures also write a diagnostic and usage text to standard error. Successful CSV output is
never mixed with help or diagnostics.

## Programmatic Acceptance

=== AC cli-help-option ===
Intent: The help option is accepted and does not emit help text on standard output.

import subprocess

result = subprocess.run(
    ["go", "run", "./cmd/csvstat", "--help"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
assert result.stdout == ""
assert result.stderr != ""
=== END AC cli-help-option ===

=== AC cli-help-usage-failure ===
Intent: A usage failure returns the usage exit status and keeps all assistance off standard output.

import subprocess

result = subprocess.run(
    ["go", "run", "./cmd/csvstat", "--help", "--unknown"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 2
assert result.stdout == ""
assert result.stderr != ""
=== END AC cli-help-usage-failure ===

## User Acceptance

- None.

## Guardrails

- Help and diagnostics are written only to standard error.
- Help must not cause a successful CSV report to be emitted.
- Keep usage text concise and aligned with the supported command forms.
