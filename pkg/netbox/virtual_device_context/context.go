package virtual_device_context

import "github.com/netbox-community/go-netbox/v4"

type Context struct {
	Identifier int32
	Name       string

	Raw *netbox.VirtualDeviceContext
}
