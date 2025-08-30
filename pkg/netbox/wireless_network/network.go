package wireless_network

import "github.com/netbox-community/go-netbox/v4"

type Network struct {
	Identifier int32
	Name       string

	Raw *netbox.WirelessLAN
}
