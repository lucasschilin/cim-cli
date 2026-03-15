package cmd

import (
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/config"
	"github.com/lucasschilin/cim-cli/internal/git"
	"github.com/spf13/cobra"
)

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show configuration",
	Run: func(cmd *cobra.Command, args []string) {
		repoRoot, _ := git.GetRepoRoot()
		if repoRoot == "" {
			fmt.Println("Not inside a git repository")
			return
		}

		var (
			cfg *config.Config
			err error
		)

		switch {
		case globalFlag:
			cfg, err = config.LoadConfigFile(config.GlobalConfigPath())

		case repoUserFlag:
			cfg, err = config.LoadConfigFile(config.RepoUserConfigPath(repoRoot))

		case repoFlag:
			cfg, err = config.LoadConfigFile(config.RepoSharedConfigPath(repoRoot))

		default:
			cfg, err = config.Resolve(repoRoot)
		}

		if err != nil {
			fmt.Println("Config error:", err)
			return
		}

		output, err := config.ToYAML(cfg)
		if err != nil {
			fmt.Println("Error serializing config:", err)
			return
		}

		fmt.Println(output)

	},
}

func init() {
	configCmd.AddCommand(configShowCmd)

	configShowCmd.Flags().BoolVar(&globalFlag, "global", false, "Show global config")
	configShowCmd.Flags().BoolVar(&repoFlag, "repo", false, "Show shared repo config")
	configShowCmd.Flags().BoolVar(&repoUserFlag, "repo-user", false, "Show user config for this repo")
}
