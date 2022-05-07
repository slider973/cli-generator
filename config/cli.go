package config

import (
	"os"
	"regexp"

	"github.com/hashicorp/go-hclog"
	"github.com/slider973/cli-generator/pkg"
)

var (
	levels *regexp.Regexp = regexp.MustCompile("^(trace|debug|info|warn|error|fatal)$")
	format *regexp.Regexp = regexp.MustCompile("^(json)$")

	// DefaultConfig values for CLI
	DefaultConfig = &Config{
		LogOptions: LogOptions{
			Level:  "info",
			Format: "logfmt",
		},
		CodePath: ".",
		Generate: []pkg.Project{},
	}
)

// Reload configuration
func (c *Config) Reload(logger hclog.Logger) (*Config, error) {
	var (
		// v   *Config = DefaultConfig
		cout *Config
		err  error
	)

	// Parse config file if needed
	if cout, err = LoadFile(logger, c.ConfigFile); err != nil {
		return nil, err
	}

	// Merge overwritting
	// if err = mergo.Merge(v, c, mergo.WithOverride); err != nil {
	// 	return err
	// }

	return cout, nil
}

// LogFlagParse level logs
func (c *LogOptions) LogFlagParse(name string) hclog.Logger {
	var level string

	if levels.MatchString(c.Level) {
		level = c.Level
	} else {
		level = "INFO"
	}

	return hclog.New(&hclog.LoggerOptions{
		Name:       name,
		Level:      hclog.LevelFromString(level),
		JSONFormat: format.MatchString(c.Format),
		Output:     os.Stdout,
	})
}
