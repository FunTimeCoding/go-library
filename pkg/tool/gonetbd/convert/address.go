package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func Address(a *internet_address.Address) server.Address {
	result := server.Address{
		Identifier: a.Identifier,
		Address:    a.Address.String(),
	}

	if a.Name != "" {
		result.Name = &a.Name
	}

	if a.ObjectType != "" {
		result.ObjectType = &a.ObjectType
	}

	return result
}
