package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createPage(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	space, e := r.RequireString(constant.SpaceIdentifier)

	if e != nil {
		return response.Fail("space_identifier is required: %v", e)
	}

	parent, f := r.RequireString(constant.ParentIdentifier)

	if f != nil {
		return response.Fail("parent_identifier is required: %v", f)
	}

	title, g := r.RequireString(parameter.Title)

	if g != nil {
		return response.Fail("title is required: %v", g)
	}

	body, h := r.RequireString(parameter.Body)

	if h != nil {
		return response.Fail("body is required: %v", h)
	}

	result, i := s.confluence.CreatePage(space, parent, title, body)

	if i != nil {
		return s.captureFail(i, "page not created")
	}

	return response.SuccessAny(result)
}
