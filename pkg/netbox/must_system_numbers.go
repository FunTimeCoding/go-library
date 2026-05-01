package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/system_number"
)

func (c *Client) MustSystemNumbers() []*system_number.Number {
	result, e := c.SystemNumbers()
	errors.PanicOnError(e)

	return result
}
