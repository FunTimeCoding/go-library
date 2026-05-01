package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListManufacturers(
	w http.ResponseWriter,
	_ *http.Request,
) {
	manufacturers := s.client.MustManufacturers()
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
