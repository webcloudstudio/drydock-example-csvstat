# FEATURE: Frequency

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | List the most frequent non-null values of one column ordered by count descending then value ascending. |
| Depends On  | FEATURE-NullValues.md, FEATURE-ProfileOutput.md, FEATURE-Arguments.md, FEATURE-InputRejection.md |
| Provides    | --freq behavior |
| Consumes    | internal/profile.IsNull |

## Purpose

`csvstat --freq <name> <file>` replaces the profile with a frequency listing of one column.

## Behavior

- Output header is `value,count`, written with the `encoding/csv` writer and `\n` line endings.
- One row per distinct non-null value of the named column, values compared as exact trimmed
  strings.
- Rows are ordered by count descending, then by value ascending by byte value.
- Null values (empty, `NA`, `N/A`, `NULL` ignoring case, after trimming) are not listed.
- A column with no non-null values writes only the header and exits 0.
- At most 5 rows are written when `--top` is not given. `--top` itself is specified in
  FEATURE-FrequencyTop.md.
- A name absent from the header exits 3 with nothing on standard output.

## Programmatic Acceptance

=== AC frequency-ordering ===
Intent: Frequencies list non-null trimmed values by count descending then byte-ordered value, matching counts computed from the input file.
Requires: executable=go; scope=test

import csv
import io
from collections import Counter
from drydock_go import cli_runner

data = "sources/orders.csv"
column = "customer"
with open(data, newline="") as handle:
    records = list(csv.reader(handle))
index = [name.strip() for name in records[0]].index(column)
nulls = {"", "na", "n/a", "null"}
values = [row[index].strip() for row in records[1:]]
counts = Counter(v for v in values if v.lower() not in nulls)
ordered = sorted(counts.items(), key=lambda item: (-item[1], item[0].encode()))
expected = [[value, str(count)] for value, count in ordered][:5]

code, out, err = cli_runner("--freq", column, data)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["value", "count"], rows[0]
assert rows[1:] == expected, (rows[1:], expected)
=== END AC frequency-ordering ===

=== AC frequency-tie-order-by-bytes ===
Intent: Values with equal counts are ordered ascending by byte value.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

values = ["b", "a", "B", "c", "a", "b"]
with tempfile.TemporaryDirectory() as tmp:
    path = os.path.join(tmp, "ties.csv")
    with open(path, "w", newline="") as handle:
        handle.write("v\n" + "\n".join(values) + "\n")
    code, out, err = cli_runner("--freq", "v", path)
    assert code == 0, (code, out, err)
    rows = list(csv.reader(io.StringIO(out)))[1:]
    expected_order = sorted(set(values), key=lambda v: (-values.count(v), v.encode()))
    assert [row[0] for row in rows] == expected_order, (rows, expected_order)
    assert [int(row[1]) for row in rows] == [values.count(v) for v in expected_order], rows
=== END AC frequency-tie-order-by-bytes ===

=== AC frequency-excludes-nulls ===
Intent: Null values are not listed, and counts equal the non-null occurrences in the file.
Requires: executable=go; scope=test

import csv
import io
from collections import Counter
from drydock_go import cli_runner

data = "sources/employees.csv"
column = "manager_id"
with open(data, newline="") as handle:
    records = list(csv.reader(handle))
index = [name.strip() for name in records[0]].index(column)
nulls = {"", "na", "n/a", "null"}
values = [row[index].strip() for row in records[1:]]
counts = Counter(v for v in values if v.lower() not in nulls)

code, out, err = cli_runner("--freq", column, "--top", str(len(values)), data)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))[1:]
listed = {row[0]: int(row[1]) for row in rows}
assert listed == dict(counts), (listed, dict(counts))
assert not any(row[0].lower() in nulls for row in rows), rows
=== END AC frequency-excludes-nulls ===

=== AC frequency-default-limit ===
Intent: Without --top at most five rows are written, even when more distinct values exist.
Requires: executable=go; scope=test

import csv
import io
from drydock_go import cli_runner

data = "sources/employees.csv"
column = "name"
with open(data, newline="") as handle:
    records = list(csv.reader(handle))
index = [name.strip() for name in records[0]].index(column)
distinct = {row[index].strip() for row in records[1:]}
assert len(distinct) > 5

code, out, err = cli_runner("--freq", column, data)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))[1:]
assert len(rows) == 5, rows
=== END AC frequency-default-limit ===

=== AC frequency-all-null-column ===
Intent: A column with no non-null values writes only the header and exits 0.
Requires: executable=go; scope=test

import csv
import io
from drydock_go import cli_runner

code, out, err = cli_runner("--freq", "battery", "sources/sensors.csv")
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert rows == [["value", "count"]], rows
=== END AC frequency-all-null-column ===

=== AC frequency-absent-column ===
Intent: An absent column name exits 3 with nothing on standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("--freq", "nosuch", "sources/orders.csv")
assert code == 3, (code, out, err)
assert out == "", out
=== END AC frequency-absent-column ===

## User Acceptance

- None.

## Guardrails

- Frequencies are computed from the file examined; no value list is hardcoded.
- Values are compared as exact strings after trimming; no case folding.
- Nothing is written to standard output when the command fails.
