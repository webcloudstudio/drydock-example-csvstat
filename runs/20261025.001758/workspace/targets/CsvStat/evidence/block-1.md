# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-24
- resulting state: closed/verified
- story points (combined assembled cost): 13881
- execution id: 20261025.003040.578Z-7fa8824b

## Stories built
- Stand up the Go module, package layout, and the CLI entry contract everything else builds into. (foundation) [story]

## Acceptance tooling authorization
- ARCHITECTURE.md#arch-go-mod-stdlib-only: executable=go; scope=test; authorization=existing Target environment
- ARCHITECTURE.md#arch-cli-run-compiles: executable=go; scope=test; authorization=existing Target environment

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- implements: go.md (SP 7578)
- compass: COMPASS.md (SP 963)
- implements: ARCHITECTURE.md (SP 2320)
- stack: common.md (SP 2776)

## Build directory changes
- bin/test.sh
- cmd/csvstat/main.go
- go.mod
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvdata/doc.go
- internal/decimal/doc.go
- internal/freq/doc.go
- internal/infer/doc.go
- internal/profile/doc.go

## Pre-build acceptance observation
- RED: arch-go-mod-stdlib-only (ARCHITECTURE.md)
  intent: The application root is a Go module that declares no third-party requirements.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      gomod = PosixPath('go.mod')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "arch-go-mod-stdlib-only.py", line 4, in <module>
        assert gomod.is_file()
               ^^^^^^^^^^^^^^^
    AssertionError
- RED: arch-package-layout (ARCHITECTURE.md)
  intent: Each internal package named by the architecture exists and holds Go source.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      name = 'cli'
      directory = PosixPath('internal/cli')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "arch-package-layout.py", line 6, in <module>
        assert directory.is_dir(), name
               ^^^^^^^^^^^^^^^^^^
    AssertionError: cli
- RED: arch-cli-run-compiles (ARCHITECTURE.md)
  intent: The internal packages, including the cli.Run entry contract, compile and vet cleanly.
  return code: 1
  stdout:
    pattern ./internal/...: directory prefix internal does not contain main module or its selected dependencies
  stderr:
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'build', './internal/...'], returncode=1, stdout='', stderr='pattern ./internal/...: directory prefix internal does not contain main module or its selected dependencies\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "arch-cli-run-compiles.py", line 6, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: arch-run-signature (ARCHITECTURE.md)
  intent: internal/cli exposes Run taking args and two writers and returning an int exit code.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      text = ''
    --- drydock: end values ---
    Traceback (most recent call last):
      File "arch-run-signature.py", line 8, in <module>
        assert "func Run(" in text
               ^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: arch-go-mod-stdlib-only (ARCHITECTURE.md)
  intent: The application root is a Go module that declares no third-party requirements.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: arch-package-layout (ARCHITECTURE.md)
  intent: Each internal package named by the architecture exists and holds Go source.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: arch-cli-run-compiles (ARCHITECTURE.md)
  intent: The internal packages, including the cli.Run entry contract, compile and vet cleanly.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: arch-run-signature (ARCHITECTURE.md)
  intent: internal/cli exposes Run taking args and two writers and returning an int exit code.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
The build directory is empty apart from `sources/` and Go 1.24.6 is installed, which meets the 1.22 floor. I'll create the module, the package layout, the CLI entry contract with usage validation, a launcher, and tests.Files are written. Next I'll make the launcher executable, then run the gate, the formatting check, and the build that the architecture's acceptance checks depend on.Gate, build, and smoke checks pass. I'll run the four declared architecture acceptance checks directly, since they are this step's Definition of Done.All four declared checks pass. I'll write the reusable compact and close with the result.

<reusable-compact filename="ARCHITECTURE.md">
Module: `csvstat` (go 1.22, no requires, no go.sum).
Packages: cmd/csvstat (main only, os.Exit), internal/cli (Run(args []string, stdout, stderr io.Writer) int), internal/csvdata, internal/infer, internal/profile, internal/freq, internal/decimal. bin/test.sh = go vet ./... then go test ./....
Exit codes: 0 success; 1 operational (file cannot be opened/read); 2 usage (missing/extra FILE, unknown option, missing option value, --top not positive, --top without --freq, --column with --freq); 3 rejected input (empty file, empty/repeated header, ragged record, column absent from header).
Failure contract: nothing on stdout; exactly one stderr line naming the file (and line number for ragged records). Stderr is never graded.
Rules: stdlib only; no floating point for sum/mean/median; only cmd/csvstat calls os.Exit; usage validated before file open; app code never reads sources/ data.
</reusable-compact>

<blueprint-decisions>
[{"spec":"ARCHITECTURE.md","severity":"Low","subject":"Option placement and help handling","decision":"Options are parsed with the standard library flag package, so they must precede FILE (for example `csvstat --freq --top 3 f.csv`). Options after the file are rejected as an extra argument with exit 2. -h prints usage to stdout and exits 0, which is outside the 0/1/2/3 table. Is that acceptable, or should interleaved options or help exit 2 be adopted on replan?"},{"spec":"ARCHITECTURE.md","severity":"Low","subject":"Profiling stub in the foundation step","decision":"Valid usage reaches a stub in internal/cli that writes one stderr line and returns exit 1, without opening the file. This is a placeholder for the profiling stories, not a real operational failure. Later blocks replace it."}]
</blue-prints-decisions>

RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/csvstat/main.go
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvdata/doc.go
- internal/infer/doc.go
- internal/profile/doc.go
- internal/freq/doc.go
- internal/decimal/doc.go
- bin/test.sh

SUMMARY:
- Module `csvstat` on go 1.22 with no requires. No go.sum is produced. Toolchain is go1.24.6, which meets the 1.22 floor.
- `cmd/csvstat/main.go` is the only caller of `os.Exit`. It delegates to `cli.Run(args, stdout, stderr) int`.
- `internal/cli` defines the exit codes 0/1/2/3 and parses options with a local `flag.FlagSet` (ContinueOnError, output discarded). Usage rules are enforced before any file access: exactly one FILE, `--top` requires `--freq` and must be positive, `--column` conflicts with `--freq`. Usage errors write one stderr line and exit 2, with nothing on stdout.
- Packages `csvdata`, `infer`, `profile`, `freq`, and `decimal` exist as documented placeholders for later stories.
- `bin/test.sh` is a POSIX launcher that runs `go vet ./...` then `go test ./...`. It is executable.
- Tests: `internal/cli/cli_test.go` covers each usage error (exit 2, empty stdout, one stderr line), help (exit 0), valid usage not returning 2, and option capture. Every test checks the returned code and the state of the output writers.
- Verification: `sh bin/test.sh` exits 0 (cli tests pass). `gofmt -l .` prints nothing. `go build ./internal/...` and `go build ./cmd/csvstat` succeed. Smoke checks return exit 2 for no arguments and for `--top 0`. The four declared architecture checks (go-mod-stdlib-only, package-layout, cli-run-compiles, run-signature) all pass when run directly.
- I did not run `sources/run.sh`. The terminal story runs it once at the end, and profiling is not implemented yet.
- Choices for you to review are in the blueprint-decisions payload: options must precede FILE, `-h` exits 0, and valid usage reaches a stub that exits 1.

BLOCKERS:
- None.
