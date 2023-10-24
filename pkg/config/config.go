package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TelegramToken string `yaml:"telegram_token"`
	TelegramHost  string `yaml:"telegram_host"`
	BonbastProxy  string `yaml:"bonbast_proxy"`
}

func ParseConfig() (*Config, error) {
	f, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var c Config

	if err = yaml.NewDecoder(f).Decode(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
