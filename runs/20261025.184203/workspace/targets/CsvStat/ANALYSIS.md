# Blueprint Analysis: drydock CsvStat Uat Case

## Commander Expectations

- assert csvstat profiles every column of a CSV file and lists column value frequencies as the brief specifies
- assert the supplied grader `sources/run.sh` exits zero against the completed application
- assert csvstat is a Go module that builds with the standard library only
- assert csvstat rejects malformed input with exit 3 and writes nothing to standard output

## Crew

| Crew | Charge |
|---|---|
| Commander | Defines intent and decides what done means. |
| Team Lead | Confirms epic completeness and stakeholder expectations. |
| Planning Crew | Authors atomic specifications and the ordered Manifest. |
| Shipyard Crew | Builds the tickets without synchronous Commander access. |

## Story List

### Feature: Command-Line Interface

| ID | Story | High-level AC |
|---|---|---|
| CLI-001 | Reject invocations with the wrong arguments | Missing or extra file, unknown option, or missing option value exits 2 with nothing on stdout. |
| CLI-002 | Reject invalid option combinations | `--top` must be a positive integer and valid only with `--freq`; `--column` cannot combine with `--freq`; each case exits 2. |

### Feature: File Input

| ID | Story | High-level AC |
|---|---|---|
| INPUT-001 | Read the file as UTF-8 RFC 4180 CSV | The first record is the header and the rest are data records, read with `encoding/csv`. |
| INPUT-002 | Report an unreadable file | A missing or unreadable file exits 1 with nothing on stdout and one stderr line naming the file. |
| INPUT-003 | Trim header names and values | Leading and trailing spaces are removed from every header name and value before any other rule. |
| INPUT-004 | Reject an empty file | An empty file exits 3 with nothing on stdout. |
| INPUT-005 | Reject empty or repeated header names | A header with an empty or repeated name exits 3 with nothing on stdout. |
| INPUT-006 | Reject ragged records | A record whose field count differs from the header exits 3 and names its line number on stderr. |

### Feature: Null Handling

| ID | Story | High-level AC |
|---|---|---|
| NULL-001 | Treat empty and NA-style values as null | Empty values and `NA`, `N/A`, `NULL` ignoring case are null; all other values are non-null. |

### Feature: Column Types

| ID | Story | High-level AC |
|---|---|---|
| TYPE-001 | Infer integer and number types | A column of values matching `-?[0-9]+` is `integer`; one matching `-?[0-9]+(\.[0-9]+)?` is `number`; the first satisfied type wins. |
| TYPE-002 | Infer the date type | A column of `YYYY-MM-DD` values that are real calendar dates is `date`; `2025-02-30` is not a date. |
| TYPE-003 | Infer the boolean type | A column whose non-null values all equal `true` or `false` ignoring case is `boolean`. |
| TYPE-004 | Fall back to text and empty | Any other non-null column is `text`; a column with no non-null values is `empty`. |

### Feature: Column Statistics

| ID | Story | High-level AC |
|---|---|---|
| STAT-001 | Count rows and nulls | `rows` counts data records and `nulls` counts null values for each column. |
| STAT-002 | Count distinct non-null values | `distinct` counts different non-null values compared as exact trimmed strings. |
| STAT-003 | Report min and max for numeric columns | `integer` and `number` columns report their smallest and largest values in their own numeric form. |
| STAT-004 | Report min and max for date columns | `date` columns report the earliest and latest values. |
| STAT-005 | Report min and max for boolean columns | `min` is `false` if any value is false, else `true`; `max` is `true` if any value is true, else `false`. |
| STAT-006 | Report min and max for text columns | `text` columns report the smallest and largest values by byte value. |
| STAT-007 | Report sum for numeric columns | `integer` reports the integer total and `number` the decimal total; other types leave `sum` empty. |
| STAT-008 | Report mean for numeric columns | `integer` and `number` report the total divided by the non-null count; other types leave `mean` empty. |
| STAT-009 | Report median for numeric columns | `integer` and `number` report the middle sorted value, or the mean of the two middle values when the count is even. |

### Feature: Decimal Formatting

| ID | Story | High-level AC |
|---|---|---|
| FMT-001 | Write decimals to two places | `number` min, max, and sum, and every mean and median, have exactly two decimals, rounded half away from zero; integer min, max, and sum are plain integers. |

### Feature: Profile Output

| ID | Story | High-level AC |
|---|---|---|
| PROF-001 | Write the profile table | Output begins with the `column,type,rows,nulls,distinct,min,max,sum,mean,median` header and has one row per column in file order, with RFC 4180 quoting and `\n` line endings. |
| PROF-002 | Profile a single named column | `--column <name>` limits the output to that one column row. |

### Feature: Frequency Listing

| ID | Story | High-level AC |
|---|---|---|
| FREQ-001 | List value frequencies | `--freq` writes `value,count` rows for distinct non-null values, ordered by count descending then value ascending by byte value. |
| FREQ-002 | Limit frequency rows with --top | At most `--top` rows are listed, with a default of 5. |
| FREQ-003 | Report frequencies for no-value and absent columns | A column with no non-null values writes only the header and exits 0; an absent column exits 3. |

