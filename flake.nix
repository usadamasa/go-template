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
    in
    {
      devShells = forAllSystems (
        pkgs:
        let
          # nixpkgs に無い go-arch-metrics plugin 同梱の CLI。plugin のリリースタグをそのまま使う
          go-arch-metrics = pkgs.buildGoModule.override { go = pkgs.go_1_27; } rec {
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
              # go.mod の go ディレクティブと同じ版。ずれたら GOTOOLCHAIN=local で即座に失敗させる
              go_1_27
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
