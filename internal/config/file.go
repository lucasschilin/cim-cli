package config

import (
	"os"
	"path/filepath"
)

func EnsureConfigFile(path string) error {

	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.WriteFile(path, []byte(DefaultConfigTemplate), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
