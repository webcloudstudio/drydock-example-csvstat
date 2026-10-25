# MANIFEST: CsvStat
updated:     2026-10-25T01:25:19+00:00
state:       approved
blocks:      3

applied_specs: |
  ARCHITECTURE.md sha256=e7aa9fbae2d0c1f547ee31aff3dfa8b5c4078dc57caa2db732820776d83011f7 commit=41f674211cc58e541d7923d9ee92147534049d03 applied_by=block-1 applied_at=2026-10-25T01:28:47+00:00 build_sha256=f9bae974a7d3a670f4d704d0ceb36567d9ed5533a4b4d11576f028ff009ed3ee

## story 1: Stand up the Go module, the csvstat entry point, and the testable CLI package.
id:           foundation
finding: ADVISORY: implemented, not verified — no governed acceptance command covers this story. Declare one in ACCEPTANCE.json to gate it.
evidence: evidence/block-1.md
summary:      Stand up the Go module, the csvstat entry point, and the testable CLI package.
type:         foundational
kind:         capability
phase:        1
block:        1
implements:   ARCHITECTURE.md
instructions: |
  Create the root go.mod (Go 1.22 or newer, standard library only), cmd/csvstat/main.go, and the
  internal/cli package exposing Run(args []string, stdout, stderr io.Writer) int. Only main calls
  os.Exit. Define the package layout under internal/ for reading, typing, statistics, and output.
  Run is a minimal stub at this step; later stories fill it in. Record the module ownership table
  and the technology stack in ARCHITECTURE.md.
stack:        common.md, go.md
stack_mode:   builder
size:         13244
provides:     go.mod, cmd/csvstat, internal/cli.Run
acceptance:   yes
state: closed/implemented

## story 2: Build the csvstat executable with go build using only the standard library.
id: build-executable
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Build the csvstat executable with go build using only the standard library.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-Build.md
covers: BLD-001
accepts: st-003
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6419
provides: csvstat executable
instructions: |
  Make go build -o <path> ./cmd/csvstat succeed from the application root with no third-party
  modules. main wires os.Args, os.Stdout, and os.Stderr into cli.Run and exits with its result.
acceptance: yes
depends: foundation
state: closed/failed

## story 3: Read the CSV file with encoding/csv and trim the header and every value.
id: csv-reader
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Read the CSV file with encoding/csv and trim the header and every value.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-CsvReader.md
covers: INP-001
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6711
provides: internal/csvio.Read
instructions: |
  Read RFC 4180 CSV with encoding/csv. The first record is the header. Trim leading and trailing
  spaces from every header name and every value. Do not reject invalid UTF-8. Blank lines are
  skipped by encoding/csv. Keep record line numbers available for diagnostics.
acceptance: yes
depends: foundation
state: closed/failed

## story 4: Report an unopenable or unreadable file, including a directory, with exit 1.
id: read-failure
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Report an unopenable or unreadable file, including a directory, with exit 1.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ReadFailure.md
covers: INP-003
accepts: st-008
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6141
provides: exit code 1 for operational failure
consumes: internal/csvio.Read
instructions: |
  A path that cannot be opened or read, including a missing file and a directory, exits 1 with
  nothing on standard output and one diagnostic line on standard error. It must not panic.
acceptance: yes
depends: csv-reader
state: closed/failed

## story 5: Classify empty, NA, N/A, and NULL values as null after trimming.
id: null-values
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Classify empty, NA, N/A, and NULL values as null after trimming.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-NullValues.md
covers: INP-004
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6153
provides: internal/profile.IsNull
consumes: internal/csvio.Read
instructions: |
  A trimmed value is null when it is empty or equals NA, N/A, or NULL ignoring case. Every other
  value is non-null. Expose the predicate for typing, counts, statistics, and frequencies.
acceptance: yes
depends: csv-reader
state: closed/failed

## story 6: Infer the integer and number column types.
id: type-numeric
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Infer the integer and number column types.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-TypeNumeric.md
covers: TYP-001
accepts: st-004
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6382
provides: internal/profile.InferType
consumes: internal/profile.IsNull
instructions: |
  A column is integer when every non-null value matches -?[0-9]+ and number when every non-null
  value matches -?[0-9]+(\.[0-9]+)?. Typing considers only non-null values and picks the first
  type in order integer, number, date, boolean, text that all of them satisfy.
