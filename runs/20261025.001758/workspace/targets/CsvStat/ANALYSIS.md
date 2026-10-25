# Blueprint Analysis: drydock CsvStat Uat Case

## Commander Expectations

- assert the product is a Go command-line tool named `csvstat` that profiles a CSV file
- assert `sh sources/run.sh` exits zero when the completed application is graded
- assert the application uses only the Go standard library and no database
- assert every column of each supplied data file is profiled with the stated statistics
- assert malformed input is rejected with exit code 3 and no standard output
- assert operational failures and usage errors use exit codes 1 and 2 respectively

## Crew

| Crew | Charge |
|---|---|
| Commander | Defines intent and decides what done means. |
| Team Lead | Confirms epic completeness and stakeholder expectations. |
| Planning Crew | Authors atomic specifications and the ordered Manifest. |
| Shipyard Crew | Builds the tickets without synchronous Commander access. |

## Story List

### Feature: Module and Build

| ID | Story | High-level AC |
|---|---|---|
| MOD-001 | Build the csvstat executable from the application root | `go build -o <path> ./cmd/csvstat` succeeds from the application root with only the standard library |
| MOD-002 | Provide the POSIX test entry point | `bin/test.sh` exits zero only when `go vet ./...` and `go test ./...` both pass |

### Feature: CSV Reading and Validation

| ID | Story | High-level AC |
|---|---|---|
| CSV-001 | Read a UTF-8 RFC 4180 CSV file with encoding/csv and trim header and values | Header and values are trimmed of leading and trailing spaces before any other rule |
| CSV-002 | Reject an empty file, empty or repeated header name, or ragged record | Exit 3, nothing on standard output, one stderr line naming the file and the line number for a ragged record |
| CSV-003 | Classify null values | Empty values and `NA`, `N/A`, `NULL` ignoring case are null |
| CSV-004 | Report operational read failures and usage errors | Unreadable file exits 1; malformed arguments exit 2; both write nothing to standard output |

### Feature: Column Type Inference

| ID | Story | High-level AC |
|---|---|---|
| TYP-001 | Infer integer columns | Every non-null value matches `-?[0-9]+` |
| TYP-002 | Infer number columns | Every non-null value matches `-?[0-9]+(\.[0-9]+)?` and the column is not integer |
| TYP-003 | Infer date columns | Every non-null value matches `YYYY-MM-DD` and is a real calendar date |
| TYP-004 | Infer boolean columns | Every non-null value equals `true` or `false` ignoring case |
| TYP-005 | Fall back to text and report empty columns | Remaining columns are text; a column with no non-null values is empty |

### Feature: Column Profile

| ID | Story | High-level AC |
|---|---|---|
| PRF-001 | Emit the profile header and one row per column in file order | Header is `column,type,rows,nulls,distinct,min,max,sum,mean,median` |
| PRF-002 | Compute rows, nulls, and distinct counts | Counts match the data record count, null count, and exact trimmed distinct non-null values |
| PRF-003 | Compute min and max per type | Integer, number, date, boolean, text, and empty rules applied |
| PRF-004 | Compute sum, mean, and median for numeric columns | Integer and number columns only; otherwise empty fields |
| PRF-005 | Format decimals to two places with half-away-from-zero rounding | Number min, max, sum, mean, and median use two decimal places; integer min, max, sum are plain integers |
| PRF-006 | Profile a single named column with `--column` | Output contains only the named column's row |

### Feature: Frequency Listing

| ID | Story | High-level AC |
|---|---|---|
| FRQ-001 | List most frequent non-null values with `--freq` | Header `value,count`; count descending then value ascending by byte value |
| FRQ-002 | Limit frequency rows with `--top` | Default 5; at most `--top` rows; positive integer required |
| FRQ-003 | Write only the header for a column with no non-null values | Output is the `value,count` header only |

### Feature: Command Line Interface

| ID | Story | High-level AC |
|---|---|---|
| CLI-001 | Parse options that precede exactly one file | Missing or extra file exits 2 |
| CLI-002 | Reject unknown options, missing option values, and invalid option combinations | `--column` with `--freq`, `--top` without `--freq`, unknown option, and missing value exit 2 |
| CLI-003 | Reject a named column absent from the header | Exit 3 with nothing on standard output |
| CLI-004 | Write CSV output with RFC 4180 quoting and `\n` line endings | Fields containing commas, quotes, or newlines are quoted |

## Story Guidance

Named stories planning must produce. Authoritative record: `STORY_GUIDANCE.json`.

