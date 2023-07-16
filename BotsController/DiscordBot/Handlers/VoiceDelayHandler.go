package Handlers

import (
	GlobalSetting "golang-discord-bot/BotsController/GlobalSetting"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var start time.Time
var commandChannel string

const TestRound int = 10

func (bot BotHandlers) VoiceDelayHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID && message.Content != "!start" {
		return
	}

	var latencyList []time.Duration
	if strings.Contains(message.Content, "!start") {
		log.Println("Received start command")
		log.Println("Start monitoring at ", time.Now())

		bot.MonitoringLatency(session, latencyList)

	}
}

func (BotHandlers) MonitoringLatency(session *discordgo.Session, latencyList []time.Duration) {
	for index := 0; ; index++ {
		latency := CalculateLatency(session)
		latencyList = append(latencyList, latency)

		log.Println("Sent message latency: " + latency.String())
		time.Sleep(5 * time.Second)

		if IsLatencyListFull(latencyList) {
			averageLatency := CalculateAverageLatency(latencyList)
			log.Println("Average latency: " + averageLatency.String())

			if IsAverageLatencyHigh(averageLatency) {
				log.Println("High Latency Alert!!!!!!  ", averageLatency)
				SendMessage(session, commandChannel, "@everyone High Latency Alert!!!! : "+averageLatency.String())
			}
			latencyList = ResetLatencyList()
		}
	}
}

func ResetLatencyList() []time.Duration {
	return make([]time.Duration, 0)
}

func CalculateAverageLatency(latencyList []time.Duration) time.Duration {
	var allLatency time.Duration
	for _, latency := range latencyList {
		allLatency = allLatency + latency
	}
	return allLatency / time.Duration(len(latencyList))
}

func CalculateLatency(session *discordgo.Session) time.Duration {
	start = time.Now()
	SendMessage(session, GlobalSetting.VoiceChannelId, "!testing")
	latency := time.Since(start)
	return latency
}

func IsAverageLatencyHigh(averageLatency time.Duration) bool {
	return averageLatency > GlobalSetting.HighLatencyThreshold
}

func IsLatencyListFull(latencyList []time.Duration) bool {
	return len(latencyList) >= GlobalSetting.LatencyTestRound
}
