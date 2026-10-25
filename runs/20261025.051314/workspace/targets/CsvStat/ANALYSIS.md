# Blueprint Analysis: drydock CsvStat Uat Case

## Commander Expectations

- assert csvstat profiles every column of a CSV file, inferring each column's type and reporting rows, nulls, distinct, min, max, sum, mean, and median
- assert csvstat lists the most frequent non-null values of one named column
- assert csvstat rejects empty, malformed, or ragged input with exit 3 and writes nothing to standard output
- assert the supplied grader `sources/run.sh` passes every one of its checks

## Crew

| Crew | Charge |
|---|---|
| Commander | Defines intent and decides what done means. |
| Team Lead | Confirms epic completeness and stakeholder expectations. |
| Planning Crew | Authors atomic specifications and the ordered Manifest. |
| Shipyard Crew | Builds the tickets without synchronous Commander access. |

## Story List

### Feature: Build and Test Gate

| ID | Story | High-level AC |
|---|---|---|
| FOUND-001 | Build the csvstat executable from the module root | `go build ./cmd/csvstat` succeeds at the application root using only the Go standard library |
| FOUND-002 | Provide the POSIX test gate | `bin/test.sh` exits zero only when `go vet ./...` and `go test ./...` both pass |

### Feature: Command Line Interface

| ID | Story | High-level AC |
|---|---|---|
| CLI-001 | Accept exactly one file after optional options | One file is required; options precede the file |
| CLI-002 | Profile a single named column | `--column <name>` writes only that column's profile row |
| CLI-003 | Select frequency mode with `--freq` | `--freq <name>` lists frequencies instead of the profile |
| CLI-004 | Reject malformed arguments with exit 2 | Missing or extra file, unknown option, missing option value, non-positive `--top`, `--top` without `--freq`, and `--column` with `--freq` each exit 2 |

### Feature: CSV Ingestion

| ID | Story | High-level AC |
|---|---|---|
| INGEST-001 | Read UTF-8 RFC 4180 CSV with a header record | The first record is the header; quoted fields containing commas parse as one value |
| INGEST-002 | Trim header names and values | Leading and trailing spaces are removed from every header name and value before other rules apply |
| INGEST-003 | Recognize null spellings | Empty values and `NA`, `N/A`, `NULL` ignoring case are null |
| INGEST-004 | Reject empty, bad-header, or ragged input | An empty file, an empty or repeated header name, or a record with a different field count exits 3 |

### Feature: Column Typing

| ID | Story | High-level AC |
|---|---|---|
| TYPE-001 | Infer integer and number columns | Every non-null value matches `-?[0-9]+` for integer, or `-?[0-9]+(\.[0-9]+)?` for number |
| TYPE-002 | Infer date columns | Every non-null value matches `YYYY-MM-DD` and is a real calendar date |
| TYPE-003 | Infer boolean and text columns, and empty for all-null columns | Boolean is `true` or `false` ignoring case; text is the fallback; a column with no non-null value is `empty` |

### Feature: Column Statistics

| ID | Story | High-level AC |
|---|---|---|
| STAT-001 | Report rows, nulls, and distinct counts | Rows counts data records; nulls counts null values; distinct counts exact trimmed non-null strings |
| STAT-002 | Report min and max by column type | Integer, number, date, boolean, and text min and max follow the brief's rules; empty columns leave them blank |
| STAT-003 | Report sum and mean for numeric columns | Integer sum is a plain integer; number sum is a decimal; mean is total over non-null count; other types leave them blank |
| STAT-004 | Report median for numeric columns | Median of sorted non-null values, or mean of the two middle values when the count is even |
| STAT-005 | Format decimal results to two places | Number min, max, sum, mean, and median, and integer mean and median, use two decimals rounded half away from zero |

### Feature: Profile Output

| ID | Story | High-level AC |
|---|---|---|
| OUT-001 | Write the profile as RFC 4180 CSV | Header `column,type,rows,nulls,distinct,min,max,sum,mean,median`, one row per column in file order, `\n` line endings, fields quoted as needed |

### Feature: Frequency Listing

