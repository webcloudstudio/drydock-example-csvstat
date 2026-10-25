# FEATURE: Type Text and Empty

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Complete the column type ladder with the text fallback and the empty type for columns without non-null values. |
| Depends On  | FEATURE-Type-Boolean.md, FEATURE-Null-Values.md |
| Provides    | type text, type empty |
| Consumes    | csvdata.IsNull |

## Purpose

Finish the first-match type ladder `integer`, `number`, `date`, `boolean`, `text`, and add `empty`. After this story every column of every readable file has a type token in the `type` field of the profile row.

## Behavior

- A column with no non-null value has type `empty`. This includes a column whose values are all null (empty, `NA`, `N/A`, `NULL` in any case) and every column of a file with a header and no data records.
- Otherwise the column takes the first of `integer`, `number`, `date`, `boolean` that every non-null value satisfies; if none is satisfied the type is `text`.
- Nulls never decide the type; they are removed before the ladder runs.
- Values are trimmed before classification.
- A date-shaped value that is not a real calendar date (`2025-02-30`) does not satisfy `date`; with no other match the column is `text`.
- Forms such as `1.`, `1e3`, `yes`, and `no` satisfy no numeric or boolean rule and fall to `text`.

## Suite Pattern

The project's own `go test` suite covers the ladder directly in `internal/infer` (one test per type, null-only column, empty slice, mixed columns, order of precedence).

## Programmatic Acceptance

=== AC typ005-ladder ===
Intent: Each column resolves to the first matching type in the ladder, with text as the fallback.
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

cases = {
    "ints": (["1", "-2", "30"], "integer"),
    "nums": (["1", "2.5", "-3.25"], "number"),
    "dates": (["2025-01-31", "2024-02-29", "2025-12-01"], "date"),
    "bools": (["true", "FALSE", "False"], "boolean"),
    "texts": (["abc", "1", "2025-01-01"], "text"),
}
columns = {name: values for name, (values, _) in cases.items()}
rows = run_profile(columns)
for name, (values, expected) in cases.items():
    assert rows[name]["type"] == expected, (name, rows[name])
    assert rows[name]["rows"] == str(len(values)), (name, rows[name])
=== END AC typ005-ladder ===

=== AC typ005-nulls-do-not-decide ===
Intent: Null values and padding are ignored when choosing the type, and an all-null column is empty.
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

cases = {
    "allnull": (["", "NA", "n/a", "NULL"], "empty"),
    "ints_nulls": (["7", "", "NULL", "8"], "integer"),
    "bools_nulls": (["TRUE", "na", "", "false"], "boolean"),
    "padded": ([" 5 ", "6 ", "  7", "N/A"], "integer"),
    "spaces_only": (["   ", " ", "", "NA"], "empty"),
}
columns = {name: values for name, (values, _) in cases.items()}
rows = run_profile(columns)
for name, (values, expected) in cases.items():
    assert rows[name]["type"] == expected, (name, rows[name])
=== END AC typ005-nulls-do-not-decide ===

=== AC typ005-header-only-empty ===
Intent: Every column of a file with a header and no data records has type empty and zero rows.
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
assert [r["column"] for r in rows] == names
for r in rows:
    assert r["type"] == "empty", r
    assert r["rows"] == str(0), r
=== END AC typ005-header-only-empty ===

=== AC typ005-text-fallback ===
Intent: Values that look close to a typed form but do not satisfy it make the column text.
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

columns = {
    "bad_date": ["2025-02-30", "2025-01-01"],
    "trailing_dot": ["1.", "2"],
    "exponent": ["1e3", "5"],
    "yes_no": ["yes", "no"],
    "bool_maybe": ["true", "maybe"],
    "plus_sign": ["+1", "2"],
}
rows = run_profile(columns)
for name in columns:
    assert rows[name]["type"] == "text", (name, rows[name])
=== END AC typ005-text-fallback ===

## User Acceptance

- None.

## Guardrails

- Do not read `sources/run.sh` or any supplied data file from application code, and do not hardcode a type.
- Typing uses only the standard library.
- A column is `empty` only when it has no non-null value; a column with at least one non-null value is never `empty`.
