package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func init() {
	const (
		command = "log_join"
	)
	addCommand(command, logJoin)
}

func logJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// Don't do anything if the event was the bot joining
	if m.Member.GuildID == s.State.User.ID {
		return
	}
	
	var highlight string
	
	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		logError(fmt.Sprintf("There was an error while getting the roles: %s", err))
		highlight = ""
	}

	for _, role := range roles {
		_, ok := config.HighlightRoles.ModChannel[role.ID]
		if (ok) {
			highlight += fmt.Sprintf(" <@&%s>", role.ID)
		}
	}
	
	logInfo(fmt.Sprintf("%s#%s joined the server.", m.User.Username, m.User.Discriminator))
	
	s.ChannelMessageSend(config.Channels.ModChannel, fmt.Sprintf(
		"%s\r\n"+
		"Ring ding ding! New user here! (＾∀＾●)ﾉｼ <@%s> has joined the server", highlight, m.User.ID))
}