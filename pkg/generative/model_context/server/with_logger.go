package server

import "log/slog"

func WithLogger(l *slog.Logger) Option {
	return func(s *Server) {
		s.logger = l
	}
}
