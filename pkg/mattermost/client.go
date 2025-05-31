package mattermost

import "github.com/mattermost/mattermost-server/v6/model"

type Client struct {
	client  *model.Client4
	team    *model.Team
	channel *model.Channel
}
