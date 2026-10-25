# Blockers: Rate Limit Hit
- target: CsvStat
- source: drydock plan CsvStat --override --plan-override-model sonnet
- declared: 2026-10-25T14:46:16-04:00

<!-- drydock-state: {"owner": "plan CsvStat", "execution": "20261025.184613.465Z-9cb0ea89", "attempts": 1} -->

## blocker-001: Provider usage limit
- kind: environmental
- class: provider rate limit
- auto-continue: after 16:50

### Error

provider rate limit 429: You've hit your session limit · resets 4:50pm (America/New_York)
