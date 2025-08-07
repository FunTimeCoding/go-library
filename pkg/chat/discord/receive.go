package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func (c *Client) receive(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case PingCommand:
		c.Send(s, c.UserChannel(s, m.Author.ID).ID, "pong")
	case CleanCommand:
		if m.GuildID == "" {
			// Direct message
			channel := c.UserChannel(s, m.Author.ID).ID
			c.Clean(s, channel, true)
			c.Send(s, channel, "Done")
		} else {
			// Channel
			c.Clean(s, m.ChannelID, false)
			c.Send(s, m.ChannelID, "Done")
		}
	case DetailsCommand:
		fmt.Printf("Channel: %+v\n", c.Channel(s, m.ChannelID))
	default:
		fmt.Printf("Message: %+v\n", m.Message)
	}
}
