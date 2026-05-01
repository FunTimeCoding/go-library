package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/platform"
)

func (c *Client) MustPlatforms() []*platform.Platform {
	result, e := c.Platforms()
	errors.PanicOnError(e)

	return result
}
