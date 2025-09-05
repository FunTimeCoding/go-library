package netbox

import (
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"net"
)

func (c *Client) PhysicalAddressStrict(a net.HardwareAddr) *physical_address.Address {
	result := c.PhysicalAddressesByHardware(a)

	if o := len(result); o > 1 {
		for _, r := range result {
			if r.Address.String() == a.String() {
				return r
			}
		}
	}

	unexpected.Count(1, len(result))

	return result[0]
}
