package mattermost

import "github.com/mattermost/mattermost-server/v6/model"

func (c *Client) SubmitDialog(v model.SubmitDialogRequest) *model.SubmitDialogResponse {
	result, r, e := c.client.SubmitInteractiveDialog(v)
	panicOnError(e, r)

	return result
}
