package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListDeviceRoles(
	w http.ResponseWriter,
	_ *http.Request,
) {
	roles := h.client.DeviceRoles()
	result := make([]server.DeviceRole, 0, len(roles))

	for _, r := range roles {
		result = append(
			result,
			server.DeviceRole{
				Identifier: r.Identifier, Name: r.Name,
			},
		)
	}

	web.EncodeNotation(w, result)
}
