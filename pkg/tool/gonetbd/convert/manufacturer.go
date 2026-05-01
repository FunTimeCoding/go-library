package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/manufacturer"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func Manufacturer(m *manufacturer.Manufacturer) server.Manufacturer {
	return server.Manufacturer{Identifier: m.Identifier, Name: m.Name}
}
