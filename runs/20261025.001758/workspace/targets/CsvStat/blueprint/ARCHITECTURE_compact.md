<!-- Compacted from ARCHITECTURE.md sha256=0f827391ef8e9c142d1bdfd2dd65ddbfda03f1b1d89fb050f5619b182327c83c on 2026-10-24 by drydock build agent -->

Module: `csvstat` (go 1.22, no requires, no go.sum).
Packages: cmd/csvstat (main only, os.Exit), internal/cli (Run(args []string, stdout, stderr io.Writer) int), internal/csvdata, internal/infer, internal/profile, internal/freq, internal/decimal. bin/test.sh = go vet ./... then go test ./....
Exit codes: 0 success; 1 operational (file cannot be opened/read); 2 usage (missing/extra FILE, unknown option, missing option value, --top not positive, --top without --freq, --column with --freq); 3 rejected input (empty file, empty/repeated header, ragged record, column absent from header).
Failure contract: nothing on stdout; exactly one stderr line naming the file (and line number for ragged records). Stderr is never graded.
Rules: stdlib only; no floating point for sum/mean/median; only cmd/csvstat calls os.Exit; usage validated before file open; app code never reads sources/ data.
