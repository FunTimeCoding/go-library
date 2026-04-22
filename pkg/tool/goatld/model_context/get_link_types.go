package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

type linkType struct {
	Name    string `json:"name"`
	Inward  string `json:"inward"`
	Outward string `json:"outward"`
}

type linkTypeResponse struct {
	IssueLinkTypes []linkType `json:"issueLinkTypes"`
}

func (s *Server) getLinkTypes(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	_, body := s.jira.Basic().GetPath(
		"rest/api/2/issueLinkType",
	)
	var parsed linkTypeResponse
	notation.DecodeStrict(body, &parsed, true)

	return response.SuccessAny(parsed.IssueLinkTypes)
}
