package mattermost

import "github.com/mattermost/mattermost/server/public/model"

func (c *Client) SubmitDialog(v model.SubmitDialogRequest) *model.SubmitDialogResponse {
	result, r, e := c.client.SubmitInteractiveDialog(c.context, v)
	panicOnError(r, e)

	return result
}
