# Evidence: Block 2 · Service (block-2)

- block type: block
- date: 2026-10-22
- resulting state: closed/failed
- story points (combined assembled cost): 5126
- execution id: 20261022.203854.435Z-44544656

## Stories built
- Parse csvstat commands, options, arguments, and exit-code categories. (cli-contract) [story]
- Provide usable help and failure diagnostics for csvstat invocations. (cli-help) [story]

## Stacked context
- stack: go_compact.md (SP 1173)
- compass: COMPASS.md (SP 628)
- implements: FEATURE-CLI-CONTRACT.md (SP 794)
- context: ARCHITECTURE_compact.md (SP 125)
- stack: common.md (SP 1812)
- implements: FEATURE-CLI-HELP.md (SP 464)

## Build directory changes
- internal/cli/cli.go
- internal/cli/cli_test.go

## Pre-build acceptance observation
- RED: cli-contract-usage (FEATURE-CLI-CONTRACT.md)
  intent: Malformed invocations return the usage-error exit status and produce no standard output.
  return code: 1
  stdout:
    usage: csvstat FILE
    exit status 2
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'run', './cmd/csvstat'], returncode=1, stdout='', stderr='usage: csvstat FILE\nexit status 2\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-contract-usage.py", line 22, in <module>
        assert result.returncode == 2
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: cli-contract-dispatch (FEATURE-CLI-CONTRACT.md)
  intent: Valid profile and frequency command forms dispatch successfully.
  return code: 1
  stdout:
    csvstat
    
    
    
    flag provided but not defined: -freq
    Usage of csvstat:
    exit status 2
  stderr:
    --- drydock: values at failure ---
      frequency = CompletedProcess(args=['go', 'run', './cmd/csvstat', '--freq', 'b', '--top', '1', '/tmp/tmpimadh4xy/data.csv'], returncode=1, stdout='', stderr='flag provided but not defined: -freq\nUsage of csvstat:\nexit status 2\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-contract-dispatch.py", line 26, in <module>
        assert frequency.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: cli-help-option (FEATURE-CLI-HELP.md)
  intent: The help option is accepted and does not emit help text on standard output.
  return code: 1
  stdout:
    Usage of csvstat:
    exit status 2
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'run', './cmd/csvstat', '--help'], returncode=1, stdout='', stderr='Usage of csvstat:\nexit status 2\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-help-option.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: cli-help-usage-failure (FEATURE-CLI-HELP.md)
  intent: A usage failure returns the usage exit status and keeps all assistance off standard output.
  return code: 1
  stdout:
    Usage of csvstat:
    exit status 2
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'run', './cmd/csvstat', '--help', '--unknown'], returncode=1, stdout='', stderr='Usage of csvstat:\nexit status 2\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-help-usage-failure.py", line 10, in <module>
        assert result.returncode == 2
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- FAIL: cli-contract-usage (FEATURE-CLI-CONTRACT.md)
  intent: Malformed invocations return the usage-error exit status and produce no standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stdout:
    csvstat: exactly one file argument is required
    Usage:
      csvstat [--column <name>] <file>
      csvstat --freq <name> [--top <n>] <file>
    
    Options:
      --column <name>  profile one column
      --freq <name>    list the most frequent values in a column
      --top <n>        limit a frequency report to a positive number of values
      --help           show this help
    
    Exit codes: 0 success, 1 file access failure, 2 usage error, 3 rejected CSV or missing column.
    exit status 2
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'run', './cmd/csvstat'], returncode=1, stdout='', stderr='csvstat: exactly one file argument is required\nUsage:\n  csvstat [--column <name>] <file>\n  csvstat --freq <name> [--top <n>] <file>\n\nOptions:\n  --column <name>  profile one column\n  --freq <name>    list the most frequent values in a column\n  --top <n>        limit a frequency report to a positive number… (+162 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-contract-usage.py", line 22, in <module>
        assert result.returncode == 2
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- PASS: cli-contract-dispatch (FEATURE-CLI-CONTRACT.md)
  intent: Valid profile and frequency command forms dispatch successfully.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    profile: all columns
    
    
    frequency: b
- PASS: cli-help-option (FEATURE-CLI-HELP.md)
  intent: The help option is accepted and does not emit help text on standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    Usage:
      csvstat [--column <name>] <file>
      csvstat --freq <name> [--top <n>] <file>
    
    Options:
      --column <name>  profile one column
      --freq <name>    list the most frequent values in a column
      --top <n>        limit a frequency report to a positive number of values
      --help           show this help
    
    Exit codes: 0 success, 1 file access failure, 2 usage error, 3 rejected CSV or missing column.
