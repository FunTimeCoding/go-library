package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"time"
)

func (h *Router) GetRecentAlerts(
	w http.ResponseWriter,
	_ *http.Request,
	params server.GetRecentAlertsParams,
) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now

	if params.Start != nil {
		start = *params.Start
	}

	if params.End != nil {
		end = *params.End
	}

	web.EncodeNotation(w, toResponse(h.store.ByTimeRange(start, end)))
}
