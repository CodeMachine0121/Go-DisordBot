package Handlers

import (
	"golang-discord-bot/BotsController/GlobalSetting"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (bot BotHandlers) VoiceChannelHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if strings.Contains(message.Content, "!latency") {
		log.Println("Start Voice Delay Handling")
		commandChannel = message.ChannelID
		SendMessage(session, GlobalSetting.VoiceChannelId, "!start")
	}

}
