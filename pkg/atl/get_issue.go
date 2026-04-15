package atl

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) GetIssue(key string) string {
	result, e := c.client.GetIssue(c.context, key)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
