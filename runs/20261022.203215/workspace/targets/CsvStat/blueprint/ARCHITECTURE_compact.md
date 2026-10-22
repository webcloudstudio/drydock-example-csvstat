<!-- Compacted from ARCHITECTURE.md sha256=51d9d1b5870f5f3f788f78d0d7e36e56c24611a4802134386be8de40d49ba705 on 2026-10-22 by drydock build agent -->

- Go 1.22+ standard-library-only CLI module.
- Executable boundary: `./cmd/csvstat`; only `main` may call `os.Exit`.
- `internal/cli` owns argument parsing, diagnostics, dispatch, and exit codes.
- `internal/csvstat` owns CSV profiling and report serialization.
- Input comes only from the user-supplied path; reports use stdout, diagnostics stderr.
