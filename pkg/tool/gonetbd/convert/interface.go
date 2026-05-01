package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/network"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func Interface(i *network.Interface) server.Interface {
	result := server.Interface{Identifier: i.Identifier, Name: i.Name}

	if i.Description != "" {
		result.Description = &i.Description
	}

	if i.Type != "" {
		t := string(i.Type)
		result.Type = &t
	}

	if len(i.PhysicalAddress) > 0 {
		p := i.PhysicalAddress.String()
		result.PhysicalAddress = &p
	}

	return result
}