acceptance: yes
depends: null-values
state: closed/failed

## story 7: Infer the date column type with real calendar validation.
id: type-date
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Infer the date column type with real calendar validation.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-TypeDate.md
covers: TYP-002
accepts: st-004
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6028
provides: date type
consumes: internal/profile.InferType
instructions: |
  A column is date when every non-null value matches YYYY-MM-DD and is a real calendar date, so
  2025-02-30 is not a date. Use time.Parse with the 2006-01-02 layout plus a strict shape check.
acceptance: yes
depends: type-numeric
state: closed/failed

## story 8: Compute the exact median for numeric columns.
id: median
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Compute the exact median for numeric columns.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-Median.md
covers: PRO-004
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6057
provides: median
consumes: internal/profile.InferType
instructions: |
  Sort the non-null numeric values exactly. The median is the middle value, or the mean of the
  two middle values when the count is even, written with two decimals rounded half away from
  zero. Other types are blank.
acceptance: yes
depends: type-numeric
state: closed/failed

## story 9: Count rows, nulls, and distinct non-null values for each column.
id: counts
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Count rows, nulls, and distinct non-null values for each column.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-Counts.md
covers: PRO-001
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6070
provides: rows, nulls, distinct
consumes: internal/profile.IsNull
instructions: |
  rows counts data records, nulls counts null values, and distinct counts different non-null
  values compared as exact trimmed strings.
acceptance: yes
depends: type-numeric
state: closed/failed

## story 10: Infer the boolean, text, and empty column types.
id: type-other
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Infer the boolean, text, and empty column types.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-TypeOther.md
covers: TYP-003
accepts: st-004
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6162
provides: boolean, text, empty types
consumes: internal/profile.InferType
instructions: |
  A column is boolean when every non-null value is true or false ignoring case, text otherwise,
  and empty when it has no non-null value.
acceptance: yes
depends: type-numeric
state: closed/failed

## story 11: Compute exact sum and mean for numeric columns, rounded once half away from zero.
id: sum-mean
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Compute exact sum and mean for numeric columns, rounded once half away from zero.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-SumMean.md
covers: PRO-003
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6281
provides: sum, mean
consumes: internal/profile.InferType
instructions: |
  Use math/big (big.Rat) for exact arithmetic. integer sum is a plain integer. number sum and
  mean are written with two decimals, rounded once half away from zero from the exact value,
  with a leading - when negative and no thousands separator. mean is the total divided by the
  non-null count; integer mean is also two decimals. Other types are blank.
acceptance: yes
depends: type-numeric
state: closed/failed

## story 12: Report min and max for every column type.
id: min-max
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Report min and max for every column type.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-MinMax.md
covers: PRO-002
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6405
provides: min, max
consumes: internal/profile.InferType
instructions: |
  integer and number compare numerically using exact math/big values and are written as plain
  integers or two-decimal numbers. date gives earliest and latest. boolean gives false if any
  value is false else true for min, and true if any value is true else false for max. text
  compares by byte value. empty is blank.
acceptance: yes
depends: type-date, type-other
state: closed/failed

## story 13: Reject empty, malformed, or ragged files with exit 3 and a diagnostic on standard error.
id: input-rejection
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Reject empty, malformed, or ragged files with exit 3 and a diagnostic on standard error.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-InputRejection.md
covers: INP-002
accepts: st-007
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6870
provides: exit code 3 for rejected input
consumes: internal/csvio.Read
instructions: |
  Reject an empty file, an empty or repeated header name, and any record whose field count differs
  from the header. Exit 3, write nothing to standard output, and write one standard error line
  naming the file and, for a record, its line number. A csv.ParseError is also a rejection.
acceptance: yes
depends: csv-reader
state: closed/failed

## story 14: Parse arguments and reject malformed usage with exit 2 and a usage line.
id: arguments
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Parse arguments and reject malformed usage with exit 2 and a usage line.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-Arguments.md
covers: ARG-001
accepts: st-008
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 7469
provides: --column, --freq, --top, exit code 2
instructions: |
  Use a local flag.FlagSet with flag.ContinueOnError and discarded flag output. Options precede
  the file. Exit 2 with nothing on standard output for a missing or extra file, an unknown option,
  an option without its value, --column with --freq, and --top without --freq. Write one standard
  error line naming the profile and frequency command forms. No --help option.
