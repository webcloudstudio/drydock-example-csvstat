# FEATURE: Frequency Mode

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Dispatches `csvstat --freq <name> [--top <n>] <file>` to the frequency listing instead of the profile. |
| Depends On  | FEATURE-FreqLimit.md, FEATURE-UsageErrors.md |
| Provides    | csvstat --freq, csvstat --freq --top |
| Consumes    | freq.Count, freq.Top, cli.ParseArgs |

## Purpose

When `--freq <name>` is given, `cli.Run` writes the frequency listing for that column and never the profile. `--top <n>` limits the rows and defaults to 5. The listing is CSV with header `value,count`, rows ordered by count descending and then value ascending by byte value, using trimmed non-null values only.

## Behavior

- The column name is matched exactly and case-sensitively against the trimmed header.
- A column name absent from the header exits 3 with nothing on standard output and one diagnostic line on standard error.
- Combinations that are usage errors (`--column` with `--freq`, `--top` without `--freq`, a non-positive `--top`) are rejected with exit 2 by the usage-errors story before this dispatch runs.
- Output is built in memory and written to standard output only on success, so a failing command writes nothing to standard output.
- A column with no non-null values writes only the header.

## Programmatic Acceptance

=== AC freq-default-matches-fixture ===
Intent: The default frequency listing for a named column equals the top five counts computed independently from the fixture.
Requires: executable=go; scope=test

import csv
import io
from collections import Counter
from drydock_go import cli_runner

path = "sources/employees.csv"
column = "department"
with open(path, newline="", encoding="utf-8") as fh:
    records = list(csv.reader(fh))
header = [h.strip() for h in records[0]]
idx = header.index(column)
values = [r[idx].strip() for r in records[1:] if r[idx].strip() != ""]
counts = Counter(values)
ordered = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))
expected = [[v, str(c)] for v, c in ordered][:5]

code, out, err = cli_runner("--freq", column, path)
print(out)
print(err)
assert code == 0, code
rows = list(csv.reader(io.StringIO(out)))
assert rows[0] == ["value", "count"], rows[0]
assert rows[1:] == expected, (rows[1:], expected)
=== END AC freq-default-matches-fixture ===

=== AC freq-top-limits-rows ===
Intent: With --top 2 exactly the first two ordered rows are listed.
Requires: executable=go; scope=test

import csv
import io
from collections import Counter
from drydock_go import cli_runner

path = "sources/sensors.csv"
column = "sensor"
top = 2
with open(path, newline="", encoding="utf-8") as fh:
    records = list(csv.reader(fh))
header = [h.strip() for h in records[0]]
idx = header.index(column)
values = [r[idx].strip() for r in records[1:] if r[idx].strip() != ""]
counts = Counter(values)
ordered = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))
expected = [[v, str(c)] for v, c in ordered][:top]

code, out, err = cli_runner("--freq", column, "--top", str(top), path)
print(out)
print(err)
assert code == 0, code
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) - 1 == top, rows
assert rows[1:] == expected, (rows[1:], expected)
=== END AC freq-top-limits-rows ===

=== AC freq-omits-null-spellings ===
Intent: Null values (empty, NA, N/A, NULL in any case) are not listed in a frequency listing.
Requires: executable=go; scope=test

import csv
import io
from collections import Counter
from drydock_go import cli_runner

path = "sources/employees.csv"
column = "manager_id"
top = 3
nulls = {"", "na", "n/a", "null"}
with open(path, newline="", encoding="utf-8") as fh:
    records = list(csv.reader(fh))
header = [h.strip() for h in records[0]]
idx = header.index(column)
values = [r[idx].strip() for r in records[1:] if r[idx].strip().lower() not in nulls]
counts = Counter(values)
ordered = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))
expected = [[v, str(c)] for v, c in ordered][:top]

code, out, err = cli_runner("--freq", column, "--top", str(top), path)
print(out)
print(err)
assert code == 0, code
rows = list(csv.reader(io.StringIO(out)))
listed = {r[0].lower() for r in rows[1:]}
assert listed.isdisjoint(nulls), listed
assert rows[1:] == expected, (rows[1:], expected)
=== END AC freq-omits-null-spellings ===

=== AC freq-values-are-trimmed ===
Intent: Padded values are trimmed before counting, so padded and unpadded spellings count together.
Requires: executable=go; scope=test

import csv
import io
from collections import Counter
from drydock_go import cli_runner

path = "sources/orders.csv"
column = "customer"
with open(path, newline="", encoding="utf-8") as fh:
    records = list(csv.reader(fh))
header = [h.strip() for h in records[0]]
idx = header.index(column)
values = [r[idx].strip() for r in records[1:] if r[idx].strip() != ""]
counts = Counter(values)
ordered = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))
expected = [[v, str(c)] for v, c in ordered][:5]

code, out, err = cli_runner("--freq", column, path)
print(out)
print(err)
assert code == 0, code
rows = list(csv.reader(io.StringIO(out)))
assert rows[1:] == expected, (rows[1:], expected)
=== END AC freq-values-are-trimmed ===

=== AC freq-all-null-column-header-only ===
Intent: A column with no non-null values lists only the header.
Requires: executable=go; scope=test

import csv
import io
from drydock_go import cli_runner

code, out, err = cli_runner("--freq", "battery", "sources/sensors.csv")
print(out)
print(err)
assert code == 0, code
rows = list(csv.reader(io.StringIO(out)))
assert rows == [["value", "count"]], rows
=== END AC freq-all-null-column-header-only ===

=== AC freq-unknown-column-rejected ===
Intent: A column name absent from the header exits 3 and writes nothing to standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("--freq", "nosuch", "sources/orders.csv")
print(out)
print(err)
assert code == 3, code
assert out == "", out
=== END AC freq-unknown-column-rejected ===

=== AC freq-name-is-case-sensitive ===
Intent: A column name differing only in case from a header name is not found and exits 3.
Requires: executable=go; scope=test

import csv
from drydock_go import cli_runner

path = "sources/employees.csv"
with open(path, newline="", encoding="utf-8") as fh:
    header = [h.strip() for h in next(csv.reader(fh))]
name = header[2]
swapped = name.upper() if name != name.upper() else name.lower()
assert swapped not in header

code, out, err = cli_runner("--freq", swapped, path)
print(out)
print(err)
assert code == 3, code
assert out == "", out
=== END AC freq-name-is-case-sensitive ===

## User Acceptance

- None.

## Guardrails

- Frequency mode never writes the profile header or any profile row.
- Application code never reads `sources/run.sh` or embeds a count or value list; every row comes from the file examined.
- A failing command writes nothing to standard output.
