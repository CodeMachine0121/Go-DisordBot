package Handlers

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func (bot BotHandlers) MessageHandler(session *discordgo.Session, messageCreate *discordgo.MessageCreate) {
	if messageCreate.Author.ID == session.State.User.ID {
		return
	}

	if strings.Contains(messageCreate.Content, "!ping") {

		_, err := session.ChannelMessageSend(messageCreate.ChannelID, "bot Testing ...")

		ErrorHandle(err)

		latencyMessage := "[+] Latency Rate: " + GetSessionLatencyRate(session) + "%"

		_, err = session.ChannelMessageSend(messageCreate.ChannelID, latencyMessage)
	}
}

func GetSessionLatencyRate(session *discordgo.Session) string {
	duration := session.LastHeartbeatAck.Sub(session.LastHeartbeatSent)
	latencyLimit := time.Second * 3
	return (duration * 100 / latencyLimit).String()
}
