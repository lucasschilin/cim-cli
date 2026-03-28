package cmd

import (
	"github.com/spf13/cobra"
)

var configParamsCmd = &cobra.Command{
	Use:   "params",
	Short: "Manage cim-cli params configuration",
}

func init() {
	configCmd.AddCommand(configParamsCmd)
}
