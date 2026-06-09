package server

import "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"

func (s *Server) captureFail(
	e error,
	message string,
) *server.ErrorJSONResponse {
	return &server.ErrorJSONResponse{
		Error:           message,
		EventIdentifier: s.reporter.CaptureException(e),
	}
}
