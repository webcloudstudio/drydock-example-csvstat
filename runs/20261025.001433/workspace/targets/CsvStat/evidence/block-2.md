# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-10-24
- resulting state: closed/verified
- story points (combined assembled cost): 49094
- execution id: 20261025.003203.441Z-8a0c66c8

## Stories built
- Accept the two csvstat command forms and the --help text. (cli-args) [story]
- Map an unreadable file to exit code 1. (cli-unreadable) [story]
- Read a well-formed CSV and write the profile header and one row per column. (input-profile) [story]
- Trim leading and trailing spaces from every header name and value. (input-trim) [story]
- Map usage errors to exit code 2. (cli-usage) [story]
- Reject empty files, bad headers, ragged records, and CSV syntax errors with exit 3. (input-reject) [story]
- Report rows, nulls, and distinct counts per column. (counts) [story]
- Infer date columns and reject invalid calendar dates. (types-date) [story]
- Infer integer and number column types. (types-numeric) [story]
- List the most frequent values of a column with --freq. (freq-list) [story]
- Limit frequency output with --top. (freq-top) [story]
- Infer boolean, text, and empty column types and compose the final type decision. (types-other) [story]
- Report min and max for every column type. (stats-minmax) [story]
- Report the median for integer and number columns. (stats-median) [story]
- Report sum and mean for integer and number columns with exact arithmetic. (stats-summean) [story]
- Write numeric statistics with two decimals, rounded half away from zero once. (stats-format) [story]
- Profile one named column with --column. (select-column) [story]
- Reject a --column or --freq name absent from the header with exit 3. (select-absent) [story]
- Provide a POSIX bin/test.sh that runs go vet and go test. (qa-testrunner) [story]

## Acceptance tooling authorization
- FEATURE-CliParsing.md#cliparse-help-test: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliParsing.md#cliparse-column-form-test: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliParsing.md#cliparse-freq-form-test: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UnreadableFile.md#unreadable-missing-path: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-UnreadableFile.md#unreadable-directory: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-UnreadableFile.md#unreadable-with-options: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-UnreadableFile.md#unreadable-diagnostic-on-stderr: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-ValueTrimming.md#trim-header-test: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ValueTrimming.md#trim-values-test: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-missing-or-extra-file: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-UsageErrors.md#usage-unknown-option: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-UsageErrors.md#usage-missing-option-value: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-UsageErrors.md#usage-bad-top: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-UsageErrors.md#usage-top-without-freq: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-UsageErrors.md#usage-column-with-freq: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-UsageErrors.md#usage-before-file-access: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-RejectMalformed.md#reject-empty-file-test: executable=go; scope=test; authorization=existing Target environment
- FEATURE-RejectMalformed.md#reject-bad-header-test: executable=go; scope=test; authorization=existing Target environment
- FEATURE-RejectMalformed.md#reject-ragged-test: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Frequency.md#freq-department: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-Frequency.md#freq-tie-order: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-Frequency.md#freq-trims-and-skips-nulls: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-Frequency.md#freq-all-null-column: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-Frequency.md#freq-output-line-endings: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-FrequencyTop.md#top-default-is-five: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-FrequencyTop.md#top-limits-rows: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-FrequencyTop.md#top-larger-than-distinct: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-FrequencyTop.md#top-rejects-non-positive: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-FrequencyTop.md#top-requires-freq: python-package=drydock_go; scope=test; authorization=Drydock test environment
- FEATURE-Median.md#median-odd-and-even-integers: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Median.md#median-number-exact-rounding: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Median.md#median-nulls-skipped: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Median.md#median-blank-for-other-types: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Median.md#median-fixture-sensors-orders: executable=go; scope=test; authorization=existing Target environment
- FEATURE-SumMean.md#summean-integer-exact: executable=go; scope=test; authorization=existing Target environment
- FEATURE-SumMean.md#summean-integer-bigint: executable=go; scope=test; authorization=existing Target environment
- FEATURE-SumMean.md#summean-number-exact-decimal: executable=go; scope=test; authorization=existing Target environment
- FEATURE-SumMean.md#summean-nulls-skipped: executable=go; scope=test; authorization=existing Target environment
- FEATURE-SumMean.md#summean-blank-for-other-types: executable=go; scope=test; authorization=existing Target environment
- FEATURE-SumMean.md#summean-fixture-employees: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DecimalFormat.md#decimalformat-number-columns: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DecimalFormat.md#decimalformat-integer-columns: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DecimalFormat.md#decimalformat-single-rounding: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DecimalFormat.md#decimalformat-no-negative-zero: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DecimalFormat.md#decimalformat-fixture-files: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnSelect.md#colselect-matches-full-profile-row: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnSelect.md#colselect-only-named-column: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnSelect.md#colselect-trimmed-header-and-values: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnSelect.md#colselect-output-format: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnSelect.md#colselect-header-only-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnAbsent.md#colabsent-column-exit-3: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnAbsent.md#colabsent-freq-exit-3: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnAbsent.md#colabsent-case-sensitive: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnAbsent.md#colabsent-present-name-still-succeeds: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnAbsent.md#colabsent-header-only-and-ragged: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TestRunner.md#testrunner-syntax: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-TestRunner.md#testrunner-runs-vet-and-test: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-TestRunner.md#testrunner-vet-failure: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-TestRunner.md#testrunner-test-failure: executable=sh; scope=test; authorization=existing Target environment
- FEATURE-TestRunner.md#testrunner-works-from-other-directory: executable=sh; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 963)
- implements: FEATURE-CliParsing.md (SP 1319)
- context: ARCHITECTURE_compact.md (SP 381)
- implements: FEATURE-UnreadableFile.md (SP 1220)
- implements: FEATURE-CsvReading.md (SP 1778)
- implements: FEATURE-ValueTrimming.md (SP 1211)
- implements: FEATURE-UsageErrors.md (SP 1932)
- implements: FEATURE-RejectMalformed.md (SP 2172)
- implements: FEATURE-ProfileCounts.md (SP 2486)
- implements: FEATURE-TypesDate.md (SP 1676)
- implements: FEATURE-TypesNumeric.md (SP 1987)
- implements: FEATURE-Frequency.md (SP 2200)
- implements: FEATURE-FrequencyTop.md (SP 2038)
- implements: FEATURE-TypesOther.md (SP 2227)
- implements: FEATURE-MinMax.md (SP 3193)
- implements: FEATURE-Median.md (SP 3169)
- implements: FEATURE-SumMean.md (SP 3681)
- implements: FEATURE-DecimalFormat.md (SP 3634)
- implements: FEATURE-ColumnSelect.md (SP 2409)
- implements: FEATURE-ColumnAbsent.md (SP 1892)
- implements: FEATURE-TestRunner.md (SP 2059)
- stack: common_compact.md (SP 1653)

