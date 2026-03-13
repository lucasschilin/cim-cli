package cmd

import (
	"fmt"

	"github.com/lucasschilin/commit-improver-cli/internal/commit"
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

		message, err := commit.ReadCommitMessage(path)
		if err != nil {
			fmt.Println("Error reading commit message:", err)
			return
		}

		fmt.Println("Original message:")
		fmt.Println(message)

		newMessage := "fix(auth): just a little test :D"

		err = commit.WriteCommitMessage(path, newMessage)
		if err != nil {
			fmt.Println("Error writing commit message:", err)
			return
		}

		fmt.Println("Message replaced")
	},
}

func init() {
	rootCmd.AddCommand(hookCmd)
}
