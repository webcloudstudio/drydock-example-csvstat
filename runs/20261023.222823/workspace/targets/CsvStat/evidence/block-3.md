# Evidence: Block 3 · Service (block-3)

- block type: block
- date: 2026-10-23
- resulting state: closed/verified
- story points (combined assembled cost): 13788
- execution id: 20261023.223704.536Z-43b6712a

## Stories built
- Read, normalize, and structurally validate CSV input. (cli-input) [story]
- Resolve named columns and classify null values. (cli-columns) [story]
- Infer each column type using the specified precedence. (profile-types) [story]
- Aggregate and order frequent non-null column values. (frequency-values) [story]
- Validate and apply frequency row limits. (frequency-limits) [story]
- Compute exact profile statistics and decimal formatting. (profile-statistics) [story]
- Render profile results as CSV. (profile-output) [story]
- Map operational, usage, and rejected-input failures. (cli-errors) [story]
- Provide the Go vet and test project gate. (test-gate) [story]
- Run the complete supplied CsvStat conformance suite. (full-verification) [story]

## Stacked context
- stack: go_compact.md (SP 1374)
- compass: COMPASS.md (SP 628)
- implements: FEATURE-CLI-INPUT.md (SP 552)
- context: employees.csv (SP 364)
- context: sensors.csv (SP 394)
- context: orders.csv (SP 413)
- context: header_only.csv (SP 5)
- context: ragged.csv (SP 20)
- context: ARCHITECTURE_compact.md (SP 146)
- implements: FEATURE-CLI-COLUMNS.md (SP 550)
- implements: FEATURE-PROFILE-TYPES.md (SP 578)
- implements: FEATURE-FREQUENCY-VALUES.md (SP 731)
- implements: FEATURE-FREQUENCY-LIMITS.md (SP 774)
- implements: FEATURE-PROFILE-STATISTICS.md (SP 1045)
- implements: FEATURE-PROFILE-OUTPUT.md (SP 729)
- implements: FEATURE-CLI-ERRORS.md (SP 793)
- implements: FEATURE-TEST-GATE.md (SP 415)
- stack: common.md (SP 1812)
- implements: FEATURE-FULL-VERIFICATION.md (SP 399)
- context: run.sh (SP 1471)

## Build directory changes
- csvstat
- csvstat-ac
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvstat/csvstat.go

