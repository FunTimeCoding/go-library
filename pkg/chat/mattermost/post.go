package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) Post(p *model.Post) *model.Post {
	result, r, e := c.client.CreatePost(c.context, p)
	panicOnError(r, e)

	return result
}
