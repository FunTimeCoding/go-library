package salt

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"
)

func (c *Client) MustJob(identifier string) *response.Job {
	result, e := c.Job(identifier)
	errors.PanicOnError(e)

	return result
}
