package model_context

import (
	"context"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) activeInstance(x context.Context) (string, bool) {
	session := server.ClientSessionFromContext(x)

	if session == nil {
		return "", false
	}

	return s.store.ActiveInstance(session.SessionID())
}
