package server

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) AddDeviceTag(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
	tag string,
) {
	d := s.client.AddTag(name, tag)
	web.EncodeNotation(w, toDevice(d))
}
