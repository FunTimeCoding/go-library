package server

import "github.com/mark3labs/mcp-go/util"

func WithLogger(l util.Logger) Option {
	return func(s *Server) {
		s.logger = l
	}
}
