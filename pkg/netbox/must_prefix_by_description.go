package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/prefix"
)

func (c *Client) MustPrefixByDescription(d string) *prefix.Prefix {
	result, e := c.PrefixByDescription(d)
	errors.PanicOnError(e)

	return result
}
