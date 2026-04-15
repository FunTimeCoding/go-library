package route

import (
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
)

func toDevices(devices []*device.Device) []server.Device {
	result := make([]server.Device, 0, len(devices))

	for _, d := range devices {
		result = append(result, toDevice(d))
	}

	return result
}
