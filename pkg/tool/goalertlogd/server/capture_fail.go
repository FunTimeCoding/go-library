package server

import "github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/server"

func (s *Server) captureFail(
	e error,
	message string,
) *server.ErrorResponse {
	return &server.ErrorResponse{
		Error:           message,
		EventIdentifier: s.reporter.CaptureException(e),
	}
}
