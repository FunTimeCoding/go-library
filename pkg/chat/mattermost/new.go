package mattermost

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chat/mattermost/user_map"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func New(o ...Option) *Client {
	result := &Client{context: context.Background()}

	for _, p := range o {
		p(result)
	}

	if result.user == nil {
		result.user = user_map.New(nil, []string{})
	}

	if result.host == "" {
		panic("host required")
	}

	if result.token == "" {
		panic("token required")
	}

	result.client = model.NewAPIv4Client(
		fmt.Sprintf("%s://%s", constant.SecureScheme, result.host),
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
