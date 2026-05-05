package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListDeviceRoles(
	w http.ResponseWriter,
	_ *http.Request,
) {
	roles := s.client.MustDeviceRoles()
	result := make([]server.DeviceRole, 0, len(roles))

	for _, r := range roles {
		result = append(
			result,
			server.DeviceRole{Identifier: r.Identifier, Name: r.Name},
		)
	}

	web.EncodeNotation(w, result)
}
