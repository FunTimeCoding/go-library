package mattermost

import (
	"github.com/funtimecoding/go-library/pkg/mattermost/user_map"
	"github.com/mattermost/mattermost-server/v6/model"
)

type Client struct {
	host  string
	token string

	teamName    string
	channelName string

	client    *model.Client4
	team      *model.Team
	channel   *model.Channel
	webSocket *model.WebSocketClient
	meCache   *model.User
	user      *user_map.Map
}
