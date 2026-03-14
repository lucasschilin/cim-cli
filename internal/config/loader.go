package config

import (
	"os"

	"go.yaml.in/yaml/v3"
)

func loadFile(path string, cfg *Config) error {

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, cfg)
}
