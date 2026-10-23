# Soundings

Per-assertion acceptance board, one row per Blueprint Programmatic Acceptance check.
`drydock plan` projects every assertion as `— UNVERIFIED`; `drydock score ac` sets the
Status column from deterministic proof results. A rerun of `drydock plan` resets Status to
`— UNVERIFIED`; rescore to refresh.

| Status | Blueprint | AC Id | Text | Binding | Baseline | Previously Passed | Last Block | Evidence | Verified At |
|---|---|---|---|---|---|---|---|---|---|
| ✓ PASS | ARCHITECTURE.md | architecture-build | The application architecture builds as the declared Go command. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | ARCHITECTURE.md | architecture-vet | The declared Go packages pass static analysis. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ~ PREPASSED | FEATURE-CLI-FOUNDATION.md | cli-foundation-build | The declared command package builds successfully. | yes | VALID existing-behavior | yes | block-3 | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-10-23T22:38:18+00:00 |
| ~ PREPASSED | FEATURE-CLI-FOUNDATION.md | cli-foundation-help | The executable provides a successful help invocation. | yes | VALID existing-behavior | yes | block-3 | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-10-23T22:38:18+00:00 |
| ~ PREPASSED | FEATURE-CLI-FOUNDATION.md | cli-foundation-usage | A missing file argument is reported as a usage failure. | yes | VALID existing-behavior | yes | block-3 | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-CLI-INPUT.md | cli-input-quoted | A supplied CSV containing quoted commas is accepted. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-CLI-INPUT.md | cli-input-ragged | A ragged CSV is rejected without standard output. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-CLI-INPUT.md | cli-input-header-only | A header-only CSV is accepted as a valid empty data set. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-CLI-COLUMNS.md | cli-columns-existing | A named existing column is accepted for profiling. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-CLI-COLUMNS.md | cli-columns-missing-profile | A missing profile column is rejected without standard output. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-CLI-COLUMNS.md | cli-columns-missing-frequency | A missing frequency column is rejected without standard output. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-PROFILE-TYPES.md | profile-types-fixtures | The supplied profile fixtures are accepted and produce profile rows. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-PROFILE-TYPES.md | profile-types-all-null | An all-null column remains a valid profile input. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-PROFILE-TYPES.md | profile-types-invalid-date-text | A column containing an invalid calendar date remains valid input rather than causing command failure. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-FREQUENCY-VALUES.md | frequency-excludes-nulls | Frequency output omits every recognized null spelling. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-FREQUENCY-VALUES.md | frequency-orders-counts | Frequency rows are ordered by descending count before value ordering. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-FREQUENCY-VALUES.md | frequency-orders-ties | Equal-frequency values are ordered by ascending byte value. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-FREQUENCY-LIMITS.md | frequency-limit-success | A valid positive limit executes the frequency command successfully and bounds the parsed result. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-FREQUENCY-LIMITS.md | frequency-limit-default | Omitting --top uses the default limit and returns a successful frequency result. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-FREQUENCY-LIMITS.md | frequency-limit-invalid | Non-positive and malformed limits are rejected as usage errors. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-FREQUENCY-LIMITS.md | top-requires-frequency | --top without --freq is rejected as a usage error. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-PROFILE-STATISTICS.md | profile-statistics-integer | Integer statistics include counts, extrema, exact total, mean, and median. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-PROFILE-STATISTICS.md | profile-statistics-decimal | Decimal statistics preserve exact values and round mean and median to two places. | no | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-PROFILE-STATISTICS.md | profile-statistics-nonnumeric | Non-numeric columns retain counts and extrema while leaving numeric fields empty. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-PROFILE-OUTPUT.md | profile-output-all-columns | A profile command emits one output row for each input column in file order. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-PROFILE-OUTPUT.md | profile-output-selected-column | The column option emits exactly the requested profile row. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-PROFILE-OUTPUT.md | profile-output-quoting | Profile output preserves comma-containing values through CSV parsing. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ~ PREPASSED | FEATURE-CLI-ERRORS.md | operational-error | An unreadable input file returns the operational failure status and no CSV output. | yes | VALID existing-behavior | yes | block-3 | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-10-23T22:38:18+00:00 |
| ~ PREPASSED | FEATURE-CLI-ERRORS.md | usage-errors | Malformed command arguments return usage status and no CSV output. | yes | VALID existing-behavior | yes | block-3 | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-CLI-ERRORS.md | rejected-input | Structurally invalid input and absent columns return rejected-input status and no CSV output. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-CLI-ERRORS.md | successful-output-isolated | A valid command remains successful and emits CSV on standard output. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
| ~ PREPASSED | FEATURE-TEST-GATE.md | test-gate-present | The project provides an executable POSIX test launcher. | yes | VALID existing-behavior | yes | block-3 | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-10-23T22:38:18+00:00 |
| ~ PREPASSED | FEATURE-TEST-GATE.md | test-gate-runs | The project test launcher completes successfully. | yes | VALID existing-behavior | yes | block-3 | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-10-23T22:38:18+00:00 |
| ~ PREPASSED | FEATURE-TEST-GATE.md | go-vet-and-tests | The launcher exercises both the Go vet and Go test commands. | yes | VALID existing-behavior | yes | block-3 | green at this block's baseline too, before its code existed — confirm the criterion exercises the story's work | 2026-10-23T22:38:18+00:00 |
| ✓ PASS | FEATURE-FULL-VERIFICATION.md | full-conformance | The supplied conformance suite passes for the completed CsvStat implementation. | yes | VALID baseline-red | yes | block-3 |  | 2026-10-23T22:38:18+00:00 |
