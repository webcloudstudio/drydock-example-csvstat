# FEATURE: Frequency Listing

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Lists the most frequent non-null values of one named column with --freq. |
| Depends On  | ARCHITECTURE.md, FEATURE-ProfileCounts.md, FEATURE-CliParsing.md |
| Provides    | value,count output |
| Consumes    | csvstat --freq <name> <file> |

## Purpose

`csvstat --freq <name> <file>` writes a two-column CSV of the distinct non-null values of the named
column and how often each occurs, instead of the profile.

## Behavior

- Header line is `value,count`.
- One row per distinct non-null value, ordered by count descending, then by value ascending by byte
  value.
- Values are trimmed before counting and compared as exact strings. Null values (empty, or `NA`,
  `N/A`, `NULL` ignoring case) are not listed.
- A column with no non-null values writes only the header.
- At most the `--top` limit of rows is written; the default limit is 5 (the limit itself is owned by
  FEATURE-FrequencyTop.md).
- Output is CSV with `\n` line endings, written with `encoding/csv`, and only after the whole file
  has been read and validated.
- Success exits 0.

## Programmatic Acceptance

=== AC freq-department ===
Intent: Frequencies for a text column are counted, ordered by count then byte value, and parsed back as CSV.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import csv
import io
import os
from collections import Counter

from drydock_go import cli_runner

path = os.path.abspath("sources/employees.csv")
col = "department"
top = 5
with open(path, newline="") as fh:
    rows = list(csv.reader(fh))
header = [h.strip() for h in rows[0]]
idx = header.index(col)
vals = [r[idx].strip() for r in rows[1:]]
vals = [v for v in vals if v != "" and v.upper() not in ("NA", "N/A", "NULL")]
ordered = sorted(Counter(vals).items(), key=lambda kv: (-kv[1], kv[0].encode()))
want = [[v, str(c)] for v, c in ordered[:top]]

code, out, err = cli_runner("--freq", col, path)
print(out)
print(err)
assert code == 0, (code, out, err)
got = list(csv.reader(io.StringIO(out)))
assert got[0] == ["value", "count"], got
assert got[1:] == want, (got[1:], want)
=== END AC freq-department ===

=== AC freq-tie-order ===
Intent: Values with equal counts are ordered ascending by byte value and the limit is honored.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import csv
import io
import os
from collections import Counter

from drydock_go import cli_runner

path = os.path.abspath("sources/sensors.csv")
col = "sensor"
top = 2
with open(path, newline="") as fh:
    rows = list(csv.reader(fh))
header = [h.strip() for h in rows[0]]
idx = header.index(col)
vals = [r[idx].strip() for r in rows[1:]]
vals = [v for v in vals if v != "" and v.upper() not in ("NA", "N/A", "NULL")]
ordered = sorted(Counter(vals).items(), key=lambda kv: (-kv[1], kv[0].encode()))
want = [[v, str(c)] for v, c in ordered[:top]]

code, out, err = cli_runner("--freq", col, "--top", str(top), path)
print(out)
print(err)
assert code == 0, (code, out, err)
got = list(csv.reader(io.StringIO(out)))
assert got[1:] == want, (got[1:], want)
assert len(got) == top + 1
=== END AC freq-tie-order ===

=== AC freq-trims-and-skips-nulls ===
Intent: Padded values are trimmed before counting and null spellings (blank, NA) are not listed.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import csv
import io
import os
from collections import Counter

from drydock_go import cli_runner

for fname, col, top in (("orders.csv", "customer", 5), ("employees.csv", "manager_id", 10)):
    path = os.path.abspath(os.path.join("sources", fname))
    with open(path, newline="") as fh:
        rows = list(csv.reader(fh))
    header = [h.strip() for h in rows[0]]
    idx = header.index(col)
    vals = [r[idx].strip() for r in rows[1:]]
    vals = [v for v in vals if v != "" and v.upper() not in ("NA", "N/A", "NULL")]
    ordered = sorted(Counter(vals).items(), key=lambda kv: (-kv[1], kv[0].encode()))
    want = [[v, str(c)] for v, c in ordered[:top]]

    code, out, err = cli_runner("--freq", col, "--top", str(top), path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    got = list(csv.reader(io.StringIO(out)))
    assert got[1:] == want, (fname, got[1:], want)
    listed = {r[0] for r in got[1:]}
    assert "" not in listed and "NA" not in listed
=== END AC freq-trims-and-skips-nulls ===

=== AC freq-all-null-column ===
Intent: A column with no non-null values writes only the header and exits 0.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import csv
import io
import os

from drydock_go import cli_runner

path = os.path.abspath("sources/sensors.csv")
code, out, err = cli_runner("--freq", "battery", path)
print(out)
print(err)
assert code == 0, (code, out, err)
got = list(csv.reader(io.StringIO(out)))
assert got == [["value", "count"]], got
=== END AC freq-all-null-column ===

=== AC freq-output-line-endings ===
Intent: Frequency output uses \n line endings only and ends with a newline.
Suite: scoped
Requires: python-package=drydock_go; scope=test

import os

from drydock_go import cli_runner

path = os.path.abspath("sources/employees.csv")
code, out, err = cli_runner("--freq", "department", path)
print(err)
assert code == 0, (code, out, err)
assert "\r" not in out
assert out.endswith("\n")
=== END AC freq-output-line-endings ===

## User Acceptance

- None.

## Guardrails

- Null values are never listed.
- Nothing is written to standard output unless the whole file has been read and validated.
- No frequency or value list is hardcoded; every value comes from the file examined.
- Application code never reads `sources/run.sh` or any supplied data file.
