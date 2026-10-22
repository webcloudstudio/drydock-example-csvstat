# Blueprint Analysis: drydock CsvStat Uat Case

## Commander Expectations

- assert the CSV profiler builds as a standard-library Go CLI and passes every supplied acceptance check.

## Crew

| Crew | Charge |
|---|---|
| Commander | Defines intent and decides what done means. |
| Team Lead | Confirms epic completeness and stakeholder expectations. |
| Planning Crew | Authors atomic specifications and the ordered Manifest. |
| Shipyard Crew | Builds the tickets without synchronous Commander access. |

## Story List

### Feature: CLI Contract

| ID | Story | High-level AC |
|---|---|---|
| CLI-001 | Parse commands and options | Valid commands, option combinations, file cardinality, and usage failures map to the specified exit codes. |
| CLI-002 | Provide CLI help and diagnostics | The CLI exposes help for its commands and writes diagnostics only to standard error on failure. |

### Feature: CSV Ingestion

| ID | Story | High-level AC |
|---|---|---|
| CSV-001 | Read and normalize RFC 4180 CSV | Headers and values are trimmed, quoted fields are parsed correctly, and records are represented in file order. |
| CSV-002 | Validate input and resolve nulls | Empty files, invalid headers, ragged records, unreadable files, and absent columns are rejected with the specified exit behavior; null spellings are recognized. |

### Feature: Column Profiling

| ID | Story | High-level AC |
|---|---|---|
| PROF-001 | Infer column types | Each column is classified using the ordered integer, number, date, boolean, text, and empty rules. |
| PROF-002 | Compute column statistics | Profile output reports rows, nulls, distinct values, extrema, and numeric sum, mean, and median with specified rounding and ordering. |
| PROF-003 | Produce frequency reports | Frequency output excludes nulls and orders values by count descending, then bytewise value ascending, respecting `--top`. |

### Feature: Delivery Verification

| ID | Story | High-level AC |
|---|---|---|
| VERIFY-001 | Provide the project test gate | `bin/test.sh` runs `go vet ./...` and `go test ./...`, succeeding only when both pass. |
| VERIFY-002 | Pass the supplied conformance suite | The complete supplied `sources/run.sh` suite passes without modifying the supplied scoring assets. |

## Story Guidance

Named stories planning must produce. Authoritative record: `STORY_GUIDANCE.json`.

None.

## Surfaced Acceptance Criteria

| ID | Story ID | Criterion |
|---|---|---|
| AC-001 | CLI-002 | When invoked for a command or invalid invocation requiring assistance, the CLI shall provide usable help text describing the available command forms and options. |

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
| sources/INSTRUCTIONS.md | instruction-to-test | sources/run.sh | The instructions define the command, rules, output, and exit codes exercised by the harness. | Treat the instructions and harness as the normative contract. |
| sources/run.sh | conformance harness | sources/employees.csv, sources/sensors.csv, sources/orders.csv, sources/header_only.csv, sources/ragged.csv | The harness builds the executable and runs it against every staged fixture. | Stage the harness and all fixtures; terminal verification depends on the complete set. |
| sources/INSTRUCTIONS.md | parser-to-normalizer | CSV-001 | RFC 4180 parsing precedes trimming and null normalization. | Sequence CSV parsing before semantic normalization. |
| CSV-001 | dependency | CSV-002 | Validation requires parsed records and header width. | Validate normalized records before profiling. |
| CSV-002 | dependency | PROF-001 | Type inference operates only on validated, normalized values. | Rejected input must not reach profiling. |
| PROF-001 | dependency | PROF-002 | Statistics formatting depends on the inferred type. | Complete type inference before statistic computation. |
| PROF-002 | dependency | PROF-003 | Frequency reporting shares the validated column model and output writer. | Reuse ingestion and output behavior while keeping frequency semantics distinct. |
| VERIFY-001 | test-kit-to-implementation | CLI-001, CSV-002, PROF-002 | The project test gate verifies the Go implementation. | Build implementation before running local tests and vet. |
| VERIFY-002 | test-kit-to-implementation | CLI-001, CSV-001, CSV-002, PROF-001, PROF-002, PROF-003 | The supplied runner compares all required exit codes and stdout bytes. | Run the full harness only as the terminal acceptance gate. |

