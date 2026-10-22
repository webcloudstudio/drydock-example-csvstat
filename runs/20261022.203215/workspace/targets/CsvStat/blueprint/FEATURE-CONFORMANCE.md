# FEATURE: Supplied Conformance Verification

| Field       | Value |
|-------------|-------|
| Version     | 20261022 V1 |
| Description | Verify the completed csvstat implementation with the supplied conformance suite. |
| Depends On  | FEATURE-TEST-GATE.md |
| Provides    | complete supplied CsvStat conformance verification |
| Consumes    | bin/test.sh project verification command |

## Purpose

Run the supplied read-only conformance harness after all implementation and project-test stories are complete.

## Verification

- Execute `sh sources/run.sh` from the application root.
- Treat the harness exit status as the acceptance verdict.
- The harness builds `./cmd/csvstat` and checks all supplied fixtures and command forms.
- Do not modify, regenerate, filter, or reinterpret the harness or fixtures.

## Programmatic Acceptance

=== AC conformance-full ===
Intent: The supplied CsvStat conformance suite passes against the completed executable.
Suite: full

import subprocess

result = subprocess.run(
    ["sh", "sources/run.sh"],
    capture_output=True,
    text=True,
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC conformance-full ===

## User Acceptance

- None.

## Guardrails

- `sources/run.sh` and every supplied CSV fixture are read-only scoring assets.
- Application code must compute results from the file it receives.
- This is the terminal full-suite verification and must run only after all implementation stories close.
