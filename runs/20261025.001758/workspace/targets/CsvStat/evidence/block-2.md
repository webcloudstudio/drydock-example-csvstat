# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-10-24
- resulting state: closed/verified
- story points (combined assembled cost): 17087
- execution id: 20261025.003324.199Z-eed4a650

## Stories built
- Build the csvstat executable from the application root with the standard library only. (MOD-001) [story]
- Parse options that precede exactly one file argument. (CLI-001) [story]
- Read a UTF-8 RFC 4180 CSV file with encoding/csv and trim header names and values. (CSV-001) [story]
- Classify null values: empty, NA, N/A, NULL ignoring case. (CSV-003) [story]
- Reject unknown options, missing option values, and invalid option combinations with exit 2. (CLI-002) [story]
- Reject an empty file, an empty or repeated header name, or a ragged record with exit 3. (CSV-002) [story]
- Exit 1 with no stdout when the file cannot be opened or read. (CSV-004) [story]
- Provide the POSIX bin/test.sh that runs go vet and go test. (MOD-002) [story]

## Acceptance tooling authorization
- FEATURE-Executable-Build.md#build-produces-binary: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Executable-Build.md#build-vet-clean: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CLI-Arguments.md#args-no-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CLI-Arguments.md#args-extra-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CLI-Arguments.md#args-option-after-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CLI-Arguments.md#args-option-then-file-accepted: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CLI-Arguments.md#args-single-file-accepted: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CLI-Option-Validation.md#opts-unknown-option: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CLI-Option-Validation.md#opts-missing-value: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CLI-Option-Validation.md#opts-top-not-positive-integer: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CLI-Option-Validation.md#opts-top-without-freq: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CLI-Option-Validation.md#opts-column-with-freq: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CLI-Option-Validation.md#opts-usage-outranks-missing-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Test-Entry.md#test-entry-passes: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Test-Entry.md#test-entry-runs-from-elsewhere: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 963)
- implements: FEATURE-Executable-Build.md (SP 1155)
- context: ARCHITECTURE_compact.md (SP 399)
- implements: FEATURE-CLI-Arguments.md (SP 1317)
- implements: FEATURE-CSV-Reading.md (SP 1713)
- implements: FEATURE-Null-Values.md (SP 1034)
- implements: FEATURE-CLI-Option-Validation.md (SP 1751)
- implements: FEATURE-CSV-Rejection.md (SP 2081)
- implements: FEATURE-Read-Failure.md (SP 984)
- implements: FEATURE-Test-Entry.md (SP 1014)
- stack: common_compact.md (SP 1653)

## Build directory changes
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvdata/csvdata.go
- internal/csvdata/csvdata_test.go
- internal/csvdata/null.go
- internal/csvdata/null_test.go
- internal/csvdata/validate.go
- internal/csvdata/validate_test.go

## Pre-build acceptance observation
- GREEN (prepassed): build-produces-binary (FEATURE-Executable-Build.md)
  intent: go build of ./cmd/csvstat from the application root exits 0 and writes an executable file.
  return code: 0
- GREEN (prepassed): build-no-third-party-requires (FEATURE-Executable-Build.md)
  intent: go.mod lists no requirements, so the build uses only the standard library.
  return code: 0
- GREEN (prepassed): build-main-is-thin (FEATURE-Executable-Build.md)
  intent: Only cmd/csvstat calls os.Exit; no internal package does.
  return code: 0
- GREEN (prepassed): build-vet-clean (FEATURE-Executable-Build.md)
  intent: go vet over the whole module exits 0.
  return code: 0
- GREEN (prepassed): args-no-file (FEATURE-CLI-Arguments.md)
  intent: Running with no arguments exits 2 and writes nothing to standard output.
  return code: 0
- GREEN (prepassed): args-extra-file (FEATURE-CLI-Arguments.md)
  intent: Two file arguments exit 2 and write nothing to standard output.
  return code: 0
