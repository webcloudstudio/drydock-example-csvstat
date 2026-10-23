# FEATURE: Profile Output

| Field       | Value |
|-------------|-------|
| Version     | 20261023 V1 |
| Description | Render profile results as RFC-compliant CSV output. |
| Depends On  | FEATURE-PROFILE-STATISTICS.md |
| Provides    | profile CSV output |
| Consumes    | rows, nulls, distinct, extrema, sum, mean, median |

## Purpose

The profile renderer writes one CSV header followed by one row per selected column in input order. A named column limits output to that column. Empty statistic values remain empty CSV fields, and fields containing commas, quotes, or newlines are encoded with RFC-compliant CSV quoting.

## Programmatic Acceptance

=== AC profile-output-all-columns ===
Intent: A profile command emits one output row for each input column in file order.

import csv
import tempfile
from pathlib import Path
from drydock_go import cli_runner

source = "first,second\n1,alpha\n2,beta\n"
expected_columns = ["first", "second"]
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    code, output, error = cli_runner(str(path), package="./cmd/csvstat")
    assert code == 0
    parsed = list(csv.reader(output.splitlines()))
    assert parsed[0][0] == "column"
    assert len(parsed) == len(expected_columns) + 1
    assert [row[0] for row in parsed[1:]] == expected_columns
=== END AC profile-output-all-columns ===

=== AC profile-output-selected-column ===
Intent: The column option emits exactly the requested profile row.

import csv
import tempfile
from pathlib import Path
from drydock_go import cli_runner

source = "first,second\n1,alpha\n2,beta\n"
selected = "second"
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    code, output, error = cli_runner("--column", selected, str(path), package="./cmd/csvstat")
    assert code == 0
    parsed = list(csv.reader(output.splitlines()))
    assert len(parsed) == 2
    assert parsed[1][0] == selected
    assert parsed[1][1] == "text"
=== END AC profile-output-selected-column ===

=== AC profile-output-quoting ===
Intent: Profile output preserves comma-containing values through CSV parsing.

import csv
import tempfile
from pathlib import Path
from drydock_go import cli_runner

value = "remote, part-time"
source = "name,notes\nAda," + '"' + value + '"' + "\n"
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    code, output, error = cli_runner("--column", "notes", str(path), package="./cmd/csvstat")
    assert code == 0
    parsed = list(csv.reader(output.splitlines()))
    assert parsed[1][5] == value
=== END AC profile-output-quoting ===

## User Acceptance

- None.

## Guardrails

- Output must be written as CSV using RFC-compliant quoting and newline handling.
- A failed command must not emit partial profile output.
