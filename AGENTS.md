# Agent Instructions

Go port of Python [inflect](https://pypi.org/project/inflect/) for English word inflection. See `README.md` for API documentation.

## Populating Context

Use the `reference` tool to explore the codebase. It provides a map of all public functions with signatures, descriptions, locations, and test coverage.

```
reference                           # List all 114 functions
reference format:groups             # Show available groups
reference group:verbs               # Functions in a group
reference filter:ordinal            # Search by name/description
reference group:numbers format:verbose  # Full details
```

Alternatively, run `make reference` for raw CSV output.

## Issue Tracking

Uses **bd** (beads). Run `bd onboard` to get started.

```bash
bd ready              # Find available work
bd show <id>          # View issue details
bd update <id> --status in_progress  # Claim work
bd close <id>         # Complete work
bd sync               # Sync with git
```

## Working on Tickets

When working on a ticket, delegate the bulk of the work to a **pi subagent** running `claude-opus-4-5`:

```bash
pi -p --model claude-opus-4-5 "your task description"
```

See `pi --help` for options. Once the subagent completes, review and verify the work before pushing to remote.

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