## Build directory changes
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/cli/modes.go
- internal/cli/parse_test.go
- internal/cli/request.go
- internal/csvio/csvio.go
- internal/csvio/csvio_test.go
- internal/decimal/decimal.go
- internal/decimal/decimal_test.go
- internal/profile/freq.go
- internal/profile/profile.go
- internal/profile/profile_test.go
- internal/profile/types.go

## Pre-build acceptance observation
- RED: cliparse-help-exits-zero (FEATURE-CliParsing.md)
  intent: csvstat --help exits 0 and writes usage naming the command and every option to standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: cliparse-help-test (FEATURE-CliParsing.md)
  intent: The package test for the help form passes.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: cliparse-column-form-test (FEATURE-CliParsing.md)
  intent: Parsing the --column form yields a profile request with the named column and the file path.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: cliparse-freq-form-test (FEATURE-CliParsing.md)
  intent: Parsing the --freq form yields a frequency request with the name, default top of 5, and the file path.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: unreadable-missing-path (FEATURE-UnreadableFile.md)
  intent: A path that does not exist exits 1 with empty standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: unreadable-directory (FEATURE-UnreadableFile.md)
  intent: A path that is a directory exits 1 with empty standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: unreadable-with-options (FEATURE-UnreadableFile.md)
  intent: The --column and --freq forms also exit 1 for an unreadable file.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: unreadable-diagnostic-on-stderr (FEATURE-UnreadableFile.md)
  intent: The failure is reported on standard error and names the file.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: csvread-header-and-columns (FEATURE-CsvReading.md)
  intent: Profiling a well-formed file exits 0 and writes the declared header then one ten-field row per input column in file order.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: csvread-line-endings (FEATURE-CsvReading.md)
  intent: Profile output uses newline line endings only.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: csvread-quoted-comma (FEATURE-CsvReading.md)
  intent: A quoted field containing a comma is one value, so the file keeps its declared column count.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: csvread-header-only (FEATURE-CsvReading.md)
  intent: A file with a header and no data records succeeds with one row per header column.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: csvread-bom-kept (FEATURE-CsvReading.md)
  intent: A byte-order mark is not stripped from the first header name.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: trim-header-names (FEATURE-ValueTrimming.md)
  intent: Padded header names are written trimmed in the profile output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: trim-only-spaces (FEATURE-ValueTrimming.md)
  intent: Only spaces are trimmed; a trailing tab stays part of the header name.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: trim-header-test (FEATURE-ValueTrimming.md)
  intent: The package test for header trimming passes.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: trim-values-test (FEATURE-ValueTrimming.md)
  intent: The package test for value trimming passes.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: usage-missing-or-extra-file (FEATURE-UsageErrors.md)
  intent: No file, or more than one file, exits 2 with empty standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: usage-unknown-option (FEATURE-UsageErrors.md)
  intent: An unknown option exits 2 with empty standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: usage-missing-option-value (FEATURE-UsageErrors.md)
  intent: Each value-taking option given as the last argument exits 2 with empty standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: usage-bad-top (FEATURE-UsageErrors.md)
  intent: A --top of zero or non-numeric text exits 2 with empty standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: usage-top-without-freq (FEATURE-UsageErrors.md)
  intent: --top without --freq exits 2 with empty standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: usage-column-with-freq (FEATURE-UsageErrors.md)
  intent: --column combined with --freq exits 2 with empty standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: usage-before-file-access (FEATURE-UsageErrors.md)
  intent: A usage error on a nonexistent path is exit 2, not exit 1, because arguments are validated before the file is opened.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: reject-ragged-supplied (FEATURE-RejectMalformed.md)
  intent: A record with a different field count than the header exits 3 with nothing on standard output and the file named on standard error.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: reject-empty-file (FEATURE-RejectMalformed.md)
  intent: A zero-byte file and a file holding only a blank line both exit 3 with nothing on standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: reject-bad-header-names (FEATURE-RejectMalformed.md)
  intent: An empty header name, a repeated header name, and a name repeated only after trimming each exit 3 with nothing on standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: reject-short-record (FEATURE-RejectMalformed.md)
  intent: A record with fewer fields than the header exits 3 with nothing on standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: reject-csv-syntax-error (FEATURE-RejectMalformed.md)
  intent: A bare quote inside an unquoted field is a CSV syntax error that exits 3 with nothing on standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: reject-valid-files-pass (FEATURE-RejectMalformed.md)
  intent: Well-formed files, including a header-only file, are not rejected.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: reject-empty-file-test (FEATURE-RejectMalformed.md)
  intent: The package test for empty-file rejection passes.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: reject-bad-header-test (FEATURE-RejectMalformed.md)
  intent: The package test for header-name rejection passes.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: reject-ragged-test (FEATURE-RejectMalformed.md)
  intent: The package test for ragged-record rejection passes.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: counts-rows-nulls-distinct (FEATURE-ProfileCounts.md)
  intent: rows, nulls, and distinct are computed from the file, with padded values trimmed before comparison.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: counts-null-spellings (FEATURE-ProfileCounts.md)
  intent: Empty, NA, N/A, and NULL in any letter case are null and are excluded from distinct.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: counts-non-null-lookalikes (FEATURE-ProfileCounts.md)
  intent: Values that merely resemble a null spelling are not null and count as distinct values.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: counts-all-null-column (FEATURE-ProfileCounts.md)
  intent: A column whose values are all null reports distinct 0 and nulls equal to rows.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: counts-header-only (FEATURE-ProfileCounts.md)
  intent: A file with a header and no data records reports zero counts for every column and exits 0.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: types-date-valid (FEATURE-TypesDate.md)
  intent: A column of real YYYY-MM-DD dates, including a leap day and nulls, is typed date.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: types-date-invalid-calendar (FEATURE-TypesDate.md)
  intent: Values that look like dates but are not real calendar dates keep the column from being date.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: types-date-wrong-shape (FEATURE-TypesDate.md)
  intent: Dates not written exactly as zero-padded YYYY-MM-DD do not qualify.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: types-numeric-integer (FEATURE-TypesNumeric.md)
  intent: A column whose non-null values all match the integer pattern is typed integer, including negatives, leading zeros, and nulls.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: types-numeric-number (FEATURE-TypesNumeric.md)
  intent: A column with at least one decimal value and otherwise integer or decimal values is typed number.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: types-numeric-rejects-lookalikes (FEATURE-TypesNumeric.md)
  intent: Each value that fails the anchored numeric patterns keeps its column from being integer or number.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: types-numeric-integer-preferred (FEATURE-TypesNumeric.md)
  intent: An all-integer column is integer, not number, and adding one decimal changes only that column to number.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: freq-department (FEATURE-Frequency.md)
  intent: Frequencies for a text column are counted, ordered by count then byte value, and parsed back as CSV.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: freq-tie-order (FEATURE-Frequency.md)
  intent: Values with equal counts are ordered ascending by byte value and the limit is honored.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: freq-trims-and-skips-nulls (FEATURE-Frequency.md)
  intent: Padded values are trimmed before counting and null spellings (blank, NA) are not listed.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: freq-all-null-column (FEATURE-Frequency.md)
  intent: A column with no non-null values writes only the header and exits 0.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: freq-output-line-endings (FEATURE-Frequency.md)
  intent: Frequency output uses \n line endings only and ends with a newline.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: top-default-is-five (FEATURE-FrequencyTop.md)
  intent: Without --top, a column with many distinct values lists exactly five rows.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: top-limits-rows (FEATURE-FrequencyTop.md)
  intent: An explicit --top n writes at most n rows, matching the ordered counts derived from the file.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: top-larger-than-distinct (FEATURE-FrequencyTop.md)
  intent: A limit above the distinct count lists every distinct non-null value.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: top-rejects-non-positive (FEATURE-FrequencyTop.md)
  intent: A --top that is not a positive integer exits 2 with nothing on standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: top-requires-freq (FEATURE-FrequencyTop.md)
  intent: --top without --freq is a usage error with exit 2 and empty standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: types-other-boolean (FEATURE-TypesOther.md)
  intent: Columns of true and false in any letter case, with nulls, are typed boolean.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: types-other-not-boolean (FEATURE-TypesOther.md)
  intent: Boolean-like words that are not true or false make the column text.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: types-other-text-fallback (FEATURE-TypesOther.md)
  intent: A column that mixes types, or holds free text, is typed text.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: types-other-empty (FEATURE-TypesOther.md)
  intent: A column with no non-null value is typed empty, whatever spellings of null it holds.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: types-other-header-only-empty (FEATURE-TypesOther.md)
  intent: Every column of a header-only file is typed empty.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: minmax-integer (FEATURE-MinMax.md)
  intent: Integer min and max are compared numerically and written as plain integers.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: minmax-integer-huge (FEATURE-MinMax.md)
  intent: Integer min and max stay exact for digit strings beyond 64-bit range.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: minmax-number (FEATURE-MinMax.md)
  intent: Number min and max are chosen by exact decimal comparison and equal the supplied values.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: minmax-date (FEATURE-MinMax.md)
  intent: Date min and max are the earliest and latest dates.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: minmax-boolean (FEATURE-MinMax.md)
  intent: Boolean min is false when any value is false and max is true when any value is true; a single-valued column reports that value for both.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: minmax-text-bytes (FEATURE-MinMax.md)
  intent: Text min and max compare by byte value, so upper-case sorts before lower-case, and padding is trimmed.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: minmax-empty-and-nulls (FEATURE-MinMax.md)
  intent: An empty column reports blank min and max, and nulls are ignored when choosing extremes.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: median-odd-and-even-integers (FEATURE-Median.md)
  intent: The median is the middle value of an odd count and the mean of the two middle values of an even count, regardless of input order.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: median-number-exact-rounding (FEATURE-Median.md)
  intent: A number median that falls on a tie is rounded half away from zero from the exact value, for positive and negative values.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: median-nulls-skipped (FEATURE-Median.md)
  intent: Null values are excluded when counting and ordering values for the median.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: median-blank-for-other-types (FEATURE-Median.md)
  intent: Text, date, boolean, and empty columns leave median blank.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: median-fixture-sensors-orders (FEATURE-Median.md)
  intent: The supplied sensors and orders files report medians computed independently from their own values.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: summean-integer-exact (FEATURE-SumMean.md)
  intent: An integer column reports the exact integer sum and the two-decimal mean of its values.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: summean-integer-bigint (FEATURE-SumMean.md)
  intent: Integer sums beyond 64 bits are computed with arbitrary precision and never overflow.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: summean-number-exact-decimal (FEATURE-SumMean.md)
  intent: Number sums and means come from exact decimal values and are rounded half away from zero once.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: summean-nulls-skipped (FEATURE-SumMean.md)
  intent: Null values are excluded from both the sum and the divisor of the mean.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: summean-blank-for-other-types (FEATURE-SumMean.md)
  intent: Text, date, boolean, and empty columns leave sum and mean blank.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: summean-fixture-employees (FEATURE-SumMean.md)
  intent: The supplied employees file reports the exact sum and mean computed independently from its own values.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: decimalformat-number-columns (FEATURE-DecimalFormat.md)
  intent: Every number min, max, sum, mean, and median is written with exactly two decimals rounded half away from zero from the exact value.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: decimalformat-integer-columns (FEATURE-DecimalFormat.md)
  intent: Integer min, max, and sum are plain integers while integer mean and median carry exactly two decimals.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: decimalformat-single-rounding (FEATURE-DecimalFormat.md)
  intent: Rounding happens once from the exact value, so a value that double rounding would push up stays down.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: decimalformat-no-negative-zero (FEATURE-DecimalFormat.md)
  intent: A negative value that rounds to zero is written as an unsigned zero with two decimals.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: decimalformat-fixture-files (FEATURE-DecimalFormat.md)
  intent: Number and integer columns of the supplied files are written in the canonical shape from values computed independently of the application.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: colselect-matches-full-profile-row (FEATURE-ColumnSelect.md)
  intent: For every column of the supplied files, --column writes the profile header and exactly the row the full profile writes for that column.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: colselect-only-named-column (FEATURE-ColumnSelect.md)
  intent: The output of --column contains one data row whose first cell is the name supplied, even for the last column of the file.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: colselect-trimmed-header-and-values (FEATURE-ColumnSelect.md)
  intent: A padded header is selected by its trimmed name and the padded values are trimmed before they are profiled.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: colselect-output-format (FEATURE-ColumnSelect.md)
  intent: Selected-column output uses newline line endings only and quotes a field containing a comma as one value.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: colselect-header-only-file (FEATURE-ColumnSelect.md)
  intent: Selecting a column of a file with a header and no rows succeeds and reports zero rows.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: colabsent-column-exit-3 (FEATURE-ColumnAbsent.md)
  intent: --column with a name absent from the header exits 3 and writes nothing to standard output.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: colabsent-freq-exit-3 (FEATURE-ColumnAbsent.md)
  intent: --freq with a name absent from the header exits 3 and writes nothing to standard output, with and without --top.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: colabsent-case-sensitive (FEATURE-ColumnAbsent.md)
  intent: A name that differs from a real header name only by letter case is absent and is rejected.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: colabsent-present-name-still-succeeds (FEATURE-ColumnAbsent.md)
  intent: Every real header name is accepted by --column and --freq, so rejection is specific to absent names.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: colabsent-header-only-and-ragged (FEATURE-ColumnAbsent.md)
  intent: An absent name is rejected for a header-only file, and a ragged file is rejected with exit 3 whatever name is given.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: testrunner-syntax (FEATURE-TestRunner.md)
  intent: bin/test.sh exists and passes a POSIX syntax check.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: testrunner-runs-vet-and-test (FEATURE-TestRunner.md)
  intent: With a passing toolchain stub, bin/test.sh exits zero and invokes both go vet and go test.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: testrunner-vet-failure (FEATURE-TestRunner.md)
  intent: When go vet fails, bin/test.sh exits non-zero.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: testrunner-test-failure (FEATURE-TestRunner.md)
  intent: When go test fails after a passing go vet, bin/test.sh exits non-zero.
  error: baseline unavailable: authorized Target tooling is not provisioned
