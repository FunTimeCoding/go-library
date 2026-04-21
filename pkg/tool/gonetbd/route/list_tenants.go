package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListTenants(
	w http.ResponseWriter,
	_ *http.Request,
) {
	tenants := h.client.Tenants()
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
