package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"net"
)

func (c *Client) PhysicalAddress(a net.HardwareAddr) (*physical_address.Address, error) {
	result, e := c.PhysicalAddressesByHardware(a)

	if e != nil {
		return nil, e
	}

	if o := len(result); o > 1 {
		for _, r := range result {
			if r.Address.String() == a.String() {
				return r, nil
			}
		}
	} else if o == 1 {
		return result[0], nil
	}

	return nil, nil
}