- RED: testrunner-works-from-other-directory (FEATURE-TestRunner.md)
  intent: bin/test.sh locates the application root itself, so a stubbed run from another directory still invokes go vet and go test.
  error: baseline unavailable: authorized Target tooling is not provisioned

## Build summary
The build directory already has a block-1 skeleton: `cli.go` is a stub and there are no real packages yet. I'll read the two files I'm going to replace, then write the packages.Starting with the input layer: `csvio` loads, trims, and validates the file.Now the CLI: request parsing, the two dispatch interfaces, and `Run`.Now the test suites. I'm writing them beside the code, and each one reads back the public result instead of matching output text.I left a garbled expression in `TestTrimValues`. Fixing it so the expected value is plainly `"\t y"`: the leading tab stays and the trailing space goes.Now I'll format, vet, and run the whole suite.The gofmt changes are mine from the format pass, so they're expected. The real error is "illegal byte order mark" at `csvio_test.go:115`, so I'll look at that line.The BOM went into the source as a raw byte instead of an escape. I'll replace it with a `\ufeff` escape.Everything passes except one test: I fed `Summarize` an untrimmed `" "`. Real input is trimmed upstream, so the test value should be `"NULL"`.The binary behaves as specified on every case I ran. I'll cross-check the numeric fixtures against an independent Python `Fraction` computation, since that's what the acceptance assertions use.The block is implemented. `go vet` and `go test -race ./...` pass across all four packages, and `sh bin/test.sh` exits 0.

