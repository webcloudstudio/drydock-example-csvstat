# CsvStat

A Drydock UAT kit that builds a Go CSV column profiler. The brief is original to this kit.
Complexity: 2 of 5, about 6 stories.

| Path | Role |
|---|---|
| `sources/INSTRUCTIONS.md` | The specification imported into the Target. |
| `sources/*.csv` | Five data files generated from seed 16 by `tools/gen_csv.py`. |
| `bin/run.sh` | The grader. Staged into the build as `sources/run.sh`; `acceptance.full` runs it. |
| `tools/` | The fixture generator and the exact-arithmetic reference that produced the answers. Not imported. |
| `inputs/` | Frozen `SEA_TRIALS.md`, `TECHNOLOGY_STACK.md`, `COMPASS.md`, `METADATA.md`. |

## Grading

`bin/run.sh [APP]` runs `go build ./cmd/csvstat` in `APP` (default: the current directory) and
makes 22 checks of exit code and standard output:

| Group | Checks |
|---|---|
| A | full profile of each data file |
| B | `--column` |
| C | `--freq`, with and without `--top`, including count ties and an all-null column |
| E | ragged file, unknown column, unreadable file |
| U | usage errors |

## Data traps

| File | Trap |
|---|---|
| `employees.csv` | quoted `"remote, part-time"`; `true`/`TRUE`/`False` count as four distinct booleans; `NA` and blank manager ids are null |
| `sensors.csv` | `N/A` temperatures; `NULL`, `null`, blank, and a single space in an all-null column (`empty`); four sensors tied at 9 |
| `orders.csv` | `"  Globex  "` and `"Acme Corp "` trim into existing values; `ref` mixes `12`, `12.5`, `A7` (text); `shipped` holds `2025-02-30` (text) |
| `header_only.csv` | zero rows; every column `empty` |
| `ragged.csv` | line 4 has an extra field; exit `3` |

The seed was chosen so that no mean or median lands on a half cent, so the answers do not depend
on the rounding mode. Regenerate with `python3 tools/gen_csv.py sources 16` and recompute with
`tools/csvstat.py`.