## Pre-build acceptance observation
- RED: cli-input-quoted (FEATURE-CLI-INPUT.md)
  intent: A supplied CSV containing quoted commas is accepted.
  return code: 1
  stdout:
    csvstat: profiling "sources/employees.csv" is not configured
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat-ac', 'sources/employees.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/employees.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-input-quoted.py", line 12, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: cli-input-ragged (FEATURE-CLI-INPUT.md)
  intent: A ragged CSV is rejected without standard output.
  return code: 1
  stdout:
    csvstat: profiling "sources/ragged.csv" is not configured
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat-ac', 'sources/ragged.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/ragged.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-input-ragged.py", line 12, in <module>
        assert result.returncode == 3
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: cli-input-header-only (FEATURE-CLI-INPUT.md)
  intent: A header-only CSV is accepted as a valid empty data set.
  return code: 1
  stdout:
    csvstat: profiling "sources/header_only.csv" is not configured
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat-ac', 'sources/header_only.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/header_only.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-input-header-only.py", line 12, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: cli-columns-existing (FEATURE-CLI-COLUMNS.md)
  intent: A named existing column is accepted for profiling.
  return code: 1
  stdout:
    csvstat: profiling "sources/employees.csv" is not configured
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat-ac', '--column', 'salary', 'sources/employees.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/employees.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-columns-existing.py", line 12, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: cli-columns-missing-profile (FEATURE-CLI-COLUMNS.md)
  intent: A missing profile column is rejected without standard output.
  return code: 1
  stdout:
    csvstat: profiling "sources/orders.csv" is not configured
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat-ac', '--column', 'nosuch', 'sources/orders.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/orders.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-columns-missing-profile.py", line 12, in <module>
        assert result.returncode == 3
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: cli-columns-missing-frequency (FEATURE-CLI-COLUMNS.md)
  intent: A missing frequency column is rejected without standard output.
  return code: 1
  stdout:
    csvstat: profiling "sources/orders.csv" is not configured
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat-ac', '--freq', 'nosuch', 'sources/orders.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/orders.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-columns-missing-frequency.py", line 12, in <module>
        assert result.returncode == 3
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: profile-types-fixtures (FEATURE-PROFILE-TYPES.md)
  intent: The supplied profile fixtures are accepted and produce profile rows.
  return code: 1
  stdout:
    csvstat: profiling "sources/employees.csv" is not configured
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat-ac', 'sources/employees.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/employees.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-types-fixtures.py", line 12, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: profile-types-all-null (FEATURE-PROFILE-TYPES.md)
  intent: An all-null column remains a valid profile input.
  return code: 1
  stdout:
    csvstat: profiling "sources/sensors.csv" is not configured
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat-ac', '--column', 'battery', 'sources/sensors.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/sensors.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-types-all-null.py", line 11, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: profile-types-invalid-date-text (FEATURE-PROFILE-TYPES.md)
  intent: A column containing an invalid calendar date remains valid input rather than causing command failure.
  return code: 1
  stdout:
    csvstat: profiling "sources/orders.csv" is not configured
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat-ac', '--column', 'shipped', 'sources/orders.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/orders.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-types-invalid-date-text.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: frequency-excludes-nulls (FEATURE-FREQUENCY-VALUES.md)
  intent: Frequency output omits every recognized null spelling.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "frequency-excludes-nulls.py", line 11, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: frequency-orders-counts (FEATURE-FREQUENCY-VALUES.md)
  intent: Frequency rows are ordered by descending count before value ordering.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "frequency-orders-counts.py", line 12, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: frequency-orders-ties (FEATURE-FREQUENCY-VALUES.md)
  intent: Equal-frequency values are ordered by ascending byte value.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "frequency-orders-ties.py", line 11, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: frequency-limit-success (FEATURE-FREQUENCY-LIMITS.md)
  intent: A valid positive limit executes the frequency command successfully and bounds the parsed result.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat', '--freq', 'department', '--top', '2', 'sources/employees.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/employees.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "frequency-limit-success.py", line 12, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: frequency-limit-default (FEATURE-FREQUENCY-LIMITS.md)
  intent: Omitting --top uses the default limit and returns a successful frequency result.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat', '--freq', 'department', 'sources/employees.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/employees.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "frequency-limit-default.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: frequency-limit-invalid (FEATURE-FREQUENCY-LIMITS.md)
  intent: Non-positive and malformed limits are rejected as usage errors.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat', '--freq', 'department', '--top', '0', 'sources/employees.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/employees.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "frequency-limit-invalid.py", line 10, in <module>
        assert result.returncode == 2
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: top-requires-frequency (FEATURE-FREQUENCY-LIMITS.md)
  intent: --top without --freq is rejected as a usage error.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat', '--top', '3', 'sources/employees.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/employees.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "top-requires-frequency.py", line 8, in <module>
        assert result.returncode == 2
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: profile-statistics-integer (FEATURE-PROFILE-STATISTICS.md)
  intent: Integer statistics include counts, extrema, exact total, mean, and median.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-statistics-integer.py", line 17, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: profile-statistics-decimal (FEATURE-PROFILE-STATISTICS.md)
  intent: Decimal statistics preserve exact values and round mean and median to two places.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-statistics-decimal.py", line 13, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: profile-statistics-nonnumeric (FEATURE-PROFILE-STATISTICS.md)
  intent: Non-numeric columns retain counts and extrema while leaving numeric fields empty.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-statistics-nonnumeric.py", line 11, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: profile-output-all-columns (FEATURE-PROFILE-OUTPUT.md)
  intent: A profile command emits one output row for each input column in file order.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-output-all-columns.py", line 12, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: profile-output-selected-column (FEATURE-PROFILE-OUTPUT.md)
  intent: The column option emits exactly the requested profile row.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-output-selected-column.py", line 12, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- RED: profile-output-quoting (FEATURE-PROFILE-OUTPUT.md)
  intent: Profile output preserves comma-containing values through CSV parsing.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "profile-output-quoting.py", line 12, in <module>
        assert code == 0
               ^^^^^^^^^
    AssertionError
