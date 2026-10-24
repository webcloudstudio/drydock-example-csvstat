# FEATURE: Reject Malformed Input

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Reject empty files, empty or repeated header names, and ragged records with exit code 3 and no standard output. |
| Depends On  | FEATURE-NullTrim.md, ARCHITECTURE.md |
| Provides    | exit 3 on rejected input |
| Consumes    | reader.Read, reader.Trim |

## Purpose

After the file is read and trimmed, csvstat validates its shape. A rejected file exits 3, writes nothing to standard output, and writes one diagnostic line to standard error naming the file and, for a ragged record, its line number.

## Rejection Rules

1. The file is empty (no records at all).
2. A header name is empty after trimming.
3. A header name repeats another after trimming (exact comparison).
4. Any data record has a different number of fields from the header.

Field-count checking is done here, not by `encoding/csv`, so the diagnostic can name the offending line. A header-only file (header and no records) is valid and is not rejected.

## Behavior

- Validation runs before any statistic is computed, so no partial output is ever written.
- Exit code is 3 for every rejection above.
- Standard output is empty on rejection.
- Standard error carries one diagnostic line with the file path and, for a ragged record, the line number.

## Programmatic Acceptance

=== AC reject-ragged-fixture ===
Intent: The supplied ragged fixture is rejected with exit 3 and nothing on standard output.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("sources/ragged.csv")
print(out)
print(err)
assert code == 3, (code, out, err)
assert out == "", out
=== END AC reject-ragged-fixture ===

=== AC reject-empty-file ===
Intent: A zero-byte file is rejected with exit 3 and nothing on standard output.
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

work = tempfile.mkdtemp()
path = os.path.join(work, "empty.csv")
with open(path, "w", encoding="utf-8") as handle:
    handle.write("")

code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 3, (code, out, err)
assert out == "", out
=== END AC reject-empty-file ===

=== AC reject-empty-header-name ===
Intent: A header with an empty or blank name is rejected with exit 3.
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

work = tempfile.mkdtemp()
for label, header in (("blank", "a, ,c"), ("empty", "a,,c")):
    path = os.path.join(work, label + ".csv")
    with open(path, "w", encoding="utf-8") as handle:
        handle.write(header + "\n1,2,3\n")
    code, out, err = cli_runner(path)
    print(label, code, out, err)
    assert code == 3, (label, code, out, err)
    assert out == "", (label, out)
=== END AC reject-empty-header-name ===

=== AC reject-repeated-header-name ===
Intent: A repeated header name, including one that repeats only after trimming, is rejected with exit 3.
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

work = tempfile.mkdtemp()
for label, header in (("plain", "a,b,a"), ("padded", "a,b, a ")):
    path = os.path.join(work, label + ".csv")
    with open(path, "w", encoding="utf-8") as handle:
        handle.write(header + "\n1,2,3\n")
    code, out, err = cli_runner(path)
    print(label, code, out, err)
    assert code == 3, (label, code, out, err)
    assert out == "", (label, out)
=== END AC reject-repeated-header-name ===

=== AC reject-short-and-long-records ===
Intent: A record with too few fields and a record with too many fields are each rejected with exit 3.
Requires: executable=go; scope=test

import os
import tempfile

from drydock_go import cli_runner

work = tempfile.mkdtemp()
cases = {
    "short": "a,b,c\n1,2,3\n4,5\n",
    "long": "a,b,c\n1,2,3\n4,5,6,7\n",
}
for label, body in cases.items():
    path = os.path.join(work, label + ".csv")
    with open(path, "w", encoding="utf-8") as handle:
        handle.write(body)
    code, out, err = cli_runner(path)
    print(label, code, out, err)
    assert code == 3, (label, code, out, err)
    assert out == "", (label, out)
=== END AC reject-short-and-long-records ===

=== AC accept-header-only ===
Intent: A header-only file is valid and is not rejected.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner("sources/header_only.csv")
print(out)
print(err)
assert code == 0, (code, out, err)
=== END AC accept-header-only ===

## User Acceptance

- The standard error diagnostic for a ragged record names the file and the line number of the offending record.

## Guardrails

- A rejected file never writes anything to standard output.
- Invalid UTF-8 bytes are not a rejection condition.
- The rejection list is exhaustive: empty file, empty header name, repeated header name, ragged record, and (in other stories) an absent column name.
