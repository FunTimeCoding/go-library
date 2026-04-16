package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
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
		})
	}

	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(result))
}
