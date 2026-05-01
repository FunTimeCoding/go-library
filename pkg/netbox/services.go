package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/service"
)

func (c *Client) Services() ([]*service.Service, error) {
	result, _, e := c.client.IpamAPI.IpamServicesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return service.NewSlice(result.Results), nil
}
