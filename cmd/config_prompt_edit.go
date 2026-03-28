package cmd

import (
	"errors"
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/editor"
	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/lucasschilin/cim-cli/internal/prompt"
	"github.com/spf13/cobra"
)

var configPromptEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit prompt",
	RunE: func(cmd *cobra.Command, args []string) error {
		repoRoot, _ := git.GetRepoRoot()
		if repoRoot == "" {
			return errors.New("Not inside a git repository")
		}

		var path string

		switch {
		case globalFlag:
			path = prompt.GlobalPromptPath()

		case repoUserFlag:
			path = prompt.RepoUserPromptPath(repoRoot)

		case repoFlag:
			path = prompt.RepoSharedPromptPath(repoRoot)

		default:
			return errors.New("You need to inform which prompt configuration to edit.\nUse one of the following flags: --global, --repo, --repo-user")
		}

		err := prompt.EnsurePromptFile(path)
		if err != nil {
			return fmt.Errorf("Error creating prompt file: %v", err)
		}

		editor.Open(path)

		return nil

	},
}

func init() {
	configPromptCmd.AddCommand(configPromptEditCmd)

	configPromptEditCmd.Flags().BoolVar(&globalFlag, "global", false, "Edit global config")
	configPromptEditCmd.Flags().BoolVar(&repoFlag, "repo", false, "Edit shared repo config")
	configPromptEditCmd.Flags().BoolVar(&repoUserFlag, "repo-user", false, "Edit user config for this repo")
}
