package Handlers

import (
	"golang-discord-bot/BotsController"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (bot BotHandlers) VoiceChannelHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if strings.Contains(message.Content, "!join") {
		voiceChannel, err := session.Channel(BotsController.VoiceChannelId)
		ErrorHandle(err)

		// mute: false, deaf: true
		voiceConn, err := session.ChannelVoiceJoin(voiceChannel.GuildID, voiceChannel.ID, false, true)
		ErrorHandle(err)

		log.Println("Joined voice channel: ", voiceConn.ChannelID)

	}

}
