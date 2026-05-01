package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"net"
)

func (c *Client) PhysicalAddressStrict(a net.HardwareAddr) (*physical_address.Address, error) {
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
	}

	if len(result) != 1 {
		return nil, fmt.Errorf(
			"expected 1 physical address %s, got %d",
			a,
			len(result),
		)
	}

	return result[0], nil
}
