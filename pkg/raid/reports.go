package raid

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Reports() string {
	result, e := c.client.GetReports(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
