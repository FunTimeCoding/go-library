package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/device_role"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
)

func DeviceRoles(v []*device_role.Role) []server.DeviceRole {
	result := make([]server.DeviceRole, 0, len(v))

	for _, r := range v {
		result = append(result, DeviceRole(r))
	}

	return result
}
