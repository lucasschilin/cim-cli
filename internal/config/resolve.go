package config

import (
	"os"
)

func Resolve(repoRoot string) (*Config, error) {

	cfg := &Config{}

	global := GlobalConfigPath()
	repoUser := RepoUserConfigPath(repoRoot)
	repoShared := RepoSharedConfigPath(repoRoot)

	if fileExists(global) {
		loadFile(global, cfg)
	}

	if fileExists(repoUser) {
		loadFile(repoUser, cfg)
	}

	if fileExists(repoShared) {
		loadFile(repoShared, cfg)
	}

	return cfg, nil
}

func fileExists(path string) bool {

	_, err := os.Stat(path)

	return err == nil
}
