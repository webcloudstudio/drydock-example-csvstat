# FEATURE: Test Gate

| Field       | Value |
|-------------|-------|
| Version     | 20261023 V1 |
| Description | Provide the project test launcher and focused automated coverage. |
| Depends On  | FEATURE-CLI-FOUNDATION.md |
| Provides    | bin/test.sh |
| Consumes    | Go module and command packages |

## Test Gate

`bin/test.sh` runs `go vet ./...` and `go test ./...` from the application root. It exits zero only when both commands succeed. The repository test suite should cover parsing, normalization, inference, statistics, rendering, frequency behavior, and exit-code mapping.

## Programmatic Acceptance

=== AC test-gate-present ===
Intent: The project provides an executable POSIX test launcher.
import os
import pathlib

path = pathlib.Path("bin/test.sh")
assert path.is_file()
assert os.access(path, os.X_OK)
=== END AC test-gate-present ===

=== AC test-gate-runs ===
Intent: The project test launcher completes successfully.
import os
import subprocess

result = subprocess.run(
    ["sh", "bin/test.sh"],
    capture_output=True,
    text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC test-gate-runs ===

=== AC go-vet-and-tests ===
Intent: The launcher exercises both the Go vet and Go test commands.
import pathlib

text = pathlib.Path("bin/test.sh").read_text()
assert "go vet ./..." in text
assert "go test ./..." in text
=== END AC go-vet-and-tests ===

## User Acceptance

- None.

## Guardrails

- Keep the launcher POSIX-compatible.
- Run commands from the application root.
- Do not make the launcher depend on the supplied grader or hardcoded fixture answers.
