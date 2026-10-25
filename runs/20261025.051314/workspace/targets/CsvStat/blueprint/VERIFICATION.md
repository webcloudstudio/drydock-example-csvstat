# FEATURE: Release Verification

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Terminal assembly story that passes the supplied grader `sources/run.sh` with every one of its 22 checks. |
| Depends On  | FEATURE-TestGate.md, FEATURE-ColumnProfile.md, FEATURE-FreqMode.md, FEATURE-UnreadableFile.md |
| Provides    | sh sources/run.sh |
| Consumes    | csvstat command, bin/test.sh |

## Purpose

This is the terminal story. After every member story has closed, the supplied grader runs unfiltered. The grader builds `./cmd/csvstat` with `go build`, makes 22 checks of exit code and standard output against five supplied CSV files, and exits zero only when the build succeeds and every check passes.

## Behavior

- The grader is run exactly as supplied: `sh sources/run.sh` from the application root.
- Its exit status is the whole verdict.
- A failure is repaired in the owning application package, never in `sources/`.

## Programmatic Acceptance

=== AC release-grader-passes ===
Intent: The supplied grader sources/run.sh, run unfiltered, exits zero.
Suite: full
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(["sh", "sources/run.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
=== END AC release-grader-passes ===

## User Acceptance

- None.

## Guardrails

- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- Application code never reads `sources/run.sh` or any supplied data file, and never hardcodes a statistic, a type, or a value list.
