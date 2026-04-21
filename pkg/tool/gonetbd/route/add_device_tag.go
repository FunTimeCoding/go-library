package route

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) AddDeviceTag(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
	tag string,
) {
	d := h.client.AddTag(name, tag)
	web.EncodeNotation(w, toDevice(d))
}
