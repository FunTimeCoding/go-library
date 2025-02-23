package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Delete(
	channel string,
	m *discordgo.Message,
) {
	errors.PanicOnError(c.client.ChannelMessageDelete(channel, m.ID))
}
