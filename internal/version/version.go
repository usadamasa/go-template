package version

import "runtime/debug"

// Info はバージョン情報を保持する構造体
type Info struct {
	Version string
	Commit  string
	Date    string
}

// テスト用に差し替え可能
var readBuildInfo = debug.ReadBuildInfo

// Resolve はldflagsの値を優先しつつ、未設定時はruntime/debug.ReadBuildInfoでフォールバックする
func Resolve(ldVersion, ldCommit, ldDate string) Info {
	if ldVersion != "dev" {
		return Info{Version: ldVersion, Commit: ldCommit, Date: ldDate}
	}

	info := Info{Version: ldVersion, Commit: ldCommit, Date: ldDate}

	bi, ok := readBuildInfo()
	if !ok {
		return info
	}

	if bi.Main.Version != "" && bi.Main.Version != "(devel)" {
		info.Version = bi.Main.Version
	}
	for _, s := range bi.Settings {
		switch s.Key {
		case "vcs.revision":
			info.Commit = s.Value
		case "vcs.time":
			info.Date = s.Value
		}
	}

	return info
}
