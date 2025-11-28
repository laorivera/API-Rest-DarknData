package services

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Port           string   `json:"port"`
	Host           string   `json:"host"`
	AllowedOrigins []string `json:"allowed_origins"`
	APIKey         string   `json:"api_key"`
}

func LoadConfig() Config {
	file, err := os.Open("env.json")
	if err != nil {
		log.Fatal("Error opening config.json:", err)
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Fatal("Error decoding config.json:", err)
	}
	return config
}

func GetKey() string {
	key := LoadConfig().APIKey
	if key == "" {
		return "hashed-development-key-here"
	}
	return key
}
