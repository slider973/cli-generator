package config

import (
	"errors"
	"io/ioutil"

	"github.com/hashicorp/go-hclog"
	"gopkg.in/yaml.v2"
)

// LoadFile parses the given YAML file into a Config.
func LoadFile(logger hclog.Logger, filename string) (*Config, error) {
	var (
		cfg     *Config = DefaultConfig
		content []byte
		err     error
	)

	if filename != "" {
		if content, err = ioutil.ReadFile(filename); err != nil {
			return nil, err
		}

		if err = yaml.Unmarshal(content, cfg); err != nil {
			return nil, err
		}

		return cfg, nil
	}

	return nil, errors.New("filename is empty")
}
