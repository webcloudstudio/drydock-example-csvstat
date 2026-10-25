# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-24
- resulting state: closed/verified
- story points (combined assembled cost): 13780
- execution id: 20261025.003108.255Z-cd35dd12

## Stories built
- Stand up the Go module, package layout, and architecture for the csvstat CLI. (foundation) [story]

## Acceptance tooling authorization
- ARCHITECTURE.md#arch-module-builds: executable=go; scope=test; authorization=existing Target environment
- ARCHITECTURE.md#arch-stdlib-only: executable=go; scope=test; authorization=existing Target environment
- ARCHITECTURE.md#arch-vet-clean: executable=go; scope=test; authorization=existing Target environment

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- implements: go.md (SP 7578)
- compass: COMPASS.md (SP 963)
- implements: ARCHITECTURE.md (SP 2220)
- stack: common.md (SP 2776)

## Build directory changes
- .gitignore
- bin/test.sh
- cmd/csvstat/main.go
- go.mod
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvio/doc.go
- internal/decimal/doc.go
- internal/profile/doc.go

## Pre-build acceptance observation
- RED: arch-module-builds (ARCHITECTURE.md)
  intent: The application root is a Go module whose main package at ./cmd/csvstat builds with go build.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "arch-module-builds.py", line 6, in <module>
        assert Path("go.mod").is_file()
               ^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: arch-stdlib-only (ARCHITECTURE.md)
  intent: The module declares no third-party requirement and the main package depends only on this module and the standard library.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      self = PosixPath('go.mod')
      mode = 'r'
      buffering = -1
      encoding = 'locale'
      errors = None
      newline = None
    --- drydock: end values ---
    Traceback (most recent call last):
      File "arch-stdlib-only.py", line 5, in <module>
        go_mod = Path("go.mod").read_text()
                 ^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
        with self.open(mode='r', encoding=encoding, errors=errors) as f:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
        return io.open(self, mode, buffering, encoding, errors, newline)
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'go.mod'
- RED: arch-vet-clean (ARCHITECTURE.md)
  intent: The foundation tree passes go vet.
  return code: 1
  stderr:
    pattern ./...: directory prefix . does not contain main module or its selected dependencies
    
    --- drydock: values at failure ---
      result = CompletedProcess(args=['go', 'vet', './...'], returncode=1, stdout='', stderr='pattern ./...: directory prefix . does not contain main module or its selected dependencies\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "arch-vet-clean.py", line 7, in <module>
        assert result.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- RED: arch-only-main-exits (ARCHITECTURE.md)
  intent: Only the main package calls os.Exit and the internal packages exist as declared.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "arch-only-main-exits.py", line 3, in <module>
        assert Path("internal/cli").is_dir()
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: arch-module-builds (ARCHITECTURE.md)
  intent: The application root is a Go module whose main package at ./cmd/csvstat builds with go build.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: arch-stdlib-only (ARCHITECTURE.md)
  intent: The module declares no third-party requirement and the main package depends only on this module and the standard library.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat/internal/cli
    csvstat/cmd/csvstat
- PASS: arch-vet-clean (ARCHITECTURE.md)
  intent: The foundation tree passes go vet.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: arch-only-main-exits (ARCHITECTURE.md)
  intent: Only the main package calls os.Exit and the internal packages exist as declared.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
The build directory has only `sources/` and an empty git repo. I'll create the module, the entry point, the `internal/cli` package with tests, the package stubs from the architecture, and `bin/test.sh`.Files are written. Now I'll run gofmt, the test script, and the four declared acceptance checks.The foundation builds, `gofmt` is clean, and `bin/test.sh` exits 0. I'll also emit the optional compact for ARCHITECTURE.md.

RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/csvstat/main.go
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvio/doc.go
- internal/profile/doc.go
- internal/decimal/doc.go
- bin/test.sh
- .gitignore

<reusable-compact filename="ARCHITECTURE.md">
Module: `csvstat`, go 1.22, no require lines (stdlib only).
Entry: `cmd/csvstat/main.go` — `os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))`; sole `os.Exit` caller.
`internal/cli.Run(args []string, stdout, stderr io.Writer) int` — exit codes: 0 success, 1 file open/read failure, 2 usage error, 3 rejected input (empty file, empty/repeated header, ragged record, CSV syntax error, unknown column).
Failing commands write nothing to stdout; one diagnostic line to stderr; output buffered until validation completes.
Packages: `internal/csvio` (open, read, trim, validate), `internal/profile` (null rules, type inference, counts, stats, frequencies), `internal/decimal` (math/big exact arithmetic, two-decimal formatting, half-away-from-zero rounding once at output).
Test entry: `sh bin/test.sh` = `go vet ./...` then `go test ./...`.
</reusable-compact>

SUMMARY:
- Created the Go module `csvstat` (go 1.22, no requires) and the `cmd/csvstat` main package, which only calls `os.Exit` on the value returned by `cli.Run`.
- `internal/cli.Run` parses flags with a local `flag.FlagSet` (ContinueOnError), returns 2 for usage errors (no file, more than one file, unknown flag), 0 for `-h` and for one file. The one-file path is a placeholder: no profiling happens yet, because csvio/profile/decimal are later stories.
- Added doc-only packages `internal/csvio`, `internal/profile`, `internal/decimal` so the package layout from ARCHITECTURE.md exists.
- Added `bin/test.sh` (POSIX sh, `go vet ./...` then `go test ./...`). I left out the tee logging from the stack template, since `set -e` with a pipe would lose the exit status under POSIX sh.
- Added `.gitignore` for `data/`, `*.log`, and the root `/csvstat` binary.
- Tests in `internal/cli/cli_test.go` assert exit codes and that stdout is empty on usage errors. They do not check stderr wording.
- Verification: `gofmt -l .` prints nothing; `sh bin/test.sh` exits 0; `go build ./cmd/csvstat` succeeds; `go list -deps` lists only module packages; the binary exits 2 on usage errors. The four ARCHITECTURE.md checks (module builds, stdlib only, vet clean, only main exits) all hold by these same checks and by reading the code. I did not run the Python-embedded versions of those checks.
- I did not run `sources/run.sh`. It is the terminal story's gate, and this block does not invoke it. The staged sources are untouched.
- I did not modify ARCHITECTURE.md, since it sits under the protected Target path. It already documents the package ownership the instructions asked for.
- No git add or commit, per the build rules.
