package virtual_chassis

import "github.com/netbox-community/go-netbox/v4"

type Chassis struct {
	Identifier int32
	Name       string

	Raw *netbox.VirtualChassis
}
