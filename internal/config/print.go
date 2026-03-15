package config

import "go.yaml.in/yaml/v3"

func ToYAML(cfg *Config) (string, error) {

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
