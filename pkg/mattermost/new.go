package mattermost

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/mattermost/user_map"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/mattermost/mattermost-server/v6/model"
)

func New(o ...OptionFunc) *Client {
	result := &Client{user: user_map.New()}

	for _, p := range o {
		p(result)
	}

	if result.host == "" {
		panic("host required")
	}

	if result.token == "" {
		panic("token required")
	}

	result.client = model.NewAPIv4Client(
		fmt.Sprintf("%s://%s", web.SecureScheme, result.host),
	)
	result.client.SetOAuthToken(result.token)

	if result.teamName != "" {
		result.team = result.Team(result.teamName)

		if result.channelName != "" {
			result.channel = result.ChannelByName(
				result.team,
				result.channelName,
			)
		}
	}

	result.RefreshSocket()

	return result
}
