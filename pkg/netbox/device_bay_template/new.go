package device_bay_template

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.DeviceBayTemplate) *Template {
	return &Template{Identifier: v.GetId(), Model: v.GetName(), Raw: v}
}
