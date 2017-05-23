package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {
	var config mainConfig
	loadConfig("config", &config, true) // Load config, panic if it fails
	
	
}