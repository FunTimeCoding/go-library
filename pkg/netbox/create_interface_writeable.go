package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) createInterfaceWriteable(q *netbox.WritableInterfaceRequest) *network.Interface {
	result, r, e := c.client.DcimAPI.DcimInterfacesCreate(
		c.context,
	).WritableInterfaceRequest(*q).Execute()
	errors.PanicOnWebError(r, e)

	return network.New(result)
}
