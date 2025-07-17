package mattermost

import "github.com/mattermost/mattermost-server/v6/model"

func (c *Client) DeletePost(p *model.Post) {
	r, e := c.client.DeletePost(p.Id)
	panicOnError(e, r)
}
