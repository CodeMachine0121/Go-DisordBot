package DiscordBot

import "log"

func ErrorHandle(err error) {
	if err != nil {
		log.Fatal("something is error, pls contact support!")
		return
	}
}
