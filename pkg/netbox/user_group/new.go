package user_group

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.Group) *Group {
	return &Group{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
