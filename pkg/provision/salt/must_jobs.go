package salt

import (
	"github.com/daixijun/go-salt/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustJobs() []salt.Job {
	result, e := c.Jobs()
	errors.PanicOnError(e)

	return result
}
