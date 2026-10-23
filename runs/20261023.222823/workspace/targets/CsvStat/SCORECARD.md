# Scorecard: CsvStat

## Specification

> **NOT YET RUN**
> drydock score spec CsvStat

## Acceptance

> **PASSED**
> drydock score ac CsvStat · 2026-10-23T22:38:41.114+00:00

PASS 27   FAIL 0   UNVERIFIED 0   PREPASSED 8   (35 AC total)

## Build

> **RECORDED**
> drydock score build CsvStat

Blocks: 3   Repaired: 1   Failed: 0
Tokens: 990,498 (cached 876,160, output 18,927)   Cache hit rate: 90.2%

## Release

> **RELEASED**
> drydock score release CsvStat · 2026-10-23T22:39:19.755+00:00

- Verdict: PASSED

CsvStat: PASSED — 8 of 8

  st-001  MET       ["sh", "sources/run.sh"] exited 0; sources/run.sh: governed gate exited 0; facts: CsvStat: 22 passed, 0 failed (of 22)
  st-002  MET       bin/test.sh:1-5; probe: sh bin/test.sh exited 0
  st-003  MET       go.mod:1-3; probe: go build ./cmd/csvstat exited 0
  st-004  MET       internal/csvstat/csvstat.go:100-136; probe: valid inference produced all required type classifications
  st-005  MET       sources/run.sh: A1-A4 passed; facts: 22 passed, 0 failed
  st-006  MET       sources/run.sh: C1-C5 passed; internal/csvstat/csvstat.go:250-270
  st-007  MET       internal/csvstat/csvstat.go:21-43; probe: all five malformed cases returned rc=3 and stdout_bytes=0; sources/run.sh: E1-E3 passed
  st-008  MET       internal/cli/cli.go:43-73; probe: unreadable returned rc=1 and usage cases returned rc=2 with stdout_bytes=0; sources/run.sh: E4 and U1-U6 passed
  reported: Build directory has uncommitted changes

## Project acceptance

| ID | Type | Criterion | Verdict | Observed |
|---|---|---|---|---|
| st-001 | technical | The completed csvstat shall make sh sources/run.sh exit zero; that script's exit status is the sole acceptance verdict. | MET | ["sh", "sources/run.sh"] exited 0; sources/run.sh: governed gate exited 0; facts: CsvStat: 22 passed, 0 failed (of 22) |
| st-002 | technical | The application shall provide a POSIX-compatible bin/test.sh that exits zero when go vet and every automated test pass. | MET | bin/test.sh:1-5; probe: sh bin/test.sh exited 0 |
| st-003 | technical | The application root shall be a Go module from which go build ./cmd/csvstat produces the csvstat executable using only the standard library. | MET | go.mod:1-3; probe: go build ./cmd/csvstat exited 0 |
| st-004 | behavioral | The application shall assign each column the first of integer, number, date, boolean, and text that every non-null value satisfies, and empty when it has no non-null value. | MET | internal/csvstat/csvstat.go:100-136; probe: valid inference produced all required type classifications |
| st-005 | behavioral | The application shall compute rows, nulls, distinct, min, max, sum, mean, and median for every column of the supplied data files, matching the known answers in sources/run.sh byte for byte. | MET | sources/run.sh: A1-A4 passed; facts: 22 passed, 0 failed |
| st-006 | behavioral | When --freq is given, the application shall list the column's non-null values by count descending then value ascending, at most --top rows. | MET | sources/run.sh: C1-C5 passed; internal/csvstat/csvstat.go:250-270 |
| st-007 | behavioral | If the file is empty, has an empty or repeated header name, has a ragged record, or a named column is absent, then the application shall exit 3 and write nothing to standard output. | MET | internal/csvstat/csvstat.go:21-43; probe: all five malformed cases returned rc=3 and stdout_bytes=0; sources/run.sh: E1-E3 passed |
| st-008 | behavioral | If the file cannot be read, then the application shall exit 1; if the arguments are malformed, then the application shall exit 2; in both cases it shall write nothing to standard output. | MET | internal/cli/cli.go:43-73; probe: unreadable returned rc=1 and usage cases returned rc=2 with stdout_bytes=0; sources/run.sh: E4 and U1-U6 passed |

**Advisory warnings:**
- Build directory has uncommitted changes
