# FEATURE: CLI CSV Output

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Write every output form as CSV with RFC 4180 quoting and \n line endings. |
| Depends On  | FEATURE-Profile-Extremes.md, FEATURE-Frequency-List.md |
| Provides    | CSV output format |
| Consumes    | profile output, frequency output |

## Purpose

Profile, `--column`, and `--freq` output all pass through one `encoding/csv` writer, so every
output form quotes and terminates records the same way.

## Behavior

- A field containing a comma, a double quote, or a newline is quoted; embedded quotes are doubled.
- Records end with `\n`, never `\r\n`.
- Parsing standard output with a standards-compliant CSV reader yields the intended field values,
  and every record of one output has the same number of fields as its header.

## Programmatic Acceptance

=== AC output-profile-quoted-field-roundtrip ===
Intent: In the employees profile every record parses to the header's field count and the notes max field equals the byte-wise largest non-null notes value.

import csv
import io
from drydock_go import cli_runner

def is_null(v):
    return v == "" or v.lower() in ("na", "n/a", "null")

path = "sources/employees.csv"
with open(path, newline="", encoding="utf-8") as f:
    rows = [[c.strip() for c in r] for r in csv.reader(f)]
i = rows[0].index("notes")
vals = [r[i] for r in rows[1:] if not is_null(r[i])]
want_max = max(vals, key=lambda v: v.encode("utf-8"))
assert "," in want_max

code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
width = len(got[0])
assert all(len(r) == width for r in got)
by_name = {r[0]: r for r in got[1:]}
head = got[0]
assert by_name["notes"][head.index("max")] == want_max
=== END AC output-profile-quoted-field-roundtrip ===

=== AC output-freq-quoted-values ===
Intent: Frequency rows for a column with comma-bearing values parse to two fields each and carry the values found in the file.

import csv
import io
from drydock_go import cli_runner

def is_null(v):
    return v == "" or v.lower() in ("na", "n/a", "null")

path = "sources/employees.csv"
with open(path, newline="", encoding="utf-8") as f:
    rows = [[c.strip() for c in r] for r in csv.reader(f)]
i = rows[0].index("notes")
counts = {}
for r in rows[1:]:
    if not is_null(r[i]):
        counts[r[i]] = counts.get(r[i], 0) + 1
want = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))
assert any("," in v for v, _ in want)

code, out, err = cli_runner("--freq", "notes", "--top", str(len(want)), path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert all(len(r) == 2 for r in got)
assert got[1:] == [[v, str(c)] for v, c in want]
=== END AC output-freq-quoted-values ===

=== AC output-newline-and-quote-roundtrip ===
Intent: A value containing a double quote and a newline survives --freq and --column output and parses back to the value supplied.

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

tricky = 'say "hi"\nthen, bye'
d = tempfile.mkdtemp()
path = os.path.join(d, "tricky.csv")
with open(path, "w", newline="", encoding="utf-8") as f:
    w = csv.writer(f, lineterminator="\n")
    w.writerow(["note"])
    w.writerow([tricky])
    w.writerow([tricky])

code, out, err = cli_runner("--freq", "note", path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert len(got) == 2
assert got[1] == [tricky, str(2)]

code, out, err = cli_runner("--column", "note", path)
print(out)
print(err)
assert code == 0
got = list(csv.reader(io.StringIO(out)))
assert len(got) == 2
head = got[0]
assert len(got[1]) == len(head)
assert got[1][head.index("min")] == tricky
assert got[1][head.index("max")] == tricky
=== END AC output-newline-and-quote-roundtrip ===

=== AC output-line-endings-are-lf ===
Intent: Profile and frequency output records end with a line feed and never a carriage return.

from drydock_go import cli_runner

lf = chr(10)
cr = chr(13)
runs = [
    ("sources/employees.csv",),
    ("--column", "notes", "sources/employees.csv"),
    ("--freq", "department", "sources/employees.csv"),
]
for args in runs:
    code, out, err = cli_runner(*args)
    print(out)
    print(err)
    assert code == 0, args
    lines = out.splitlines(keepends=True)
    assert lines
    assert all(line.endswith(lf) for line in lines), args
    assert all(not line.endswith(cr + lf) for line in lines), args
=== END AC output-line-endings-are-lf ===

## User Acceptance

- None.

## Guardrails

- All output goes through one `encoding/csv` writer; no output form is assembled by string concatenation.
- Nothing is written to standard output when a command fails.
