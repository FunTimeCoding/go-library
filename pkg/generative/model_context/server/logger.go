package server

import "github.com/mark3labs/mcp-go/util"

func (s *Server) Logger() util.Logger {
	if s.logger != nil {
		return s.logger
	}

	s.logger = util.DefaultLogger()

	return s.logger
}
