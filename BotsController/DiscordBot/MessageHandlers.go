package DiscordBot

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func (botHandlers BotHandlers) MessageHandler(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	if messageCreate.Author.ID == session.State.User.ID {
		return
	}

	if strings.Contains(messageCreate.Content, botHandlers.BotConfig.BotPrefix) {
		_, err := session.ChannelMessageSend(messageCreate.ChannelID, "bot")
		if err != nil {
			log.Fatal("something error in sending message")
			return
		}
	}
}
