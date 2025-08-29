package service_template

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.ServiceTemplate) *Template {
	return &Template{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
