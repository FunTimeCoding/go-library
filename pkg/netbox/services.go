package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/service"
)

func (c *Client) Services() []*service.Service {
	result, r, e := c.client.IpamAPI.IpamServicesList(
		c.context,
	).Limit(constant.PageLimit).Execute()
	errors.PanicOnWebError(r, e)

	return service.NewSlice(result.Results)
}
