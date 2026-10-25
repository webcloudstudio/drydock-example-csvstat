# FEATURE: Executable

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Build the csvstat executable from the module root with a thin main package over cli.Run. |
| Depends On  | ARCHITECTURE.md |
| Provides    | csvstat command, cli.Run |
| Consumes    | None |

## Purpose

`go build ./cmd/csvstat` at the application root produces the `csvstat` executable using only the
standard library. This story delivers the entry point and a minimal `Run` that later stories extend.

## Behavior

- `cmd/csvstat/main.go` is `package main`. It calls `cli.Run(os.Args[1:], os.Stdout, os.Stderr)` and
  passes the returned code to `os.Exit`. It contains no other logic.
- `internal/cli` exports `Run(args []string, stdout, stderr io.Writer) int`. In this story it returns 0.
  Later stories add parsing, dispatch, and failure mapping.
- The import path of the cli package is `<module>/internal/cli`, where the module path is the one in
  `go.mod`.
- Every import in the module is either standard library or module-local.

## Tests For The Project Suite

The builder adds `internal/cli` tests that call `Run` with `bytes.Buffer` writers.

## Programmatic Acceptance

=== AC exe-builds ===
Intent: The csvstat executable builds with go build and runs to a defined exit code on a valid file.
Requires: executable=go; scope=test

import csv
import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as workdir:
    path = os.path.join(workdir, "data.csv")
    with open(path, "w", newline="", encoding="utf-8") as handle:
        writer = csv.writer(handle, lineterminator="\n")
        writer.writerow(["a", "b"])
        writer.writerow(["1", "2"])
    code, out, err = cli_runner(path)
    print(code, out, err)
    assert isinstance(code, int)
    assert 0 <= code <= 3
    assert isinstance(out, str)
=== END AC exe-builds ===

=== AC exe-stdlib-only ===
Intent: Every import in the module is standard library or module-local, with no third-party path.

import re
from pathlib import Path

files = list(Path("cmd").rglob("*.go")) + list(Path("internal").rglob("*.go"))
assert files
paths = []
for source in files:
    text = source.read_text(encoding="utf-8")
    blocks = re.findall(r"import\s*\((.*?)\)", text, re.S)
    blocks += re.findall(r"^import\s+(?:\w+\s+)?(\"[^\"]+\")", text, re.M)
    for block in blocks:
        paths += re.findall(r"\"([^\"]+)\"", block)
assert paths
for path in paths:
    first = path.split("/")[0]
    assert "." not in first, path
=== END AC exe-stdlib-only ===

=== AC exe-main-owns-exit ===
Intent: Only the main package calls os.Exit and main delegates to cli.Run.

from pathlib import Path

main_source = Path("cmd/csvstat/main.go").read_text(encoding="utf-8")
assert "package main" in main_source
assert "os.Exit(" in main_source
assert "cli.Run(" in main_source
for source in Path("internal").rglob("*.go"):
    assert "os.Exit(" not in source.read_text(encoding="utf-8"), source
=== END AC exe-main-owns-exit ===

=== AC exe-no-grader-read ===
Intent: Application code never references the supplied grader script.

from pathlib import Path

files = list(Path("cmd").rglob("*.go")) + list(Path("internal").rglob("*.go"))
for source in files:
    if source.name.endswith("_test.go"):
        continue
    assert "run.sh" not in source.read_text(encoding="utf-8"), source
=== END AC exe-no-grader-read ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no third-party module.
- Application code never reads `sources/run.sh` or any supplied data file.
- Do not modify anything under `sources/`.
