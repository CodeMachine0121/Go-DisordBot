package DiscordBot

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	signal "os/signal"
	"syscall"
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

func Stop() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err := goBot.Close()
	if err != nil {
		log.Fatal("something error in shutting down bot ")
		return
	}
}
