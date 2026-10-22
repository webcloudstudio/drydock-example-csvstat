# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-22
- resulting state: closed/verified
- story points (combined assembled cost): 7869
- execution id: 20261022.203556.991Z-81540f24

## Stories built
- Establish the Go module and application boundaries for csvstat. (architecture) [story]

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- implements: go.md (SP 4743)
- compass: COMPASS.md (SP 628)
- implements: ARCHITECTURE.md (SP 625)
- stack: common.md (SP 1812)

## Build directory changes
- cmd/csvstat/main.go
- go.mod
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvstat/csvstat.go

## Pre-build acceptance observation
- RED: architecture-module (ARCHITECTURE.md)
  intent: The application exposes the required Go module and executable source boundary.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      root = PosixPath('.')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "architecture-module.py", line 5, in <module>
        assert (root / "go.mod").is_file()
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: architecture-module (ARCHITECTURE.md)
  intent: The application exposes the required Go module and executable source boundary.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
<reusable-compact filename="ARCHITECTURE.md">
- Go 1.22+ standard-library-only CLI module.
- Executable boundary: `./cmd/csvstat`; only `main` may call `os.Exit`.
- `internal/cli` owns argument parsing, diagnostics, dispatch, and exit codes.
- `internal/csvstat` owns CSV profiling and report serialization.
- Input comes only from the user-supplied path; reports use stdout, diagnostics stderr.
</reusable-compact>

RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/csvstat/main.go
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvstat/csvstat.go

SUMMARY:
Established the Go module and application package boundaries. `go build ./cmd/csvstat`, `go vet ./...`, formatting, and race-enabled tests pass.

BLOCKERS:
- None
