{
  description = "Development environment (Go, lint, and CI tools)";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs =
    { nixpkgs, ... }:
    let
      systems = [
        "aarch64-darwin"
        "aarch64-linux"
        "x86_64-darwin"
        "x86_64-linux"
      ];
      forAllSystems = f: nixpkgs.lib.genAttrs systems (system: f nixpkgs.legacyPackages.${system});
      # nixos-unstable の go_1_27 は 1.27.3 で、net/http の MaxBytesReader のデータ競合 (CVE-2026-97054) と
      # cgo 経由でビルド時に任意コードが実行される cmd/go などの脆弱性 (CVE-2026-97033〜97043) が残る。
      # channel の go_1_27 が 1.27.4 より古いときだけソースを差し替える。
      # nixpkgs が追いついたら (flake.lock の go_1_27 が 1.27.4 以上になったら) この分岐ごと消し、go_1_27 に戻す
      devGo =
        pkgs:
        if pkgs.lib.versionOlder pkgs.go_1_27.version "1.27.4" then
          pkgs.go_1_27.overrideAttrs (old: rec {
            version = "1.27.4";
            src = pkgs.fetchurl {
              url = "https://go.dev/dl/go${version}.src.tar.gz";
              hash = "sha256-c8+4VxrsDAEXdofbhxY9Rqog6XeWcncokaoXZU2aYJk=";
            };
          })
        else
          pkgs.go_1_27;
    in
    {
      devShells = forAllSystems (
        pkgs:
        let
          go = devGo pkgs;
          # nixpkgs に無い go-arch-metrics plugin 同梱の CLI。plugin のリリースタグをそのまま使う
          go-arch-metrics = pkgs.buildGoModule.override { inherit go; } rec {
            pname = "go-arch-metrics";
            version = "2026.0927.0";
            src = pkgs.fetchFromGitHub {
              owner = "usadamasa";
              repo = "go-arch-metrics";
              tag = version;
              hash = "sha256-cwYDbuqIM7uFikWvZdtYf7LAO/5Bup2u9508V9Jz90w=";
            };
            vendorHash = "sha256-g+yaVIx4jxpAQ/+WrGKxhVeliYx7nLQe/zsGpxV4Fn4=";
            subPackages = [
              "cmd/analyze-modularity"
              "cmd/analyze-arch-lint"
            ];
          };
        in
        {
          default = pkgs.mkShellNoCC {
            packages = with pkgs; [
              # go.mod の go ディレクティブと同じバージョン。ずれたら GOTOOLCHAIN=local で即座に失敗させる
              go
              go-task
              gotools
              golangci-lint
              gosec
              govulncheck
              go-arch-lint
              go-arch-metrics
              goreleaser
              yamllint
            ];
            env.GOTOOLCHAIN = "local";
          };
        }
      );
    };
}
