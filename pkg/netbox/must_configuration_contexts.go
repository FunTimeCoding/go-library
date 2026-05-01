package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/configuration_context"
)

func (c *Client) MustConfigurationContexts() []*configuration_context.Context {
	result, e := c.ConfigurationContexts()
	errors.PanicOnError(e)

	return result
}
