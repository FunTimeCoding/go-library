package server

import (
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func toInterfaces(v []*network.Interface) []server.Interface {
	result := make([]server.Interface, 0, len(v))

	for _, i := range v {
		e := server.Interface{Identifier: i.Identifier, Name: i.Name}

		if i.Description != "" {
			e.Description = &i.Description
		}

		if i.Type != "" {
			e.Type = new(string(i.Type))
		}

		if len(i.PhysicalAddress) > 0 {
			e.PhysicalAddress = new(i.PhysicalAddress.String())
		}

		result = append(result, e)
	}

	return result
}
