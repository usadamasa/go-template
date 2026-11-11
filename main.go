package main

import (
	"os"

	"github.com/usadamasa/go-template/cmd"
	"github.com/usadamasa/go-template/internal/log"
	versionpkg "github.com/usadamasa/go-template/internal/version"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	info := versionpkg.Resolve(version, commit, date)
	cmd.SetVersion(info.DisplayString())
	if err := cmd.Execute(); err != nil {
		log.Errorln(err)
		os.Exit(1)
	}
}
