package salt

import (
	"github.com/daixijun/go-salt/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Job(identifier string) *salt.Job {
	result, e := c.client.LookupJID(c.context, identifier)
	errors.PanicOnError(e)

	return result
}
