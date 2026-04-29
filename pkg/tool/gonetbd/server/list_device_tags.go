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
	web.EncodeNotation(w, s.client.DeviceTagNames(name))
}
