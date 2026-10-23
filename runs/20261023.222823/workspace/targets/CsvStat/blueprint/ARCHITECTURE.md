# ARCHITECTURE: CsvStat

| Field       | Value |
|-------------|-------|
| Version     | 20261023 V1 |
| Description | Architecture and module boundaries for the CsvStat command-line CSV profiler. |
| Depends On  | — |
| Provides    | csvstat architecture |
| Consumes    | — |

## Purpose

CsvStat is a standard-library Go command-line pipeline. It reads one RFC 4180 CSV file, normalizes records, validates structure, profiles columns or reports frequencies, renders CSV output, and maps failures to documented exit codes.

## Modules and Boundaries

- `cmd/csvstat`: process startup, argument forwarding, and final `os.Exit`.
- `internal/cli`: flag parsing, command dispatch, diagnostics, and exit-code mapping.
- `internal/csvstat`: CSV loading, normalization, type inference, statistics, frequency aggregation, and rendering.
- `bin/test.sh`: project verification launcher.

Application logic returns values and errors through testable functions. Only the command entry point terminates the process.

## Technology Stack

- Go 1.22 or newer with the standard library only.
- `encoding/csv` for RFC 4180 parsing and output.
- POSIX shell for `bin/test.sh`.
- Exact numeric arithmetic uses standard-library facilities such as `math/big`.

## Contracts

Successful commands write CSV only to standard output. Diagnostics go to standard error. Failed commands write no standard output. Exit codes are 1 for operational failures, 2 for usage failures, and 3 for rejected input.

## Programmatic Acceptance

=== AC architecture-build ===
Intent: The application architecture builds as the declared Go command.
import subprocess

result = subprocess.run(
    ["go", "build", "./cmd/csvstat"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC architecture-build ===

=== AC architecture-vet ===
Intent: The declared Go packages pass static analysis.
import subprocess

result = subprocess.run(
    ["go", "vet", "./..."],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC architecture-vet ===

## User Acceptance

- None.

## Guardrails

- Use only the Go standard library.
- Never read or modify files under `sources/` from application code.
- Keep diagnostics off standard output.
