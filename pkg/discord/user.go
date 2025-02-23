package discord

import "github.com/bwmarrin/discordgo"

func (c *Client) User() *discordgo.User {
	return c.client.State.User
}
