# FEATURE: CLI Unknown Column

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Look up a named column by exact trimmed header name and reject a name absent from the header with exit 3. |
| Depends On  | FEATURE-Profile-Rows.md |
| Provides    | column lookup, exit 3 for unknown column |
| Consumes    | csvdata.Read |

## Purpose

`--column <name>` and `--freq <name>` each name a column. This feature provides the single lookup
both options share and the rejection when the name is not a header column.

## Behavior

- The lookup compares the requested name with the trimmed header names by exact string equality.
  It is case-sensitive.
- When the name is absent, the command exits 3, writes nothing to standard output, and writes one
  diagnostic line to standard error naming the file and the column.
- The lookup runs after the file has been read and validated, so a file rejected for another
  reason (empty, bad header, ragged record) reports that rejection first.
- A name that is present proceeds to the normal profile or frequency path with exit 0.
- Usage validation (exit 2) completes before the file is opened and therefore before this lookup.

## Interfaces

- Package `internal/csvdata` (or `internal/cli`) exposes a lookup returning the column index and a
  found flag; `internal/cli.Run` maps "not found" to exit 3.

## Programmatic Acceptance

=== AC cli003-column-unknown ===
Intent: A --column name absent from the header exits 3 with nothing on standard output.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

names = ["id", "qty"]
body = io.StringIO()
writer = csv.writer(body, lineterminator="\n")
writer.writerow(names)
writer.writerow(["1", "10"])
writer.writerow(["2", "20"])

missing = "nosuch"
assert missing not in names

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "t.csv")
    with open(path, "w", newline="") as f:
        f.write(body.getvalue())
    code, out, err = cli_runner("--column", missing, path)
print(out)
print(err)
assert code == 3
assert out == ""
=== END AC cli003-column-unknown ===

=== AC cli003-freq-unknown ===
Intent: A --freq name absent from the header exits 3 with nothing on standard output.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

names = ["id", "qty"]
body = io.StringIO()
writer = csv.writer(body, lineterminator="\n")
writer.writerow(names)
writer.writerow(["1", "10"])

missing = "nosuch"
assert missing not in names

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "t.csv")
    with open(path, "w", newline="") as f:
        f.write(body.getvalue())
    code, out, err = cli_runner("--freq", missing, path)
print(out)
print(err)
assert code == 3
assert out == ""
=== END AC cli003-freq-unknown ===

=== AC cli003-lookup-exact ===
Intent: The lookup is exact and case-sensitive, and matches trimmed header names.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

present = "qty"
wrong_case = present.upper()
assert wrong_case != present

body = io.StringIO()
writer = csv.writer(body, lineterminator="\n")
writer.writerow(["id", " " + present + " "])
writer.writerow(["1", "10"])

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "t.csv")
    with open(path, "w", newline="") as f:
        f.write(body.getvalue())
    code_ok, out_ok, err_ok = cli_runner("--column", present, path)
    code_bad, out_bad, err_bad = cli_runner("--column", wrong_case, path)
print(out_ok)
print(err_ok)
print(err_bad)
assert code_ok == 0
columns = [r["column"] for r in csv.DictReader(io.StringIO(out_ok))]
assert present in columns
assert code_bad == 3
assert out_bad == ""
=== END AC cli003-lookup-exact ===

=== AC cli003-rejection-precedes-lookup ===
Intent: A ragged file reports its own rejection (exit 3) and writes nothing even when the column name is valid.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

header = ["a", "b"]
body = io.StringIO()
writer = csv.writer(body, lineterminator="\n")
writer.writerow(header)
writer.writerow(["1", "2"])
writer.writerow(["3", "4", "5"])

with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "t.csv")
    with open(path, "w", newline="") as f:
        f.write(body.getvalue())
    code, out, err = cli_runner("--column", header[0], path)
print(out)
print(err)
assert code == 3
assert out == ""
=== END AC cli003-rejection-precedes-lookup ===

## User Acceptance

- None.

## Guardrails

- An unknown column is exit 3 (rejected), never exit 2 and never exit 1.
- A failing command writes nothing to standard output.
- Application code never reads `sources/run.sh` or any supplied data file.
