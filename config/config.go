package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PathToLogFile string `yaml:"PathToLogFile"`
}

func NewConfig() *Config {
	return &Config{
		PathToLogFile: "",
	}
}

func (c *Config) GetPathToLogFile() string {
	return c.PathToLogFile
}

func ReadConfig(filePath string) (*Config, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg := NewConfig()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
