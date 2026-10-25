# Evidence: Block 4 · Service (block-4)

- block type: block
- date: 2026-10-25
- resulting state: closed/verified
- story points (combined assembled cost): 13293
- execution id: 20261025.054944.501Z-bdb56842

## Stories built
- Reject malformed arguments with exit 2 and a usage line on standard error. (usage-errors) [story]
- Profile a single named column with --column. (column-profile) [story]
- Write one stderr diagnostic line and nothing on stdout when a command fails. (failure-diagnostics) [story]
- Exit 1 when the file cannot be opened or read. (unreadable-file) [story]
- Select frequency mode with --freq. (freq-mode) [story]

## Acceptance tooling authorization
- FEATURE-UsageErrors.md#usage-errors-exit-two: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-errors-single-stderr-line: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-errors-before-file-open: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UsageErrors.md#usage-errors-valid-forms-accepted: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnProfile.md#column-profile-single-row: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnProfile.md#column-profile-trimmed-header: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnProfile.md#column-profile-unknown-column: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnProfile.md#column-profile-case-sensitive: executable=go; scope=test; authorization=existing Target environment
- FEATURE-ColumnProfile.md#column-profile-rejected-input: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FailureDiagnostics.md#diag-ragged-fixture: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FailureDiagnostics.md#diag-ragged-line-number: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FailureDiagnostics.md#diag-empty-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FailureDiagnostics.md#diag-blank-lines-only: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FailureDiagnostics.md#diag-empty-header-name: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FailureDiagnostics.md#diag-repeated-header-name: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FailureDiagnostics.md#diag-failure-with-column-option: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UnreadableFile.md#unreadable-missing-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UnreadableFile.md#unreadable-missing-file-freq: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UnreadableFile.md#unreadable-missing-file-column: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UnreadableFile.md#unreadable-directory-as-file: executable=go; scope=test; authorization=existing Target environment
- FEATURE-UnreadableFile.md#unreadable-readable-file-still-succeeds: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqMode.md#freq-default-matches-fixture: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqMode.md#freq-top-limits-rows: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqMode.md#freq-omits-null-spellings: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqMode.md#freq-values-are-trimmed: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqMode.md#freq-all-null-column-header-only: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqMode.md#freq-unknown-column-rejected: executable=go; scope=test; authorization=existing Target environment
- FEATURE-FreqMode.md#freq-name-is-case-sensitive: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 963)
- implements: FEATURE-UsageErrors.md (SP 1544)
- context: ARCHITECTURE_compact.md (SP 282)
- implements: FEATURE-ColumnProfile.md (SP 1879)
- implements: FEATURE-FailureDiagnostics.md (SP 2150)
- implements: FEATURE-UnreadableFile.md (SP 1357)
- implements: FEATURE-FreqMode.md (SP 2706)

## Build directory changes
- internal/cli/args.go
- internal/cli/args_test.go
- internal/cli/cli.go
- internal/cli/column_test.go
- internal/cli/freq_test.go
- internal/cli/profile.go

## Pre-build acceptance observation
- RED: usage-errors-exit-two (FEATURE-UsageErrors.md)
  intent: Every malformed argument form exits 2 and writes nothing to standard output.
  return code: 1
  stdout:
    [] 2
    csvstat: missing file argument
    
    ['/tmp/tmpftcd3ltg/in.csv', '/tmp/tmpftcd3ltg/other.csv'] 2
    csvstat: unexpected argument "/tmp/tmpftcd3ltg/other.csv" after file or extra file
    
    ['--frobnicate', '/tmp/tmpftcd3ltg/in.csv'] 2
    csvstat: flag provided but not defined: -frobnicate
    
    ['--column'] 2
    csvstat: flag needs an argument: -column
    
    ['--freq'] 2
    csvstat: flag needs an argument: -freq
    
    ['--freq', 'sensor', '--top'] 2
    csvstat: flag needs an argument: -top
    
    ['--freq', 'sensor', '--top', '0', '/tmp/tmpftcd3ltg/in.csv'] 0
  stderr:
    --- drydock: values at failure ---
      code = 0
      args = ['--freq', 'sensor', '--top', '0', '/tmp/tmpftcd3ltg/in.csv']
    --- drydock: end values ---
    Traceback (most recent call last):
      File "usage-errors-exit-two.py", line 30, in <module>
        assert code == 2, (args, code)
               ^^^^^^^^^
    AssertionError: (['--freq', 'sensor', '--top', '0', '/tmp/tmpftcd3ltg/in.csv'], 0)