| Story ID | Provenance | Gate | Note |
|---|---|---|---|
| CSV-002 | plan | `sh sources/run.sh` | sources/run.sh checks E1 (ragged.csv rejected with exit 3) |
| CLI-002 | plan | `sh sources/run.sh` | sources/run.sh checks U1–U6 (usage errors, exit 2) |
| FRQ-001 | plan | `sh sources/run.sh` | sources/run.sh checks C1–C5 (frequency output) |

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | CSV-002 | A failing command writes exactly one diagnostic line to standard error naming the file and, for a ragged record, its line number |
| AC-002 | PRF-005 | Decimal values are rounded half away from zero from the exact value with a leading `-` when negative and no thousands separator |
| AC-003 | TYP-005 | A column whose non-null values are all integer and also all date-shaped is typed by the first matching type in the order integer, number, date, boolean, text |
| AC-004 | FRQ-001 | Null values are never listed in frequency output |
| AC-005 | PRF-003 | Text min and max are compared by byte value |

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
| sources/run.sh | test-kit-to-implementation | ./cmd/csvstat | `sources/run.sh` builds `./cmd/csvstat` and checks 22 outputs | Stage run.sh and all data files; implementation must satisfy every check |
| sources/INSTRUCTIONS.md | instruction-to-test | MOD-002 | `bin/test.sh` requirement | Implementation must provide the POSIX test entry point |
| sources/employees.csv | dependency | PRF-*, FRQ-* | Known answers in `sources/run.sh` A1, B1, B2, C1, C4 | Stage with run.sh |
| sources/sensors.csv | dependency | PRF-*, FRQ-* | Known answers in `sources/run.sh` A2, B3, C2, C5 | Stage with run.sh |
| sources/orders.csv | dependency | PRF-*, FRQ-* | Known answers in `sources/run.sh` A3, C3, E2, E3 | Stage with run.sh |
| sources/header_only.csv | dependency | PRF-* | Known answer in `sources/run.sh` A4 | Stage with run.sh |
| sources/ragged.csv | dependency | CSV-002 | Known answer in `sources/run.sh` E1 | Stage with run.sh |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| sources/INSTRUCTIONS.md | author intent | context | prompt-only |
| sources/run.sh | conformance harness | context | stage |
| sources/employees.csv | test helper | context | stage |
| sources/sensors.csv | test helper | context | stage |
| sources/orders.csv | test helper | context | stage |
| sources/header_only.csv | test helper | context | stage |
| sources/ragged.csv | test helper | context | stage |

## Planning Instructions

### Delivery Shape

The system is a single Go executable `csvstat` built from `./cmd/csvstat` in a Go module rooted at the application root. Inputs are one CSV file path plus options (`--column`, `--freq`, `--top`). Outputs are CSV profile or frequency rows on standard output; diagnostics go to standard error. Execution is a single-pass read of the file, then per-column inference and statistics, then formatted output.

### Story Realization Map

- MOD-001: `sources/INSTRUCTIONS.md` Build section; `sources/run.sh` build step. Capability.
- MOD-002: `sources/INSTRUCTIONS.md` Tests section. Test harness.
- CSV-001 through CSV-004: `sources/INSTRUCTIONS.md` Reading the file, Exit codes; `sources/run.sh` E1–E4, U1–U6. Capability.
- TYP-001 through TYP-005: `sources/INSTRUCTIONS.md` Types; `sources/run.sh` A1–A4. Capability.
- PRF-001 through PRF-006: `sources/INSTRUCTIONS.md` Profile output and Decimals; `sources/run.sh` A1–A4, B1–B3. Capability.
- FRQ-001 through FRQ-003: `sources/INSTRUCTIONS.md` Frequency output; `sources/run.sh` C1–C5. Capability.
- CLI-001 through CLI-004: `sources/INSTRUCTIONS.md` Command and Output format; `sources/run.sh` U1–U6. Capability.

### Test and Acceptance Strategy

Focused story tests run in `bin/test.sh` via `go test ./...`, one feature slice per story (`Suite: scoped`). The terminal verification story depends on every implementation story and gates on the full `sh sources/run.sh` run (`Suite: full`). The 22 checks in `run.sh` are the sole acceptance verdict; no story asserts a passed-count target beyond the full suite.

### Sequencing and Dependencies

- MOD-001 before every other story.
- CSV-001 and CSV-002 before TYP-* and PRF-*.
- TYP-* before PRF-003, PRF-004, PRF-005.
- PRF-* and FRQ-* before CLI-004 output-format verification.
- CLI-* verification and all feature stories before the terminal `sh sources/run.sh` gate.
- Fixtures are staged before verification.

### Source Conflicts and Gaps

- None. The sources define one command, one type order, one exit-code table, and one output format without contradiction. Where the sources are silent, the narrowest behavior consistent with the stated rules is selected: a column whose values are all date-shaped but not all valid dates is typed `text`, since the date rule requires a real calendar date.

## Analysis Notes
generated: 2026-10-24
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.001758/workspace/targets/CsvStat/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 7
  stories: 28
  stack: Go 1.22 standard library
  display_name: not proposed
  short_description: not proposed

Selected behavior notes: the first-match type order is integer, number, date, boolean, text; `--top` defaults to 5; the frequency listing and profile both exclude null values from distinct, min, max, sum, mean, and median computations.
