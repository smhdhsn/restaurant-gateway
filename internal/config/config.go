package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// envVar holds the environment variable for application's mode.
const envVar = "APP_MODE"

// defaultMode holds the default configuration file's name.
const defaultMode = "local"

// Config holds the application's configurations.
type Config struct {
	Server   ServerConf             `yaml:"server"`
	Services map[string]ServiceConf `yaml:"services"`
	Auth     AuthConf               `yaml:"auth"`
}

// LoadConf loads configuration files from yaml file.
func LoadConf() (*Config, error) {
	cfg := new(Config)
	if err := readFile(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// readFile reads configurations from a file.
func readFile(c *Config) error {
	mode := defaultMode
	if m := os.Getenv(envVar); m != "" {
		mode = m
	}

	f, err := os.Open(path.Join("./config/", mode, "/config.yaml"))
	if err != nil {
		return err
	}

	return yaml.NewDecoder(f).Decode(c)
}
