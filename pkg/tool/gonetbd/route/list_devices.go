package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) ListDevices(
	w http.ResponseWriter,
	_ *http.Request,
	params server.ListDevicesParams,
) {
	var devices []server.Device

	if params.Query != nil && *params.Query != "" {
		devices = toDevices(h.client.DevicesByMatch(*params.Query))
	} else {
		devices = toDevices(h.client.Devices())
	}

	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(devices))
}
