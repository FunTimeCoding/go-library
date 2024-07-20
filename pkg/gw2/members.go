package gw2

import (
	"github.com/MrGunflame/gw2api"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Members(identifier string) []*gw2api.GuildMember {
	result, e := c.client.GuildMembers(identifier)
	errors.PanicOnError(e)

	return result
}