- FAIL: cli-help-usage-failure (FEATURE-CLI-HELP.md)
  intent: A usage failure returns the usage exit status and keeps all assistance off standard output.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 1
  stdout:
    csvstat: help cannot be combined with other arguments
    Usage:
      csvstat [--column <name>] <file>
      csvstat --freq <name> [--top <n>] <file>
    
    Options:
      --column <name>  profile one column
      --freq <name>    list the most frequent values in a column
      --top <n>        limit a frequency report to a positive number of values
      --help           show this help
    
    Exit codes: 0 success, 1 file access failure, 2 usage error, 3 rejected CSV or missing column.
    exit status 2
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'run', './cmd/csvstat', '--help', '--unknown'], returncode=1, stdout='', stderr='csvstat: help cannot be combined with other arguments\nUsage:\n  csvstat [--column <name>] <file>\n  csvstat --freq <name> [--top <n>] <file>\n\nOptions:\n  --column <name>  profile one column\n  --freq <name>    list the most frequent values in a column\n  --top <n>        limit a frequen… (+192 chars)
    --- drydock: end values ---
    Traceback (most recent call last):
      File "cli-help-usage-failure.py", line 10, in <module>
        assert result.returncode == 2
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Repair attempts
- attempt 0 (initial build): failed; 2/4 checks model=gpt-5.6-luna; execution 20261022.203643.347Z-d50010a9; reason: programmatic acceptance failed: cli-contract-usage, cli-help-usage-failure
- attempt 1 (repair 1): failed; 2/4 checks model=gpt-5.6-luna; execution 20261022.203805.869Z-0497c5b3; reason: programmatic acceptance failed: cli-contract-usage, cli-help-usage-failure
- attempt 2 (repair 2): failed; 2/4 checks model=gpt-5.6-luna; execution 20261022.203854.435Z-44544656; stopped: deterministic acceptance score did not improve on 2 consecutive calls; reason: programmatic acceptance failed: cli-contract-usage, cli-help-usage-failure

## Agent self-report (advisory)
The build agent declared a failure. This is advisory only; the programmatic acceptance above is the authority for this block's outcome.
- summary: `go run` normalizes application exit code 2 to 1.
- detail:
    Directly built `csvstat` returns 2 for every malformed invocation, but each authoritative Python check using `go run ./cmd/csvstat` observes return code 1. The criteria must invoke the built executable or account for `go run` behavior.

## Failure
- summary: programmatic acceptance failed: cli-contract-usage, cli-help-usage-failure
- detail:
    Block "Block 2 · Service" [block-2] failed its acceptance criteria.
      Story "Parse csvstat commands, options, arguments, and exit-code categories." [cli-contract] does not meet its own acceptance criteria:
        - AC cli-contract-usage — Malformed invocations return the usage-error exit status and produce no standard output.
            assertion: assert result.returncode == 2 → AssertionError
            raised at: cli-contract-usage.py:22
            process exit code: 1
            values at failure:
              result = CompletedProcess(args=['go', 'run', './cmd/csvstat'], returncode=1, stdout='', stderr='csvstat: exactly one file argument is required\nUsage:\n  csvstat [--column <name>] <file>\n  csvstat --freq <name> [--top <n>] <file>\n\nOptions:\n  --column <name>  profile one column\n  --freq <name>    list the most frequent values in a column\n  --top <n>        limit a frequency report to a positive number… (+162 chars)
            observed output:
              csvstat: exactly one file argument is required
              Usage:
                csvstat [--column <name>] <file>
                csvstat --freq <name> [--top <n>] <file>
              Options:
                --column <name>  profile one column
                --freq <name>    list the most frequent values in a column
                --top <n>        limit a frequency report to a positive number of values
                --help           show this help
              Exit codes: 0 success, 1 file access failure, 2 usage error, 3 rejected CSV or missing column.
              exit status 2
            check stderr:
              Traceback (most recent call last):
                File "cli-contract-usage.py", line 22, in <module>
                  assert result.returncode == 2
                         ^^^^^^^^^^^^^^^^^^^^^^
              AssertionError
      Story "Provide usable help and failure diagnostics for csvstat invocations." [cli-help] does not meet its own acceptance criteria:
        - AC cli-help-usage-failure — A usage failure returns the usage exit status and keeps all assistance off standard output.
            assertion: assert result.returncode == 2 → AssertionError
            raised at: cli-help-usage-failure.py:10
            process exit code: 1
            values at failure:
              result = CompletedProcess(args=['go', 'run', './cmd/csvstat', '--help', '--unknown'], returncode=1, stdout='', stderr='csvstat: help cannot be combined with other arguments\nUsage:\n  csvstat [--column <name>] <file>\n  csvstat --freq <name> [--top <n>] <file>\n\nOptions:\n  --column <name>  profile one column\n  --freq <name>    list the most frequent values in a column\n  --top <n>        limit a frequen… (+192 chars)
            observed output:
              csvstat: help cannot be combined with other arguments
              Usage:
                csvstat [--column <name>] <file>
                csvstat --freq <name> [--top <n>] <file>
              Options:
                --column <name>  profile one column
                --freq <name>    list the most frequent values in a column
                --top <n>        limit a frequency report to a positive number of values
                --help           show this help
              Exit codes: 0 success, 1 file access failure, 2 usage error, 3 rejected CSV or missing column.
              exit status 2
            check stderr:
              Traceback (most recent call last):
                File "cli-help-usage-failure.py", line 10, in <module>
                  assert result.returncode == 2
                         ^^^^^^^^^^^^^^^^^^^^^^
              AssertionError

## Build summary
AC_BROKEN: cli-contract-usage, cli-help-usage-failure

RESULT: FAILED

FILES CHANGED:
- None

SUMMARY:
The CLI parser and help behavior are correct; direct execution returns status 2 and `go test ./...` passes. The declared checks invoke `go run`, which reports application exit status 2 as return code 1, causing both assertions to fail.

BLOCKERS:
- Acceptance criteria require `go run` to preserve exit code 2, which it does not in this environment.

FAILURE_SUMMARY: `go run` normalizes application exit code 2 to 1.
FAILURE_DETAIL: Directly built `csvstat` returns 2 for every malformed invocation, but each authoritative Python check using `go run ./cmd/csvstat` observes return code 1. The criteria must invoke the built executable or account for `go run` behavior.
