package discord

import "github.com/bwmarrin/discordgo"

func (c *Client) Guilds() []*discordgo.Guild {
	return c.client.State.Guilds
}
