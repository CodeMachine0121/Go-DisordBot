package Handlers

import (
	"encoding/json"
	"log"
	"os"
)

type BotHandlers struct {
	BotConfig Config
}

func ErrorHandle(err error) {
	if err != nil {
		log.Fatal("something is error, pls contact support!")
		return
	}
}

type Config struct {
	Token     string
	BotPrefix string
}

func ReadConfig() Config {
	config := Config{}

	configFile, _ := os.ReadFile("./setting.json")
	err := json.Unmarshal(configFile, &config)

	if err != nil {
		return Config{}
	}
	return config
}
