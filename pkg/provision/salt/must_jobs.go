package salt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"
)

func (c *Client) MustJobs() []response.Job {
	result, e := c.Jobs()
	errors.PanicOnError(e)

	return result
}
