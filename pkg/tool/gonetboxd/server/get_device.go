package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetDevice(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
) {
	d, e := s.client.DeviceByNameStrict(name)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.EncodeNotation(w, convert.Device(d))
}
