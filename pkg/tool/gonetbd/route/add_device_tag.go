package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) AddDeviceTag(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
	tag string,
) {
	d := h.client.AddTag(name, tag)
	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(toDevice(d)))
}
