# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-24
- resulting state: closed/verified
- story points (combined assembled cost): 10350
- execution id: 20261024.205632.520Z-5e6416b7

## Stories built
- Stand up the Go module, the cmd/csvstat entry point, and the internal/cli Run skeleton. (foundation) [story]

## Acceptance tooling authorization
- ARCHITECTURE.md#architecture-module-layout: executable=go; scope=test; authorization=existing Target environment
- ARCHITECTURE.md#architecture-executable-builds: executable=go; scope=test; authorization=existing Target environment

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- implements: go.md (SP 7578)
- compass: COMPASS.md (SP 963)
- implements: ARCHITECTURE.md (SP 1667)

## Build directory changes
- cmd/csvstat/main.go
- go.mod
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/freq/doc.go
- internal/output/doc.go
- internal/profile/doc.go
- internal/reader/doc.go
- internal/typing/doc.go

## Pre-build acceptance observation
- RED: architecture-module-layout (ARCHITECTURE.md)
  intent: The application root is a Go 1.22+ module with no requirements and the declared entry point and CLI package directory.
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
      File "architecture-module-layout.py", line 4, in <module>
        mod = Path("go.mod").read_text()
              ^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
        with self.open(mode='r', encoding=encoding, errors=errors) as f:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
        return io.open(self, mode, buffering, encoding, errors, newline)
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'go.mod'
- RED: architecture-executable-builds (ARCHITECTURE.md)
  intent: The cmd/csvstat main package builds and the built executable runs to a normal process exit.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      found = []
    --- drydock: end values ---
    Traceback (most recent call last):
      File "architecture-executable-builds.py", line 3, in <module>
        code, out, err = cli_runner()
                         ^^^^^^^^^^^^
      File "/tmp/drydock-acceptance-jtr7vna2/drydock_go.py", line 114, in cli_runner
        package = package or _main_package()
                             ^^^^^^^^^^^^^^^
      File "/tmp/drydock-acceptance-jtr7vna2/drydock_go.py", line 82, in _main_package
        raise ValueError(f'expected one main package under ./cmd, found {found or "none"}; pass package=')
    ValueError: expected one main package under ./cmd, found none; pass package=
- GREEN (prepassed): architecture-only-main-exits (ARCHITECTURE.md)
  intent: No non-test Go source outside cmd/ calls os.Exit.
  return code: 0

## Post-build programmatic acceptance
- PASS: architecture-module-layout (ARCHITECTURE.md)
  intent: The application root is a Go 1.22+ module with no requirements and the declared entry point and CLI package directory.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: architecture-executable-builds (ARCHITECTURE.md)
  intent: The cmd/csvstat main package builds and the built executable runs to a normal process exit.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
  stdout:
    csvstat: expected exactly one FILE argument
    usage: csvstat FILE
- PASS: architecture-only-main-exits (ARCHITECTURE.md)
  intent: No non-test Go source outside cmd/ calls os.Exit.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
Go 1.24.6 meets the 1.22 floor, and the working tree is clean with only `sources/` present. I'll create the module, the entry point, the `Run` skeleton, the empty internal package stubs, and tests.The `ErrHelp` branch I wrote returns the same code as the other error, so it adds nothing. I'll collapse it before going further.Next, the internal package stubs the build block asks for, so later stories add code without restructuring. Each one holds only a package doc comment because the Go toolchain needs a package clause to compile.Now I'll verify: formatting, vet, the full test suite, and a build of the executable. I'll also run the two architecture checks that apply to this step, the module-layout and only-main-exits assertions, by reading the files back.RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/csvstat/main.go
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/reader/doc.go
- internal/typing/doc.go
- internal/profile/doc.go
- internal/freq/doc.go
- internal/output/doc.go

SUMMARY:
The Go module `csvstat` is in place with `go 1.22`, no `require` lines, and no third-party modules. `cmd/csvstat/main.go` is a thin `main` that calls `cli.Run` and is the only place that calls `os.Exit`. `internal/cli.Run(args, stdout, stderr) int` uses a local `flag.FlagSet` with `flag.ContinueOnError`. It returns 2 for usage errors (zero or two-plus arguments, unknown flag, `-h`) and writes nothing to stdout. A single file argument returns 0. The internal `reader`, `typing`, `profile`, `freq`, and `output` packages exist as package-doc stubs for later stories.

Verification, run in the build directory with Go 1.24.6:
- `gofmt -l .` printed nothing.
- `go vet ./...` passed.
- `go test -count=1 ./...` passed: `internal/cli` is ok, and the other packages have no test files yet.
- `go build ./cmd/csvstat` succeeded. The binary exited 2 with no args and with two args, and 0 with one arg.
- Architecture checks, read back directly: the go.mod version is 1.22 with no `require`, `cmd/csvstat/main.go` and `internal/cli` exist, and no non-test Go file under `internal/` calls `os.Exit`.

Tests in `internal/cli/cli_test.go` are table-driven. They check return codes and that stdout stays empty. They don't compare stderr wording.

Choices I made:
- A valid single-file invocation returns 0 with no output. The profiling stories have not been built, so this is a placeholder, not a real result. Later stories should replace it.
- `-h` returns 2, to match the usage-error contract, not 0.

Not done, by scope:
- `bin/test.sh` appears in the architecture and COMPASS, but no story in this block covers it, so I did not create it.
- I did not run `sources/run.sh`. It is the terminal story's gate.

<blueprint-decisions>
[{"spec":"ARCHITECTURE.md","severity":"Low","subject":"Valid single-file invocation behavior in the Run skeleton","decision":"Options were: return 0 with no output, or return a non-success code with a 'not implemented' diagnostic. I chose 0 with no output because exit 1 means the file cannot be read and the contract has no code for unimplemented work. Profiling stories will replace this path. Is that acceptable, or should the skeleton fail until profiling lands?"}]
</blueprint-decisions>

<reusable-compact filename="ARCHITECTURE.md">
Module: `csvstat`, go 1.22, no requires. Entry: `cmd/csvstat` (package main) calls `os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))`.
Contract: `internal/cli.Run(args []string, stdout, stderr io.Writer) int`; flag.FlagSet with ContinueOnError; diagnostics to stderr only.
Exit codes: 0 success; 1 file cannot be opened/read; 2 usage error; 3 rejected input (empty file, empty or repeated header, ragged record, absent column).
Failing commands write nothing to stdout; output is buffered and written only on success.
Packages: internal/reader (Read, Trim, IsNull, rejection checks); internal/typing (Infer, IsDate); internal/profile (Counts, MinMax, NumericStats); internal/freq (Top); internal/output (WriteProfile, WriteFreq).
Only main calls os.Exit (non-test code under internal/ must not).
Stats: sum/mean/median via math/big exact values, rounded once half away from zero when written.
</reusable-compact>
