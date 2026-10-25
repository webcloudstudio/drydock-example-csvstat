# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-10-25
- resulting state: closed/verified
- story points (combined assembled cost): 19517
- execution id: 20261025.053417.096Z-02b57a21

## Stories built
- Build the csvstat executable from the module root with a thin main over cli.Run. (executable) [story]
- Read UTF-8 RFC 4180 CSV with the first record as header. (csv-reading) [story]
- Trim leading and trailing spaces from every header name and value. (value-trimming) [story]
- Recognise null spellings: empty, NA, N/A, NULL ignoring case. (null-values) [story]
- Infer integer and number column types. (numeric-types) [story]
- Infer date columns of real YYYY-MM-DD calendar dates. (date-type) [story]
- Infer boolean, text, and empty column types. (boolean-text-types) [story]
- Parse options and exactly one file argument, with options preceding the file. (args-parsing) [story]
- Reject empty files, empty or repeated header names, and ragged records. (input-rejection) [story]

## Acceptance tooling authorization
- FEATURE-Executable.md#exe-builds: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvReading.md#read-header-and-records: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvReading.md#read-quoted-comma: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvReading.md#read-trailing-newline: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvReading.md#read-line-numbers: executable=go; scope=test; authorization=existing Target environment
- FEATURE-CsvReading.md#read-header-only-and-empty: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Trimming.md#trim-header-names: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Trimming.md#trim-values: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Trimming.md#trim-quoted-padding: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Trimming.md#trim-scope: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NullValues.md#null-empty: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NullValues.md#null-spellings: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NullValues.md#null-non-null-values: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NullValues.md#null-quoted-empty: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NullValues.md#null-spaces-only: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NumericTypes.md#numeric-integer: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NumericTypes.md#numeric-number: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NumericTypes.md#numeric-integer-preferred: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NumericTypes.md#numeric-nulls-ignored: executable=go; scope=test; authorization=existing Target environment
- FEATURE-NumericTypes.md#numeric-malformed: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateType.md#date-column: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateType.md#date-invalid-calendar: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateType.md#date-leap-day: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateType.md#date-shape-strict: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DateType.md#date-nulls-ignored: executable=go; scope=test; authorization=existing Target environment
- FEATURE-BooleanTextTypes.md#type-boolean-mixed-case: executable=go; scope=test; authorization=existing Target environment
- FEATURE-BooleanTextTypes.md#type-text-fallback: executable=go; scope=test; authorization=existing Target environment
- FEATURE-BooleanTextTypes.md#type-empty-all-null: executable=go; scope=test; authorization=existing Target environment
- FEATURE-BooleanTextTypes.md#type-empty-no-rows: executable=go; scope=test; authorization=existing Target environment
- FEATURE-BooleanTextTypes.md#type-boolean-lookalikes: executable=go; scope=test; authorization=existing Target environment
- FEATURE-BooleanTextTypes.md#type-order-precedence: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Arguments.md#args-file-only: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Arguments.md#args-column: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Arguments.md#args-freq-top: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Arguments.md#args-structural-errors: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Arguments.md#args-valid-forms-not-usage-errors: executable=go; scope=test; authorization=existing Target environment
- FEATURE-InputRejection.md#reject-unit-empty-and-blank: executable=go; scope=test; authorization=existing Target environment
- FEATURE-InputRejection.md#reject-unit-headers: executable=go; scope=test; authorization=existing Target environment
- FEATURE-InputRejection.md#reject-unit-ragged-and-valid: executable=go; scope=test; authorization=existing Target environment
- FEATURE-InputRejection.md#reject-cli-ragged: executable=go; scope=test; authorization=existing Target environment
- FEATURE-InputRejection.md#reject-cli-empty-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-InputRejection.md#reject-cli-blank-lines-only: executable=go; scope=test; authorization=existing Target environment
- FEATURE-InputRejection.md#reject-cli-repeated-header: executable=go; scope=test; authorization=existing Target environment
- FEATURE-InputRejection.md#reject-cli-empty-header-name: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 963)
- implements: FEATURE-Executable.md (SP 1392)
- context: ARCHITECTURE_compact.md (SP 282)
- implements: FEATURE-CsvReading.md (SP 1824)
- implements: FEATURE-Trimming.md (SP 1379)
- implements: FEATURE-NullValues.md (SP 1393)
- implements: FEATURE-NumericTypes.md (SP 1494)
- implements: FEATURE-DateType.md (SP 1362)
- implements: FEATURE-BooleanTextTypes.md (SP 1602)
- implements: FEATURE-Arguments.md (SP 2348)
- implements: FEATURE-InputRejection.md (SP 2529)

## Build directory changes
- internal/cli/args.go
- internal/cli/args_test.go
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvdata/null.go
- internal/csvdata/null_test.go
- internal/csvdata/read.go
- internal/csvdata/read_test.go
- internal/csvdata/trim.go
- internal/csvdata/trim_test.go
- internal/csvdata/validate.go
- internal/csvdata/validate_test.go
- internal/profile/infer.go
- internal/profile/infer_test.go

## Pre-build acceptance observation
- GREEN (prepassed): exe-builds (FEATURE-Executable.md)
  intent: The csvstat executable builds with go build and runs to a defined exit code on a valid file.
  return code: 0
  stdout:
    1  csvstat: not implemented
- GREEN (prepassed): exe-stdlib-only (FEATURE-Executable.md)
  intent: Every import in the module is standard library or module-local, with no third-party path.
  return code: 0
- GREEN (prepassed): exe-main-owns-exit (FEATURE-Executable.md)
  intent: Only the main package calls os.Exit and main delegates to cli.Run.
  return code: 0
- GREEN (prepassed): exe-no-grader-read (FEATURE-Executable.md)
  intent: Application code never references the supplied grader script.
  return code: 0
