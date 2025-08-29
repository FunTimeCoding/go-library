package platform

import "github.com/netbox-community/go-netbox/v4"

type Platform struct {
	Identifier int32
	Name       string

	Raw *netbox.Platform
}
