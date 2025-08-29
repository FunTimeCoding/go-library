package virtual_device_context

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.VirtualDeviceContext) *Context {
	return &Context{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
