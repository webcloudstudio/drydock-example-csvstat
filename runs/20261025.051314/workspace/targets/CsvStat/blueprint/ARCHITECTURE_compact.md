<!-- Compacted from ARCHITECTURE.md sha256=e12f94e5c90e8e9b4581c3d4e02add14222bb76448d04c486ee6d71e88a43bf6 on 2026-10-25 by drydock build agent -->

Module: `csvstat`, `go 1.22`, no `require`.
Layout: `cmd/csvstat` (package main), `internal/cli`, `internal/csvdata`, `internal/profile`, `internal/freq`, `bin/test.sh` (vet then test).
Entry: `cli.Run(args []string, stdout, stderr io.Writer) int`; `args` excludes program name; only `cmd/csvstat/main.go` calls `os.Exit`.
Exit codes: 0 success; 1 operational (open/read failure); 2 usage; 3 rejected input.
Stdout written only on success; diagnostics to stderr only.
Rules: `sources/` read-only; no reading `sources/run.sh` or data files in app code; no embedded answers; stdlib only.
