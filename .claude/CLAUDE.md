# .claude/CLAUDE.md

Project-specific Claude Code configuration.

## Build and Test

```bash
task build      # Build the binary
task test       # Run tests
task lint       # Run linters (yamllint, golangci-lint)
task lint:sec   # Run govulncheck and gosec
task arch-lint  # Dependency direction + arch-lint config health + modularity
task ci         # Full CI check
```

## Code Style

- Use `internal/log` instead of `fmt.Print*` for console output
- Follow Go conventions and idioms
- Run `task format` before committing

## Architecture

- `cmd/` - Cobra CLI commands
- `internal/` - Private packages (must not depend on `cmd/`; enforced by `.go-arch-lint.yml`)
- Entry point: `main.go`
- Metrics thresholds and dependency rules follow the `go-arch-metrics` plugin skills
