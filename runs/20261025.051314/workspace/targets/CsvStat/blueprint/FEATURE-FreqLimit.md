# FEATURE: Frequency Limit

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Limit the number of frequency rows with --top, defaulting to five. |
| Depends On  | FEATURE-FreqCount.md |
| Provides    | freq.Top |
| Consumes    | freq.Count |

## Purpose

After the frequency listing is ordered, at most `--top` rows are written. `--top` is a positive
integer and defaults to 5 when not given.

## Behavior

- The limit applies after ordering, so the rows listed are the first N of the full ordered list.
- A limit larger than the number of distinct values lists every value.
- The header `value,count` is always written on success.
- Validation of `--top` (non-positive, non-integer, without `--freq`) belongs to the usage-errors
  story and exits 2.

## Programmatic Acceptance

=== AC freq-limit-default-five ===
Intent: Without --top the listing has at most five rows, the first five of the ordered list.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from collections import Counter
from drydock_go import cli_runner

values = []
for i in range(7):
    values += [f"v{i}"] * (8 - i)
rows = [["v"]] + [[x] for x in values]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--freq", "v", path)
print(out)
print(err)
assert code == 0, code
parsed = list(csv.reader(io.StringIO(out)))
counts = Counter(values)
ordered = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))
default_top = 5
assert parsed[0] == ["value", "count"]
assert parsed[1:] == [[k, str(n)] for k, n in ordered[:default_top]]
assert len(ordered) > default_top
=== END AC freq-limit-default-five ===

=== AC freq-limit-explicit-top ===
Intent: --top N lists exactly the first N ordered rows.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from collections import Counter
from drydock_go import cli_runner

values = []
for i in range(7):
    values += [f"v{i}"] * (8 - i)
rows = [["v"]] + [[x] for x in values]
counts = Counter(values)
ordered = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    for n in (1, 2, 3):
        code, out, err = cli_runner("--freq", "v", "--top", str(n), path)
        print(out)
        print(err)
        assert code == 0, code
        parsed = list(csv.reader(io.StringIO(out)))
        assert parsed[0] == ["value", "count"]
        assert parsed[1:] == [[k, str(c)] for k, c in ordered[:n]]
=== END AC freq-limit-explicit-top ===

=== AC freq-limit-larger-than-distinct ===
Intent: A --top larger than the distinct count lists every value.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from collections import Counter
from drydock_go import cli_runner

values = ["a", "b", "b", "c", "c", "c"]
rows = [["v"]] + [[x] for x in values]
counts = Counter(values)
ordered = sorted(counts.items(), key=lambda kv: (-kv[1], kv[0].encode("utf-8")))
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--freq", "v", "--top", str(len(ordered) + 50), path)
print(out)
print(err)
assert code == 0, code
parsed = list(csv.reader(io.StringIO(out)))
assert parsed[1:] == [[k, str(n)] for k, n in ordered]
=== END AC freq-limit-larger-than-distinct ===

=== AC freq-limit-all-null-column ===
Intent: A column with no non-null values lists only the header even with --top given.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile
from drydock_go import cli_runner

rows = [["gap"], [""], ["NA"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    code, out, err = cli_runner("--freq", "gap", "--top", "3", path)
print(out)
print(err)
assert code == 0, code
assert list(csv.reader(io.StringIO(out))) == [["value", "count"]]
=== END AC freq-limit-all-null-column ===

## User Acceptance

- None.

## Guardrails

- The limit is applied after ordering, never before.
- The default limit is 5.
