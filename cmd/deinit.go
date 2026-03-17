package cmd

import (
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/spf13/cobra"
)

var deinitCmd = &cobra.Command{
	Use:   "deinit",
	Short: "Remove cim-cli git hook",
	RunE: func(cmd *cobra.Command, args []string) error {
		repoRoot, err := git.GetRepoRoot()
		if err != nil {
			return fmt.Errorf("Not inside a git repository: %v", err)
		}

		err = git.RemoveCommitMsgHook(repoRoot)
		if err != nil {
			return fmt.Errorf("Error removing cim-cli hook: %v", err)
		}

		fmt.Println("cim-cli hook removed :(")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deinitCmd)
}
