package cmd

import (
	"github.com/spf13/cobra"
)

var version = "dev"

var rootCmd = &cobra.Command{
	Use:     "go-template",
	Short:   "go-template - Short description",
	Long:    `go-template is a CLI application built with Cobra.`,
	Version: version,
}

func init() {
	rootCmd.SetVersionTemplate("{{.Name}} version {{.Version}} ({{.Date}}, {{.Commit}})\n")
}

func SetVersion(displayVersion string) {
	version = displayVersion
	rootCmd.Version = displayVersion
	rootCmd.SetVersionTemplate("{{.Name}} version {{.Version}}\n")
}

func Execute() error {
	return rootCmd.Execute()
}
