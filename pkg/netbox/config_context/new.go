package config_context

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.ConfigContext) *Context {
	return &Context{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
