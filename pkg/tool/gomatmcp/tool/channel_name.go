package tool

import (
	"github.com/mattermost/mattermost/server/public/model"
	"strings"
)

func (t *Tool) channelDisplayName(
	ch *model.Channel,
) string {
	if ch.DisplayName != "" {
		return ch.DisplayName
	}

	if ch.Type == model.ChannelTypeDirect {
		me := t.client.Me()
		parts := strings.Split(ch.Name, "__")
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

	return ch.Name
}
