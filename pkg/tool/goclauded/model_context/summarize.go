package model_context

import (
	"context"
	library "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) summarize(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	c := s.resolveCaller(x, constant.Summarize)

	if c.Callsign == "" {
		return response.Fail("unknown session - announce first to bind your identity")
	}

	body, e := q.RequireString(constant.Body)

	if e != nil {
		return response.Fail("body is required: %v", e)
	}

	if f := s.service.Summarize(c.SessionIdentifier, c.Callsign, body); f != nil {
		return s.captureFail(f, library.UnexpectedError)
	}

	return response.Success("Session summary recorded for %s", c.Callsign)
}
