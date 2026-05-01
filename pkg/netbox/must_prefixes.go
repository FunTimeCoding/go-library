package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/prefix"
)

func (c *Client) MustPrefixes() []*prefix.Prefix {
	result, e := c.Prefixes()
	errors.PanicOnError(e)

	return result
}
