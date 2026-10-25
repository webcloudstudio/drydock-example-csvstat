# Evidence: Block 5 · Service (block-5)

- block type: block
- date: 2026-10-25
- resulting state: closed/verified
- story points (combined assembled cost): 6549
- execution id: 20261025.055501.273Z-18678fb5

## Stories built
- Provide bin/test.sh running go vet and go test and the project test suite. (test-gate) [story]

## Acceptance tooling authorization
- FEATURE-TestGate.md#testgate-script-exits-zero: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TestGate.md#testgate-runs-from-other-directory: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TestGate.md#testgate-go-test-and-vet-pass: executable=go; scope=test; authorization=existing Target environment
- FEATURE-TestGate.md#testgate-suite-has-tests: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 963)
- implements: FEATURE-TestGate.md (SP 1429)
- context: ARCHITECTURE_compact.md (SP 282)
- stack: common_compact.md (SP 1653)

## Pre-build acceptance observation
- GREEN (prepassed): testgate-script-exits-zero (FEATURE-TestGate.md)
  intent: Running bin/test.sh from the application root exits zero.
  return code: 0
  stdout:
    ?   	csvstat/cmd/csvstat	[no test files]
    ok  	csvstat/internal/cli	0.018s
    ok  	csvstat/internal/csvdata	0.006s
    ok  	csvstat/internal/freq	0.007s
    ok  	csvstat/internal/profile	0.007s
- GREEN (prepassed): testgate-runs-from-other-directory (FEATURE-TestGate.md)
  intent: bin/test.sh locates the application root itself and exits zero when started from another directory.
  return code: 0
  stdout:
    ?   	csvstat/cmd/csvstat	[no test files]
    ok  	csvstat/internal/cli	(cached)
    ok  	csvstat/internal/csvdata	(cached)
    ok  	csvstat/internal/freq	(cached)
    ok  	csvstat/internal/profile	(cached)
- GREEN (prepassed): testgate-go-test-and-vet-pass (FEATURE-TestGate.md)
  intent: go vet and go test pass directly at the application root, the two commands the script wraps.
  return code: 0
  stdout:
    ?   	csvstat/cmd/csvstat	[no test files]
    ok  	csvstat/internal/cli	(cached)
    ok  	csvstat/internal/csvdata	(cached)
    ok  	csvstat/internal/freq	(cached)
    ok  	csvstat/internal/profile	(cached)
- GREEN (prepassed): testgate-suite-has-tests (FEATURE-TestGate.md)
  intent: The project ships Go test files in more than one internal package, so the gate is not vacuous.
  return code: 0
  stdout:
    ['internal/cli/args_test.go', 'internal/cli/cli_test.go', 'internal/cli/column_test.go', 'internal/cli/freq_test.go', 'internal/cli/profile_test.go', 'internal/csvdata/null_test.go', 'internal/csvdata/read_test.go', 'internal/csvdata/trim_test.go', 'internal/csvdata/validate_test.go', 'internal/freq/count_test.go', 'internal/freq/top_test.go', 'internal/profile/counts_test.go', 'internal/profile/format_test.go', 'internal/profile/infer_test.go', 'internal/profile/median_test.go', 'internal/profile/minmax_test.go', 'internal/profile/sum_test.go']

## Post-build programmatic acceptance
- PASS: testgate-script-exits-zero (FEATURE-TestGate.md)
  intent: Running bin/test.sh from the application root exits zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ?   	csvstat/cmd/csvstat	[no test files]
    ok  	csvstat/internal/cli	(cached)
    ok  	csvstat/internal/csvdata	(cached)
    ok  	csvstat/internal/freq	(cached)
    ok  	csvstat/internal/profile	(cached)
- PASS: testgate-runs-from-other-directory (FEATURE-TestGate.md)
  intent: bin/test.sh locates the application root itself and exits zero when started from another directory.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ?   	csvstat/cmd/csvstat	[no test files]
    ok  	csvstat/internal/cli	(cached)
    ok  	csvstat/internal/csvdata	(cached)
    ok  	csvstat/internal/freq	(cached)
    ok  	csvstat/internal/profile	(cached)
- PASS: testgate-go-test-and-vet-pass (FEATURE-TestGate.md)
  intent: go vet and go test pass directly at the application root, the two commands the script wraps.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ?   	csvstat/cmd/csvstat	[no test files]
    ok  	csvstat/internal/cli	(cached)
    ok  	csvstat/internal/csvdata	(cached)
    ok  	csvstat/internal/freq	(cached)
    ok  	csvstat/internal/profile	(cached)
- PASS: testgate-suite-has-tests (FEATURE-TestGate.md)
  intent: The project ships Go test files in more than one internal package, so the gate is not vacuous.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    ['internal/cli/args_test.go', 'internal/cli/cli_test.go', 'internal/cli/column_test.go', 'internal/cli/freq_test.go', 'internal/cli/profile_test.go', 'internal/csvdata/null_test.go', 'internal/csvdata/read_test.go', 'internal/csvdata/trim_test.go', 'internal/csvdata/validate_test.go', 'internal/freq/count_test.go', 'internal/freq/top_test.go', 'internal/profile/counts_test.go', 'internal/profile/format_test.go', 'internal/profile/infer_test.go', 'internal/profile/median_test.go', 'internal/profile/minmax_test.go', 'internal/profile/sum_test.go']

## Build summary
The gate script already exists from an earlier block. I'll run it and the ACs before changing anything.The gate passes from the root and from another directory. Next I'll check the test suite against the spec's coverage list.Coverage looks complete against the spec list. Last check: the tests should write their own inputs to temp dirs and never touch `sources/`.RESULT: SUCCESS

FILES CHANGED:
- none

SUMMARY:
`bin/test.sh` already satisfies FEATURE-TestGate, so I made no edits. It was added in block-1 (commit dcb0bc4) and matches the spec: a POSIX `sh` script with `set -eu` that changes to the application root from its own path, runs `go vet ./...`, then `go test ./...`, and exits non-zero if either fails.

Checks run:
- `sh bin/test.sh` from the application root: exit 0. `go vet` passed, and `go test` passed for `internal/cli`, `csvdata`, `freq`, and `profile`.
- `sh <absolute path>/bin/test.sh` from `/tmp`: exit 0.
- `go vet ./...` run directly: exit 0.
- `go test -count=1 -v ./...`: all packages `ok`. Test counts from this run: 17 `*_test.go` files, and the `go test` run reported no failures.
- `sh -n bin/test.sh`: syntax OK.

Acceptance assertions, all passing:
- `testgate-script-exits-zero`
- `testgate-runs-from-other-directory`
- `testgate-go-test-and-vet-pass`
- `testgate-suite-has-tests`: test files are in four packages, so the `len(packages) > 1` check holds.

Coverage against the spec list, by test function name:
- Types: integer, number, date, boolean, text, and empty.
- Statistics: rows, nulls, distinct, min, max, sum, mean, and median.
- Null handling and trimming.
- Rounding: half-away-from-zero, and no negative zero.
- Output forms: profile, `--column`, and `--freq`.
- Exit codes 0, 1, 2, and 3.

The tests do not reference `sources/`, and the three test files that write inputs use temp directories.

I did not run `sources/run.sh`. It is the terminal story, and this block is not the terminal step.

BLOCKERS:
- none
