package main

import (
	"fmt"
	
	"github.com/bwmarrin/discordgo"
)

func init() {
	commandList = append(commandList, welcomeUser)
}

func welcomeUser(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// Don't do anything if the event was the bot joining
	if m.Member.GuildID == s.State.User.ID {
		return
	}
	
	s.ChannelMessageSend(config.Channels.GreetingChannel, fmt.Sprintf("Welcome, <@%s>", m.User.ID))
}