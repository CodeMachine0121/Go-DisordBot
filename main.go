package main

import (
	"golang-discord-bot/BotsController/DiscordBot"
)

func main() {
	DiscordBot.Start()
	<-make(chan struct{})
}
