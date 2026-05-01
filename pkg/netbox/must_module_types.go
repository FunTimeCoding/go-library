package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/module_type"
)

func (c *Client) MustModuleTypes() []*module_type.Type {
	result, e := c.ModuleTypes()
	errors.PanicOnError(e)

	return result
}
