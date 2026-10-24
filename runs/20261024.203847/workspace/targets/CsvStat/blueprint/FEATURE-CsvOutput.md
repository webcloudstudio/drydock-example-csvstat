# FEATURE: CSV Output

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Write profile and frequency results as RFC 4180 CSV with newline line endings, only on success. |
| Depends On  | FEATURE-Counts.md, FEATURE-MinMax.md, FEATURE-NumericStats.md, FEATURE-Frequencies.md, FEATURE-ColumnSelect.md |
| Provides    | output.WriteProfile, output.WriteFreq |
| Consumes    | profile.Counts, profile.MinMax, profile.NumericStats, freq.Top |

## Purpose

Serialize the profile and frequency results to standard output.

## Behavior

1. Profile output has the header `column,type,rows,nulls,distinct,min,max,sum,mean,median`, then
   one row per column in file order, or the single row of the column named by `--column`.
2. Frequency output has the header `value,count`, then the ranked rows.
3. Writing uses `encoding/csv` with `UseCRLF` false, so every line ends with `\n`. A field
   containing a comma, a double quote, or a newline is quoted as RFC 4180 specifies.
4. Output is buffered and written to standard output only when the whole command succeeds. A
   failing command writes nothing to standard output and one diagnostic line to standard error
   naming the file and, for a rejected record, its line number.
5. Empty statistic fields are written as empty strings.

## Programmatic Acceptance

=== AC output-profile-shape ===
Intent: Profile output has the declared header and one row per column in file order.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

columns = ["zeta", "alpha", "mid"]
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as f:
        w = csv.writer(f)
        w.writerow(columns)
        w.writerow(["1", "x", "true"])
        w.writerow(["2", "y", "false"])
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["column", "type", "rows", "nulls", "distinct", "min", "max", "sum", "mean", "median"]
assert [r[0] for r in rows[1:]] == columns
assert all(len(r) == len(rows[0]) for r in rows)
=== END AC output-profile-shape ===

=== AC output-newline-endings ===
Intent: Output lines end with a line feed and never a carriage return.
Requires: executable=go; scope=test

import csv
import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as f:
        w = csv.writer(f, lineterminator="\r\n")
        w.writerow(["a", "b"])
        w.writerow(["1", "x"])
        w.writerow(["2", "y"])
    code, out_profile, err = cli_runner(path)
    code2, out_freq, err2 = cli_runner("--freq", "b", path)
assert code == 0
assert code2 == 0
assert "\r" not in out_profile
assert "\r" not in out_freq
assert out_profile.endswith("\n")
assert out_freq.endswith("\n")
=== END AC output-newline-endings ===

=== AC output-profile-quoting ===
Intent: A profile field containing a comma or a double quote round-trips through RFC 4180 quoting.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

with_comma = "remote, part-time"
with_quote = 'say "hi" now'
for value in (with_comma, with_quote):
    with tempfile.TemporaryDirectory() as tmp:
        path = os.path.join(tmp, "data.csv")
        with open(path, "w", newline="") as f:
            w = csv.writer(f)
            w.writerow(["notes"])
            w.writerow([value])
        code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0
    rows = list(csv.reader(io.StringIO(out)))
    header = rows[0]
    row = dict(zip(header, rows[1]))
    assert row["column"] == "notes"
    assert row["min"] == value
    assert row["max"] == value
=== END AC output-profile-quoting ===

=== AC output-freq-quoting ===
Intent: A frequency value containing a comma round-trips through RFC 4180 quoting.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

value = "relocating to Austin, TX"
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as f:
        w = csv.writer(f)
        w.writerow(["notes"])
        w.writerow([value])
        w.writerow([value])
    code, out, err = cli_runner("--freq", "notes", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["value", "count"]
assert rows[1] == [value, "2"]
=== END AC output-freq-quoting ===

=== AC output-column-single-row ===
Intent: --column yields the header and exactly the named column's row.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

target = "beta"
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as f:
        w = csv.writer(f)
        w.writerow(["alpha", target, "gamma"])
        w.writerow(["1", "2", "3"])
    code, out, err = cli_runner("--column", target, path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2
assert rows[1][0] == target
=== END AC output-column-single-row ===

=== AC output-failure-silent-stdout ===
Intent: A rejected file writes nothing to standard output.
Requires: executable=go; scope=test

import csv
import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "ragged.csv")
    with open(path, "w", newline="") as f:
        w = csv.writer(f)
        w.writerow(["a", "b"])
        w.writerow(["1", "2"])
        w.writerow(["3", "4", "5"])
    code, out, err = cli_runner(path)
    code2, out2, err2 = cli_runner("--freq", "a", path)
print(err)
assert code == 3
assert len(out) == 0
assert code2 == 3
assert len(out2) == 0
=== END AC output-failure-silent-stdout ===

## User Acceptance

- None.

## Guardrails

- Diagnostics go to standard error and are never part of standard output.
- Output is never partially written before a failure is detected.
