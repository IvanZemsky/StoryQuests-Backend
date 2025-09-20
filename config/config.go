package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port     int `yaml:"port"`
	Database struct {
		Name        string `yaml:"name"`
		Port        int    `yaml:"port"`
		Host        string `yaml:"host"`
		UserName    string `yaml:"username"`
		Password    string `yaml:"password"`
		ClusterCode string `yaml:"cluster_code"`
		ClusterName string `yaml:"cluster_name"`
	} `yaml:"database"`
	DBType string `yaml:"db_type"`
	Origin string `yaml:"origin"`
	JWT_secret string `yaml:"jwt_secret"`
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
