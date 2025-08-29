package user

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.User) *User {
	return &User{Identifier: d.GetId(), Name: d.GetDisplay(), Raw: d}
}
