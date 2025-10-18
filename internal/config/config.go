package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func Read() Config {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}
	}

	var c Config
	if err := json.Unmarshal(data, &c); err != nil {
		return Config{}
	}

	return c
}


func (c *Config) SetUser(name string) error{
	c.CurrentUsername = name
	return write(*c)
}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	filepath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	errWrite := os.WriteFile(filepath, jsonData, 0644)
	if errWrite != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := homeDir + "/" + configFileName
	return filePath, nil
}