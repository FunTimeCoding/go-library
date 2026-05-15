package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListAddresses(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
) {
	addresses, e := s.client.DeviceAddresses(name)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.EncodeNotation(w, convert.Addresses(addresses))
}
