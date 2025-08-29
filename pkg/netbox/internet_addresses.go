package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
	"github.com/netbox-community/go-netbox/v4"
	"log"
)

func (c *Client) InternetAddresses() []*internet_address.Address {
	if len(c.cache.InternetAddresses) > 0 {
		return c.cache.InternetAddresses
	}

	var result []netbox.IPAddress
	var offset int32

	for {
		page := c.internetAddressesOffset(offset)
		result = append(result, page...)

		if len(page) < int(constant.PageLimit) {
			break
		}

		offset += constant.PageLimit
	}

	if len(result) == 0 {
		log.Panicf("no addresses found")
	}

	c.cache.InternetAddresses = internet_address.NewSlice(result)

	return c.cache.InternetAddresses
}
