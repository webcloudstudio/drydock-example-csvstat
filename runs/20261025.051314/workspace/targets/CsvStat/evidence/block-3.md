# Evidence: Block 3 · Service (block-3)

- block type: block
- date: 2026-10-25
- resulting state: closed/verified
- story points (combined assembled cost): 19507
- execution id: 20261025.054012.840Z-d903a989

## Stories built
- Report rows, nulls, and distinct counts per column. (row-counts) [story]
- Compute exact sum and mean for integer and number columns. (sum-mean) [story]
- Compute the median for integer and number columns. (median) [story]
- Compute min and max by column type. (min-max) [story]
- Format decimal results to two places, rounded half away from zero once. (decimal-format) [story]
- Count and order distinct non-null values of a column. (freq-count) [story]
- Limit frequency rows with --top, default 5. (freq-limit) [story]
- Write the column profile as RFC 4180 CSV for the default invocation. (profile-output) [story]

## Acceptance tooling authorization
- FEATURE-RowCounts.md#rowcounts-rows-nulls-distinct: executable=go; scope=test; authorization=existing Target environment
- FEATURE-RowCounts.md#rowcounts-header-only: executable=go; scope=test; authorization=existing Target environment
- FEATURE-RowCounts.md#rowcounts-distinct-exact: executable=go; scope=test; authorization=existing Target environment
- FEATURE-RowCounts.md#rowcounts-all-null: executable=go; scope=test; authorization=existing Target environment
- FEATURE-SumMean.md#summean-integer-sum: executable=go; scope=test; authorization=existing Target environment
- FEATURE-SumMean.md#summean-number-no-drift: executable=go; scope=test; authorization=existing Target environment
- FEATURE-SumMean.md#summean-mean-exact: executable=go; scope=test; authorization=existing Target environment
- FEATURE-SumMean.md#summean-ignores-nulls: executable=go; scope=test; authorization=existing Target environment
- FEATURE-SumMean.md#summean-non-numeric: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Median.md#median-odd: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Median.md#median-even: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Median.md#median-numeric-order: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Median.md#median-ignores-nulls: executable=go; scope=test; authorization=existing Target environment
- FEATURE-Median.md#median-non-numeric: executable=go; scope=test; authorization=existing Target environment
- FEATURE-MinMax.md#minmax-integer: executable=go; scope=test; authorization=existing Target environment
- FEATURE-MinMax.md#minmax-number: executable=go; scope=test; authorization=existing Target environment
- FEATURE-MinMax.md#minmax-date: executable=go; scope=test; authorization=existing Target environment
- FEATURE-MinMax.md#minmax-boolean: executable=go; scope=test; authorization=existing Target environment
- FEATURE-MinMax.md#minmax-text: executable=go; scope=test; authorization=existing Target environment
- FEATURE-MinMax.md#minmax-empty: executable=go; scope=test; authorization=existing Target environment
- FEATURE-MinMax.md#minmax-single-value: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DecimalFormat.md#decfmt-two-places: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DecimalFormat.md#decfmt-half-away: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DecimalFormat.md#decfmt-no-negative-zero: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DecimalFormat.md#decfmt-no-separator: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DecimalFormat.md#decfmt-large-exact: executable=go; scope=test; authorization=existing Target environment
- FEATURE-DecimalFormat.md#decfmt-integer: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqCount.md#freq-count-ordering: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqCount.md#freq-count-nulls-and-trim: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqCount.md#freq-count-all-null: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqCount.md#freq-count-header-only: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqCount.md#freq-count-quoted-value: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqCount.md#freq-count-rejected-input: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqLimit.md#freq-limit-default-five: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqLimit.md#freq-limit-explicit-top: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqLimit.md#freq-limit-larger-than-distinct: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqLimit.md#freq-limit-all-null-column: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ProfileOutput.md#profile-output-columns: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ProfileOutput.md#profile-output-quoting: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ProfileOutput.md#profile-output-line-endings: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ProfileOutput.md#profile-output-header-only: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ProfileOutput.md#profile-output-rounding: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ProfileOutput.md#profile-output-rejected-input: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 963)
- implements: FEATURE-RowCounts.md (SP 1422)
- context: ARCHITECTURE_compact.md (SP 282)
- implements: FEATURE-SumMean.md (SP 1518)
- implements: FEATURE-Median.md (SP 1404)
- implements: FEATURE-MinMax.md (SP 2146)
- implements: FEATURE-DecimalFormat.md (SP 1933)
- implements: FEATURE-FreqCount.md (SP 2301)
- implements: FEATURE-FreqLimit.md (SP 1768)
- implements: FEATURE-ProfileOutput.md (SP 2928)

## Build directory changes
- internal/cli/cli.go
- internal/cli/freq.go
- internal/cli/freq_test.go
- internal/cli/profile.go
- internal/cli/profile_test.go
- internal/freq/count.go
- internal/freq/count_test.go
- internal/freq/top.go
- internal/freq/top_test.go
- internal/profile/counts.go
- internal/profile/counts_test.go
- internal/profile/format.go
- internal/profile/format_test.go
- internal/profile/median.go
- internal/profile/median_test.go
- internal/profile/minmax.go
- internal/profile/minmax_test.go
- internal/profile/numeric.go
- internal/profile/sum.go
- internal/profile/sum_test.go

## Pre-build acceptance observation
- UNVERIFIED: rowcounts-rows-nulls-distinct (FEATURE-RowCounts.md)
  intent: Counts reports rows, nulls, and distinct for a mixed column.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:22.497744273-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:22.503957301-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:22.504010218-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:22.504643942-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:22.504660988-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:22.497744273-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:22.503957301-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:22.504010218-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "rowcounts-rows-nulls-distinct.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:22.497744273-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:22.503957301-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:22.504010218-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:22.504643942-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:22.504660988-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}\n', '')
- UNVERIFIED: rowcounts-header-only (FEATURE-RowCounts.md)
  intent: A column with no data records reports zero for every count.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:23.331319757-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:23.338557493-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:23.338611362-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:23.339751038-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:23.339776078-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:23.331319757-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:23.338557493-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:23.338611362-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "rowcounts-header-only.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:23.331319757-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:23.338557493-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:23.338611362-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:23.339751038-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.008s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:23.339776078-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}\n', '')
- UNVERIFIED: rowcounts-distinct-exact (FEATURE-RowCounts.md)
  intent: Distinct compares exact strings and counts repeats once.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:24.178912485-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:24.185030291-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:24.185075656-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:24.185651208-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:24.18567589-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:24.178912485-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:24.185030291-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:24.185075656-04:00","Action":"output","Package":"csvstat/… (+354 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "rowcounts-distinct-exact.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:24.178912485-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:24.185030291-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:24.185075656-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:24.185651208-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:24.18567589-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}\n', '')
- UNVERIFIED: rowcounts-all-null (FEATURE-RowCounts.md)
  intent: An all-null column has zero distinct values and every row null.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:25.072258384-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:25.078523096-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:25.078568086-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:25.079136642-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:25.079160958-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:25.072258384-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:25.078523096-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:25.078568086-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "rowcounts-all-null.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:25.072258384-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:25.078523096-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:25.078568086-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:25.079136642-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:25.079160958-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}\n', '')
- UNVERIFIED: summean-integer-sum (FEATURE-SumMean.md)
  intent: Integer sums are exact, including negative values.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:25.969600195-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:25.975865847-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:25.975917201-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:25.976599048-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:25.976625312-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:25.969600195-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:25.975865847-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:25.975917201-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summean-integer-sum.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:25.969600195-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:25.975865847-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:25.975917201-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:25.976599048-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:25.976625312-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}\n', '')
