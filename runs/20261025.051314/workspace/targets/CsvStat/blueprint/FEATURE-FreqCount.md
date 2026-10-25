# FEATURE: Frequency Count

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Count and order the distinct non-null values of one column for frequency output. |
| Depends On  | FEATURE-NullValues.md, FEATURE-InputRejection.md |
| Provides    | freq.Count |
| Consumes    | csvdata.Read, csvdata.IsNull |

## Purpose

`internal/freq` counts the trimmed non-null values of a named column and returns them in listing
order. The listing is CSV with header `value,count`.

## Behavior

- Values are trimmed before counting, so padded spellings count together. Comparison is by exact
  string after trimming.
- Null values (empty, `NA`, `N/A`, `NULL` ignoring case) are not listed and not counted.
- Order is count descending, then value ascending by byte value.
- A column with no non-null values produces only the header row.
- A header-only file produces only the header row.
- Output is written with `encoding/csv`, so a value containing a comma or quote is quoted.
- Rejected input (ragged, empty, bad header) exits 3 with empty standard output.

## Programmatic Acceptance

=== AC freq-count-ordering ===
Intent: Values are listed by count descending then by byte value ascending, with counts matching the input.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from collections import Counter
from drydock_go import cli_runner

values = ["b", "a", "b", "c", "a", "b", "z", "c", "B", "B"]
rows = [["v"]] + [[x] for x in values]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--freq", "v", path)
print(out)
print(err)
assert code == 0, code
parsed = list(csv.reader(io.StringIO(out)))
assert parsed[0] == ["value", "count"]
counts = Counter(values)
ordered = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))
assert parsed[1:] == [[k, str(n)] for k, n in ordered]
=== END AC freq-count-ordering ===

=== AC freq-count-nulls-and-trim ===
Intent: Null spellings are omitted and padded values count with their trimmed form.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from collections import Counter
from drydock_go import cli_runner

values = ["", "NA", "n/a", "NULL", "null", "x", " x ", "y"]
rows = [["v"]] + [[x] for x in values]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--freq", "v", path)
print(out)
print(err)
assert code == 0, code
parsed = list(csv.reader(io.StringIO(out)))
nulls = {"", "na", "n/a", "null"}
counts = Counter(v.strip() for v in values if v.strip().lower() not in nulls)
ordered = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))
assert parsed[1:] == [[k, str(n)] for k, n in ordered]
=== END AC freq-count-nulls-and-trim ===

=== AC freq-count-all-null ===
Intent: A column with no non-null value lists only the header.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

rows = [["id", "gap"], ["1", ""], ["2", "NA"], ["3", " "]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--freq", "gap", path)
print(out)
print(err)
assert code == 0, code
parsed = list(csv.reader(io.StringIO(out)))
assert parsed == [["value", "count"]]
=== END AC freq-count-all-null ===

=== AC freq-count-header-only ===
Intent: A header-only file lists only the header for any of its columns.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

rows = [["sku", "price"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--freq", "sku", path)
print(out)
print(err)
assert code == 0, code
assert list(csv.reader(io.StringIO(out))) == [["value", "count"]]
=== END AC freq-count-header-only ===

=== AC freq-count-quoted-value ===
Intent: A value containing a comma is written quoted and parses back to the original value.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

value = "relocating, TX"
rows = [["notes"], [value], [value], ["other"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--freq", "notes", path)
print(out)
print(err)
assert code == 0, code
parsed = list(csv.reader(io.StringIO(out)))
assert parsed[1] == [value, str(2)]
assert parsed[2] == ["other", str(1)]
=== END AC freq-count-quoted-value ===

=== AC freq-count-rejected-input ===
Intent: A ragged file in frequency mode exits 3 and writes nothing to standard output.
Requires: executable=go; scope=test

import csv
import os
import tempfile
from drydock_go import cli_runner

rows = [["a", "b"], ["1", "2"], ["3"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "ragged.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--freq", "a", path)
print(err)
assert code == 3, code
assert out == ""
=== END AC freq-count-rejected-input ===

## User Acceptance

- None.

## Guardrails

- Null values are never listed.
- Ordering is by byte value, not locale collation.
- No frequency value is embedded in application code.
