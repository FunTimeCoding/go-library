package mattermost

import "github.com/mattermost/mattermost-server/v6/model"

func (c *Client) Post(p *model.Post) *model.Post {
	result, r, e := c.client.CreatePost(p)
	panicOnError(e, r)

	return result
}
