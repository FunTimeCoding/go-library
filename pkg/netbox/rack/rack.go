package rack

import "github.com/netbox-community/go-netbox/v4"

type Rack struct {
	Identifier int32
	Name       string
	Link       string

	Raw *netbox.Rack
}
