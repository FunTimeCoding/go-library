package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListManufacturers(
	w http.ResponseWriter,
	_ *http.Request,
) {
	manufacturers, e := s.client.Manufacturers()

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	result := make([]server.Manufacturer, 0, len(manufacturers))

	for _, m := range manufacturers {
		result = append(
			result,
			server.Manufacturer{
				Identifier: m.Identifier, Name: m.Name,
			},
		)
	}

	web.EncodeNotation(w, result)
}
