# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-23
- resulting state: closed/verified
- story points (combined assembled cost): 6226
- execution id: 20261023.223246.637Z-942037e8

## Stories built
- Define the CsvStat command architecture and module boundaries. (architecture) [story]

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- implements: go.md (SP 4944)
- compass: COMPASS.md (SP 628)
- implements: ARCHITECTURE.md (SP 584)

## Build directory changes
- bin/test.sh
- cmd/csvstat/main.go
- go.mod
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvstat/csvstat.go
- internal/csvstat/csvstat_test.go

## Pre-build acceptance observation
- RED: architecture-build (ARCHITECTURE.md)
  intent: The application architecture builds as the declared Go command.
  return code: 1
  stdout:
    go: cannot find main module, but found .git/config in /mnt/c/Users/barlo/projects/drydock/uat/CsvStat
    	to create a module there, run:
    	cd ../../../.. && go mod init
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'build', './cmd/csvstat'], returncode=1, stdout='', stderr='go: cannot find main module, but found .git/config in /mnt/c/Users/barlo/projects/drydock/uat/CsvStat\n\tto create a module there, run:\n\tcd ../../../.. && go mod init\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "architecture-build.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: architecture-vet (ARCHITECTURE.md)
  intent: The declared Go packages pass static analysis.
  return code: 1
  stdout:
    pattern ./...: directory prefix . does not contain main module or its selected dependencies
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'vet', './...'], returncode=1, stdout='', stderr='pattern ./...: directory prefix . does not contain main module or its selected dependencies\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "architecture-vet.py", line 10, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: architecture-build (ARCHITECTURE.md)
  intent: The application architecture builds as the declared Go command.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: architecture-vet (ARCHITECTURE.md)
  intent: The declared Go packages pass static analysis.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
<reusable-compact filename="ARCHITECTURE.md">
- Go CLI architecture: `cmd/csvstat` handles process startup and `os.Exit`; `internal/cli` handles flags, dispatch, diagnostics, and exit codes; `internal/csvstat` owns CSV loading, normalization, profiling, frequencies, and rendering; `bin/test.sh` runs verification.
- Standard library only; Go 1.22+.
- Success writes CSV to stdout; diagnostics use stderr.
- Exit codes: 1 operational failure, 2 usage failure, 3 rejected input.
</reusable-compact>

RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/csvstat/main.go
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvstat/csvstat.go
- internal/csvstat/csvstat_test.go
- bin/test.sh

SUMMARY:
Established the Go module, command boundary, internal package boundaries, CSV loading foundation, tests, and verification launcher. `go build ./cmd/csvstat`, `go vet ./...`, and `./bin/test.sh` pass.

BLOCKERS:
- None
