package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"time"
)

func (h *Router) GetTopAlerts(
	w http.ResponseWriter,
	_ *http.Request,
	params server.GetTopAlertsParams,
) {
	n := 25
	now := time.Now()
	start := now.Add(-7 * 24 * time.Hour)
	end := now

	if params.N != nil {
		n = *params.N
	}

	if params.Start != nil {
		start = *params.Start
	}

	if params.End != nil {
		end = *params.End
	}

	records := h.store.Top(n, start, end)
	result := make([]server.TopAlertsResponse, 0, len(records))

	for _, r := range records {
		result = append(
			result,
			server.TopAlertsResponse{
				Name:            r.Name,
				Count:           r.Count,
				AverageDuration: r.AverageDuration.String(),
				CurrentlyFiring: r.CurrentlyFiring,
				Severity:        r.Severity,
			},
		)
	}

	web.EncodeNotation(w, result)
}