- GREEN (prepassed): args-option-after-file (FEATURE-CLI-Arguments.md)
  intent: An option placed after the file is an extra positional and exits 2 with empty standard output.
  return code: 0
- GREEN (prepassed): args-option-then-file-accepted (FEATURE-CLI-Arguments.md)
  intent: Options followed by exactly one file are not a usage error.
  return code: 0
  stdout:
    csvstat: profiling sources/orders.csv is not implemented yet
- GREEN (prepassed): args-single-file-accepted (FEATURE-CLI-Arguments.md)
  intent: A single file with no options is not a usage error.
  return code: 0
  stdout:
    csvstat: profiling sources/sensors.csv is not implemented yet
- RED: csv-reading-valid-file-exits-zero (FEATURE-CSV-Reading.md)
  intent: A readable CSV whose header and values carry padding spaces is read successfully by the built executable.
  return code: 1
  stdout:
    csvstat: profiling /tmp/tmpeydp5g59/padded.csv is not implemented yet
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-reading-valid-file-exits-zero.py", line 12, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: csv-reading-quoted-comma-exits-zero (FEATURE-CSV-Reading.md)
  intent: A quoted field containing a comma and a doubled quote does not make the file ragged.
  return code: 1
  stdout:
    csvstat: profiling /tmp/tmpwbn12q_a/quoted.csv is not implemented yet
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-reading-quoted-comma-exits-zero.py", line 12, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- UNVERIFIED: csv-reading-trims-header-and-values (FEATURE-CSV-Reading.md)
  intent: The reader trims the header and every value before any other rule.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T20:32:15.331470318-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-24T20:32:15.331519471-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-24T20:32:15.331523665-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T20:32:15.331470318-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-24T20:32:15.331519471-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-24T20:32:15.331523665-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-reading-trims-header-and-values.py", line 5, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T20:32:15.331470318-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-24T20:32:15.331519471-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-24T20:32:15.331523665-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: csv-reading-quoted-field (FEATURE-CSV-Reading.md)
  intent: The reader decodes RFC 4180 quoted fields containing commas.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T20:32:15.697439793-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-24T20:32:15.697482781-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-24T20:32:15.697486243-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T20:32:15.697439793-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-24T20:32:15.697482781-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-24T20:32:15.697486243-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-reading-quoted-field.py", line 5, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T20:32:15.697439793-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-24T20:32:15.697482781-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-24T20:32:15.697486243-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: csv-reading-ragged-records-returned (FEATURE-CSV-Reading.md)
  intent: The reader returns records of differing width with their line numbers so rejection can name the line.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T20:32:16.230867002-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-24T20:32:16.230928773-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-24T20:32:16.230932369-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T20:32:16.230867002-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-24T20:32:16.230928773-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-24T20:32:16.230932369-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-reading-ragged-records-returned.py", line 5, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T20:32:16.230867002-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-24T20:32:16.230928773-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-24T20:32:16.230932369-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: null-values-sentinels (FEATURE-Null-Values.md)
  intent: Empty and the NA, N/A, NULL sentinels in any case are null.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T20:32:16.647720351-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-24T20:32:16.647928084-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-24T20:32:16.647948809-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T20:32:16.647720351-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-24T20:32:16.647928084-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-24T20:32:16.647948809-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "null-values-sentinels.py", line 5, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T20:32:16.647720351-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-24T20:32:16.647928084-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-24T20:32:16.647948809-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: null-values-lookalikes (FEATURE-Null-Values.md)
  intent: Values that merely resemble a null sentinel are not null.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T20:32:17.100217771-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-24T20:32:17.100262238-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-24T20:32:17.100265814-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T20:32:17.100217771-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-24T20:32:17.100262238-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-24T20:32:17.100265814-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "null-values-lookalikes.py", line 5, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T20:32:17.100217771-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-24T20:32:17.100262238-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-24T20:32:17.100265814-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- RED: null-values-file-with-nulls-reads (FEATURE-Null-Values.md)
  intent: A file containing null sentinels of every spelling is read by the built executable without error.
  return code: 1
  stdout:
    csvstat: profiling /tmp/tmp14xaq4ce/nulls.csv is not implemented yet
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "null-values-file-with-nulls-reads.py", line 15, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- GREEN (prepassed): opts-unknown-option (FEATURE-CLI-Option-Validation.md)
  intent: An unknown option exits 2 and writes nothing to standard output.
  return code: 0
