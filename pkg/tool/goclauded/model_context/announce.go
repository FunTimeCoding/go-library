package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func (s *Server) announce(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := q.RequireString(constant.SessionName)

	if e != nil {
		return response.Fail("name is required: %v", e)
	}

	topic, e := q.RequireString(constant.Topic)

	if e != nil {
		return response.Fail("topic is required: %v", e)
	}

	var files []string

	if raw, okay := q.GetArguments()[constant.Files]; okay {
		if list, okay := raw.([]any); okay {
			for _, item := range list {
				if str, okay := item.(string); okay {
					files = append(files, str)
				}
			}
		}
	}

	if cs := server.ClientSessionFromContext(x); cs != nil {
		s.service.BindModelContextSession(name, cs.SessionID())
	}

	c := s.resolveCaller(x, constant.Announce)
	s.service.Announce(
		c.SessionIdentifier,
		name,
		topic,
		join.NewLine(files),
	)

	return response.Success(
		fmt.Sprintf(
			"Announced as %s: %s",
			name,
			topic,
		),
	)
}
