package DiscordBot

import (
	"golang-discord-bot/BotsController/DiscordBot/Handlers"
	"log"

	"github.com/bwmarrin/discordgo"
)

var goBot *discordgo.Session

func Start() {
	botConfig := Handlers.GetConfigWithToken(Handlers.GetVaultToken())
	botHandlers := Handlers.BotHandlers{BotConfig: botConfig}
	goBot, _ = discordgo.New("Bot " + botHandlers.BotConfig.Token)

	// goBot.AddHandler(botHandlers.MessageHandler)
	goBot.AddHandler(botHandlers.VoiceChannelHandler)
	goBot.AddHandler(botHandlers.VoiceDelayHandler)

	err := goBot.Open()
	Handlers.ErrorHandle(err)
	log.Println("Start Discord Bot")
}