- GREEN (prepassed): opts-missing-value (FEATURE-CLI-Option-Validation.md)
  intent: Each value-taking option given without its value exits 2 with empty standard output.
  return code: 0
- GREEN (prepassed): opts-top-not-positive-integer (FEATURE-CLI-Option-Validation.md)
  intent: --top of zero, a negative number, or a non-integer exits 2 with empty standard output.
  return code: 0
- GREEN (prepassed): opts-top-without-freq (FEATURE-CLI-Option-Validation.md)
  intent: --top without --freq exits 2 with empty standard output.
  return code: 0
- GREEN (prepassed): opts-column-with-freq (FEATURE-CLI-Option-Validation.md)
  intent: --column combined with --freq exits 2 with empty standard output.
  return code: 0
- GREEN (prepassed): opts-usage-outranks-missing-file (FEATURE-CLI-Option-Validation.md)
  intent: A usage error is reported as exit 2 even when the named file does not exist.
  return code: 0
- RED: csv-rejection-ragged-supplied-file (FEATURE-CSV-Rejection.md)
  intent: The supplied ragged file is rejected with exit 3 and nothing on standard output.
  return code: 1
  stdout:
    csvstat: profiling /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.001758/build/CsvStat/sources/ragged.csv is not implemented yet
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-rejection-ragged-supplied-file.py", line 8, in <module>
        assert code == 3
               ^^^^^^^^^
    AssertionError
- RED: csv-rejection-empty-file (FEATURE-CSV-Rejection.md)
  intent: A zero-byte file is rejected with exit 3 and nothing on standard output.
  return code: 1
  stdout:
    csvstat: profiling /tmp/tmpgf8l9w3g/empty.csv is not implemented yet
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-rejection-empty-file.py", line 11, in <module>
        assert code == 3
               ^^^^^^^^^
    AssertionError
- RED: csv-rejection-empty-header-name (FEATURE-CSV-Rejection.md)
  intent: A header name that is empty or only spaces after trimming is rejected with exit 3.
  return code: 1
  stdout:
    csvstat: profiling /tmp/tmpq_0lo50k/h0.csv is not implemented yet
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-rejection-empty-header-name.py", line 13, in <module>
        assert code == 3
               ^^^^^^^^^
    AssertionError
- RED: csv-rejection-repeated-header-name (FEATURE-CSV-Rejection.md)
  intent: A repeated header name, including one that repeats only after trimming, is rejected with exit 3.
  return code: 1
  stdout:
    csvstat: profiling /tmp/tmptf6idmiw/r0.csv is not implemented yet
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-rejection-repeated-header-name.py", line 13, in <module>
        assert code == 3
               ^^^^^^^^^
    AssertionError
- RED: csv-rejection-short-and-long-records (FEATURE-CSV-Rejection.md)
  intent: A record with too few or too many fields is rejected with exit 3 and nothing on standard output.
  return code: 1
  stdout:
    csvstat: profiling /tmp/tmp0robcndp/w0.csv is not implemented yet
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-rejection-short-and-long-records.py", line 13, in <module>
        assert code == 3
               ^^^^^^^^^
    AssertionError
