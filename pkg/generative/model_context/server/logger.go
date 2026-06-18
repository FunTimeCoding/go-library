package server

import "log/slog"

func (s *Server) Logger() *slog.Logger {
	if s.logger != nil {
		return s.logger
	}

	s.logger = slog.Default()

	return s.logger
}
