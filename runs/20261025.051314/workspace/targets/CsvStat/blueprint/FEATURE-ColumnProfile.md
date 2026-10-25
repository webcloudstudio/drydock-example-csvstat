# FEATURE: Column Profile

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Profile a single named column when --column is given. |
| Depends On  | FEATURE-ProfileOutput.md, FEATURE-UsageErrors.md |
| Provides    | csvstat --column |
| Consumes    | csvstat <file>, exit code 2 |

## Purpose

`csvstat --column <name> <file>` writes the profile header and only the row for the named column.

## Behavior

- The output has the same header and the same row content for that column as the full profile.
- The name is matched exactly and case-sensitively against the trimmed header.
- A name not present in the header exits 3 with empty standard output.
- Rejected input (ragged, empty, bad header) exits 3 with empty standard output.
- `--column` combined with `--freq` is a usage error (exit 2), owned by the usage-errors story.

## Programmatic Acceptance

=== AC column-profile-single-row ===
Intent: --column writes the profile header and exactly one row, equal to that column's row in the full profile.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

rows = [["id", "price", "label"], ["1", "1.50", "a"], ["2", "2.25", "b"], ["4", "3.00", "c"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code_all, out_all, err_all = cli_runner(path)
    assert code_all == 0, code_all
    full = list(csv.reader(io.StringIO(out_all)))
    for name in rows[0]:
        code, out, err = cli_runner("--column", name, path)
        print(out)
        print(err)
        assert code == 0, code
        parsed = list(csv.reader(io.StringIO(out)))
        assert parsed[0] == full[0]
        assert len(parsed) == 2
        assert parsed[1][0] == name
        assert parsed[1] == [r for r in full[1:] if r[0] == name][0]
=== END AC column-profile-single-row ===

=== AC column-profile-trimmed-header ===
Intent: A padded header name is trimmed, so the trimmed name selects the column and is reported.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

padded = "  qty "
rows = [[padded, "other"], ["1", "x"], ["2", "y"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--column", padded.strip(), path)
print(out)
print(err)
assert code == 0, code
parsed = list(csv.reader(io.StringIO(out)))
assert len(parsed) == 2
assert parsed[1][0] == padded.strip()
assert parsed[1][1] == "integer"
=== END AC column-profile-trimmed-header ===

=== AC column-profile-unknown-column ===
Intent: A column name absent from the header exits 3 with empty standard output.
Requires: executable=go; scope=test

import csv
import os
import tempfile
from drydock_go import cli_runner

rows = [["qty", "price"], ["1", "1.50"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--column", "nosuch", path)
print(err)
assert code == 3, code
assert out == ""
=== END AC column-profile-unknown-column ===

=== AC column-profile-case-sensitive ===
Intent: Column names match case-sensitively, so a different case is an unknown column.
Requires: executable=go; scope=test

import csv
import os
import tempfile
from drydock_go import cli_runner

name = "qty"
rows = [[name, "price"], ["1", "1.50"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--column", name.upper(), path)
print(err)
assert code == 3, code
assert out == ""
=== END AC column-profile-case-sensitive ===

=== AC column-profile-rejected-input ===
Intent: A ragged file with --column exits 3 and writes nothing to standard output.
Requires: executable=go; scope=test

import csv
import os
import tempfile
from drydock_go import cli_runner

rows = [["a", "b"], ["1", "2"], ["3", "4", "5"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "ragged.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--column", "a", path)
print(err)
assert code == 3, code
assert out == ""
=== END AC column-profile-rejected-input ===

## User Acceptance

- None.

## Guardrails

- Standard output stays empty whenever the command fails.
- Column names are never matched case-insensitively.
