package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListDevices(
	w http.ResponseWriter,
	_ *http.Request,
	v server.ListDevicesParams,
) {
	var devices []server.Device

	if v.Query != nil && *v.Query != "" {
		devices = toDevices(s.client.DevicesByMatch(*v.Query))
	} else {
		devices = toDevices(s.client.Devices())
	}

	web.EncodeNotation(w, devices)
}
