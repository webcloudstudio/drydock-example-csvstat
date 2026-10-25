# FEATURE: Type Boolean

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Infer boolean columns whose non-null values are true or false in any letter case. |
| Depends On  | FEATURE-Type-Date.md |
| Provides    | type boolean |
| Consumes    | csvdata.IsNull |

## Purpose

Extend the type ladder so columns of true/false values are typed `boolean`.

## Behavior

- A column that is not `integer`, `number`, or `date` is `boolean` when every non-null trimmed
  value equals `true` or `false` ignoring case.
- `1`, `0`, `yes`, `no`, `t`, and `f` are not boolean. A column of `0` and `1` is `integer`
  because the ladder takes the first match.
- Null values are ignored; a column with no non-null value is never `boolean`.

## Programmatic Acceptance

=== AC typ004-boolean-detected ===
Intent: Columns of true/false in any letter case, with padding and nulls, are typed boolean.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

columns = {
    "mixed_case": ["true", "FALSE", "True", "false"],
    "upper": ["TRUE", "FALSE", "TRUE", "FALSE"],
    "padded": [" true ", "false", "false", "true"],
    "with_nulls": ["true", "", "NA", "N/A"],
    "single_value": ["false", "false", "false", "false"],
}
names = list(columns)
count = len(columns[names[0]])
body = io.StringIO()
writer = csv.writer(body, lineterminator="\n")
writer.writerow(names)
for i in range(count):
    writer.writerow([columns[k][i] for k in names])

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "t.csv")
    with open(path, "w", newline="") as f:
        f.write(body.getvalue())
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0
rows = {r["column"]: r for r in csv.DictReader(io.StringIO(out))}
for name in names:
    assert rows[name]["type"] == "boolean", name
=== END AC typ004-boolean-detected ===

=== AC typ004-non-boolean-rejected ===
Intent: Yes/no words, a stray word among booleans, and short forms are not boolean, and a 0/1 column stays integer.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

columns = {
    "zero_one": ["0", "1", "1", "0"],
    "yes_no": ["yes", "no", "yes", "no"],
    "stray": ["true", "false", "maybe", "true"],
    "short": ["t", "f", "t", "f"],
}
names = list(columns)
count = len(columns[names[0]])
body = io.StringIO()
writer = csv.writer(body, lineterminator="\n")
writer.writerow(names)
for i in range(count):
    writer.writerow([columns[k][i] for k in names])

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "t.csv")
    with open(path, "w", newline="") as f:
        f.write(body.getvalue())
    code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0
rows = {r["column"]: r for r in csv.DictReader(io.StringIO(out))}
assert rows["zero_one"]["type"] == "integer"
for name in names[1:]:
    assert rows[name]["type"] != "boolean", name
=== END AC typ004-non-boolean-rejected ===

## User Acceptance

- None.

## Guardrails

- Inference uses only the values of the file examined.
- Integer, number, and date take precedence over boolean whenever their patterns are satisfied.
