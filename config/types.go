package config

import "github.com/slider973/cli-generator/pkg"

// LogOptions struct parameters for logging level and format
type LogOptions struct {
	Level  string `yaml:"level,omitempty"`
	Format string `yaml:"format,omitempty"`

	// Catches all undefined fields and must be empty after parsing.
	XXX map[string]interface{} `yaml:",inline"`
}

// Config CLI structs
type Config struct {
	ConfigFile string        `yaml:"-"`
	LogOptions LogOptions    `yaml:"log_options,omitempty"`
	CodePath   string        `yaml:"code_path,omitempty"`
	Generate   []pkg.Project `yaml:"generate,omitempty"`
}
