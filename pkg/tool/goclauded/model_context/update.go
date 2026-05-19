package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

func (s *Server) update(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.Update)

	if c.Name == "" {
		return response.Fail("unknown session - announce first to bind your identity")
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

	s.service.Store.UpdateTopic(c.Name, topic, strings.Join(files, "\n"))
	s.service.Store.CreateCompletion(
		c.SessionIdentifier,
		c.Name,
		constant.Update,
		topic,
		"",
	)
	s.service.Store.LogEvent(
		c.SessionIdentifier,
		constant.Update,
		c.Name,
		"",
		topic,
	)

	return response.Success(
		fmt.Sprintf("Updated scope: %s", topic),
	)
}
