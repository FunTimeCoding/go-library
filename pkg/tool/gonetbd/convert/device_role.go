package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/device_role"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func DeviceRole(r *device_role.Role) server.DeviceRole {
	return server.DeviceRole{Identifier: r.Identifier, Name: r.Name}
}
