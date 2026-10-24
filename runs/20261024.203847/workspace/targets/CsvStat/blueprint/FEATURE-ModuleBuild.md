# FEATURE: Module Build

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | The application root builds the csvstat executable with go build using only the standard library. |
| Depends On  | ARCHITECTURE.md |
| Provides    | go build ./cmd/csvstat |

## Purpose

`go build ./cmd/csvstat` run from the application root produces the `csvstat` executable for the host platform. The executable has no runtime dependency beyond itself, and the module uses only the Go standard library.

## Behavior

- The root `go.mod` declares no `require` directives and no `replace` directives.
- No `go.sum` file is needed and none is committed.
- `go build -o <path> ./cmd/csvstat` succeeds from the application root and writes an executable file.
- The built binary starts and exits normally with no arguments.

## Programmatic Acceptance

=== AC modulebuild-go-build-succeeds ===
Intent: go build of ./cmd/csvstat from the application root succeeds and writes an executable file.
Requires: executable=go; scope=test

import os
import subprocess
import tempfile
from pathlib import Path

with tempfile.TemporaryDirectory() as tmp:
    target = Path(tmp) / "csvstat-built"
    result = subprocess.run(
        ["go", "build", "-o", str(target), "./cmd/csvstat"],
        capture_output=True, text=True,
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
    assert target.is_file()
    assert os.access(target, os.X_OK)
=== END AC modulebuild-go-build-succeeds ===

=== AC modulebuild-standard-library-only ===
Intent: The module declares no requirements or replacements and has no checksum file.

from pathlib import Path
import re

mod = Path("go.mod").read_text()
assert re.search(r"^\s*require\b", mod, re.M) is None
assert re.search(r"^\s*replace\b", mod, re.M) is None
assert not Path("go.sum").exists()
=== END AC modulebuild-standard-library-only ===

=== AC modulebuild-binary-runs ===
Intent: The built executable starts and ends with a normal process exit status.
Requires: executable=go; scope=test

import subprocess
import tempfile
from pathlib import Path

with tempfile.TemporaryDirectory() as tmp:
    target = Path(tmp) / "csvstat-built"
    build = subprocess.run(
        ["go", "build", "-o", str(target), "./cmd/csvstat"],
        capture_output=True, text=True,
    )
    print(build.stderr)
    assert build.returncode == 0
    run = subprocess.run([str(target)], capture_output=True, text=True)
    print(run.stdout)
    print(run.stderr)
    assert run.returncode >= 0
=== END AC modulebuild-binary-runs ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no third-party modules.
- Do not modify anything under `sources/`.
