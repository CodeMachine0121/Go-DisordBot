package DiscordBot

import (
	"golang-discord-bot/BotsController/DiscordBot/Handlers"
	"log"

	"github.com/bwmarrin/discordgo"
)

var goBot *discordgo.Session

func Start() {
	botHandlers := Handlers.BotHandlers{BotConfig: ReadConfig()}
	goBot, _ = discordgo.New("Bot " + botHandlers.BotConfig.Token)

	// goBot.AddHandler(botHandlers.MessageHandler)
	goBot.AddHandler(botHandlers.VoiceChannelHandler)

	err := goBot.Open()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Start Discord Bot")
}
