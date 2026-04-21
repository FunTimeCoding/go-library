package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListDevices(
	w http.ResponseWriter,
	_ *http.Request,
	params server.ListDevicesParams,
) {
	var devices []server.Device

	if params.Query != nil && *params.Query != "" {
		devices = toDevices(h.client.DevicesByMatch(*params.Query))
	} else {
		devices = toDevices(h.client.Devices())
	}

	web.EncodeNotation(w, devices)
}
