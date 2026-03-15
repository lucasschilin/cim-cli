package cmd

import (
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/config"
	"github.com/lucasschilin/cim-cli/internal/editor"
	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/spf13/cobra"
)

var editorFlag bool

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit configuration",
	Run: func(cmd *cobra.Command, args []string) {
		repoRoot, _ := git.GetRepoRoot()

		var path string

		switch {
		case globalFlag:
			path = config.GlobalConfigPath()

		case repoUserFlag:
			path = config.RepoUserConfigPath(repoRoot)

		case repoFlag:
			path = config.RepoSharedConfigPath(repoRoot)

		default:
			fmt.Println(
				"You need to inform which configuration to edit.\nUse one of the following flags: --global, --repo, --repo-user",
			)
			return
		}

		err := config.EnsureConfigFile(path)
		if err != nil {
			fmt.Println("Error creating config file", err)
			return
		}

		if editorFlag != false {
			editor.Open(path)
			return
		}

		editor.Open(path)

	},
}

func init() {
	configCmd.AddCommand(configEditCmd)

	configEditCmd.Flags().BoolVar(&globalFlag, "global", false, "Edit global config")
	configEditCmd.Flags().BoolVar(&repoFlag, "repo", false, "Edit shared repo config")
	configEditCmd.Flags().BoolVar(&repoUserFlag, "repo-user", false, "Edit user config for this repo")

	configEditCmd.Flags().BoolVarP(&editorFlag, "editor", "e", false, "Edit config using editor")
}
