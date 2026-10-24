# FEATURE: Type Inference

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Infer each column's type among integer, number, boolean, text, and empty from its non-null values. |
| Depends On  | FEATURE-NullTrim.md, ARCHITECTURE.md |
| Provides    | typing.Infer |
| Consumes    | reader.IsNull, reader.Trim |

## Purpose

Give every column a type so later stories choose the right min, max, and numeric statistics.

## Behavior

- Inference runs over the trimmed, non-null values of one column. Null values never influence the type.
- A column takes the first of `integer`, `number`, `date`, `boolean`, `text` that every non-null value satisfies. A column with no non-null value is `empty`.
- `integer`: every value matches `-?[0-9]+`.
- `number`: every value matches `-?[0-9]+(\.[0-9]+)?`. A column mixing integer and decimal values is `number`.
- `boolean`: every value equals `true` or `false`, ignoring case.
- `text`: anything else.
- `date` is recognized by `typing.IsDate`, supplied by the date-type story. In this story the date check is a hook, `func(string) bool`, that defaults to returning false, so a column of dates is `text` until that story installs the real check. The hook sits between `number` and `boolean` in the order.
- `typing.Infer(values []string) string` takes already-trimmed values, skips nulls via `reader.IsNull`, and returns the type token.
- The profile output writes the token in the `type` column.

## Programmatic Acceptance

=== AC type-integer-and-number ===
Intent: Integer, mixed integer-and-decimal, and negative columns are typed integer and number as the contract states.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

cols = {
    "ints": ["1", "-20", "300", "0"],
    "mixed": ["1", "2.5", "-3", "4.25"],
    "decs": ["0.5", "-1.75", "10.0", "3.14"],
}
names = list(cols)
nrows = len(cols[names[0]])
lines = [",".join(names)]
for i in range(nrows):
    lines.append(",".join(cols[n][i] for n in names))
text = "\n".join(lines) + "\n"
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = {r["column"]: r for r in csv.DictReader(io.StringIO(out))}
assert rows["ints"]["type"] == "integer"
assert rows["mixed"]["type"] == "number"
assert rows["decs"]["type"] == "number"
=== END AC type-integer-and-number ===

=== AC type-boolean-case-insensitive ===
Intent: A column of true and false in any letter case is boolean, while a column mixing a boolean with other text is text.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

cols = {
    "flag": ["true", "FALSE", "True", "false"],
    "maybe": ["true", "false", "yes", "true"],
}
names = list(cols)
nrows = len(cols[names[0]])
lines = [",".join(names)]
for i in range(nrows):
    lines.append(",".join(cols[n][i] for n in names))
text = "\n".join(lines) + "\n"
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = {r["column"]: r for r in csv.DictReader(io.StringIO(out))}
assert rows["flag"]["type"] == "boolean"
assert rows["maybe"]["type"] == "text"
=== END AC type-boolean-case-insensitive ===

=== AC type-text-fallback ===
Intent: A column with one non-numeric value among integers falls back to text.

import csv
import io
import os
import tempfile

from drydb_placeholder import nothing  # noqa: F401
=== END AC type-text-fallback ===

=== AC type-nulls-ignored ===
Intent: Null spellings (empty, NA, N/A, NULL in any case) do not change the inferred type, and an all-null column is empty.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

cols = {
    "n": ["5", "", "NA", "n/a", "Null", "7"],
    "allnull": ["", "NA", "N/A", "NULL", "null", " "],
    "word": ["red", "", "blue", "NA", "green", "NULL"],
}
names = list(cols)
nrows = len(cols[names[0]])
lines = [",".join(names)]
for i in range(nrows):
    lines.append(",".join(cols[n][i] for n in names))
text = "\n".join(lines) + "\n"
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="") as fh:
        fh.write(text)
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = {r["column"]: r for r in csv.DictReader(io.StringIO(out))}
assert rows["n"]["type"] == "integer"
assert rows["allnull"]["type"] == "empty"
assert rows["word"]["type"] == "text"
=== END AC type-nulls-ignored ===

=== AC type-header-only-empty ===
Intent: A file with a header and no data records types every column as empty.

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

names = ["a", "b", "c"]
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
assert all(r["type"] == "empty" for r in rows)
=== END AC type-header-only-empty ===

## User Acceptance

- None.

## Guardrails

- Types are computed from the file given; no column name, type, or value list is hardcoded.
- Application code never reads `sources/run.sh` or any supplied data file.
- Standard library only.
