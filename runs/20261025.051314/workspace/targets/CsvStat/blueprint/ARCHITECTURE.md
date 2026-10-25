# ARCHITECTURE: CsvStat

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Module layout, package ownership, entry contract, and exit-code policy for the csvstat command-line CSV profiler. |
| Depends On  | None |
| Provides    | go.mod, internal/cli, internal/csvdata, internal/profile, internal/freq, cmd/csvstat |

## Purpose

csvstat is a single Go command-line executable. It reads one CSV file, infers a type for each column,
and writes either a column profile or a frequency listing to standard output as CSV. Failures write one
diagnostic line to standard error and nothing to standard output.

## Technology Stack

Derived from `TECHNOLOGY_STACK.md`; that file remains the decision of record.

| Technology | Where it applies |
|---|---|
| Go 1.22 or newer, standard library only | The whole application; `go.mod` declares no third-party requirement |
| `go test` with `go vet` | Project test suite behind `bin/test.sh` |
| CSV (RFC 4180, `encoding/csv`) | Input reading and output writing |
| POSIX shell | `bin/test.sh` and the supplied grader `sources/run.sh` |

## Module Layout

```text
go.mod                    module csvstat, go 1.22 or newer, no require directive
cmd/csvstat/main.go       package main; calls cli.Run and passes the result to os.Exit
internal/cli              argument parsing, dispatch, exit-code mapping, diagnostics
internal/csvdata          CSV reading, trimming, null recognition, input rejection
internal/profile          type inference, counts, min/max, sum/mean/median, decimal formatting
internal/freq             frequency counting, ordering, limiting
bin/test.sh               POSIX test gate (go vet ./... then go test ./...)
sources/                  read-only scoring assets; never modified
```

Each `internal/` package carries a `doc.go` declaring its package so the layout exists before any
behavior is built.

## Entry Contract

`cli.Run(args []string, stdout, stderr io.Writer) int` is the single entry point. `args` excludes the
program name. Only `cmd/csvstat/main.go` calls `os.Exit`; no package under `internal/` does. Tests call
`Run` with `bytes.Buffer` writers.

Standard output is written only on success: output is built in memory and flushed after every check
has passed.

## Exit Codes

| Code | Meaning |
|---|---|
| 0 | Success |
| 1 | Operational failure: the file cannot be opened or read |
| 2 | Usage error: missing or extra file, unknown option, missing option value, non-positive or malformed `--top`, `--top` without `--freq`, `--column` with `--freq` |
| 3 | Rejected input: empty file, empty or repeated header name, ragged or malformed record, or a column name absent from the header |

## Package Ownership

| Boundary | Owning package | Allowed low-level access |
|---|---|---|
| Command line, exit codes, stderr diagnostics | `internal/cli` | `flag`, `io`, `os.Open` for the input file |
| CSV parsing, trimming, null spellings, rejection | `internal/csvdata` | `encoding/csv`, `io`, `strings` |
| Type inference and statistics | `internal/profile` | `math/big`, `time`, `regexp`, `sort` |
| Frequency counting | `internal/freq` | `sort` |
| Configuration, file stores, external services, persistence | None; the application has no persistent state, no database, and no network access | Not applicable |

## Technical Decisions

- Sums, means, and medians are computed exactly with `math/big` and rounded once, half away from zero,
  when written.
- Application code never reads `sources/run.sh` or any supplied data file and never embeds a statistic,
  a type, or a value list. Every value comes from the file examined.
- Diagnostics go to standard error and are never compared by the grader.

## Programmatic Acceptance

=== AC arch-go-mod ===
Intent: The root go.mod declares a module, Go 1.22 or newer, and no third-party requirement.
Requires: executable=go; scope=test

import re
from pathlib import Path

text = Path("go.mod").read_text(encoding="utf-8")
lines = [line.strip() for line in text.splitlines()]
modules = [line for line in lines if line.startswith("module ")]
assert len(modules) == 1, lines
assert modules[0].split(None, 1)[1].strip() != ""
goline = [line for line in lines if re.match(r"^go\s+\d+\.\d+", line)]
assert len(goline) == 1, lines
match = re.match(r"^go\s+(\d+)\.(\d+)", goline[0])
assert (int(match.group(1)), int(match.group(2))) >= (1, 22)
assert not any(line.startswith("require") for line in lines), lines
=== END AC arch-go-mod ===

=== AC arch-layout ===
Intent: The application packages exist under internal/ and the command directory exists under cmd/.

from pathlib import Path

for package in ("cli", "csvdata", "profile", "freq"):
    directory = Path("internal") / package
    assert directory.is_dir(), directory
    assert list(directory.glob("*.go")), directory
assert Path("cmd/csvstat").is_dir()
=== END AC arch-layout ===

## User Acceptance

- None.

## Guardrails

- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- Application code never reads `sources/run.sh` or any supplied data file and never embeds an answer.
- Use only the Go standard library; add no third-party module.
- Only `cmd/csvstat/main.go` calls `os.Exit`.
