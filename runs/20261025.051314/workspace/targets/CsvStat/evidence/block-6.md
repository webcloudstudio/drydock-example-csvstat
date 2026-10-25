# Evidence: Block 6 · Feature (block-6)

- block type: block
- date: 2026-10-25
- resulting state: closed/verified
- story points (combined assembled cost): 5418
- execution id: 20261025.055938.959Z-572e3679

## Stories built
- Pass the supplied grader sources/run.sh with all 22 checks. (release-verification) [story]

## Acceptance tooling authorization
- VERIFICATION.md#release-grader-passes: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- stack: go_compact.md (SP 2105)
- compass: COMPASS.md (SP 963)
- implements: VERIFICATION.md (SP 628)
- stack: common_compact.md (SP 1653)

## Pre-build acceptance observation
- GREEN (prepassed): release-grader-passes (VERIFICATION.md)
  intent: The supplied grader sources/run.sh, run unfiltered, exits zero.
  return code: 0
  stdout:
    PASS A1: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/employees.csv
    PASS A2: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS A3: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv
    PASS A4: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/header_only.csv
    PASS B1: csvstat --column salary /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/employees.csv
    PASS B2: csvstat --column notes /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/employees.csv
    PASS B3: csvstat --column temp_c /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS C1: csvstat --freq department /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/employees.csv
    PASS C2: csvstat --freq sensor --top 2 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS C3: csvstat --freq customer /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv
    PASS C4: csvstat --freq manager_id --top 3 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/employees.csv
    PASS C5: csvstat --freq battery /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS E1: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/ragged.csv
    PASS E2: csvstat --column nosuch /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv
    PASS E3: csvstat --freq nosuch /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv
    PASS E4: csvstat /tmp/tmp.ARG33xZrIP/no-such-file.csv
    PASS U1: csvstat 
    PASS U2: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS U3: csvstat --freq sensor --top 0 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS U4: csvstat --top 3 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS U5: csvstat --column qty --freq qty /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv
    PASS U6: csvstat --frobnicate /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv
    CsvStat: 22 passed, 0 failed (of 22)

## Post-build programmatic acceptance
- PASS: release-grader-passes (VERIFICATION.md)
  intent: The supplied grader sources/run.sh, run unfiltered, exits zero.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    PASS A1: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/employees.csv
    PASS A2: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS A3: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv
    PASS A4: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/header_only.csv
    PASS B1: csvstat --column salary /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/employees.csv
    PASS B2: csvstat --column notes /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/employees.csv
    PASS B3: csvstat --column temp_c /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS C1: csvstat --freq department /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/employees.csv
    PASS C2: csvstat --freq sensor --top 2 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS C3: csvstat --freq customer /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv
    PASS C4: csvstat --freq manager_id --top 3 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/employees.csv
    PASS C5: csvstat --freq battery /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS E1: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/ragged.csv
    PASS E2: csvstat --column nosuch /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv
    PASS E3: csvstat --freq nosuch /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv
    PASS E4: csvstat /tmp/tmp.Iu4iJrbDIw/no-such-file.csv
    PASS U1: csvstat 
    PASS U2: csvstat /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS U3: csvstat --freq sensor --top 0 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS U4: csvstat --top 3 /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/sensors.csv
    PASS U5: csvstat --column qty --freq qty /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv
    PASS U6: csvstat --frobnicate /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.051314/build/CsvStat/sources/orders.csv
    CsvStat: 22 passed, 0 failed (of 22)

## Build summary
I'll start by inspecting the build directory and the grader script.The tree is clean and the earlier blocks are committed. I'll run the grader unfiltered and the project's test gate.The grader passes all 22 checks. The test gate reported cached results, so I'll rerun it uncached and confirm nothing in the tree changed.RESULT: SUCCESS

FILES CHANGED:
- (none)

SUMMARY:
This is a verification-only block. The existing implementation already satisfies it, so I made no changes.

- `sh sources/run.sh` from the application root exited 0 and reported `CsvStat: 22 passed, 0 failed (of 22)`. The grader builds `./cmd/csvstat` with `go build` and runs all 22 checks (A1–A4, B1–B3, C1–C5, E1–E4, U1–U6).
- `sh bin/test.sh` (`go vet ./...` then `go test ./...`) exited 0. Its test results were cached, so I reran `go test -count=1 ./...` uncached. All four packages with tests (`internal/cli`, `internal/csvdata`, `internal/freq`, `internal/profile`) passed.
- `git status` is clean, so no file under `sources/` or elsewhere was modified.

BLOCKERS:
- None.
