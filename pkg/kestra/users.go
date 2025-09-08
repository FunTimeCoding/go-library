package kestra

import "github.com/kestra-io/client-sdk/go-sdk"

func (c *Client) Users() []kestra_api_client.IAMUserControllerApiUserSummary {
	result, r, e := c.client.UsersAPI.ListUsers(
		c.context,
	).Page(0).Size(10).Execute()
	panicOnError(e, r)

	return result.Results
}