**Verification gap:** I did not run the declared Python acceptance checks, because the `drydock_go` helper (`cli_runner`, `run_test`) isn't in this environment. I also didn't run `sources/run.sh`, which the COMPASS reserves for the terminal story. Instead I exercised the built binary by hand. The fixture numbers I checked (employees salary sum/mean/median, sensors temp_c count/min/max/sum/mean) match an independent Python `Fraction` computation.

**Choices I made where the spec was open:**
- `--help` is a registered bool flag, so `flag` never prints its own usage. Usage goes to stdout only on help; every other usage error goes to stderr. Flag-package diagnostics are discarded and re-reported through one path.
- `--top` accepts digits only, so `+5` is a usage error. A value too large for `int` is treated as unbounded.
- `Run` holds two small interfaces, `Profiler` and `FrequencyLister`, so later stories plug in without touching parsing.
- Output is buffered and written only after validation. Unknown column names are exit 3 with the check after load, as specified.
- I replaced the block-1 skeleton test `TestRunExitCodes` "one file" case, which pointed at a nonexistent `a.csv` and now correctly exits 1. The new version uses a real temp file.

**Files changed:**
- `internal/cli/cli.go` (rewritten), `internal/cli/cli_test.go` (rewritten)
- `internal/cli/request.go`, `internal/cli/modes.go`, `internal/cli/parse_test.go` (new)
- `internal/csvio/csvio.go`, `internal/csvio/csvio_test.go` (new)
- `internal/decimal/decimal.go`, `internal/decimal/decimal_test.go` (new)
- `internal/profile/profile.go`, `types.go`, `freq.go`, `profile_test.go` (new)

