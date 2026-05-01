package netbox

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDeleteInterface(identifier int32) {
	errors.PanicOnError(c.DeleteInterface(identifier))
}
