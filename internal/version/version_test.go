package version

import (
	"runtime/debug"
	"testing"
)

func TestResolve_LdflagsSet(t *testing.T) {
	// ldflagsで設定済みの場合はそのまま返す
	info := Resolve("v1.2.3", "abc1234", "2024-01-01T00:00:00Z")

	if info.Version != "v1.2.3" {
		t.Errorf("Version = %q, want %q", info.Version, "v1.2.3")
	}
	if info.Commit != "abc1234" {
		t.Errorf("Commit = %q, want %q", info.Commit, "abc1234")
	}
	if info.Date != "2024-01-01T00:00:00Z" {
		t.Errorf("Date = %q, want %q", info.Date, "2024-01-01T00:00:00Z")
	}
}

func TestResolve_DevWithBuildInfo(t *testing.T) {
	// readBuildInfoを差し替えてBuildInfoありのケースをテスト
	original := readBuildInfo
	t.Cleanup(func() { readBuildInfo = original })

	readBuildInfo = func() (*debug.BuildInfo, bool) {
		return &debug.BuildInfo{
			Main: debug.Module{
				Version: "v0.3.0",
			},
			Settings: []debug.BuildSetting{
				{Key: "vcs.revision", Value: "def5678abcdef5678abcdef5678abcdef5678abc"},
				{Key: "vcs.time", Value: "2024-06-15T12:00:00Z"},
			},
		}, true
	}

	info := Resolve("dev", "none", "unknown")

	if info.Version != "v0.3.0" {
		t.Errorf("Version = %q, want %q", info.Version, "v0.3.0")
	}
	if info.Commit != "def5678abcdef5678abcdef5678abcdef5678abc" {
		t.Errorf("Commit = %q, want %q", info.Commit, "def5678abcdef5678abcdef5678abcdef5678abc")
	}
	if info.Date != "2024-06-15T12:00:00Z" {
		t.Errorf("Date = %q, want %q", info.Date, "2024-06-15T12:00:00Z")
	}
}

func TestResolve_DevWithBuildInfoDevel(t *testing.T) {
	// BuildInfoのバージョンが(devel)の場合はdevのまま
	original := readBuildInfo
	t.Cleanup(func() { readBuildInfo = original })

	readBuildInfo = func() (*debug.BuildInfo, bool) {
		return &debug.BuildInfo{
			Main: debug.Module{
				Version: "(devel)",
			},
			Settings: []debug.BuildSetting{
				{Key: "vcs.revision", Value: "aaa1111"},
				{Key: "vcs.time", Value: "2024-03-01T09:00:00Z"},
			},
		}, true
	}

	info := Resolve("dev", "none", "unknown")

	if info.Version != "dev" {
		t.Errorf("Version = %q, want %q", info.Version, "dev")
	}
	if info.Commit != "aaa1111" {
		t.Errorf("Commit = %q, want %q", info.Commit, "aaa1111")
	}
	if info.Date != "2024-03-01T09:00:00Z" {
		t.Errorf("Date = %q, want %q", info.Date, "2024-03-01T09:00:00Z")
	}
}

func TestResolve_DevWithoutBuildInfo(t *testing.T) {
	// BuildInfoが取得できない場合はデフォルト値のまま
	original := readBuildInfo
	t.Cleanup(func() { readBuildInfo = original })

	readBuildInfo = func() (*debug.BuildInfo, bool) {
		return nil, false
	}

	info := Resolve("dev", "none", "unknown")

	if info.Version != "dev" {
		t.Errorf("Version = %q, want %q", info.Version, "dev")
	}
	if info.Commit != "none" {
		t.Errorf("Commit = %q, want %q", info.Commit, "none")
	}
	if info.Date != "unknown" {
		t.Errorf("Date = %q, want %q", info.Date, "unknown")
	}
}
