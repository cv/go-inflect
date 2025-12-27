# Claude Workflow for go-inflect

## Project Overview

Go port of the Python [inflect](https://pypi.org/project/inflect/) library for English word inflection.

## Development Workflow

```bash
bd ready              # See available issues
bd update <id> --status in_progress  # Claim work
go test -v            # Run tests
gofmt -w . && go vet  # Lint
git commit && git push
bd close <id> -r "reason"
bd sync               # Sync with remote
```

## Current State (2025-12-26)

### Completed: 39 tasks
- **Core**: An/A, Plural/Singular, Ordinal/OrdinalWord, NumberToWords
- **Joining**: Join, JoinWithConj, JoinWithSep
- **Verbs**: PresentParticiple
- **Comparison**: Compare, CompareNouns
- **Custom Defs**: DefNoun, DefA/DefAn, DefVerb, DefAdj + Undef/Reset
- **Classical Mode**: ClassicalAll, ClassicalAncient, ClassicalPersons, ClassicalZero, ClassicalHerd, ClassicalNames
- **Utility**: No, Num, Gender
- **Docs**: README.md, CONTRIBUTING.md, LICENSE

### Remaining: 11 tasks
- CI/CD setup
- Inflect() inline text parsing
- Regex patterns in custom definitions
- Comma/semicolon detection in Join
- NumberToWords options (decimal, threshold, comma, group)

### Quality
- 96.1% test coverage
- 5,228 lines of code
- 66 benchmarks
