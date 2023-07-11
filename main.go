package main

import (
	"golang-discord-bot/DicordBot"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	DicordBot.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	return
}
