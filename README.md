# PROJECT_NAME

Short description of what this project does.

## Installation

```bash
go install github.com/OWNER/REPO@latest
```

Or download from [Releases](https://github.com/OWNER/REPO/releases).

## Usage

```bash
BINARY_NAME --help
```

## Development

### Prerequisites

- [Go](https://golang.org/) 1.24+
- [aqua](https://aquaproj.github.io/) for tool management
- [direnv](https://direnv.net/) (optional, recommended)

### Setup

```bash
# Install development tools
aqua install

# Allow direnv (if using direnv)
direnv allow
```

### Commands

```bash
# Build
task build

# Run tests
task test

# Run linters
task lint

# Format code
task format

# Full CI check (format, lint, test, build)
task ci

# Clean build artifacts
task clean
```

## Template Customization

When using this as a template, replace the following placeholders:

| Placeholder | Description |
|-------------|-------------|
| `OWNER` | GitHub owner/organization name |
| `REPO` | Repository name |
| `BINARY_NAME` | Binary/executable name |
| `PROJECT_NAME` | Display name for the project |
| `[YEAR]` | Copyright year in LICENSE |

Files to update:
- `go.mod` - module path
- `main.go` - import paths
- `cmd/root.go` - command name and descriptions
- `Taskfile.yaml` - BINARY_NAME variable
- `.goreleaser.yaml` - project_name and build id
- `LICENSE` - year and owner
- `README.md` - this file

## License

MIT License - see [LICENSE](LICENSE) for details.
