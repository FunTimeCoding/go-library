package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) CreateDeviceType(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body generated.CreateDeviceTypeRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	m := h.client.ManufacturerByName(body.Manufacturer)
	t := h.client.CreateDeviceType(body.Model, m)
	w.Header().Set(constant.ContentType, constant.Object)
	w.WriteHeader(http.StatusCreated)
	errors.PanicOnError(
		json.NewEncoder(w).Encode(generated.DeviceType{
			Identifier: t.Identifier, Model: t.Model, Manufacturer: &m.Name,
		}),
	)
}
