package config

type Config struct {
	Provider    string `yaml:"provider"`
	Model       string `yaml:"model"`
	APIKey      string `yaml:"api_key"`
	Language    string `yaml:"language"`
	DiffLimit   int    `yaml:"diff_limit"`
	AutoConfirm bool   `yaml:"auto_confirm"`
}
