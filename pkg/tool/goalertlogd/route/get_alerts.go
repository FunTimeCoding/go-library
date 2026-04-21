package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) GetAlerts(
	w http.ResponseWriter,
	_ *http.Request,
	params server.GetAlertsParams,
) {
	web.EncodeNotation(w, toResponse(h.store.ByName(params.Name)))
}
