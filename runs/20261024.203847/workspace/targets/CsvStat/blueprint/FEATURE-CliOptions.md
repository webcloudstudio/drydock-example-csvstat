# FEATURE: CLI Options

| Field       | Value |
|-------------|-------|
| Version     | 20261024 V1 |
| Description | Parse the --column, --freq, and --top options and the single file argument into a request value. |
| Depends On  | ARCHITECTURE.md |
| Provides    | --column, --freq, --top |
| Consumes    | cli.Run |

## Purpose

`internal/cli` parses arguments with a local `flag.FlagSet` created with `flag.ContinueOnError` and produces a request value that the profile and frequency stories consume. This story owns the accepted forms; the rejection of invalid forms belongs to the usage-errors story.

## Command Forms

```text
csvstat [--column <name>] <file>
csvstat --freq <name> [--top <n>] <file>
```

| Option | Meaning |
|---|---|
| `--column <name>` | Profile only the named column. |
| `--freq <name>` | List the most frequent values of the named column instead of profiling. |
| `--top <n>` | With `--freq`, list at most `<n>` values. A positive integer; default `5`. |

## Behavior

- Options precede exactly one file argument.
- With no option, every column is profiled.
- `--column <name>` produces a profile containing only that column.
- `--freq <name>` produces the frequency list with header `value,count`.
- `--top` defaults to 5 when `--freq` is given without it.
- The parsed request carries: file path, column name (optional), freq name (optional), and top count.
- Only `main` calls `os.Exit`; `Run` returns the exit code.

## Programmatic Acceptance

=== AC options-no-option-profiles-all ===
Intent: With only a file argument, every column is profiled in file order.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

names = ["alpha", "beta", "gamma"]
work = tempfile.mkdtemp()
path = os.path.join(work, "data.csv")
with open(path, "w", encoding="utf-8") as handle:
    handle.write(",".join(names) + "\n1,2,3\n4,5,6\n")

code, out, err = cli_runner(path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert [row[0] for row in rows[1:]] == names, rows
=== END AC options-no-option-profiles-all ===

=== AC options-column-profiles-one ===
Intent: --column profiles exactly the named column.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

names = ["alpha", "beta", "gamma"]
chosen = names[1]
work = tempfile.mkdtemp()
path = os.path.join(work, "data.csv")
with open(path, "w", encoding="utf-8") as handle:
    handle.write(",".join(names) + "\n1,2,3\n4,5,6\n")

code, out, err = cli_runner("--column", chosen, path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) == 2, rows
assert rows[1][0] == chosen, rows
=== END AC options-column-profiles-one ===

=== AC options-freq-default-top ===
Intent: --freq without --top lists at most five values.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

default_top = 5
distinct = 8
work = tempfile.mkdtemp()
path = os.path.join(work, "data.csv")
with open(path, "w", encoding="utf-8") as handle:
    handle.write("v\n")
    for index in range(distinct):
        handle.write("item%d\n" % index)

code, out, err = cli_runner("--freq", "v", path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) - 1 == default_top, rows
=== END AC options-freq-default-top ===

=== AC options-freq-explicit-top ===
Intent: --top limits the frequency list to the requested number of rows.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

limit = 2
distinct = 6
work = tempfile.mkdtemp()
path = os.path.join(work, "data.csv")
with open(path, "w", encoding="utf-8") as handle:
    handle.write("v\n")
    for index in range(distinct):
        handle.write("item%d\n" % index)

code, out, err = cli_runner("--freq", "v", "--top", str(limit), path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) - 1 == limit, rows
=== END AC options-freq-explicit-top ===

=== AC options-freq-top-larger-than-distinct ===
Intent: A --top larger than the number of distinct values lists every distinct value.
Requires: executable=go; scope=test

import csv
import io
import os
import tempfile

from drydock_go import cli_runner

distinct = 3
limit = 50
work = tempfile.mkdtemp()
path = os.path.join(work, "data.csv")
with open(path, "w", encoding="utf-8") as handle:
    handle.write("v\n")
    for index in range(distinct):
        handle.write("item%d\n" % index)

code, out, err = cli_runner("--freq", "v", "--top", str(limit), path)
print(out)
print(err)
assert code == 0, (code, out, err)
rows = list(csv.reader(io.StringIO(out)))
assert len(rows) - 1 == distinct, rows
=== END AC options-freq-top-larger-than-distinct ===

## User Acceptance

- None.

## Guardrails

- Only `main` calls `os.Exit`; `Run` returns an exit code.
- Parsing uses a local `flag.FlagSet` with `flag.ContinueOnError`, never the global flag set.
- Options always precede the file argument.
- Only the Go standard library is used.
