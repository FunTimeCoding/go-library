package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) DeletePost(p *model.Post) {
	r, e := c.client.DeletePost(c.context, p.Id)
	panicOnError(e, r)
}