- GREEN (prepassed): operational-error (FEATURE-CLI-ERRORS.md)
  intent: An unreadable input file returns the operational failure status and no CSV output.
  return code: 0
- GREEN (prepassed): usage-errors (FEATURE-CLI-ERRORS.md)
  intent: Malformed command arguments return usage status and no CSV output.
  return code: 0
- RED: rejected-input (FEATURE-CLI-ERRORS.md)
  intent: Structurally invalid input and absent columns return rejected-input status and no CSV output.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat', 'sources/ragged.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/ragged.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "rejected-input.py", line 10, in <module>
        assert result.returncode == 3
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: successful-output-isolated (FEATURE-CLI-ERRORS.md)
  intent: A valid command remains successful and emits CSV on standard output.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['./csvstat', '--column', 'qty', 'sources/orders.csv'], returncode=1, stdout='', stderr='csvstat: profiling "sources/orders.csv" is not configured\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "successful-output-isolated.py", line 8, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- GREEN (prepassed): test-gate-present (FEATURE-TEST-GATE.md)
  intent: The project provides an executable POSIX test launcher.
  return code: 0
- GREEN (prepassed): test-gate-runs (FEATURE-TEST-GATE.md)
  intent: The project test launcher completes successfully.
  return code: 0
  stdout:
    ?   	csvstat/cmd/csvstat	[no test files]
    ok  	csvstat/internal/cli	(cached)
    ok  	csvstat/internal/csvstat	(cached)
- GREEN (prepassed): go-vet-and-tests (FEATURE-TEST-GATE.md)
  intent: The launcher exercises both the Go vet and Go test commands.
  return code: 0
