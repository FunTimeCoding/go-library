package discord

import (
	"github.com/bwmarrin/discordgo"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) UserChannel(
	s *discordgo.Session,
	user string,
) *discordgo.Channel {
	result, e := s.UserChannelCreate(user)
	errors.PanicOnError(e)

	return result
}
