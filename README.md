# go-template

Short description of what this project does.

## Installation

```bash
go install github.com/usadamasa/go-template@latest
```

Or download from [Releases](https://github.com/usadamasa/go-template/releases).

## Usage

```bash
go-template --help
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

## CI/CD Setup

### GitHub Actions ワークフロー

| ワークフロー | トリガー | 内容 |
|---|---|---|
| CI (`ci.yaml`) | PR, merge group | テスト､リント |
| tagpr (`tagpr.yaml`) | main push | リリースPR作成､タグ付け､GoReleaser でリリース |

### 必要な GitHub リポジトリ設定

tagpr と GoReleaser が動作するために以下の Secret の設定が必要:

| 種類 | 名前 | 説明 |
|---|---|---|
| Secret | `TAGPR_PRIVATE_KEY` | tagpr 用 GitHub App の Private Key |

App ID は秘密情報ではないため `tagpr.yaml` に直接書いてあり､リポジトリごとの
Variable 設定は不要｡`GITHUB_TOKEN` は GitHub Actions が自動的に提供するため設定不要｡

### セットアップ手順

1. GitHub App `usadamasa-tagpr` をこのリポジトリにインストールする
   (<https://github.com/settings/installations/101558454>)
2. Private Key を Secret として設定する:

```bash
gh secret set TAGPR_PRIVATE_KEY -R <owner>/<repo> < /path/to/private-key.pem
```

App を新しく作り直す場合に必要な権限は Contents: Read & Write､
Pull Requests: Read & Write､Issues: Read の 3 つ｡

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