| ID | Story | High-level AC |
|---|---|---|
| FREQ-001 | Count and order distinct non-null values | Header `value,count`; count descending, then value ascending by byte; nulls omitted; header only when no non-null value exists |
| FREQ-002 | Limit frequency rows with `--top` | At most `--top` rows are listed; default is 5 |

### Feature: Failure Reporting

| ID | Story | High-level AC |
|---|---|---|
| ERR-001 | Write a diagnostic line on failure | A failing command writes nothing to standard output and one stderr line naming the file, and the line number for a rejected record |
| ERR-002 | Exit 1 when the file cannot be read | A missing or unreadable file exits 1 and writes nothing to standard output |

### Feature: Acceptance Verification

| ID | Story | High-level AC |
|---|---|---|
| VERIFY-001 | Pass the supplied grader | `sh sources/run.sh` exits zero with every one of its 22 checks passing |

## Story Guidance

Named stories planning must produce. Authoritative record: `STORY_GUIDANCE.json`.

None.

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | INGEST-001 | A line ending after the last record does not create an extra data record. |
| AC-002 | INGEST-003 | A quoted empty field is an empty value and is counted as null. |
| AC-003 | INGEST-004 | A file containing only blank lines is rejected as empty with exit 3. |
| AC-004 | STAT-005 | A number total or mean that rounds to zero is written as `0.00` without a leading minus sign. |
| AC-005 | TYPE-003 | Boolean min and max are written as lowercase `true` or `false` regardless of input case. |
| AC-006 | CLI-004 | A usage error writes one stderr line that shows the two command forms. |
| AC-007 | STAT-002 | A column with exactly one non-null value reports that value as both min and max. |

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
| `sources/run.sh` | instruction-to-test | `sources/INSTRUCTIONS.md` | Check IDs and expected output mirror the Command and Profile output sections | VERIFY-001 depends on every implementation story |
| `sources/run.sh` | test-kit-to-implementation | `./cmd/csvstat` | `go build -o "$BIN" ./cmd/csvstat` at the application root | FOUND-001 precedes every check |
| `sources/employees.csv`, `sources/sensors.csv`, `sources/orders.csv`, `sources/header_only.csv`, `sources/ragged.csv` | dependency | `sources/run.sh` | `DATA` resolves to `sources/` when `employees.csv` is present beside the script | Stage all five fixtures under `sources/` |
| `sources/INSTRUCTIONS.md` | instruction-to-test | `bin/test.sh` suite | "Tests" section requires coverage of types, statistics, nulls, rounding, outputs, and exit codes | FOUND-002 owns the suite; stories add their own slices |
| `sources/INSTRUCTIONS.md` | dependency | `sources/ragged.csv`, `sources/header_only.csv` | Rejection and empty-profile rules | INGEST-004 and STAT-001 depend on these fixtures |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| sources/INSTRUCTIONS.md | normative specification and conformance test suite | context | prompt-only |
| sources/run.sh | conformance harness | context | stage |
| sources/employees.csv | asset | context | stage |
| sources/sensors.csv | asset | context | stage |
| sources/orders.csv | asset | context | stage |
| sources/header_only.csv | asset | context | stage |
| sources/ragged.csv | asset | context | stage |

## Planning Instructions

### Delivery Shape

The system is a single Go command-line executable, `csvstat`, in a Go module whose main package is `./cmd/csvstat`. It reads one CSV file given on the command line and writes either a column profile or a frequency listing to standard output as CSV. Inputs are one file and its options. Outputs are CSV on standard output and a diagnostic line on standard error for failures. Exit codes are 0 for success, 1 for an unreadable file, 2 for usage errors, and 3 for rejected input. The execution flow is: parse arguments, read and trim the file, reject malformed input, infer column types, then compute either the profile or the frequencies.

### Story Realization Map

