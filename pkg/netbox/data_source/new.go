package data_source

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.DataSource) *Source {
	return &Source{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