RESULT: SUCCESS

FILES CHANGED:
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/cli/request.go
- internal/cli/modes.go
- internal/cli/parse_test.go
- internal/csvio/csvio.go
- internal/csvio/csvio_test.go
- internal/decimal/decimal.go
- internal/decimal/decimal_test.go
- internal/profile/profile.go
- internal/profile/types.go
- internal/profile/freq.go
- internal/profile/profile_test.go

SUMMARY:
Implemented the full csvstat CLI for block 2 in Go, standard library only:
- Argument parsing for the profile and `--freq` forms, `--column`, `--top`, and `--help`, with exit mapping 0/1/2/3.
- CSV loading with trimming, rejection of empty files, bad headers, ragged records, and syntax errors (exit 3), and unreadable paths or directories (exit 1).
- Profile rows with rows/nulls/distinct, type inference (integer, number, date, boolean, text, empty), min/max, exact sum/mean/median via `math/big`, and half-away-from-zero two-decimal output.
- `--freq` listings with count-then-byte ordering and `--top` limiting.

Go tests: the required names `TestParseHelp`, `TestParseColumnForm`, `TestParseFreqForm`, `TestTrimHeaderNames`, `TestTrimValues`, `TestRejectEmptyFile`, `TestRejectBadHeader`, and `TestRejectRaggedRecord` are present and pass. Coverage also includes the boundary and error paths in each package. `go test -race ./...` passed.

BLOCKERS:
- None blocking. The declared Python acceptance checks were not executed because the `drydock_go` helper is absent here.

<blueprint-decisions>
[{"spec":"FEATURE-FrequencyTop.md","severity":"Low","subject":"--top accepts digits only","decision":"Options were to accept a leading '+' or to accept digits only. I accept digits only, so '+5' is a usage error (exit 2). A value too large for int is treated as unbounded. Is that acceptable, or should '+5' be accepted?"}]
</blueprint-decisions>
