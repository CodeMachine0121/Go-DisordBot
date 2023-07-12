package DiscordBot

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Token     string
	BotPrefix string
}

func ReadConfig() Config {
	file, _ := ioutil.ReadFile("./setting.json")
	config := Config{}
	err := json.Unmarshal(file, &config)
	if err != nil {
		return Config{}
	}
	return config
}
