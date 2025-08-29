package netbox

import (
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
	"net"
)

func (c *Client) FindUnusedAddress(sub *net.IPNet) net.IP {
	return internet_address.FindUnused(
		internet_address.FilterSubnet(c.InternetAddresses(), sub),
		sub,
	)
}
