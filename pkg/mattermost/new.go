package mattermost

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/mattermost/mattermost-server/v6/model"
)

func New(
	host string,
	token string,
	team string,
	channel string,
) *Client {
	client := model.NewAPIv4Client(
		fmt.Sprintf("%s://%s", web.SecureScheme, host),
	)
	client.SetOAuthToken(token)
	result := &Client{client: client}

	if team != "" {
		result.team = result.Team(team)

		if channel != "" {
			result.channel = result.Channel(result.team, channel)
		}
	}

	return result
}
