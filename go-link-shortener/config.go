package main

import (
	"encoding/json"
	"log"
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
	HTMLPath     string `json:"HTML_PATH"`
}

func loadConfig(configPath string) error {
	confFile, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	jsonParser := json.NewDecoder(confFile)
	jsonParser.Decode(&config)
	return nil
}