- GREEN (prepassed): usage-errors-single-stderr-line (FEATURE-UsageErrors.md)
  intent: A usage error writes exactly one non-blank diagnostic line to standard error.
  return code: 0
  stdout:
    csvstat: missing file argument
- RED: usage-errors-before-file-open (FEATURE-UsageErrors.md)
  intent: A usage error is reported as 2 even when the named file does not exist.
  return code: 1
  stdout:
    csvstat: open /tmp/tmputn6e6sx/no-such-file.csv: no such file or directory
  stderr:
    --- drydock: values at failure ---
      code = 1
    --- drydock: end values ---
    Traceback (most recent call last):
      File "usage-errors-before-file-open.py", line 9, in <module>
        assert code == 2, code
               ^^^^^^^^^
    AssertionError: 1
- GREEN (prepassed): usage-errors-valid-forms-accepted (FEATURE-UsageErrors.md)
  intent: Well-formed invocations are not rejected as usage errors.
  return code: 0
  stdout:
    ['/tmp/tmp6ss5r7fo/in.csv'] 0
    
    ['--column', 'qty', '/tmp/tmp6ss5r7fo/in.csv'] 0
    
    ['--freq', 'sensor', '/tmp/tmp6ss5r7fo/in.csv'] 0
    
    ['--freq', 'sensor', '--top', '2', '/tmp/tmp6ss5r7fo/in.csv'] 0
- RED: column-profile-single-row (FEATURE-ColumnProfile.md)
  intent: --column writes the profile header and exactly one row, equal to that column's row in the full profile.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      parsed = []
      full = [['column', 'type', 'rows', 'nulls', 'distinct', 'min', 'max', 'sum', 'mean', 'median'], ['id', 'integer', '3', '0', '3', '1', '4', '7', '2.33', '2.00'], ['price', 'number', '3', '0', '3', '1.50', '3.00', '6.75', '2.25', '2.25'], ['label', 'text', '3', '0', '3', 'a', 'c', '', '', '']]
    --- drydock: end values ---
    Traceback (most recent call last):
      File "column-profile-single-row.py", line 21, in <module>
        assert parsed[0] == full[0]
               ~~~~~~^^^
    IndexError: list index out of range
- RED: column-profile-trimmed-header (FEATURE-ColumnProfile.md)
  intent: A padded header name is trimmed, so the trimmed name selects the column and is reported.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      parsed = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "column-profile-trimmed-header.py", line 18, in <module>
        assert len(parsed) == 2
               ^^^^^^^^^^^^^^^^
    AssertionError
- RED: column-profile-unknown-column (FEATURE-ColumnProfile.md)
  intent: A column name absent from the header exits 3 with empty standard output.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 0
    --- drydock: end values ---
    Traceback (most recent call last):
      File "column-profile-unknown-column.py", line 13, in <module>
        assert code == 3, code
               ^^^^^^^^^
    AssertionError: 0
- RED: column-profile-case-sensitive (FEATURE-ColumnProfile.md)
  intent: Column names match case-sensitively, so a different case is an unknown column.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      code = 0
    --- drydock: end values ---
    Traceback (most recent call last):
      File "column-profile-case-sensitive.py", line 14, in <module>
        assert code == 3, code
               ^^^^^^^^^
    AssertionError: 0
- GREEN (prepassed): column-profile-rejected-input (FEATURE-ColumnProfile.md)
  intent: A ragged file with --column exits 3 and writes nothing to standard output.
  return code: 0
  stdout:
    csvstat: /tmp/tmpeuoy9ipj/ragged.csv: line 3: record has 3 fields, header has 2
- GREEN (prepassed): diag-ragged-fixture (FEATURE-FailureDiagnostics.md)
  intent: The supplied ragged fixture exits 3, writes nothing to standard output, and writes one diagnostic line.
  return code: 0
  stdout:
    csvstat: sources/ragged.csv: line 4: record has 4 fields, header has 3
