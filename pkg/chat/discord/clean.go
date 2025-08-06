package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (c *Client) Clean(
	s *discordgo.Session,
	channel string,
	onlyOwn bool,
) {
	for i, m := range c.Messages(s, channel) {
		if onlyOwn && m.Author.ID != s.State.User.ID {
			continue
		}

		fmt.Printf("Delete %d: %s\n", i, m.ID)
		c.Delete(channel, m)
	}
}
