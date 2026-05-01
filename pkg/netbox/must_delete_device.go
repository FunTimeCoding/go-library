package netbox

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDeleteDevice(identifier int32) {
	errors.PanicOnError(c.DeleteDevice(identifier))
}
