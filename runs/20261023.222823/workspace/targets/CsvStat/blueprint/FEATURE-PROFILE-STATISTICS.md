# FEATURE: Profile Statistics

| Field       | Value |
|-------------|-------|
| Version     | 20261023 V1 |
| Description | Compute exact profile statistics and decimal formatting for normalized columns. |
| Depends On  | FEATURE-PROFILE-TYPES.md |
| Provides    | rows, nulls, distinct, extrema, sum, mean, median |
| Consumes    | inferred column types, normalized records |

## Purpose

This capability calculates profile statistics after input normalization and type inference.

For every column it reports row count, null count, distinct non-null values, type-specific minimum and maximum, and numeric sum, mean, and median. Numeric calculations use exact arithmetic and are rounded once, half away from zero, to two decimal places when rendered.

Integer values remain plain integers. Non-numeric sum, mean, and median fields remain empty. Boolean extrema are false when any false value exists and true only when all non-null values are true. Text extrema use bytewise ordering.

## Programmatic Acceptance

=== AC profile-statistics-integer ===
Intent: Integer statistics include counts, extrema, exact total, mean, and median.

import csv
import tempfile
from pathlib import Path
from drydock_go import cli_runner

rows = [
    ["value"],
    ["-3"],
    ["1"],
    ["4"],
    ["NA"],
]
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text("\n".join(",".join(row) for row in rows) + "\n")
    code, output, error = cli_runner(str(path), package="./cmd/csvstat")
    assert code == 0
    parsed = list(csv.reader(output.splitlines()))
    header, result = parsed[0], parsed[1]
    assert result[0] == "value"
    assert result[1] == "integer"
    assert int(result[2]) == len(rows) - 1
    assert int(result[3]) == 1
    values = [-3, 1, 4]
    assert int(result[4]) == len(set(values))
    assert int(result[5]) == min(values)
    assert int(result[6]) == max(values)
    assert int(result[7]) == sum(values)
    assert result[8] == "0.67"
    assert result[9] == "1.00"
=== END AC profile-statistics-integer ===

=== AC profile-statistics-decimal ===
Intent: Decimal statistics preserve exact values and round mean and median to two places.

import csv
import tempfile
from decimal import Decimal, ROUND_HALF_UP
from pathlib import Path
from drydock_go import cli_runner

values = [Decimal("-1.01"), Decimal("-1.00"), Decimal("2.00"), Decimal("2.01")]
source = "amount\n" + "\n".join(str(value) for value in values) + "\n"
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    code, output, error = cli_runner(str(path), package="./cmd/csvstat")
    assert code == 0
    result = list(csv.reader(output.splitlines()))[1]
    total = sum(values)
    mean = (total / len(values)).quantize(Decimal("0.01"), rounding=ROUND_HALF_UP)
    median = ((values[1] + values[2]) / 2).quantize(Decimal("0.01"), rounding=ROUND_HALF_UP)
    assert result[1] == "number"
    assert result[7] == f"{total:.2f}"
    assert result[8] == f"{mean:.2f}"
    assert result[9] == f"{median:.2f}"
=== END AC profile-statistics-decimal ===

=== AC profile-statistics-nonnumeric ===
Intent: Non-numeric columns retain counts and extrema while leaving numeric fields empty.

import csv
import tempfile
from pathlib import Path
from drydock_go import cli_runner

source = "flag\nfalse\ntrue\nfalse\n"
with tempfile.TemporaryDirectory() as directory:
    path = Path(directory) / "input.csv"
    path.write_text(source)
    code, output, error = cli_runner(str(path), package="./cmd/csvstat")
    assert code == 0
    result = list(csv.reader(output.splitlines()))[1]
    assert result[1] == "boolean"
    assert int(result[2]) == 3
    assert int(result[3]) == 0
    assert int(result[4]) == 2
    assert result[5] == "false"
    assert result[6] == "true"
    assert result[7:10] == ["", "", ""]
=== END AC profile-statistics-nonnumeric ===

## User Acceptance

- None.

## Guardrails

- Statistics must be computed from normalized input values, never from hardcoded fixture answers.
- Decimal calculations must use exact arithmetic and round only during output formatting.