- RED: full-conformance (FEATURE-FULL-VERIFICATION.md)
  intent: The supplied conformance suite passes for the completed CsvStat implementation.
  return code: 1
  stdout:
    FAIL A1: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv
      exit: expected 0, got 1
      1,9d0
      < column,type,rows,nulls,distinct,min,max,sum,mean,median
      < id,integer,24,0,24,1001,1024,24300,1012.50,1012.50
      < name,text,24,0,23,Ada,Tony,,,
      < department,text,24,0,4,Engineering,Support,,,
      < hire_date,date,24,0,24,2015-08-27,2025-06-02,,,
      < salary,number,24,0,24,43699.25,147150.99,2085635.10,86901.46,79496.68
      < remote,boolean,24,0,4,false,true,,,
      < manager_id,integer,24,15,4,2,5,30,3.33,4.00
      < notes,text,24,10,4,mentor,"remote, part-time",,,
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv" is not configured
    FAIL A2: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
      exit: expected 0, got 1
      1,8d0
      < column,type,rows,nulls,distinct,min,max,sum,mean,median
      < reading_id,integer,36,0,36,1,36,666,18.50,18.50
      < recorded_at,text,36,0,36,2025-03-01T08:00,2025-03-03T19:00,,,
      < sensor,text,36,0,4,east,west,,,
      < temp_c,number,36,3,32,-12.21,29.81,207.06,6.27,2.41
      < humidity,integer,36,7,25,18,92,1698,58.55,60.00
      < ok,boolean,36,0,2,false,true,,,
      < battery,empty,36,36,0,,,,,
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv" is not configured
    FAIL A3: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv
      exit: expected 0, got 1
      1,9d0
      < column,type,rows,nulls,distinct,min,max,sum,mean,median
      < order_id,integer,30,0,30,5000,5029,150435,5014.50,5014.50
      < order_date,date,30,0,19,2025-02-01,2025-02-27,,,
      < customer,text,30,0,5,Acme Corp,Umbrella,,,
      < qty,integer,30,0,9,1,12,206,6.87,6.00
      < unit_price,number,30,0,30,15.27,231.18,4177.97,139.27,131.31
      < coupon,text,30,18,2,SAVE10,WELCOME,,,
      < ref,text,30,0,5,-3,A7,,,
      < shipped,text,30,0,20,2025-02-30,2025-03-27,,,
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv" is not configured
    FAIL A4: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/header_only.csv
      exit: expected 0, got 1
      1,4d0
      < column,type,rows,nulls,distinct,min,max,sum,mean,median
      < sku,empty,0,0,0,,,,,
      < price,empty,0,0,0,,,,,
      < in_stock,empty,0,0,0,,,,,
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/header_only.csv" is not configured
    FAIL B1: csvstat --column salary /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv
      exit: expected 0, got 1
      1,2d0
      < column,type,rows,nulls,distinct,min,max,sum,mean,median
      < salary,number,24,0,24,43699.25,147150.99,2085635.10,86901.46,79496.68
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv" is not configured
    FAIL B2: csvstat --column notes /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv
      exit: expected 0, got 1
      1,2d0
      < column,type,rows,nulls,distinct,min,max,sum,mean,median
      < notes,text,24,10,4,mentor,"remote, part-time",,,
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv" is not configured
    FAIL B3: csvstat --column temp_c /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
      exit: expected 0, got 1
      1,2d0
      < column,type,rows,nulls,distinct,min,max,sum,mean,median
      < temp_c,number,36,3,32,-12.21,29.81,207.06,6.27,2.41
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv" is not configured
    FAIL C1: csvstat --freq department /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv
      exit: expected 0, got 1
      1,5d0
      < value,count
      < Sales,8
      < Engineering,7
      < Support,7
      < Finance,2
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv" is not configured
    FAIL C2: csvstat --freq sensor --top 2 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
      exit: expected 0, got 1
      1,3d0
      < value,count
      < east,9
      < north,9
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv" is not configured
    FAIL C3: csvstat --freq customer /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv
      exit: expected 0, got 1
      1,6d0
      < value,count
      < Globex,10
      < Acme Corp,9
      < Umbrella,7
      < Initech,3
      < Hooli,1
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv" is not configured
    FAIL C4: csvstat --freq manager_id --top 3 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv
      exit: expected 0, got 1
      1,4d0
      < value,count
      < 4,4
      < 2,3
      < 3,1
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv" is not configured
    FAIL C5: csvstat --freq battery /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
      exit: expected 0, got 1
      1d0
      < value,count
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv" is not configured
    FAIL E1: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/ragged.csv
      exit: expected 3, got 1
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/ragged.csv" is not configured
    FAIL E2: csvstat --column nosuch /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv
      exit: expected 3, got 1
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv" is not configured
    FAIL E3: csvstat --freq nosuch /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv
      exit: expected 3, got 1
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv" is not configured
    PASS E4: csvstat /tmp/tmp.yoy25k1rNj/no-such-file.csv
    PASS U1: csvstat 
    PASS U2: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
    FAIL U3: csvstat --freq sensor --top 0 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
      exit: expected 2, got 1
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv" is not configured
    FAIL U4: csvstat --top 3 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
      exit: expected 2, got 1
      stderr: csvstat: profiling "/mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv" is not configured
    PASS U5: csvstat --column qty --freq qty /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv
    PASS U6: csvstat --frobnicate /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv
    CsvStat: 5 passed, 17 failed (of 22)
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['sh', 'sources/run.sh'], returncode=1, stdout='FAIL A1: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv\n  exit: expected 0, got 1\n  1,9d0\n  < column,type,rows,nulls,distinct,min,max,sum,mean,median\n  < id,integer,24,0,24,1001,1024,24300,1012.50,1012.50\n  < name,text,24,0,23,Ada,Tony,,,\n  < department,text… (+7861 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "full-conformance.py", line 12, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: cli-input-quoted (FEATURE-CLI-INPUT.md)
  intent: A supplied CSV containing quoted commas is accepted.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    id,integer,24,0,24,1001,1024,24300,1012.50,1012.50
    name,text,24,0,23,Ada,Tony,,,
    department,text,24,0,4,Engineering,Support,,,
    hire_date,date,24,0,24,2015-08-27,2025-06-02,,,
    salary,number,24,0,24,43699.25,147150.99,2085635.10,86901.46,79496.68
    remote,boolean,24,0,4,false,true,,,
    manager_id,integer,24,15,4,2,5,30,3.33,4.00
    notes,text,24,10,4,mentor,"remote, part-time",,,
- PASS: cli-input-ragged (FEATURE-CLI-INPUT.md)
  intent: A ragged CSV is rejected without standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: sources/ragged.csv: record line 4 has 4 fields, want 3
