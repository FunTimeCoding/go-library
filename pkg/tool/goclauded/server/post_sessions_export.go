package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostSessionsExport(
	w http.ResponseWriter,
	_ *http.Request,
) {
	basePath := s.sessionExportPath
	var paths []string

	for _, session := range s.service.Sessions() {
		path := s.exportSession(session, basePath)
		paths = append(paths, path)
	}

	web.EncodeNotation(
		w,
		server.ExportResponse{Paths: paths},
	)
}
