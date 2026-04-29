package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) DeletePost(p *model.Post) error {
	_, e := c.client.DeletePost(c.context, p.Id)

	return e
}
