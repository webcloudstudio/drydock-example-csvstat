# ARCHITECTURE: CsvStat

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Module layout, package ownership, and boundaries of the csvstat command-line CSV column profiler. |
| Depends On  | None |
| Provides    | go.mod, cmd/csvstat, internal/cli.Run |
| Consumes    | None |

## Purpose

CsvStat is a single-binary command-line profiler written in Go 1.22 or newer using only the
standard library. It reads one CSV file, infers each column's type, reports counts, extremes, and
for numeric columns the sum, mean, and median, and can list the most frequent values of one column.
There is no persistence and no network I/O.

## Technology Stack

Derived from `TECHNOLOGY_STACK.md`:

| Technology | Where it applies |
|---|---|
| Go 1.22 or newer | The whole application; standard library only, no third-party modules. |
| go test and go vet | The test runner behind `bin/test.sh`. |
| CSV (RFC 4180 via `encoding/csv`) | Input reading and output writing. |
| POSIX sh | `bin/test.sh` and the supplied grader `sources/run.sh`. |

## Module Layout

```text
go.mod                    # module path, go 1.22 or newer, no require lines
cmd/csvstat/main.go       # package main; calls internal/cli.Run and passes the code to os.Exit
internal/cli/             # argument parsing, usage text, dispatch, exit-code mapping
internal/csvio/           # file opening, CSV reading, trimming, validation
internal/profile/         # null rules, type inference, counts, statistics, frequencies
internal/decimal/         # exact arithmetic with math/big and two-decimal formatting
bin/test.sh               # POSIX sh: go vet ./... then go test ./...
```

`internal/cli.Run(args []string, stdout, stderr io.Writer) int` is the single entry point. Only
`cmd/csvstat/main.go` calls `os.Exit`. Behavior is tested with `bytes.Buffer`.

## Module Ownership

| Boundary | Owning package | Low-level APIs it may access |
|---|---|---|
| Command-line arguments, usage text, exit codes | `internal/cli` | `flag`, `io` |
| File access and CSV parsing | `internal/csvio` | `os`, `encoding/csv`, `strings` |
| Column semantics (null, type, counts, stats, frequency) | `internal/profile` | `time`, `sort`, `bytes` |
| Exact numbers and rounding | `internal/decimal` | `math/big` |
| Persistence, configuration files, external services | none | not applicable |

## Exit Codes

| Code | Meaning |
|---|---|
| 0 | Success. |
| 1 | The file cannot be opened or read. |
| 2 | Usage error. |
| 3 | Rejected input: empty file, empty or repeated header name, ragged record, CSV syntax error, or a column name absent from the header. |

A failing command writes nothing to standard output and one diagnostic line to standard error.
Output is buffered until validation completes.

## Constraints

- Application code never reads `sources/run.sh` or any supplied data file and never embeds a
  statistic, type, or value list.
- Sums, means, and medians are exact (`math/big`) and rounded once, half away from zero, at output.
- Input bytes are compared as bytes; a byte-order mark is not stripped.

## Programmatic Acceptance

=== AC arch-module-builds ===
Intent: The application root is a Go module whose main package at ./cmd/csvstat builds with go build.
Requires: executable=go; scope=test

import subprocess
import sys
import tempfile
from pathlib import Path

assert Path("go.mod").is_file()
assert Path("cmd/csvstat/main.go").is_file()
with tempfile.TemporaryDirectory() as tmp:
    target = str(Path(tmp) / "csvstat")
    result = subprocess.run(["go", "build", "-o", target, "./cmd/csvstat"], capture_output=True, text=True)
    print(result.stdout)
    print(result.stderr, file=sys.stderr)
    assert result.returncode == 0
    assert Path(target).is_file()
=== END AC arch-module-builds ===

=== AC arch-stdlib-only ===
Intent: The module declares no third-party requirement and the main package depends only on this module and the standard library.
Requires: executable=go; scope=test

import subprocess
import sys
from pathlib import Path

go_mod = Path("go.mod").read_text()
module_line = [ln for ln in go_mod.splitlines() if ln.startswith("module ")]
assert len(module_line) == 1
module_path = module_line[0].split()[1]
assert "require" not in go_mod
result = subprocess.run(
    ["go", "list", "-deps", "-f", "{{if not .Standard}}{{.ImportPath}}{{end}}", "./cmd/csvstat"],
    capture_output=True, text=True,
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
for line in result.stdout.split():
    assert line.startswith(module_path), line
=== END AC arch-stdlib-only ===

=== AC arch-vet-clean ===
Intent: The foundation tree passes go vet.
Requires: executable=go; scope=test

import subprocess
import sys

result = subprocess.run(["go", "vet", "./..."], capture_output=True, text=True)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC arch-vet-clean ===

=== AC arch-only-main-exits ===
Intent: Only the main package calls os.Exit and the internal packages exist as declared.

from pathlib import Path

assert Path("internal/cli").is_dir()
exit_call = "os." + "Exit("
for path in Path("internal").rglob("*.go"):
    if path.name.endswith("_test.go"):
        continue
    assert exit_call not in path.read_text(), str(path)
assert exit_call in Path("cmd/csvstat/main.go").read_text()
=== END AC arch-only-main-exits ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no third-party module may appear in `go.mod`.
- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- Application code never reads `sources/run.sh` or a supplied data file and never hardcodes a statistic, type, or value list.
- No database and no network access.
