package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"time"
)

func (s *Server) GetRecentAlerts(
	w http.ResponseWriter,
	_ *http.Request,
	v server.GetRecentAlertsParams,
) {
	now := time.Now()
	start := now.Add(-1 * time.Hour)
	end := now

	if v.Start != nil {
		start = *v.Start
	}

	if v.End != nil {
		end = *v.End
	}

	web.EncodeNotation(w, toResponse(s.store.ByTimeRange(start, end)))
}
