package DiscordBot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var goBot *discordgo.Session

func Start() {
	botHandlers := BotHandlers{BotConfig: ReadConfig()}
	goBot, _ = discordgo.New("Bot " + botHandlers.BotConfig.Token)

	goBot.AddHandler(botHandlers.MessageHandler)
	goBot.AddHandler(botHandlers.MessageHandler)

	err := goBot.Open()
	if err != nil {
		log.Fatal("something error in enabling discord bot")
	}

	log.Println("Start Discord Bot")
}
