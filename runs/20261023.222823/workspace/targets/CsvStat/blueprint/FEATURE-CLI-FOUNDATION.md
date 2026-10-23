# FEATURE: CLI Foundation

| Field       | Value |
|-------------|-------|
| Version     | 20261023 V1 |
| Description | Provides the csvstat executable and documented command invocation surface. |
| Depends On  | ARCHITECTURE.md |
| Provides    | ./cmd/csvstat, csvstat help |
| Consumes    | csvstat architecture |

## Intent

Provide a Go module and executable at `./cmd/csvstat` with profile and frequency invocation forms.

## Invocation

- `csvstat [--column <name>] <file>`
- `csvstat --freq <name> [--top <n>] <file>`

The command supports useful help text for these forms and keeps process termination in `main`.

## Programmatic Acceptance

=== AC cli-foundation-build ===
Intent: The declared command package builds successfully.
import subprocess

result = subprocess.run(
    ["go", "build", "-o", "csvstat-ac", "./cmd/csvstat"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC cli-foundation-build ===

=== AC cli-foundation-help ===
Intent: The executable provides a successful help invocation.
import subprocess

result = subprocess.run(
    ["./csvstat-ac", "--help"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC cli-foundation-help ===

=== AC cli-foundation-usage ===
Intent: A missing file argument is reported as a usage failure.
import subprocess

result = subprocess.run(
    ["./csvstat-ac"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 2
=== END AC cli-foundation-usage ===

## User Acceptance

- None.

## Guardrails

- The executable is a Go main package at `./cmd/csvstat`.
- Only `main` calls `os.Exit`.
- No third-party modules are introduced.
