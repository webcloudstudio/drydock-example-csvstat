<!-- Compacted from ARCHITECTURE.md sha256=57a234b80547446b389d3155cd03f148f68875c3d778548285f38cedcb095b0a on 2026-10-24 by drydock build agent -->

Module: `csvstat`, go 1.22, no require lines (stdlib only).
Entry: `cmd/csvstat/main.go` — `os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))`; sole `os.Exit` caller.
`internal/cli.Run(args []string, stdout, stderr io.Writer) int` — exit codes: 0 success, 1 file open/read failure, 2 usage error, 3 rejected input (empty file, empty/repeated header, ragged record, CSV syntax error, unknown column).
Failing commands write nothing to stdout; one diagnostic line to stderr; output buffered until validation completes.
Packages: `internal/csvio` (open, read, trim, validate), `internal/profile` (null rules, type inference, counts, stats, frequencies), `internal/decimal` (math/big exact arithmetic, two-decimal formatting, half-away-from-zero rounding once at output).
Test entry: `sh bin/test.sh` = `go vet ./...` then `go test ./...`.
