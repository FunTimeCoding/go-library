package server

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) RemoveDeviceTag(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
	tag string,
) {
	d := s.client.RemoveTag(name, tag)
	web.EncodeNotation(w, toDevice(d))
}
