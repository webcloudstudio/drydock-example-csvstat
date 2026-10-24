# FEATURE: Frequencies

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | List the most frequent non-null values of one column with --freq, ordered by count then byte value. |
| Depends On  | FEATURE-NullTrim.md, FEATURE-CliOptions.md |
| Provides    | freq.Top |
| Consumes    | reader.Trim, reader.IsNull, --freq, --top |

## Purpose

With `--freq <name>`, csvstat lists the most frequent values of the named column instead of
profiling. The command is `csvstat --freq <name> [--top <n>] <file>`.

## Behavior

1. The column is found by exact, case-sensitive comparison against the trimmed header names.
2. Each value is trimmed before counting. Null values (empty, `NA`, `N/A`, `NULL` ignoring case)
   are not listed and not counted.
3. One row is produced per distinct non-null value, with its count.
4. Rows are ordered by count descending, then by value ascending by byte value.
5. At most `--top` rows are returned. `--top` defaults to 5 and is a positive integer.
6. A column with no non-null values produces only the header row.
7. A column name absent from the header exits 3 with nothing on standard output and one diagnostic
   on standard error.

## Output

Header `value,count`, then the ranked rows. Serialization is owned by FEATURE-CsvOutput.md; this
feature supplies the ordered list of (value, count) pairs as `freq.Top`.

## Programmatic Acceptance

=== AC freq-order-and-default-top ===
Intent: Frequencies are ordered by count descending then byte value ascending, and default to five rows.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from collections import Counter

from drydock_go import cli_runner

values = ["delta"] * 4 + ["alpha"] * 3 + ["charlie"] * 3 + ["bravo"] * 2 + ["echo"] + ["foxtrot"] + ["golf"]
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as f:
        w = csv.writer(f)
        w.writerow(["k"])
        for v in values:
            w.writerow([v])
    code, out, err = cli_runner("--freq", "k", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["value", "count"]
counts = Counter(values)
ranked = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0]))[:5]
assert rows[1:] == [[v, str(c)] for v, c in ranked]
=== END AC freq-order-and-default-top ===

=== AC freq-top-limit ===
Intent: --top limits the number of rows listed.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from collections import Counter

from drydock_go import cli_runner

values = ["a"] * 5 + ["b"] * 4 + ["c"] * 3 + ["d"] * 2 + ["e"]
top = 2
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as f:
        w = csv.writer(f)
        w.writerow(["k"])
        for v in values:
            w.writerow([v])
    code, out, err = cli_runner("--freq", "k", "--top", str(top), path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == top + 1
ranked = sorted(Counter(values).items(), key=lambda kv: (-kv[1], kv[0]))[:top]
assert rows[1:] == [[v, str(c)] for v, c in ranked]
=== END AC freq-top-limit ===

=== AC freq-trim-and-nulls ===
Intent: Values are trimmed before counting and null spellings are not listed.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

core = "Globex"
raw = [core, "  " + core + "  ", core + " ", "", "NA", "n/a", "NULL", "Other"]
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as f:
        w = csv.writer(f)
        w.writerow(["k"])
        for v in raw:
            w.writerow([v])
    code, out, err = cli_runner("--freq", "k", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
listed = {r[0]: int(r[1]) for r in rows[1:]}
assert listed == {core: 3, "Other": 1}
=== END AC freq-trim-and-nulls ===

=== AC freq-all-null-column ===
Intent: A column with no non-null values writes only the header.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as f:
        w = csv.writer(f)
        w.writerow(["k", "other"])
        w.writerow(["", "x"])
        w.writerow(["NULL", "y"])
        w.writerow([" ", "z"])
    code, out, err = cli_runner("--freq", "k", path)
print(out)
print(err)
assert code == 0
rows = list(csv.reader(io.StringIO(out)))
assert rows == [["value", "count"]]
=== END AC freq-all-null-column ===

=== AC freq-absent-column ===
Intent: A column name absent from the header exits 3 and writes nothing to standard output.
Requires: executable=go; scope=test

import csv
import os
import tempfile

from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as f:
        w = csv.writer(f)
        w.writerow(["k"])
        w.writerow(["v"])
    code, out, err = cli_runner("--freq", "nosuch", path)
print(err)
assert code == 3
assert len(out) == 0
=== END AC freq-absent-column ===

=== AC freq-case-sensitive-name ===
Intent: The --freq column name is matched case-sensitively against the header.
Requires: executable=go; scope=test

import csv
import os
import tempfile

from drydock_go import cli_runner

name = "Kind"
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "data.csv")
    with open(path, "w", newline="") as f:
        w = csv.writer(f)
        w.writerow([name])
        w.writerow(["v"])
    code, out, err = cli_runner("--freq", name.lower(), path)
print(err)
assert code == 3
assert len(out) == 0
=== END AC freq-case-sensitive-name ===

## User Acceptance

- None.

## Guardrails

- Null values are never listed.
- The ordering tie-break is byte value, not locale collation.
- A failing command writes nothing to standard output.
