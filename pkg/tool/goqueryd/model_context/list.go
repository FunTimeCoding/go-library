package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) list(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	collection, e := q.RequireString(constant.Collection)

	if e != nil {
		return response.Fail("collection is required: %v", e)
	}

	documents, f := s.store.ListDocuments(collection)

	if f != nil {
		return s.captureFail(f, "document listing failed")
	}

	return response.Success(notation.MarshalIndent(documents))
}
