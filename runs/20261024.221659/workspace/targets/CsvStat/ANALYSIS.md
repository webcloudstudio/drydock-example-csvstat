# Blueprint Analysis: drydock CsvStat Uat Case

## Commander Expectations

- assert the product runs as a standard-library-only Go command-line tool that profiles a CSV file
- assert the product infers a column type for every column and reports its statistics
- assert the product lists the most frequent values of a named column
- assert malformed input is rejected with exit code 3 and no standard output
- assert the supplied grader `sources/run.sh` exits zero against the completed application

## Crew

| Crew | Charge |
|---|---|
| Commander | Defines intent and decides what done means. |
| Team Lead | Confirms epic completeness and stakeholder expectations. |
| Planning Crew | Authors atomic specifications and the ordered Manifest. |
| Shipyard Crew | Builds the tickets without synchronous Commander access. |

## Story List

### Feature: Input Reading and Validation

| ID | Story | High-level AC |
|---|---|---|
| CSVSTAT-001 | Read a UTF-8 RFC 4180 CSV file with trimmed header and values | Header is the first record; every name and value is trimmed of leading and trailing spaces |
| CSVSTAT-002 | Reject an empty file, empty or repeated header name, or ragged record | Exit 3 with no standard output and a diagnostic naming the file and line number |
| CSVSTAT-003 | Classify null values | Empty, `NA`, `N/A`, or `NULL` ignoring case is null |

### Feature: Column Type Inference

| ID | Story | High-level AC |
|---|---|---|
| CSVSTAT-004 | Infer integer and number column types | First matching type across non-null values; integer before number |
| CSVSTAT-005 | Infer date column type with calendar validation | `YYYY-MM-DD` and a real calendar date such as `2025-02-30` is rejected |
| CSVSTAT-006 | Infer boolean and text column types | Boolean is `true` or `false` ignoring case; text is the fallback |
| CSVSTAT-007 | Report `empty` type for columns with no non-null values | Type is `empty` and statistic fields are empty |

### Feature: Column Profile Statistics

| ID | Story | High-level AC |
|---|---|---|
| CSVSTAT-008 | Report rows, nulls, and distinct counts per column | Counts follow the Read and Null rules; distinct compares trimmed exact strings |
| CSVSTAT-009 | Report min and max for integer, number, date, boolean, and text | Each type uses its ordering rule as specified in Types |
| CSVSTAT-010 | Report sum and mean for integer and number columns | Integer total is plain; number totals and means use two decimal places |
| CSVSTAT-011 | Report median for integer and number columns | Middle value, or mean of two middle values when count is even |
| CSVSTAT-012 | Format decimal output with two places, half away from zero | `1012.50`, `-12.21`, `0.00`; no thousands separator |

### Feature: Profile Command

| ID | Story | High-level AC |
|---|---|---|
| CSVSTAT-013 | Emit the profile CSV for all columns of a file | Header `column,type,rows,nulls,distinct,min,max,sum,mean,median`; one row per column in file order |
| CSVSTAT-014 | Profile a single named column with `--column` | Only the named column is emitted; unknown name exits 3 |

### Feature: Frequency Command

| ID | Story | High-level AC |
|---|---|---|
| CSVSTAT-015 | List the most frequent values with `--freq` | Header `value,count`; count descending then value ascending by byte value; null values excluded |
| CSVSTAT-016 | Limit frequency output with `--top` | Positive integer; default 5; valid only with `--freq` |
| CSVSTAT-017 | Write only the header for a column with no non-null values | Output is the header line only |

### Feature: Command Line Interface

| ID | Story | High-level AC |
|---|---|---|
| CSVSTAT-018 | Parse options that precede exactly one file argument | Options before file; exactly one file required |
| CSVSTAT-019 | Map usage errors to exit code 2 with no standard output | Missing or extra file, unknown option, missing option value, invalid `--top`, `--column` with `--freq`, `--top` without `--freq` |
| CSVSTAT-020 | Map operational read failures to exit code 1 with no standard output | Unopenable or unreadable file exits 1 |
| CSVSTAT-021 | Emit a diagnostic on standard error for each failing command | One line naming the file and, for a rejected record, its line number |

### Feature: Build and Test Harness

| ID | Story | High-level AC |
|---|---|---|
| CSVSTAT-022 | Provide a Go module whose main package builds at `./cmd/csvstat` | `go build ./cmd/csvstat` succeeds using only the standard library |
| CSVSTAT-023 | Provide `bin/test.sh` that runs `go vet` and `go test` | Exits zero only when both pass |
| CSVSTAT-024 | Provide an automated test suite covering types, statistics, nulls, rounding, output forms, and exit codes | Tests write their own files; suite passes under `go test ./...` |
| CSVSTAT-025 | Run the supplied grader as the terminal verification | Full `sh sources/run.sh` run exits zero against the completed application |

Story count: 25 stories across 7 features.

## Story Guidance

Named stories planning must produce. Authoritative record: `STORY_GUIDANCE.json`.

None.

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | CSVSTAT-002 | When a record has a different field count from the header, the application shall name the file and the record's line number on standard error. |
| AC-002 | CSVSTAT-004 | The integer type shall be assigned before number when a column satisfies both. |
| AC-003 | CSVSTAT-008 | The distinct count shall compare trimmed values as exact strings. |
| AC-004 | CSVSTAT-010 | The sum shall be written as a plain integer for integer columns and with two decimal places for number columns. |
| AC-005 | CSVSTAT-012 | Rounding shall be half away from zero applied to the exact value. |
| AC-006 | CSVSTAT-015 | Null values shall never appear in frequency output. |
| AC-007 | CSVSTAT-017 | A column with no non-null values shall produce a frequency output containing only the header line. |
| AC-008 | CSVSTAT-019 | A failing command shall write nothing to standard output. |

