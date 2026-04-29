package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Alerts() string {
	result, e := c.client.GetAlerts(c.context, &client.GetAlertsParams{})
	errors.PanicOnError(e)

	return web.ReadString(result)
}
