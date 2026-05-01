package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/module"
)

func (c *Client) MustModules() []*module.Module {
	result, e := c.Modules()
	errors.PanicOnError(e)

	return result
}
