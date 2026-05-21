package model_context

import (
	"context"
	telemetry "github.com/funtimecoding/go-library/pkg/telemetry/constant"
	"github.com/funtimecoding/go-library/pkg/telemetry/record"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) resolveCaller(
	x context.Context,
	tool string,
) *caller {
	session := server.ClientSessionFromContext(x)

	if session == nil {
		return &caller{}
	}

	modelContextSessionIdentifier := session.SessionID()
	name, sessionIdentifier := s.service.ResolveModelContextSession(modelContextSessionIdentifier)
	s.logger.Structured(
		"model_context_tool_call",
		"tool",
		tool,
		"model_context_session_identifier",
		modelContextSessionIdentifier,
		constant.SessionName,
		name,
	)
	s.telemetry.Record(
		record.New(
			tool,
			telemetry.ModelContext,
			name,
			telemetry.Success,
		),
	)

	return &caller{Callsign: name, SessionIdentifier: sessionIdentifier}
}
