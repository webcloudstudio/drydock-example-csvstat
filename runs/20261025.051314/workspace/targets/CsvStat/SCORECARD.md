# Scorecard: CsvStat

## Specification

> **NOT YET RUN**
> drydock score spec CsvStat

## Acceptance

> **PASSED**
> drydock score ac CsvStat · 2026-10-25T06:08:07.405+00:00

PASS 94   FAIL 0   UNVERIFIED 0   PREPASSED 32   (126 AC total)

## Build

> **RECORDED**
> drydock score build CsvStat

Blocks: 6   Repaired: 0   Failed: 0
Tokens: 2,594,811 (cached 2,485,312, output 109,399)   Cache hit rate: 100.0%

## Release

> **RELEASED**
> drydock score release CsvStat · 2026-10-25T06:09:36.478+00:00

- Verdict: PASSED

CsvStat: PASSED — 8 of 8

  st-001  MET       ["sh", "sources/run.sh"] exited 0; governed_gate full: outcome PASS, return_code 0, stdout 'CsvStat: 22 passed, 0 failed (of 22)'; sources/run.sh:44-50 builds ./cmd/csvstat and exits 1 on build failure; it exits 0 only when every check passes
  st-002  MET       bin/test.sh:1-6 is a POSIX sh script (set -eu, go vet, go test) and `sh -n` parses it; go vet ./... exited 0; go test ./...: ok csvstat/internal/cli, csvdata, freq, profile; cmd/csvstat has no test files; script exit 0
  st-003  MET       go.mod:1-3 'module csvstat', 'go 1.22', no require directives; go build ./cmd/csvstat exited 0; grep of all Go imports: only stdlib packages and csvstat/internal/*
  st-004  MET       internal/profile/infer.go:31-47 ordered switch: empty, integer, number, date, boolean, text; nulls filtered via csvdata.IsNull; probe: [1,2,-3] -> integer; [1,2.5] -> number; [1.0,2] -> number; [1,NA,'',2] -> integer; [2025-01-01,NA,2024-02-29] -> date; [2025-01-01,2025-02-30] -> text (invalid calendar date); [true,FALSE,True] -> boolean; [true,1] -> text; [1e3,2] -> text; [' 7 ',8] -> integer; [NA,'',N/A] -> empty; grader A1-A4 and B1-B3 also match the expected type column for every column of the supplied files
  st-005  MET       governed_gate: PASS A1 (employees.csv), A2 (sensors.csv), A3 (orders.csv), A4 (header_only.csv), B1-B3 (single-column runs); sources/run.sh:58-103 expected outputs are cmp-compared against actual stdout, exit codes included
  st-006  MET       internal/freq/count.go:27-35 sorts by count desc then value asc; null values skipped before counting; internal/freq/top.go:7-16 keeps the first n entries; grader PASS C1-C5 (including --top 2 and --top 3); probe: values b,a,b,a,c,c,c,NA x5,B,'' -> c 3, a 2, b 2, B 1 (NA and empty excluded, 'B' before 'a' by byte order); --top 2 -> c,a; --top 99 -> all rows
  st-007  MET       internal/csvdata/validate.go:23-50 rejects nil header, empty name, repeated name, and field count mismatch; internal/cli/cli.go:38-42 returns exitRejected on Validate failure before any output is written; grader PASS E1 (ragged, exit 3), E2 and E3 (absent column, exit 3), with empty expected stdout; probe: empty file, '\n\n\n', 'a,,c', 'a,a', 'a, a', 'a,b / 1,2,3', plus absent column and ragged under --freq and --column: all rc=3, stdout empty
  st-008  MET       internal/cli/cli.go:31-33 usage errors return exitUsage (2); cli.go:44-47 open failures return exitOperational (1); grader PASS E4 (missing file, exit 1) and U1-U6 (usage errors, exit 2), each with empty stdout; probe: nonexistent path -> rc 1, empty stdout; directory path -> rc 1, empty stdout; no args, --column without value, --top abc, --top without --freq, --top 0, --column with --freq, repeated --column, unknown flag, two files -> rc 2, empty stdout

## Project acceptance

| ID | Type | Criterion | Verdict | Observed |
|---|---|---|---|---|
| st-001 | technical | The completed csvstat shall make sh sources/run.sh exit zero; that script's exit status is the sole acceptance verdict. | MET | ["sh", "sources/run.sh"] exited 0; governed_gate full: outcome PASS, return_code 0, stdout 'CsvStat: 22 passed, 0 failed (of 22)'; sources/run.sh:44-50 builds ./cmd/csvstat and exits 1 on build failure; it exits 0 only when every check passes |
| st-002 | technical | The application shall provide a POSIX-compatible bin/test.sh that exits zero when go vet and every automated test pass. | MET | bin/test.sh:1-6 is a POSIX sh script (set -eu, go vet, go test) and `sh -n` parses it; go vet ./... exited 0; go test ./...: ok csvstat/internal/cli, csvdata, freq, profile; cmd/csvstat has no test files; script exit 0 |
| st-003 | technical | The application root shall be a Go module from which go build ./cmd/csvstat produces the csvstat executable using only the standard library. | MET | go.mod:1-3 'module csvstat', 'go 1.22', no require directives; go build ./cmd/csvstat exited 0; grep of all Go imports: only stdlib packages and csvstat/internal/* |
| st-004 | behavioral | The application shall assign each column the first of integer, number, date, boolean, and text that every non-null value satisfies, and empty when it has no non-null value. | MET | internal/profile/infer.go:31-47 ordered switch: empty, integer, number, date, boolean, text; nulls filtered via csvdata.IsNull; probe: [1,2,-3] -> integer; [1,2.5] -> number; [1.0,2] -> number; [1,NA,'',2] -> integer; [2025-01-01,NA,2024-02-29] -> date; [2025-01-01,2025-02-30] -> text (invalid calendar date); [true,FALSE,True] -> boolean; [true,1] -> text; [1e3,2] -> text; [' 7 ',8] -> integer; [NA,'',N/A] -> empty; grader A1-A4 and B1-B3 also match the expected type column for every column of the supplied files |
| st-005 | behavioral | The application shall compute rows, nulls, distinct, min, max, sum, mean, and median for every column of the supplied data files, matching the known answers in sources/run.sh byte for byte. | MET | governed_gate: PASS A1 (employees.csv), A2 (sensors.csv), A3 (orders.csv), A4 (header_only.csv), B1-B3 (single-column runs); sources/run.sh:58-103 expected outputs are cmp-compared against actual stdout, exit codes included |
| st-006 | behavioral | When --freq is given, the application shall list the column's non-null values by count descending then value ascending, at most --top rows. | MET | internal/freq/count.go:27-35 sorts by count desc then value asc; null values skipped before counting; internal/freq/top.go:7-16 keeps the first n entries; grader PASS C1-C5 (including --top 2 and --top 3); probe: values b,a,b,a,c,c,c,NA x5,B,'' -> c 3, a 2, b 2, B 1 (NA and empty excluded, 'B' before 'a' by byte order); --top 2 -> c,a; --top 99 -> all rows |
| st-007 | behavioral | If the file is empty, has an empty or repeated header name, has a ragged record, or a named column is absent, then the application shall exit 3 and write nothing to standard output. | MET | internal/csvdata/validate.go:23-50 rejects nil header, empty name, repeated name, and field count mismatch; internal/cli/cli.go:38-42 returns exitRejected on Validate failure before any output is written; grader PASS E1 (ragged, exit 3), E2 and E3 (absent column, exit 3), with empty expected stdout; probe: empty file, '\n\n\n', 'a,,c', 'a,a', 'a, a', 'a,b / 1,2,3', plus absent column and ragged under --freq and --column: all rc=3, stdout empty |
| st-008 | behavioral | If the file cannot be read, then the application shall exit 1; if the arguments are malformed, then the application shall exit 2; in both cases it shall write nothing to standard output. | MET | internal/cli/cli.go:31-33 usage errors return exitUsage (2); cli.go:44-47 open failures return exitOperational (1); grader PASS E4 (missing file, exit 1) and U1-U6 (usage errors, exit 2), each with empty stdout; probe: nonexistent path -> rc 1, empty stdout; directory path -> rc 1, empty stdout; no args, --column without value, --top abc, --top without --freq, --top 0, --column with --freq, repeated --column, unknown flag, two files -> rc 2, empty stdout |

**Ranked improvements:**
1. Add grader checks for the st-007 cases the grader does not cover today: empty file, blank-only input, empty header name, and repeated header name. Those rules currently rest only on probes, and a future change could break them without the gate noticing.
2. Add grader checks for the st-004 inference edge cases: an impossible calendar date such as 2025-02-30 must be text, a mix of integer and date must be text, and a column of only nulls must be empty.
3. Add a grader check for an unreadable directory path (exit 1, empty stdout) so the st-008 read-failure branch is covered by the gate.
