# Blockers: Rate Limit Hit
- target: CsvStat
- source: drydock plan CsvStat --override --plan-override-model sonnet
- owner: plan CsvStat
- declared: 2026-10-24T22:22:15Z
- execution: 20261024.222213.429Z-7be13be5
- attempts: 1

## blocker-001: Provider usage limit
- kind: environmental
- class: provider rate limit
- auto-continue: after 20:10
### Error
provider rate limit 429: You've hit your session limit · resets 8:10pm (America/New_York)
### Commander Resolution
