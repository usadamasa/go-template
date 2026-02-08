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

tagpr と GoReleaser が動作するために以下の Variables / Secrets の設定が必要:

| 種類 | 名前 | 説明 |
|---|---|---|
| Variable | `TAGPR_APP_ID` | tagpr 用 GitHub App の App ID |
| Secret | `TAGPR_PRIVATE_KEY` | tagpr 用 GitHub App の Private Key |

`GITHUB_TOKEN` は GitHub Actions が自動的に提供するため設定不要｡

### セットアップ手順

1. [GitHub App を作成](https://docs.github.com/en/apps/creating-github-apps) し､以下の権限を付与:
   - Contents: Read & Write
   - Pull Requests: Read & Write
   - Issues: Read
2. 作成した App をリポジトリにインストール
3. 以下のコマンドで Variables / Secrets を設定:

```bash
# App ID を Variable として設定
gh variable set TAGPR_APP_ID --body "<YOUR_APP_ID>"

# Private Key を Secret として設定
gh secret set TAGPR_PRIVATE_KEY < /path/to/private-key.pem
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
