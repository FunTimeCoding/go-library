package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) complete(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.Complete)

	if c.Callsign == "" {
		return response.Fail("unknown session - announce first to bind your identity")
	}

	message, e := q.RequireString(constant.Message)

	if e != nil {
		return response.Fail("message is required: %v", e)
	}

	topic := s.service.CompleteTask(c.Callsign)

	if topic == "" {
		topic = q.GetString(constant.Topic, "")
	}

	if topic == "" {
		return response.Fail("no active topic - provide a topic parameter or announce first")
	}

	s.service.Complete(c.SessionIdentifier, c.Callsign, topic, message)

	return response.Success(
		fmt.Sprintf("Completed %s: %s", topic, message),
	)
}