- RED: csv-rejection-header-only-is-valid (FEATURE-CSV-Rejection.md)
  intent: A header with no data records is not rejected.
  return code: 1
  stdout:
    csvstat: profiling /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.001758/build/CsvStat/sources/header_only.csv is not implemented yet
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-rejection-header-only-is-valid.py", line 8, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- UNVERIFIED: csv-rejection-diagnostic-names-file-and-line (FEATURE-CSV-Rejection.md)
  intent: The single diagnostic line names the file and the ragged record's line number.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-24T20:32:33.494367726-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-24T20:32:33.497890919-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-24T20:32:33.497937108-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-24T20:32:33.498409843-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"ok  \tcsvstat/internal/cli\t0.004s [no tests to run]\n"}
    {"Time":"2026-10-24T20:32:33.498426775-04:00","Action":"pass","Package":"csvstat/internal/cli","Elapsed":0.004}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-24T20:32:33.494367726-04:00","Action":"start","Package":"csvstat/internal/cli"}\n{"Time":"2026-10-24T20:32:33.497890919-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-24T20:32:33.497937108-04:00","Action":"output","Package":"csvstat/internal… (+331 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "csv-rejection-diagnostic-names-file-and-line.py", line 5, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-24T20:32:33.494367726-04:00","Action":"start","Package":"csvstat/internal/cli"}\n{"Time":"2026-10-24T20:32:33.497890919-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-24T20:32:33.497937108-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"PASS\\n"}\n{"Time":"2026-10-24T20:32:33.498409843-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"ok  \\tcsvstat/internal/cli\\t0.004s [no tests to run]\\n"}\n{"Time":"2026-10-24T20:32:33.498426775-04:00","Action":"pass","Package":"csvstat/internal/cli","Elapsed":0.004}\n', '')
- GREEN (prepassed): read-failure-missing-file (FEATURE-Read-Failure.md)
  intent: A path that does not exist exits 1 with nothing on standard output.
  return code: 0
  stdout:
    csvstat: profiling /tmp/tmpikogouai/no-such-file.csv is not implemented yet
- GREEN (prepassed): read-failure-directory (FEATURE-Read-Failure.md)
  intent: A path that is a directory cannot be read and exits 1 with nothing on standard output.
  return code: 0
  stdout:
    csvstat: profiling /tmp/tmpazovif6m is not implemented yet
- RED: read-failure-distinct-from-rejection (FEATURE-Read-Failure.md)
  intent: A readable but ragged file exits 3 while an unreadable one exits 1, so the two codes are not conflated.
  return code: 1
  stdout:
    csvstat: profiling /tmp/tmp5ybzmi8y/absent.csv is not implemented yet
  stderr:
    --- drydock: values at failure ---
      rejected_code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "read-failure-distinct-from-rejection.py", line 12, in <module>
        assert rejected_code == 3
               ^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): test-entry-syntax (FEATURE-Test-Entry.md)
  intent: bin/test.sh exists and is syntactically valid POSIX sh.
  return code: 0
- GREEN (prepassed): test-entry-passes (FEATURE-Test-Entry.md)
  intent: sh bin/test.sh exits 0 when go vet and go test both pass.
  return code: 0
  stdout:
    ?   	csvstat/cmd/csvstat	[no test files]
    ok  	csvstat/internal/cli	0.014s
    ?   	csvstat/internal/csvdata	[no test files]
    ?   	csvstat/internal/decimal	[no test files]
    ?   	csvstat/internal/freq	[no test files]
    ?   	csvstat/internal/infer	[no test files]
    ?   	csvstat/internal/profile	[no test files]
- GREEN (prepassed): test-entry-runs-from-elsewhere (FEATURE-Test-Entry.md)
  intent: bin/test.sh exits 0 when invoked from a different working directory.
  return code: 0
  stdout:
    ?   	csvstat/cmd/csvstat	[no test files]
    ok  	csvstat/internal/cli	(cached)
    ?   	csvstat/internal/csvdata	[no test files]
    ?   	csvstat/internal/decimal	[no test files]
    ?   	csvstat/internal/freq	[no test files]
    ?   	csvstat/internal/infer	[no test files]
    ?   	csvstat/internal/profile	[no test files]
- GREEN (prepassed): test-entry-suite-nonempty (FEATURE-Test-Entry.md)
  intent: The module contains at least one Go test file.
  return code: 0

