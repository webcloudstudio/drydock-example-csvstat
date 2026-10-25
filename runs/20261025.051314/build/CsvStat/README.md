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

csvstat accepts two command forms:

```text
csvstat [--column <name>] <file>
csvstat --freq <name> [--top <n>] <file>
```

Options precede the file and exactly one file is required. This story delivers the parser and its
wiring into `Run`. Cross-option validation (`--top` without `--freq`, `--column` with `--freq`,
non-positive `--top`) belongs to the usage-errors story.

Complete `profile.InferType` so every column receives exactly one type: the first of `integer`,
`number`, `date`, `boolean`, `text` that every non-null value satisfies, or `empty` when there is
no non-null value.

`csvstat --column <name> <file>` writes the profile header and only the row for the named column.

`internal/csvdata` turns an input stream into a header and data records. Trimming, null recognition, and
rejection of malformed input are separate stories that build on this reader.

Extend `profile.InferType` so date columns are recognised and sort chronologically in later
statistics.

Exact values are rounded once, when written. This story provides the formatting functions the
profile output story applies to each statistic.

`go build ./cmd/csvstat` at the application root produces the `csvstat` executable using only the
standard library. This story delivers the entry point and a minimal `Run` that later stories extend.

A command that fails reports the failure on standard error as one line and leaves standard output empty. Rejected input exits 3. The diagnostic names the file and, for a rejected record, its line number. Standard error is never compared by the grader, so its wording is free; its shape (one line, nothing on standard output) is the contract.

`internal/freq` counts the trimmed non-null values of a named column and returns them in listing
order. The listing is CSV with header `value,count`.

After the frequency listing is ordered, at most `--top` rows are written. `--top` is a positive
integer and defaults to 5 when not given.

When `--freq <name>` is given, `cli.Run` writes the frequency listing for that column and never the profile. `--top <n>` limits the rows and defaults to 5. The listing is CSV with header `value,count`, rows ordered by count descending and then value ascending by byte value, using trimmed non-null values only.

Malformed input must never produce partial output. Rejection is detected in `internal/csvdata`
and mapped to exit code 3 by `internal/cli`.

Numeric columns report a `median`. It is computed exactly and rounded once at output time by the
decimal formatting story.

The profile reports `min` and `max` for every column. The rule depends on the column type. This
story implements the extremes computation in `internal/profile`; two-decimal formatting of
`number` results belongs to the decimal formatting story.

Statistics, type inference and frequency counting all need one shared definition of "null". This
feature provides it as a function in `internal/csvdata` so no other package re-implements it.

Column typing decides which statistics apply. This feature starts `profile.InferType` in
`internal/profile` with the two numeric types. Later features extend the same function with
`date`, `boolean`, `text`, and `empty`.

`csvstat <file>` profiles every column of the file. The output is CSV on standard output with
the header `column,type,rows,nulls,distinct,min,max,sum,mean,median`, followed by one row per
column in file order. Line endings are `\n`.

For each column the profile reports `rows`, `nulls`, and `distinct`. This story implements the
counting function in `internal/profile`. Rendering the counts into the profile CSV belongs to the
profile output story.

Numeric columns report `sum` and `mean`. Exactness matters because the output is rounded once, at
write time, from the exact value. This story computes the exact values; formatting belongs to the
decimal formatting story.

`bin/test.sh` is the project's single test entry point. It is a POSIX `sh` script that changes to the application root, runs `go vet ./...`, then `go test ./...`, and exits non-zero if either fails. The project suite, written beside the code, covers each type, each statistic, null handling, rounding, each output form, and each exit code, using files the tests write to temporary directories.

Every header name and every value is trimmed of leading and trailing spaces before null recognition,
typing, counting, or comparison. Padded text such as `  Globex  ` and `Acme Corp ` therefore count as
`Globex` and `Acme Corp`.

Exit code 1 is the operational-failure code. It is used when the file named on the command line does not exist, is not permitted to be read, or fails while being read. It is distinct from exit 3, which means the file was read but its content was rejected.

Malformed invocations exit 2, write nothing to standard output, and write one diagnostic line to
standard error that shows the two command forms:

```text
csvstat [--column <name>] <file>
csvstat --freq <name> [--top <n>] <file>
```

## Stack

Go

## Verification

```bash
bash bin/test.sh
```

## Next Steps

- Complete the first-run checks above.
- Run `bash bin/test.sh` before making release or deployment changes.
