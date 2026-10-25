# BIG ERRORS — action required

- Command: `build`
- Phase: build step
- State: Failed
- Timestamp: 2026-10-25T01:49:54+00:00
- Execution ID: 20261025.013803.939Z-a6311d9d
- Challenge Execution ID: -
- Classification: programmatic acceptance failed: profile-quoting-and-trim
- Evidence / logs: /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261025.011759/workspace/targets/CsvStat/evidence/block-2.md

## Diagnostic

Block "Block 2 · Service" [block-2] failed its acceptance criteria.
  Story "Write the profile table for every column in file order." [profile-output] does not meet its own acceptance criteria:
    - AC profile-quoting-and-trim — A value containing a comma is quoted on output, and padding spaces in header names and values are trimmed.
        assertion: assert code == 0, (code, out, err) → AssertionError: (3, '', 'csvstat: rejected input: /tmp/tmp9phg2ko5/data.csv: parse error on line 2, column 3: bare " in non-quoted-field\n')
        raised at: profile-quoting-and-trim.py:20
        process exit code: 1
        error: AssertionError: (3, '', 'csvstat: rejected input: /tmp/tmp9phg2ko5/data.csv: parse error on line 2, column 3: bare " in non-quoted-field\n')
        values at failure:
… (11 more lines truncated)

## Recovery

Run: drydock build CsvStat
  to continue the build. This story resumes where it left off and is retried against the checks it failed.
