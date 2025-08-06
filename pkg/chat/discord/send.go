package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Send(
	s *discordgo.Session,
	channel string,
	v string,
) *discordgo.Message {
	result, e := s.ChannelMessageSend(channel, v)
	errors.PanicOnError(e)

	return result
}
