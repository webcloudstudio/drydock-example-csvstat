# FEATURE: Conformance

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Assemble every csvstat story and pass the supplied conformance grader sh sources/run.sh. |
| Depends On  | FEATURE-TestRunner.md, FEATURE-FrequencyTop.md, FEATURE-ColumnSelect.md, FEATURE-ReadFailure.md |
| Provides    | conformance with sources/run.sh |
| Consumes    | csvstat executable, bin/test.sh |

## Purpose

This assembly story runs once every implementation story has closed. The supplied grader builds
`./cmd/csvstat` with `go build` and makes 22 checks of exit code and standard output against the
five supplied CSV files. The grader's exit status is the verdict.

## Behavior

- `sh sources/run.sh` is run from the application root, unfiltered, and exits zero.
- A failing check is repaired in the member stories' code. Files under `sources/` are never
  modified, rewritten, regenerated, or substituted, and application code never reads
  `sources/run.sh` or the supplied data files or embeds any answer.

## Programmatic Acceptance

=== AC grader-passes ===
Intent: The complete csvstat passes every check of the supplied grader.
Suite: full
Requires: executable=go; scope=test

import subprocess

result = subprocess.run(["sh", "sources/run.sh"], capture_output=True, text=True)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0, result.returncode
=== END AC grader-passes ===

## User Acceptance

- None.

## Guardrails

- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- Do not read `sources/run.sh` or any supplied data file from application code, and do not
  hardcode a statistic, a type, or a value list.
- The grader is run unfiltered; no check is skipped or reinterpreted.
