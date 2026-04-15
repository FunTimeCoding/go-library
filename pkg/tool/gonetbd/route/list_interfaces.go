package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) ListInterfaces(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
) {
	d := h.client.DeviceByNameStrict(name)
	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(
		json.NewEncoder(w).Encode(toInterfaces(h.client.DeviceInterfaces(d.Identifier))),
	)
}