- UNVERIFIED: summean-number-no-drift (FEATURE-SumMean.md)
  intent: Decimal sums carry no floating-point drift.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:26.839094326-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:26.845379616-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:26.845411837-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:26.846083044-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:26.846108438-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:26.839094326-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:26.845379616-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:26.845411837-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summean-number-no-drift.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:26.839094326-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:26.845379616-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:26.845411837-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:26.846083044-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:26.846108438-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}\n', '')
- UNVERIFIED: summean-mean-exact (FEATURE-SumMean.md)
  intent: The mean is the exact total divided by the non-null count.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:27.736031784-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:27.743965035-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:27.744011201-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:27.74463083-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:27.74464745-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:27.736031784-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:27.743965035-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:27.744011201-04:00","Action":"output","Package":"csvstat/… (+353 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summean-mean-exact.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:27.736031784-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:27.743965035-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:27.744011201-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:27.74463083-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.008s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:27.74464745-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}\n', '')
- UNVERIFIED: summean-ignores-nulls (FEATURE-SumMean.md)
  intent: The mean divisor is the non-null count.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:28.676136692-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:28.683313364-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:28.683360716-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:28.684004527-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:28.684026134-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:28.676136692-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:28.683313364-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:28.683360716-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summean-ignores-nulls.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:28.676136692-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:28.683313364-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:28.683360716-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:28.684004527-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.008s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:28.684026134-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}\n', '')
