package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) push(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	collection, e := q.RequireString(constant.Collection)

	if e != nil {
		return response.Fail("collection is required: %v", e)
	}

	path, f := q.RequireString(constant.Path)

	if f != nil {
		return response.Fail("path is required: %v", f)
	}

	body, g := q.RequireString(constant.Body)

	if g != nil {
		return response.Fail("body is required: %v", g)
	}

	sourceType := q.GetString(constant.SourceType, "")

	if h := s.store.PushDocument(
		collection,
		path,
		body,
		sourceType,
		s.ollama,
	); h != nil {
		return s.captureFail(h, "push failed")
	}

	return response.Success("pushed %s/%s", collection, path)
}
