package virtual_network_group

import "github.com/netbox-community/go-netbox/v4"

type Group struct {
	Identifier int32
	Name       string

	Raw *netbox.VLANGroup
}