## Post-build programmatic acceptance
- PASS: build-produces-binary (FEATURE-Executable-Build.md)
  intent: go build of ./cmd/csvstat from the application root exits 0 and writes an executable file.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: build-no-third-party-requires (FEATURE-Executable-Build.md)
  intent: go.mod lists no requirements, so the build uses only the standard library.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: build-main-is-thin (FEATURE-Executable-Build.md)
  intent: Only cmd/csvstat calls os.Exit; no internal package does.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: build-vet-clean (FEATURE-Executable-Build.md)
  intent: go vet over the whole module exits 0.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: args-no-file (FEATURE-CLI-Arguments.md)
  intent: Running with no arguments exits 2 and writes nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: args-extra-file (FEATURE-CLI-Arguments.md)
  intent: Two file arguments exit 2 and write nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: args-option-after-file (FEATURE-CLI-Arguments.md)
  intent: An option placed after the file is an extra positional and exits 2 with empty standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: args-option-then-file-accepted (FEATURE-CLI-Arguments.md)
  intent: Options followed by exactly one file are not a usage error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: args-single-file-accepted (FEATURE-CLI-Arguments.md)
  intent: A single file with no options is not a usage error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: csv-reading-valid-file-exits-zero (FEATURE-CSV-Reading.md)
  intent: A readable CSV whose header and values carry padding spaces is read successfully by the built executable.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: csv-reading-quoted-comma-exits-zero (FEATURE-CSV-Reading.md)
  intent: A quoted field containing a comma and a doubled quote does not make the file ragged.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- UNVERIFIED: csv-reading-trims-header-and-values (FEATURE-CSV-Reading.md)
  intent: The reader trims the header and every value before any other rule.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: csv-reading-quoted-field (FEATURE-CSV-Reading.md)
  intent: The reader decodes RFC 4180 quoted fields containing commas.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: csv-reading-ragged-records-returned (FEATURE-CSV-Reading.md)
  intent: The reader returns records of differing width with their line numbers so rejection can name the line.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: null-values-sentinels (FEATURE-Null-Values.md)
  intent: Empty and the NA, N/A, NULL sentinels in any case are null.
  provisioning: not required
  error: excluded check: ERROR environment
- UNVERIFIED: null-values-lookalikes (FEATURE-Null-Values.md)
  intent: Values that merely resemble a null sentinel are not null.
  provisioning: not required
  error: excluded check: ERROR environment
- PASS: null-values-file-with-nulls-reads (FEATURE-Null-Values.md)
  intent: A file containing null sentinels of every spelling is read by the built executable without error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: opts-unknown-option (FEATURE-CLI-Option-Validation.md)
  intent: An unknown option exits 2 and writes nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: opts-missing-value (FEATURE-CLI-Option-Validation.md)
  intent: Each value-taking option given without its value exits 2 with empty standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: opts-top-not-positive-integer (FEATURE-CLI-Option-Validation.md)
  intent: --top of zero, a negative number, or a non-integer exits 2 with empty standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: opts-top-without-freq (FEATURE-CLI-Option-Validation.md)
  intent: --top without --freq exits 2 with empty standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: opts-column-with-freq (FEATURE-CLI-Option-Validation.md)
  intent: --column combined with --freq exits 2 with empty standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: opts-usage-outranks-missing-file (FEATURE-CLI-Option-Validation.md)
  intent: A usage error is reported as exit 2 even when the named file does not exist.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: csv-rejection-ragged-supplied-file (FEATURE-CSV-Rejection.md)
  intent: The supplied ragged file is rejected with exit 3 and nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.001758/build/CsvStat/sources/ragged.csv: line 4: record has 4 fields, header has 3
- PASS: csv-rejection-empty-file (FEATURE-CSV-Rejection.md)
  intent: A zero-byte file is rejected with exit 3 and nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmp8i7jc3g9/empty.csv: file has no records
