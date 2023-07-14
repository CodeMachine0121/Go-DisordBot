package DiscordBot

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (botHandlers BotHandlers) VoiceChannelHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	const VoiceChannelId string = "433651657214263327"

	log.Println("Received message: " + message.Content)

	if strings.Contains(message.Content, "!join") {
		voiceChannel, err := session.Channel(VoiceChannelId)
		ErrorHandle(err)

		voiceConn, err := session.ChannelVoiceJoin(voiceChannel.GuildID, voiceChannel.ID, false, true)
		ErrorHandle(err)

		log.Printf("Joined voice channel: %s\n"+voiceChannel.Name+" with ID: ", voiceConn.ChannelID)
	}

}
