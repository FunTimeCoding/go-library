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
	if v.Query != nil && *v.Query != "" {
		devices, e := s.client.DevicesByMatch(*v.Query)

		if e != nil {
			s.captureDetail(w, e)

			return
		}

		web.EncodeNotation(w, convert.Devices(devices))

		return
	}

	devices, e := s.client.Devices()

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.EncodeNotation(w, convert.Devices(devices))
}
