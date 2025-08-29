package user

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.User) []*User {
	var result []*User

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
