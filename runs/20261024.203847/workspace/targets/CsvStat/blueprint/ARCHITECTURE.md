# ARCHITECTURE: CsvStat

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Module layout, package boundaries, and exit-code contract for the csvstat command-line CSV profiler. |
| Depends On  | None |
| Provides    | cli.Run, cmd/csvstat |

## Purpose

csvstat is a single-pass command-line tool. It reads one CSV file, validates it, infers column types, and writes either a column profile or a value-frequency list as CSV on standard output. It has no persistence, network use, or configuration file.

## Technology Stack

Derived from `TECHNOLOGY_STACK.md`:

- Go 1.22 or newer, standard library only, no third-party modules. The root `go.mod` declares no requirements.
- `go test` with `go vet` behind `bin/test.sh`.
- CSV read and written through `encoding/csv`.
- POSIX sh for `bin/test.sh` and the supplied grader `sources/run.sh`.

## Module Layout

```text
go.mod                    module csvstat, go 1.22 or newer, no require lines
cmd/csvstat/main.go       package main; calls cli.Run and passes the result to os.Exit
internal/cli              Run(args []string, stdout, stderr io.Writer) int
internal/reader           Read, Trim, IsNull, rejection checks
internal/typing           Infer, IsDate
internal/profile          Counts, MinMax, NumericStats
internal/freq             Top
internal/output           WriteProfile, WriteFreq
bin/test.sh               go vet ./... then go test ./...
```

Only `main` calls `os.Exit`. `cli.Run` uses a local `flag.FlagSet` with `flag.ContinueOnError` and is tested with `bytes.Buffer` writers.

## Exit Codes

| Code | Meaning |
|---|---|
| 0 | Success |
| 1 | The file cannot be opened or read |
| 2 | Usage error |
| 3 | Rejected input: empty file, empty or repeated header name, ragged record, or an absent column name |

A failing command writes nothing to standard output and one or more diagnostic lines to standard error. Output is buffered and written only on success.

## Module Ownership

| Boundary | Owner | Notes |
|---|---|---|
| Persistence | none | No persistent state exists. |
| Configuration | none | Behavior is set only by command-line options. |
| File access | internal/cli | Opens the named file read-only and hands a reader to internal/reader. |
| External services | none | No network use. |

## Constraints

- Statistics (sum, mean, median) are computed with `math/big` exact values and rounded once, half away from zero, when written.
- Application code never reads `sources/run.sh` or any supplied data file and never embeds an expected value.

## Programmatic Acceptance

=== AC architecture-module-layout ===
Intent: The application root is a Go 1.22+ module with no requirements and the declared entry point and CLI package directory.
Requires: executable=go; scope=test

from pathlib import Path
import re

mod = Path("go.mod").read_text()
found = re.search(r"^go (\d+)\.(\d+)", mod, re.M)
assert found is not None, mod
assert (int(found.group(1)), int(found.group(2))) >= (1, 22)
assert re.search(r"^\s*require\b", mod, re.M) is None
assert Path("cmd/csvstat/main.go").is_file()
assert Path("internal/cli").is_dir()
=== END AC architecture-module-layout ===

=== AC architecture-executable-builds ===
Intent: The cmd/csvstat main package builds and the built executable runs to a normal process exit.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner()
print(out)
print(err)
assert isinstance(code, int)
assert code >= 0
=== END AC architecture-executable-builds ===

=== AC architecture-only-main-exits ===
Intent: No non-test Go source outside cmd/ calls os.Exit.

from pathlib import Path

offenders = []
for path in Path("internal").rglob("*.go"):
    if path.name.endswith("_test.go"):
        continue
    if "os.Exit(" in path.read_text():
        offenders.append(str(path))
assert offenders == [], offenders
=== END AC architecture-only-main-exits ===

## User Acceptance

- None.

## Guardrails

- Do not add third-party modules; `go.mod` carries no `require` lines.
- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- Do not read `sources/run.sh` or any supplied data file from application code, and do not hardcode a statistic, a type, or a value list.
- Diagnostics go to standard error and are never part of the compared output.
