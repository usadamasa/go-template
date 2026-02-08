package cmd

import (
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:     "go-template",
	Short:   "go-template - Short description",
	Long:    `go-template is a CLI application built with Cobra.`,
	Version: version,
}

func init() {
	rootCmd.SetVersionTemplate("{{.Name}} version {{.Version}} ({{.Date}}, {{.Commit}})\n")
}

func SetVersion(v, c, d string) {
	version = v
	commit = c
	date = d
	rootCmd.Version = v
	rootCmd.SetVersionTemplate("{{.Name}} version " + v + " (" + d + ", " + c + ")\n")
}

func Execute() error {
	return rootCmd.Execute()
}