- PASS: cli-input-header-only (FEATURE-CLI-INPUT.md)
  intent: A header-only CSV is accepted as a valid empty data set.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    sku,empty,0,0,0,,,,,
    price,empty,0,0,0,,,,,
    in_stock,empty,0,0,0,,,,,
- PASS: cli-columns-existing (FEATURE-CLI-COLUMNS.md)
  intent: A named existing column is accepted for profiling.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    salary,number,24,0,24,43699.25,147150.99,2085635.10,86901.46,79496.68
- PASS: cli-columns-missing-profile (FEATURE-CLI-COLUMNS.md)
  intent: A missing profile column is rejected without standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: sources/orders.csv: column "nosuch" not found
- PASS: cli-columns-missing-frequency (FEATURE-CLI-COLUMNS.md)
  intent: A missing frequency column is rejected without standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: sources/orders.csv: column "nosuch" not found
- PASS: profile-types-fixtures (FEATURE-PROFILE-TYPES.md)
  intent: The supplied profile fixtures are accepted and produce profile rows.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    id,integer,24,0,24,1001,1024,24300,1012.50,1012.50
    name,text,24,0,23,Ada,Tony,,,
    department,text,24,0,4,Engineering,Support,,,
    hire_date,date,24,0,24,2015-08-27,2025-06-02,,,
    salary,number,24,0,24,43699.25,147150.99,2085635.10,86901.46,79496.68
    remote,boolean,24,0,4,false,true,,,
    manager_id,integer,24,15,4,2,5,30,3.33,4.00
    notes,text,24,10,4,mentor,"remote, part-time",,,
    
    
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    reading_id,integer,36,0,36,1,36,666,18.50,18.50
    recorded_at,text,36,0,36,2025-03-01T08:00,2025-03-03T19:00,,,
    sensor,text,36,0,4,east,west,,,
    temp_c,number,36,3,32,-12.21,29.81,207.06,6.27,2.41
    humidity,integer,36,7,25,18,92,1698,58.55,60.00
    ok,boolean,36,0,2,false,true,,,
    battery,empty,36,36,0,,,,,
    
    
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    order_id,integer,30,0,30,5000,5029,150435,5014.50,5014.50
    order_date,date,30,0,19,2025-02-01,2025-02-27,,,
    customer,text,30,0,5,Acme Corp,Umbrella,,,
    qty,integer,30,0,9,1,12,206,6.87,6.00
    unit_price,number,30,0,30,15.27,231.18,4177.97,139.27,131.31
    coupon,text,30,18,2,SAVE10,WELCOME,,,
    ref,text,30,0,5,-3,A7,,,
    shipped,text,30,0,20,2025-02-30,2025-03-27,,,
    
    
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    sku,empty,0,0,0,,,,,
    price,empty,0,0,0,,,,,
    in_stock,empty,0,0,0,,,,,
- PASS: profile-types-all-null (FEATURE-PROFILE-TYPES.md)
  intent: An all-null column remains a valid profile input.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    battery,empty,36,36,0,,,,,
- PASS: profile-types-invalid-date-text (FEATURE-PROFILE-TYPES.md)
  intent: A column containing an invalid calendar date remains valid input rather than causing command failure.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    shipped,text,30,0,20,2025-02-30,2025-03-27,,,