### Feature: Conformance

| ID | Story | High-level AC |
|---|---|---|
| CONF-001 | Pass the supplied conformance suite | `sh sources/run.sh` runs unfiltered against the built executable and exits zero with every check passing. |

## Story Guidance

Named stories planning must produce. Authoritative record: `STORY_GUIDANCE.json`.

| Story ID | Provenance | Gate | Note |
|---|---|---|---|
| INPUT-002 | plan | — | sources/run.sh E4: missing file exits 1 |
| CLI-001 | plan | — | sources/run.sh U1, U2, U6: usage errors with exit 2 |
| CLI-002 | plan | — | sources/run.sh U3, U4, U5: option combination and --top usage errors |
| INPUT-006 | plan | — | sources/run.sh E1: ragged.csv rejected with exit 3 |
| PROF-001 | plan | — | sources/run.sh section A: whole-file profiles A1 to A4 |
| PROF-002 | plan | — | sources/run.sh section B: single-column profiles B1 to B3, and E2 absent column |
| FREQ-001 | plan | — | sources/run.sh section C: frequency listings C1, C3 |
| FREQ-002 | plan | — | sources/run.sh section C: --top limits C2, C4 |
| CONF-001 | plan | — | sources/run.sh whole harness, run unfiltered |

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | CLI-001 | `--help` is not a defined option; it is rejected as an unknown option with exit 2. |
| AC-002 | INPUT-004 | A file of zero bytes, or one containing only blank lines, has no CSV records and is rejected as empty with exit 3. |
| AC-003 | STAT-007 | Integer totals are exact for any number of values and do not overflow. |
| AC-004 | FMT-001 | Mean, median, and decimal results are rounded from the exact value, not from binary floating point. |

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
| sources/INSTRUCTIONS.md | instruction-to-test | sources/run.sh | Command, Exit codes, and Output format sections map to checks A1–U6 | Stories derive from the brief; CONF-001 proves them |
| sources/run.sh | dependency | ./cmd/csvstat | Runs `go build -o <tmp> ./cmd/csvstat` before any check | Build path must be `./cmd/csvstat` (st-003) |
| sources/run.sh | test-kit-to-implementation | sources/employees.csv, sensors.csv, orders.csv, header_only.csv, ragged.csv | Fixtures resolved from the harness directory or `../sources` | Stage all five fixtures under `sources/` |
| sources/employees.csv | dependency | sources/run.sh | A1, B1, B2, C1, C4, E-series not used | Feeds TYPE, STAT, PROF, FREQ |
| sources/sensors.csv | dependency | sources/run.sh | A2, B3, C2, C5 | Feeds NULL-001, TYPE-004, FREQ-003 |
| sources/orders.csv | dependency | sources/run.sh | A3, C3, E2, E3, U2 | Padded text and invalid date exercise INPUT-003, TYPE-002 |
| sources/header_only.csv | dependency | sources/run.sh | A4 | Zero-row profile yields `empty` types |
| sources/ragged.csv | dependency | sources/run.sh | E1 (line 4 has an extra field) | Exercises INPUT-006 |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| sources/INSTRUCTIONS.md | author intent | compass | prompt-only |
| sources/run.sh | conformance harness | context | stage |
| sources/employees.csv | asset | context | stage |
| sources/sensors.csv | asset | context | stage |
| sources/orders.csv | asset | context | stage |
| sources/header_only.csv | asset | context | stage |
| sources/ragged.csv | asset | context | stage |

## Planning Instructions

### Delivery Shape

A single-command Go CLI. It takes one CSV file path plus options, reads the file with `encoding/csv`, and writes CSV to standard output. Diagnostics go to standard error. Exit codes are 0 (success), 1 (unreadable file), 2 (usage), and 3 (rejected input). The application root is a Go module whose main package is `./cmd/csvstat`, using the standard library only.

### Story Realization Map

