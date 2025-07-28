package mattermost

import "github.com/mattermost/mattermost-server/v6/model"

func (c *Client) OpenDialog(v model.OpenDialogRequest) {
	r, e := c.client.OpenInteractiveDialog(v)
	panicOnError(e, r)
}
