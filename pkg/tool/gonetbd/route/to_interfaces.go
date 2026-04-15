package route

import (
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
)

func toInterfaces(interfaces []*network.Interface) []server.Interface {
	result := make([]server.Interface, 0, len(interfaces))

	for _, i := range interfaces {
		entry := server.Interface{
			Identifier: i.Identifier,
			Name:       i.Name,
		}

		if i.Description != "" {
			entry.Description = &i.Description
		}

		if i.Type != "" {
			s := string(i.Type)
			entry.Type = &s
		}

		if len(i.PhysicalAddress) > 0 {
			s := i.PhysicalAddress.String()
			entry.PhysicalAddress = &s
		}

		result = append(result, entry)
	}

	return result
}
