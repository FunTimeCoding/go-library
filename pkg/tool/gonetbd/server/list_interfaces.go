package server

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListInterfaces(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
) {
	d := s.client.DeviceByNameStrict(name)
	web.EncodeNotation(
		w,
		toInterfaces(s.client.DeviceInterfaces(d.Identifier)),
	)
}
