package cmd

import (
	"fmt"
	"runtime"

	"github.com/lucasschilin/cim-cli/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version info",
	RunE: func(cmd *cobra.Command, args []string) error {

		goVersion := runtime.Version()
		osArch := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

		fmt.Printf("cim-cli %s\n", version.Version)
		fmt.Printf("commit: %s\n", version.Commit)
		fmt.Printf("built at: %s\n", version.Date)
		fmt.Printf("go: %s\n", goVersion)
		fmt.Printf("os/arch: %s\n", osArch)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
