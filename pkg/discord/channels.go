package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Channels(guild string) []*discordgo.Channel {
	result, e := c.client.GuildChannels(guild)
	errors.PanicOnError(e)

	return result
}