- UNVERIFIED: summean-non-numeric (FEATURE-SumMean.md)
  intent: Non-numeric column types have no sum or mean.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:29.573702902-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:29.579822103-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:29.57987389-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:29.580401498-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:29.580432397-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:29.573702902-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:29.579822103-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:29.57987389-04:00","Action":"output","Package":"csvstat/i… (+354 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "summean-non-numeric.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:29.573702902-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:29.579822103-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:29.57987389-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:29.580401498-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:29.580432397-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}\n', '')
- UNVERIFIED: median-odd (FEATURE-Median.md)
  intent: An odd count yields the middle value of the sorted values.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:30.751565798-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:30.757714291-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:30.757766224-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:30.758470295-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:30.758502268-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:30.751565798-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:30.757714291-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:30.757766224-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "median-odd.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:30.751565798-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:30.757714291-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:30.757766224-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:30.758470295-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:30.758502268-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}\n', '')
- UNVERIFIED: median-even (FEATURE-Median.md)
  intent: An even count yields the exact mean of the two middle values.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:31.545206027-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:31.55176983-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:31.551806989-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:31.55265952-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:31.552682969-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:31.545206027-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:31.55176983-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:31.551806989-04:00","Action":"output","Package":"csvstat/i… (+353 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "median-even.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:31.545206027-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:31.55176983-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:31.551806989-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:31.55265952-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:31.552682969-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}\n', '')
- UNVERIFIED: median-numeric-order (FEATURE-Median.md)
  intent: Values are sorted numerically, not lexically.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:32.538576092-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:32.544538077-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:32.544588518-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:32.545218131-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:32.545241936-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:32.538576092-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:32.544538077-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:32.544588518-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "median-numeric-order.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:32.538576092-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:32.544538077-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:32.544588518-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:32.545218131-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:32.545241936-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}\n', '')
- UNVERIFIED: median-ignores-nulls (FEATURE-Median.md)
  intent: Null values do not count toward the median.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:33.621218796-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:33.62892202-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:33.628963515-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:33.629711204-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:33.629786452-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:33.621218796-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:33.62892202-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:33.628963515-04:00","Action":"output","Package":"csvstat/i… (+354 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "median-ignores-nulls.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:33.621218796-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:33.62892202-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:33.628963515-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:33.629711204-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.008s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:33.629786452-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}\n', '')
- UNVERIFIED: median-non-numeric (FEATURE-Median.md)
  intent: Non-numeric and empty columns have no median.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:34.564875139-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:34.571000533-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:34.571039578-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:34.571770213-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:34.571795048-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:34.564875139-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:34.571000533-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:34.571039578-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "median-non-numeric.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:34.564875139-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:34.571000533-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:34.571039578-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:34.571770213-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:34.571795048-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}\n', '')
- UNVERIFIED: minmax-integer (FEATURE-MinMax.md)
  intent: Integer extremes are ordered numerically with exact arithmetic.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:35.503964219-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:35.50934459-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:35.50938526-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:35.510105658-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:35.510125028-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:35.503964219-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:35.50934459-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:35.50938526-04:00","Action":"output","Package":"csvstat/in… (+353 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "minmax-integer.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:35.503964219-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:35.50934459-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:35.50938526-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:35.510105658-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:35.510125028-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}\n', '')
- UNVERIFIED: minmax-number (FEATURE-MinMax.md)
  intent: Number extremes are ordered exactly as decimals.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:36.662457008-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:36.669605626-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:36.669651064-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:36.670276904-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:36.670300002-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:36.662457008-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:36.669605626-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:36.669651064-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "minmax-number.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:36.662457008-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:36.669605626-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:36.669651064-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:36.670276904-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.008s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:36.670300002-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}\n', '')
- UNVERIFIED: minmax-date (FEATURE-MinMax.md)
  intent: Date extremes are the earliest and latest dates.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:37.624131357-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:37.631182522-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:37.6312245-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:37.631925576-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:37.631947124-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:37.624131357-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:37.631182522-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:37.6312245-04:00","Action":"output","Package":"csvstat/in… (+353 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "minmax-date.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:37.624131357-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:37.631182522-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:37.6312245-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:37.631925576-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:37.631947124-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}\n', '')
- UNVERIFIED: minmax-boolean (FEATURE-MinMax.md)
  intent: Boolean extremes follow the false/true rule and are lowercase regardless of input case.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:38.617094295-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:38.6222852-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:38.622327316-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:38.6229366-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:38.622958756-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:38.617094295-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:38.6222852-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:38.622327316-04:00","Action":"output","Package":"csvstat/in… (+351 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "minmax-boolean.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:38.617094295-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:38.6222852-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:38.622327316-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:38.6229366-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:38.622958756-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}\n', '')
- UNVERIFIED: minmax-text (FEATURE-MinMax.md)
  intent: Text extremes are ordered by byte value.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:39.676956243-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:39.683369031-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:39.683453318-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:39.683904659-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:39.683929457-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:39.676956243-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:39.683369031-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:39.683453318-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "minmax-text.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:39.676956243-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:39.683369031-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:39.683453318-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:39.683904659-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:39.683929457-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}\n', '')
- UNVERIFIED: minmax-empty (FEATURE-MinMax.md)
  intent: An empty column has blank minimum and maximum.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:40.72414908-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:40.729497291-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:40.729552287-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:40.730068193-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:40.730091372-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:40.72414908-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:40.729497291-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:40.729552287-04:00","Action":"output","Package":"csvstat/i… (+354 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "minmax-empty.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:40.72414908-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:40.729497291-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:40.729552287-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:40.730068193-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.006s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:40.730091372-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}\n', '')
- UNVERIFIED: minmax-single-value (FEATURE-MinMax.md)
  intent: A single non-null value is both the minimum and the maximum.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:41.606201139-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:41.610113707-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:41.610154334-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:41.610795588-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.004s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:41.610816013-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.005}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:41.606201139-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:41.610113707-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:41.610154334-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "minmax-single-value.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:41.606201139-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:41.610113707-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:41.610154334-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:41.610795588-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.004s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:41.610816013-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.005}\n', '')
- UNVERIFIED: decfmt-two-places (FEATURE-DecimalFormat.md)
  intent: FormatDecimal always writes exactly two decimal places.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:42.707036109-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:42.714569596-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:42.714607536-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:42.715895432-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.009s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:42.715916794-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:42.707036109-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:42.714569596-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:42.714607536-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "decfmt-two-places.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:42.707036109-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:42.714569596-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:42.714607536-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:42.715895432-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.009s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:42.715916794-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}\n', '')
- UNVERIFIED: decfmt-half-away (FEATURE-DecimalFormat.md)
  intent: Exact ties round half away from zero for both signs.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:43.85313935-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:43.857103101-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:43.857135941-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:43.857687079-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.004s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:43.857710503-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.005}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:43.85313935-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:43.857103101-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:43.857135941-04:00","Action":"output","Package":"csvstat/i… (+354 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "decfmt-half-away.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:43.85313935-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:43.857103101-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:43.857135941-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:43.857687079-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.004s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:43.857710503-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.005}\n', '')
- UNVERIFIED: decfmt-no-negative-zero (FEATURE-DecimalFormat.md)
  intent: A negative value that rounds to zero is written without a minus sign.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:44.867534359-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:44.874526427-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:44.874575624-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:44.875334826-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:44.875355162-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:44.867534359-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:44.874526427-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:44.874575624-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "decfmt-no-negative-zero.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:44.867534359-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:44.874526427-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:44.874575624-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:44.875334826-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.008s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:44.875355162-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}\n', '')
- UNVERIFIED: decfmt-no-separator (FEATURE-DecimalFormat.md)
  intent: Large values are written without a thousands separator.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:45.892838498-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:45.898036457-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:45.89808348-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:45.898595767-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.005s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:45.898621339-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:45.892838498-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:45.898036457-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:45.89808348-04:00","Action":"output","Package":"csvstat/i… (+354 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "decfmt-no-separator.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:45.892838498-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:45.898036457-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:45.89808348-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:45.898595767-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.005s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:45.898621339-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}\n', '')
- UNVERIFIED: decfmt-large-exact (FEATURE-DecimalFormat.md)
  intent: Values beyond float64 precision keep every digit.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:46.872910534-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:46.881353692-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:46.881409274-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:46.882052983-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.009s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:46.882089302-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:46.872910534-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:46.881353692-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:46.881409274-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "decfmt-large-exact.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:46.872910534-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:46.881353692-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:46.881409274-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:46.882052983-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.009s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:46.882089302-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}\n', '')
- UNVERIFIED: decfmt-integer (FEATURE-DecimalFormat.md)
  intent: FormatInteger writes plain integers with no decimal point.
  return code: 1
  error: unverified acceptance: not-run
  stdout:
    {"Time":"2026-10-25T01:38:47.788782227-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:38:47.795056391-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\n"}
    {"Time":"2026-10-25T01:38:47.795105019-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:38:47.795820068-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s [no tests to run]\n"}
    {"Time":"2026-10-25T01:38:47.795841973-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "not-run"}
    --- drydock: values at failure ---
      result = GoResult(passed=False, status='not-run', return_code=0, stdout='{"Time":"2026-10-25T01:38:47.788782227-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:47.795056391-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:47.795105019-04:00","Action":"output","Package":"csvstat/… (+355 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "decfmt-integer.py", line 6, in <module>
        assert result.passed, (result.status, result.stdout, result.stderr)
               ^^^^^^^^^^^^^
    AssertionError: ('not-run', '{"Time":"2026-10-25T01:38:47.788782227-04:00","Action":"start","Package":"csvstat/internal/profile"}\n{"Time":"2026-10-25T01:38:47.795056391-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"testing: warning: no tests to run\\n"}\n{"Time":"2026-10-25T01:38:47.795105019-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\\n"}\n{"Time":"2026-10-25T01:38:47.795820068-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \\tcsvstat/internal/profile\\t0.007s [no tests to run]\\n"}\n{"Time":"2026-10-25T01:38:47.795841973-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}\n', '')
- RED: freq-count-ordering (FEATURE-FreqCount.md)
  intent: Values are listed by count descending then by byte value ascending, with counts matching the input.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      parsed = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-count-ordering.py", line 19, in <module>
        assert parsed[0] == ["value", "count"]
               ~~~~~~^^^
    IndexError: list index out of range
- RED: freq-count-nulls-and-trim (FEATURE-FreqCount.md)
  intent: Null spellings are omitted and padded values count with their trimmed form.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      parsed = []
      ordered = [('x', 2), ('y', 1)]
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-count-nulls-and-trim.py", line 22, in <module>
        assert parsed[1:] == [[k, str(n)] for k, n in ordered]
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: freq-count-all-null (FEATURE-FreqCount.md)
  intent: A column with no non-null value lists only the header.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      parsed = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-count-all-null.py", line 17, in <module>
        assert parsed == [["value", "count"]]
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: freq-count-header-only (FEATURE-FreqCount.md)
  intent: A header-only file lists only the header for any of its columns.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      out = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-count-header-only.py", line 16, in <module>
        assert list(csv.reader(io.StringIO(out))) == [["value", "count"]]
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: freq-count-quoted-value (FEATURE-FreqCount.md)
  intent: A value containing a comma is written quoted and parses back to the original value.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      parsed = []
      value = 'relocating, TX'
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-count-quoted-value.py", line 18, in <module>
        assert parsed[1] == [value, str(2)]
               ~~~~~~^^^
    IndexError: list index out of range
- GREEN (prepassed): freq-count-rejected-input (FEATURE-FreqCount.md)
  intent: A ragged file in frequency mode exits 3 and writes nothing to standard output.
  return code: 0
  stdout:
    csvstat: /tmp/tmpot6vvcp9/ragged.csv: line 3: record has 1 fields, header has 2
- RED: freq-limit-default-five (FEATURE-FreqLimit.md)
  intent: Without --top the listing has at most five rows, the first five of the ordered list.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      parsed = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-limit-default-five.py", line 24, in <module>
        assert parsed[0] == ["value", "count"]
               ~~~~~~^^^
    IndexError: list index out of range
- RED: freq-limit-explicit-top (FEATURE-FreqLimit.md)
  intent: --top N lists exactly the first N ordered rows.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      parsed = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-limit-explicit-top.py", line 24, in <module>
        assert parsed[0] == ["value", "count"]
               ~~~~~~^^^
    IndexError: list index out of range
- RED: freq-limit-larger-than-distinct (FEATURE-FreqLimit.md)
  intent: A --top larger than the distinct count lists every value.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      parsed = []
      ordered = [('c', 3), ('b', 2), ('a', 1)]
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-limit-larger-than-distinct.py", line 21, in <module>
        assert parsed[1:] == [[k, str(n)] for k, n in ordered]
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: freq-limit-all-null-column (FEATURE-FreqLimit.md)
  intent: A column with no non-null values lists only the header even with --top given.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      out = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-limit-all-null-column.py", line 16, in <module>
        assert list(csv.reader(io.StringIO(out))) == [["value", "count"]]
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: profile-output-columns (FEATURE-ProfileOutput.md)
  intent: The profile has the declared header and one row per column in file order, with types, counts, integer sum, and decimal sum and mean derived from the input.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      header = ['column', 'type', 'rows', 'nulls', 'distinct', 'min', 'max', 'sum', 'mean', 'median']
      parsed = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-output-columns.py", line 19, in <module>
        assert parsed[0] == header
               ~~~~~~^^^
    IndexError: list index out of range
- RED: profile-output-quoting (FEATURE-ProfileOutput.md)
  intent: A field containing a comma and a double quote is written quoted so the parsed value equals the trimmed input value.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      parsed = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-output-quoting.py", line 18, in <module>
        assert len(parsed) == 2
               ^^^^^^^^^^^^^^^^
    AssertionError
- RED: profile-output-line-endings (FEATURE-ProfileOutput.md)
  intent: Output uses newline line endings with one line per column plus the header and a trailing newline.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      out = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-output-line-endings.py", line 16, in <module>
        assert out.endswith("\n")
               ^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: profile-output-header-only (FEATURE-ProfileOutput.md)
  intent: A header-only file yields one empty-typed row per column with zero rows.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      rows = [['sku', 'price', 'in_stock']]
      body = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-output-header-only.py", line 17, in <module>
        assert [r[0] for r in body] == rows[0]
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: profile-output-rounding (FEATURE-ProfileOutput.md)
  intent: Number results round half away from zero once, and a total or mean that rounds to zero carries no minus sign.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      pos_mean = Decimal('0.13')
      body = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-output-rounding.py", line 21, in <module>
        assert body[0][8] == str(pos_mean)
               ~~~~^^^
    IndexError: list index out of range
- GREEN (prepassed): profile-output-rejected-input (FEATURE-ProfileOutput.md)
  intent: A ragged file exits 3 and writes nothing to standard output.
  return code: 0
  stdout:
    csvstat: /tmp/tmpxrg_4m8g/ragged.csv: line 3: record has 3 fields, header has 2

## Post-build programmatic acceptance
- PASS: rowcounts-rows-nulls-distinct (FEATURE-RowCounts.md)
  intent: Counts reports rows, nulls, and distinct for a mixed column.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:44:48.786017178-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:44:48.792182676-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestCountsRowsNullsDistinct"}
    {"Time":"2026-10-25T01:44:48.7922087-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestCountsRowsNullsDistinct","Output":"=== RUN   TestCountsRowsNullsDistinct\n"}
    {"Time":"2026-10-25T01:44:48.792227068-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestCountsRowsNullsDistinct","Output":"--- PASS: TestCountsRowsNullsDistinct (0.00s)\n"}
    {"Time":"2026-10-25T01:44:48.792229079-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestCountsRowsNullsDistinct","Elapsed":0}
    {"Time":"2026-10-25T01:44:48.792231787-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:44:48.792951582-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:44:48.792975018-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: rowcounts-header-only (FEATURE-RowCounts.md)
  intent: A column with no data records reports zero for every count.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:44:49.956998883-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:44:49.96343725-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestCountsHeaderOnly"}
    {"Time":"2026-10-25T01:44:49.963485903-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestCountsHeaderOnly","Output":"=== RUN   TestCountsHeaderOnly\n"}
    {"Time":"2026-10-25T01:44:49.963557294-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestCountsHeaderOnly","Output":"--- PASS: TestCountsHeaderOnly (0.00s)\n"}
    {"Time":"2026-10-25T01:44:49.963585529-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestCountsHeaderOnly","Elapsed":0}
    {"Time":"2026-10-25T01:44:49.963621186-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:44:49.964066592-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:44:49.964086012-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: rowcounts-distinct-exact (FEATURE-RowCounts.md)
  intent: Distinct compares exact strings and counts repeats once.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:44:51.113778826-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:44:51.118557752-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestCountsDistinctExactStrings"}
    {"Time":"2026-10-25T01:44:51.118580687-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestCountsDistinctExactStrings","Output":"=== RUN   TestCountsDistinctExactStrings\n"}
    {"Time":"2026-10-25T01:44:51.118598973-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestCountsDistinctExactStrings","Output":"--- PASS: TestCountsDistinctExactStrings (0.00s)\n"}
    {"Time":"2026-10-25T01:44:51.11860061-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestCountsDistinctExactStrings","Elapsed":0}
    {"Time":"2026-10-25T01:44:51.118603653-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:44:51.119140161-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.005s\n"}
    {"Time":"2026-10-25T01:44:51.119164134-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.005}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: rowcounts-all-null (FEATURE-RowCounts.md)
  intent: An all-null column has zero distinct values and every row null.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:44:52.485201269-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:44:52.491860265-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestCountsAllNull"}
    {"Time":"2026-10-25T01:44:52.491886245-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestCountsAllNull","Output":"=== RUN   TestCountsAllNull\n"}
    {"Time":"2026-10-25T01:44:52.4919083-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestCountsAllNull","Output":"--- PASS: TestCountsAllNull (0.00s)\n"}
    {"Time":"2026-10-25T01:44:52.491909965-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestCountsAllNull","Elapsed":0}
    {"Time":"2026-10-25T01:44:52.491913552-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:44:52.492711908-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:44:52.492734213-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: summean-integer-sum (FEATURE-SumMean.md)
  intent: Integer sums are exact, including negative values.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:44:53.708804253-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:44:53.716082765-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestSumIntegerExact"}
    {"Time":"2026-10-25T01:44:53.716110304-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestSumIntegerExact","Output":"=== RUN   TestSumIntegerExact\n"}
    {"Time":"2026-10-25T01:44:53.716131031-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestSumIntegerExact","Output":"--- PASS: TestSumIntegerExact (0.00s)\n"}
    {"Time":"2026-10-25T01:44:53.716136844-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestSumIntegerExact","Elapsed":0}
    {"Time":"2026-10-25T01:44:53.71614-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:44:53.716776427-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s\n"}
    {"Time":"2026-10-25T01:44:53.716798182-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: summean-number-no-drift (FEATURE-SumMean.md)
  intent: Decimal sums carry no floating-point drift.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:44:54.979892106-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:44:54.985259722-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestSumNumberNoFloatDrift"}
    {"Time":"2026-10-25T01:44:54.98531269-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestSumNumberNoFloatDrift","Output":"=== RUN   TestSumNumberNoFloatDrift\n"}
    {"Time":"2026-10-25T01:44:54.985322097-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestSumNumberNoFloatDrift","Output":"--- PASS: TestSumNumberNoFloatDrift (0.00s)\n"}
    {"Time":"2026-10-25T01:44:54.985324442-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestSumNumberNoFloatDrift","Elapsed":0}
    {"Time":"2026-10-25T01:44:54.985328467-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:44:54.9859396-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s\n"}
    {"Time":"2026-10-25T01:44:54.985968474-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: summean-mean-exact (FEATURE-SumMean.md)
  intent: The mean is the exact total divided by the non-null count.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:44:56.159257928-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:44:56.165417464-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMeanExactRational"}
    {"Time":"2026-10-25T01:44:56.165459664-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMeanExactRational","Output":"=== RUN   TestMeanExactRational\n"}
    {"Time":"2026-10-25T01:44:56.165466451-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMeanExactRational","Output":"--- PASS: TestMeanExactRational (0.00s)\n"}
    {"Time":"2026-10-25T01:44:56.165468078-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMeanExactRational","Elapsed":0}
    {"Time":"2026-10-25T01:44:56.165470423-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:44:56.166160361-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:44:56.16618031-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: summean-ignores-nulls (FEATURE-SumMean.md)
  intent: The mean divisor is the non-null count.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:44:57.408591485-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:44:57.41468512-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestSumMeanIgnoresNulls"}
    {"Time":"2026-10-25T01:44:57.414732558-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestSumMeanIgnoresNulls","Output":"=== RUN   TestSumMeanIgnoresNulls\n"}
    {"Time":"2026-10-25T01:44:57.4147405-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestSumMeanIgnoresNulls","Output":"--- PASS: TestSumMeanIgnoresNulls (0.00s)\n"}
    {"Time":"2026-10-25T01:44:57.414742169-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestSumMeanIgnoresNulls","Elapsed":0}
    {"Time":"2026-10-25T01:44:57.414745982-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:44:57.41546477-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:44:57.415489038-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: summean-non-numeric (FEATURE-SumMean.md)
  intent: Non-numeric column types have no sum or mean.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:44:58.707764821-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:44:58.712808307-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestSumMeanNonNumeric"}
    {"Time":"2026-10-25T01:44:58.712832778-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestSumMeanNonNumeric","Output":"=== RUN   TestSumMeanNonNumeric\n"}
    {"Time":"2026-10-25T01:44:58.712860818-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestSumMeanNonNumeric","Output":"--- PASS: TestSumMeanNonNumeric (0.00s)\n"}
    {"Time":"2026-10-25T01:44:58.712863234-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestSumMeanNonNumeric","Elapsed":0}
    {"Time":"2026-10-25T01:44:58.712865746-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:44:58.713314026-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.005s\n"}
    {"Time":"2026-10-25T01:44:58.713335622-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: median-odd (FEATURE-Median.md)
  intent: An odd count yields the middle value of the sorted values.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:00.072038983-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:00.078309236-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMedianOddCount"}
    {"Time":"2026-10-25T01:45:00.078340674-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMedianOddCount","Output":"=== RUN   TestMedianOddCount\n"}
    {"Time":"2026-10-25T01:45:00.078370955-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMedianOddCount","Output":"--- PASS: TestMedianOddCount (0.00s)\n"}
    {"Time":"2026-10-25T01:45:00.07837445-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMedianOddCount","Elapsed":0}
    {"Time":"2026-10-25T01:45:00.078378069-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:00.07939977-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:45:00.079429573-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: median-even (FEATURE-Median.md)
  intent: An even count yields the exact mean of the two middle values.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:01.312754701-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:01.318387365-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMedianEvenCountExact"}
    {"Time":"2026-10-25T01:45:01.318439534-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMedianEvenCountExact","Output":"=== RUN   TestMedianEvenCountExact\n"}
    {"Time":"2026-10-25T01:45:01.318448238-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMedianEvenCountExact","Output":"--- PASS: TestMedianEvenCountExact (0.00s)\n"}
    {"Time":"2026-10-25T01:45:01.318449906-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMedianEvenCountExact","Elapsed":0}
    {"Time":"2026-10-25T01:45:01.318453655-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:01.318885445-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s\n"}
    {"Time":"2026-10-25T01:45:01.318916043-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: median-numeric-order (FEATURE-Median.md)
  intent: Values are sorted numerically, not lexically.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:02.558164255-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:02.564960729-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMedianNumericOrder"}
    {"Time":"2026-10-25T01:45:02.565009046-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMedianNumericOrder","Output":"=== RUN   TestMedianNumericOrder\n"}
    {"Time":"2026-10-25T01:45:02.565017835-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMedianNumericOrder","Output":"--- PASS: TestMedianNumericOrder (0.00s)\n"}
    {"Time":"2026-10-25T01:45:02.56501954-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMedianNumericOrder","Elapsed":0}
    {"Time":"2026-10-25T01:45:02.565022218-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:02.565678208-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:45:02.565695421-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: median-ignores-nulls (FEATURE-Median.md)
  intent: Null values do not count toward the median.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:03.836867819-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:03.84308753-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMedianIgnoresNulls"}
    {"Time":"2026-10-25T01:45:03.843118288-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMedianIgnoresNulls","Output":"=== RUN   TestMedianIgnoresNulls\n"}
    {"Time":"2026-10-25T01:45:03.843139104-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMedianIgnoresNulls","Output":"--- PASS: TestMedianIgnoresNulls (0.00s)\n"}
    {"Time":"2026-10-25T01:45:03.843140652-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMedianIgnoresNulls","Elapsed":0}
    {"Time":"2026-10-25T01:45:03.843143484-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:03.843912336-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:45:03.843932822-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: median-non-numeric (FEATURE-Median.md)
  intent: Non-numeric and empty columns have no median.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:05.057281986-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:05.063838022-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMedianNonNumeric"}
    {"Time":"2026-10-25T01:45:05.063894509-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMedianNonNumeric","Output":"=== RUN   TestMedianNonNumeric\n"}
    {"Time":"2026-10-25T01:45:05.06390151-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMedianNonNumeric","Output":"--- PASS: TestMedianNonNumeric (0.00s)\n"}
    {"Time":"2026-10-25T01:45:05.063903745-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMedianNonNumeric","Elapsed":0}
    {"Time":"2026-10-25T01:45:05.063908431-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:05.064528969-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s\n"}
    {"Time":"2026-10-25T01:45:05.064565904-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: minmax-integer (FEATURE-MinMax.md)
  intent: Integer extremes are ordered numerically with exact arithmetic.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:06.29104355-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:06.297509443-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMinMaxInteger"}
    {"Time":"2026-10-25T01:45:06.297538085-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxInteger","Output":"=== RUN   TestMinMaxInteger\n"}
    {"Time":"2026-10-25T01:45:06.297555978-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxInteger","Output":"--- PASS: TestMinMaxInteger (0.00s)\n"}
    {"Time":"2026-10-25T01:45:06.297557966-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMinMaxInteger","Elapsed":0}
    {"Time":"2026-10-25T01:45:06.297560719-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:06.299090304-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.008s\n"}
    {"Time":"2026-10-25T01:45:06.29911463-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: minmax-number (FEATURE-MinMax.md)
  intent: Number extremes are ordered exactly as decimals.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:07.533086742-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:07.539058462-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMinMaxNumber"}
    {"Time":"2026-10-25T01:45:07.539093476-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxNumber","Output":"=== RUN   TestMinMaxNumber\n"}
    {"Time":"2026-10-25T01:45:07.539101457-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxNumber","Output":"--- PASS: TestMinMaxNumber (0.00s)\n"}
    {"Time":"2026-10-25T01:45:07.539106237-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMinMaxNumber","Elapsed":0}
    {"Time":"2026-10-25T01:45:07.539109457-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:07.53965864-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s\n"}
    {"Time":"2026-10-25T01:45:07.539696265-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: minmax-date (FEATURE-MinMax.md)
  intent: Date extremes are the earliest and latest dates.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:08.750289995-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:08.756604051-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMinMaxDate"}
    {"Time":"2026-10-25T01:45:08.756628534-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxDate","Output":"=== RUN   TestMinMaxDate\n"}
    {"Time":"2026-10-25T01:45:08.756648543-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxDate","Output":"--- PASS: TestMinMaxDate (0.00s)\n"}
    {"Time":"2026-10-25T01:45:08.756650309-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMinMaxDate","Elapsed":0}
    {"Time":"2026-10-25T01:45:08.756652813-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:08.757593163-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:45:08.757620928-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: minmax-boolean (FEATURE-MinMax.md)
  intent: Boolean extremes follow the false/true rule and are lowercase regardless of input case.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:10.095338382-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:10.103047535-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean"}
    {"Time":"2026-10-25T01:45:10.103073561-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean","Output":"=== RUN   TestMinMaxBoolean\n"}
    {"Time":"2026-10-25T01:45:10.103089985-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/mixed_case"}
    {"Time":"2026-10-25T01:45:10.103091349-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/mixed_case","Output":"=== RUN   TestMinMaxBoolean/mixed_case\n"}
    {"Time":"2026-10-25T01:45:10.103094295-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/mixed_case","Output":"--- PASS: TestMinMaxBoolean/mixed_case (0.00s)\n"}
    {"Time":"2026-10-25T01:45:10.103097083-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/mixed_case","Elapsed":0}
    {"Time":"2026-10-25T01:45:10.103100751-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/all_true"}
    {"Time":"2026-10-25T01:45:10.103101609-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/all_true","Output":"=== RUN   TestMinMaxBoolean/all_true\n"}
    {"Time":"2026-10-25T01:45:10.103124226-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/all_true","Output":"--- PASS: TestMinMaxBoolean/all_true (0.00s)\n"}
    {"Time":"2026-10-25T01:45:10.103135352-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/all_true","Elapsed":0}
    {"Time":"2026-10-25T01:45:10.103162131-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/all_false"}
    {"Time":"2026-10-25T01:45:10.103180289-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/all_false","Output":"=== RUN   TestMinMaxBoolean/all_false\n"}
    {"Time":"2026-10-25T01:45:10.10318431-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/all_false","Output":"--- PASS: TestMinMaxBoolean/all_false (0.00s)\n"}
    {"Time":"2026-10-25T01:45:10.10318615-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/all_false","Elapsed":0}
    {"Time":"2026-10-25T01:45:10.103196817-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/mixed"}
    {"Time":"2026-10-25T01:45:10.103198276-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/mixed","Output":"=== RUN   TestMinMaxBoolean/mixed\n"}
    {"Time":"2026-10-25T01:45:10.10320234-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/mixed","Output":"--- PASS: TestMinMaxBoolean/mixed (0.00s)\n"}
    {"Time":"2026-10-25T01:45:10.103203649-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean/mixed","Elapsed":0}
    {"Time":"2026-10-25T01:45:10.103219482-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean","Output":"--- PASS: TestMinMaxBoolean (0.00s)\n"}
    {"Time":"2026-10-25T01:45:10.10322152-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMinMaxBoolean","Elapsed":0}
    {"Time":"2026-10-25T01:45:10.103222788-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:10.10454123-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.009s\n"}
    {"Time":"2026-10-25T01:45:10.104595636-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.009}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: minmax-text (FEATURE-MinMax.md)
  intent: Text extremes are ordered by byte value.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:11.341615814-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:11.347476176-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMinMaxText"}
    {"Time":"2026-10-25T01:45:11.347498015-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxText","Output":"=== RUN   TestMinMaxText\n"}
    {"Time":"2026-10-25T01:45:11.3475155-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxText","Output":"--- PASS: TestMinMaxText (0.00s)\n"}
    {"Time":"2026-10-25T01:45:11.347516992-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMinMaxText","Elapsed":0}
    {"Time":"2026-10-25T01:45:11.347519596-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:11.348186709-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s\n"}
    {"Time":"2026-10-25T01:45:11.348211049-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: minmax-empty (FEATURE-MinMax.md)
  intent: An empty column has blank minimum and maximum.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:12.537378798-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:12.543318307-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMinMaxEmpty"}
    {"Time":"2026-10-25T01:45:12.543361615-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxEmpty","Output":"=== RUN   TestMinMaxEmpty\n"}
    {"Time":"2026-10-25T01:45:12.543377955-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxEmpty","Output":"--- PASS: TestMinMaxEmpty (0.00s)\n"}
    {"Time":"2026-10-25T01:45:12.543380218-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMinMaxEmpty","Elapsed":0}
    {"Time":"2026-10-25T01:45:12.54338296-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:12.543931249-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s\n"}
    {"Time":"2026-10-25T01:45:12.54395439-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: minmax-single-value (FEATURE-MinMax.md)
  intent: A single non-null value is both the minimum and the maximum.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:13.756239483-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:13.760792443-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestMinMaxSingleValue"}
    {"Time":"2026-10-25T01:45:13.760815742-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxSingleValue","Output":"=== RUN   TestMinMaxSingleValue\n"}
    {"Time":"2026-10-25T01:45:13.760836686-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestMinMaxSingleValue","Output":"--- PASS: TestMinMaxSingleValue (0.00s)\n"}
    {"Time":"2026-10-25T01:45:13.760838453-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestMinMaxSingleValue","Elapsed":0}
    {"Time":"2026-10-25T01:45:13.760841176-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:13.761374662-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.005s\n"}
    {"Time":"2026-10-25T01:45:13.761397073-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.005}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: decfmt-two-places (FEATURE-DecimalFormat.md)
  intent: FormatDecimal always writes exactly two decimal places.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:15.017744462-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:15.023843728-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestFormatDecimalTwoPlaces"}
    {"Time":"2026-10-25T01:45:15.023887648-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestFormatDecimalTwoPlaces","Output":"=== RUN   TestFormatDecimalTwoPlaces\n"}
    {"Time":"2026-10-25T01:45:15.023893893-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestFormatDecimalTwoPlaces","Output":"--- PASS: TestFormatDecimalTwoPlaces (0.00s)\n"}
    {"Time":"2026-10-25T01:45:15.023897223-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestFormatDecimalTwoPlaces","Elapsed":0}
    {"Time":"2026-10-25T01:45:15.02390051-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:15.024369502-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s\n"}
    {"Time":"2026-10-25T01:45:15.024394333-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: decfmt-half-away (FEATURE-DecimalFormat.md)
  intent: Exact ties round half away from zero for both signs.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:16.243341376-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:16.248863974-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestFormatDecimalHalfAwayFromZero"}
    {"Time":"2026-10-25T01:45:16.248971516-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestFormatDecimalHalfAwayFromZero","Output":"=== RUN   TestFormatDecimalHalfAwayFromZero\n"}
    {"Time":"2026-10-25T01:45:16.248990273-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestFormatDecimalHalfAwayFromZero","Output":"--- PASS: TestFormatDecimalHalfAwayFromZero (0.00s)\n"}
    {"Time":"2026-10-25T01:45:16.248993181-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestFormatDecimalHalfAwayFromZero","Elapsed":0}
    {"Time":"2026-10-25T01:45:16.248995941-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:16.249568738-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s\n"}
    {"Time":"2026-10-25T01:45:16.249587886-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: decfmt-no-negative-zero (FEATURE-DecimalFormat.md)
  intent: A negative value that rounds to zero is written without a minus sign.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:17.396608701-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:17.401947445-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestFormatDecimalNoNegativeZero"}
    {"Time":"2026-10-25T01:45:17.401972514-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestFormatDecimalNoNegativeZero","Output":"=== RUN   TestFormatDecimalNoNegativeZero\n"}
    {"Time":"2026-10-25T01:45:17.401992867-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestFormatDecimalNoNegativeZero","Output":"--- PASS: TestFormatDecimalNoNegativeZero (0.00s)\n"}
    {"Time":"2026-10-25T01:45:17.401994431-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestFormatDecimalNoNegativeZero","Elapsed":0}
    {"Time":"2026-10-25T01:45:17.401997924-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:17.402585448-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s\n"}
    {"Time":"2026-10-25T01:45:17.402609353-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.006}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: decfmt-no-separator (FEATURE-DecimalFormat.md)
  intent: Large values are written without a thousands separator.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:18.613288639-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:18.619253045-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestFormatDecimalNoSeparator"}
    {"Time":"2026-10-25T01:45:18.619293897-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestFormatDecimalNoSeparator","Output":"=== RUN   TestFormatDecimalNoSeparator\n"}
    {"Time":"2026-10-25T01:45:18.619311011-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestFormatDecimalNoSeparator","Output":"--- PASS: TestFormatDecimalNoSeparator (0.00s)\n"}
    {"Time":"2026-10-25T01:45:18.619313734-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestFormatDecimalNoSeparator","Elapsed":0}
    {"Time":"2026-10-25T01:45:18.619317015-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:18.619981206-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.006s\n"}
    {"Time":"2026-10-25T01:45:18.620003736-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.007}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: decfmt-large-exact (FEATURE-DecimalFormat.md)
  intent: Values beyond float64 precision keep every digit.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:19.872604227-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:19.87960769-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestFormatDecimalLargeExact"}
    {"Time":"2026-10-25T01:45:19.879662111-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestFormatDecimalLargeExact","Output":"=== RUN   TestFormatDecimalLargeExact\n"}
    {"Time":"2026-10-25T01:45:19.879670917-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestFormatDecimalLargeExact","Output":"--- PASS: TestFormatDecimalLargeExact (0.00s)\n"}
    {"Time":"2026-10-25T01:45:19.879672946-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestFormatDecimalLargeExact","Elapsed":0}
    {"Time":"2026-10-25T01:45:19.879676503-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:19.88029857-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:45:19.880321855-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: decfmt-integer (FEATURE-DecimalFormat.md)
  intent: FormatInteger writes plain integers with no decimal point.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    {"Time":"2026-10-25T01:45:21.041760769-04:00","Action":"start","Package":"csvstat/internal/profile"}
    {"Time":"2026-10-25T01:45:21.048690513-04:00","Action":"run","Package":"csvstat/internal/profile","Test":"TestFormatInteger"}
    {"Time":"2026-10-25T01:45:21.048725207-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestFormatInteger","Output":"=== RUN   TestFormatInteger\n"}
    {"Time":"2026-10-25T01:45:21.048747904-04:00","Action":"output","Package":"csvstat/internal/profile","Test":"TestFormatInteger","Output":"--- PASS: TestFormatInteger (0.00s)\n"}
    {"Time":"2026-10-25T01:45:21.04874991-04:00","Action":"pass","Package":"csvstat/internal/profile","Test":"TestFormatInteger","Elapsed":0}
    {"Time":"2026-10-25T01:45:21.048753346-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"PASS\n"}
    {"Time":"2026-10-25T01:45:21.0494562-04:00","Action":"output","Package":"csvstat/internal/profile","Output":"ok  \tcsvstat/internal/profile\t0.007s\n"}
    {"Time":"2026-10-25T01:45:21.049487802-04:00","Action":"pass","Package":"csvstat/internal/profile","Elapsed":0.008}
  stderr:
    DRYDOCK_NATIVE_RESULT={"status": "pass"}
- PASS: freq-count-ordering (FEATURE-FreqCount.md)
  intent: Values are listed by count descending then by byte value ascending, with counts matching the input.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    b,3
    B,2
    a,2
    c,2
    z,1
- PASS: freq-count-nulls-and-trim (FEATURE-FreqCount.md)
  intent: Null spellings are omitted and padded values count with their trimmed form.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    x,2
    y,1
- PASS: freq-count-all-null (FEATURE-FreqCount.md)
  intent: A column with no non-null value lists only the header.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
- PASS: freq-count-header-only (FEATURE-FreqCount.md)
  intent: A header-only file lists only the header for any of its columns.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
- PASS: freq-count-quoted-value (FEATURE-FreqCount.md)
  intent: A value containing a comma is written quoted and parses back to the original value.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    "relocating, TX",2
    other,1
- PASS: freq-count-rejected-input (FEATURE-FreqCount.md)
  intent: A ragged file in frequency mode exits 3 and writes nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmp9gegk44b/ragged.csv: line 3: record has 1 fields, header has 2
- PASS: freq-limit-default-five (FEATURE-FreqLimit.md)
  intent: Without --top the listing has at most five rows, the first five of the ordered list.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    v0,8
    v1,7
    v2,6
    v3,5
    v4,4
- PASS: freq-limit-explicit-top (FEATURE-FreqLimit.md)
  intent: --top N lists exactly the first N ordered rows.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    v0,8
    
    
    value,count
    v0,8
    v1,7
    
    
    value,count
    v0,8
    v1,7
    v2,6
- PASS: freq-limit-larger-than-distinct (FEATURE-FreqLimit.md)
  intent: A --top larger than the distinct count lists every value.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    c,3
    b,2
    a,1
- PASS: freq-limit-all-null-column (FEATURE-FreqLimit.md)
  intent: A column with no non-null values lists only the header even with --top given.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
- PASS: profile-output-columns (FEATURE-ProfileOutput.md)
  intent: The profile has the declared header and one row per column in file order, with types, counts, integer sum, and decimal sum and mean derived from the input.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    id,integer,3,0,3,1,4,7,2.33,2.00
    label,text,3,0,3,a,c,,,
    price,number,3,0,3,1.50,3.00,6.75,2.25,2.25
- PASS: profile-output-quoting (FEATURE-ProfileOutput.md)
  intent: A field containing a comma and a double quote is written quoted so the parsed value equals the trimmed input value.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    note,text,1,0,1,"x, ""y""","x, ""y""",,,
- PASS: profile-output-line-endings (FEATURE-ProfileOutput.md)
  intent: Output uses newline line endings with one line per column plus the header and a trailing newline.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    a,integer,2,0,2,1,2,3,1.50,1.50
    b,text,2,0,2,x,y,,,
    c,boolean,2,0,2,false,true,,,
- PASS: profile-output-header-only (FEATURE-ProfileOutput.md)
  intent: A header-only file yields one empty-typed row per column with zero rows.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    sku,empty,0,0,0,,,,,
    price,empty,0,0,0,,,,,
    in_stock,empty,0,0,0,,,,,
- PASS: profile-output-rounding (FEATURE-ProfileOutput.md)
  intent: Number results round half away from zero once, and a total or mean that rounds to zero carries no minus sign.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    pos,number,2,0,1,0.13,0.13,0.25,0.13,0.13
    neg,number,2,0,1,-0.13,-0.13,-0.25,-0.13,-0.13
    tiny,number,2,0,2,0.00,0.00,0.00,0.00,0.00
- PASS: profile-output-rejected-input (FEATURE-ProfileOutput.md)
  intent: A ragged file exits 3 and writes nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpc_n9cwju/ragged.csv: line 3: record has 3 fields, header has 2

## Build summary
I'll start by inspecting the build directory to see what already exists.The build directory already has earlier blocks in place. I'll read the existing packages before adding anything.The earlier blocks give me the parsing, validation, null, trim, and type-inference layers. This block adds the statistics in `internal/profile`, the frequency package, and the CLI output. I'll write the profile statistics first.Next, the profile tests. The named acceptance tests come first, followed by extra coverage for boundaries and the no-reorder rule.Now the frequency package, with its own tests.Now the CLI. The default invocation writes the profile, `--freq` writes the frequency listing, and output goes to stdout only after everything has been computed.Next, the CLI tests. They read the output back through `encoding/csv` instead of matching substrings.I spotted a bug in one of my test inputs. A bare blank line is skipped by `encoding/csv`, so that test needs a quoted empty field to produce a null row. Fixing that, then running `gofmt` and the suite.The `freq_test.go` change is just my own `gofmt` import sort, so that stays. The real failure is my test expectation: `RatString` renders a fraction as `a/b`, so the mean should be `500000000000000000015/2`, and the code is correct. Fixing the expectation.The gate is green: `go vet` and `go test ./...` pass across all packages, and 48 profile tests run. The Go-side acceptance assertions are covered by the named tests. The CLI-side assertions are Python scripts that depend on a `drydock_go` helper that isn't on disk. I'll confirm the named tests really ran rather than skipped, then reproduce those CLI checks with a shim kept outside the build directory.All 28 declared Go assertions pass as named tests. Next I'll check whether the real `drydock_go` helper exists in the Rigging tree before writing a stand-in.The real helper is shipped in Rigging. I'll run the CLI acceptance scripts against it, with the working directory set to the build directory so `cli_runner` builds the real binary. The scripts go in `/tmp`, outside the build tree.All 16 CLI acceptance scripts pass against the real binary, and the 28 named Go assertions pass. I'll confirm the changed-file list from git, without staging or committing anything.All eight stories in this block are implemented, and the declared acceptance checks pass.

- **Go-side assertions:** all 28 named tests in `internal/profile` pass when run individually through the same `-run` pattern the shipped helper uses.
- **CLI-side assertions:** all 16 `freq-count`, `freq-limit`, and `profile-output` checks pass. I ran them with the shipped `drydock_go.cli_runner` from Rigging, with the working directory set to the build directory. The driver lives in `/tmp/acrun/ac_cli.py`, outside the build tree.
- **Gate:** `sh bin/test.sh` (vet, then the full test suite) exits 0.

RESULT: SUCCESS

FILES CHANGED:
- internal/cli/cli.go
- internal/cli/profile.go
- internal/cli/profile_test.go
- internal/cli/freq.go
- internal/cli/freq_test.go
- internal/profile/counts.go
- internal/profile/counts_test.go
- internal/profile/numeric.go
- internal/profile/sum.go
- internal/profile/sum_test.go
- internal/profile/median.go
- internal/profile/median_test.go
- internal/profile/minmax.go
- internal/profile/minmax_test.go
- internal/profile/format.go
- internal/profile/format_test.go
- internal/freq/count.go
- internal/freq/count_test.go
- internal/freq/top.go
- internal/freq/top_test.go

SUMMARY:
- **Statistics (`internal/profile`):** `Counts`, `SumMean`, `Median`, `MinMax`, `FormatDecimal`, and `FormatInteger`. Sums, means, and medians use `big.Rat`. Each function ignores null values itself. `FormatDecimal` rounds once, half away from zero, on the exact rational, with no negative zero and no separator.
- **Frequencies (`internal/freq`):** `Count` lists trimmed non-null values ordered by count descending, then byte value. `Top` applies the limit after ordering.
- **CLI:** `Run` now takes stdout and builds all output in memory before writing it. `--freq` writes `value,count` and defaults to 5 rows. The default invocation writes the ten-column profile, with integer min, max, and sum as plain integers, and number statistics with two decimals. Output is `encoding/csv` with `\n` line endings. A rejected input still exits 3 with empty stdout.
- **Tests:** the CLI tests parse stdout with `encoding/csv` and compare cells, rather than matching substrings.

Choices I made where the specs were silent:
- **Unknown `--freq` column:** exits 2 with empty stdout.
- **`--top` below zero:** lists no rows. Rejecting it belongs to the usage-errors story, so I did not add validation here.
- **`--column` alone:** keeps its existing behaviour (exit 0, no output), because column output is outside this block.
- **`--freq` together with `--column`:** `--freq` takes precedence.

I did not run `sources/run.sh`. The Verification Protocol reserves it for the terminal story, and this block is a service block.

<blueprint-decisions>
[{"spec":"FEATURE-FreqCount.md","severity":"Low","subject":"Unknown --freq column name","decision":"The specs do not say what happens when the --freq column is not in the header. I exit 2 (usage) with empty stdout and a diagnostic on stderr. Exit 3 would also be defensible as rejected input. Is 2 acceptable, or should this be 3?"},{"spec":"FEATURE-FreqLimit.md","severity":"Low","subject":"Non-positive --top with --freq","decision":"Validation of --top belongs to the usage-errors story, which is not in this block. Until then --top 0 or a negative value lists no rows with exit 0. Confirm that the usage-errors story will reject these with exit 2."},{"spec":"FEATURE-ProfileOutput.md","severity":"Low","subject":"--column without --freq","decision":"Column output is not part of this block, so --column keeps the prior behaviour: exit 0 with no output. The profile is written only when neither --column nor --freq is given, as the spec says. Confirm this is acceptable until the column story lands."}]
</blueprint-decisions>
