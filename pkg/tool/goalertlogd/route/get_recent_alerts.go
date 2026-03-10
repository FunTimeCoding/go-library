package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
	"time"
)

func (h *Handler) GetRecentAlerts(
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

	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(
		json.NewEncoder(w).Encode(
			toResponse(h.store.ByTimeRange(start, end)),
		),
	)
}
