<!-- Compacted from ARCHITECTURE.md sha256=0523e9a1384a42626d0062f018639c66c4784890cecd223013d13f942f05e05a on 2026-10-23 by drydock build agent -->

- Go CLI architecture: `cmd/csvstat` handles process startup and `os.Exit`; `internal/cli` handles flags, dispatch, diagnostics, and exit codes; `internal/csvstat` owns CSV loading, normalization, profiling, frequencies, and rendering; `bin/test.sh` runs verification.
- Standard library only; Go 1.22+.
- Success writes CSV to stdout; diagnostics use stderr.
- Exit codes: 1 operational failure, 2 usage failure, 3 rejected input.
