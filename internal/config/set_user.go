package config

import (
	"encoding/json"
	"os"
)

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) SetUser(user string) error {
	c.CurrentUsername = user
	err := write(*c)
	if err != nil {
		return err
	}
	return nil
}
