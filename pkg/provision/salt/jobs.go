package salt

import (
	"github.com/daixijun/go-salt/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Jobs() []salt.Job {
	result, e := c.client.ListJobs(c.context)
	errors.PanicOnError(e)

	return result
}
