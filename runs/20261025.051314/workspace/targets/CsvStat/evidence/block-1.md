# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-25
- resulting state: closed/verified
- story points (combined assembled cost): 13517
- execution id: 20261025.053142.773Z-2fde1648

## Stories built
- Stand up the Go module, directory layout, and architecture for csvstat. (foundation) [story]

## Acceptance tooling authorization
- ARCHITECTURE.md#arch-go-mod: executable=go; scope=test; authorization=existing Target environment

## Reusable compacts
- ARCHITECTURE_compact.md

## Stacked context
- implements: go.md (SP 7578)
- compass: COMPASS.md (SP 963)
- implements: ARCHITECTURE.md (SP 2005)
- stack: common.md (SP 2776)

## Build directory changes
- bin/test.sh
- cmd/csvstat/main.go
- go.mod
- internal/cli/cli.go
- internal/cli/doc.go
- internal/csvdata/doc.go
- internal/freq/doc.go
- internal/profile/doc.go

## Pre-build acceptance observation
- RED: arch-go-mod (ARCHITECTURE.md)
  intent: The root go.mod declares a module, Go 1.22 or newer, and no third-party requirement.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      self = PosixPath('go.mod')
      mode = 'r'
      buffering = -1
      encoding = 'utf-8'
      errors = None
      newline = None
    --- drydock: end values ---
    Traceback (most recent call last):
      File "arch-go-mod.py", line 4, in <module>
        text = Path("go.mod").read_text(encoding="utf-8")
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1027, in read_text
        with self.open(mode='r', encoding=encoding, errors=errors) as f:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
      File "/home/barlo/.local/share/uv/python/cpython-3.12.13-linux-x86_64-gnu/lib/python3.12/pathlib.py", line 1013, in open
        return io.open(self, mode, buffering, encoding, errors, newline)
               ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'go.mod'
- RED: arch-layout (ARCHITECTURE.md)
  intent: The application packages exist under internal/ and the command directory exists under cmd/.
  return code: 1
  stderr:
    --- drydock: values at failure ---
      directory = PosixPath('internal/cli')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "arch-layout.py", line 5, in <module>
        assert directory.is_dir(), directory
               ^^^^^^^^^^^^^^^^^^
    AssertionError: internal/cli

## Post-build programmatic acceptance
- PASS: arch-go-mod (ARCHITECTURE.md)
  intent: The root go.mod declares a module, Go 1.22 or newer, and no third-party requirement.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: arch-layout (ARCHITECTURE.md)
  intent: The application packages exist under internal/ and the command directory exists under cmd/.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
Go 1.24.6 meets the 1.22 floor and the build directory is clean apart from `sources/`. Creating the module, layout, entry point, and test gate now.Build, vet, and the stub run all work. I'll tidy the unused-parameter handling, then run the two declared acceptance checks from the build directory.Foundation is in place and both declared checks pass. Writing the summary.

<reusable-compact filename="ARCHITECTURE.md">
Module: `csvstat`, `go 1.22`, no `require`.
Layout: `cmd/csvstat` (package main), `internal/cli`, `internal/csvdata`, `internal/profile`, `internal/freq`, `bin/test.sh` (vet then test).
Entry: `cli.Run(args []string, stdout, stderr io.Writer) int`; `args` excludes program name; only `cmd/csvstat/main.go` calls `os.Exit`.
Exit codes: 0 success; 1 operational (open/read failure); 2 usage; 3 rejected input.
Stdout written only on success; diagnostics to stderr only.
Rules: `sources/` read-only; no reading `sources/run.sh` or data files in app code; no embedded answers; stdlib only.
</reusable-compact>

<blueprint-decisions>
[{"spec":"ARCHITECTURE.md","severity":"Low","subject":"Stub behavior of cli.Run","decision":"The foundation has no argument parsing yet, so cli.Run writes 'csvstat: not implemented' to stderr and returns 1. I chose 1 over 0 so the stub cannot pass as a successful run. Replace it when the CLI story lands. Is that acceptable?"}]
</blueprint-decisions>

RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/csvstat/main.go
- internal/cli/doc.go
- internal/cli/cli.go
- internal/csvdata/doc.go
- internal/profile/doc.go
- internal/freq/doc.go
- bin/test.sh

SUMMARY:
- Created the root Go module `csvstat` (`go 1.22`, no `require`). The toolchain is go1.24.6, which meets the 1.22 floor.
- Created the layout: `cmd/csvstat` (package main, calls `cli.Run` and passes the result to `os.Exit`), `internal/cli`, `internal/csvdata`, `internal/profile`, `internal/freq` (each with a `doc.go`), and `bin/`.
- `internal/cli` declares `Run(args []string, stdout, stderr io.Writer) int` as a stub. It writes a diagnostic to stderr and returns 1.
- `bin/test.sh` runs `go vet ./...` then `go test ./...`. It exits 0.
- Checks run: `gofmt -l .` is clean. `sh bin/test.sh` exits 0 (no test files yet, so no tests ran). `go build ./cmd/csvstat` succeeds. The two declared acceptance checks, `arch-go-mod` and `arch-layout`, both pass when run from the build directory.
- Did not run `sources/run.sh`, since this is not the terminal story. Did not read it or any data file.
- Coverage: no tests yet. The foundation has no behavior beyond the stub.

BLOCKERS:
- The build instructions ask for the package ownership, entry contract, and exit-code table to be documented in `ARCHITECTURE.md`. That file lives under the Target's protected `blueprint/` directory, outside the authorized build directory, so I did not write it. The contract is captured in the reusable compact above. Someone with access to the Target should add the documentation there, or the build guardrail should be widened to allow a build-side copy.
