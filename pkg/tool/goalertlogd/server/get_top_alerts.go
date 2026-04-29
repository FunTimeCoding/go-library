package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"time"
)

func (s *Server) GetTopAlerts(
	w http.ResponseWriter,
	_ *http.Request,
	v server.GetTopAlertsParams,
) {
	n := 25
	now := time.Now()
	start := now.Add(-7 * 24 * time.Hour)
	end := now

	if v.N != nil {
		n = *v.N
	}

	if v.Start != nil {
		start = *v.Start
	}

	if v.End != nil {
		end = *v.End
	}

	records := s.store.Top(n, start, end)
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
