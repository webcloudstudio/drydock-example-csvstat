# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-10-24
- resulting state: closed/verified
- story points (combined assembled cost): 37629
- execution id: 20261024.211617.637Z-3c41261c

## Stories built
- Prove the application root builds csvstat with go build using only the standard library. (module-build) [story]
- Read a UTF-8 RFC 4180 file with encoding/csv into a header and data records. (csv-reading) [story]
- Trim header names and values and classify null values. (null-trim) [story]
- Exit 1 when the file cannot be opened or read. (unreadable-file) [story]
- Count rows, nulls, and distinct non-null values per column. (row-counts) [story]
- Reject empty files, empty or repeated header names, and ragged records with exit 3. (reject-malformed) [story]
- Parse --column, --freq, --top, and the single file argument. (cli-options) [story]
- Profile only the column named by --column, rejecting an absent name with exit 3. (column-select) [story]
- Reject usage errors with exit 2 and both synopses on standard error. (usage-errors) [story]
- Infer integer, number, boolean, text, and empty column types. (type-inference) [story]
- Recognize date columns with real calendar validation. (date-type) [story]
- List the most frequent non-null values of a column with --freq. (frequencies) [story]
- Compute min and max per column type. (min-max) [story]
- Compute exact sum, mean, and median for integer and number columns. (numeric-stats) [story]
- Write profile and frequency output as CSV with newline line endings. (csv-output) [story]
- Provide bin/test.sh running go vet and go test from the application root. (test-script) [story]

## Acceptance tooling authorization
- FEATURE-ModuleBuild.md#modulebuild-go-build-succeeds: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ModuleBuild.md#modulebuild-binary-runs: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvReading.md#csvreading-header-and-rows: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvReading.md#csvreading-quoted-comma: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvReading.md#csvreading-invalid-utf8-accepted: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvReading.md#csvreading-ragged-returned: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvReading.md#csvreading-line-numbers: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvReading.md#csvreading-empty-input: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NullTrim.md#nulltrim-trim-spaces: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NullTrim.md#nulltrim-null-spellings: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NullTrim.md#nulltrim-trim-before-null: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NullTrim.md#nulltrim-trim-table: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UnreadableFile.md#unreadable-missing-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UnreadableFile.md#unreadable-directory-path: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UnreadableFile.md#unreadable-not-confused-with-rejected: executable=go; scope=test; authorization=existing Target environment
- FEATURE-RejectMalformed.md#reject-ragged-fixture: executable=go; scope=test; authorization=existing Target environment
- FEATURE-RejectMalformed.md#reject-empty-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-RejectMalformed.md#reject-empty-header-name: executable=go; scope=test; authorization=existing Target environment
- FEATURE-RejectMalformed.md#reject-repeated-header-name: executable=go; scope=test; authorization=existing Target environment
- FEATURE-RejectMalformed.md#reject-short-and-long-records: executable=go; scope=test; authorization=existing Target environment
- FEATURE-RejectMalformed.md#accept-header-only: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliOptions.md#options-no-option-profiles-all: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliOptions.md#options-column-profiles-one: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliOptions.md#options-freq-default-top: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliOptions.md#options-freq-explicit-top: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CliOptions.md#options-freq-top-larger-than-distinct: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnSelect.md#column-select-one-row: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnSelect.md#column-select-matches-full-profile-row: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnSelect.md#column-select-absent-name: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnSelect.md#column-select-case-sensitive: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnSelect.md#column-select-trimmed-header: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-missing-or-extra-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-unknown-option: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-missing-option-value: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-invalid-top: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-top-without-freq: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-column-with-freq: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-decided-before-file-open: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Frequencies.md#freq-order-and-default-top: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Frequencies.md#freq-top-limit: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Frequencies.md#freq-trim-and-nulls: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Frequencies.md#freq-all-null-column: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Frequencies.md#freq-absent-column: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Frequencies.md#freq-case-sensitive-name: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvOutput.md#output-profile-shape: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvOutput.md#output-newline-endings: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvOutput.md#output-profile-quoting: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvOutput.md#output-freq-quoting: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvOutput.md#output-column-single-row: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvOutput.md#output-failure-silent-stdout: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TestScript.md#testscript-exists-and-passes: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TestScript.md#testscript-fails-on-failing-test: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TestScript.md#testscript-fails-on-vet-finding: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TestScript.md#testscript-passes-on-clean-module: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 963)
- implements: FEATURE-ModuleBuild.md (SP 1050)
- context: ARCHITECTURE_compact.md (SP 401)
- implements: FEATURE-CsvReading.md (SP 1829)
- implements: FEATURE-NullTrim.md (SP 1286)
- implements: FEATURE-UnreadableFile.md (SP 945)
- implements: FEATURE-Counts.md (SP 1808)
- implements: FEATURE-RejectMalformed.md (SP 1904)
- implements: FEATURE-CliOptions.md (SP 2024)
- implements: FEATURE-ColumnSelect.md (SP 1755)
- implements: FEATURE-UsageErrors.md (SP 1879)
- implements: FEATURE-TypeInference.md (SP 2180)
- implements: FEATURE-DateType.md (SP 1706)
- implements: FEATURE-Frequencies.md (SP 2363)
- implements: FEATURE-MinMax.md (SP 2496)
- implements: FEATURE-NumericStats.md (SP 2846)
- implements: FEATURE-CsvOutput.md (SP 2404)
- implements: FEATURE-TestScript.md (SP 1713)
- stack: common.md (SP 2776)

## Build directory changes
- bin/test.sh
- cmd/csvstat/main_test.go
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/freq/freq.go
- internal/freq/freq_test.go
- internal/output/output.go
- internal/output/output_test.go
- internal/profile/profile.go
- internal/profile/profile_test.go
- internal/reader/reader.go
- internal/reader/reader_test.go
- internal/reader/trim.go
- internal/reader/trim_test.go
- internal/reader/validate.go
- internal/reader/validate_test.go
- internal/typing/typing.go
- internal/typing/typing_test.go

## Pre-build acceptance observation
- GREEN (prepassed): modulebuild-go-build-succeeds (FEATURE-ModuleBuild.md)
  intent: go build of ./cmd/csvstat from the application root succeeds and writes an executable file.
  return code: 0
- GREEN (prepassed): modulebuild-standard-library-only (FEATURE-ModuleBuild.md)
  intent: The module declares no requirements or replacements and has no checksum file.
  return code: 0
- GREEN (prepassed): modulebuild-binary-runs (FEATURE-ModuleBuild.md)
  intent: The built executable starts and ends with a normal process exit status.
  return code: 0
  stdout:
    csvstat: expected exactly one FILE argument
    usage: csvstat FILE
