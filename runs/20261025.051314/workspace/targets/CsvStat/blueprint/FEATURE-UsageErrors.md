# FEATURE: Usage Errors

| Field       | Value |
|-------------|-------|
| Version     | 20261025 V1 |
| Description | Reject malformed command-line arguments with exit status 2 and a single usage line on standard error. |
| Depends On  | FEATURE-Arguments.md |
| Provides    | exit code 2 |
| Consumes    | cli.ParseArgs |

## Purpose

Malformed invocations exit 2, write nothing to standard output, and write one diagnostic line to
standard error that shows the two command forms:

```text
csvstat [--column <name>] <file>
csvstat --freq <name> [--top <n>] <file>
```

## Behavior

Exit 2 for each of:

- no file argument
- more than one file argument
- an unknown option
- an option given without its value
- `--top` that is not a positive integer (zero, negative, or non-numeric)
- `--top` without `--freq`
- `--column` together with `--freq`
- a repeated option

Usage errors are decided from the arguments alone, before the file is opened, and exit 2 even when
the file exists and is valid.

## Programmatic Acceptance

=== AC usage-errors-exit-two ===
Intent: Every malformed argument form exits 2 and writes nothing to standard output.
Requires: executable=go; scope=test

import csv
import os
import tempfile
from drydock_go import cli_runner

rows = [["qty", "sensor"], ["1", "east"], ["2", "west"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    other = os.path.join(d, "other.csv")
    for p in (path, other):
        with open(p, "w", newline="", encoding="utf-8") as f:
            csv.writer(f, lineterminator="\n").writerows(rows)
    cases = [
        [],
        [path, other],
        ["--frobnicate", path],
        ["--column"],
        ["--freq"],
        ["--freq", "sensor", "--top"],
        ["--freq", "sensor", "--top", "0", path],
        ["--freq", "sensor", "--top", "-1", path],
        ["--freq", "sensor", "--top", "abc", path],
        ["--top", "3", path],
        ["--column", "qty", "--freq", "qty", path],
    ]
    for args in cases:
        code, out, err = cli_runner(*args)
        print(args, code)
        print(err)
        assert code == 2, (args, code)
        assert out == "", (args, out)
=== END AC usage-errors-exit-two ===

=== AC usage-errors-single-stderr-line ===
Intent: A usage error writes exactly one non-blank diagnostic line to standard error.
Requires: executable=go; scope=test

from drydock_go import cli_runner

code, out, err = cli_runner()
print(err)
assert code == 2, code
assert out == ""
lines = [ln for ln in err.splitlines() if ln.strip()]
assert len(lines) == 1, lines
=== END AC usage-errors-single-stderr-line ===

=== AC usage-errors-before-file-open ===
Intent: A usage error is reported as 2 even when the named file does not exist.
Requires: executable=go; scope=test

import os
import tempfile
from drydock_go import cli_runner

with tempfile.TemporaryDirectory() as d:
    missing = os.path.join(d, "no-such-file.csv")
    code, out, err = cli_runner("--top", "3", missing)
print(err)
assert code == 2, code
assert out == ""
=== END AC usage-errors-before-file-open ===

=== AC usage-errors-valid-forms-accepted ===
Intent: Well-formed invocations are not rejected as usage errors.
Requires: executable=go; scope=test

import csv
import os
import tempfile
from drydock_go import cli_runner

rows = [["qty", "sensor"], ["1", "east"], ["2", "west"]]
with tempfile.TemporaryDirectory() as d:
    path = os.path.join(d, "in.csv")
    with open(path, "w", newline="", encoding="utf-8") as f:
        csv.writer(f, lineterminator="\n").writerows(rows)
    for args in ([path], ["--column", "qty", path], ["--freq", "sensor", path],
                 ["--freq", "sensor", "--top", "2", path]):
        code, out, err = cli_runner(*args)
        print(args, code)
        print(err)
        assert code == 0, (args, code)
=== END AC usage-errors-valid-forms-accepted ===

## User Acceptance

- None.

## Guardrails

- A usage error never writes to standard output.
- Usage errors never depend on the contents of the file.
