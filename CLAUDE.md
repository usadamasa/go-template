# CLAUDE.md

This is a Go project template with a comprehensive development setup.

## Project Structure

```
.
├── cmd/           # CLI command definitions (Cobra)
├── internal/      # Private application code
│   └── log/       # Logging package (forbidigo-compliant)
├── main.go        # Application entry point
└── Taskfile.yaml  # Task runner configuration
```

## Development Commands

```bash
task build    # Build the binary
task test     # Run tests with race detection
task lint     # Run yamllint and golangci-lint
task format   # Format code (go mod tidy, goimports, go fmt)
task dev      # format + lint + test
task ci       # format + lint + test + build
task clean    # Remove build artifacts
```

## Code Guidelines

### Logging

Do not use `fmt.Print*` for console output. Use `internal/log` package instead:

```go
import "github.com/OWNER/REPO/internal/log"

// Use package-level functions
log.Println("Hello")
log.Printf("Value: %d", 42)
log.Errorln("Something went wrong")

// Or create a custom logger
logger := log.New(os.Stdout, os.Stderr)
logger.Println("Custom output")
```

This is enforced by the `forbidigo` linter. Exceptions are made for:
- Test files (`*_test.go`)
- Entry point (`main.go`)
- CLI commands (`cmd/`)
- The log package itself (`internal/log/`)

### Testing

- Write tests for all new functionality
- Use table-driven tests where appropriate
- Run `task test` before committing

## CI/CD

- **CI Workflow**: Runs on PRs - lints and tests
- **TagPR Workflow**: Runs on main push - creates release PRs and GitHub releases
- Uses aqua for tool version management
- GoReleaser for cross-platform builds

## Template Placeholders

| Placeholder | Replace With |
|-------------|--------------|
| `OWNER` | GitHub owner name |
| `REPO` | Repository name |
| `BINARY_NAME` | Executable name |
| `PROJECT_NAME` | Display name |
| `[YEAR]` | Copyright year |
