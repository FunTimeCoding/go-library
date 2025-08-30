package service_template

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.ServiceTemplate) *Template {
	return &Template{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
