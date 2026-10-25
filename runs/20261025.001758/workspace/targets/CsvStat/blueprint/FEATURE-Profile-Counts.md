# FEATURE: Profile Counts

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Compute the null count and the exact distinct count of non-null values for every profiled column. |
| Depends On  | FEATURE-Profile-Rows.md, FEATURE-Null-Values.md |
| Provides    | profile nulls, distinct |
| Consumes    | csvdata.IsNull |

## Purpose

Fill the `nulls` and `distinct` fields of each profile row. The `rows` field is already produced by the profile rows story.

## Behavior

- `nulls` is the number of values in the column that are null: empty, or `NA`, `N/A`, `NULL` in any letter case, after trimming.
- `distinct` is the number of different non-null values compared as exact strings after trimming. Case differences count as different values (`Abc` and `abc`), and so do `1`, `01`, and `1.0`.
- `0`, `none`, and `nan` are not null.
- A column with no non-null values has `distinct` 0; a header-only file reports `nulls` 0 and `distinct` 0 for every column.
- Both counts are written as plain integers.

## Programmatic Acceptance

=== AC prf002-counts ===
Intent: nulls and distinct match counts derived from the fixture values by the stated rules.
Suite: scoped

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

def run_profile(columns, *opts):
    names = list(columns)
    n = len(columns[names[0]])
    buf = io.StringIO()
    w = csv.writer(buf, lineterminator="\n")
    w.writerow(names)
    for i in range(n):
        w.writerow([columns[c][i] for c in names])
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write(buf.getvalue())
        code, out, err = cli_runner(*opts, path)
    print(out)
    print(err)
    assert code == 0, code
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}

def is_null(v):
    return v.strip().lower() in ("", "na", "n/a", "null")

columns = {
    "a": ["x", "x", "y", "", "NA"],
    "b": ["Abc", "abc", "ABC", "NULL", "N/A"],
    "c": [" x ", "x", "x ", "", " "],
    "d": ["1", "01", "1.0", "1", "NULL"],
    "allnull": ["", "", "NA", "null", "n/A"],
}
rows = run_profile(columns)
for name, values in columns.items():
    trimmed = [v.strip() for v in values]
    expected_nulls = sum(1 for v in trimmed if is_null(v))
    expected_distinct = len(set(v for v in trimmed if not is_null(v)))
    assert rows[name]["rows"] == str(len(values)), (name, rows[name])
    assert rows[name]["nulls"] == str(expected_nulls), (name, rows[name])
    assert rows[name]["distinct"] == str(expected_distinct), (name, rows[name])
=== END AC prf002-counts ===

=== AC prf002-non-null-lookalikes ===
Intent: Values such as 0, none, and nan are counted as distinct non-null values.
Suite: scoped

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

def run_profile(columns, *opts):
    names = list(columns)
    n = len(columns[names[0]])
    buf = io.StringIO()
    w = csv.writer(buf, lineterminator="\n")
    w.writerow(names)
    for i in range(n):
        w.writerow([columns[c][i] for c in names])
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write(buf.getvalue())
        code, out, err = cli_runner(*opts, path)
    print(out)
    print(err)
    assert code == 0, code
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}

values = ["0", "0", "none", "nan", ""]
rows = run_profile({"e": values})
expected_nulls = sum(1 for v in values if v == "")
expected_distinct = len(set(v for v in values if v != ""))
assert rows["e"]["nulls"] == str(expected_nulls), rows["e"]
assert rows["e"]["distinct"] == str(expected_distinct), rows["e"]
=== END AC prf002-non-null-lookalikes ===

=== AC prf002-header-only ===
Intent: A header-only file reports zero rows, nulls, and distinct for every column.
Suite: scoped

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

names = ["sku", "price", "in_stock"]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as f:
        f.write(",".join(names) + "\n")
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, code
rows = list(csv.DictReader(io.StringIO(out)))
assert len(rows) == len(names)
for r in rows:
    assert r["rows"] == str(0), r
    assert r["nulls"] == str(0), r
    assert r["distinct"] == str(0), r
=== END AC prf002-header-only ===

## User Acceptance

- None.

## Guardrails

- Distinct compares exact trimmed strings; never normalize case or numeric form.
- Null values are never counted as distinct values.
- Do not hardcode any count; every value comes from the file examined.