acceptance: yes
depends: build-executable
state: closed/failed

## story 15: Write the profile table for every column in file order.
id: profile-output
finding: programmatic acceptance failed: profile-quoting-and-trim: AssertionError: (3, '', 'csvstat: rejected input: /tmp/tmp9phg2ko5/data.csv: parse error on line 2, column 3: bare " in non-quoted-field\n')
evidence: evidence/block-2.md
summary: Write the profile table for every column in file order.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ProfileOutput.md
covers: OUT-001
accepts: st-005
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 7622
provides: profile table on standard output
consumes: internal/profile
instructions: |
  Write the header column,type,rows,nulls,distinct,min,max,sum,mean,median and one row per
  column in file order using encoding/csv's writer with \n line endings. Quoting follows the
  writer. Nothing is written to standard output on failure.
acceptance: yes
depends: counts, min-max, sum-mean, median, type-date, type-other
state: closed/failed

## story 16: Profile only the column named by --column and exit 3 when it is absent.
id: column-select
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Profile only the column named by --column and exit 3 when it is absent.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-ColumnSelect.md
covers: SEL-001
accepts: st-007
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6605
provides: --column behavior
consumes: profile table
instructions: |
  With --column, output the header and the single named column row. A name absent from the
  header exits 3 with nothing on standard output.
acceptance: yes
depends: profile-output, arguments, input-rejection
state: closed/failed

## story 17: List the most frequent values of a column ordered by count then value.
id: frequency
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: List the most frequent values of a column ordered by count then value.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-Frequency.md
covers: FRQ-001
accepts: st-006, st-007
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 7259
provides: --freq behavior
consumes: internal/profile.IsNull
instructions: |
  With --freq, write the header value,count and one row per distinct non-null value ordered by
  count descending, then value ascending by bytes. Nulls are not listed. An all-null column
  writes only the header. An absent column name exits 3. The default limit is 5 rows.
acceptance: yes
depends: null-values, profile-output, arguments, input-rejection
state: closed/failed

## story 18: Limit frequency rows with --top and reject non-positive values with exit 2.
id: frequency-top
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Limit frequency rows with --top and reject non-positive values with exit 2.
type: service
kind: capability
phase: 1
block: 2
implements: FEATURE-FrequencyTop.md
covers: FRQ-002
accepts: st-006, st-008
context: ARCHITECTURE_compact.md
stack: go.md
stack_mode: consumer
size: 6361
provides: --top behavior
consumes: --freq behavior
instructions: |
  --top n lists at most n frequency rows, default 5. A value that is not a positive integer
  exits 2 with nothing on standard output.
acceptance: yes
depends: frequency
state: closed/failed

## story 19: Provide the POSIX bin/test.sh that runs go vet and go test.
id: test-runner
finding: step block-2 failed: profile-quoting-and-trim
evidence: evidence/block-2.md
summary: Provide the POSIX bin/test.sh that runs go vet and go test.
type: service
kind: test harness
phase: 1
block: 2
implements: FEATURE-TestRunner.md
covers: BLD-002
accepts: st-002
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 8025
provides: bin/test.sh
instructions: |
  Write bin/test.sh in POSIX sh. From the application root it runs go vet ./... and then
  go test ./..., exiting zero only when both pass. The project test suite covers each type, each
  statistic, null handling, rounding, each output form, and each exit code, using files the tests
  write themselves. Tests call cli.Run with bytes.Buffer writers.
acceptance: yes
depends: frequency-top, column-select, read-failure, build-executable
state: closed/failed

## story 20: Pass the supplied conformance grader sh sources/run.sh.
id: conformance
summary: Pass the supplied conformance grader sh sources/run.sh.
type: feature
kind: integration
phase: 1
block: 3
implements: FEATURE-Conformance.md
covers: CNF-001
accepts: st-001
context: ARCHITECTURE_compact.md
stack: common.md, go.md
stack_mode: consumer
size: 7358
provides: conformance with sources/run.sh
instructions: |
  Assembly story. Once every implementation story has closed, run the supplied grader unfiltered
  and make it exit zero. Declare Suite: full on its single assertion and assert only the exit
  status. Fix defects in the member stories' code, never in sources/.
acceptance: yes
depends: test-runner, frequency-top, column-select, read-failure
state: pending
