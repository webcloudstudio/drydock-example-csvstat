# FEATURE: Full Verification

| Field       | Value |
|-------------|-------|
| Version     | 20261023 V1 |
| Description | Execute the complete supplied CsvStat conformance suite. |
| Depends On  | ARCHITECTURE.md, FEATURE-CLI-FOUNDATION.md, FEATURE-CLI-INPUT.md, FEATURE-CLI-COLUMNS.md, FEATURE-PROFILE-TYPES.md, FEATURE-PROFILE-STATISTICS.md, FEATURE-PROFILE-OUTPUT.md, FEATURE-FREQUENCY-VALUES.md, FEATURE-FREQUENCY-LIMITS.md, FEATURE-CLI-ERRORS.md, FEATURE-TEST-GATE.md |
| Provides    | full CsvStat conformance verification |
| Consumes    | bin/test.sh, ./cmd/csvstat, complete CsvStat implementation |

## Verification Scope

The supplied `sources/run.sh` is the authoritative release conformance suite. It builds `./cmd/csvstat`, exercises all supplied CSV fixtures and command forms, and returns zero only when every check passes. Run it from the completed application root without filtering, modification, or reinterpretation.

## Programmatic Acceptance

=== AC full-conformance ===
Intent: The supplied conformance suite passes for the completed CsvStat implementation.
Suite: full

import os
import subprocess

result = subprocess.run(
    ["sh", "sources/run.sh"],
    capture_output=True,
    text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr)
assert result.returncode == 0
=== END AC full-conformance ===

## User Acceptance

- None.

## Guardrails

- Execute the supplied suite exactly as staged.
- Do not modify, regenerate, filter, or reinterpret any supplied source or fixture.
- Treat the suite exit status as the release verification result.
