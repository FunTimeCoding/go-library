package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
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
		devices = convert.Devices(s.client.MustDevicesByMatch(*v.Query))
	} else {
		devices = convert.Devices(s.client.MustDevices())
	}

	web.EncodeNotation(w, devices)
}
