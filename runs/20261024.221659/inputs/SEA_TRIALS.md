# Sea Trials: CsvStat

## Policy

| Consequence | On FAIL | On INCONCLUSIVE |
|---|---|---|
| blocks  | fail   | attest |
| scores  | score  | score  |
| attests | report | report |

## st-001: The supplied grader passes
Type: technical
Required: yes
Criterion: The completed csvstat shall make sh sources/run.sh exit zero; that script's exit status is the sole acceptance verdict.
Testability: deterministic
Consequence: blocks
Verification: proof
Command: ["sh", "sources/run.sh"]
Pattern: ubiquitous

## st-002: Test suite passes
Type: technical
Required: yes
Criterion: The application shall provide a POSIX-compatible bin/test.sh that exits zero when go vet and every automated test pass.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: ubiquitous

## st-003: The executable builds with go build
Type: technical
Required: yes
Criterion: The application root shall be a Go module from which go build ./cmd/csvstat produces the csvstat executable using only the standard library.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: ubiquitous

## st-004: Column types are inferred
Type: behavioral
Required: yes
Criterion: The application shall assign each column the first of integer, number, date, boolean, and text that every non-null value satisfies, and empty when it has no non-null value.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: ubiquitous

## st-005: Profiles match the known answers
Type: behavioral
Required: yes
Criterion: The application shall compute rows, nulls, distinct, min, max, sum, mean, and median for every column of the supplied data files, matching the known answers in sources/run.sh byte for byte.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: ubiquitous

## st-006: Frequencies are ordered and limited
Type: behavioral
Required: yes
Criterion: When --freq is given, the application shall list the column's non-null values by count descending then value ascending, at most --top rows.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: event

## st-007: Malformed input is rejected
Type: behavioral
Required: yes
Criterion: If the file is empty, has an empty or repeated header name, has a ragged record, or a named column is absent, then the application shall exit 3 and write nothing to standard output.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: unwanted

## st-008: Errors map to exit codes
Type: behavioral
Required: yes
Criterion: If the file cannot be read, then the application shall exit 1; if the arguments are malformed, then the application shall exit 2; in both cases it shall write nothing to standard output.
Testability: deterministic
Consequence: blocks
Verification: proof
Pattern: unwanted
