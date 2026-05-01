package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListTenants(
	w http.ResponseWriter,
	_ *http.Request,
) {
	tenants := s.client.MustTenants()
	result := make([]server.Tenant, 0, len(tenants))

	for _, t := range tenants {
		result = append(
			result,
			server.Tenant{
				Identifier: t.Identifier, Name: t.Name,
			},
		)
	}

	web.EncodeNotation(w, result)
}
