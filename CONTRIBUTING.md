# Contributing to go-inflect

## Reporting Bugs

Create an issue with:
- Steps to reproduce
- Expected vs. actual behavior
- Go version and OS
- Minimal code example

## Suggesting Features

1. Check existing issues first
2. Describe the use case
3. Reference Python inflect behavior if applicable
4. Show example API usage

## Development

### Setup

```bash
git clone https://github.com/cv/go-inflect.git
cd go-inflect
./.githooks/setup.sh  # enables pre-commit hooks
```

### Workflow

```bash
make help     # see all targets
make test     # run tests
make lint     # run linter
```

The pre-commit hook runs build, test, and lint automatically.

### Code Style

- Follow [Effective Go](https://go.dev/doc/effective_go)
- Use table-driven tests
- Document all exported symbols
- Keep PRs focused—one feature or fix per PR

## License

Contributions are licensed under [MIT](LICENSE).
