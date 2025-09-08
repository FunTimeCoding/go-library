package kestra

func (c *Client) Login() map[string]interface{} {
	result, r, e := c.client.DefaultAPI.Login(
		c.context,
	).Username(c.user).Password(c.password).Execute()
	panicOnError(e, r)

	return result
}
