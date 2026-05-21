package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostSessionExport(
	w http.ResponseWriter,
	_ *http.Request,
	identifier string,
) {
	session := s.service.Resolve(identifier)

	if session.Identifier == "" {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	web.EncodeNotation(
		w,
		server.ExportResponse{
			Paths: []string{
				s.exportSession(
					session,
					s.sessionExportPath,
				),
			},
		},
	)
}