- GREEN (prepassed): diag-ragged-line-number (FEATURE-FailureDiagnostics.md)
  intent: The diagnostic for a ragged record carries the number of the line holding that record.
  return code: 0
  stdout:
    csvstat: /tmp/tmpno0e759c/ragged_case.csv: line 4: record has 3 fields, header has 2
- GREEN (prepassed): diag-empty-file (FEATURE-FailureDiagnostics.md)
  intent: An empty file is rejected with exit 3 and nothing on standard output.
  return code: 0
  stdout:
    csvstat: /tmp/tmpn8m8f9bu/empty.csv: no header record: input is empty or blank
- GREEN (prepassed): diag-blank-lines-only (FEATURE-FailureDiagnostics.md)
  intent: A file containing only blank lines is rejected as empty with exit 3.
  return code: 0
  stdout:
    csvstat: /tmp/tmpslhzo9pj/blank.csv: no header record: input is empty or blank
- GREEN (prepassed): diag-empty-header-name (FEATURE-FailureDiagnostics.md)
  intent: An empty header name is rejected with exit 3 and nothing on standard output.
  return code: 0
  stdout:
    csvstat: /tmp/tmpkoag8lt7/emptyname.csv: header column 2 is empty
- GREEN (prepassed): diag-repeated-header-name (FEATURE-FailureDiagnostics.md)
  intent: A repeated header name, including one repeated only after trimming, is rejected with exit 3.
  return code: 0
  stdout:
    csvstat: /tmp/tmpfbats8s8/repeated.csv: header name "dup" is repeated
- GREEN (prepassed): diag-failure-with-column-option (FEATURE-FailureDiagnostics.md)
  intent: A rejected file under --column still exits 3 with nothing on standard output.
  return code: 0
  stdout:
    csvstat: sources/ragged.csv: line 4: record has 4 fields, header has 3
- GREEN (prepassed): unreadable-missing-file (FEATURE-UnreadableFile.md)
  intent: A path that does not exist exits 1 and writes nothing to standard output.
  return code: 0
  stdout:
    csvstat: open /tmp/tmpd_yhi93b/no-such-file.csv: no such file or directory
- GREEN (prepassed): unreadable-missing-file-freq (FEATURE-UnreadableFile.md)
  intent: A missing file under --freq also exits 1 with nothing on standard output.
  return code: 0
  stdout:
    csvstat: open /tmp/tmpe1w5vm97/absent.csv: no such file or directory
- GREEN (prepassed): unreadable-missing-file-column (FEATURE-UnreadableFile.md)
  intent: A missing file under --column also exits 1 with nothing on standard output.
  return code: 0
  stdout:
    csvstat: open /tmp/tmpn988au76/absent.csv: no such file or directory
- GREEN (prepassed): unreadable-directory-as-file (FEATURE-UnreadableFile.md)
  intent: A directory named as the file cannot be read and exits 1 with nothing on standard output.
  return code: 0
  stdout:
    csvstat: /tmp/tmp9rh9iyrs: read csv: read /tmp/tmp9rh9iyrs: is a directory
- GREEN (prepassed): unreadable-readable-file-still-succeeds (FEATURE-UnreadableFile.md)
  intent: A readable fixture is not treated as unreadable and exits 0.
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    sku,empty,0,0,0,,,,,
    price,empty,0,0,0,,,,,
    in_stock,empty,0,0,0,,,,,
- GREEN (prepassed): freq-default-matches-fixture (FEATURE-FreqMode.md)
  intent: The default frequency listing for a named column equals the top five counts computed independently from the fixture.
  return code: 0
  stdout:
    value,count
    Sales,8
    Engineering,7
    Support,7
    Finance,2
- GREEN (prepassed): freq-top-limits-rows (FEATURE-FreqMode.md)
  intent: With --top 2 exactly the first two ordered rows are listed.
  return code: 0
  stdout:
    value,count
    east,9
    north,9
