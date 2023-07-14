package Handlers

import (
	"github.com/bwmarrin/discordgo"
	"golang-discord-bot/BotsController/DiscordBot"
	"log"
	"strings"
	"time"
)

func (bot BotHandlers) VoiceDelayHandler(session *discordgo.Session, message *discordgo.MessageCreate) {

	const VoiceChannelId string = "433651657214263327"

	if strings.Contains(message.Content, "!latency") {
		start := time.Now()
		msg, err := session.ChannelMessageSend(VoiceChannelId, "Test")
		DiscordBot.ErrorHandle(err)

		if msg.ID == message.ID {
			latency := time.Now().Sub(start)
			log.Println("Voice Latency: ", latency)
		}
	}
}