- UNVERIFIED: read-header-and-records (FEATURE-CsvReading.md)
  intent: The first record is the header and later records are data records.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:45.18726585-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:45.187330412-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:45.187335205-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:45.18726585-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:45.187330412-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:45.187335205-04:00","Action":"skip","Packa… (+68 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "read-header-and-records.py", line 7, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:45.18726585-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:45.187330412-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:45.187335205-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: read-quoted-comma (FEATURE-CsvReading.md)
  intent: A quoted field containing a comma parses as one value.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:46.2053361-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:46.205394178-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:46.205398178-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:46.2053361-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:46.205394178-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:46.205398178-04:00","Action":"skip","Packag… (+67 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "read-quoted-comma.py", line 7, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:46.2053361-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:46.205394178-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:46.205398178-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: read-trailing-newline (FEATURE-CsvReading.md)
  intent: A line ending after the last record does not create an extra data record.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:47.224846235-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:47.224912986-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:47.2249173-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:47.224846235-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:47.224912986-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:47.2249173-04:00","Action":"skip","Packag… (+67 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "read-trailing-newline.py", line 7, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:47.224846235-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:47.224912986-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:47.2249173-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: read-line-numbers (FEATURE-CsvReading.md)
  intent: Each record carries the source line on which it starts.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:48.240426667-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:48.240516398-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:48.240520762-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:48.240426667-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:48.240516398-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:48.240520762-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "read-line-numbers.py", line 7, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:48.240426667-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:48.240516398-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:48.240520762-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: read-header-only-and-empty (FEATURE-CsvReading.md)
  intent: A header-only input has zero records and an input with no record has no header.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestReadHeaderOnly not-run
    {"Time":"2026-10-25T01:32:49.243852843-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:49.243898499-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:49.243902377-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:49.243852843-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:49.243898499-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:49.243902377-04:00","Action":"skip","Pack… (+69 chars)
      name = 'TestReadHeaderOnly'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "read-header-only-and-empty.py", line 9, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestReadHeaderOnly', 'not-run', '{"Time":"2026-10-25T01:32:49.243852843-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:49.243898499-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:49.243902377-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: trim-header-names (FEATURE-Trimming.md)
  intent: Leading and trailing spaces are removed from header names.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:50.244150378-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:50.244207365-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:50.244212678-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:50.244150378-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:50.244207365-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:50.244212678-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "trim-header-names.py", line 7, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:50.244150378-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:50.244207365-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:50.244212678-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: trim-values (FEATURE-Trimming.md)
  intent: Leading and trailing spaces are removed from every value.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:51.30725068-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:51.307321977-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:51.3073264-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:51.30725068-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:51.307321977-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:51.3073264-04:00","Action":"skip","Package… (+66 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "trim-values.py", line 7, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:51.30725068-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:51.307321977-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:51.3073264-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: trim-quoted-padding (FEATURE-Trimming.md)
  intent: Padding inside a quoted field is trimmed while the quoted comma is kept.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:52.205404276-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:52.205457722-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:52.205461552-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:52.205404276-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:52.205457722-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:52.205461552-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "trim-quoted-padding.py", line 7, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:52.205404276-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:52.205457722-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:52.205461552-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: trim-scope (FEATURE-Trimming.md)
  intent: Only spaces are trimmed, interior spaces remain, and line numbers and field counts are unchanged.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestTrimOnlySpaces not-run
    {"Time":"2026-10-25T01:32:53.207612288-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:53.20766843-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:53.207672116-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:53.207612288-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:53.20766843-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:53.207672116-04:00","Action":"skip","Packa… (+68 chars)
      name = 'TestTrimOnlySpaces'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "trim-scope.py", line 9, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestTrimOnlySpaces', 'not-run', '{"Time":"2026-10-25T01:32:53.207612288-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:53.20766843-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:53.207672116-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: null-empty (FEATURE-NullValues.md)
  intent: The empty string is reported null by csvdata.IsNull.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:54.130958072-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:54.1310074-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:54.131011177-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:54.130958072-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:54.1310074-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:54.131011177-04:00","Action":"skip","Packag… (+67 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "null-empty.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:54.130958072-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:54.1310074-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:54.131011177-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: null-spellings (FEATURE-NullValues.md)
  intent: NA, N/A and NULL are null in any letter case.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:54.987141482-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:54.98719425-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:54.987198299-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:54.987141482-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:54.98719425-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:54.987198299-04:00","Action":"skip","Packa… (+68 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "null-spellings.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:54.987141482-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:54.98719425-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:54.987198299-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: null-non-null-values (FEATURE-NullValues.md)
  intent: Ordinary values such as 0, false and near-miss spellings are not null.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:55.751331408-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:55.751387548-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:55.751391292-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:55.751331408-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:55.751387548-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:55.751391292-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "null-non-null-values.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:55.751331408-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:55.751387548-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:55.751391292-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: null-quoted-empty (FEATURE-NullValues.md)
  intent: A quoted empty field read from a CSV file is null.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:56.534071284-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:56.534408087-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:56.53441238-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:56.534071284-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:56.534408087-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:56.53441238-04:00","Action":"skip","Packa… (+68 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "null-quoted-empty.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:56.534071284-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:56.534408087-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:56.53441238-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: null-spaces-only (FEATURE-NullValues.md)
  intent: A field of only spaces is null once trimmed.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:57.367785641-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:32:57.367840426-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:57.367844049-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:57.367785641-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:57.367840426-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:57.367844049-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "null-spaces-only.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:57.367785641-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:32:57.367840426-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:57.367844049-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: numeric-integer (FEATURE-NumericTypes.md)
  intent: A column of integer-shaped values, including negatives, is typed integer.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:58.232698792-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:32:58.232753412-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:58.232757204-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:58.232698792-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:32:58.232753412-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:58.232757204-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numeric-integer.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:58.232698792-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:32:58.232753412-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:58.232757204-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: numeric-number (FEATURE-NumericTypes.md)
  intent: A column with a decimal value among integers is typed number.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:59.063508316-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:32:59.063563342-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:59.063567118-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:59.063508316-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:32:59.063563342-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:59.063567118-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numeric-number.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:59.063508316-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:32:59.063563342-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:59.063567118-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: numeric-integer-preferred (FEATURE-NumericTypes.md)
  intent: Integer is chosen over number when both patterns match every value.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:32:59.85405653-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:32:59.854110139-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:32:59.854114184-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:32:59.85405653-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:32:59.854110139-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:59.854114184-04:00","Action":"skip","Packa… (+68 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numeric-integer-preferred.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:32:59.85405653-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:32:59.854110139-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:32:59.854114184-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: numeric-nulls-ignored (FEATURE-NumericTypes.md)
  intent: Null spellings among numeric values do not change the inferred type.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:00.525075608-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:00.525136211-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:00.525140339-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:00.525075608-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:00.525136211-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:00.525140339-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numeric-nulls-ignored.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:00.525075608-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:00.525136211-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:00.525140339-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: numeric-malformed (FEATURE-NumericTypes.md)
  intent: Values such as +1, 1., .5, 1e3 and 1,5 are not numeric.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:01.366154131-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:01.36620385-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:01.366207653-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:01.366154131-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:01.36620385-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:01.366207653-04:00","Action":"skip","Packa… (+68 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "numeric-malformed.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:01.366154131-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:01.36620385-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:01.366207653-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: date-column (FEATURE-DateType.md)
  intent: A column of valid YYYY-MM-DD values is typed date.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:02.169841141-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:02.169904724-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:02.169909096-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:02.169841141-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:02.169904724-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:02.169909096-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-column.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:02.169841141-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:02.169904724-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:02.169909096-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: date-invalid-calendar (FEATURE-DateType.md)
  intent: A column containing an impossible date such as 2025-02-30 is not typed date.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:03.042783289-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:03.042858794-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:03.042862611-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:03.042783289-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:03.042858794-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:03.042862611-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-invalid-calendar.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:03.042783289-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:03.042858794-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:03.042862611-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: date-leap-day (FEATURE-DateType.md)
  intent: A leap-year February 29 is a date and a non-leap-year February 29 is not.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:03.920400587-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:03.920452914-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:03.920456652-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:03.920400587-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:03.920452914-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:03.920456652-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-leap-day.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:03.920400587-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:03.920452914-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:03.920456652-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: date-shape-strict (FEATURE-DateType.md)
  intent: Values with the wrong shape (unpadded, slashes, timestamps) are not dates.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:04.720669819-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:04.720729517-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:04.720733496-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:04.720669819-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:04.720729517-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:04.720733496-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-shape-strict.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:04.720669819-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:04.720729517-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:04.720733496-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: date-nulls-ignored (FEATURE-DateType.md)
  intent: Null spellings among dates do not change the inferred date type.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:05.482570687-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:05.482622854-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:05.48262672-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:05.482570687-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:05.482622854-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:05.48262672-04:00","Action":"skip","Packa… (+68 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "date-nulls-ignored.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:05.482570687-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:05.482622854-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:05.48262672-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: type-boolean-mixed-case (FEATURE-BooleanTextTypes.md)
  intent: Values true and false in any letter case form a boolean column.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:06.289917353-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:06.290081734-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:06.290086314-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:06.289917353-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:06.290081734-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:06.290086314-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "type-boolean-mixed-case.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:06.289917353-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:06.290081734-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:06.290086314-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: type-text-fallback (FEATURE-BooleanTextTypes.md)
  intent: A column matching no earlier type is text.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:07.115507447-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:07.115555284-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:07.115559402-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:07.115507447-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:07.115555284-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:07.115559402-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "type-text-fallback.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:07.115507447-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:07.115555284-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:07.115559402-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: type-empty-all-null (FEATURE-BooleanTextTypes.md)
  intent: A column whose values are all null spellings is typed empty.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:07.960164037-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:07.960216853-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:07.960221082-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:07.960164037-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:07.960216853-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:07.960221082-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "type-empty-all-null.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:07.960164037-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:07.960216853-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:07.960221082-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: type-empty-no-rows (FEATURE-BooleanTextTypes.md)
  intent: A column with zero data rows is typed empty.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:08.752287504-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:08.752351691-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:08.752356198-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:08.752287504-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:08.752351691-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:08.752356198-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "type-empty-no-rows.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:08.752287504-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:08.752351691-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:08.752356198-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: type-boolean-lookalikes (FEATURE-BooleanTextTypes.md)
  intent: Values such as yes, 1 or t do not make a boolean column.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:09.697209083-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:09.697257554-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:09.697261583-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:09.697209083-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:09.697257554-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:09.697261583-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "type-boolean-lookalikes.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:09.697209083-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:09.697257554-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:09.697261583-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: type-order-precedence (FEATURE-BooleanTextTypes.md)
  intent: Earlier types win over later ones when several would match.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:10.607811982-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:33:10.607869352-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \tcsvstat/internal/profile\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:10.607872952-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:10.607811982-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:10.607869352-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:10.607872952-04:00","Action":"skip","Pack… (+69 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "type-order-precedence.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:10.607811982-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:33:10.607869352-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"?   \\tcsvstat/internal/profile\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:10.607872952-04:00","Action":"skip","Package":"csvstat/internal/profile","Elapsed":0}\n', '')
- UNVERIFIED: args-file-only (FEATURE-Arguments.md)
  intent: A lone file argument parses.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:11.423440834-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-25T01:33:11.423493506-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"?   \tcsvstat/internal/cli\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:11.423497915-04:00","Action":"skip","Package":"csvstat/internal/cli","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:11.423440834-04:00","Action":"start","Package":"csvstat/internal/cli"}\n{"Time":"2026-10-25T01:33:11.423493506-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"?   \\tcsvstat/internal/cli\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:11.423497915-04:00","Action":"skip","Package":"csvsta… (+53 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "args-file-only.py", line 7, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:11.423440834-04:00","Action":"start","Package":"csvstat/internal/cli"}\n{"Time":"2026-10-25T01:33:11.423493506-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"?   \\tcsvstat/internal/cli\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:11.423497915-04:00","Action":"skip","Package":"csvstat/internal/cli","Elapsed":0}\n', '')
- UNVERIFIED: args-column (FEATURE-Arguments.md)
  intent: The --column option and its value parse before the file.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:12.557387562-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-25T01:33:12.557431902-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"?   \tcsvstat/internal/cli\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:12.557436026-04:00","Action":"skip","Package":"csvstat/internal/cli","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:12.557387562-04:00","Action":"start","Package":"csvstat/internal/cli"}\n{"Time":"2026-10-25T01:33:12.557431902-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"?   \\tcsvstat/internal/cli\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:12.557436026-04:00","Action":"skip","Package":"csvsta… (+53 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "args-column.py", line 7, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:12.557387562-04:00","Action":"start","Package":"csvstat/internal/cli"}\n{"Time":"2026-10-25T01:33:12.557431902-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"?   \\tcsvstat/internal/cli\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:12.557436026-04:00","Action":"skip","Package":"csvstat/internal/cli","Elapsed":0}\n', '')
- UNVERIFIED: args-freq-top (FEATURE-Arguments.md)
  intent: The --freq and --top options and their values parse before the file.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:13.629408169-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-25T01:33:13.629450869-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"?   \tcsvstat/internal/cli\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:13.629455253-04:00","Action":"skip","Package":"csvstat/internal/cli","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:13.629408169-04:00","Action":"start","Package":"csvstat/internal/cli"}\n{"Time":"2026-10-25T01:33:13.629450869-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"?   \\tcsvstat/internal/cli\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:13.629455253-04:00","Action":"skip","Package":"csvsta… (+53 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "args-freq-top.py", line 7, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:33:13.629408169-04:00","Action":"start","Package":"csvstat/internal/cli"}\n{"Time":"2026-10-25T01:33:13.629450869-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"?   \\tcsvstat/internal/cli\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:13.629455253-04:00","Action":"skip","Package":"csvstat/internal/cli","Elapsed":0}\n', '')
- UNVERIFIED: args-structural-errors (FEATURE-Arguments.md)
  intent: Missing file, extra file, unknown option, missing value, and options after the file are parse errors.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    TestParseArgsMissingFile not-run
    {"Time":"2026-10-25T01:33:14.820272561-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-25T01:33:14.820317631-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"?   \tcsvstat/internal/cli\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:14.820321733-04:00","Action":"skip","Package":"csvstat/internal/cli","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:14.820272561-04:00","Action":"start","Package":"csvstat/internal/cli"}\n{"Time":"2026-10-25T01:33:14.820317631-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"?   \\tcsvstat/internal/cli\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:14.820321733-04:00","Action":"skip","Package":"csvsta… (+53 chars)
      name = 'TestParseArgsMissingFile'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "args-structural-errors.py", line 16, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestParseArgsMissingFile', 'not-run', '{"Time":"2026-10-25T01:33:14.820272561-04:00","Action":"start","Package":"csvstat/internal/cli"}\n{"Time":"2026-10-25T01:33:14.820317631-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"?   \\tcsvstat/internal/cli\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:14.820321733-04:00","Action":"skip","Package":"csvstat/internal/cli","Elapsed":0}\n', '')
- GREEN (prepassed): args-valid-forms-not-usage-errors (FEATURE-Arguments.md)
  intent: The valid command forms are accepted by the executable and never exit with the usage-error code.
  return code: 0
  stdout:
    ['/tmp/tmpq7tlvj09/data.csv'] 1  csvstat: not implemented
    
    ['--column', 'a', '/tmp/tmpq7tlvj09/data.csv'] 1  csvstat: not implemented
    
    ['--freq', 'a', '/tmp/tmpq7tlvj09/data.csv'] 1  csvstat: not implemented
    
    ['--freq', 'a', '--top', '3', '/tmp/tmpq7tlvj09/data.csv'] 1  csvstat: not implemented
- RED: args-options-after-file-exit (FEATURE-Arguments.md)
  intent: An option placed after the file is rejected by the executable with the usage-error exit status and no standard output.
  return code: 1
  stdout:
    1  csvstat: not implemented
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "args-options-after-file-exit.py", line 14, in <module>
        assert code == 2
               ^^^^^^^^^
    AssertionError
- UNVERIFIED: reject-unit-empty-and-blank (FEATURE-InputRejection.md)
  intent: Empty files and files of only blank lines are rejected by the reader.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:22.268341343-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:33:22.268398515-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:22.268402452-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:22.268341343-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:33:22.268398515-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:22.268402452-04:00","Action":"skip","Pack… (+69 chars)
      name = 'TestRejectEmptyFile'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-unit-empty-and-blank.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestRejectEmptyFile', 'not-run', '{"Time":"2026-10-25T01:33:22.268341343-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:33:22.268398515-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:22.268402452-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: reject-unit-headers (FEATURE-InputRejection.md)
  intent: Empty and repeated header names are rejected by the reader.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:23.127763383-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:33:23.127946096-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:23.127951686-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:23.127763383-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:33:23.127946096-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:23.127951686-04:00","Action":"skip","Pack… (+69 chars)
      name = 'TestRejectEmptyHeaderName'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-unit-headers.py", line 7, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestRejectEmptyHeaderName', 'not-run', '{"Time":"2026-10-25T01:33:23.127763383-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:33:23.127946096-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:23.127951686-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- UNVERIFIED: reject-unit-ragged-and-valid (FEATURE-InputRejection.md)
  intent: Ragged records are rejected while header-only and newline-terminated files are accepted.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:33:23.968363485-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:33:23.968440093-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \tcsvstat/internal/csvdata\t[no test files]\n"}
    {"Time":"2026-10-25T01:33:23.968446918-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:33:23.968363485-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:33:23.968440093-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:23.968446918-04:00","Action":"skip","Pack… (+69 chars)
      name = 'TestRejectRaggedRecord'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-unit-ragged-and-valid.py", line 11, in <module>
        assert result.passed, (name, result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('TestRejectRaggedRecord', 'not-run', '{"Time":"2026-10-25T01:33:23.968363485-04:00","Action":"start","Package":"csvstat/internal/csvdata"}\n{"Time":"2026-10-25T01:33:23.968440093-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"?   \\tcsvstat/internal/csvdata\\t[no test files]\\n"}\n{"Time":"2026-10-25T01:33:23.968446918-04:00","Action":"skip","Package":"csvstat/internal/csvdata","Elapsed":0}\n', '')
- RED: reject-cli-ragged (FEATURE-InputRejection.md)
  intent: The executable exits 3 with empty standard output for a ragged record.
  return code: 1
  stdout:
    csvstat: not implemented
  stderr:
    --- drydock: values at failure ---
      code = 1
      out = ''
      err = 'csvstat: not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-cli-ragged.py", line 13, in <module>
        assert code == 3, (code, out, err)
               ^^^^^^^^^
    AssertionError: (1, '', 'csvstat: not implemented\n')
- RED: reject-cli-empty-file (FEATURE-InputRejection.md)
  intent: The executable exits 3 with empty standard output for a zero-byte file.
  return code: 1
  stdout:
    csvstat: not implemented
  stderr:
    --- drydock: values at failure ---
      code = 1
      out = ''
      err = 'csvstat: not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-cli-empty-file.py", line 13, in <module>
        assert code == 3, (code, out, err)
               ^^^^^^^^^
    AssertionError: (1, '', 'csvstat: not implemented\n')
- RED: reject-cli-blank-lines-only (FEATURE-InputRejection.md)
  intent: The executable exits 3 with empty standard output for a file holding only blank lines.
  return code: 1
  stdout:
    csvstat: not implemented
  stderr:
    --- drydock: values at failure ---
      code = 1
      out = ''
      err = 'csvstat: not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-cli-blank-lines-only.py", line 13, in <module>
        assert code == 3, (code, out, err)
               ^^^^^^^^^
    AssertionError: (1, '', 'csvstat: not implemented\n')
- RED: reject-cli-repeated-header (FEATURE-InputRejection.md)
  intent: The executable exits 3 with empty standard output when a header name repeats.
  return code: 1
  stdout:
    csvstat: not implemented
  stderr:
    --- drydock: values at failure ---
      code = 1
      out = ''
      err = 'csvstat: not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-cli-repeated-header.py", line 16, in <module>
        assert code == 3, (code, out, err)
               ^^^^^^^^^
    AssertionError: (1, '', 'csvstat: not implemented\n')
- RED: reject-cli-empty-header-name (FEATURE-InputRejection.md)
  intent: The executable exits 3 with empty standard output when a header name is empty.
  return code: 1
  stdout:
    csvstat: not implemented
  stderr:
    --- drydock: values at failure ---
      code = 1
      out = ''
      err = 'csvstat: not implemented\n'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "reject-cli-empty-header-name.py", line 16, in <module>
        assert code == 3, (code, out, err)
               ^^^^^^^^^
    AssertionError: (1, '', 'csvstat: not implemented\n')

## Post-build programmatic acceptance
- PASS: exe-builds (FEATURE-Executable.md)
  intent: The csvstat executable builds with go build and runs to a defined exit code on a valid file.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    0
- PASS: exe-stdlib-only (FEATURE-Executable.md)
  intent: Every import in the module is standard library or module-local, with no third-party path.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: exe-main-owns-exit (FEATURE-Executable.md)
  intent: Only the main package calls os.Exit and main delegates to cli.Run.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: exe-no-grader-read (FEATURE-Executable.md)
  intent: Application code never references the supplied grader script.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: read-header-and-records (FEATURE-CsvReading.md)
  intent: The first record is the header and later records are data records.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:04.24743743-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:04.254178537-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestReadHeaderAndRecords"}
    {"Time":"2026-10-25T01:37:04.254216939-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestReadHeaderAndRecords","Output":"=== RUN   TestReadHeaderAndRecords\n"}
    {"Time":"2026-10-25T01:37:04.254225198-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestReadHeaderAndRecords","Output":"--- PASS: TestReadHeaderAndRecords (0.00s)\n"}
    {"Time":"2026-10-25T01:37:04.254226861-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestReadHeaderAndRecords","Elapsed":0}
    {"Time":"2026-10-25T01:37:04.254229735-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:04.255485403-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.008s\n"}
    {"Time":"2026-10-25T01:37:04.255509323-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: read-quoted-comma (FEATURE-CsvReading.md)
  intent: A quoted field containing a comma parses as one value.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:05.185526136-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:05.1924836-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestReadQuotedComma"}
    {"Time":"2026-10-25T01:37:05.192507417-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestReadQuotedComma","Output":"=== RUN   TestReadQuotedComma\n"}
    {"Time":"2026-10-25T01:37:05.192527562-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestReadQuotedComma","Output":"--- PASS: TestReadQuotedComma (0.00s)\n"}
    {"Time":"2026-10-25T01:37:05.192529502-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestReadQuotedComma","Elapsed":0}
    {"Time":"2026-10-25T01:37:05.192532343-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:05.193107825-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.007s\n"}
    {"Time":"2026-10-25T01:37:05.193130716-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: read-trailing-newline (FEATURE-CsvReading.md)
  intent: A line ending after the last record does not create an extra data record.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:06.056307758-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:06.063082873-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestReadTrailingNewline"}
    {"Time":"2026-10-25T01:37:06.063113037-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestReadTrailingNewline","Output":"=== RUN   TestReadTrailingNewline\n"}
    {"Time":"2026-10-25T01:37:06.063135834-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestReadTrailingNewline","Output":"--- PASS: TestReadTrailingNewline (0.00s)\n"}
    {"Time":"2026-10-25T01:37:06.063137641-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestReadTrailingNewline","Elapsed":0}
    {"Time":"2026-10-25T01:37:06.063140418-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:06.063918245-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.007s\n"}
    {"Time":"2026-10-25T01:37:06.063940548-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: read-line-numbers (FEATURE-CsvReading.md)
  intent: Each record carries the source line on which it starts.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:07.05596808-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:07.060597269-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestReadRecordLineNumbers"}
    {"Time":"2026-10-25T01:37:07.060646553-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestReadRecordLineNumbers","Output":"=== RUN   TestReadRecordLineNumbers\n"}
    {"Time":"2026-10-25T01:37:07.06066495-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestReadRecordLineNumbers","Output":"--- PASS: TestReadRecordLineNumbers (0.00s)\n"}
    {"Time":"2026-10-25T01:37:07.060667856-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestReadRecordLineNumbers","Elapsed":0}
    {"Time":"2026-10-25T01:37:07.06067104-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:07.061186653-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.005s\n"}
    {"Time":"2026-10-25T01:37:07.061216436-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.005}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: read-header-only-and-empty (FEATURE-CsvReading.md)
  intent: A header-only input has zero records and an input with no record has no header.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestReadHeaderOnly pass
    {"Time":"2026-10-25T01:37:08.079826746-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:08.085929937-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestReadHeaderOnly"}
    {"Time":"2026-10-25T01:37:08.085957656-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestReadHeaderOnly","Output":"=== RUN   TestReadHeaderOnly\n"}
    {"Time":"2026-10-25T01:37:08.085979269-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestReadHeaderOnly","Output":"--- PASS: TestReadHeaderOnly (0.00s)\n"}
    {"Time":"2026-10-25T01:37:08.085983793-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestReadHeaderOnly","Elapsed":0}
    {"Time":"2026-10-25T01:37:08.085986693-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:08.086612749-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.007s\n"}
    {"Time":"2026-10-25T01:37:08.08663365-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.007}
    
    TestReadNoRecords pass
    {"Time":"2026-10-25T01:37:08.94585664-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:08.951356732-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestReadNoRecords"}
    {"Time":"2026-10-25T01:37:08.951383142-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestReadNoRecords","Output":"=== RUN   TestReadNoRecords\n"}
    {"Time":"2026-10-25T01:37:08.951401931-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestReadNoRecords","Output":"--- PASS: TestReadNoRecords (0.00s)\n"}
    {"Time":"2026-10-25T01:37:08.951405177-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestReadNoRecords","Elapsed":0}
    {"Time":"2026-10-25T01:37:08.951408008-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:08.951688291-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.006s\n"}
    {"Time":"2026-10-25T01:37:08.951710245-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: trim-header-names (FEATURE-Trimming.md)
  intent: Leading and trailing spaces are removed from header names.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:09.908714411-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:09.916041607-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestTrimHeaderNames"}
    {"Time":"2026-10-25T01:37:09.916066111-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestTrimHeaderNames","Output":"=== RUN   TestTrimHeaderNames\n"}
    {"Time":"2026-10-25T01:37:09.916094995-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestTrimHeaderNames","Output":"--- PASS: TestTrimHeaderNames (0.00s)\n"}
    {"Time":"2026-10-25T01:37:09.916098081-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestTrimHeaderNames","Elapsed":0}
    {"Time":"2026-10-25T01:37:09.91610102-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:09.916601167-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.008s\n"}
    {"Time":"2026-10-25T01:37:09.91662401-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: trim-values (FEATURE-Trimming.md)
  intent: Leading and trailing spaces are removed from every value.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:10.619176268-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:10.625090924-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestTrimValues"}
    {"Time":"2026-10-25T01:37:10.625120191-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestTrimValues","Output":"=== RUN   TestTrimValues\n"}
    {"Time":"2026-10-25T01:37:10.625140628-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestTrimValues","Output":"--- PASS: TestTrimValues (0.00s)\n"}
    {"Time":"2026-10-25T01:37:10.625142232-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestTrimValues","Elapsed":0}
    {"Time":"2026-10-25T01:37:10.625145038-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:10.6258272-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.006s\n"}
    {"Time":"2026-10-25T01:37:10.625851064-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: trim-quoted-padding (FEATURE-Trimming.md)
  intent: Padding inside a quoted field is trimmed while the quoted comma is kept.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:11.225015892-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:11.230486359-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestTrimQuotedPadding"}
    {"Time":"2026-10-25T01:37:11.230517472-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestTrimQuotedPadding","Output":"=== RUN   TestTrimQuotedPadding\n"}
    {"Time":"2026-10-25T01:37:11.230537862-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestTrimQuotedPadding","Output":"--- PASS: TestTrimQuotedPadding (0.00s)\n"}
    {"Time":"2026-10-25T01:37:11.230539626-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestTrimQuotedPadding","Elapsed":0}
    {"Time":"2026-10-25T01:37:11.230542438-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:11.230958715-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.006s\n"}
    {"Time":"2026-10-25T01:37:11.230981293-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: trim-scope (FEATURE-Trimming.md)
  intent: Only spaces are trimmed, interior spaces remain, and line numbers and field counts are unchanged.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestTrimOnlySpaces pass
    {"Time":"2026-10-25T01:37:12.154065311-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:12.165096181-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestTrimOnlySpaces"}
    {"Time":"2026-10-25T01:37:12.165131065-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestTrimOnlySpaces","Output":"=== RUN   TestTrimOnlySpaces\n"}
    {"Time":"2026-10-25T01:37:12.165152667-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestTrimOnlySpaces","Output":"--- PASS: TestTrimOnlySpaces (0.00s)\n"}
    {"Time":"2026-10-25T01:37:12.165155855-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestTrimOnlySpaces","Elapsed":0}
    {"Time":"2026-10-25T01:37:12.165158897-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:12.16588822-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.012s\n"}
    {"Time":"2026-10-25T01:37:12.165910299-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.012}
    
    TestTrimKeepsLines pass
    {"Time":"2026-10-25T01:37:12.982729413-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:12.988516431-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestTrimKeepsLines"}
    {"Time":"2026-10-25T01:37:12.988544199-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestTrimKeepsLines","Output":"=== RUN   TestTrimKeepsLines\n"}
    {"Time":"2026-10-25T01:37:12.988563694-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestTrimKeepsLines","Output":"--- PASS: TestTrimKeepsLines (0.00s)\n"}
    {"Time":"2026-10-25T01:37:12.988565608-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestTrimKeepsLines","Elapsed":0}
    {"Time":"2026-10-25T01:37:12.98856861-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:12.989177671-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.006s\n"}
    {"Time":"2026-10-25T01:37:12.989198437-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: null-empty (FEATURE-NullValues.md)
  intent: The empty string is reported null by csvdata.IsNull.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:13.95082594-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:13.95813958-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestIsNullEmpty"}
    {"Time":"2026-10-25T01:37:13.958164607-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestIsNullEmpty","Output":"=== RUN   TestIsNullEmpty\n"}
    {"Time":"2026-10-25T01:37:13.958186211-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestIsNullEmpty","Output":"--- PASS: TestIsNullEmpty (0.00s)\n"}
    {"Time":"2026-10-25T01:37:13.958187936-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestIsNullEmpty","Elapsed":0}
    {"Time":"2026-10-25T01:37:13.958191139-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:13.958615709-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.008s\n"}
    {"Time":"2026-10-25T01:37:13.958638948-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: null-spellings (FEATURE-NullValues.md)
  intent: NA, N/A and NULL are null in any letter case.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:14.938035447-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:14.944420871-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestIsNullSpellings"}
    {"Time":"2026-10-25T01:37:14.944444566-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestIsNullSpellings","Output":"=== RUN   TestIsNullSpellings\n"}
    {"Time":"2026-10-25T01:37:14.944464779-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestIsNullSpellings","Output":"--- PASS: TestIsNullSpellings (0.00s)\n"}
    {"Time":"2026-10-25T01:37:14.944466777-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestIsNullSpellings","Elapsed":0}
    {"Time":"2026-10-25T01:37:14.944469719-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:14.945146032-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.007s\n"}
    {"Time":"2026-10-25T01:37:14.945168184-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: null-non-null-values (FEATURE-NullValues.md)
  intent: Ordinary values such as 0, false and near-miss spellings are not null.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:15.949586583-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:15.955670056-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestIsNullNonNullValues"}
    {"Time":"2026-10-25T01:37:15.955709377-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestIsNullNonNullValues","Output":"=== RUN   TestIsNullNonNullValues\n"}
    {"Time":"2026-10-25T01:37:15.955716402-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestIsNullNonNullValues","Output":"--- PASS: TestIsNullNonNullValues (0.00s)\n"}
    {"Time":"2026-10-25T01:37:15.955718066-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestIsNullNonNullValues","Elapsed":0}
    {"Time":"2026-10-25T01:37:15.955721046-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:15.956301401-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.007s\n"}
    {"Time":"2026-10-25T01:37:15.956323365-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: null-quoted-empty (FEATURE-NullValues.md)
  intent: A quoted empty field read from a CSV file is null.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:16.890125027-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:16.895365912-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestQuotedEmptyFieldIsNull"}
    {"Time":"2026-10-25T01:37:16.895410355-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestQuotedEmptyFieldIsNull","Output":"=== RUN   TestQuotedEmptyFieldIsNull\n"}
    {"Time":"2026-10-25T01:37:16.895418949-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestQuotedEmptyFieldIsNull","Output":"--- PASS: TestQuotedEmptyFieldIsNull (0.00s)\n"}
    {"Time":"2026-10-25T01:37:16.895421055-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestQuotedEmptyFieldIsNull","Elapsed":0}
    {"Time":"2026-10-25T01:37:16.895424072-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:16.896054683-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.006s\n"}
    {"Time":"2026-10-25T01:37:16.896085699-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: null-spaces-only (FEATURE-NullValues.md)
  intent: A field of only spaces is null once trimmed.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:17.897475455-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:17.904310098-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestSpacesOnlyFieldIsNull"}
    {"Time":"2026-10-25T01:37:17.904335277-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestSpacesOnlyFieldIsNull","Output":"=== RUN   TestSpacesOnlyFieldIsNull\n"}
    {"Time":"2026-10-25T01:37:17.904364023-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestSpacesOnlyFieldIsNull","Output":"--- PASS: TestSpacesOnlyFieldIsNull (0.00s)\n"}
    {"Time":"2026-10-25T01:37:17.904366746-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestSpacesOnlyFieldIsNull","Elapsed":0}
    {"Time":"2026-10-25T01:37:17.904370283-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:17.904966175-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.007s\n"}
    {"Time":"2026-10-25T01:37:17.9050003-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: numeric-integer (FEATURE-NumericTypes.md)
  intent: A column of integer-shaped values, including negatives, is typed integer.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:18.628074915-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:18.634824436-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestInferIntegerColumn"}
    {"Time":"2026-10-25T01:37:18.634879638-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferIntegerColumn","Output":"=== RUN   TestInferIntegerColumn\n"}
    {"Time":"2026-10-25T01:37:18.63490155-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferIntegerColumn","Output":"--- PASS: TestInferIntegerColumn (0.00s)\n"}
    {"Time":"2026-10-25T01:37:18.634903966-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestInferIntegerColumn","Elapsed":0}
    {"Time":"2026-10-25T01:37:18.634906926-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:18.635281617-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:37:18.635303731-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: numeric-number (FEATURE-NumericTypes.md)
  intent: A column with a decimal value among integers is typed number.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:19.551113461-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:19.55881388-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestInferNumberColumn"}
    {"Time":"2026-10-25T01:37:19.558844847-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferNumberColumn","Output":"=== RUN   TestInferNumberColumn\n"}
    {"Time":"2026-10-25T01:37:19.55886904-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferNumberColumn","Output":"--- PASS: TestInferNumberColumn (0.00s)\n"}
    {"Time":"2026-10-25T01:37:19.558870781-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestInferNumberColumn","Elapsed":0}
    {"Time":"2026-10-25T01:37:19.558873719-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:19.559599353-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s\n"}
    {"Time":"2026-10-25T01:37:19.559625636-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: numeric-integer-preferred (FEATURE-NumericTypes.md)
  intent: Integer is chosen over number when both patterns match every value.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:20.518439997-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:20.525519925-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestIntegerPreferredOverNumber"}
    {"Time":"2026-10-25T01:37:20.525563487-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestIntegerPreferredOverNumber","Output":"=== RUN   TestIntegerPreferredOverNumber\n"}
    {"Time":"2026-10-25T01:37:20.525583584-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestIntegerPreferredOverNumber","Output":"--- PASS: TestIntegerPreferredOverNumber (0.00s)\n"}
    {"Time":"2026-10-25T01:37:20.525585879-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestIntegerPreferredOverNumber","Elapsed":0}
    {"Time":"2026-10-25T01:37:20.525589619-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:20.526214707-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s\n"}
    {"Time":"2026-10-25T01:37:20.526231255-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: numeric-nulls-ignored (FEATURE-NumericTypes.md)
  intent: Null spellings among numeric values do not change the inferred type.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:21.398018816-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:21.40413221-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestNullsIgnoredWhenInferringNumeric"}
    {"Time":"2026-10-25T01:37:21.404153863-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestNullsIgnoredWhenInferringNumeric","Output":"=== RUN   TestNullsIgnoredWhenInferringNumeric\n"}
    {"Time":"2026-10-25T01:37:21.404273026-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestNullsIgnoredWhenInferringNumeric","Output":"--- PASS: TestNullsIgnoredWhenInferringNumeric (0.00s)\n"}
    {"Time":"2026-10-25T01:37:21.40428909-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestNullsIgnoredWhenInferringNumeric","Elapsed":0}
    {"Time":"2026-10-25T01:37:21.404292679-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:21.404867776-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:37:21.404889235-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: numeric-malformed (FEATURE-NumericTypes.md)
  intent: Values such as +1, 1., .5, 1e3 and 1,5 are not numeric.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:22.348533607-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:22.354019112-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric"}
    {"Time":"2026-10-25T01:37:22.354058675-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric","Output":"=== RUN   TestMalformedNumericsAreNotNumeric\n"}
    {"Time":"2026-10-25T01:37:22.354065397-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/+1"}
    {"Time":"2026-10-25T01:37:22.354066968-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/+1","Output":"=== RUN   TestMalformedNumericsAreNotNumeric/+1\n"}
    {"Time":"2026-10-25T01:37:22.354070772-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/+1","Output":"--- PASS: TestMalformedNumericsAreNotNumeric/+1 (0.00s)\n"}
    {"Time":"2026-10-25T01:37:22.354072157-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/+1","Elapsed":0}
    {"Time":"2026-10-25T01:37:22.354074428-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/1."}
    {"Time":"2026-10-25T01:37:22.354075341-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/1.","Output":"=== RUN   TestMalformedNumericsAreNotNumeric/1.\n"}
    {"Time":"2026-10-25T01:37:22.354076653-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/1.","Output":"--- PASS: TestMalformedNumericsAreNotNumeric/1. (0.00s)\n"}
    {"Time":"2026-10-25T01:37:22.354077698-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/1.","Elapsed":0}
    {"Time":"2026-10-25T01:37:22.354078776-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/.5"}
    {"Time":"2026-10-25T01:37:22.354079606-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/.5","Output":"=== RUN   TestMalformedNumericsAreNotNumeric/.5\n"}
    {"Time":"2026-10-25T01:37:22.35408181-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/.5","Output":"--- PASS: TestMalformedNumericsAreNotNumeric/.5 (0.00s)\n"}
    {"Time":"2026-10-25T01:37:22.354082896-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/.5","Elapsed":0}
    {"Time":"2026-10-25T01:37:22.354083868-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/1e3"}
    {"Time":"2026-10-25T01:37:22.354084717-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/1e3","Output":"=== RUN   TestMalformedNumericsAreNotNumeric/1e3\n"}
    {"Time":"2026-10-25T01:37:22.354085865-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/1e3","Output":"--- PASS: TestMalformedNumericsAreNotNumeric/1e3 (0.00s)\n"}
    {"Time":"2026-10-25T01:37:22.354086928-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/1e3","Elapsed":0}
    {"Time":"2026-10-25T01:37:22.354087867-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/1,5"}
    {"Time":"2026-10-25T01:37:22.354088703-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/1,5","Output":"=== RUN   TestMalformedNumericsAreNotNumeric/1,5\n"}
    {"Time":"2026-10-25T01:37:22.35410183-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/1,5","Output":"--- PASS: TestMalformedNumericsAreNotNumeric/1,5 (0.00s)\n"}
    {"Time":"2026-10-25T01:37:22.354103001-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/1,5","Elapsed":0}
    {"Time":"2026-10-25T01:37:22.354103963-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/0x10"}
    {"Time":"2026-10-25T01:37:22.354104794-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/0x10","Output":"=== RUN   TestMalformedNumericsAreNotNumeric/0x10\n"}
    {"Time":"2026-10-25T01:37:22.354105964-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/0x10","Output":"--- PASS: TestMalformedNumericsAreNotNumeric/0x10 (0.00s)\n"}
    {"Time":"2026-10-25T01:37:22.354106973-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/0x10","Elapsed":0}
    {"Time":"2026-10-25T01:37:22.354108651-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/-"}
    {"Time":"2026-10-25T01:37:22.354109879-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/-","Output":"=== RUN   TestMalformedNumericsAreNotNumeric/-\n"}
    {"Time":"2026-10-25T01:37:22.354111062-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/-","Output":"--- PASS: TestMalformedNumericsAreNotNumeric/- (0.00s)\n"}
    {"Time":"2026-10-25T01:37:22.354112053-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric/-","Elapsed":0}
    {"Time":"2026-10-25T01:37:22.35411319-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric","Output":"--- PASS: TestMalformedNumericsAreNotNumeric (0.00s)\n"}
    {"Time":"2026-10-25T01:37:22.354114492-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMalformedNumericsAreNotNumeric","Elapsed":0}
    {"Time":"2026-10-25T01:37:22.354115527-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:22.354826808-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s\n"}
    {"Time":"2026-10-25T01:37:22.354846208-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: date-column (FEATURE-DateType.md)
  intent: A column of valid YYYY-MM-DD values is typed date.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:23.24871304-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:23.254870599-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestInferDateColumn"}
    {"Time":"2026-10-25T01:37:23.254895337-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferDateColumn","Output":"=== RUN   TestInferDateColumn\n"}
    {"Time":"2026-10-25T01:37:23.254920911-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferDateColumn","Output":"--- PASS: TestInferDateColumn (0.00s)\n"}
    {"Time":"2026-10-25T01:37:23.254923655-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestInferDateColumn","Elapsed":0}
    {"Time":"2026-10-25T01:37:23.254926499-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:23.255489455-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:37:23.255513136-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: date-invalid-calendar (FEATURE-DateType.md)
  intent: A column containing an impossible date such as 2025-02-30 is not typed date.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:24.14669276-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:24.153179423-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestInvalidCalendarDateIsNotDate"}
    {"Time":"2026-10-25T01:37:24.15321835-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInvalidCalendarDateIsNotDate","Output":"=== RUN   TestInvalidCalendarDateIsNotDate\n"}
    {"Time":"2026-10-25T01:37:24.153226674-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInvalidCalendarDateIsNotDate","Output":"--- PASS: TestInvalidCalendarDateIsNotDate (0.00s)\n"}
    {"Time":"2026-10-25T01:37:24.153228507-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestInvalidCalendarDateIsNotDate","Elapsed":0}
    {"Time":"2026-10-25T01:37:24.153231633-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:24.153920681-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:37:24.153939837-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: date-leap-day (FEATURE-DateType.md)
  intent: A leap-year February 29 is a date and a non-leap-year February 29 is not.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:25.112969271-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:25.119393019-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestLeapDayHandling"}
    {"Time":"2026-10-25T01:37:25.119431934-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestLeapDayHandling","Output":"=== RUN   TestLeapDayHandling\n"}
    {"Time":"2026-10-25T01:37:25.119462582-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestLeapDayHandling","Output":"--- PASS: TestLeapDayHandling (0.00s)\n"}
    {"Time":"2026-10-25T01:37:25.119465399-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestLeapDayHandling","Elapsed":0}
    {"Time":"2026-10-25T01:37:25.119468466-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:25.120189213-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:37:25.120214062-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: date-shape-strict (FEATURE-DateType.md)
  intent: Values with the wrong shape (unpadded, slashes, timestamps) are not dates.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:26.09447334-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:26.100616946-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict"}
    {"Time":"2026-10-25T01:37:26.100647801-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict","Output":"=== RUN   TestDateShapeIsStrict\n"}
    {"Time":"2026-10-25T01:37:26.100673601-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict/2025-2-3"}
    {"Time":"2026-10-25T01:37:26.100675287-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict/2025-2-3","Output":"=== RUN   TestDateShapeIsStrict/2025-2-3\n"}
    {"Time":"2026-10-25T01:37:26.100679162-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict/2025-2-3","Output":"--- PASS: TestDateShapeIsStrict/2025-2-3 (0.00s)\n"}
    {"Time":"2026-10-25T01:37:26.100680492-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict/2025-2-3","Elapsed":0}
    {"Time":"2026-10-25T01:37:26.100683394-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict/2025/02/03"}
    {"Time":"2026-10-25T01:37:26.100684359-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict/2025/02/03","Output":"=== RUN   TestDateShapeIsStrict/2025/02/03\n"}
    {"Time":"2026-10-25T01:37:26.100685771-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict/2025/02/03","Output":"--- PASS: TestDateShapeIsStrict/2025/02/03 (0.00s)\n"}
    {"Time":"2026-10-25T01:37:26.100686935-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict/2025/02/03","Elapsed":0}
    {"Time":"2026-10-25T01:37:26.10068809-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict/2025-02-03T10:00"}
    {"Time":"2026-10-25T01:37:26.10068912-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict/2025-02-03T10:00","Output":"=== RUN   TestDateShapeIsStrict/2025-02-03T10:00\n"}
    {"Time":"2026-10-25T01:37:26.100693168-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict/2025-02-03T10:00","Output":"--- PASS: TestDateShapeIsStrict/2025-02-03T10:00 (0.00s)\n"}
    {"Time":"2026-10-25T01:37:26.100694682-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict/2025-02-03T10:00","Elapsed":0}
    {"Time":"2026-10-25T01:37:26.100695995-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict","Output":"--- PASS: TestDateShapeIsStrict (0.00s)\n"}
    {"Time":"2026-10-25T01:37:26.100697456-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestDateShapeIsStrict","Elapsed":0}
    {"Time":"2026-10-25T01:37:26.100698537-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:26.101441807-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:37:26.10146516-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: date-nulls-ignored (FEATURE-DateType.md)
  intent: Null spellings among dates do not change the inferred date type.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:27.072619059-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:27.080664982-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestNullsIgnoredWhenInferringDate"}
    {"Time":"2026-10-25T01:37:27.08071535-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestNullsIgnoredWhenInferringDate","Output":"=== RUN   TestNullsIgnoredWhenInferringDate\n"}
    {"Time":"2026-10-25T01:37:27.080725919-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestNullsIgnoredWhenInferringDate","Output":"--- PASS: TestNullsIgnoredWhenInferringDate (0.00s)\n"}
    {"Time":"2026-10-25T01:37:27.080731274-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestNullsIgnoredWhenInferringDate","Elapsed":0}
    {"Time":"2026-10-25T01:37:27.080735756-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:27.081603671-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.009s\n"}
    {"Time":"2026-10-25T01:37:27.081628714-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: type-boolean-mixed-case (FEATURE-BooleanTextTypes.md)
  intent: Values true and false in any letter case form a boolean column.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:27.961436966-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:27.969290067-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestInferBooleanColumnMixedCase"}
    {"Time":"2026-10-25T01:37:27.969339109-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferBooleanColumnMixedCase","Output":"=== RUN   TestInferBooleanColumnMixedCase\n"}
    {"Time":"2026-10-25T01:37:27.969349981-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferBooleanColumnMixedCase","Output":"--- PASS: TestInferBooleanColumnMixedCase (0.00s)\n"}
    {"Time":"2026-10-25T01:37:27.969352369-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestInferBooleanColumnMixedCase","Elapsed":0}
    {"Time":"2026-10-25T01:37:27.969356272-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:27.970035119-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s\n"}
    {"Time":"2026-10-25T01:37:27.970057508-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: type-text-fallback (FEATURE-BooleanTextTypes.md)
  intent: A column matching no earlier type is text.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:28.860232748-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:28.866634696-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestInferTextFallback"}
    {"Time":"2026-10-25T01:37:28.866671901-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferTextFallback","Output":"=== RUN   TestInferTextFallback\n"}
    {"Time":"2026-10-25T01:37:28.866697569-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferTextFallback","Output":"--- PASS: TestInferTextFallback (0.00s)\n"}
    {"Time":"2026-10-25T01:37:28.866699816-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestInferTextFallback","Elapsed":0}
    {"Time":"2026-10-25T01:37:28.866702395-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:28.867633698-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:37:28.867653402-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: type-empty-all-null (FEATURE-BooleanTextTypes.md)
  intent: A column whose values are all null spellings is typed empty.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:29.799626716-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:29.806464127-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestInferEmptyWhenAllNull"}
    {"Time":"2026-10-25T01:37:29.806512165-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferEmptyWhenAllNull","Output":"=== RUN   TestInferEmptyWhenAllNull\n"}
    {"Time":"2026-10-25T01:37:29.806519493-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferEmptyWhenAllNull","Output":"--- PASS: TestInferEmptyWhenAllNull (0.00s)\n"}
    {"Time":"2026-10-25T01:37:29.806522764-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestInferEmptyWhenAllNull","Elapsed":0}
    {"Time":"2026-10-25T01:37:29.806526098-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:29.808124865-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s\n"}
    {"Time":"2026-10-25T01:37:29.808147918-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: type-empty-no-rows (FEATURE-BooleanTextTypes.md)
  intent: A column with zero data rows is typed empty.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:30.827000229-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:30.834110887-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestInferEmptyWhenNoRows"}
    {"Time":"2026-10-25T01:37:30.83415495-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferEmptyWhenNoRows","Output":"=== RUN   TestInferEmptyWhenNoRows\n"}
    {"Time":"2026-10-25T01:37:30.834162273-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestInferEmptyWhenNoRows","Output":"--- PASS: TestInferEmptyWhenNoRows (0.00s)\n"}
    {"Time":"2026-10-25T01:37:30.834163946-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestInferEmptyWhenNoRows","Elapsed":0}
    {"Time":"2026-10-25T01:37:30.834166705-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:30.834798145-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s\n"}
    {"Time":"2026-10-25T01:37:30.83482003-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: type-boolean-lookalikes (FEATURE-BooleanTextTypes.md)
  intent: Values such as yes, 1 or t do not make a boolean column.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:31.802748854-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:31.809149841-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText"}
    {"Time":"2026-10-25T01:37:31.809202001-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText","Output":"=== RUN   TestBooleanLookalikesAreText\n"}
    {"Time":"2026-10-25T01:37:31.809216993-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/yes"}
    {"Time":"2026-10-25T01:37:31.809218825-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/yes","Output":"=== RUN   TestBooleanLookalikesAreText/yes\n"}
    {"Time":"2026-10-25T01:37:31.809222404-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/yes","Output":"--- PASS: TestBooleanLookalikesAreText/yes (0.00s)\n"}
    {"Time":"2026-10-25T01:37:31.809223714-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/yes","Elapsed":0}
    {"Time":"2026-10-25T01:37:31.809226776-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/no"}
    {"Time":"2026-10-25T01:37:31.809227729-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/no","Output":"=== RUN   TestBooleanLookalikesAreText/no\n"}
    {"Time":"2026-10-25T01:37:31.809229073-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/no","Output":"--- PASS: TestBooleanLookalikesAreText/no (0.00s)\n"}
    {"Time":"2026-10-25T01:37:31.809230166-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/no","Elapsed":0}
    {"Time":"2026-10-25T01:37:31.809231277-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/t"}
    {"Time":"2026-10-25T01:37:31.809232161-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/t","Output":"=== RUN   TestBooleanLookalikesAreText/t\n"}
    {"Time":"2026-10-25T01:37:31.809233381-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/t","Output":"--- PASS: TestBooleanLookalikesAreText/t (0.00s)\n"}
    {"Time":"2026-10-25T01:37:31.809234395-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/t","Elapsed":0}
    {"Time":"2026-10-25T01:37:31.809235419-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/f"}
    {"Time":"2026-10-25T01:37:31.809236298-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/f","Output":"=== RUN   TestBooleanLookalikesAreText/f\n"}
    {"Time":"2026-10-25T01:37:31.809237416-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/f","Output":"--- PASS: TestBooleanLookalikesAreText/f (0.00s)\n"}
    {"Time":"2026-10-25T01:37:31.80923941-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/f","Elapsed":0}
    {"Time":"2026-10-25T01:37:31.809240538-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/y"}
    {"Time":"2026-10-25T01:37:31.809241437-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/y","Output":"=== RUN   TestBooleanLookalikesAreText/y\n"}
    {"Time":"2026-10-25T01:37:31.809242589-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/y","Output":"--- PASS: TestBooleanLookalikesAreText/y (0.00s)\n"}
    {"Time":"2026-10-25T01:37:31.809243682-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/y","Elapsed":0}
    {"Time":"2026-10-25T01:37:31.80924466-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/n"}
    {"Time":"2026-10-25T01:37:31.809245544-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/n","Output":"=== RUN   TestBooleanLookalikesAreText/n\n"}
    {"Time":"2026-10-25T01:37:31.809246679-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/n","Output":"--- PASS: TestBooleanLookalikesAreText/n (0.00s)\n"}
    {"Time":"2026-10-25T01:37:31.809251282-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/n","Elapsed":0}
    {"Time":"2026-10-25T01:37:31.809252401-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/on"}
    {"Time":"2026-10-25T01:37:31.809253347-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/on","Output":"=== RUN   TestBooleanLookalikesAreText/on\n"}
    {"Time":"2026-10-25T01:37:31.809256709-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/on","Output":"--- PASS: TestBooleanLookalikesAreText/on (0.00s)\n"}
    {"Time":"2026-10-25T01:37:31.809268491-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/on","Elapsed":0}
    {"Time":"2026-10-25T01:37:31.809279062-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/off"}
    {"Time":"2026-10-25T01:37:31.809280635-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/off","Output":"=== RUN   TestBooleanLookalikesAreText/off\n"}
    {"Time":"2026-10-25T01:37:31.809282953-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/off","Output":"--- PASS: TestBooleanLookalikesAreText/off (0.00s)\n"}
    {"Time":"2026-10-25T01:37:31.809284241-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText/off","Elapsed":0}
    {"Time":"2026-10-25T01:37:31.809285616-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText","Output":"--- PASS: TestBooleanLookalikesAreText (0.00s)\n"}
    {"Time":"2026-10-25T01:37:31.809287069-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestBooleanLookalikesAreText","Elapsed":0}
    {"Time":"2026-10-25T01:37:31.809290079-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:31.809948943-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:37:31.80996994-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: type-order-precedence (FEATURE-BooleanTextTypes.md)
  intent: Earlier types win over later ones when several would match.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:32.629015121-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:37:32.636833246-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestTypeOrderPrecedence"}
    {"Time":"2026-10-25T01:37:32.63687074-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestTypeOrderPrecedence","Output":"=== RUN   TestTypeOrderPrecedence\n"}
    {"Time":"2026-10-25T01:37:32.636878182-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestTypeOrderPrecedence","Output":"--- PASS: TestTypeOrderPrecedence (0.00s)\n"}
    {"Time":"2026-10-25T01:37:32.636879886-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestTypeOrderPrecedence","Elapsed":0}
    {"Time":"2026-10-25T01:37:32.636883168-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:32.637366668-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s\n"}
    {"Time":"2026-10-25T01:37:32.63738575-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: args-file-only (FEATURE-Arguments.md)
  intent: A lone file argument parses.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:33.688349533-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-25T01:37:33.694552143-04:00","Action":"run","Package":"csvstat/internal/cli","Test":"TestParseArgsFileOnly"}
    {"Time":"2026-10-25T01:37:33.694576045-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsFileOnly","Output":"=== RUN   TestParseArgsFileOnly\n"}
    {"Time":"2026-10-25T01:37:33.694598775-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsFileOnly","Output":"--- PASS: TestParseArgsFileOnly (0.00s)\n"}
    {"Time":"2026-10-25T01:37:33.694600692-04:00","Action":"pass","Package":"csvstat/internal/cli","Test":"TestParseArgsFileOnly","Elapsed":0}
    {"Time":"2026-10-25T01:37:33.694603716-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:33.695274694-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"ok  \tcsvstat/internal/cli\t0.007s\n"}
    {"Time":"2026-10-25T01:37:33.695300587-04:00","Action":"pass","Package":"csvstat/internal/cli","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: args-column (FEATURE-Arguments.md)
  intent: The --column option and its value parse before the file.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:34.692365441-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-25T01:37:34.699595269-04:00","Action":"run","Package":"csvstat/internal/cli","Test":"TestParseArgsColumn"}
    {"Time":"2026-10-25T01:37:34.699618967-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsColumn","Output":"=== RUN   TestParseArgsColumn\n"}
    {"Time":"2026-10-25T01:37:34.6996479-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsColumn","Output":"--- PASS: TestParseArgsColumn (0.00s)\n"}
    {"Time":"2026-10-25T01:37:34.69965014-04:00","Action":"pass","Package":"csvstat/internal/cli","Test":"TestParseArgsColumn","Elapsed":0}
    {"Time":"2026-10-25T01:37:34.699653365-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:34.700216905-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"ok  \tcsvstat/internal/cli\t0.008s\n"}
    {"Time":"2026-10-25T01:37:34.700237863-04:00","Action":"pass","Package":"csvstat/internal/cli","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: args-freq-top (FEATURE-Arguments.md)
  intent: The --freq and --top options and their values parse before the file.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:35.690587817-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-25T01:37:35.696270222-04:00","Action":"run","Package":"csvstat/internal/cli","Test":"TestParseArgsFreqTop"}
    {"Time":"2026-10-25T01:37:35.696294494-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsFreqTop","Output":"=== RUN   TestParseArgsFreqTop\n"}
    {"Time":"2026-10-25T01:37:35.696315337-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsFreqTop","Output":"--- PASS: TestParseArgsFreqTop (0.00s)\n"}
    {"Time":"2026-10-25T01:37:35.696318175-04:00","Action":"pass","Package":"csvstat/internal/cli","Test":"TestParseArgsFreqTop","Elapsed":0}
    {"Time":"2026-10-25T01:37:35.696321068-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:35.69713717-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"ok  \tcsvstat/internal/cli\t0.006s\n"}
    {"Time":"2026-10-25T01:37:35.697149845-04:00","Action":"pass","Package":"csvstat/internal/cli","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: args-structural-errors (FEATURE-Arguments.md)
  intent: Missing file, extra file, unknown option, missing value, and options after the file are parse errors.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    TestParseArgsMissingFile pass
    {"Time":"2026-10-25T01:37:36.596695296-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-25T01:37:36.602794748-04:00","Action":"run","Package":"csvstat/internal/cli","Test":"TestParseArgsMissingFile"}
    {"Time":"2026-10-25T01:37:36.602828673-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsMissingFile","Output":"=== RUN   TestParseArgsMissingFile\n"}
    {"Time":"2026-10-25T01:37:36.602851472-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsMissingFile","Output":"--- PASS: TestParseArgsMissingFile (0.00s)\n"}
    {"Time":"2026-10-25T01:37:36.602853135-04:00","Action":"pass","Package":"csvstat/internal/cli","Test":"TestParseArgsMissingFile","Elapsed":0}
    {"Time":"2026-10-25T01:37:36.602855825-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:36.603444581-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"ok  \tcsvstat/internal/cli\t0.007s\n"}
    {"Time":"2026-10-25T01:37:36.603467871-04:00","Action":"pass","Package":"csvstat/internal/cli","Elapsed":0.007}
    
    TestParseArgsExtraFile pass
    {"Time":"2026-10-25T01:37:37.317931998-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-25T01:37:37.3236519-04:00","Action":"run","Package":"csvstat/internal/cli","Test":"TestParseArgsExtraFile"}
    {"Time":"2026-10-25T01:37:37.323680321-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsExtraFile","Output":"=== RUN   TestParseArgsExtraFile\n"}
    {"Time":"2026-10-25T01:37:37.323702275-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsExtraFile","Output":"--- PASS: TestParseArgsExtraFile (0.00s)\n"}
    {"Time":"2026-10-25T01:37:37.323705275-04:00","Action":"pass","Package":"csvstat/internal/cli","Test":"TestParseArgsExtraFile","Elapsed":0}
    {"Time":"2026-10-25T01:37:37.323708127-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:37.324254414-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"ok  \tcsvstat/internal/cli\t0.006s\n"}
    {"Time":"2026-10-25T01:37:37.324275085-04:00","Action":"pass","Package":"csvstat/internal/cli","Elapsed":0.006}
    
    TestParseArgsUnknownOption pass
    {"Time":"2026-10-25T01:37:38.008943047-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-25T01:37:38.01501789-04:00","Action":"run","Package":"csvstat/internal/cli","Test":"TestParseArgsUnknownOption"}
    {"Time":"2026-10-25T01:37:38.015051956-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsUnknownOption","Output":"=== RUN   TestParseArgsUnknownOption\n"}
    {"Time":"2026-10-25T01:37:38.015072622-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsUnknownOption","Output":"--- PASS: TestParseArgsUnknownOption (0.00s)\n"}
    {"Time":"2026-10-25T01:37:38.015074361-04:00","Action":"pass","Package":"csvstat/internal/cli","Test":"TestParseArgsUnknownOption","Elapsed":0}
    {"Time":"2026-10-25T01:37:38.015076918-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:38.015775585-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"ok  \tcsvstat/internal/cli\t0.007s\n"}
    {"Time":"2026-10-25T01:37:38.015823147-04:00","Action":"pass","Package":"csvstat/internal/cli","Elapsed":0.007}
    
    TestParseArgsMissingValue pass
    {"Time":"2026-10-25T01:37:38.695379145-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-25T01:37:38.702380346-04:00","Action":"run","Package":"csvstat/internal/cli","Test":"TestParseArgsMissingValue"}
    {"Time":"2026-10-25T01:37:38.702407943-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsMissingValue","Output":"=== RUN   TestParseArgsMissingValue\n"}
    {"Time":"2026-10-25T01:37:38.702425449-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsMissingValue","Output":"--- PASS: TestParseArgsMissingValue (0.00s)\n"}
    {"Time":"2026-10-25T01:37:38.702427106-04:00","Action":"pass","Package":"csvstat/internal/cli","Test":"TestParseArgsMissingValue","Elapsed":0}
    {"Time":"2026-10-25T01:37:38.702429509-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:38.702964928-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"ok  \tcsvstat/internal/cli\t0.007s\n"}
    {"Time":"2026-10-25T01:37:38.702994303-04:00","Action":"pass","Package":"csvstat/internal/cli","Elapsed":0.008}
    
    TestParseArgsOptionsAfterFile pass
    {"Time":"2026-10-25T01:37:39.387014672-04:00","Action":"start","Package":"csvstat/internal/cli"}
    {"Time":"2026-10-25T01:37:39.393015184-04:00","Action":"run","Package":"csvstat/internal/cli","Test":"TestParseArgsOptionsAfterFile"}
    {"Time":"2026-10-25T01:37:39.393045473-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsOptionsAfterFile","Output":"=== RUN   TestParseArgsOptionsAfterFile\n"}
    {"Time":"2026-10-25T01:37:39.393064664-04:00","Action":"output","Package":"csvstat/internal/cli","Test":"TestParseArgsOptionsAfterFile","Output":"--- PASS: TestParseArgsOptionsAfterFile (0.00s)\n"}
    {"Time":"2026-10-25T01:37:39.393068795-04:00","Action":"pass","Package":"csvstat/internal/cli","Test":"TestParseArgsOptionsAfterFile","Elapsed":0}
    {"Time":"2026-10-25T01:37:39.393071556-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:39.39377624-04:00","Action":"output","Package":"csvstat/internal/cli","Output":"ok  \tcsvstat/internal/cli\t0.006s\n"}
    {"Time":"2026-10-25T01:37:39.393801745-04:00","Action":"pass","Package":"csvstat/internal/cli","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: args-valid-forms-not-usage-errors (FEATURE-Arguments.md)
  intent: The valid command forms are accepted by the executable and never exit with the usage-error code.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ['/tmp/tmp8syan6ue/data.csv'] 0  
    ['--column', 'a', '/tmp/tmp8syan6ue/data.csv'] 0  
    ['--freq', 'a', '/tmp/tmp8syan6ue/data.csv'] 0  
    ['--freq', 'a', '--top', '3', '/tmp/tmp8syan6ue/data.csv'] 0
- PASS: args-options-after-file-exit (FEATURE-Arguments.md)
  intent: An option placed after the file is rejected by the executable with the usage-error exit status and no standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    2  csvstat: unexpected argument "--column" after file or extra file
- PASS: reject-unit-empty-and-blank (FEATURE-InputRejection.md)
  intent: Empty files and files of only blank lines are rejected by the reader.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:44.968829031-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:44.974514926-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestRejectEmptyFile"}
    {"Time":"2026-10-25T01:37:44.974548256-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestRejectEmptyFile","Output":"=== RUN   TestRejectEmptyFile\n"}
    {"Time":"2026-10-25T01:37:44.974582122-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestRejectEmptyFile","Output":"--- PASS: TestRejectEmptyFile (0.00s)\n"}
    {"Time":"2026-10-25T01:37:44.974585047-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestRejectEmptyFile","Elapsed":0}
    {"Time":"2026-10-25T01:37:44.974589609-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:44.97499702-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.006s\n"}
    {"Time":"2026-10-25T01:37:44.97501915-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.006}
    
    
    {"Time":"2026-10-25T01:37:45.667039255-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:45.673780391-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestRejectBlankLinesOnly"}
    {"Time":"2026-10-25T01:37:45.673805566-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestRejectBlankLinesOnly","Output":"=== RUN   TestRejectBlankLinesOnly\n"}
    {"Time":"2026-10-25T01:37:45.673834967-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestRejectBlankLinesOnly","Output":"--- PASS: TestRejectBlankLinesOnly (0.00s)\n"}
    {"Time":"2026-10-25T01:37:45.67383664-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestRejectBlankLinesOnly","Elapsed":0}
    {"Time":"2026-10-25T01:37:45.673839901-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:45.674409852-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.007s\n"}
    {"Time":"2026-10-25T01:37:45.67444698-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: reject-unit-headers (FEATURE-InputRejection.md)
  intent: Empty and repeated header names are rejected by the reader.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:46.566293776-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:46.572102788-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestRejectEmptyHeaderName"}
    {"Time":"2026-10-25T01:37:46.572142884-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestRejectEmptyHeaderName","Output":"=== RUN   TestRejectEmptyHeaderName\n"}
    {"Time":"2026-10-25T01:37:46.572160198-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestRejectEmptyHeaderName","Output":"--- PASS: TestRejectEmptyHeaderName (0.00s)\n"}
    {"Time":"2026-10-25T01:37:46.572163018-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestRejectEmptyHeaderName","Elapsed":0}
    {"Time":"2026-10-25T01:37:46.572165978-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:46.57277421-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.006s\n"}
    {"Time":"2026-10-25T01:37:46.572794093-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.007}
    
    
    {"Time":"2026-10-25T01:37:47.385332249-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:47.394076055-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestRejectRepeatedHeaderName"}
    {"Time":"2026-10-25T01:37:47.394101073-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestRejectRepeatedHeaderName","Output":"=== RUN   TestRejectRepeatedHeaderName\n"}
    {"Time":"2026-10-25T01:37:47.394120614-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestRejectRepeatedHeaderName","Output":"--- PASS: TestRejectRepeatedHeaderName (0.00s)\n"}
    {"Time":"2026-10-25T01:37:47.394122637-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestRejectRepeatedHeaderName","Elapsed":0}
    {"Time":"2026-10-25T01:37:47.394125492-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:47.395444494-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.010s\n"}
    {"Time":"2026-10-25T01:37:47.395469005-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.01}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: reject-unit-ragged-and-valid (FEATURE-InputRejection.md)
  intent: Ragged records are rejected while header-only and newline-terminated files are accepted.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:37:48.340257992-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:48.346056782-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestRejectRaggedRecord"}
    {"Time":"2026-10-25T01:37:48.346124866-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestRejectRaggedRecord","Output":"=== RUN   TestRejectRaggedRecord\n"}
    {"Time":"2026-10-25T01:37:48.346132117-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestRejectRaggedRecord","Output":"--- PASS: TestRejectRaggedRecord (0.00s)\n"}
    {"Time":"2026-10-25T01:37:48.346133678-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestRejectRaggedRecord","Elapsed":0}
    {"Time":"2026-10-25T01:37:48.346136325-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:48.34650459-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.006s\n"}
    {"Time":"2026-10-25T01:37:48.346548626-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.006}
    
    
    {"Time":"2026-10-25T01:37:49.054273419-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:49.060512007-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestHeaderOnlyFileAccepted"}
    {"Time":"2026-10-25T01:37:49.06054568-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestHeaderOnlyFileAccepted","Output":"=== RUN   TestHeaderOnlyFileAccepted\n"}
    {"Time":"2026-10-25T01:37:49.060569919-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestHeaderOnlyFileAccepted","Output":"--- PASS: TestHeaderOnlyFileAccepted (0.00s)\n"}
    {"Time":"2026-10-25T01:37:49.060573596-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestHeaderOnlyFileAccepted","Elapsed":0}
    {"Time":"2026-10-25T01:37:49.06057652-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:49.061346751-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.007s\n"}
    {"Time":"2026-10-25T01:37:49.061376763-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.007}
    
    
    {"Time":"2026-10-25T01:37:49.759591876-04:00","Action":"start","Package":"csvstat/internal/csvdata"}
    {"Time":"2026-10-25T01:37:49.765170818-04:00","Action":"run","Package":"csvstat/internal/csvdata","Test":"TestTrailingNewlineAddsNoRecord"}
    {"Time":"2026-10-25T01:37:49.765211105-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestTrailingNewlineAddsNoRecord","Output":"=== RUN   TestTrailingNewlineAddsNoRecord\n"}
    {"Time":"2026-10-25T01:37:49.765226715-04:00","Action":"output","Package":"csvstat/internal/csvdata","Test":"TestTrailingNewlineAddsNoRecord","Output":"--- PASS: TestTrailingNewlineAddsNoRecord (0.00s)\n"}
    {"Time":"2026-10-25T01:37:49.765229069-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Test":"TestTrailingNewlineAddsNoRecord","Elapsed":0}
    {"Time":"2026-10-25T01:37:49.765231848-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"PASS\n"}
    {"Time":"2026-10-25T01:37:49.765948301-04:00","Action":"output","Package":"csvstat/internal/csvdata","Output":"ok  \tcsvstat/internal/csvdata\t0.006s\n"}
    {"Time":"2026-10-25T01:37:49.765970308-04:00","Action":"pass","Package":"csvstat/internal/csvdata","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: reject-cli-ragged (FEATURE-InputRejection.md)
  intent: The executable exits 3 with empty standard output for a ragged record.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpu0crkyhl/ragged.csv: line 3: record has 3 fields, header has 2
- PASS: reject-cli-empty-file (FEATURE-InputRejection.md)
  intent: The executable exits 3 with empty standard output for a zero-byte file.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmp69unls8r/empty.csv: no header record: input is empty or blank
- PASS: reject-cli-blank-lines-only (FEATURE-InputRejection.md)
  intent: The executable exits 3 with empty standard output for a file holding only blank lines.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmp0dcdqex0/blank.csv: no header record: input is empty or blank
- PASS: reject-cli-repeated-header (FEATURE-InputRejection.md)
  intent: The executable exits 3 with empty standard output when a header name repeats.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpr9p9yy_o/dup.csv: header name "id" is repeated
- PASS: reject-cli-empty-header-name (FEATURE-InputRejection.md)
  intent: The executable exits 3 with empty standard output when a header name is empty.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpf2csyqaw/emptyname.csv: header column 2 is empty

## Build summary
The build directory has the scaffold (`go.mod`, `cmd/csvstat/main.go`, a stub `cli.Run`, package docs) and no implementation yet. Block 2 covers all nine stories, so I'll build the full set: the `csvdata`, `profile`, and `cli` packages, plus their tests named as the specs require.Implementation is in place. Next, the tests, named as the specs require. Each one acts on the code and reads back the returned value, not the output text.Now I'll format, then run the project's declared gate (`bin/test.sh`: vet, then test).The one failure is a wrong expectation in my test, not a code bug. `20250203` is all digits, so it matches `integer` before `date` is ever considered, as the type order requires. The spec only says it is not a date, so I'll assert `integer` for that value. The `cli_test.go` change is just gofmt realignment of my own file.The gate is green: vet passes, and 58 tests pass with 0 failures. The static rules hold: no `run.sh` in application code, `os.Exit` only in `main`, and all imports are stdlib or module-local. Last, I'll run the built binary against the staged data files from the shell to confirm the exit codes the acceptance checks expect.Block 2 is implemented and `sh bin/test.sh` passes: vet is clean and all 58 tests pass.

Valid files exit 0 and print nothing, because profile output belongs to later stories. Rejected input exits 3, usage errors exit 2, and an unreadable file exits 1. I ran the built binary on the staged CSVs: the five valid files exit 0, and `ragged.csv` exits 3 with `line 4` in the diagnostic. I did not run `sources/run.sh` or the Python acceptance criteria, since the `drydock_go` helper isn't present here. The Go tests carry the same names and assertions.

Choices I made where the specs left room:
- **Read vs. Validate:** the reader returns no error for empty input, per the CSV-reading spec. Rejection lives in a separate `csvdata.Validate`, which the cli calls after `Read` and `Trim`. Header-level rejections report line 0, because the table does not record the header's line.
- **Option presence:** `Options` has `ColumnSet` and `FreqSet` alongside `TopSet`, so an empty `--column ""` is distinguishable from an absent option.
- **Exit mapping:** a CSV syntax error maps to 3 as rejected input. Any other read failure maps to 1.
- **Number parsing:** `--top` is parsed as base-10 only, so `0x10` is rejected.

FILES CHANGED:
- internal/cli/cli.go
- internal/cli/args.go
- internal/cli/args_test.go
- internal/cli/cli_test.go
- internal/csvdata/read.go
- internal/csvdata/read_test.go
- internal/csvdata/trim.go
- internal/csvdata/trim_test.go
- internal/csvdata/null.go
- internal/csvdata/null_test.go
- internal/csvdata/validate.go
- internal/csvdata/validate_test.go
- internal/profile/infer.go
- internal/profile/infer_test.go

SUMMARY:
- `cmd/csvstat/main.go` was already correct and is unchanged. `cli.Run` now parses args, opens the file, reads, trims, and validates, mapping outcomes to exit codes 0/1/2/3.
- `csvdata` provides `Read`, `Trim`, `IsNull`, `Validate`, and `ErrRejected`, with a `RejectedError` that carries the line number for ragged records.
- `profile.InferType` returns empty, integer, number, date, boolean, or text in that order.
- `cli.ParseArgs` uses a local `flag.FlagSet` with `ContinueOnError`, rejects repeated options, and stops at the first positional argument.
- Static gates checked by hand: no `run.sh` reference in application code, `os.Exit` only in `main`, and stdlib or module-local imports only.
- Test counts: 58 passing across `cli`, `csvdata`, and `profile`, 0 failing.

<blueprint-decisions>
[{"spec":"FEATURE-CsvReading.md","severity":"Low","subject":"Where empty-input rejection happens","decision":"Read returns a nil Header with no error for empty or blank-only input, as the spec requires. Rejection is a separate csvdata.Validate, called by cli after Read and Trim, which returns ErrRejected. Reading and rejecting stay separate, and the input-rejection story still holds. Is that acceptable?"}]
</blueprint-decisions>

RESULT: SUCCESS

FILES CHANGED:
- internal/cli/cli.go
- internal/cli/args.go
- internal/cli/args_test.go
- internal/cli/cli_test.go
- internal/csvdata/read.go
- internal/csvdata/read_test.go
- internal/csvdata/trim.go
- internal/csvdata/trim_test.go
- internal/csvdata/null.go
- internal/csvdata/null_test.go
- internal/csvdata/validate.go
- internal/csvdata/validate_test.go
- internal/profile/infer.go
- internal/profile/infer_test.go

SUMMARY:
Implemented all nine Block 2 stories. `csvdata` covers reading, trimming, null recognition, and rejection. `profile.InferType` covers integer, number, date, boolean, text, and empty. `cli` covers argument parsing and the Run pipeline with exit codes 0/1/2/3. `sh bin/test.sh` passes: vet is clean, and 58 tests pass with 0 failures. The built binary exits 0 on the valid staged CSVs, 3 on `ragged.csv` with empty stdout, and 2 on options after the file. The Python acceptance criteria were not run, because the `drydock_go` helper is not present in this environment.

BLOCKERS:
- None.
