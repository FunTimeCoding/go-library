package server

import (
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func toAddresses(v []*internet_address.Address) []server.Address {
	result := make([]server.Address, 0, len(v))

	for _, a := range v {
		e := server.Address{
			Identifier: a.Identifier,
			Address:    a.Address.String(),
		}

		if a.Name != "" {
			e.Name = &a.Name
		}

		if a.ObjectType != "" {
			e.ObjectType = &a.ObjectType
		}

		result = append(result, e)
	}

	return result
}
