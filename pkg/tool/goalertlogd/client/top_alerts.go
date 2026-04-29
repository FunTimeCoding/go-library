package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
	"time"
)

func (c *Client) TopAlerts(
	n int,
	start time.Time,
	end time.Time,
) string {
	result, e := c.client.GetTopAlerts(
		c.context,
		&client.GetTopAlertsParams{
			N:     &n,
			Start: &start,
			End:   &end,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
