package Handlers

import (
	"golang-discord-bot/BotsController/DiscordBot"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (bot BotHandlers) VoiceChannelHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	const VoiceChannelId string = "433651657214263327"

	if strings.Contains(message.Content, "!join") {
		voiceChannel, err := session.Channel(VoiceChannelId)
		DiscordBot.ErrorHandle(err)

		// mute: false, deaf: true
		voiceConn, err := session.ChannelVoiceJoin(voiceChannel.GuildID, voiceChannel.ID, false, true)
		DiscordBot.ErrorHandle(err)

		log.Println("Joined voice channel: ", voiceConn.ChannelID)
	}

}
