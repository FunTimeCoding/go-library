package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) OpenDialog(v model.OpenDialogRequest) {
	r, e := c.client.OpenInteractiveDialog(c.context, v)
	panicOnError(e, r)
}