- PASS: frequency-excludes-nulls (FEATURE-FREQUENCY-VALUES.md)
  intent: Frequency output omits every recognized null spelling.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: frequency-orders-counts (FEATURE-FREQUENCY-VALUES.md)
  intent: Frequency rows are ordered by descending count before value ordering.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: frequency-orders-ties (FEATURE-FREQUENCY-VALUES.md)
  intent: Equal-frequency values are ordered by ascending byte value.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: frequency-limit-success (FEATURE-FREQUENCY-LIMITS.md)
  intent: A valid positive limit executes the frequency command successfully and bounds the parsed result.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: frequency-limit-default (FEATURE-FREQUENCY-LIMITS.md)
  intent: Omitting --top uses the default limit and returns a successful frequency result.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: frequency-limit-invalid (FEATURE-FREQUENCY-LIMITS.md)
  intent: Non-positive and malformed limits are rejected as usage errors.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: top-requires-frequency (FEATURE-FREQUENCY-LIMITS.md)
  intent: --top without --freq is rejected as a usage error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: profile-statistics-integer (FEATURE-PROFILE-STATISTICS.md)
  intent: Integer statistics include counts, extrema, exact total, mean, and median.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: profile-statistics-decimal (FEATURE-PROFILE-STATISTICS.md)
  intent: Decimal statistics preserve exact values and round mean and median to two places.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: profile-statistics-nonnumeric (FEATURE-PROFILE-STATISTICS.md)
  intent: Non-numeric columns retain counts and extrema while leaving numeric fields empty.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: profile-output-all-columns (FEATURE-PROFILE-OUTPUT.md)
  intent: A profile command emits one output row for each input column in file order.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: profile-output-selected-column (FEATURE-PROFILE-OUTPUT.md)
  intent: The column option emits exactly the requested profile row.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: profile-output-quoting (FEATURE-PROFILE-OUTPUT.md)
  intent: Profile output preserves comma-containing values through CSV parsing.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: operational-error (FEATURE-CLI-ERRORS.md)
  intent: An unreadable input file returns the operational failure status and no CSV output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: usage-errors (FEATURE-CLI-ERRORS.md)
  intent: Malformed command arguments return usage status and no CSV output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: rejected-input (FEATURE-CLI-ERRORS.md)
  intent: Structurally invalid input and absent columns return rejected-input status and no CSV output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: successful-output-isolated (FEATURE-CLI-ERRORS.md)
  intent: A valid command remains successful and emits CSV on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: test-gate-present (FEATURE-TEST-GATE.md)
  intent: The project provides an executable POSIX test launcher.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: test-gate-runs (FEATURE-TEST-GATE.md)
  intent: The project test launcher completes successfully.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ?   	csvstat/cmd/csvstat	[no test files]
    ok  	csvstat/internal/cli	(cached)
    ok  	csvstat/internal/csvstat	(cached)
- PASS: go-vet-and-tests (FEATURE-TEST-GATE.md)
  intent: The launcher exercises both the Go vet and Go test commands.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: full-conformance (FEATURE-FULL-VERIFICATION.md)
  intent: The supplied conformance suite passes for the completed CsvStat implementation.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    PASS A1: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv
    PASS A2: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
    PASS A3: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv
    PASS A4: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/header_only.csv
    PASS B1: csvstat --column salary /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv
    PASS B2: csvstat --column notes /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv
    PASS B3: csvstat --column temp_c /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
    PASS C1: csvstat --freq department /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv
    PASS C2: csvstat --freq sensor --top 2 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
    PASS C3: csvstat --freq customer /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv
    PASS C4: csvstat --freq manager_id --top 3 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/employees.csv
    PASS C5: csvstat --freq battery /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
    PASS E1: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/ragged.csv
    PASS E2: csvstat --column nosuch /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv
    PASS E3: csvstat --freq nosuch /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv
    PASS E4: csvstat /tmp/tmp.Tw1TfnTGIP/no-such-file.csv
    PASS U1: csvstat 
    PASS U2: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
    PASS U3: csvstat --freq sensor --top 0 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
    PASS U4: csvstat --top 3 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/sensors.csv
    PASS U5: csvstat --column qty --freq qty /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv
    PASS U6: csvstat --frobnicate /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261023.222823/build/CsvStat/sources/orders.csv
    CsvStat: 22 passed, 0 failed (of 22)

## Repair attempts
- attempt 0 (initial build): failed; 15/30 checks; 22/22 cases model=gpt-5.6-luna; execution 20261023.223528.933Z-cb51eaa7; reason: programmatic acceptance failed: cli-input-quoted, cli-input-ragged, cli-input-header-only, cli-columns-existing, cli-columns-missing-profile, cli-columns-missi…
- attempt 1 (repair 1): built; 30/30 checks; 22/22 cases model=gpt-5.6-luna; execution 20261023.223704.536Z-43b6712a

## Build summary
RESULT: SUCCESS

FILES CHANGED:
- csvstat
- csvstat-ac

SUMMARY:
Refreshed stale executables from the current Go implementation. `bin/test.sh` passes, and the full CsvStat conformance suite passes all 22 checks.

BLOCKERS:
- None
