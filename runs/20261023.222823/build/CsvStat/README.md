# drydock CsvStat Uat Case

Drydock uat - CsvStat builds a Go CSV column profiler. The source is an original brief: infer each column's type (integer, number, date, boolean, text, empty), report rows, nulls, distinct, min, max, sum, mean, and median, and list a column's most frequent values. The supplied grader, sources/run.sh, builds ./cmd/csvstat with go build and runs 22 checks of exit code and output against five seeded random CSV files with known answers: quoted commas, null spellings, padded text, an invalid calendar date, a header-only file, and a ragged file. Standard library only, no database. Run with drydock uat CsvStat. Source Repo: none, the brief is original to this kit.

## Intent

Build a command-line CSV column profiler, the Go executable `./cmd/csvstat`, as described in
`sources/INSTRUCTIONS.md`. It infers each column's type, reports counts, extremes, and for numeric
columns sum, mean, and median, and lists a column's most frequent values.

Correctness is measured by `sources/run.sh`: it builds the executable with `go build` and makes
22 checks of exit code and standard output against five supplied CSV files. The goal is every
check passing.

## What It Does

The command reports failures through exit status and diagnostics. It must never write profile or frequency CSV to standard output after a failure.

The frequency command accepts `--freq <name> [--top <n>] <file>`. The default limit is five. `--top` must be a positive integer and is valid only with `--freq`. The limit is applied after frequency ordering.

The frequency capability counts each distinct non-null normalized value in the requested column. It excludes empty values and the case-insensitive null spellings `NA`, `N/A`, and `NULL`.

Rows are ordered by descending count and then by bytewise ascending value.

The profile renderer writes one CSV header followed by one row per selected column in input order. A named column limits output to that column. Empty statistic values remain empty CSV fields, and fields containing commas, quotes, or newlines are encoded with RFC-compliant CSV quoting.

This capability calculates profile statistics after input normalization and type inference.

For every column it reports row count, null count, distinct non-null values, type-specific minimum and maximum, and numeric sum, mean, and median. Numeric calculations use exact arithmetic and are rounded once, half away from zero, to two decimal places when rendered.

Integer values remain plain integers. Non-numeric sum, mean, and median fields remain empty. Boolean extrema are false when any false value exists and true only when all non-null values are true. Text extrema use bytewise ordering.

## Stack

Go

## Verification

```bash
bash bin/test.sh
```

## Next Steps

- Complete the first-run checks above.
- Run `bash bin/test.sh` before making release or deployment changes.
