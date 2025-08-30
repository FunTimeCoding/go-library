package tunnel

import "github.com/netbox-community/go-netbox/v4"

type Tunnel struct {
	Identifier int32
	Name       string

	Raw *netbox.Tunnel
}
