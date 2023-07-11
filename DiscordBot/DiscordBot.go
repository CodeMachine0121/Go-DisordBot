package DiscordBot

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"os"
	signal "os/signal"
	"strings"
	"syscall"
)

type Config struct {
	Token     string
	BotPrefix string
}

var goBot *discordgo.Session
var config Config

func Start() {
	config = ReadConfig()
	goBot, _ = discordgo.New("Bot " + config.Token)

	goBot.AddHandler(messageHandler)
	goBot.AddHandler(messageHandler)

	err := goBot.Open()
	if err != nil {
		log.Fatal("something error in enabling discord bot")
	}
}

func messageHandler(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	if messageCreate.Author.ID == session.State.User.ID {
		return
	}

	if strings.Contains(messageCreate.Content, config.BotPrefix) {
		_, err := session.ChannelMessageSend(messageCreate.ChannelID, "bot")
		if err != nil {
			log.Fatal("something error in sending message")
			return
		}
	}
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