- GREEN (prepassed): freq-omits-null-spellings (FEATURE-FreqMode.md)
  intent: Null values (empty, NA, N/A, NULL in any case) are not listed in a frequency listing.
  return code: 0
  stdout:
    value,count
    4,4
    2,3
    3,1
- GREEN (prepassed): freq-values-are-trimmed (FEATURE-FreqMode.md)
  intent: Padded values are trimmed before counting, so padded and unpadded spellings count together.
  return code: 0
  stdout:
    value,count
    Globex,10
    Acme Corp,9
    Umbrella,7
    Initech,3
    Hooli,1
- GREEN (prepassed): freq-all-null-column-header-only (FEATURE-FreqMode.md)
  intent: A column with no non-null values lists only the header.
  return code: 0
  stdout:
    value,count
- RED: freq-unknown-column-rejected (FEATURE-FreqMode.md)
  intent: A column name absent from the header exits 3 and writes nothing to standard output.
  return code: 1
  stdout:
    csvstat: no column named "nosuch"
  stderr:
    --- drydock: values at failure ---
      code = 2
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-unknown-column-rejected.py", line 6, in <module>
        assert code == 3, code
               ^^^^^^^^^
    AssertionError: 2
- RED: freq-name-is-case-sensitive (FEATURE-FreqMode.md)
  intent: A column name differing only in case from a header name is not found and exits 3.
  return code: 1
  stdout:
    csvstat: no column named "DEPARTMENT"
  stderr:
    --- drydock: values at failure ---
      code = 2
    --- drydock: end values ---
    Traceback (most recent call last):
      File "freq-name-is-case-sensitive.py", line 14, in <module>
        assert code == 3, code
               ^^^^^^^^^
    AssertionError: 2

## Post-build programmatic acceptance
- PASS: usage-errors-exit-two (FEATURE-UsageErrors.md)
  intent: Every malformed argument form exits 2 and writes nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    [] 2
    csvstat: missing file argument; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
    
    ['/tmp/tmptcxshsru/in.csv', '/tmp/tmptcxshsru/other.csv'] 2
    csvstat: unexpected argument "/tmp/tmptcxshsru/other.csv" after file or extra file; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
    
    ['--frobnicate', '/tmp/tmptcxshsru/in.csv'] 2
    csvstat: flag provided but not defined: -frobnicate; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
    
    ['--column'] 2
    csvstat: flag needs an argument: -column; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
    
    ['--freq'] 2
    csvstat: flag needs an argument: -freq; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
    
    ['--freq', 'sensor', '--top'] 2
    csvstat: flag needs an argument: -top; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
    
    ['--freq', 'sensor', '--top', '0', '/tmp/tmptcxshsru/in.csv'] 2
    csvstat: --top 0 is not a positive integer; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
    
    ['--freq', 'sensor', '--top', '-1', '/tmp/tmptcxshsru/in.csv'] 2
    csvstat: --top -1 is not a positive integer; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
    
    ['--freq', 'sensor', '--top', 'abc', '/tmp/tmptcxshsru/in.csv'] 2
    csvstat: invalid value "abc" for flag -top: "abc" is not a decimal integer; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
    
    ['--top', '3', '/tmp/tmptcxshsru/in.csv'] 2
    csvstat: --top requires --freq; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
    
    ['--column', 'qty', '--freq', 'qty', '/tmp/tmptcxshsru/in.csv'] 2
    csvstat: --column cannot be combined with --freq; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
- PASS: usage-errors-single-stderr-line (FEATURE-UsageErrors.md)
  intent: A usage error writes exactly one non-blank diagnostic line to standard error.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: missing file argument; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
- PASS: usage-errors-before-file-open (FEATURE-UsageErrors.md)
  intent: A usage error is reported as 2 even when the named file does not exist.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: --top requires --freq; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>
- PASS: usage-errors-valid-forms-accepted (FEATURE-UsageErrors.md)
  intent: Well-formed invocations are not rejected as usage errors.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ['/tmp/tmpsctabn7b/in.csv'] 0
    
    ['--column', 'qty', '/tmp/tmpsctabn7b/in.csv'] 0
    
    ['--freq', 'sensor', '/tmp/tmpsctabn7b/in.csv'] 0
    
    ['--freq', 'sensor', '--top', '2', '/tmp/tmpsctabn7b/in.csv'] 0
