package Handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (bot BotHandlers) MessageHandler(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	if messageCreate.Author.ID == session.State.User.ID {
		return
	}

	if strings.Contains(messageCreate.Content, "!ping") {
		_, err := session.ChannelMessageSend(messageCreate.ChannelID, "bot Testing ...")
		ErrorHandle(err)
	}
}
