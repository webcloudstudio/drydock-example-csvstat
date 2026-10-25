# CSVSTAT

- Short description: Drydock uat - CsvStat builds a Go CSV column profiler. The source is an original brief: infer each column's type (integer, number, date, boolean, text, empty), report rows, nulls, distinct, min, max, sum, mean, and median, and list a column's most frequent values. The supplied grader, sources/run.sh, builds ./cmd/csvstat with go build and runs 22 checks of exit code and output against five seeded random CSV files with known answers: quoted commas, null spellings, padded text, an invalid calendar date, a header-only file, and a ragged file. Standard library only, no database. Run with drydock uat CsvStat. Source Repo: none, the brief is original to this kit.
- Status: BLOCKED: UAT STAGE

1 of 4 receipt claims proven. Open `index.html` for the linked proof kit; verify it with `sha256sum -c SHA256SUMS`.

## Receipt

| Claim | Verdict | Recorded outcome | Proof |
|---|---|---|---|
| Lifecycle completed | FAIL | 13 lifecycle commands executed; the run ended at 13-after-initial-build-build-status. | [result.json](result.json) |
| External conformance suite passed | UNPROVEN | The configured external conformance suite was not executed. | — |
| Release score passed | UNPROVEN | No release score was recorded for this run. | — |
| Integrity verification passed | PASS | 495 files digested; verify with sha256sum -c SHA256SUMS. | [SHA256SUMS](SHA256SUMS) |

## Run facts

- Drydock: `0.8.0` (commit `a05037a0923cb82e4a9ead403dbccac1a17a46af`)
- Provider and model: `claude` / `claude-haiku-5-5, claude-sonnet-5-5`
- Platform: `Linux-6.6.87.2-microsoft-standard-WSL2-x86_64-with-glibc2.35` on Python `3.12.13`
- Run: `20261025.001433`
- Elapsed: 1503.5s
- LLM calls: 25
- Tokens: cached 2,779,365; uncached 108; output 204,718
- LLM elapsed: 1127.2s
- Build passes: 1; repairs: 6 attempts allowed
- Conformance: not executed
- Advisory scores: none recorded

## RUN SUMMARY

- Input specification: [`sources/INSTRUCTIONS.md`](sources/INSTRUCTIONS.md)
- Delivered Code: [`build/CsvStat/`](build/CsvStat)
- Test Results: not recorded for this run.

## Shortfall

- Degraded: initial-build-1 exited 1

## Commands

| # | Command | Exit | Elapsed | Output |
|---|---|---|---|---|
| 01-init | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock init CsvStat` | 0 | 2.6s | [stdout](evidence/commands/01-init.stdout.log) · [stderr](evidence/commands/01-init.stderr.log) |
| 02-import-sources | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock import CsvStat sources --format markdown` | 0 | 4.3s | [stdout](evidence/commands/02-import-sources.stdout.log) · [stderr](evidence/commands/02-import-sources.stderr.log) |
| 03-analyze | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock analyze CsvStat` | 0 | 187.2s | [stdout](evidence/commands/03-analyze.stdout.log) · [stderr](evidence/commands/03-analyze.stderr.log) · [llm](evidence/commands/03-analyze.llm.log) |
| 04-plan | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock plan CsvStat --override --plan-override-model sonnet` | 0 | 540.6s | [stdout](evidence/commands/04-plan.stdout.log) · [stderr](evidence/commands/04-plan.stderr.log) · [llm](evidence/commands/04-plan.llm.log) |
| 05-plan-verify | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock plan verify CsvStat --plan-override-model sonnet` | 1 | 41.9s | [stdout](evidence/commands/05-plan-verify.stdout.log) · [stderr](evidence/commands/05-plan-verify.stderr.log) |
| 06-plan-repair | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock plan repair CsvStat --plan-override-model sonnet` | 0 | 122.5s | [stdout](evidence/commands/06-plan-repair.stdout.log) · [stderr](evidence/commands/06-plan-repair.stderr.log) · [llm](evidence/commands/06-plan-repair.llm.log) |
| 07-plan-verify-after-repair | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock plan verify CsvStat --plan-override-model sonnet` | 0 | 40.2s | [stdout](evidence/commands/07-plan-verify-after-repair.stdout.log) · [stderr](evidence/commands/07-plan-verify-after-repair.stderr.log) |
| 08-after-plan-build-status | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock build status CsvStat` | 0 | 1.6s | [stdout](evidence/commands/08-after-plan-build-status.stdout.log) · [stderr](evidence/commands/08-after-plan-build-status.stderr.log) |
| 09-after-plan-target-status | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock status CsvStat` | 0 | 1.3s | [stdout](evidence/commands/09-after-plan-target-status.stdout.log) · [stderr](evidence/commands/09-after-plan-target-status.stderr.log) |
| 10-after-plan-workspace-status | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock status` | 0 | 0.8s | [stdout](evidence/commands/10-after-plan-workspace-status.stdout.log) · [stderr](evidence/commands/10-after-plan-workspace-status.stderr.log) |
| 11-initial-ready | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock status CsvStat --ready` | 0 | 1.8s | [stdout](evidence/commands/11-initial-ready.stdout.log) · [stderr](evidence/commands/11-initial-ready.stderr.log) |
| 12-initial-build-1 | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock build CsvStat --override --repair-attempts 6` | 1 | 549.0s | [stdout](evidence/commands/12-initial-build-1.stdout.log) · [stderr](evidence/commands/12-initial-build-1.stderr.log) · [llm](evidence/commands/12-initial-build-1.llm.log) |
| 13-after-initial-build-build-status | `/mnt/c/Users/barlo/projects/drydock/.venv/bin/python -m drydock build status CsvStat` | 1 | 0.9s | [stdout](evidence/commands/13-after-initial-build-build-status.stdout.log) · [stderr](evidence/commands/13-after-initial-build-build-status.stderr.log) |

## Evidence

- [`evidence/commands/`](evidence/commands) — stdout and stderr of every command
- [`evidence/prompts/`](evidence/prompts) — the assembled prompt for every LLM call
- [`evidence/provider_raw/`](evidence/provider_raw) — unmodified provider transcripts
- [`evidence/llm_logs/`](evidence/llm_logs) — call banners and token accounting
- [`result.json`](result.json) — the machine-readable record of this run
