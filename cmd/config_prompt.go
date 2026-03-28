package cmd

import (
	"github.com/spf13/cobra"
)

var configPromptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Manage cim-cli custom prompt configuration",
}

func init() {
	configCmd.AddCommand(configPromptCmd)
}
