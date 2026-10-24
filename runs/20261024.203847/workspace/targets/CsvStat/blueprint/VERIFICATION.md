# FEATURE: Verification

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Terminal release verification: the supplied conformance harness passes in full. |
| Depends On  | FEATURE-ModuleBuild.md, FEATURE-TestScript.md, FEATURE-CsvReading.md, FEATURE-NullTrim.md, FEATURE-RejectMalformed.md, FEATURE-UnreadableFile.md, FEATURE-CliOptions.md, FEATURE-ColumnSelect.md, FEATURE-UsageErrors.md, FEATURE-TypeInference.md, FEATURE-DateType.md, FEATURE-Counts.md, FEATURE-MinMax.md, FEATURE-NumericStats.md, FEATURE-Frequencies.md, FEATURE-CsvOutput.md |
| Provides    | sh sources/run.sh |
| Consumes    | go build ./cmd/csvstat |

## Purpose

Once every implementation story is closed, the supplied harness `sources/run.sh` builds
`./cmd/csvstat` and runs its 22 checks against the supplied data files. This is the only story that
runs the whole suite. Nothing under `sources/` is modified, and application code never reads
`sources/run.sh` or embeds an answer.

## Programmatic Acceptance

=== AC full-conformance ===
Intent: The supplied grader builds csvstat and passes all of its checks.
Suite: full
Requires: executable=go; scope=test

import os
import subprocess
import sys

result = subprocess.run(
    ["sh", "sources/run.sh"],
    capture_output=True,
    text=True,
    env={**os.environ},
)
print(result.stdout)
print(result.stderr, file=sys.stderr)
assert result.returncode == 0
=== END AC full-conformance ===

## User Acceptance

- None.

## Guardrails

- Do not modify, rewrite, regenerate, or substitute any file under `sources/`.
- Application code does not read `sources/run.sh` or any supplied data file, and does not hardcode
  a statistic, a type, or a value list.
