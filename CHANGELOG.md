# Changelog

## [v0.0.8](https://github.com/usadamasa/go-template/compare/v0.0.7...v0.0.8) - 2026-10-25

- chore(deps): Bump actions/create-github-app-token from 3.2.0 to 3.3.0 by @dependabot[bot] in https://github.com/usadamasa/go-template/pull/25
- chore(deps): Bump Songmu/tagpr from 1.20.3 to 1.21.3 by @dependabot[bot] in https://github.com/usadamasa/go-template/pull/24
- setup-nix を gh-manage の共有 action へ移行する by @usadamasa in https://github.com/usadamasa/go-template/pull/30
- ci: pull_request の types に edited を足す by @usadamasa in https://github.com/usadamasa/go-template/pull/32

## [v0.0.7](https://github.com/usadamasa/go-template/compare/v0.0.6...v0.0.7) - 2026-10-21

- ci: Dependabot で gomod・github-actions・nix を週次で更新する by @usadamasa in https://github.com/usadamasa/go-template/pull/22
- devShell の Go を 1.27.3 に上げて GO-2026-6809 を解消する by @usadamasa in https://github.com/usadamasa/go-template/pull/27

## [v0.0.6](https://github.com/usadamasa/go-template/compare/v0.0.5...v0.0.6) - 2026-10-17

- chore: ツールチェーンを更新し go-arch-metrics と pinact を組み込む by @usadamasa in https://github.com/usadamasa/go-template/pull/13
- chore: GitHub Actions の固定を pinact から gh-actions-lock へ移行する by @usadamasa in https://github.com/usadamasa/go-template/pull/15
- chore: gh-actions-lock を v0.2.0 へ上げ ci-status-check を ubuntu-slim にする by @usadamasa in https://github.com/usadamasa/go-template/pull/16
- chore: 開発ツールの管理を aqua から Nix の devShell へ移す by @usadamasa in https://github.com/usadamasa/go-template/pull/17
- docs: devShell の読み込み手順から nix develop の補足を消す by @usadamasa in https://github.com/usadamasa/go-template/pull/18
- ci: リリースの Go と goreleaser を Nix の devShell から取る by @usadamasa in https://github.com/usadamasa/go-template/pull/19
- ci: tagpr の App トークン取得を client-id で指定する by @usadamasa in https://github.com/usadamasa/go-template/pull/20
- ci: tagpr の App トークンの権限を絞る by @usadamasa in https://github.com/usadamasa/go-template/pull/21

## [v0.0.5](https://github.com/usadamasa/go-template/compare/v0.0.4...v0.0.5) - 2026-09-06

- chore(tagpr): App ID をワークフローに直書きして repository variable を不要にする by @usadamasa in https://github.com/usadamasa/go-template/pull/11

## [v0.0.4](https://github.com/usadamasa/go-template/compare/v0.0.3...v0.0.4) - 2026-02-08
- バージョン表示フォーマットをorm-discovery-mcp-goに合わせる by @usadamasa in https://github.com/usadamasa/go-template/pull/9

## [v0.0.3](https://github.com/usadamasa/go-template/compare/v0.0.2...v0.0.3) - 2026-02-08
- バージョン表示でcommit/dateが不明な場合に括弧を省略 by @usadamasa in https://github.com/usadamasa/go-template/pull/7

## [v0.0.2](https://github.com/usadamasa/go-template/compare/v0.0.1...v0.0.2) - 2026-02-08
- go install時のバージョン情報をReadBuildInfoでフォールバック by @usadamasa in https://github.com/usadamasa/go-template/pull/4

## [v0.0.2](https://github.com/usadamasa/go-template/compare/v0.0.1...v0.0.2) - 2026-02-08
- go install時のバージョン情報をReadBuildInfoでフォールバック by @usadamasa in https://github.com/usadamasa/go-template/pull/4

## [v0.0.1](https://github.com/usadamasa/go-template/commits/v0.0.1) - 2026-02-08
- テンプレートplaceholderを実際の値に置換 by @usadamasa in https://github.com/usadamasa/go-template/pull/1
- tagprワークフローを1ジョブ構成にリファクタリング by @usadamasa in https://github.com/usadamasa/go-template/pull/2
