package device_bay_template

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.DeviceBayTemplate) *Template {
	return &Template{Identifier: d.GetId(), Model: d.GetName(), Raw: d}
}
