package config

import (
	"os"
	"path/filepath"
)

func GlobalConfigPath() string {
	home, _ := os.UserHomeDir()

	return filepath.Join(home, ".cim-cli", "config.yaml")
}

func RepoSharedConfigPath(repoRoot string) string {
	return filepath.Join(repoRoot, ".cim-cli", "config.yaml")
}

func RepoUserConfigPath(repoRoot string) string {
	return filepath.Join(repoRoot, ".git", ".cim-cli", "config.yaml")
}
