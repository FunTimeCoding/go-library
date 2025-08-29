package user

import "github.com/netbox-community/go-netbox/v4"

type User struct {
	Identifier int32
	Name       string

	Raw *netbox.User
}
