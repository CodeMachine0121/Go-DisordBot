package main

import (
	"golang-discord-bot/BotHandlers/DiscordBot"
)

func main() {
	DiscordBot.Start()
	<-make(chan struct{})
}
