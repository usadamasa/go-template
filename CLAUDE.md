# CLAUDE.md

This is a Go project template with a comprehensive development setup.

## Project Structure

```
.
├── cmd/           # CLI command definitions (Cobra)
├── internal/      # Private application code
│   ├── log/       # Logging package (forbidigo-compliant)
│   └── version/   # Version resolution (ldflags / runtime/debug)
├── aqua/          # Local aqua registry + policy for tools missing from the standard registry
├── main.go        # Application entry point
└── Taskfile.yaml  # Task runner configuration
```

## Development Commands

```bash
task build          # Build the binary
task test           # Run tests with race detection
task lint           # Run yamllint and golangci-lint
task lint:sec       # Run govulncheck and gosec
task arch-lint      # go-arch-lint + analyze-arch-lint + analyze-modularity
task pinact         # Pin GitHub Actions to commit SHAs
task pinact:verify  # Verify pinned actions
task format         # Format code (go mod tidy, goimports, go fmt)
task dev            # format + lint + test
task ci             # format + lint + lint:sec + arch-lint + pinact:verify + test + build
task clean          # Remove build artifacts
```

Tools are managed by aqua (`aqua.yaml`). `spm-go` and `analyze-*` come from the local registry
(`aqua/registry.yaml`), which requires `AQUA_POLICY_CONFIG=aqua/policy.yaml`. `.envrc`, the Taskfile,
and CI set it; export it yourself when calling those tools directly.

## Code Guidelines

### Logging

Do not use `fmt.Print*` for console output. Use `internal/log` package instead:

```go
import "github.com/usadamasa/go-template/internal/log"

log.Println("Hello")
log.Printf("Value: %d", 42)
log.Errorln("Something went wrong")
log.Errorf("failed: %v", err)

// In tests, swap the writers for a buffer
log.Out = &buf
```

The API is kept to four functions plus the `Out` / `Err` writers on purpose: a larger facade trips
the public-ratio gate in `task arch-lint` (analyze-modularity).

This is enforced by the `forbidigo` linter. Exceptions are made for:
- Test files (`*_test.go`)
- Entry point (`main.go`)
- CLI commands (`cmd/`)
- The log package itself (`internal/log/`)

### Architecture Metrics

The lint setup follows the `go-arch-metrics` plugin
(`go-arch-metrics:setup` / `measure` / `evaluate` / `remediate` skills):

- `.golangci.yml` enables gocognit / gocyclo / cyclop / funlen / nestif / maintidx with the
  recommended thresholds. Raise a threshold only to just above the measured maximum, then tighten.
- `.go-arch-lint.yml` declares dependency direction by role (`cli`, `shared_internal`), not per
  package. Adding a package under `internal/` should not require a config change.
  `analyze-arch-lint --strict .` guards the config itself from degrading into a package-tree copy.
- Take a baseline with `go-arch-metrics:measure` before and after structural changes.

### Testing

- Write tests for all new functionality
- Use table-driven tests where appropriate
- Run `task test` before committing

## CI/CD

- **CI Workflow**: Runs on PRs - test, lint (+ security), architecture lint, pinact verify
- **TagPR Workflow**: Runs on main push - creates release PRs and GitHub releases
- Uses aqua for tool version management
- GitHub Actions are pinned to commit SHAs by pinact (`task pinact` to update)
- GoReleaser for cross-platform builds

## Template Placeholders

| Placeholder | Replace With |
|-------------|--------------|
| `OWNER` | GitHub owner name |
| `REPO` | Repository name |
| `BINARY_NAME` | Executable name |
| `PROJECT_NAME` | Display name |
| `[YEAR]` | Copyright year |
