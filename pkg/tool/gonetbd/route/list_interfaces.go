package route

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListInterfaces(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
) {
	d := h.client.DeviceByNameStrict(name)
	web.EncodeNotation(
		w,
		toInterfaces(h.client.DeviceInterfaces(d.Identifier)),
	)
}
