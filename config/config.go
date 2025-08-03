package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port     int `yaml:"port"`
	Database struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"database"`
}

func ReadConfig(path string) (*Config, error) {
	var config Config
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
