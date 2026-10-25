# FEATURE: Executable Build

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | The csvstat executable builds from the application root with go build using only the standard library. |
| Depends On  | ARCHITECTURE.md |
| Provides    | csvstat executable |
| Consumes    | internal/cli.Run |

## Purpose

`go build -o <path> ./cmd/csvstat` succeeds from the module root and produces a runnable `csvstat` binary. The main package is a thin shell: it calls `cli.Run(os.Args[1:], os.Stdout, os.Stderr)` and passes the result to `os.Exit`.

## Trigger

A developer or the grader runs `go build -o <temporary path> ./cmd/csvstat` from the application root.

## Behavior

- `cmd/csvstat/main.go` is `package main` and contains the only `os.Exit` call in the module.
- A minimal `cli.Run` exists so the build succeeds before later stories add behavior.
- `go.mod` declares no `require` entries.

## Reads

- `go.mod`, `cmd/csvstat/main.go`, `internal/cli`.

## Writes

- The built binary at the path given to `-o`.

## Programmatic Acceptance

=== AC build-produces-binary ===
Intent: go build of ./cmd/csvstat from the application root exits 0 and writes an executable file.
Requires: executable=go; scope=test

import os
import subprocess
import tempfile

with tempfile.TemporaryDirectory() as work:
    target = os.path.join(work, "csvstat")
    result = subprocess.run(
        ["go", "build", "-o", target, "./cmd/csvstat"], capture_output=True, text=True
    )
    print(result.stdout)
    print(result.stderr)
    assert result.returncode == 0
    assert os.path.isfile(target)
    assert os.access(target, os.X_OK)
=== END AC build-produces-binary ===

=== AC build-no-third-party-requires ===
Intent: go.mod lists no requirements, so the build uses only the standard library.

import pathlib

lines = [ln.strip() for ln in pathlib.Path("go.mod").read_text().splitlines()]
assert not any(ln.startswith("require") for ln in lines)
assert not any(ln.startswith("replace") for ln in lines)
=== END AC build-no-third-party-requires ===

=== AC build-main-is-thin ===
Intent: Only cmd/csvstat calls os.Exit; no internal package does.

import pathlib

main_text = pathlib.Path("cmd/csvstat/main.go").read_text()
assert "package main" in main_text
assert "os.Exit(" in main_text
for path in pathlib.Path("internal").rglob("*.go"):
    if path.name.endswith("_test.go"):
        continue
    assert "os.Exit(" not in path.read_text(), str(path)
=== END AC build-main-is-thin ===

=== AC build-vet-clean ===
Intent: go vet over the whole module exits 0.
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(["go", "vet", "./..."], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC build-vet-clean ===

## User Acceptance

- None.

## Guardrails

- Standard library only.
- Do not modify anything under `sources/`.
- `main` contains no CSV logic; behavior lives in `internal/` packages.
