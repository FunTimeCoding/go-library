package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) DefaultTeam() *model.Team {
	return c.team
}
