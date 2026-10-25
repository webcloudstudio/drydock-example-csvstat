# ARCHITECTURE: CsvStat

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Module layout, package boundaries, exit-code contract, and technical decisions for the csvstat command-line CSV profiler. |
| Depends On  | — |
| Provides    | go.mod, internal/cli.Run |
| Consumes    | — |

## Purpose

CsvStat is a single Go executable, `./cmd/csvstat`, that profiles one CSV file. It infers each column's type, reports counts, extremes, and for numeric columns sum, mean, and median, and lists the most frequent values of one column. There is no database, no network, and no persistent state.

## Technology Stack

Derived from `TECHNOLOGY_STACK.md`, which remains the decision of record.

| Technology | Where it applies |
|---|---|
| Go 1.22 or newer, standard library only | All application code; `go.mod` declares no `require` entries |
| go test and go vet | Test runner behind `bin/test.sh` |
| CSV via `encoding/csv` | Reading input (RFC 4180) and writing every output form |
| POSIX sh | `bin/test.sh` and the supplied grader `sources/run.sh` |

## Module Layout

```text
go.mod
cmd/csvstat/main.go        package main; calls cli.Run and passes the result to os.Exit
internal/cli/              Run(args []string, stdout, stderr io.Writer) int; option parsing, dispatch, exit codes
internal/csvdata/          reading, trimming, null classification, structural validation
internal/infer/            type inference ladder
internal/profile/          rows, nulls, distinct, min, max, sum, mean, median
internal/freq/             frequency listing
internal/decimal/          exact big.Rat formatting, one rounding, half away from zero
bin/test.sh                POSIX launcher: go vet ./... then go test ./...
```

## Module Ownership

| Boundary | Owning module | May access |
|---|---|---|
| Process start and exit | `cmd/csvstat` | `os.Args`, `os.Stdout`, `os.Stderr`, `os.Exit` (the only caller of `os.Exit`) |
| Option parsing, exit codes, diagnostics | `internal/cli` | A local `flag.FlagSet` with `flag.ContinueOnError`, the injected writers |
| File access and CSV decoding | `internal/csvdata` | `os.Open`, `encoding/csv` reader |
| Type rules | `internal/infer` | Trimmed non-null values only |
| Statistics | `internal/profile`, `internal/decimal` | `math/big`; no floating point for sums, means, or medians |
| Output encoding | `internal/cli` with `encoding/csv` writer | The injected stdout writer |
| Persistence, configuration, external services | None | The product has none |

## Exit-Code Contract

| Code | Meaning |
|---|---|
| 0 | Success |
| 1 | Operational failure: the file cannot be opened or read |
| 2 | Usage error: missing or extra file, unknown option, missing option value, `--top` not a positive integer, `--top` without `--freq`, `--column` with `--freq` |
| 3 | Rejected input: empty file, empty or repeated header name, ragged record, or a column name absent from the header |

A failing command writes nothing to standard output and exactly one diagnostic line to standard error naming the file and, for a rejected record, its line number. Standard error is never compared by the grader.

## Technical Decisions

- Application code never reads `sources/run.sh` or any supplied data file and never embeds a statistic, type, or value list. Every value is computed from the file it is given.
- Sums, means, and medians are exact (`math/big`) and rounded once, half away from zero, when written.
- All output goes through one `encoding/csv` writer, so quoting follows RFC 4180 and lines end with `\n`.
- Usage validation completes before the file is opened, so a usage error outranks a missing file.
- Only `main` calls `os.Exit`; `cli.Run` returns the exit code so tests drive it with `bytes.Buffer`.

## Programmatic Acceptance

=== AC arch-go-mod-stdlib-only ===
Intent: The application root is a Go module that declares no third-party requirements.
Requires: executable=go; scope=test

import pathlib

gomod = pathlib.Path("go.mod")
assert gomod.is_file()
lines = [ln.strip() for ln in gomod.read_text().splitlines()]
assert any(ln.startswith("module ") for ln in lines)
assert any(ln.startswith("go ") for ln in lines)
assert not any(ln.startswith("require") for ln in lines)
assert not pathlib.Path("go.sum").exists() or pathlib.Path("go.sum").stat().st_size == 0
=== END AC arch-go-mod-stdlib-only ===

=== AC arch-package-layout ===
Intent: Each internal package named by the architecture exists and holds Go source.

import pathlib

packages = ["cli", "csvdata", "infer", "profile", "freq", "decimal"]
for name in packages:
    directory = pathlib.Path("internal") / name
    assert directory.is_dir(), name
    sources = [p for p in directory.glob("*.go") if not p.name.endswith("_test.go")]
    assert sources, name
assert pathlib.Path("bin").is_dir()
=== END AC arch-package-layout ===

=== AC arch-cli-run-compiles ===
Intent: The internal packages, including the cli.Run entry contract, compile and vet cleanly.
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(["go", "build", "./internal/..."], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC arch-cli-run-compiles ===

=== AC arch-run-signature ===
Intent: internal/cli exposes Run taking args and two writers and returning an int exit code.

import pathlib

text = "\n".join(
    p.read_text()
    for p in pathlib.Path("internal/cli").glob("*.go")
    if not p.name.endswith("_test.go")
)
assert "func Run(" in text
assert "io.Writer" in text
=== END AC arch-run-signature ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no third-party module may appear in `go.mod`.
- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- Application code must not read `sources/run.sh` or any supplied data file, and must not hardcode a statistic, type, or value list.
- Only `cmd/csvstat` calls `os.Exit`.
- Sums, means, and medians are never computed with floating point.
