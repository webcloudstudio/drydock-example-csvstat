# ARCHITECTURE: CsvStat Application

| Field       | Value |
|-------------|-------|
| Version     | 20261022 V1 |
| Description | Defines the Go module, executable boundary, and internal ownership for the csvstat CLI. |
| Depends On  | — |
| Provides    | Go module, ./cmd/csvstat entry point, internal profiling packages |
| Consumes    | — |

## Purpose

CsvStat is a standard-library-only Go command-line application. The executable is built from
`./cmd/csvstat`; reusable behavior belongs in internal packages.

## Module and Boundaries

- `cmd/csvstat` owns process startup and is the only package permitted to call `os.Exit`.
- `internal/cli` owns argument parsing, command dispatch, diagnostics, and exit-code selection.
- `internal/csvstat` owns CSV reading, normalization, validation, type inference, statistics,
  frequencies, and CSV serialization.
- Input is read from the user-supplied file path only.
- Successful reports are written to standard output; diagnostics are written to standard error.

## Technology Stack

- Go 1.22 or newer.
- Standard library only.
- `encoding/csv` for RFC 4180 parsing and output.
- `math/big` or equivalent exact standard-library arithmetic for decimal calculations.
- POSIX shell for the project test launcher.

## Module Ownership

| Boundary | Owner | Allowed low-level APIs |
|---|---|---|
| Process startup | `cmd/csvstat` | `os`, `io`, CLI package |
| CLI behavior | `internal/cli` | `flag`, `os`, `io` |
| CSV and profiling domain | `internal/csvstat` | `encoding/csv`, `math/big`, `time`, standard library |
| Persistent state | None | No persistent state |
| Configuration | None | No configuration store |
| External services | None | No external services |

## Programmatic Acceptance

=== AC architecture-module ===
Intent: The application exposes the required Go module and executable source boundary.

import pathlib
import subprocess

root = pathlib.Path(".")
assert (root / "go.mod").is_file()
assert (root / "cmd" / "csvstat").is_dir()

result = subprocess.run(
    ["go", "build", "./cmd/csvstat"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC architecture-module ===

## User Acceptance

- None.

## Guardrails

- Use only the Go standard library.
- Do not place application logic in `cmd/csvstat`.
- Do not add persistent storage, network services, or third-party dependencies.
- Only the process entry point may terminate the process directly.
