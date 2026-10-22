# FEATURE: Project Test Gate

| Field       | Value |
|-------------|-------|
| Version     | 20261022 V1 |
| Description | Provide the project-wide Go vet and test verification command. |
| Depends On  | FEATURE-COLUMN-STATISTICS.md, FEATURE-FREQUENCY.md, FEATURE-CLI-HELP.md |
| Provides    | bin/test.sh project verification command |
| Consumes    | profile and frequency implementation |

## Purpose

Provide a POSIX-compatible project gate that runs static analysis and the complete Go test suite from the application root.

## Behavior

- `bin/test.sh` runs `go vet ./...`.
- `bin/test.sh` runs `go test ./...`.
- The script exits zero only when both commands succeed.
- The script preserves failure visibility and is executable from the project root.
- Focused Go tests cover parsing, ingestion, validation, inference, arithmetic, serialization, frequencies, and exit behavior.

## Programmatic Acceptance

=== AC test-gate-present ===
Intent: The project test gate exists at the declared path and is executable.
import os
from pathlib import Path

path = Path("bin/test.sh")
assert path.is_file()
assert os.access(path, os.X_OK)
=== END AC test-gate-present ===

=== AC test-gate-passes ===
Intent: The project test gate completes successfully for the implemented project.
import subprocess

result = subprocess.run(["sh", "bin/test.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC test-gate-passes ===

=== AC test-gate-builds ===
Intent: The declared csvstat package remains buildable by the Go toolchain.
import subprocess

result = subprocess.run(
    ["go", "build", "./cmd/csvstat"],
    capture_output=True, text=True
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC test-gate-builds ===

## User Acceptance

- None.

## Guardrails

- Do not modify files under `sources/`.
- The script must use the project toolchain and standard library only.
- A passing script must not mask failures from either `go vet` or `go test`.
