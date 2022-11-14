package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"app"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func New(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
