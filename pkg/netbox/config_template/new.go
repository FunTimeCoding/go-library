package config_template

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.ConfigTemplate) *Template {
	return &Template{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