- FOUND-001: `sources/INSTRUCTIONS.md` (Stack, Run Harness); related `sources/run.sh` build step. Requires a capability (module and entry point).
- FOUND-002: `sources/INSTRUCTIONS.md` (Tests). Requires a test harness (`bin/test.sh`).
- CLI-001 to CLI-004: `sources/INSTRUCTIONS.md` (Command, Exit codes); `sources/run.sh` checks U1 to U6. Requires a capability.
- INGEST-001 to INGEST-004: `sources/INSTRUCTIONS.md` (Reading the file); `sources/ragged.csv`, `sources/header_only.csv`, `sources/run.sh` E1, E4. Requires a capability.
- TYPE-001 to TYPE-003: `sources/INSTRUCTIONS.md` (Types); `sources/sensors.csv`, `sources/orders.csv`, `sources/employees.csv`. Requires a capability.
- STAT-001 to STAT-005: `sources/INSTRUCTIONS.md` (Profile output, Decimals); `sources/run.sh` A1 to A4, B1 to B3. Requires a capability.
- OUT-001: `sources/INSTRUCTIONS.md` (Output format); `sources/run.sh` byte comparisons. Requires a capability.
- FREQ-001 and FREQ-002: `sources/INSTRUCTIONS.md` (Frequency output); `sources/run.sh` C1 to C5. Requires a capability.
- ERR-001 and ERR-002: `sources/INSTRUCTIONS.md` (Output format, Exit codes); `sources/run.sh` E1 to E4. Requires a capability.
- VERIFY-001: `sources/run.sh` as the complete grader. Requires an acceptance contract with `Suite: full`.

### Test and Acceptance Strategy

Each implementation story runs its own focused tests and never runs the whole grader. A story that exercises the grader declares `Suite: scoped` and bounds the run with the runner's `--pattern` selector. `VERIFY-001` is the single terminal story. It depends on every other story and declares `Suite: full`, which runs `sh sources/run.sh` unfiltered and requires exit zero. The project-wide criteria in `SEA_TRIALS.md` are proved by the terminal run and by `bin/test.sh`. Sample tests prove a unit only. Full conformance is proved by the grader.

### Sequencing and Dependencies

- FOUND-001 precedes all other stories; it is build-before-test.
- INGEST-001 precedes INGEST-002, INGEST-003, and INGEST-004.
- INGEST-003 and TYPE-001 precede TYPE-002, TYPE-003, and every statistic story.
- TYPE-003 precedes STAT-002 through STAT-005.
- STAT-005 precedes OUT-001, because the profile rows carry formatted decimals.
- FREQ-001 precedes FREQ-002.
- CLI-001 precedes CLI-002, CLI-003, and CLI-004.
- ERR-001 and ERR-002 follow INGEST-004 and CLI-004.
- FOUND-002 follows the stories whose tests it runs.
- VERIFY-001 follows all other stories.
- External dependencies: Go 1.22 or newer. No network access and no database are used.

### Source Conflicts and Gaps

No conflicts were found between `sources/INSTRUCTIONS.md`, the fixtures, and `sources/run.sh`. Spot checks by hand, not by execution, matched the expected values: the employee `id` sum of 24300, the `manager_id` statistics (9 non-null values, sum 30, mean 3.33, median 4.00), the `orders.csv` `ref` and `shipped` text typing, and the `customer` distinct count with Globex at 10. No blocker or questionnaire is required.

## Analysis Notes
generated: 2026-10-25
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/workspace/targets/CsvStat/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 9
  stories: 24
  stack: Go 1.22 or newer, standard library only
  display_name: not proposed
  short_description: not proposed

- `COMPASS.md` exists, so it is not emitted. `BLOCKERS.md` and every questionnaire are omitted because no decision is open.
- The prior `SEA_TRIALS.md` IDs st-001 through st-008 are preserved. st-002 gains a `Command`. st-009 is new.
- st-009 is checked by a source grep and cannot prove the absence of an embedded answer. It is partial coverage.
- The earlier concern about `orders.csv` is resolved: `ref` and `shipped` are text because they contain non-numeric and invalid-date values, which matches the expected output.
- Surfaced Acceptance Criteria AC-001 to AC-007 cover gaps the brief does not state explicitly: trailing newlines, quoted empty fields, blank-only files, zero-rounding signs, boolean case, usage help text, and single-value min and max.
