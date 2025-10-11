package kestra

import "github.com/kestra-io/client-sdk/go-sdk"

func (c *Client) Users() []kestra_api_client.IAMUserControllerApiUserSummary {
	r, s, e := c.client.UsersAPI.ListUsers(
		c.context,
	).Page(0).Size(10).Execute()
	panicOnError(e, s)

	return r.Results
}
