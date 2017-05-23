package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func init() {
	const (
		command = "welcome"
	)
	addCommand(command, welcomeUser)
}

func welcomeUser(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// Don't do anything if the event was the bot joining
	if m.User.ID == s.State.User.ID {
		return
	}

	s.ChannelMessageSend(config.Channels.GreetingChannel, fmt.Sprintf(
		"Welcome to the ASTOST English Discord, <@%s> (｡＾＿＾｡)づ！！！\n"+
			"To make sure we know who you are, please post a link to your ASTOST profile or let us know what your username is. Sometimes it takes us a while to check this so please be patient.\n"+
			"If you need help with anything, @ one of the admins or moderators in the userlist and we'll get back to you as soon as possible.\n"+
			"Once again, welcome to ATOAST! ＼(｡＾ ｏ ＾｡)／", m.User.ID))
}
