# FEATURE: Build Executable

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | The csvstat executable builds from the application root with go build using only the standard library. |
| Depends On  | ARCHITECTURE.md |
| Provides    | csvstat executable |
| Consumes    | internal/cli.Run |

## Purpose

`go build -o <path> ./cmd/csvstat` succeeds from the application root and produces the `csvstat` executable. `main` passes `os.Args[1:]`, `os.Stdout`, and `os.Stderr` to `cli.Run` and exits with the returned code. The supplied grader builds the executable this way.

## Trigger

A developer or the grader runs `go build -o <path> ./cmd/csvstat`.

## Behavior

- The module is Go 1.22 or newer and imports only the standard library and its own packages.
- `main` is the only caller of `os.Exit`.
- The built binary is runnable and exits with the code `cli.Run` returns.

## Programmatic Acceptance

=== AC build-go-build ===
Intent: go build of ./cmd/csvstat succeeds and writes an executable file.
Requires: executable=go; scope=test

import os
import subprocess
import sys
import tempfile

with tempfile.TemporaryDirectory() as work:
    target = os.path.join(work, "csvstat")
    result = subprocess.run(
        ["go", "build", "-o", target, "./cmd/csvstat"],
        capture_output=True,
        text=True,
    )
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
    assert os.path.isfile(target)
    assert os.access(target, os.X_OK)
=== END AC build-go-build ===

=== AC build-standard-library-only ===
Intent: Every non-test Go source file imports only standard library packages or packages of this module.

import os
import re

with open("go.mod", encoding="utf-8") as handle:
    gomod = handle.read()
module = re.search(r"^module\s+(\S+)", gomod, re.MULTILINE).group(1)

import_line = re.compile(r'^\s*(?:[A-Za-z_.]+\s+)?"([^"]+)"\s*$')
checked = 0
for root, dirs, files in os.walk("."):
    dirs[:] = [d for d in dirs if d not in ("sources", ".git", "vendor")]
    for name in files:
        if not name.endswith(".go") or name.endswith("_test.go"):
            continue
        in_block = False
        with open(os.path.join(root, name), encoding="utf-8") as handle:
            for line in handle:
                stripped = line.strip()
                if stripped.startswith("import ("):
                    in_block = True
                    continue
                if in_block and stripped == ")":
                    in_block = False
                    continue
                single = re.match(r'^import\s+(?:[A-Za-z_.]+\s+)?"([^"]+)"', stripped)
                path = None
                if single:
                    path = single.group(1)
                elif in_block:
                    found = import_line.match(line)
                    if found:
                        path = found.group(1)
                if path is None:
                    continue
                checked += 1
                first = path.split("/")[0]
                assert "." not in first or path.startswith(module), (name, path)
assert checked > 0
=== END AC build-standard-library-only ===

=== AC build-binary-runs ===
Intent: The built executable runs and returns an exit code in the documented range 0 to 3.

from drydock_go import cli_runner

code, out, err = cli_runner()
print(out)
print(err)
assert code in (0, 1, 2, 3)
=== END AC build-binary-runs ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no third-party modules.
- The executable's main package is `./cmd/csvstat`.
- Do not modify any file under `sources/`.
