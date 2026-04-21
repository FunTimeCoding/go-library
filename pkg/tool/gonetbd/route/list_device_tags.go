package route

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListDeviceTags(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
) {
	web.EncodeNotation(w, h.client.DeviceTagNames(name))
}
