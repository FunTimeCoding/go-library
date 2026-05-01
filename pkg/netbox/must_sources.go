package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/source"
)

func (c *Client) MustSources() []*source.Source {
	result, e := c.Sources()
	errors.PanicOnError(e)

	return result
}
