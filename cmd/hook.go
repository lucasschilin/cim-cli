package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Intercept commit messages",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Commit message file not provided")
			return
		}

		path := args[0]

		fmt.Println("Hook called")
		fmt.Println("Commit message file:", path)
	},
}

func init() {
	rootCmd.AddCommand(hookCmd)
}
