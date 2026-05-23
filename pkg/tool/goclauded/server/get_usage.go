package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetUsage(
	w http.ResponseWriter,
	_ *http.Request,
) {
	result := s.service.Usage()

	if result == nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	web.EncodeNotation(
		w,
		server.UsageResponse{
			SessionPercent:      result.SessionPercent,
			SessionReset:        result.SessionReset,
			WeeklyAllPercent:    result.WeeklyAllPercent,
			WeeklyAllReset:      result.WeeklyAllReset,
			WeeklySonnetPercent: result.WeeklySonnetPercent,
			WeeklyDesignPercent: result.WeeklyDesignPercent,
			RoutineRuns:         result.RoutineRuns,
			CreditSpent:         result.CreditSpent,
			CreditReset:         result.CreditReset,
			CreditPercent:       result.CreditPercent,
			LastUpdated:         result.LastUpdated,
		},
	)
}
