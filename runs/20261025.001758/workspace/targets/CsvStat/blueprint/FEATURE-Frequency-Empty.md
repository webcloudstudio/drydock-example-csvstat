# FEATURE: Frequency Empty Column

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Write only the header when a column has no non-null values. |
| Depends On  | FEATURE-Frequency-List.md |
| Provides    | empty frequency output |
| Consumes    | frequency list |

## Purpose

When the column named by `--freq` has no non-null value, or the file has no data records, the
command succeeds and writes only the header `value,count`.

## Behavior

- Exit status is 0.
- Standard output contains exactly one record, the header `value,count`.
- Values that are all null (empty, whitespace only, `NA`, `N/A`, `NULL` in any case) count as no
  values.

## Programmatic Acceptance

=== AC empty-column-header-only ===
Intent: --freq on the all-null battery column of sensors.csv exits 0 and writes only the value,count header.

import csv
import io
from drydock_go import cli_runner

code, out, err = cli_runner("--freq", "battery", "sources/sensors.csv")
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert got == [["value", "count"]]
assert len(out.splitlines()) == 1
=== END AC empty-column-header-only ===

=== AC empty-header-only-file ===
Intent: --freq on a file with a header and no data records exits 0 and writes only the value,count header.

import csv
import io
from drydock_go import cli_runner

code, out, err = cli_runner("--freq", "sku", "sources/header_only.csv")
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert got == [["value", "count"]]
=== END AC empty-header-only-file ===

=== AC empty-all-null-spellings ===
Intent: A column whose values are only null spellings writes only the header, while a populated column of the same file is still listed.

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

nulls = ["NA", "n/a", "NULL", "null", "", "N/A"]
d = tempfile.mkdtemp()
path = os.path.join(d, "nulls.csv")
with open(path, "w", newline="", encoding="utf-8") as f:
    w = csv.writer(f, lineterminator="\n")
    w.writerow(["x", "y"])
    for i, v in enumerate(nulls):
        w.writerow([v, "k" + str(i % 2)])
code, out, err = cli_runner("--freq", "x", path)
print(out)
print(err)
assert code == 0
assert list(csv.reader(io.StringIO(out))) == [["value", "count"]]
code, out, err = cli_runner("--freq", "y", path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert len(got) == 1 + 2
=== END AC empty-all-null-spellings ===

## User Acceptance

- None.

## Guardrails

- An all-null column never produces an error exit.
- Null values are never listed, even when no other value exists.
