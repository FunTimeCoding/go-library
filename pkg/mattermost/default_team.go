package mattermost

import "github.com/mattermost/mattermost-server/v6/model"

func (c *Client) DefaultTeam() *model.Team {
	return c.team
}
