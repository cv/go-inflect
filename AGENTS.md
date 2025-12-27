# Agent Instructions

Go port of Python [inflect](https://pypi.org/project/inflect/) for English word inflection. See `README.md` for API documentation.

## Issue Tracking

Uses **bd** (beads). Run `bd onboard` to get started.

```bash
bd ready              # Find available work
bd show <id>          # View issue details
bd update <id> --status in_progress  # Claim work
bd close <id>         # Complete work
bd sync               # Sync with git
```

## Session Completion

**Work is NOT complete until `git push` succeeds.**

1. File issues for any remaining work
2. Run quality gates if code changed (pre-commit hook handles this)
3. Close finished issues, update in-progress ones
4. Push to remote:
   ```bash
   git pull --rebase && bd sync && git push
   ```
5. Verify: `git status` must show "up to date with origin"
