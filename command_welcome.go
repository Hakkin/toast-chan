package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const (
	command = "welcome"
)

func init() {
	addCommand(command, welcomeUser)
}

func welcomeUser(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// Don't do anything if the event was the bot joining
	if m.Member.GuildID == s.State.User.ID {
		return
	}
	
	logInfo(fmt.Sprintf("%s#%s joined the server, sending welcome.", m.User.Username, m.User.Discriminator))

	s.ChannelMessageSend(config.Channels.GreetingChannel, fmt.Sprintf(
		"Welcome to the ASTOST English Discord, <@%s>!\n"+
			"As a safety precaution, we must verify new users before they are able to enter the main channels..\n"+
			"Please post a link to your ASTOST profile or tell us your ASTOST username to expedite this process.\n"+
			"It may take us some time to verify users, so please be patient.\n"+
			"If you need help with anything, @ one of the admins or moderators in the user list.", m.User.ID))
}
