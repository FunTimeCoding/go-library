package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/virtual_machine"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func VirtualMachine(m *virtual_machine.Machine) server.VirtualMachine {
	result := server.VirtualMachine{Identifier: m.Identifier, Name: m.Name}

	if m.Raw.Cluster.IsSet() && m.Raw.Cluster.Get() != nil {
		n := m.Raw.Cluster.Get().GetName()
		result.Cluster = &n
	}

	if m.Raw.Site.IsSet() && m.Raw.Site.Get() != nil {
		n := m.Raw.Site.Get().GetName()
		result.Site = &n
	}

	if len(m.Tags) > 0 {
		result.Tags = &m.Tags
	}

	return result
}
