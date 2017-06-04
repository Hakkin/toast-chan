package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type mainConfigS struct {
	Token    string
}

var config mainConfigS

func init() {

}

func main() {
	// Load config, panic if there is an error, so no need to check it
	loadConfig("config", &config, true)

	bot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		logFatal("There was an error creating the Discord sesion: " + err.Error())
	}

	// Starting bot
	err = bot.Open()
	if err != nil {
		logFatal("There was an error while opening a connection to Discord: " + err.Error())
	}

	// Register commands
	for name, function := range commandList {
		logInfo(fmt.Sprintf("Registering command [%s]", name))
		bot.AddHandler(function)
	}

	// Run until CTRL+C
	logInfo("Bot running, CTRL+C to shut down.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly shut down
	logInfo("Bot shutting down...")
	bot.Close()
}
