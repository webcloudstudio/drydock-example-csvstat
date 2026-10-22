# COMPASS: CsvStat

## Compass

Build a command-line CSV column profiler, the Go executable `./cmd/csvstat`, as described in
`sources/INSTRUCTIONS.md`. It infers each column's type, reports counts, extremes, and for numeric
columns sum, mean, and median, and lists a column's most frequent values.

Correctness is measured by `sources/run.sh`: it builds the executable with `go build` and makes
22 checks of exit code and standard output against five supplied CSV files. The goal is every
check passing.

## Constraints

- Implement in Go 1.22 or newer with the standard library only. No third-party modules.
- The application root is a Go module; the executable's main package is `./cmd/csvstat`.
- Compute sums, means, and medians exactly (for example with `math/big`) and round once, half away
  from zero, when writing.
- Diagnostics go to standard error and are never compared.

## Guardrails

- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- Do not read `sources/run.sh` or any supplied data file from application code, and do not
  hardcode a statistic, a type, or a value list. Every value comes from the file examined.

## Verification Protocol

| Story kind | Command | Asserts |
|---|---|---|
| Every behavioral story | `sh bin/test.sh` | exit `0` |
| Terminal story (once, last) | `sh sources/run.sh` | exit `0` |
