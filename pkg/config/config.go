package config

import (
	"fmt"
	"os"
)

const (
	CONFIG_FILE_DIR       = ".mypocket"
	CONFIG_FILE_NAME      = "config"
	CONFIG_FILE_EXTENSION = "yaml"
)

type ConfigInterface interface {
	GetConfigFilePath() (string, error)
}

type Config struct{}

func NewConfig() ConfigInterface {
	return &Config{}
}

func (u *Config) GetConfigFilePath() (string, error) {

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%s/%s/%s.%s", home, CONFIG_FILE_DIR, CONFIG_FILE_NAME, CONFIG_FILE_EXTENSION)
	return path, nil
}