## Source Inventory

| Path | Content kind | Disposition | Reason |
|---|---|---|---|
| `sources/INSTRUCTIONS.md` | markdown | analyzed | readable UTF-8 |
| `sources/employees.csv` | text | analyzed | readable UTF-8 |
| `sources/header_only.csv` | text | analyzed | readable UTF-8 |
| `sources/orders.csv` | text | analyzed | readable UTF-8 |
| `sources/ragged.csv` | text | analyzed | readable UTF-8 |
| `sources/run.sh` | code | analyzed | readable UTF-8 |
| `sources/sensors.csv` | text | analyzed | readable UTF-8 |

## Relationship Model

| Source or group | Relationship type | Related source or group | Evidence | Delivery implication |
|---|---|---|---|---|
| `sources/INSTRUCTIONS.md` | instruction-to-test | `sources/run.sh` | Command, exit code, and known-answer sections in INSTRUCTIONS.md match checks A1–E4 and U1–U6 | Profile and usage behavior must match the grader byte for byte |
| `sources/run.sh` | test-kit-to-implementation | `cmd/csvstat` | Builds `./cmd/csvstat` and runs 22 checks against it | Build target path is fixed; grader is staged and run from application root |
| `sources/employees.csv` | dependency | `sources/run.sh` | Known answers A1, B1, B2, C1, C4 | Staged data files must be present at build-relative `sources/` |
| `sources/sensors.csv` | dependency | `sources/run.sh` | Known answers A2, B3, C2, C5 | Staged data file must be present |
| `sources/orders.csv` | dependency | `sources/run.sh` | Known answers A3, C3, E2, E3, U2, U5, U6 | Staged data file must be present |
| `sources/header_only.csv` | dependency | `sources/run.sh` | Known answer A4 | Staged data file must be present |
| `sources/ragged.csv` | dependency | `sources/run.sh` | Known answer E1 | Staged data file must be present |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| sources/INSTRUCTIONS.md | normative specification | context | prompt-only |
| sources/run.sh | conformance harness | compass | stage |
| sources/employees.csv | conformance harness | context | stage |
| sources/sensors.csv | conformance harness | context | stage |
| sources/orders.csv | conformance harness | context | stage |
| sources/header_only.csv | conformance harness | context | stage |
| sources/ragged.csv | conformance harness | context | stage |

## Planning Instructions

### Delivery Shape

The system is a single-binary Go command-line profiler. Input is one CSV file path plus options. Output is CSV on standard output for profile or frequency forms, or a single diagnostic line on standard error with exit codes 1, 2, or 3 on failure. The required flow is: parse arguments, open and read the file with `encoding/csv`, validate header and records, classify nulls and types, compute statistics per column, then emit output.

### Story Realization Map

- CSVSTAT-001 through CSVSTAT-003: `sources/INSTRUCTIONS.md` "Reading the file" section; exercised by `sources/run.sh` checks A1–A4 and E1.
- CSVSTAT-004 through CSVSTAT-007: `sources/INSTRUCTIONS.md` "Types" section; exercised by A1–A4.
- CSVSTAT-008 through CSVSTAT-012: `sources/INSTRUCTIONS.md` "Profile output" and "Decimals" sections; exercised by A1–A4 and B1–B3.
- CSVSTAT-013 and CSVSTAT-014: `sources/INSTRUCTIONS.md` "Command" and "Profile output"; exercised by A1–A4, B1–B3.
- CSVSTAT-015 through CSVSTAT-017: `sources/INSTRUCTIONS.md` "Frequency output"; exercised by C1–C5.
- CSVSTAT-018 through CSVSTAT-021: `sources/INSTRUCTIONS.md` "Command", "Output format", and "Exit codes"; exercised by E1–E4 and U1–U6.
- CSVSTAT-022 through CSVSTAT-024: `sources/INSTRUCTIONS.md` "Stack" and "Tests" sections.
- CSVSTAT-025: terminal verification gating on `sources/run.sh`.

### Test and Acceptance Strategy

Focused story tests run under `go test` on constructed files written by the tests. The terminal verification story (CSVSTAT-025) runs `sh sources/run.sh` unfiltered and gates on exit zero. The grader is not a substitute for per-story tests.

### Sequencing and Dependencies

Module and build (CSVSTAT-022) precedes all other stories. Input reading (CSVSTAT-001 to CSVSTAT-003) precedes type inference and statistics. Type inference precedes statistics. Profile and frequency commands depend on statistics and reading. CLI parsing precedes the commands. Test suite (CSVSTAT-024) and `bin/test.sh` (CSVSTAT-023) follow the implementation they cover. CSVSTAT-025 depends on every implementation story.

### Source Conflicts and Gaps

None. The sources are internally consistent. The rounding rule in the Decimals section and the known answers in `run.sh` agree on the two-decimal form. The mean of `mean` for `qty` in A3 (6.87) and median (6.00) are taken as stated by the grader.

## Analysis Notes
generated: 2026-10-24
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261024.221659/workspace/targets/CsvStat/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 7
  stories: 25
  stack: Go 1.22+ standard library
  display_name: not proposed
  short_description: not proposed

Observations: The `ragged.csv` line-number diagnostic is a stderr-only requirement and is not compared by the grader. `sources/run.sh` is the only file that must be staged as a harness; its data files are staged alongside it.
