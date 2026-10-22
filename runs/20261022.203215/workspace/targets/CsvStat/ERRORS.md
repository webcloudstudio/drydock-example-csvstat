# BIG ERRORS — action required

- Command: `build`
- Phase: build step
- State: Failed
- Timestamp: 2026-10-22T20:39:26+00:00
- Execution ID: 20261022.203854.435Z-44544656
- Challenge Execution ID: -
- Classification: programmatic acceptance failed: cli-contract-usage, cli-help-usage-failure
- Evidence / logs: /mnt/c/Users/barlo/projects/drydock/uat/CsvStat/runs/20261022.203215/workspace/targets/CsvStat/evidence/block-2.md

## Diagnostic

Block "Block 2 · Service" [block-2] failed its acceptance criteria.
  Story "Parse csvstat commands, options, arguments, and exit-code categories." [cli-contract] does not meet its own acceptance criteria:
    - AC cli-contract-usage — Malformed invocations return the usage-error exit status and produce no standard output.
        assertion: assert result.returncode == 2 → AssertionError
        raised at: cli-contract-usage.py:22
        process exit code: 1
        values at failure:
… (44 more lines truncated)

## Recovery

Run: drydock build CsvStat
  to continue the build. This story resumes where it left off and is retried against the checks it failed.
