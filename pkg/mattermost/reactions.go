package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Reactions(p *model.Post) []*model.Reaction {
	result, r, e := c.client.GetReactions(c.context, p.Id)
	panicOnError(e, r)

	return result
}
