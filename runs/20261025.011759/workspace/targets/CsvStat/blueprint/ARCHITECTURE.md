# ARCHITECTURE: CsvStat

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Module layout, boundaries, and technical decisions for the csvstat command-line CSV column profiler. |
| Depends On  | None |
| Provides    | go.mod, cmd/csvstat, internal/cli.Run |
| Consumes    | None |

## Purpose

`csvstat` is a single-binary Go command-line tool. It reads one CSV file, infers a type for each column, and writes either a profile table or a frequency list to standard output as CSV. Diagnostics go to standard error. Exit codes are 0 (success), 1 (operational failure), 2 (usage error), and 3 (rejected input).

## Technology Stack

Derived from `TECHNOLOGY_STACK.md`, which remains the decision of record.

| Technology | Where it applies |
|---|---|
| Go 1.22 or newer, standard library only | Every package; no third-party modules and no `require` line in `go.mod` |
| go test and go vet | The project suite behind `bin/test.sh` |
| CSV (RFC 4180, `encoding/csv`) | Reading input and writing output |
| POSIX sh | `bin/test.sh` and the supplied grader `sources/run.sh` |

There is no database and no persistent state, so there is no `DATABASE.md`.

## Layout

```text
go.mod                    module declaration, Go 1.22 or newer, no dependencies
cmd/csvstat/main.go       process startup; the only caller of os.Exit
internal/cli/             Run(args []string, stdout, stderr io.Writer) int
internal/csvio/           reading, trimming, and rejecting CSV input
internal/profile/         null rule, type inference, counts, statistics, frequencies
internal/output/          profile and frequency CSV writers
bin/test.sh               POSIX test runner
```

## Modules and Boundaries

| Module | Owns | Allowed to use |
|---|---|---|
| `cmd/csvstat` | Wiring `os.Args[1:]`, `os.Stdout`, `os.Stderr` into `cli.Run`; calling `os.Exit` | `internal/cli` |
| `internal/cli` | Argument parsing with a local `flag.FlagSet` (`flag.ContinueOnError`), exit-code mapping, diagnostics, orchestration | `csvio`, `profile`, `output` |
| `internal/csvio` | Opening and reading the file, trimming, structural rejection | `os`, `encoding/csv` |
| `internal/profile` | Null classification, typing, counts, exact arithmetic with `math/big`, frequencies | `math/big`, `time`, `sort` |
| `internal/output` | CSV writing through `encoding/csv` with `\n` line endings | `encoding/csv` |

Persistence, configuration, file-store, and external-service boundaries: none. The only file access is reading the one input path named on the command line, owned by `internal/csvio`.

## Technical Decisions

- Only `main` calls `os.Exit`; `Run` returns the exit code so tests drive it with `bytes.Buffer` writers.
- Nothing is written to standard output until the whole input has been read and validated, so a failing command writes nothing to standard output.
- Sums, means, and medians are exact (`big.Rat`) and rounded once, half away from zero, when written.
- Application code never reads `sources/run.sh` or any supplied data file and embeds no answer.
- Invalid UTF-8 is accepted; blank lines are skipped by `encoding/csv`.

## Programmatic Acceptance

=== AC arch-module-layout ===
Intent: The application root is a Go module of Go 1.22 or newer with the csvstat main package and the cli package present.
Requires: executable=go; scope=test

import os
import re

assert os.path.isfile("go.mod")
with open("go.mod", encoding="utf-8") as handle:
    gomod = handle.read()
match = re.search(r"^go\s+(\d+)\.(\d+)", gomod, re.MULTILINE)
assert match is not None
major, minor = int(match.group(1)), int(match.group(2))
assert (major, minor) >= (1, 22)
assert os.path.isfile(os.path.join("cmd", "csvstat", "main.go"))
assert os.path.isdir(os.path.join("internal", "cli"))
=== END AC arch-module-layout ===

=== AC arch-no-dependencies ===
Intent: go.mod declares no third-party requirement, so the standard library is the only dependency.

import re

with open("go.mod", encoding="utf-8") as handle:
    gomod = handle.read()
assert re.search(r"^\s*require\b", gomod, re.MULTILINE) is None
=== END AC arch-no-dependencies ===

=== AC arch-cli-run-compiles ===
Intent: The module builds and vets cleanly, so cli.Run and the main package compile together.
Requires: executable=go; scope=test

import subprocess
import sys

build = subprocess.run(["go", "build", "./..."], capture_output=True, text=True)
print(build.stdout)
print(build.stderr, file=sys.stderr)
assert build.returncode == 0

vet = subprocess.run(["go", "vet", "./..."], capture_output=True, text=True)
print(vet.stdout)
print(vet.stderr, file=sys.stderr)
assert vet.returncode == 0
=== END AC arch-cli-run-compiles ===

## User Acceptance

- None.

## Guardrails

- Standard library only; no third-party modules.
- Only `main` calls `os.Exit`.
- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- Do not read `sources/run.sh` or any supplied data file from application code, and do not hardcode a statistic, a type, or a value list.
- Diagnostics go to standard error and are never part of the output contract.
