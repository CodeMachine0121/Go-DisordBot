package Handlers

import (
	"golang-discord-bot/BotsController"
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
		// var allLatency time.Duration
		for index := 0; ; index++ {
			latency := CalculateLatency(session)
			latencyList = append(latencyList, latency)
			// allLatency += latency
			log.Println("Sent message latency: " + latency.String())
			time.Sleep(5 * time.Second)

			if len(latencyList) == 10 {
				averageLatency := CalculateAverageLatency(latencyList)
				log.Println("Average latency: " + averageLatency.String())

				if averageLatency > time.Second*2 {
					log.Println("High Latency Alert!!!!!!  ", averageLatency)
					SendMessage(session, commandChannel, "@everyone High Latency Alert!!!! : "+averageLatency.String())
				}

				// reset list
				latencyList = make([]time.Duration, 0)
				log.Println("Empty list: ", latencyList)
			}
		}

	}
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
	SendMessage(session, BotsController.VoiceChannelId, "!testing")
	latency := time.Since(start)
	return latency
}
