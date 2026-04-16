package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
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

	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(result))
}
