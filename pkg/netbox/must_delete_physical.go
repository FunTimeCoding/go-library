package netbox

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDeletePhysical(identifier int32) {
	errors.PanicOnError(c.DeletePhysical(identifier))
}
