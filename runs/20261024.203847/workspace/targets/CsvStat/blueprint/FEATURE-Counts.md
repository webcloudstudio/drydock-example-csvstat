# FEATURE: Counts

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Count rows, null values, and distinct non-null values for each column. |
| Depends On  | FEATURE-NullTrim.md, ARCHITECTURE.md |
| Provides    | profile.Counts |
| Consumes    | reader.IsNull, reader.Trim |

## Purpose

Supply the `rows`, `nulls`, and `distinct` fields of each profile row.

## Behavior

- `rows` is the number of data records (the header is not counted). A header-only file has `rows` 0 for every column.
- `nulls` is the number of values in the column that are null: empty after trimming, or equal to `NA`, `N/A`, or `NULL` ignoring case.
- `distinct` is the number of different non-null values, compared as exact strings after trimming. Comparison is case-sensitive, so `Acme` and `acme` are two values; `  Globex  ` and `Globex` are one.
- An all-null or zero-row column has `distinct` 0.
- `profile.Counts(values []string) (rows, nulls, distinct int)` takes trimmed values.

## Programmatic Acceptance

=== AC counts-rows-nulls-distinct ===
Intent: rows, nulls, and distinct match the counts derived from the supplied column values.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

NULLS = {"", "na", "n/a", "null"}
values = ["apple", "pear", "apple", "", "NA", "n/a", "Null", "fig", "pear", "apple"]
text = "fruit\n" + "\n".join(values) + "\n"
expected_rows = len(values)
expected_nulls = sum(1 for v in values if v.strip().lower() in NULLS)
expected_distinct = len({v.strip() for v in values if v.strip().lower() not in NULLS})
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.DictReader(io.StringIO(out)))
assert len(rows) == 1
r = rows[0]
assert int(r["rows"]) == expected_rows
assert int(r["nulls"]) == expected_nulls
assert int(r["distinct"]) == expected_distinct
=== END AC counts-rows-nulls-distinct ===

=== AC counts-trimmed-distinct ===
Intent: Padded and unpadded spellings of one value count once, while differing case counts separately.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

values = ["  Globex  ", "Globex", "Globex ", "globex", "Acme"]
text = "customer\n" + "\n".join(values) + "\n"
expected_distinct = len({v.strip() for v in values})
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.DictReader(io.StringIO(out)))
assert int(rows[0]["distinct"]) == expected_distinct
assert int(rows[0]["nulls"]) == 0
assert int(rows[0]["rows"]) == len(values)
=== END AC counts-trimmed-distinct ===

=== AC counts-all-null-column ===
Intent: A column of only null spellings has nulls equal to rows and distinct zero.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

values = ["", "NA", "N/A", "NULL", "null", " "]
text = "a,b\n" + "\n".join(v + ",x" for v in values) + "\n"
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = {r["column"]: r for r in csv.DictReader(io.StringIO(out))}
assert int(rows["a"]["rows"]) == len(values)
assert int(rows["a"]["nulls"]) == len(values)
assert int(rows["a"]["distinct"]) == 0
assert int(rows["b"]["nulls"]) == 0
assert int(rows["b"]["distinct"]) == 1
=== END AC counts-all-null-column ===

=== AC counts-header-only ===
Intent: A header-only file reports zero rows, zero nulls, and zero distinct for every column.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

names = ["sku", "price", "in_stock"]
text = ",".join(names) + "\n"
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.DictReader(io.StringIO(out)))
assert [r["column"] for r in rows] == names
for r in rows:
    assert int(r["rows"]) == 0
    assert int(r["nulls"]) == 0
    assert int(r["distinct"]) == 0
=== END AC counts-header-only ===

## User Acceptance

- None.

## Guardrails

- Counts are computed from the file given; none is hardcoded.
- Distinct comparison is exact and case-sensitive on trimmed strings.
