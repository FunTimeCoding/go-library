package service

import "github.com/netbox-community/go-netbox/v4"

func NewSlice(v []netbox.Service) []*Service {
	var result []*Service

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
