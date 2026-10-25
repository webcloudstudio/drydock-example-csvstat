# AUTHORITATIVE PROJECT METADATA — FIELDS SHOULD BE CURRENT

name: CsvStat
display_name: drydock CsvStat Uat Case
short_description: Drydock uat - CsvStat builds a Go CSV column profiler. The source is an original brief: infer each column's type (integer, number, date, boolean, text, empty), report rows, nulls, distinct, min, max, sum, mean, and median, and list a column's most frequent values. The supplied grader, sources/run.sh, builds ./cmd/csvstat with go build and runs 22 checks of exit code and output against five seeded random CSV files with known answers: quoted commas, null spellings, padded text, an invalid calendar date, a header-only file, and a ragged file. Standard library only, no database. Run with drydock uat CsvStat. Source Repo: none, the brief is original to this kit.
stack: Go
version: 0.12
build_state: built
build_sub_state: complete
last_analyzed: 2026-10-25
last_planned: 2026-10-25
last_built: 2026-10-25
build_dir: 
brand: 
git_repo: 
release_tag: 
updated: 2026-10-25T06:03:58+00:00
