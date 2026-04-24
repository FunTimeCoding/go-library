package tool

import (
	"github.com/mattermost/mattermost/server/public/model"
	"strings"
)

func (t *Tool) channelDisplayName(c *model.Channel) string {
	if c.DisplayName != "" {
		return c.DisplayName
	}

	if c.Type == model.ChannelTypeDirect {
		me := t.client.Me()
		parts := strings.Split(c.Name, "__")
		resolved := false

		for _, p := range parts {
			if p != me.Id {
				other := t.client.User(p)

				if other != nil {
					return other.Username
				}

				resolved = true

				break
			}
		}

		if !resolved {
			return me.Username + " (self)"
		}
	}

	return c.Name
}
