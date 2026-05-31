package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostBackfill(
	w http.ResponseWriter,
	r *http.Request,
) {
	result := s.service.BackfillAllSessions()
	web.EncodeNotation(
		w,
		server.BackfillResponse{
			Enriched: result.Enriched,
			Skipped:  result.Skipped,
		},
	)
}
