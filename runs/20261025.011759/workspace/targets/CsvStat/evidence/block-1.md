# Evidence: Block 1 · Foundational (block-1)

- block type: block
- date: 2026-10-24
- resulting state: closed/verified
- story points (combined assembled cost): 13440
- execution id: 20261025.012754.773Z-86ba1825

## Stories built
- Stand up the Go module, the csvstat entry point, and the testable CLI package. (foundation) [story]

## Acceptance tooling authorization
- ARCHITECTURE.md#arch-module-layout: executable=go; scope=test; authorization=existing Target environment
- ARCHITECTURE.md#arch-cli-run-compiles: executable=go; scope=test; authorization=existing Target environment

## Stacked context
- implements: go.md (SP 7578)
- compass: COMPASS.md (SP 963)
- implements: ARCHITECTURE.md (SP 1959)
- stack: common.md (SP 2776)

## Build directory changes
- bin/test.sh
- cmd/csvstat/main.go
- go.mod
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvio/csvio.go
- internal/output/output.go
- internal/profile/profile.go

## Pre-build acceptance observation
- RED: arch-module-layout (ARCHITECTURE.md)
  intent: The application root is a Go module of Go 1.22 or newer with the csvstat main package and the cli package present.
  return code: 1
  stderr:
    Traceback (most recent call last):
      File "arch-module-layout.py", line 4, in <module>
        assert os.path.isfile("go.mod")
               ^^^^^^^^^^^^^^^^^^^^^^^^
    AssertionError
- UNVERIFIED: arch-no-dependencies (ARCHITECTURE.md)
  intent: go.mod declares no third-party requirement, so the standard library is the only dependency.
  return code: 1
  error: unverified acceptance: the assertion raised FileNotFoundError ([Errno 2] No such file or directory: 'go.mod') in its own frame, before reaching the code under test. This is a fault in the acceptance kit or its environment, not a defect in the build.
  stderr:
    Traceback (most recent call last):
      File "arch-no-dependencies.py", line 3, in <module>
        with open("go.mod", encoding="utf-8") as handle:
             ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
    FileNotFoundError: [Errno 2] No such file or directory: 'go.mod'
- RED: arch-cli-run-compiles (ARCHITECTURE.md)
  intent: The module builds and vets cleanly, so cli.Run and the main package compile together.
  return code: 1
  stderr:
    pattern ./...: directory prefix . does not contain main module or its selected dependencies
    
    --- drydock: values at failure ---
      build = CompletedProcess(args=['go', 'build', './...'], returncode=1, stdout='', stderr='pattern ./...: directory prefix . does not contain main module or its selected dependencies\n')
    --- drydock: end values ---
    Traceback (most recent call last):
      File "arch-cli-run-compiles.py", line 7, in <module>
        assert build.returncode == 0
               ^^^^^^^^^^^^^^^^^^^^^
    AssertionError

## Post-build programmatic acceptance
- PASS: arch-module-layout (ARCHITECTURE.md)
  intent: The application root is a Go module of Go 1.22 or newer with the csvstat main package and the cli package present.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: arch-no-dependencies (ARCHITECTURE.md)
  intent: go.mod declares no third-party requirement, so the standard library is the only dependency.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0
- PASS: arch-cli-run-compiles (ARCHITECTURE.md)
  intent: The module builds and vets cleanly, so cli.Run and the main package compile together.
  target interpreter: /mnt/c/Users/barlo/projects/drydock/.venv/bin/python
  provisioning: not required
  return code: 0

## Build summary
Files are written. Now I'll run the declared acceptance checks: the build, vet, the test script, and the module assertions.RESULT: SUCCESS

FILES CHANGED:
- go.mod
- cmd/csvstat/main.go
- internal/cli/cli.go
- internal/cli/cli_test.go
- internal/csvio/csvio.go
- internal/profile/profile.go
- internal/output/output.go
- bin/test.sh

SUMMARY:
- **Module:** `go.mod` declares `module csvstat` with `go 1.22`, and has no `require` lines. The toolchain is go1.24.6, which meets the 1.22 floor.
- **Entry point:** `cmd/csvstat/main.go` is the only caller of `os.Exit`. It passes `os.Args[1:]`, stdout and stderr to `cli.Run`.
- **CLI package:** `internal/cli/cli.go` exposes `Run(args []string, stdout, stderr io.Writer) int`. It uses a local `flag.FlagSet` with `ContinueOnError`. A bad flag returns 2, `-h` returns 0, and anything else returns 0. This is a stub; later stories fill it in.
- **Package layout:** `internal/csvio`, `internal/profile`, and `internal/output` exist as package-only stubs with doc comments, so the layout in ARCHITECTURE.md is in place.
- **Tests:** `internal/cli/cli_test.go` covers the unknown-flag exit code 2 with empty stdout, `-h` exit 0, and no-args exit 0. Each test checks the return value and stdout, not message wording.
- **Test runner:** `bin/test.sh` is POSIX sh. It checks gofmt, runs `go vet ./...`, then `go test -race -count=1 ./...`. I did not add the tee logging from the common stack guidance, because the script has to return a plain exit status and `pipefail` is not POSIX.

Checks run:
- `sh bin/test.sh`: exit 0. The `internal/cli` tests pass under `-race`.
- `go build ./...` and `go vet ./...`: both exit 0.
- `gofmt -l .`: no output.
- Built binary: `-bogus` exits 2 and no args exits 0.
- The arch-module-layout, arch-no-dependencies, and arch-cli-run-compiles assertions are satisfied by the files above. I did not run the Python AC snippets themselves.
- `sources/` is unchanged. Nothing was staged or committed.

Not done:
- ARCHITECTURE.md's module ownership table and technology stack are already recorded in the protected Target `blueprint/` directory. I did not touch it, because the build guardrail protects it.
- The terminal grader `sources/run.sh` belongs to a later step, so I did not run it.

BLOCKERS:
- None.