- PASS: column-profile-single-row (FEATURE-ColumnProfile.md)
  intent: --column writes the profile header and exactly one row, equal to that column's row in the full profile.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    id,integer,3,0,3,1,4,7,2.33,2.00
    
    
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    price,number,3,0,3,1.50,3.00,6.75,2.25,2.25
    
    
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    label,text,3,0,3,a,c,,,
- PASS: column-profile-trimmed-header (FEATURE-ColumnProfile.md)
  intent: A padded header name is trimmed, so the trimmed name selects the column and is reported.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    qty,integer,2,0,2,1,2,3,1.50,1.50
- PASS: column-profile-unknown-column (FEATURE-ColumnProfile.md)
  intent: A column name absent from the header exits 3 with empty standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmp6i6ckalr/in.csv: no column named "nosuch"
- PASS: column-profile-case-sensitive (FEATURE-ColumnProfile.md)
  intent: Column names match case-sensitively, so a different case is an unknown column.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpwp9zjwxb/in.csv: no column named "QTY"
- PASS: column-profile-rejected-input (FEATURE-ColumnProfile.md)
  intent: A ragged file with --column exits 3 and writes nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmp_51zz7ki/ragged.csv: line 3: record has 3 fields, header has 2
- PASS: diag-ragged-fixture (FEATURE-FailureDiagnostics.md)
  intent: The supplied ragged fixture exits 3, writes nothing to standard output, and writes one diagnostic line.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: sources/ragged.csv: line 4: record has 4 fields, header has 3
- PASS: diag-ragged-line-number (FEATURE-FailureDiagnostics.md)
  intent: The diagnostic for a ragged record carries the number of the line holding that record.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmph618i3f8/ragged_case.csv: line 4: record has 3 fields, header has 2
- PASS: diag-empty-file (FEATURE-FailureDiagnostics.md)
  intent: An empty file is rejected with exit 3 and nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmp0ljanamc/empty.csv: no header record: input is empty or blank
- PASS: diag-blank-lines-only (FEATURE-FailureDiagnostics.md)
  intent: A file containing only blank lines is rejected as empty with exit 3.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpstrfntbn/blank.csv: no header record: input is empty or blank
- PASS: diag-empty-header-name (FEATURE-FailureDiagnostics.md)
  intent: An empty header name is rejected with exit 3 and nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpv4uvv38u/emptyname.csv: header column 2 is empty
- PASS: diag-repeated-header-name (FEATURE-FailureDiagnostics.md)
  intent: A repeated header name, including one repeated only after trimming, is rejected with exit 3.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmpdrugtg3a/repeated.csv: header name "dup" is repeated
- PASS: diag-failure-with-column-option (FEATURE-FailureDiagnostics.md)
  intent: A rejected file under --column still exits 3 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: sources/ragged.csv: line 4: record has 4 fields, header has 3
- PASS: unreadable-missing-file (FEATURE-UnreadableFile.md)
  intent: A path that does not exist exits 1 and writes nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: open /tmp/tmply63mrxu/no-such-file.csv: no such file or directory
- PASS: unreadable-missing-file-freq (FEATURE-UnreadableFile.md)
  intent: A missing file under --freq also exits 1 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: open /tmp/tmp6nqlag7r/absent.csv: no such file or directory
- PASS: unreadable-missing-file-column (FEATURE-UnreadableFile.md)
  intent: A missing file under --column also exits 1 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: open /tmp/tmpgh5gw1hw/absent.csv: no such file or directory
- PASS: unreadable-directory-as-file (FEATURE-UnreadableFile.md)
  intent: A directory named as the file cannot be read and exits 1 with nothing on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: /tmp/tmphq_dcnzw: read csv: read /tmp/tmphq_dcnzw: is a directory
- PASS: unreadable-readable-file-still-succeeds (FEATURE-UnreadableFile.md)
  intent: A readable fixture is not treated as unreadable and exits 0.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    column,type,rows,nulls,distinct,min,max,sum,mean,median
    sku,empty,0,0,0,,,,,
    price,empty,0,0,0,,,,,
    in_stock,empty,0,0,0,,,,,
