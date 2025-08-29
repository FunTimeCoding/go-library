package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
	"log"
)

func (c *Client) PhysicalAddresses() []*physical_address.Address {
	if len(c.cache.PhysicalAddresses) > 0 {
		return c.cache.PhysicalAddresses
	}

	var result []netbox.MACAddress
	var offset int32

	for {
		page := c.physicalAddressesOffset(offset)
		result = append(result, page...)

		if len(page) < int(constant.PageLimit) {
			break
		}

		offset += constant.PageLimit
	}

	if len(result) == 0 {
		log.Panicf("no addresses found")
	}

	c.cache.PhysicalAddresses = physical_address.NewSlice(result)

	return c.cache.PhysicalAddresses
}
