package config

type Uri struct {
	InputPath       string `yaml:"inputPath"`
	Method          string `yaml:"method"`
	RedirectionHost string `yaml:"redirectionHost"`
	RedirectionPath string `yaml:"redirectionPath"`
	AccessLevel     int    `yaml:"accessLevel"`
	Enabled         bool   `yaml:"enabled"`
}