- UNVERIFIED: csvreading-header-and-rows (FEATURE-CsvReading.md)
  intent: The first record is the header and later records are data rows in order.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T16:57:24.289722881-04:00","Action":"start","Package":"csvstat/internal/reader"}
    {"Time":"2026-10-24T16:57:24.289778466-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \tcsvstat/internal/reader\t[no test files]\n"}
    {"Time":"2026-10-24T16:57:24.289783307-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T16:57:24.289722881-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:24.289778466-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:24.289783307-04:00","Action":"skip","Package… (+65 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csvreading-header-and-rows.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T16:57:24.289722881-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:24.289778466-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:24.289783307-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}\n', '')
- UNVERIFIED: csvreading-quoted-comma (FEATURE-CsvReading.md)
  intent: A quoted comma stays inside a single value.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T16:57:24.687640364-04:00","Action":"start","Package":"csvstat/internal/reader"}
    {"Time":"2026-10-24T16:57:24.687688321-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \tcsvstat/internal/reader\t[no test files]\n"}
    {"Time":"2026-10-24T16:57:24.687693939-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T16:57:24.687640364-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:24.687688321-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:24.687693939-04:00","Action":"skip","Package… (+65 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csvreading-quoted-comma.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T16:57:24.687640364-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:24.687688321-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:24.687693939-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}\n', '')
- UNVERIFIED: csvreading-invalid-utf8-accepted (FEATURE-CsvReading.md)
  intent: Invalid UTF-8 bytes do not cause rejection.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T16:57:25.093111389-04:00","Action":"start","Package":"csvstat/internal/reader"}
    {"Time":"2026-10-24T16:57:25.093156041-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \tcsvstat/internal/reader\t[no test files]\n"}
    {"Time":"2026-10-24T16:57:25.093159603-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T16:57:25.093111389-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:25.093156041-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:25.093159603-04:00","Action":"skip","Package… (+65 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csvreading-invalid-utf8-accepted.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T16:57:25.093111389-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:25.093156041-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:25.093159603-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}\n', '')
- UNVERIFIED: csvreading-ragged-returned (FEATURE-CsvReading.md)
  intent: A record with a different field count is returned for later rejection rather than failing the read.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T16:57:25.508759421-04:00","Action":"start","Package":"csvstat/internal/reader"}
    {"Time":"2026-10-24T16:57:25.508891819-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \tcsvstat/internal/reader\t[no test files]\n"}
    {"Time":"2026-10-24T16:57:25.508913275-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T16:57:25.508759421-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:25.508891819-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:25.508913275-04:00","Action":"skip","Package… (+65 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csvreading-ragged-returned.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T16:57:25.508759421-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:25.508891819-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:25.508913275-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}\n', '')
- UNVERIFIED: csvreading-line-numbers (FEATURE-CsvReading.md)
  intent: Each row carries the input line on which it starts.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T16:57:25.837336104-04:00","Action":"start","Package":"csvstat/internal/reader"}
    {"Time":"2026-10-24T16:57:25.83737622-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \tcsvstat/internal/reader\t[no test files]\n"}
    {"Time":"2026-10-24T16:57:25.83737972-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T16:57:25.837336104-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:25.83737622-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:25.83737972-04:00","Action":"skip","Package":… (+63 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csvreading-line-numbers.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T16:57:25.837336104-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:25.83737622-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:25.83737972-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}\n', '')
- UNVERIFIED: csvreading-empty-input (FEATURE-CsvReading.md)
  intent: Empty input yields an empty table and no error.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T16:57:26.209837388-04:00","Action":"start","Package":"csvstat/internal/reader"}
    {"Time":"2026-10-24T16:57:26.209879999-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \tcsvstat/internal/reader\t[no test files]\n"}
    {"Time":"2026-10-24T16:57:26.209883338-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T16:57:26.209837388-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:26.209879999-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:26.209883338-04:00","Action":"skip","Package… (+65 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csvreading-empty-input.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T16:57:26.209837388-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:26.209879999-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:26.209883338-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}\n', '')
- UNVERIFIED: nulltrim-trim-spaces (FEATURE-NullTrim.md)
  intent: Trim removes surrounding spaces only.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T16:57:26.614305646-04:00","Action":"start","Package":"csvstat/internal/reader"}
    {"Time":"2026-10-24T16:57:26.61435336-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \tcsvstat/internal/reader\t[no test files]\n"}
    {"Time":"2026-10-24T16:57:26.614357705-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T16:57:26.614305646-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:26.61435336-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:26.614357705-04:00","Action":"skip","Package"… (+64 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "nulltrim-trim-spaces.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T16:57:26.614305646-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:26.61435336-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:26.614357705-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}\n', '')
- UNVERIFIED: nulltrim-null-spellings (FEATURE-NullTrim.md)
  intent: Empty, NA, N/A, and NULL in any case are null and nothing else is.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T16:57:26.902438551-04:00","Action":"start","Package":"csvstat/internal/reader"}
    {"Time":"2026-10-24T16:57:26.90248168-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \tcsvstat/internal/reader\t[no test files]\n"}
    {"Time":"2026-10-24T16:57:26.902485306-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T16:57:26.902438551-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:26.90248168-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:26.902485306-04:00","Action":"skip","Package"… (+64 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "nulltrim-null-spellings.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T16:57:26.902438551-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:26.90248168-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:26.902485306-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}\n', '')
- UNVERIFIED: nulltrim-trim-before-null (FEATURE-NullTrim.md)
  intent: A value made only of spaces is null once trimmed.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T16:57:27.269775447-04:00","Action":"start","Package":"csvstat/internal/reader"}
    {"Time":"2026-10-24T16:57:27.269820802-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \tcsvstat/internal/reader\t[no test files]\n"}
    {"Time":"2026-10-24T16:57:27.269825012-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T16:57:27.269775447-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:27.269820802-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:27.269825012-04:00","Action":"skip","Package… (+65 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "nulltrim-trim-before-null.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T16:57:27.269775447-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:27.269820802-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:27.269825012-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}\n', '')
- UNVERIFIED: nulltrim-trim-table (FEATURE-NullTrim.md)
  intent: TrimTable trims header names and all row values and keeps line numbers.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T16:57:27.670566464-04:00","Action":"start","Package":"csvstat/internal/reader"}
    {"Time":"2026-10-24T16:57:27.670708132-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \tcsvstat/internal/reader\t[no test files]\n"}
    {"Time":"2026-10-24T16:57:27.670723641-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T16:57:27.670566464-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:27.670708132-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:27.670723641-04:00","Action":"skip","Package… (+65 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "nulltrim-trim-table.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T16:57:27.670566464-04:00","Action":"start","Package":"csvstat/internal/reader"}\n{"Time":"2026-10-24T16:57:27.670708132-04:00","Action":"output","Package":"csvstat/internal/reader","Output":"?   \\tcsvstat/internal/reader\\t[no test files]\\n"}\n{"Time":"2026-10-24T16:57:27.670723641-04:00","Action":"skip","Package":"csvstat/internal/reader","Elapsed":0}\n', '')
- RED: unreadable-missing-file (FEATURE-UnreadableFile.md)
  intent: A path that does not exist exits 1 with nothing on standard output.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 0
      out = ''
      err = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "unreadable-missing-file.py", line 13, in <module>
        assert code == 1, (code, out, err)
               ^^^^^^^^^
    AssertionError: (0, '', '')
- RED: unreadable-directory-path (FEATURE-UnreadableFile.md)
  intent: A path that names a directory cannot be read as a file and exits 1 with nothing on standard output.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 0
      out = ''
      err = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "unreadable-directory-path.py", line 10, in <module>
        assert code == 1, (code, out, err)
               ^^^^^^^^^
    AssertionError: (0, '', '')
- RED: unreadable-not-confused-with-rejected (FEATURE-UnreadableFile.md)
  intent: A readable but malformed file exits 3, not 1, so the two failure classes stay distinct.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 0
      out = ''
      err = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "unreadable-not-confused-with-rejected.py", line 7, in <module>
        assert code == 3, (code, out, err)
               ^^^^^^^^^
    AssertionError: (0, '', '')
- RED: counts-rows-nulls-distinct (FEATURE-Counts.md)
  intent: rows, nulls, and distinct match the counts derived from the supplied column values.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "counts-rows-nulls-distinct.py", line 23, in <module>
        assert len(rows) == 1
               ^^^^^^^^^^^^^^
    AssertionError
- RED: counts-trimmed-distinct (FEATURE-Counts.md)
  intent: Padded and unpadded spellings of one value count once, while differing case counts separately.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      expected_distinct = 3
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "counts-trimmed-distinct.py", line 20, in <module>
        assert int(rows[0]["distinct"]) == expected_distinct
                   ~~~~^^^
    IndexError: list index out of range
- RED: counts-all-null-column (FEATURE-Counts.md)
  intent: A column of only null spellings has nulls equal to rows and distinct zero.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      values = ['', 'NA', 'N/A', 'NULL', 'null', ' ']
      rows = {}
    --- drydock: end values ---
    Traceback (most recent call last):
      File "counts-all-null-column.py", line 19, in <module>
        assert int(rows["a"]["rows"]) == len(values)
                   ~~~~^^^^^
    KeyError: 'a'
- RED: counts-header-only (FEATURE-Counts.md)
  intent: A header-only file reports zero rows, zero nulls, and zero distinct for every column.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      names = ['sku', 'price', 'in_stock']
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "counts-header-only.py", line 19, in <module>
        assert [r["column"] for r in rows] == names
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: reject-ragged-fixture (FEATURE-RejectMalformed.md)
  intent: The supplied ragged fixture is rejected with exit 3 and nothing on standard output.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 0
      out = ''
      err = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-ragged-fixture.py", line 6, in <module>
        assert code == 3, (code, out, err)
               ^^^^^^^^^
    AssertionError: (0, '', '')
- RED: reject-empty-file (FEATURE-RejectMalformed.md)
  intent: A zero-byte file is rejected with exit 3 and nothing on standard output.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 0
      out = ''
      err = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-empty-file.py", line 14, in <module>
        assert code == 3, (code, out, err)
               ^^^^^^^^^
    AssertionError: (0, '', '')
- RED: reject-empty-header-name (FEATURE-RejectMalformed.md)
  intent: A header with an empty or blank name is rejected with exit 3.
  return code: 1
  stdout:
    blank 0
  stderr:
    --- drydock: values at failure ---
      code = 0
      label = 'blank'
      out = ''
      err = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-empty-header-name.py", line 13, in <module>
        assert code == 3, (label, code, out, err)
               ^^^^^^^^^
    AssertionError: ('blank', 0, '', '')
- RED: reject-repeated-header-name (FEATURE-RejectMalformed.md)
  intent: A repeated header name, including one that repeats only after trimming, is rejected with exit 3.
  return code: 1
  stdout:
    plain 0
  stderr:
    --- drydock: values at failure ---
      code = 0
      label = 'plain'
      out = ''
      err = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-repeated-header-name.py", line 13, in <module>
        assert code == 3, (label, code, out, err)
               ^^^^^^^^^
    AssertionError: ('plain', 0, '', '')
- RED: reject-short-and-long-records (FEATURE-RejectMalformed.md)
  intent: A record with too few fields and a record with too many fields are each rejected with exit 3.
  return code: 1
  stdout:
    short 0
  stderr:
    --- drydock: values at failure ---
      code = 0
      label = 'short'
      out = ''
      err = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-short-and-long-records.py", line 17, in <module>
        assert code == 3, (label, code, out, err)
               ^^^^^^^^^
    AssertionError: ('short', 0, '', '')
- GREEN (prepassed): accept-header-only (FEATURE-RejectMalformed.md)
  intent: A header-only file is valid and is not rejected.
  return code: 0
- RED: options-no-option-profiles-all (FEATURE-CliOptions.md)
  intent: With only a file argument, every column is profiled in file order.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
      names = ['alpha', 'beta', 'gamma']
    --- drydock: end values ---
    Traceback (most recent call last):
      File "options-no-option-profiles-all.py", line 19, in <module>
        assert [row[0] for row in rows[1:]] == names, rows
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError: []
- RED: options-column-profiles-one (FEATURE-CliOptions.md)
  intent: --column profiles exactly the named column.
  return code: 1
  stdout:
    flag provided but not defined: -column
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'flag provided but not defined: -column\nusage: csvstat FILE\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "options-column-profiles-one.py", line 18, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'flag provided but not defined: -column\nusage: csvstat FILE\n')
- RED: options-freq-default-top (FEATURE-CliOptions.md)
  intent: --freq without --top lists at most five values.
  return code: 1
  stdout:
    flag provided but not defined: -freq
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'flag provided but not defined: -freq\nusage: csvstat FILE\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "options-freq-default-top.py", line 20, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'flag provided but not defined: -freq\nusage: csvstat FILE\n')
- RED: options-freq-explicit-top (FEATURE-CliOptions.md)
  intent: --top limits the frequency list to the requested number of rows.
  return code: 1
  stdout:
    flag provided but not defined: -freq
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'flag provided but not defined: -freq\nusage: csvstat FILE\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "options-freq-explicit-top.py", line 20, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'flag provided but not defined: -freq\nusage: csvstat FILE\n')
- RED: options-freq-top-larger-than-distinct (FEATURE-CliOptions.md)
  intent: A --top larger than the number of distinct values lists every distinct value.
  return code: 1
  stdout:
    flag provided but not defined: -freq
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'flag provided but not defined: -freq\nusage: csvstat FILE\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "options-freq-top-larger-than-distinct.py", line 20, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'flag provided but not defined: -freq\nusage: csvstat FILE\n')
- RED: column-select-one-row (FEATURE-ColumnSelect.md)
  intent: --column outputs the profile header and a single row for the named column.
  return code: 1
  stdout:
    flag provided but not defined: -column
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'flag provided but not defined: -column\nusage: csvstat FILE\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "column-select-one-row.py", line 17, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'flag provided but not defined: -column\nusage: csvstat FILE\n')
- RED: column-select-matches-full-profile-row (FEATURE-ColumnSelect.md)
  intent: The row produced by --column equals that column's row in the full profile.
  return code: 1
  stdout:
    flag provided but not defined: -column
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code_one = 2
      err_one = 'flag provided but not defined: -column\nusage: csvstat FILE\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "column-select-matches-full-profile-row.py", line 17, in <module>
        assert code_one == 0, (code_one, err_one)
               ^^^^^^^^^^^^^
    AssertionError: (2, 'flag provided but not defined: -column\nusage: csvstat FILE\n')
- RED: column-select-absent-name (FEATURE-ColumnSelect.md)
  intent: A column name not in the header exits 3 with nothing on standard output.
  return code: 1
  stdout:
    flag provided but not defined: -column
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'flag provided but not defined: -column\nusage: csvstat FILE\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "column-select-absent-name.py", line 14, in <module>
        assert code == 3, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'flag provided but not defined: -column\nusage: csvstat FILE\n')
- RED: column-select-case-sensitive (FEATURE-ColumnSelect.md)
  intent: Column names match case-sensitively, so a differently cased name is absent and exits 3.
  return code: 1
  stdout:
    flag provided but not defined: -column
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'flag provided but not defined: -column\nusage: csvstat FILE\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "column-select-case-sensitive.py", line 15, in <module>
        assert code == 3, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'flag provided but not defined: -column\nusage: csvstat FILE\n')
- RED: column-select-trimmed-header (FEATURE-ColumnSelect.md)
  intent: A header name padded with spaces in the file is matched by its trimmed name.
  return code: 1
  stdout:
    flag provided but not defined: -column
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
      out = ''
      err = 'flag provided but not defined: -column\nusage: csvstat FILE\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "column-select-trimmed-header.py", line 17, in <module>
        assert code == 0, (code, out, err)
               ^^^^^^^^^
    AssertionError: (2, '', 'flag provided but not defined: -column\nusage: csvstat FILE\n')
- GREEN (prepassed): usage-missing-or-extra-file (FEATURE-UsageErrors.md)
  intent: No file argument and two file arguments each exit 2 with nothing on standard output.
  return code: 0
  stdout:
    none 2  csvstat: expected exactly one FILE argument
    usage: csvstat FILE
    
    two 2  csvstat: expected exactly one FILE argument
    usage: csvstat FILE
- GREEN (prepassed): usage-unknown-option (FEATURE-UsageErrors.md)
  intent: An unknown option, including --help and -h, exits 2 with nothing on standard output.
  return code: 0
  stdout:
    --frobnicate 2  flag provided but not defined: -frobnicate
    usage: csvstat FILE
    
    --help 2  usage: csvstat FILE
    
    -h 2  usage: csvstat FILE
- GREEN (prepassed): usage-missing-option-value (FEATURE-UsageErrors.md)
  intent: Options that need a value but have none exit 2 with nothing on standard output.
  return code: 0
  stdout:
    column 2  flag provided but not defined: -column
    usage: csvstat FILE
    
    freq 2  flag provided but not defined: -freq
    usage: csvstat FILE
    
    top 2  flag provided but not defined: -freq
    usage: csvstat FILE
- GREEN (prepassed): usage-invalid-top (FEATURE-UsageErrors.md)
  intent: A --top that is zero, negative, or not a number exits 2 with nothing on standard output.
  return code: 0
  stdout:
    0 2  flag provided but not defined: -freq
    usage: csvstat FILE
    
    -1 2  flag provided but not defined: -freq
    usage: csvstat FILE
    
    abc 2  flag provided but not defined: -freq
    usage: csvstat FILE
    
    1.5 2  flag provided but not defined: -freq
    usage: csvstat FILE
- GREEN (prepassed): usage-top-without-freq (FEATURE-UsageErrors.md)
  intent: --top without --freq exits 2 with nothing on standard output.
  return code: 0
  stdout:
    flag provided but not defined: -top
    usage: csvstat FILE
- GREEN (prepassed): usage-column-with-freq (FEATURE-UsageErrors.md)
  intent: --column combined with --freq exits 2 with nothing on standard output.
  return code: 0
  stdout:
    flag provided but not defined: -column
    usage: csvstat FILE
- GREEN (prepassed): usage-decided-before-file-open (FEATURE-UsageErrors.md)
  intent: A usage error on a nonexistent file still exits 2, not 1, because usage is checked before the file is opened.
  return code: 0
  stdout:
    flag provided but not defined: -top
    usage: csvstat FILE
- RED: type-integer-and-number (FEATURE-TypeInference.md)
  intent: Integer, mixed integer-and-decimal, and negative columns are typed integer and number as the contract states.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = {}
    --- drydock: end values ---
    Traceback (most recent call last):
      File "type-integer-and-number.py", line 28, in <module>
        assert rows["ints"]["type"] == "integer"
               ~~~~^^^^^^^^
    KeyError: 'ints'
- RED: type-boolean-case-insensitive (FEATURE-TypeInference.md)
  intent: A column of true and false in any letter case is boolean, while a column mixing a boolean with other text is text.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = {}
    --- drydock: end values ---
    Traceback (most recent call last):
      File "type-boolean-case-insensitive.py", line 27, in <module>
        assert rows["flag"]["type"] == "boolean"
               ~~~~^^^^^^^^
    KeyError: 'flag'
- RED: type-text-fallback (FEATURE-TypeInference.md)
  intent: A column with one non-numeric value among integers falls back to text.
  return code: 1
  note: proof contains no assertion, raise, or failing exit
  stderr:
    Traceback (most recent call last):
      File "type-text-fallback.py", line 6, in <module>
        from drydb_placeholder import nothing  # noqa: F401
        ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    ModuleNotFoundError: No module named 'drydb_placeholder'
- RED: type-nulls-ignored (FEATURE-TypeInference.md)
  intent: Null spellings (empty, NA, N/A, NULL in any case) do not change the inferred type, and an all-null column is empty.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = {}
    --- drydock: end values ---
    Traceback (most recent call last):
      File "type-nulls-ignored.py", line 28, in <module>
        assert rows["n"]["type"] == "integer"
               ~~~~^^^^^
    KeyError: 'n'
- RED: type-header-only-empty (FEATURE-TypeInference.md)
  intent: A file with a header and no data records types every column as empty.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      names = ['a', 'b', 'c']
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "type-header-only-empty.py", line 19, in <module>
        assert [r["column"] for r in rows] == names
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: date-valid-column (FEATURE-DateType.md)
  intent: A column of real calendar dates, including a leap day, is typed date.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-valid-column.py", line 19, in <module>
        assert len(rows) == 1
               ^^^^^^^^^^^^^^
    AssertionError
- RED: date-invalid-calendar-day (FEATURE-DateType.md)
  intent: A single impossible calendar date makes the whole column text.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-invalid-calendar-day.py", line 20, in <module>
        assert rows[0]["type"] == "text"
               ~~~~^^^
    IndexError: list index out of range
- RED: date-non-leap-and-month-bounds (FEATURE-DateType.md)
  intent: Feb 29 in a non-leap year, month 13, and day 00 each fail date typing.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      n = 'nonleap'
      rows = {}
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-non-leap-and-month-bounds.py", line 26, in <module>
        assert rows[n]["type"] == "text", n
               ~~~~^^^
    KeyError: 'nonleap'
- RED: date-nulls-and-timestamps (FEATURE-DateType.md)
  intent: Nulls do not disturb a date column, while timestamp values are not dates.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = {}
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-nulls-and-timestamps.py", line 27, in <module>
        assert rows["d"]["type"] == "date"
               ~~~~^^^^^
    KeyError: 'd'
- RED: freq-order-and-default-top (FEATURE-Frequencies.md)
  intent: Frequencies are ordered by count descending then byte value ascending, and default to five rows.
  return code: 1
  stdout:
    flag provided but not defined: -freq
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-order-and-default-top.py", line 20, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: freq-top-limit (FEATURE-Frequencies.md)
  intent: --top limits the number of rows listed.
  return code: 1
  stdout:
    flag provided but not defined: -freq
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-top-limit.py", line 21, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: freq-trim-and-nulls (FEATURE-Frequencies.md)
  intent: Values are trimmed before counting and null spellings are not listed.
  return code: 1
  stdout:
    flag provided but not defined: -freq
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-trim-and-nulls.py", line 20, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: freq-all-null-column (FEATURE-Frequencies.md)
  intent: A column with no non-null values writes only the header.
  return code: 1
  stdout:
    flag provided but not defined: -freq
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-all-null-column.py", line 19, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: freq-absent-column (FEATURE-Frequencies.md)
  intent: A column name absent from the header exits 3 and writes nothing to standard output.
  return code: 1
  stdout:
    flag provided but not defined: -freq
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-absent-column.py", line 15, in <module>
        assert code == 3
               ^^^^^^^^^
    AssertionError
- RED: freq-case-sensitive-name (FEATURE-Frequencies.md)
  intent: The --freq column name is matched case-sensitively against the header.
  return code: 1
  stdout:
    flag provided but not defined: -freq
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-case-sensitive-name.py", line 16, in <module>
        assert code == 3
               ^^^^^^^^^
    AssertionError
- RED: minmax-integer (FEATURE-MinMax.md)
  intent: Integer min and max, including negatives and a value beyond 64 bits, are exact plain integers.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      out = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "minmax-integer.py", line 18, in <module>
        r = list(csv.DictReader(io.StringIO(out)))[0]
            ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~^^^
    IndexError: list index out of range
- RED: minmax-number-two-places (FEATURE-MinMax.md)
  intent: Number min and max are written with two decimal places rounded half away from zero.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      out = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "minmax-number-two-places.py", line 23, in <module>
        r = list(csv.DictReader(io.StringIO(out)))[0]
            ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~^^^
    IndexError: list index out of range
- RED: minmax-number-negative (FEATURE-MinMax.md)
  intent: Negative decimals order numerically, not lexically, and keep their sign.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      out = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "minmax-number-negative.py", line 23, in <module>
        r = list(csv.DictReader(io.StringIO(out)))[0]
            ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~^^^
    IndexError: list index out of range
- RED: minmax-date-boolean-text (FEATURE-MinMax.md)
  intent: Date, boolean, and text columns report earliest/latest, any-false/any-true, and byte-order extremes.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = {}
    --- drydock: end values ---
    Traceback (most recent call last):
      File "minmax-date-boolean-text.py", line 29, in <module>
        assert rows["when"]["type"] == "date"
               ~~~~^^^^^^^^
    KeyError: 'when'
- RED: minmax-empty-and-nulls (FEATURE-MinMax.md)
  intent: An empty column has blank min and max, and nulls are excluded from extremes.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = {}
    --- drydock: end values ---
    Traceback (most recent call last):
      File "minmax-empty-and-nulls.py", line 28, in <module>
        assert rows["none"]["type"] == "empty"
               ~~~~^^^^^^^^
    KeyError: 'none'
- RED: numeric-integer-stats (FEATURE-NumericStats.md)
  intent: Integer sum is a plain integer while mean and median carry two places, with an even count averaging the middle pair.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      out = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numeric-integer-stats.py", line 32, in <module>
        r = list(csv.DictReader(io.StringIO(out)))[0]
            ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~^^^
    IndexError: list index out of range
- RED: numeric-integer-odd-median (FEATURE-NumericStats.md)
  intent: An odd count takes the middle sorted value, independent of input order.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      out = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numeric-integer-odd-median.py", line 29, in <module>
        r = list(csv.DictReader(io.StringIO(out)))[0]
            ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~^^^
    IndexError: list index out of range
- RED: numeric-exact-large-sum (FEATURE-NumericStats.md)
  intent: Integer totals beyond 64 bits are exact rather than overflowed.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      out = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numeric-exact-large-sum.py", line 18, in <module>
        r = list(csv.DictReader(io.StringIO(out)))[0]
            ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~^^^
    IndexError: list index out of range
- RED: numeric-number-stats-rounding (FEATURE-NumericStats.md)
  intent: Number sum, mean, and median are two places rounded half away from zero from the exact value, with nulls excluded.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      out = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numeric-number-stats-rounding.py", line 33, in <module>
        r = list(csv.DictReader(io.StringIO(out)))[0]
            ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~^^^
    IndexError: list index out of range
- RED: numeric-half-away-negative (FEATURE-NumericStats.md)
  intent: A negative exact tie rounds away from zero, so the magnitude rounds up.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      out = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numeric-half-away-negative.py", line 29, in <module>
        r = list(csv.DictReader(io.StringIO(out)))[0]
            ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~^^^
    IndexError: list index out of range
- RED: numeric-blank-for-non-numeric (FEATURE-NumericStats.md)
  intent: Text, boolean, and empty columns leave sum, mean, and median blank.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      n = 'word'
      rows = {}
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numeric-blank-for-non-numeric.py", line 29, in <module>
        assert rows[n]["sum"] == "", n
               ~~~~^^^
    KeyError: 'word'
- RED: output-profile-shape (FEATURE-CsvOutput.md)
  intent: Profile output has the declared header and one row per column in file order.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "output-profile-shape.py", line 21, in <module>
        assert rows[0] == ["column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"]
               ~~~~^^^
    IndexError: list index out of range
- RED: output-newline-endings (FEATURE-CsvOutput.md)
  intent: Output lines end with a line feed and never a carriage return.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code2 = 2
    --- drydock: end values ---
    Traceback (most recent call last):
      File "output-newline-endings.py", line 17, in <module>
        assert code2 == 0
               ^^^^^^^^^^
    AssertionError
- RED: output-profile-quoting (FEATURE-CsvOutput.md)
  intent: A profile field containing a comma or a double quote round-trips through RFC 4180 quoting.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "output-profile-quoting.py", line 22, in <module>
        header = rows[0]
                 ~~~~^^^
    IndexError: list index out of range
- RED: output-freq-quoting (FEATURE-CsvOutput.md)
  intent: A frequency value containing a comma round-trips through RFC 4180 quoting.
  return code: 1
  stdout:
    flag provided but not defined: -freq
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
    --- drydock: end values ---
    Traceback (most recent call last):
      File "output-freq-quoting.py", line 19, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: output-column-single-row (FEATURE-CsvOutput.md)
  intent: --column yields the header and exactly the named column's row.
  return code: 1
  stdout:
    flag provided but not defined: -column
    usage: csvstat FILE
  stderr:
    --- drydock: values at failure ---
      code = 2
    --- drydock: end values ---
    Traceback (most recent call last):
      File "output-column-single-row.py", line 18, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: output-failure-silent-stdout (FEATURE-CsvOutput.md)
  intent: A rejected file writes nothing to standard output.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 0
    --- drydock: end values ---
    Traceback (most recent call last):
      File "output-failure-silent-stdout.py", line 17, in <module>
        assert code == 3
               ^^^^^^^^^
    AssertionError
- RED: testscript-exists-and-passes (FEATURE-TestScript.md)
  intent: bin/test.sh exists and exits zero on the application root.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      script = PosixPath('bin/test.sh')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "testscript-exists-and-passes.py", line 5, in <module>
        assert script.is_file()
               ^^^^^^^^^^^^^^^^
    AssertionError
- RED: testscript-fails-on-failing-test (FEATURE-TestScript.md)
  intent: bin/test.sh exits non-zero when a test in the module fails.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      src = 'bin/test.sh'
      _HAS_FCOPYFILE = False
      dst = PosixPath('/tmp/tmp_ue8engf/bin/test.sh')
      _USE_CP_SENDFILE = True
      _WINDOWS = False
      file_size = 0
      COPY_BUFSIZE = 65536
    --- drydock: end values ---
    Traceback (most recent call last):
      File "testscript-fails-on-failing-test.py", line 9, in <module>
        shutil.copy("bin/test.sh", root / "bin" / "test.sh")
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/shutil.py", line 435, in copy
        copyfile(src, dst, follow_symlinks=follow_symlinks)
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/shutil.py", line 260, in copyfile
        with open(src, 'rb') as fsrc:
             ^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'bin/test.sh'
- RED: testscript-fails-on-vet-finding (FEATURE-TestScript.md)
  intent: bin/test.sh exits non-zero when go vet reports a finding even though tests pass.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      src = 'bin/test.sh'
      _HAS_FCOPYFILE = False
      dst = PosixPath('/tmp/tmpc1yevdfn/bin/test.sh')
      _USE_CP_SENDFILE = True
      _WINDOWS = False
      file_size = 0
      COPY_BUFSIZE = 65536
    --- drydock: end values ---
    Traceback (most recent call last):
      File "testscript-fails-on-vet-finding.py", line 9, in <module>
        shutil.copy("bin/test.sh", root / "bin" / "test.sh")
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/shutil.py", line 435, in copy
        copyfile(src, dst, follow_symlinks=follow_symlinks)
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/shutil.py", line 260, in copyfile
        with open(src, 'rb') as fsrc:
             ^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'bin/test.sh'
- RED: testscript-passes-on-clean-module (FEATURE-TestScript.md)
  intent: bin/test.sh exits zero for a module whose vet and tests are clean, from a different working directory.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      src = 'bin/test.sh'
      _HAS_FCOPYFILE = False
      dst = PosixPath('/tmp/tmpanlm9086/bin/test.sh')
      _USE_CP_SENDFILE = True
      _WINDOWS = False
      file_size = 0
      COPY_BUFSIZE = 65536
    --- drydock: end values ---
    Traceback (most recent call last):
      File "testscript-passes-on-clean-module.py", line 9, in <module>
        shutil.copy("bin/test.sh", root / "bin" / "test.sh")
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/shutil.py", line 435, in copy
        copyfile(src, dst, follow_symlinks=follow_symlinks)
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/shutil.py", line 260, in copyfile
        with open(src, 'rb') as fsrc:
             ^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'bin/test.sh'

## Post-build programmatic acceptance
- PASS: modulebuild-go-build-succeeds (FEATURE-ModuleBuild.md)
  intent: go build of ./cmd/csvstat from the application root succeeds and writes an executable file.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: modulebuild-standard-library-only (FEATURE-ModuleBuild.md)
  intent: The module declares no requirements or replacements and has no checksum file.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: modulebuild-binary-runs (FEATURE-ModuleBuild.md)
  intent: The built executable starts and ends with a normal process exit status.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: expected exactly one file argument
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
- UNVERIFIED: csvreading-header-and-rows (FEATURE-CsvReading.md)
  intent: The first record is the header and later records are data rows in order.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: csvreading-quoted-comma (FEATURE-CsvReading.md)
  intent: A quoted comma stays inside a single value.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: csvreading-invalid-utf8-accepted (FEATURE-CsvReading.md)
  intent: Invalid UTF-8 bytes do not cause rejection.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: csvreading-ragged-returned (FEATURE-CsvReading.md)
  intent: A record with a different field count is returned for later rejection rather than failing the read.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: csvreading-line-numbers (FEATURE-CsvReading.md)
  intent: Each row carries the input line on which it starts.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: csvreading-empty-input (FEATURE-CsvReading.md)
  intent: Empty input yields an empty table and no error.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: nulltrim-trim-spaces (FEATURE-NullTrim.md)
  intent: Trim removes surrounding spaces only.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: nulltrim-null-spellings (FEATURE-NullTrim.md)
  intent: Empty, NA, N/A, and NULL in any case are null and nothing else is.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: nulltrim-trim-before-null (FEATURE-NullTrim.md)
  intent: A value made only of spaces is null once trimmed.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: nulltrim-trim-table (FEATURE-NullTrim.md)
  intent: TrimTable trims header names and all row values and keeps line numbers.
  provisioning: not required
  error: excluded check: ERROR environment
- PASS: unreadable-missing-file (FEATURE-UnreadableFile.md)
  intent: A path that does not exist exits 1 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: open /tmp/tmpiot9jaip/no-such-file.csv: no such file or directory
- PASS: unreadable-directory-path (FEATURE-UnreadableFile.md)
  intent: A path that names a directory cannot be read as a file and exits 1 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmp7ze8rj3z: read /tmp/tmp7ze8rj3z: is a directory
- PASS: unreadable-not-confused-with-rejected (FEATURE-UnreadableFile.md)
  intent: A readable but malformed file exits 3, not 1, so the two failure classes stay distinct.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: sources/ragged.csv: line 4: record has 4 fields, header has 3
- PASS: counts-rows-nulls-distinct (FEATURE-Counts.md)
  intent: rows, nulls, and distinct match the counts derived from the supplied column values.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    fruit,text,10,4,3,apple,pear,,,
- PASS: counts-trimmed-distinct (FEATURE-Counts.md)
  intent: Padded and unpadded spellings of one value count once, while differing case counts separately.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    customer,text,5,0,3,Acme,globex,,,
- PASS: counts-all-null-column (FEATURE-Counts.md)
  intent: A column of only null spellings has nulls equal to rows and distinct zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    a,empty,6,6,0,,,,,
    b,text,6,0,1,x,x,,,
- PASS: counts-header-only (FEATURE-Counts.md)
  intent: A header-only file reports zero rows, zero nulls, and zero distinct for every column.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    sku,empty,0,0,0,,,,,
    price,empty,0,0,0,,,,,
    in_stock,empty,0,0,0,,,,,
- PASS: reject-ragged-fixture (FEATURE-RejectMalformed.md)
  intent: The supplied ragged fixture is rejected with exit 3 and nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: sources/ragged.csv: line 4: record has 4 fields, header has 3
- PASS: reject-empty-file (FEATURE-RejectMalformed.md)
  intent: A zero-byte file is rejected with exit 3 and nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpe9vm94rh/empty.csv: file is empty
- PASS: reject-empty-header-name (FEATURE-RejectMalformed.md)
  intent: A header with an empty or blank name is rejected with exit 3.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    blank 3  csvstat: /tmp/tmpxozym9_k/blank.csv: header name 2 is empty
    
    empty 3  csvstat: /tmp/tmpxozym9_k/empty.csv: header name 2 is empty
- PASS: reject-repeated-header-name (FEATURE-RejectMalformed.md)
  intent: A repeated header name, including one that repeats only after trimming, is rejected with exit 3.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    plain 3  csvstat: /tmp/tmpzaw55e6q/plain.csv: header name "a" is repeated
    
    padded 3  csvstat: /tmp/tmpzaw55e6q/padded.csv: header name "a" is repeated
- PASS: reject-short-and-long-records (FEATURE-RejectMalformed.md)
  intent: A record with too few fields and a record with too many fields are each rejected with exit 3.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    short 3  csvstat: /tmp/tmpgercdd0z/short.csv: line 3: record has 2 fields, header has 3
    
    long 3  csvstat: /tmp/tmpgercdd0z/long.csv: line 3: record has 4 fields, header has 3
- PASS: accept-header-only (FEATURE-RejectMalformed.md)
  intent: A header-only file is valid and is not rejected.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    sku,empty,0,0,0,,,,,
    price,empty,0,0,0,,,,,
    in_stock,empty,0,0,0,,,,,
- PASS: options-no-option-profiles-all (FEATURE-CliOptions.md)
  intent: With only a file argument, every column is profiled in file order.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    alpha,integer,2,0,2,1,4,5,2.50,2.50
    beta,integer,2,0,2,2,5,7,3.50,3.50
    gamma,integer,2,0,2,3,6,9,4.50,4.50
- PASS: options-column-profiles-one (FEATURE-CliOptions.md)
  intent: --column profiles exactly the named column.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    beta,integer,2,0,2,2,5,7,3.50,3.50
- PASS: options-freq-default-top (FEATURE-CliOptions.md)
  intent: --freq without --top lists at most five values.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    item0,1
    item1,1
    item2,1
    item3,1
    item4,1
- PASS: options-freq-explicit-top (FEATURE-CliOptions.md)
  intent: --top limits the frequency list to the requested number of rows.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    item0,1
    item1,1
- PASS: options-freq-top-larger-than-distinct (FEATURE-CliOptions.md)
  intent: A --top larger than the number of distinct values lists every distinct value.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    item0,1
    item1,1
    item2,1
- PASS: column-select-one-row (FEATURE-ColumnSelect.md)
  intent: --column outputs the profile header and a single row for the named column.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    qty,integer,2,0,2,10,20,30,15.00,15.00
- PASS: column-select-matches-full-profile-row (FEATURE-ColumnSelect.md)
  intent: The row produced by --column equals that column's row in the full profile.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    salary,number,24,0,24,43699.25,147150.99,2085635.10,86901.46,79496.68
- PASS: column-select-absent-name (FEATURE-ColumnSelect.md)
  intent: A column name not in the header exits 3 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpc8ls4iz0/data.csv: no column named "nosuch"
- PASS: column-select-case-sensitive (FEATURE-ColumnSelect.md)
  intent: Column names match case-sensitively, so a differently cased name is absent and exits 3.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpfoxzxg8x/data.csv: no column named "QTY"
- PASS: column-select-trimmed-header (FEATURE-ColumnSelect.md)
  intent: A header name padded with spaces in the file is matched by its trimmed name.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    qty,integer,1,0,1,10,10,10,10.00,10.00
- PASS: usage-missing-or-extra-file (FEATURE-UsageErrors.md)
  intent: No file argument and two file arguments each exit 2 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    none 2  csvstat: expected exactly one file argument
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
    
    two 2  csvstat: expected exactly one file argument
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
- PASS: usage-unknown-option (FEATURE-UsageErrors.md)
  intent: An unknown option, including --help and -h, exits 2 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    --frobnicate 2  csvstat: flag provided but not defined: -frobnicate
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
    
    --help 2  csvstat: no help option is defined
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
    
    -h 2  csvstat: no help option is defined
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
- PASS: usage-missing-option-value (FEATURE-UsageErrors.md)
  intent: Options that need a value but have none exit 2 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column 2  csvstat: flag needs an argument: -column
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
    
    freq 2  csvstat: flag needs an argument: -freq
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
    
    top 2  csvstat: expected exactly one file argument
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
- PASS: usage-invalid-top (FEATURE-UsageErrors.md)
  intent: A --top that is zero, negative, or not a number exits 2 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    0 2  csvstat: --top must be a positive integer
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
    
    -1 2  csvstat: --top must be a positive integer
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
    
    abc 2  csvstat: invalid value "abc" for flag -top: parse error
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
    
    1.5 2  csvstat: invalid value "1.5" for flag -top: parse error
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
- PASS: usage-top-without-freq (FEATURE-UsageErrors.md)
  intent: --top without --freq exits 2 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: --top requires --freq
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
- PASS: usage-column-with-freq (FEATURE-UsageErrors.md)
  intent: --column combined with --freq exits 2 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: --column and --freq cannot be combined
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
- PASS: usage-decided-before-file-open (FEATURE-UsageErrors.md)
  intent: A usage error on a nonexistent file still exits 2, not 1, because usage is checked before the file is opened.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: --top requires --freq
    usage: csvstat [--column <name>] <file>
           csvstat --freq <name> [--top <n>] <file>
- PASS: type-integer-and-number (FEATURE-TypeInference.md)
  intent: Integer, mixed integer-and-decimal, and negative columns are typed integer and number as the contract states.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    ints,integer,4,0,4,-20,300,281,70.25,0.50
    mixed,number,4,0,4,-3.00,4.25,4.75,1.19,1.75
    decs,number,4,0,4,-1.75,10.00,11.89,2.97,1.82
- PASS: type-boolean-case-insensitive (FEATURE-TypeInference.md)
  intent: A column of true and false in any letter case is boolean, while a column mixing a boolean with other text is text.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    flag,boolean,4,0,4,false,true,,,
    maybe,text,4,0,3,false,yes,,,
- UNVERIFIED: type-text-fallback (FEATURE-TypeInference.md)
  intent: A column with one non-numeric value among integers falls back to text.
  provisioning: not required
  error: excluded check: ALWAYS_PASSES vacuous
- PASS: type-nulls-ignored (FEATURE-TypeInference.md)
  intent: Null spellings (empty, NA, N/A, NULL in any case) do not change the inferred type, and an all-null column is empty.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    n,integer,6,4,2,5,7,12,6.00,6.00
    allnull,empty,6,6,0,,,,,
    word,text,6,3,3,blue,red,,,
- PASS: type-header-only-empty (FEATURE-TypeInference.md)
  intent: A file with a header and no data records types every column as empty.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    a,empty,0,0,0,,,,,
    b,empty,0,0,0,,,,,
    c,empty,0,0,0,,,,,
- PASS: date-valid-column (FEATURE-DateType.md)
  intent: A column of real calendar dates, including a leap day, is typed date.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    when,date,4,0,4,2015-08-27,2025-12-01,,,
- PASS: date-invalid-calendar-day (FEATURE-DateType.md)
  intent: A single impossible calendar date makes the whole column text.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    shipped,text,3,0,3,2025-02-30,2025-03-18,,,
- PASS: date-non-leap-and-month-bounds (FEATURE-DateType.md)
  intent: Feb 29 in a non-leap year, month 13, and day 00 each fail date typing.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    nonleap,text,1,0,1,2025-02-29,2025-02-29,,,
    month13,text,1,0,1,2025-13-01,2025-13-01,,,
    day00,text,1,0,1,2025-01-00,2025-01-00,,,
- PASS: date-nulls-and-timestamps (FEATURE-DateType.md)
  intent: Nulls do not disturb a date column, while timestamp values are not dates.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    d,date,4,2,2,2025-03-01,2025-03-02,,,
    ts,text,4,0,4,2025-03-01T08:00,2025-03-03T11:00,,,
- PASS: freq-order-and-default-top (FEATURE-Frequencies.md)
  intent: Frequencies are ordered by count descending then byte value ascending, and default to five rows.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    delta,4
    alpha,3
    charlie,3
    bravo,2
    echo,1
- PASS: freq-top-limit (FEATURE-Frequencies.md)
  intent: --top limits the number of rows listed.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    a,5
    b,4
- PASS: freq-trim-and-nulls (FEATURE-Frequencies.md)
  intent: Values are trimmed before counting and null spellings are not listed.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    Globex,3
    Other,1
- PASS: freq-all-null-column (FEATURE-Frequencies.md)
  intent: A column with no non-null values writes only the header.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
- PASS: freq-absent-column (FEATURE-Frequencies.md)
  intent: A column name absent from the header exits 3 and writes nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpn48516k0/data.csv: no column named "nosuch"
- PASS: freq-case-sensitive-name (FEATURE-Frequencies.md)
  intent: The --freq column name is matched case-sensitively against the header.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmp20mqtwwh/data.csv: no column named "kind"
- PASS: minmax-integer (FEATURE-MinMax.md)
  intent: Integer min and max, including negatives and a value beyond 64 bits, are exact plain integers.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    n,integer,6,0,6,-99,12345678901234567890123,12345678901234567890064,2057613150205761315010.67,2.50
- PASS: minmax-number-two-places (FEATURE-MinMax.md)
  intent: Number min and max are written with two decimal places rounded half away from zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    p,number,5,0,5,2.01,231.18,355.96,71.19,15.27
- PASS: minmax-number-negative (FEATURE-MinMax.md)
  intent: Negative decimals order numerically, not lexically, and keep their sign.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    t,number,5,0,5,-12.21,29.81,14.95,2.99,-0.44
- PASS: minmax-date-boolean-text (FEATURE-MinMax.md)
  intent: Date, boolean, and text columns report earliest/latest, any-false/any-true, and byte-order extremes.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    when,date,4,0,4,2015-08-27,2025-06-02,,,
    allt,boolean,4,0,3,true,true,,,
    mix,boolean,4,0,3,false,true,,,
    word,text,4,0,4,Zed,remote,,,
- PASS: minmax-empty-and-nulls (FEATURE-MinMax.md)
  intent: An empty column has blank min and max, and nulls are excluded from extremes.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    none,empty,4,4,0,,,,,
    n,integer,4,2,2,3,9,12,6.00,6.00
- PASS: numeric-integer-stats (FEATURE-NumericStats.md)
  intent: Integer sum is a plain integer while mean and median carry two places, with an even count averaging the middle pair.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    n,integer,4,0,4,1,10,16,4.00,2.50
- PASS: numeric-integer-odd-median (FEATURE-NumericStats.md)
  intent: An odd count takes the middle sorted value, independent of input order.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    n,integer,5,0,5,1,100,118,23.60,5.00
- PASS: numeric-exact-large-sum (FEATURE-NumericStats.md)
  intent: Integer totals beyond 64 bits are exact rather than overflowed.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    n,integer,3,0,1,9223372036854775807,9223372036854775807,27670116110564327421,9223372036854775807.00,9223372036854775807.00
- PASS: numeric-number-stats-rounding (FEATURE-NumericStats.md)
  intent: Number sum, mean, and median are two places rounded half away from zero from the exact value, with nulls excluded.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    t,number,8,2,6,-12.21,29.33,40.64,6.77,1.27
- PASS: numeric-half-away-negative (FEATURE-NumericStats.md)
  intent: A negative exact tie rounds away from zero, so the magnitude rounds up.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    t,number,2,0,2,-2.00,-1.01,-3.01,-1.50,-1.50
- PASS: numeric-blank-for-non-numeric (FEATURE-NumericStats.md)
  intent: Text, boolean, and empty columns leave sum, mean, and median blank.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    word,text,3,0,3,a,c,,,
    flag,boolean,3,0,2,false,true,,,
    none,empty,3,3,0,,,,,
- PASS: output-profile-shape (FEATURE-CsvOutput.md)
  intent: Profile output has the declared header and one row per column in file order.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    zeta,integer,2,0,2,1,2,3,1.50,1.50
    alpha,text,2,0,2,x,y,,,
    mid,boolean,2,0,2,false,true,,,
- PASS: output-newline-endings (FEATURE-CsvOutput.md)
  intent: Output lines end with a line feed and never a carriage return.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: output-profile-quoting (FEATURE-CsvOutput.md)
  intent: A profile field containing a comma or a double quote round-trips through RFC 4180 quoting.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    notes,text,1,0,1,"remote, part-time","remote, part-time",,,
    
    
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    notes,text,1,0,1,"say ""hi"" now","say ""hi"" now",,,
- PASS: output-freq-quoting (FEATURE-CsvOutput.md)
  intent: A frequency value containing a comma round-trips through RFC 4180 quoting.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    "relocating to Austin, TX",2
- PASS: output-column-single-row (FEATURE-CsvOutput.md)
  intent: --column yields the header and exactly the named column's row.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    beta,integer,1,0,1,2,2,2,2.00,2.00
- PASS: output-failure-silent-stdout (FEATURE-CsvOutput.md)
  intent: A rejected file writes nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpy_iow0rj/ragged.csv: line 3: record has 3 fields, header has 2
- PASS: testscript-exists-and-passes (FEATURE-TestScript.md)
  intent: bin/test.sh exists and exits zero on the application root.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ok  	csvstat/cmd/csvstat	(cached)
    ok  	csvstat/internal/cli	0.013s
    ok  	csvstat/internal/freq	(cached)
    ok  	csvstat/internal/output	(cached)
    ok  	csvstat/internal/profile	(cached)
    ok  	csvstat/internal/reader	0.006s
    ok  	csvstat/internal/typing	(cached)
- PASS: testscript-fails-on-failing-test (FEATURE-TestScript.md)
  intent: bin/test.sh exits non-zero when a test in the module fails.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    --- FAIL: TestFails (0.00s)
        x_test.go:6: expected failure
    FAIL
    FAIL	tmpmod	0.001s
    FAIL
- PASS: testscript-fails-on-vet-finding (FEATURE-TestScript.md)
  intent: bin/test.sh exits non-zero when go vet reports a finding even though tests pass.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    # tmpmod
    # [tmpmod]
    ./x.go:6:9: fmt.Sprintf format %d has arg "text" of wrong type string
- PASS: testscript-passes-on-clean-module (FEATURE-TestScript.md)
  intent: bin/test.sh exits zero for a module whose vet and tests are clean, from a different working directory.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ok  	tmpmod	0.001s

## Repair attempts
- attempt 0 (initial build): failed; 67/79 checks model=haiku; execution 20261024.210112.444Z-6c546620; reason: programmatic acceptance failed: counts-rows-nulls-distinct
- attempt 1 (repair 1): built; 68/79 checks model=haiku; execution 20261024.211617.637Z-3c41261c

## Build summary
Two things to address. `counts-rows-nulls-distinct` is a real bug: `encoding/csv` silently skips blank lines, so a single-column file with an empty value loses that row (9 rows instead of 10). The other failures are harness-side (`environment` errors, and the vacuous `type-text-fallback`).RESULT: SUCCESS

`counts-rows-nulls-distinct` had a real bug, now fixed. `go vet ./...` and `go test -race ./...` pass in all seven packages. I did not run the Python acceptance criteria themselves, since the `drydock_go` harness isn't available here.

**The bug:** `encoding/csv` silently skips blank lines. In a single-column file, an empty value written as a blank line is a record, so the criterion's 10 values came back as 9 rows. I reran that input and csvstat now reports `rows` 10, `nulls` 4, `distinct` 3, as the criterion expects.

**The fix:** `reader.Read` now marks each blank line outside a quoted field as `""` before parsing. It uses `encoding/csv` for the parsing and keeps line numbers unchanged. Blank lines are dropped again when the header has more than one column, so a trailing blank line in a multi-column file is still ignored, as before. Blank lines inside a quoted multi-line field are untouched. I added three reader tests covering these cases.

**Not fixable in code:**
- **`csvreading-*` and `nulltrim-*` (10 checks):** these failed with `ERROR environment`, meaning the harness could not run them, not that an assertion failed. The Go tests they name exist and pass under `go test ./...`. I can't diagnose the harness error from here.
- **`type-text-fallback`:** its body only imports `drydb_placeholder`, which doesn't exist, so it cannot assert anything. `TestInferTextFallback` covers the behaviour it describes.

AC_BROKEN: type-text-fallback

FILES CHANGED:
- internal/reader/reader.go
- internal/reader/reader_test.go

SUMMARY: `reader.Read` now treats a blank line in a single-column file as a record with one empty value, which fixes `counts-rows-nulls-distinct`. Test status: `go vet ./...` and `go test -race ./...` pass in all seven packages. The rest of the earlier work is unchanged.

<blueprint-decisions>
[{"spec":"FEATURE-CsvReading.md","severity":"Low","subject":"Blank lines","decision":"encoding/csv skips blank lines. Reader makes them explicit as empty records when the header has one column, so an empty value in a single-column file counts as a row and a null. In multi-column files blank lines are still skipped. Is that acceptable, or should blank lines in multi-column files be rejected as ragged?"}]
</blueprint-decisions>
