package cmd

import (
	"github.com/lucasschilin/cim-cli/internal/version"
	"github.com/spf13/cobra"
)

var (
	globalFlag   bool
	repoFlag     bool
	repoUserFlag bool
)

var rootCmd = &cobra.Command{
	Use:     "cim-cli",
	Short:   "Commit Improver CLI",
	Version: version.Version,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.SetVersionTemplate("cim-cli {{.Version}}\n")
}
