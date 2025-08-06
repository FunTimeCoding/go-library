package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Channel(
	s *discordgo.Session,
	channel string,
) *discordgo.Channel {
	result, e := s.Channel(channel)
	errors.PanicOnError(e)

	return result
}
