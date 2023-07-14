package Handlers

import (
	"github.com/bwmarrin/discordgo"
	"golang-discord-bot/BotsController/GlobalSetting"
	"log"
	"strings"
	"time"
)

var Start time.Time

func (bot BotHandlers) VoiceDelayHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.Contains(message.Content, "!latency") {
		log.Println("Start Voice Delay Handling")

		Start = time.Now()
		_, err := session.ChannelMessageSend(GlobalSetting.VoiceChannelId, "!end")
		latency := time.Now().Sub(Start)
		ErrorHandle(err)

		log.Println("Voice Latency: ", latency)
		_, err = session.ChannelMessageSend(message.ChannelID, "Voice Latency: "+latency.String())
		ErrorHandle(err)
	}
}
