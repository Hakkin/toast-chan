package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var config mainConfig
var commandList []interface{}

func main() {
	loadConfig("config", &config, true) // Load config, panic if there is an error, so no need to check it
	fmt.Println(config)
	
	bot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		panic(err) // TODO: Change to proper logging once made
	}
	
	// Starting bot
	err = bot.Open()
	if err != nil {
		panic(err) // TODO: Change to proper logging once made
	}
	
	// Register commands
	for _, function := range commandList {
		bot.AddHandler(function)
	}
	
	
	// Run until CTRL+C
	fmt.Println("Bot running, CTRL+C to shut down.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly shut down
	fmt.Println("Bot shutting down...")
	bot.Close()
}