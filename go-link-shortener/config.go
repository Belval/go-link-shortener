package main

import (
	"encoding/json"
	"os"
)

var (
	config *Config
)

// Config : Service configuration struct
type Config struct {
	Domain       string `json:"DOMAIN"`
	URLLength    int    `json:"URL_LENGTH"`
	Port         string `json:"PORT"`
	DatabasePath string `json:"DATABASE_PATH"`
}

func loadConfig() error {
	confFile, err := os.Open("config.json")
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(confFile)
	jsonParser.Decode(&config)
	return nil
}
