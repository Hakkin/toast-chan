package main

import (
	"fmt"
	
	"github.com/bwmarrin/discordgo"
)

func welcomeUser(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	fmt.Println(m.Member)

	// Don't do anything if the event was the bot joining
	if m.Member.GuildID == s.State.User.ID {
		return
	}
	
	s.ChannelMessageSend(config.Channels.GreetingChannel, fmt.Sprintf("Welcome, <@%s>", m.User.ID))
}