## Source Roles

| Path | Role | Plan disposition | Build disposition |
|---|---|---|---|
| sources/INSTRUCTIONS.md | normative specification and conformance test suite | prompt-only | prompt-only |
| sources/run.sh | conformance harness | stage | stage |
| sources/employees.csv | asset | stage | stage |
| sources/sensors.csv | asset | stage | stage |
| sources/orders.csv | asset | stage | stage |
| sources/header_only.csv | asset | stage | stage |
| sources/ragged.csv | asset | stage | stage |

## Planning Instructions

### Delivery Shape

The system is a local Go command-line application. It accepts one CSV file plus optional profiling or frequency arguments, parses and validates the file, normalizes values, infers column types, computes requested statistics, and writes CSV results to standard output. Failures write no standard output and use the specified exit code.

### Story Realization Map

| Story ID | Blueprint scope | Evidence | Related files | Delivery shape |
|---|---|---|---|---|
| CLI-001 | Argument parsing and exit-code contract | `sources/INSTRUCTIONS.md` Command, Options, Exit codes | `cmd/csvstat` | capability |
| CLI-002 | Help and stderr diagnostics | `sources/INSTRUCTIONS.md` Output format, CLI gap checklist | `cmd/csvstat` | capability |
| CSV-001 | RFC 4180 reader and trimming | `sources/INSTRUCTIONS.md` Reading the file | `cmd/csvstat` | capability |
| CSV-002 | Header, record, file, column, and null validation | `sources/INSTRUCTIONS.md` Reading the file, Exit codes | `cmd/csvstat` | capability |
| PROF-001 | Ordered type inference and calendar validation | `sources/INSTRUCTIONS.md` Types | `cmd/csvstat` | capability |
| PROF-002 | Profile statistics and exact CSV serialization | `sources/INSTRUCTIONS.md` Profile output, Decimals | `cmd/csvstat` | capability |
| PROF-003 | Frequency aggregation, ordering, and limiting | `sources/INSTRUCTIONS.md` Frequency output | `cmd/csvstat` | capability |
| VERIFY-001 | POSIX test wrapper | `sources/INSTRUCTIONS.md` Tests | `bin/test.sh` | test harness |
| VERIFY-002 | Complete supplied runner verification | `sources/run.sh` and all supplied CSV fixtures | `sources/run.sh` | acceptance contract |

### Test and Acceptance Strategy

Focused tests cover argument parsing, CSV quoting and trimming, null handling, validation, type inference, calendar dates, exact decimal rounding, profile serialization, frequency ordering, and exit codes. `bin/test.sh` is the local build-quality gate.

The terminal verification story is VERIFY-002 and is the only story that runs the complete supplied suite. Its acceptance assertion declares `Suite: full`. Other stories use focused tests or bounded story slices and do not run the complete runner.

### Sequencing and Dependencies

Implement and build the CLI contract first, then CSV parsing and normalization, then validation, type inference, statistics, frequencies, and output serialization. Add focused tests and the POSIX test gate after the relevant implementation exists. Stage fixtures before running VERIFY-002. VERIFY-002 depends on every implementation story.

### Source Conflicts and Gaps

No conflicting definitions or blockers were found. The only checklist gap was unspecified CLI help behavior; it is surfaced under CLI-002. No persistence, authentication, external service, or deployment infrastructure is described or required.

## Analysis Notes
generated: 2026-10-22
blueprint: /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261022.203215/workspace/targets/CsvStat/blueprint

Quality: Ready
  blockers: 0
  questions: 0
  features: 4
  stories: 9
  stack: Go 1.22+ with standard library
  display_name: not proposed
  short_description: not proposed

No unresolved analysis issues.
