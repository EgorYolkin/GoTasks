package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigInterface interface {
	NewConfig(string) (Config, error)
}

type Config struct {
	Bot struct {
		TelegramApiKey string `yaml:"telegram_api_key"`
	} `yaml:"bot"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		DBName   string `yaml:"dbname"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

func NewConfig(configFilePath string) (Config, error) {
	f, err := os.Open(configFilePath)

	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	var config Config
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(&config); err != nil {
		return Config{}, nil
	}

	return config, nil
}
