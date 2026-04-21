package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListManufacturers(
	w http.ResponseWriter,
	_ *http.Request,
) {
	manufacturers := h.client.Manufacturers()
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
