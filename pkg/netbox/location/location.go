package location

import "github.com/netbox-community/go-netbox/v4"

type Location struct {
	Identifier int32
	Name       string

	Raw *netbox.Location
}