- PASS: csv-rejection-empty-header-name (FEATURE-CSV-Rejection.md)
  intent: A header name that is empty or only spaces after trimming is rejected with exit 3.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpbd732xph/h0.csv: header name 2 is empty
    
    
    csvstat: /tmp/tmpbd732xph/h1.csv: header name 2 is empty
- PASS: csv-rejection-repeated-header-name (FEATURE-CSV-Rejection.md)
  intent: A repeated header name, including one that repeats only after trimming, is rejected with exit 3.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpdsa3r7rz/r0.csv: header name "a" is repeated
    
    
    csvstat: /tmp/tmpdsa3r7rz/r1.csv: header name "a" is repeated
- PASS: csv-rejection-short-and-long-records (FEATURE-CSV-Rejection.md)
  intent: A record with too few or too many fields is rejected with exit 3 and nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpqjytqru9/w0.csv: line 2: record has 2 fields, header has 3
    
    
    csvstat: /tmp/tmpqjytqru9/w1.csv: line 2: record has 4 fields, header has 3
- PASS: csv-rejection-header-only-is-valid (FEATURE-CSV-Rejection.md)
  intent: A header with no data records is not rejected.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- UNVERIFIED: csv-rejection-diagnostic-names-file-and-line (FEATURE-CSV-Rejection.md)
  intent: The single diagnostic line names the file and the ragged record's line number.
  provisioning: not required
  error: excluded check: ERROR environment
- PASS: read-failure-missing-file (FEATURE-Read-Failure.md)
  intent: A path that does not exist exits 1 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmp_k3b7f31/no-such-file.csv: open /tmp/tmp_k3b7f31/no-such-file.csv: no such file or directory
- PASS: read-failure-directory (FEATURE-Read-Failure.md)
  intent: A path that is a directory cannot be read and exits 1 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpqbji7ay9: read /tmp/tmpqbji7ay9: is a directory
- PASS: read-failure-distinct-from-rejection (FEATURE-Read-Failure.md)
  intent: A readable but ragged file exits 3 while an unreadable one exits 1, so the two codes are not conflated.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpfacqp1bv/absent.csv: open /tmp/tmpfacqp1bv/absent.csv: no such file or directory
- PASS: test-entry-syntax (FEATURE-Test-Entry.md)
  intent: bin/test.sh exists and is syntactically valid POSIX sh.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: test-entry-passes (FEATURE-Test-Entry.md)
  intent: sh bin/test.sh exits 0 when go vet and go test both pass.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ?   	csvstat/cmd/csvstat	[no test files]
    ok  	csvstat/internal/cli	0.013s
    ok  	csvstat/internal/csvdata	0.010s
    ?   	csvstat/internal/decimal	[no test files]
    ?   	csvstat/internal/freq	[no test files]
    ?   	csvstat/internal/infer	[no test files]
    ?   	csvstat/internal/profile	[no test files]
- PASS: test-entry-runs-from-elsewhere (FEATURE-Test-Entry.md)
  intent: bin/test.sh exits 0 when invoked from a different working directory.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ?   	csvstat/cmd/csvstat	[no test files]
    ok  	csvstat/internal/cli	(cached)
    ok  	csvstat/internal/csvdata	(cached)
    ?   	csvstat/internal/decimal	[no test files]
    ?   	csvstat/internal/freq	[no test files]
    ?   	csvstat/internal/infer	[no test files]
    ?   	csvstat/internal/profile	[no test files]
- PASS: test-entry-suite-nonempty (FEATURE-Test-Entry.md)
  intent: The module contains at least one Go test file.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