- CLI-001 — argument parsing for file and options; INSTRUCTIONS "Command" and "Exit codes" (2); run.sh U1, U2, U6; capability.
- CLI-002 — option validation; INSTRUCTIONS options table; run.sh U3, U4, U5; capability.
- INPUT-001 — CSV reader; INSTRUCTIONS "Reading the file" 1; all fixtures; capability.
- INPUT-002 — open/read error mapping; INSTRUCTIONS Exit codes (1); run.sh E4; capability.
- INPUT-003 — trimming; INSTRUCTIONS "Reading the file" 2; orders.csv padded customers; sensors.csv battery value " "; capability.
- INPUT-004 — empty-file rejection; INSTRUCTIONS "Reading the file" 3; test-written files; capability.
- INPUT-005 — header validation; INSTRUCTIONS "Reading the file" 3; test-written files; capability.
- INPUT-006 — ragged rejection; INSTRUCTIONS "Reading the file" 3; ragged.csv line 4; run.sh E1; capability.
- NULL-001 — null rule; INSTRUCTIONS "Reading the file" 4; sensors.csv NULL/null/N/A; employees.csv NA; capability.
- TYPE-001 — integer and number inference; INSTRUCTIONS Types; run.sh A1 id and salary; capability.
- TYPE-002 — date inference; orders.csv order_date and shipped (2025-02-30 makes shipped text); run.sh A3; capability.
- TYPE-003 — boolean inference; employees.csv remote; run.sh A1; capability.
- TYPE-004 — text and empty fallback; sensors.csv battery; header_only.csv; run.sh A2, A4; capability.
- STAT-001 — rows and nulls; run.sh A1–A4; capability.
- STAT-002 — distinct; run.sh A1 name (23), A3 customer (5); capability.
- STAT-003 — numeric min and max; run.sh A1 id, salary; A2 temp_c; capability.
- STAT-004 — date min and max; run.sh A1 hire_date; A3 order_date; capability.
- STAT-005 — boolean min and max; run.sh A1 remote; capability.
- STAT-006 — text min and max; run.sh A1 notes; A3 ref ("-3", "A7"); capability.
- STAT-007 — sum; run.sh A1 salary 2085635.10; A3 unit_price; capability.
- STAT-008 — mean; run.sh A1 id 1012.50; A3 qty 6.87; capability.
- STAT-009 — median; run.sh A1 manager_id 4.00; A2 humidity 60.00; capability.
- FMT-001 — decimal rounding; run.sh A2 temp_c (-12.21); A3 qty; capability.
- PROF-001 — profile table; INSTRUCTIONS "Profile output" and "Output format"; run.sh A1–A4 headers; capability.
- PROF-002 — `--column`; run.sh B1–B3; E2 absent column; capability.
- FREQ-001 — frequency listing; run.sh C1 (Engineering before Support tie), C3; capability.
- FREQ-002 — `--top`; run.sh C2, C4; capability.
- FREQ-003 — no-value and absent frequency; run.sh C5, E3; capability.
- CONF-001 — terminal conformance gate; INSTRUCTIONS "Run Harness"; full run.sh; acceptance contract, Suite: full.

### Test and Acceptance Strategy

Each implementation story gets focused `go test` cases that write their own temporary files, and declares `Suite: scoped` with the slice it owns. The harness has no scope selector, so no story runs it. CONF-001 is the single terminal verification story. It depends on every other story and declares `Suite: full`, running `sh sources/run.sh` unfiltered. `bin/test.sh` (st-002) runs `go vet ./...` and `go test ./...` for the whole project.

### Sequencing and Dependencies

- Input reading (INPUT-001 to INPUT-006) precedes null handling and type inference.
- NULL-001 precedes TYPE-001 to TYPE-004.
- TYPE stories precede all STAT stories; STAT-008 and STAT-009 follow STAT-007.
- STAT stories precede FMT-001 and PROF-001.
- PROF-002 and FREQ-001 to FREQ-003 follow PROF-001 and the statistics they reuse.
- CLI-001 and CLI-002 may be built first.
- Fixtures under `sources/` are staged before any run that uses them.
- CONF-001 is last.

### Source Conflicts and Gaps

No cross-source conflicts were found, and no blockers or questions remain. Ambiguities were resolved narrowly as listed in Analysis Notes.

## Analysis Notes
generated: 2026-10-25
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.184203/workspace/targets/CsvStat/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 9
  stories: 29
  stack: Go 1.22+ (standard library only)
  display_name: not proposed
  short_description: not proposed

- Narrow resolutions taken under the unattended policy:
  - `--help` and `-h` are undefined and are rejected as unknown options with exit 2 (AC-001).
  - A file with zero bytes, or with only blank lines, has no CSV records and is empty (exit 3).
  - A leading UTF-8 byte-order mark is not stripped; `encoding/csv` keeps it in the first header name.
  - Integer totals use exact arithmetic, and decimals round from the exact value.
  - Date validity follows the Gregorian calendar, including leap years.
  - Text ordering and frequency ties use byte order of the trimmed string.
- Spot checks against the fixtures agree with the harness: 18 empty `coupon` values in orders.csv, the C1 tie order (Engineering before Support), and A3 `ref` min/max.
- Deployment target: the local `csvstat` executable built by `go build`, covered by st-001 and st-003. No hosted deployment is stated.
- Success measure: the supplied grader is the acceptance verdict (st-001). No numeric threshold is stated, so no measurement questions are raised.
- Help text baseline (cli): csvstat has one command and no sub-verbs; usage behavior is covered by AC-001 and CLI-001.
- Auth, PII, persistence, external calls, and concurrency do not apply: the CLI reads one local file and holds no shared state.
- Prior SEA_TRIALS.md fields `Testability`, `Consequence`, and its policy table are not in the current Sea Trials schema and are dropped. Criteria, IDs, and the wording of st-001 through st-008 are preserved. st-002 gains `Command: ["sh", "bin/test.sh"]`.
- st-009 is added as a guardrail from the brief's prohibition: application code never reads `sources/run.sh`.
- Story count is 29, below the 80 threshold, so no story-count questionnaire is raised.
- Identity: DISPLAY_NAME and SHORT_DESCRIPTION are already set, so no identity questionnaire is emitted.
