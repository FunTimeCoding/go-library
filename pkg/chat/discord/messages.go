package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Messages(
	s *discordgo.Session,
	channel string,
) []*discordgo.Message {
	result, e := s.ChannelMessages(
		channel,
		MessageLimit,
		"",
		"",
		"",
	)
	errors.PanicOnError(e)

	return result
}
