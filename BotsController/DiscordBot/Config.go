package DiscordBot

import (
	"encoding/json"
	"os"
)

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
