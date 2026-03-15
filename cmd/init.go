package cmd

import (
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/config"
	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Commit Improver Hook",
	Run: func(cmd *cobra.Command, args []string) {
		repoRoot, err := git.GetRepoRoot()
		if err != nil {
			fmt.Println("Not inside a git repository")
			return
		}
		_, err = config.Resolve(repoRoot)
		if err != nil {
			fmt.Println("Config error:", err)
			return
		}

		err = git.InstallCommitMsgHook(repoRoot)
		if err != nil {
			fmt.Println("Error installing cim-cli hook:", err)
			return
		}

		fmt.Println("cim-cli hook installed :D")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
