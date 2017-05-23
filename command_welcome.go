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
	
	logInfo(fmt.Sprintf("%s (%s#%s) joined the server, sending welcome.", m.User.ID, m.User.Username, m.User.Discriminator))

	s.ChannelMessageSend(config.Channels.GreetingChannel, fmt.Sprintf(
		"Welcome to the ASTOST English Discord, <@%s>!\n"+
			"We must verify new users before they are able to enter the main rooms as a safety precaution.\n"+
			"Please post a link to your ASTOST profile or tell us your ASTOST username for verification.\n"+
			"Please be patient, it may take some time to verify users.\n"+
			"If you need help with anything, please @ one of the Admins or Moderators in the user list.", m.User.ID))
}
