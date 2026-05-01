package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/prefix"
)

func (c *Client) MustPrefixByName(n string) *prefix.Prefix {
	result, e := c.PrefixByName(n)
	errors.PanicOnError(e)

	return result
}
