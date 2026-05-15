package server

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListDeviceTags(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
) {
	tags, e := s.client.DeviceTagNames(name)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.EncodeNotation(w, tags)
}
