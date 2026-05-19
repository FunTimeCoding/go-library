package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/sweep"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostSweep(
	w http.ResponseWriter,
	r *http.Request,
) {
	result := sweep.Run(s.harborPath)
	web.EncodeNotation(
		w,
		server.SweepResponse{
			Copied:  result.Copied,
			Updated: result.Updated,
			Skipped: result.Skipped,
		},
	)
}
