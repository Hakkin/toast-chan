package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type newUserConfigS struct {
	Channels struct {
		GreetingChannel string
		ModChannel      string
	}
	HighlightRoles map[string]bool
	Messages struct {
		Greeting string
		Notify string
	}
}

var newUserConfig newUserConfigS

func init() {
	// Load config, panic if there is an error, so no need to check it
	loadConfig("newUser", &newUserConfig, true)
	addCommand("newUser", newUser)
}

func newUser(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	// Don't do anything if the event was the bot joining
	if m.User.ID == s.State.User.ID {
		return
	}
	
	logInfo(fmt.Sprintf("%s#%s joined the server.", m.User.Username, m.User.Discriminator))
	
	// Notify mods
	s.ChannelMessageSend(newUserConfig.Channels.ModChannel, parseMessage(newUserConfig.Messages.Notify, s, m))

	// Welcome user
	s.ChannelMessageSend(newUserConfig.Channels.GreetingChannel, parseMessage(newUserConfig.Messages.Greeting, s, m))
}

func parseMessage(message string, s *discordgo.Session, m *discordgo.GuildMemberAdd) string {
	parsedMessage := message

	// {{user}} replace
	parsedMessage = strings.Replace(parsedMessage, "{{user}}", "<@" + m.User.ID + ">", -1)
	
	// {{highlight}} replace
	parsedMessage = strings.Replace(parsedMessage, "{{highlight}}", getHighlightString(s, m), -1)
	
	return parsedMessage
}

func getHighlightString(s *discordgo.Session, m *discordgo.GuildMemberAdd) string {
	var highlight string

	roles, err := s.GuildRoles(m.GuildID)
	if err != nil {
		logError(fmt.Sprintf("There was an error while getting the roles: %s", err))
		return ""
	}

	for _, role := range roles {
		_, ok := newUserConfig.HighlightRoles[role.ID]
		if ok {
			highlight += fmt.Sprintf(" <@&%s>", role.ID)
		}
	}
	
	return highlight
}