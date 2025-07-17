package mattermost

import "github.com/mattermost/mattermost-server/v6/model"

func (c *Client) Reactions(p *model.Post) []*model.Reaction {
	result, r, e := c.client.GetReactions(p.Id)
	panicOnError(e, r)

	return result
}
