<!-- Compacted from ARCHITECTURE.md sha256=523486215aa0fe9a99cda0460e112edd7b537c3b37017825b01e6352c1e713c6 on 2026-10-24 by drydock build agent -->

Module: `csvstat`, go 1.22, no requires. Entry: `cmd/csvstat` (package main) calls `os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))`.
Contract: `internal/cli.Run(args []string, stdout, stderr io.Writer) int`; flag.FlagSet with ContinueOnError; diagnostics to stderr only.
Exit codes: 0 success; 1 file cannot be opened/read; 2 usage error; 3 rejected input (empty file, empty or repeated header, ragged record, absent column).
Failing commands write nothing to stdout; output is buffered and written only on success.
Packages: internal/reader (Read, Trim, IsNull, rejection checks); internal/typing (Infer, IsDate); internal/profile (Counts, MinMax, NumericStats); internal/freq (Top); internal/output (WriteProfile, WriteFreq).
Only main calls os.Exit (non-test code under internal/ must not).
Stats: sum/mean/median via math/big exact values, rounded once half away from zero when written.
