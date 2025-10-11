package jenkins

import (
	"github.com/bndr/gojenkins"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Plugins() []gojenkins.Plugin {
	r, e := c.client.GetPlugins(c.context, 2)
	errors.PanicOnError(e)

	return r.Raw.Plugins
}
