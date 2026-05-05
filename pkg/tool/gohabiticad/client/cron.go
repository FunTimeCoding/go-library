package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Cron() string {
	result, e := c.client.RunCron(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
