package user

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.User) *User {
	return &User{Identifier: v.GetId(), Name: v.GetDisplay(), Raw: v}
}
