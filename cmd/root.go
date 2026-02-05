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
	Use:     "BINARY_NAME",
	Short:   "PROJECT_NAME - Short description",
	Long:    `PROJECT_NAME is a CLI application built with Cobra.`,
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
