package route

import (
	"github.com/funtimecoding/go-library/pkg/netbox/internet_address"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
)

func toAddresses(addresses []*internet_address.Address) []server.Address {
	result := make([]server.Address, 0, len(addresses))

	for _, a := range addresses {
		entry := server.Address{
			Identifier: a.Identifier,
			Address:    a.Address.String(),
		}

		if a.Name != "" {
			entry.Name = &a.Name
		}

		if a.ObjectType != "" {
			entry.ObjectType = &a.ObjectType
		}

		result = append(result, entry)
	}

	return result
}
