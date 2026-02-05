# .claude/CLAUDE.md

Project-specific Claude Code configuration.

## Build and Test

```bash
task build    # Build the binary
task test     # Run tests
task lint     # Run linters
task ci       # Full CI check
```

## Code Style

- Use `internal/log` instead of `fmt.Print*` for console output
- Follow Go conventions and idioms
- Run `task format` before committing

## Architecture

- `cmd/` - Cobra CLI commands
- `internal/` - Private packages
- Entry point: `main.go`