I'll start by inspecting the build directory and the staged sources.Block 1 left a skeleton in place. I'll read those files before writing anything.The skeleton from block 1 has the CLI shell, but two things conflict with this block's specs. `--freq` is declared as a boolean, while the spec form is `--freq <name>`. The file is also never read. I'll read the two files I need to change with the Read tool before editing.The block-1 tests encode `--freq` as a bare flag, so they need updating with the rest. I'll start with the `csvdata` package: reading, trimming, null rule, and rejection.Next, the CLI. `--freq` becomes a name-taking string, `--top` defaults to 5, and the file is now read and validated, with rejected content mapped to exit 3.Now the CLI tests. I'm updating the block-1 cases to the `--freq <name>` form and adding exit-code tests for each read and rejection outcome.Next, the `csvdata` tests. The names of the five tests come directly from the specs; the others cover boundaries and the read-failure classification.I wrote a duplicate table row in `validate_test.go`: the "whitespace-only" case is the same table as the empty-name case, since trimming happens in `Read`. I'll remove it, then run the gate.The gate passes: `gofmt` is clean, `go vet` and `go test` pass. The Python acceptance criteria need the `drydock_go` harness, which isn't on this machine, so I'll exercise the same cases directly against a built binary. The script goes in `/tmp`, outside the build tree.`/tmp/csvstat-check` already exists as a file, so I won't touch it. I'll use a fresh `mktemp` directory.All 43 block-2 cases match their expected exit codes with empty stdout. Last, I'll confirm the diagnostic text, that `sources/` is untouched, and that the tree has only the intended changes.Block 2 is built. `sh bin/test.sh` passes (vet and tests both green), and every acceptance case for this block behaves as specified.

**Verification**
- `go test ./...` passes across `cli` and `csvdata`; `gofmt -l` is clean.
- I ran the block's Python acceptance cases as equivalent shell checks against a built binary: 43 cases, 0 failures. The `drydock_go` harness isn't on this machine, so I couldn't run the Python ACs themselves.
- I didn't run `sources/run.sh`. It's the terminal gate, and its checks for profiling output belong to later blocks, so they will fail until then.

**What changed**
- **`internal/csvdata`** (new files): `Read` parses with `encoding/csv`, trims spaces from each header name and value, and records each record's start line. `IsNull` is the single null rule. `Validate` rejects an empty file, an empty or repeated header name, or a ragged record. A CSV syntax error is classified as a rejection, and an unreadable file returns a `*fs.PathError`.
- **`internal/cli/cli.go`**: after option validation, the file is read and validated. Rejected content exits 3, an unreadable file exits 1, and a valid file exits 0 with nothing on stdout. Profiling output remains for later blocks.
- **`--freq` is now `--freq <name>`.** The block-1 skeleton declared it as a boolean flag, which contradicts the spec form and would make `--freq department file.csv` a usage error. `--top` now defaults to 5.
- **`internal/cli/cli_test.go`**: updated the block-1 cases to the `--freq <name>` form and added exit-code tests for valid, missing, directory, and rejected files.

**Choices to review**
- `TestRejectionDiagnosticNamesFileAndLine` checks the stderr line for the file path and `line 4`. The spec asks for that, and the test is the declared AC. It is a substring check, which goes against the general test rule, so if you'd rather use a structured check I can change it.
- Read and validation live in `csvdata`, which matches its package doc. The `cli` package only maps errors to exit codes.
- `bin/test.sh` is unchanged from block 1. It already satisfies the FEATURE-Test-Entry spec. I did not add the log-to-`data/logs` behavior from the common rules, because that conflicts with the POSIX-only requirement in that spec.

No files under `sources/` were modified, and I did not commit anything.

RESULT: SUCCESS

FILES CHANGED:
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvdata/csvdata.go
- internal/csvdata/csvdata_test.go
- internal/csvdata/null.go
- internal/csvdata/null_test.go
- internal/csvdata/validate.go
- internal/csvdata/validate_test.go

SUMMARY:
Implemented the CSV reader with header/value trimming, the null rule, structural validation (exit 3), read-failure classification (exit 1), and option handling (exit 2). `--freq` now takes a name, and `--top` defaults to 5. Go tests pass (`sh bin/test.sh` exit 0). The 43 acceptance cases for this block were checked via shell against the built binary and all match. Python ACs were not run because the `drydock_go` harness is unavailable here. Pass/fail counts: `go test` 2 packages ok; `sources/run.sh` not invoked (terminal gate).

BLOCKERS:
- none
