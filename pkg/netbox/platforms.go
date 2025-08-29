package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/platform"
)

func (c *Client) Platforms() []*platform.Platform {
	result, _, e := c.client.DcimAPI.DcimPlatformsList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnError(e)

	return platform.NewSlice(result.Results)
}
