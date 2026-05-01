package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/module_bay"
)

func (c *Client) MustModuleBays() []*module_bay.Bay {
	result, e := c.ModuleBays()
	errors.PanicOnError(e)

	return result
}
