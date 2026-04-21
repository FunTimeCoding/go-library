package route

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListAddresses(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
) {
	web.EncodeNotation(w, toAddresses(h.client.DeviceAddresses(name)))
}
