package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/device_type"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func DeviceType(t *device_type.Type) server.DeviceType {
	result := server.DeviceType{
		Identifier: t.Identifier,
		Model:      t.Model,
	}

	if t.Raw.Manufacturer.Name != "" {
		result.Manufacturer = &t.Raw.Manufacturer.Name
	}

	return result
}