- PASS: freq-default-matches-fixture (FEATURE-FreqMode.md)
  intent: The default frequency listing for a named column equals the top five counts computed independently from the fixture.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    Sales,8
    Engineering,7
    Support,7
    Finance,2
- PASS: freq-top-limits-rows (FEATURE-FreqMode.md)
  intent: With --top 2 exactly the first two ordered rows are listed.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    east,9
    north,9
- PASS: freq-omits-null-spellings (FEATURE-FreqMode.md)
  intent: Null values (empty, NA, N/A, NULL in any case) are not listed in a frequency listing.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    4,4
    2,3
    3,1
- PASS: freq-values-are-trimmed (FEATURE-FreqMode.md)
  intent: Padded values are trimmed before counting, so padded and unpadded spellings count together.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
    Globex,10
    Acme Corp,9
    Umbrella,7
    Initech,3
    Hooli,1
- PASS: freq-all-null-column-header-only (FEATURE-FreqMode.md)
  intent: A column with no non-null values lists only the header.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    value,count
- PASS: freq-unknown-column-rejected (FEATURE-FreqMode.md)
  intent: A column name absent from the header exits 3 and writes nothing to standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: sources/orders.csv: no column named "nosuch"
- PASS: freq-name-is-case-sensitive (FEATURE-FreqMode.md)
  intent: A column name differing only in case from a header name is not found and exits 3.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: sources/employees.csv: no column named "DEPARTMENT"

## Build summary
The baseline gate is green, but several declared behaviors are missing: `--column` is a stub that writes nothing, `--top`/`--column`/`--freq` combinations aren't rejected, and an unknown `--freq` column exits 2 instead of 3. I'll make those changes now.Now I'll update the one existing test that encoded the old exit-2 behavior for an unknown `--freq` column, and add tests for the new paths.I accidentally wrote a near-duplicate freq test; removing it before running the suite.Gate is green. Next I'll exercise the built binary against the declared acceptance cases directly, since the Python `cli_runner` harness isn't available here.RESULT: SUCCESS

The five stories in this block are implemented. `sh bin/test.sh` (vet, then the full Go test suite) passes. I did not run the Python acceptance criteria through `drydock_go.cli_runner`, because that helper isn't on disk here. I checked the same cases by running the built binary directly, and they match the declared exit codes and stdout/stderr shape.

FILES CHANGED:
- internal/cli/args.go
- internal/cli/cli.go
- internal/cli/profile.go
- internal/cli/args_test.go
- internal/cli/freq_test.go
- internal/cli/column_test.go (new)

SUMMARY:
- **usage-errors:** `ParseArgs` now rejects `--column` with `--freq`, `--top` without `--freq`, and a non-positive `--top`, in addition to the existing checks. Every usage error prints one stderr line that includes both command forms and exits 2 before the file is opened.
- **column-profile:** `--column` was a stub that exited 0 with no output. It now writes the profile header and only that column's row, matched exactly and case-sensitively against the trimmed header. An absent name exits 3.
- **freq-mode:** An unknown `--freq` column now exits 3 instead of 2, as the feature spec requires. I changed the existing test that asserted exit 2 to assert exit 3, and added case-sensitivity coverage.
- **failure-diagnostics / unreadable-file:** Already behaving as specified. I verified that rejected input exits 3, unreadable or missing files and directories exit 1, and stdout stays empty on every failure.

Direct checks on the binary:
- All 11 malformed-argument forms exit 2 with empty stdout.
- Valid forms exit 0.
- Unknown and wrong-case column names exit 3.
- Missing file, directory, `--freq` on a missing file, and `--column` on a missing file exit 1.
- The ragged fixture exits 3 with the line number in the diagnostic.
- `sources/` is untouched, and I did not run `sources/run.sh`.

Choices I made:
- The usage diagnostic is a single line: `csvstat: <reason>; usage: csvstat [--column <name>] <file> | csvstat --freq <name> [--top <n>] <file>`.
- `--column` output reuses the profile row builder, so it is guaranteed to equal the full profile's row for that column.
