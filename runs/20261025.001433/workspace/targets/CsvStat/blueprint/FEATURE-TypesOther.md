# FEATURE: Types Other

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infer the boolean, text, and empty column types and compose the final type decision. |
| Depends On  | FEATURE-TypesNumeric.md, FEATURE-TypesDate.md |
| Provides    | profile field type values boolean, text, empty |
| Consumes    | integer and number tests from FEATURE-TypesNumeric.md, date test from FEATURE-TypesDate.md |

## Purpose

Complete the type vocabulary and compose the single type decision reported in the `type` field of each profile row.

## Behavior

- `boolean`: every non-null value equals `true` or `false`, ignoring case (`True`, `FALSE`, `tRuE` qualify). Values such as `yes`, `1`, `t` do not.
- `text`: no earlier type fits.
- `empty`: the column has no non-null value (including a column of zero data records).
- The final type is the first of `integer`, `number`, `date`, `boolean`, `text` that every non-null value satisfies; `empty` replaces that decision when no non-null value exists.
- A column mixing types, such as integers and dates, or numbers and booleans, is `text`.
- The decision is a single function in `internal/profile` that the statistics stories branch on.

## Programmatic Acceptance

=== AC types-other-boolean ===
Intent: Columns of true and false in any letter case, with nulls, are typed boolean.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def profile(lines):
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write("".join(line + "\n" for line in lines))
        code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}


values = ["true", "False", "TRUE", "false", "tRuE", "NA", ""]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
assert rows["v"]["type"] == "boolean"
=== END AC types-other-boolean ===

=== AC types-other-not-boolean ===
Intent: Boolean-like words that are not true or false make the column text.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def profile(lines):
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write("".join(line + "\n" for line in lines))
        code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}


for bad in ["yes", "no", "t", "f", "0", "1", "truthy"]:
    values = ["true", "false", bad]
    lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
    rows = profile(lines)
    assert rows["v"]["type"] == "text", (bad, rows["v"]["type"])
=== END AC types-other-not-boolean ===

=== AC types-other-text-fallback ===
Intent: A column that mixes types, or holds free text, is typed text.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def profile(lines):
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write("".join(line + "\n" for line in lines))
        code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}


columns = {
    "words": ["alpha", "beta", "gamma"],
    "intdate": ["1", "2025-02-01"],
    "numbool": ["1.5", "true"],
    "intword": ["1", "2", "x"],
}
names = list(columns)
height = max(len(v) for v in columns.values())
lines = [",".join(names)]
for i in range(height):
    lines.append(",".join(columns[n][i] if i < len(columns[n]) else "NA" for n in names))
rows = profile(lines)
for n in names:
    assert rows[n]["type"] == "text", (n, rows[n]["type"])
=== END AC types-other-text-fallback ===

=== AC types-other-empty ===
Intent: A column with no non-null value is typed empty, whatever spellings of null it holds.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner


def profile(lines):
    with tempfile.TemporaryDirectory() as d:
        path = os.path.join(d, "in.csv")
        with open(path, "w", newline="") as f:
            f.write("".join(line + "\n" for line in lines))
        code, out, err = cli_runner(path)
    print(out)
    print(err)
    assert code == 0, (code, out, err)
    return {r["column"]: r for r in csv.DictReader(io.StringIO(out))}


values = ["", "NA", "N/A", "NULL", " ", "null"]
lines = ["id,v"] + [f"{i},{v}" for i, v in enumerate(values)]
rows = profile(lines)
assert rows["v"]["type"] == "empty"
assert rows["id"]["type"] == "integer"
=== END AC types-other-empty ===

=== AC types-other-header-only-empty ===
Intent: Every column of a header-only file is typed empty.

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
assert code == 0, (code, out, err)
parsed = list(csv.DictReader(io.StringIO(out)))
assert [r["column"] for r in parsed] == names
for r in parsed:
    assert r["type"] == "empty"
=== END AC types-other-header-only-empty ===

## User Acceptance

- None.

## Guardrails

- Only `true` and `false` (any letter case) are boolean values; no other spelling is accepted.
- A column with no non-null value is `empty`, never `text`.
- Types are computed from the file examined; no type is hardcoded.
