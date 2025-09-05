package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"net"
)

func (c *Client) PhysicalAddressesByHardware(a net.HardwareAddr) []*physical_address.Address {
	result, _, e := c.client.DcimAPI.DcimMacAddressesList(
		c.context,
	).MacAddress([]string{a.String()}).Execute()
	errors.PanicOnError(e)

	return physical_address.NewSlice(result.Results)
}
