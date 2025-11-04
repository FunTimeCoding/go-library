package kestra

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) Login() map[string]interface{} {
	result, r, e := c.client.DefaultAPI.Login(
		c.context,
	).Username(c.user).Password(c.password).Execute()
	errors.PanicOnWebError(r, e)

	return result
}
