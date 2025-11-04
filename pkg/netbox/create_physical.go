package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) CreatePhysical(
	a net.HardwareAddr,
	description string,
) *physical_address.Address {
	q := netbox.NewMACAddressRequest(a.String())

	if description != "" {
		q.SetDescription(description)
	}

	result, r, e := c.client.DcimAPI.DcimMacAddressesCreate(
		c.context,
	).MACAddressRequest(*q).Execute()
	errors.PanicOnWebError(r, e)

	return physical_address.New(result)
}
