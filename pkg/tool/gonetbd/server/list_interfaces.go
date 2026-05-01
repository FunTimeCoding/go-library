package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListInterfaces(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
) {
	d := s.client.MustDeviceByNameStrict(name)
	web.EncodeNotation(
		w,
		convert.Interfaces(s.client.MustDeviceInterfaces(d.Identifier)),
	)
}
