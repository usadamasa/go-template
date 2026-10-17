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

- [Nix](https://nixos.org/) with flakes enabled
- [direnv](https://direnv.net/) + [nix-direnv](https://github.com/nix-community/nix-direnv)

### Setup

```bash
# Load the devShell via nix-direnv
direnv allow
```

Go と開発ツール (go-task、goimports、golangci-lint、gosec、govulncheck、go-arch-lint、
analyze-modularity / analyze-arch-lint、goreleaser、yamllint) は `flake.nix` の devShell で揃える。
版は `flake.lock` が固定する nixpkgs で決まり、CI も同じ flake を使う。
nixpkgs に無い `analyze-*` は `flake.nix` の中で go-arch-metrics のタグからビルドする。

`go.mod` の `go` ディレクティブは devShell の `go_1_27` と揃え、`GOTOOLCHAIN=local` でずれたら失敗させる。
ツールの版を上げるときは `nix flake update` で `flake.lock` を更新する。

### Commands

```bash
# Build
task build

# Run tests
task test

# Run linters (yamllint, golangci-lint)
task lint

# Security checks (govulncheck, gosec)
task lint:sec

# Architecture lint (go-arch-lint, analyze-arch-lint, analyze-modularity)
task arch-lint

# Format code
task format

# Full CI check
task ci

# Clean build artifacts
task clean
```

### Architecture Metrics

lint 構成は Claude Code plugin [go-arch-metrics](https://github.com/usadamasa/go-arch-metrics) の
`setup` skill に沿っている｡

| ファイル | 役割 |
|---|---|
| `.golangci.yml` | 認知的複雑度・循環複雑度・関数長・ネスト深さ・保守性指数のしきい値 |
| `.go-arch-lint.yml` | パッケージ依存方向 (役割単位の component: `cli` → `shared_internal`) |
| `flake.nix` | nixpkgs に無い `analyze-modularity` / `analyze-arch-lint` を go-arch-metrics のタグからビルド |

ベースライン測定は `go-arch-metrics:measure`､数値の解釈は `go-arch-metrics:evaluate` を使う｡

## CI/CD Setup

### GitHub Actions ワークフロー

| ワークフロー | トリガー | 内容 |
|---|---|---|
| CI (`ci.yaml`) | PR, merge group | テスト､リント (+ セキュリティ)､アーキテクチャリント､actions-lock 検証 |
| tagpr (`tagpr.yaml`) | main push | リリースPR作成､タグ付け､GoReleaser でリリース |

GitHub Actions は `uses:` にタグを書き、対応するコミット SHA を `.github/workflows/actions.lock` に記録している｡
バージョンを上げるときは `uses:` のタグを書き換えて `gh actions-lock` を実行し、workflow と lockfile を同じコミットに入れる｡

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
- `.go-arch-lint.yml` - vendors (add third-party packages as you introduce them)
- `LICENSE` - year and owner
- `README.md` - this file

## License

MIT License - see [LICENSE](LICENSE) for details.
