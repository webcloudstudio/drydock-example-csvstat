# FEATURE: Null Values

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Classify a trimmed value as null when it is empty or equals NA, N/A, or NULL ignoring case. |
| Depends On  | ARCHITECTURE.md, FEATURE-CSV-Reading.md |
| Provides    | csvdata.IsNull |
| Consumes    | csvdata.Read |

## Purpose

Null handling is one rule shared by every statistic: nulls count toward `rows` and `nulls` but are
ignored by `distinct`, `min`, `max`, `sum`, `mean`, `median`, type inference and frequencies.

## Behavior

- `internal/csvdata` exposes `IsNull(value string) bool`, applied to already-trimmed values.
- It returns true for the empty string and for `NA`, `N/A` and `NULL` in any letter case.
- It returns false for every other value, including `0`, `nan`, `none`, `n/a/` and `NAN`.
- No other package re-implements the null rule; all call `csvdata.IsNull`.

## Required unit tests

The build agent writes these tests in `internal/csvdata` with exactly these names:

- `TestIsNullEmptyAndSentinels` covers the empty string and every case variant of `NA`, `N/A`, `NULL`.
- `TestIsNullRejectsLookalikes` covers `0`, `nan`, `none`, `NAN`, `nul` and `n/a/`.

## Programmatic Acceptance

=== AC null-values-sentinels ===
Intent: Empty and the NA, N/A, NULL sentinels in any case are null.

from drydock_go import run_test

result = run_test("./internal/csvdata", "TestIsNullEmptyAndSentinels")
print(result.stdout)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC null-values-sentinels ===

=== AC null-values-lookalikes ===
Intent: Values that merely resemble a null sentinel are not null.

from drydock_go import run_test

result = run_test("./internal/csvdata", "TestIsNullRejectsLookalikes")
print(result.stdout)
assert result.passed, (result.status, result.stdout, result.stderr)
=== END AC null-values-lookalikes ===

=== AC null-values-file-with-nulls-reads ===
Intent: A file containing null sentinels of every spelling is read by the built executable without error.

import os
import tempfile
from drydock_go import cli_runner

sentinels = ["", "NA", "na", "N/A", "n/a", "NULL", "null"]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "nulls.csv")
    with open(path, "w", newline="") as f:
        f.write("id,v\n")
        for i, s in enumerate(sentinels):
            f.write(f"{i},{s}\n")
    code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0
=== END AC null-values-file-with-nulls-reads ===

## User Acceptance

- None.

## Guardrails

- A value is trimmed before it is tested; a value of only spaces is therefore null.
- `0`, `nan` and `none` are never null.
